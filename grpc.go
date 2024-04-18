package header

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/subiz/header/common"
	"github.com/subiz/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const (
	CtxKey = "pcontext"
)

// WithShardRedirect creates a dial option that learns on "shard_addrs" reponse
// header to send requests to correct shard
// see https://www.notion.so/Shard-service-c002bcb0b00c47669bce547be646cd9f
// for the overall design
func WithShardRedirect() grpc.DialOption {
	lock := &sync.Mutex{}
	// GRPC connections to all shard workers, mapping host (user-3.user:2000) to connection
	conn := make(map[string]*grpc.ClientConn)
	// list of current shard worker addresses (order is important)
	addrs := []string{}

	f := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// has data learned from last request
		if len(addrs) > 0 {
			// looking for shard key in header or account_id field in parameter
			md, _ := metadata.FromOutgoingContext(ctx)
			pkey := strings.Join(md["shard_key"], "")
			if pkey == "" {
				pkey = GetAccountId(ctx, req)
			}
			if pkey != "" {
				// finding the shard number
				shardNumber := int(crc32.ChecksumIEEE([]byte(pkey))) % len(addrs)
				host := addrs[shardNumber]
				co, ok := conn[host]
				if !ok {
					lock.Lock()
					if co, ok = conn[host]; !ok {
						co = DialGrpc(host)
						// copy on write
						copyConn := make(map[string]*grpc.ClientConn)
						for k, v := range conn {
							copyConn[k] = v
						}
						copyConn[host] = co
						conn = copyConn
					}
					lock.Unlock()
				}
				var header metadata.MD // variable to store header and trailer
				opts = append([]grpc.CallOption{grpc.Header(&header)}, opts...)
				err := co.Invoke(ctx, method, req, reply, opts...)
				if len(header["shard_addrs"]) > 0 {
					addrs = header.Get("shard_addrs")
				}
				return err
			}
		}

		// no sharding parameter, perform the request anyway
		var header metadata.MD
		opts = append([]grpc.CallOption{grpc.Header(&header)}, opts...)
		err := invoker(ctx, method, req, reply, cc, opts...)
		if len(header["shard_addrs"]) > 0 {
			addrs = header.Get("shard_addrs")
		}
		return err
	}
	return grpc.WithChainUnaryInterceptor(f)
}

func ToGrpcCtx(pctx *common.Context) context.Context {
	data, err := proto.Marshal(pctx)
	if err != nil {
		panic(fmt.Sprintf("unable to marshal cred, %v", pctx))
	}
	cred64 := base64.StdEncoding.EncodeToString(data)
	return metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs(CtxKey, cred64))
}

func FromGrpcCtx(ctx context.Context) *common.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md, ok = metadata.FromOutgoingContext(ctx)
		if !ok {
			return nil
		}
	}
	cred64 := strings.Join(md[CtxKey], "")
	if cred64 == "" {
		return nil
	}
	data, err := base64.StdEncoding.DecodeString(cred64)
	if err != nil {
		panic(fmt.Sprintf("%v, %s: %s", err, "wrong base64 ", cred64))
	}

	pctx := &common.Context{}
	if err = proto.Unmarshal(data, pctx); err != nil {
		panic(fmt.Sprintf("%v, %s: %s", err, "unable to unmarshal cred ", cred64))
	}
	return pctx
}

func RecoverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (ret interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = log.EServer(e, log.M{"_function_name": info.FullMethod, "__skip_stack": 1}) // wrap error
			} else {
				err = log.EServiceUnavailable(nil, log.M{"base": r, "_function_name": info.FullMethod, "__skip_stack": 1})
			}
		}
	}()
	return handler(ctx, req)
}

func WithErrorStack() grpc.DialOption {
	return grpc.WithChainUnaryInterceptor(func(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = log.EServer(err, log.M{"grpc_code": codes.Canceled.String(), "_function_name": method, "__skip_stack": "3"})
			}
		}()

		err = invoker(ctx, method, req, reply, cc, opts...)
		if err == nil {
			return nil
		}

		grpcerr, ok := status.FromError(err)
		if !ok {
			// very strange error
			return log.EServiceUnavailable(err, log.M{"grpc_code": "nono", "_function_name": method, "__skip_stack": "2"}) // report
		}

		ourerr := log.FromString(grpcerr.Message())
		if ourerr != nil {
			return log.WrapStack(ourerr, 2)
		}

		switch grpcerr.Code() {
		case codes.Unavailable:
			return log.EServiceUnavailable(err, log.M{"grpc_code": grpcerr.Code().String(), "_function_name": method, "__skip_stack": "3"}) // report
		default:
			// good grpc err but not our error
			// codes.Internal
			input, _ := json.Marshal(req)
			cred := FromGrpcCtx(ctx).GetCredential()
			return log.EServer(err, log.M{"grpc_code": grpcerr.Code().String(), "cred": cred, "input": string(input), "_function_name": method, "__skip_stack": "3"}) // report
		}
	})
}

// forward proxy a GRPC calls to another host, header and trailer are preserved
// parameters:
//
//	host: host address which will be redirected to
//	method: the full RPC method string, i.e., /package.service/method.
//	returnedType: type of returned value
//	in: value of input (in request) parameter
//
// this method returns output just like a normal GRPC call
func forward(cc *grpc.ClientConn, method string, returnedType reflect.Type, ctx context.Context, in interface{}, extraHeader metadata.MD) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	extraHeader = metadata.Join(extraHeader, md)
	outctx := metadata.NewOutgoingContext(context.Background(), extraHeader)

	out := reflect.New(returnedType).Interface()
	var header, trailer metadata.MD
	err := cc.Invoke(outctx, method, in, out, grpc.Header(&header), grpc.Trailer(&trailer))
	grpc.SendHeader(ctx, header)
	grpc.SetTrailer(ctx, trailer)

	return out, err
}

// NewStatefulSetShardInterceptor create a shard interceptor compatible with kubernetes
// statefulset, each pod is a shard, shard number is extract from pod ordinal number
func NewStatefulSetShardInterceptor(grpcport, shards int) grpc.UnaryServerInterceptor {
	hostname, _ := os.Hostname()
	sp := strings.Split(hostname, "-")
	if len(sp) < 2 {
		panic("invalid hostname" + hostname)
	}
	ordinal := sp[len(sp)-1]
	pari64, _ := strconv.ParseInt(ordinal, 10, 0)
	ordinal_num := int(pari64)

	hosts := make([]string, 0)
	for i := 0; i < shards; i++ {
		// convo-${i}.convo:{port}
		hosts = append(hosts, sp[0]+"-"+strconv.Itoa(i)+"."+sp[0]+":"+strconv.Itoa(grpcport))
	}

	return NewServerShardInterceptor(hosts, ordinal_num)
}

// NewShardInterceptor makes a GRPC server intercepter that can be used for sharding
// see https://www.notion.so/Shard-service-c002bcb0b00c47669bce547be646cd9f
// for more details about the design
func NewServerShardInterceptor(serviceAddrs []string, id int) grpc.UnaryServerInterceptor {
	// holds the current maximum number of shards
	numShard := len(serviceAddrs)

	// in order to proxy (forward) the request to another grpc host,
	// we must have an output object of the request's method (so we can marshal the response).
	// we are going to build a map of returning type for all methods of the server. And do it
	// only once time for each method name right before the first request.
	returnedTypeM := make(map[string]reflect.Type)

	// GRPC connections to shard workers
	// mapping worker address (user-2.user:8080) to a GRPC connection
	lock := &sync.Mutex{}
	conn := make(map[string]*grpc.ClientConn)

	return func(ctx context.Context, in interface{}, sinfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				if e, ok := r.(error); ok {
					err = log.EServer(e, log.M{"_function_name": sinfo.FullMethod, "__skip_stack": 1}) // wrap error
				} else {
					err = log.EServiceUnavailable(nil, log.M{"base": r, "_function_name": sinfo.FullMethod, "__skip_stack": 1})
				}
			}
		}()

		// looking for shard_key in header or in account_id parameter
		md, _ := metadata.FromIncomingContext(ctx)
		pkey := strings.Join(md["shard_key"], "")

		if pkey == "" {
			pkey = GetAccountId(ctx, in)
			if pkey == "" {
				// no sharding parameter, perform the request anyway
				return handler(ctx, in)
			}
		}

		// find the correct shard

		parindex := int(crc32.ChecksumIEEE([]byte(pkey))) % numShard

		// process if this is the correct shard
		if int(parindex) == id {
			return handler(ctx, in)
		}

		// the request have been proxied two times. We give up to prevent looping
		redirectOfRedirect := len(strings.Join(md["shard_redirected_2"], "")) > 0
		if redirectOfRedirect {
			return nil, status.Errorf(codes.Internal, "Sharding inconsistent")
		}

		// the request just have been proxied and still
		// doesn't arrived to the correct host
		// this happend when total_shards is not consistent between servers. We will wait for
		// 5 secs and then proxy one more time. Hoping that the consistent will be resolved
		justRedirect := len(strings.Join(md["shard_redirected"], "")) > 0
		extraHeader := metadata.New(nil)
		if justRedirect {
			// mark the the request have been proxied twice
			extraHeader.Set("shard_redirected_2", "true")
			time.Sleep(5 * time.Second)
		} else {
			// mark the the request have been proxied once
			extraHeader.Set("shard_redirected", "true")
		}

		header := metadata.New(nil)
		header.Set("shard_addrs", serviceAddrs...)
		grpc.SendHeader(ctx, header)

		// use cache host connection or create a new one
		host := serviceAddrs[parindex]
		lock.Lock()
		cc, ok := conn[host]

		if !ok {
			cc = DialGrpc(host)
			conn[host] = cc
		}

		// making a map of returning type for all methods of the server
		returntype := returnedTypeM[sinfo.FullMethod]
		if returntype == nil {
			returntype = getReturnType(sinfo.Server, sinfo.FullMethod)
			returnedTypeM[sinfo.FullMethod] = returntype
		}
		lock.Unlock()
		return forward(cc, sinfo.FullMethod, returntype, ctx, in, extraHeader)
	}
}

// getReturnType returns the return types for a GRPC method
// the method name should be full method name (i.e., /package.service/method)
// For example, with handler
//
//	(s *server) func Goodbye() string {}
//	(s *server) func Ping(_ context.Context, _ *pb.Ping) (*pb.Pong, error) {}
//	(s *server) func Hello(_ context.Context, _ *pb.Empty) (*pb.String, error) {}
func getReturnType(server interface{}, fullmethod string) reflect.Type {
	t := reflect.TypeOf(server)
	for i := 0; i < t.NumMethod(); i++ {
		methodType := t.Method(i).Type

		if !strings.HasSuffix(fullmethod, "/"+t.Method(i).Name) {
			continue
		}

		if methodType.NumOut() != 2 || methodType.NumIn() < 2 {
			continue
		}

		// the first parameter should context and the second one should be a pointer
		if methodType.In(1).Name() != "Context" || methodType.In(2).Kind() != reflect.Ptr {
			continue
		}

		// the first output should be a pointer and the second one should be an error
		if methodType.Out(0).Kind() != reflect.Ptr || methodType.Out(1).Name() != "error" {
			continue
		}

		return methodType.Out(0).Elem()

	}
	return nil
}

func GetAccountId(ctx context.Context, message interface{}) string {
	msgrefl := message.(proto.Message).ProtoReflect()
	accIdDesc := msgrefl.Descriptor().Fields().ByName("account_id")
	accid := ""
	if accIdDesc == nil {
		accid = ""
	} else {
		accid = msgrefl.Get(accIdDesc).String()
	}

	if accid == "" {
		accid = FromGrpcCtx(ctx).GetCredential().GetAccountId()
	}

	return accid
}

var prommetrics *grpcprom.ClientMetrics

func init() {
	prommetrics = grpcprom.NewClientMetrics(grpcprom.WithClientCounterOptions())
	prometheus.MustRegister(prommetrics)

	conf := &ServerConfig{}

	// read fom file
	if dat, _ := os.ReadFile(".server_config"); len(dat) > 0 {
		json.Unmarshal(dat, conf)
	}

	makeSureServerConfiguation(conf)
	_serverConfiguration = *conf
	go func() {
		for {
			conf, err := fetchServerConfig()
			if err == nil && conf != nil {
				makeSureServerConfiguation(conf)
				_serverConfiguration = *conf
			}

			time.Sleep(30 * time.Second)
		}
	}()
}

func fetchServerConfig() (*ServerConfig, error) {
	resp, err := http.Get("https://config.subiz.net/subiz")
	if err != nil {
		return nil, log.EServer(err, log.M{"MSG": "CANNOT FETCH CONFIG"})
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, log.EServer(err, log.M{"_payload": bodyBytes, "MSG": "CANNOT FETCH CONFIG"})
	}

	conf := &ServerConfig{}
	err = json.Unmarshal(bodyBytes, conf)
	if err == nil {
		return conf, nil
	}
	return nil, log.EServer(err, log.M{"_payload": bodyBytes, "MSG": "CANNOT FETCH CONFIG INVALIDJSON"})
}

func makeSureServerConfiguation(conf *ServerConfig) {
	if conf == nil {
		return
	}
	// make sure service config is good
	if conf.Services == nil {
		conf.Services = map[string]ServiceConfiguration{}
	}
	for _, s := range conf.Services {
		s.PortStr = strconv.Itoa(int(s.Port))
	}

	if conf.StaggingAccounts == nil {
		conf.StaggingAccounts = map[string]bool{}
	}

	if conf.ProductionAccounts == nil {
		conf.ProductionAccounts = map[string]bool{}
	}

	if conf.DevelopmentAccounts == nil {
		conf.DevelopmentAccounts = map[string]bool{}
	}

	if conf.XAccounts == nil {
		conf.XAccounts = map[string]bool{}
	}
}

func DialGrpc(service string, opts ...grpc.DialOption) *grpc.ClientConn {
	opts = append([]grpc.DialOption{}, opts...)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// opts = append(opts, sgrpc.WithCache())
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(50*1024*1024)))
	opts = append(opts, WithErrorStack())
	opts = append(opts, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	opts = append(opts, grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:    time.Duration(10) * time.Second,
		Timeout: time.Duration(60) * time.Second,
	}))
	opts = append(opts, grpc.WithChainUnaryInterceptor(prommetrics.UnaryClientInterceptor()))
	for {
		conn, err := grpc.Dial(service, opts...)
		if err != nil {
			fmt.Println("CANNOT CONNECT TO", service, ". Err", err, ". Retry in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}
		return conn
	}
}

func NewShardServer(port, shards int) *grpc.Server {
	return grpc.NewServer(grpc.KeepaliveParams(
		keepalive.ServerParameters{
			MaxConnectionAge: time.Duration(20) * time.Second,
		},
	), grpc.ChainUnaryInterceptor(NewStatefulSetShardInterceptor(port, shards)))
}

func NewServer() *grpc.Server {
	return grpc.NewServer(grpc.UnaryInterceptor(RecoverInterceptor),
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionAge: time.Duration(20) * time.Second,
			},
		))
}

var (
	dialLock             = &sync.Mutex{}
	ticketClients        map[string]TicketMgrClient
	_serverConfiguration ServerConfig
)

type ServiceConfiguration struct {
	PortStr   string // must fill using port
	Name      string `json:"name,omitempty"`
	Port      int64  `json:"port,omitempty"`
	NumShards int64  `json:"num_shards,omitempty"` // must >= 1
}

type ServerConfig struct {
	Services            map[string]ServiceConfiguration `json:"services,omitempty"`
	StaggingAccounts    map[string]bool                 `json:"stagging_accounts,omitempty"`
	ProductionAccounts  map[string]bool                 `json:"production_accounts,omitempty"`
	DevelopmentAccounts map[string]bool                 `json:"development_accounts,omitempty"`
	XAccounts           map[string]bool                 `json:"x_accounts,omitempty"`
}

func GetTicketClient(accid string) TicketMgrClient {
	const service = "ticket"
	num := ""
	if _serverConfiguration.XAccounts[accid] {
		num = "x"
	} else if _serverConfiguration.DevelopmentAccounts[accid] {
		num = "dev"
	} else if _serverConfiguration.StaggingAccounts[accid] {
		num = "stg"
	} else {
		if numShard := _serverConfiguration.Services[service].NumShards; numShard > 0 {
			num = strconv.Itoa(int(crc32.ChecksumIEEE([]byte(accid)) % uint32(numShard)))
		}
	}

	serviceLocation := service + "-" + num + "." + service + ":" + _serverConfiguration.Services[service].PortStr
	if client, has := ticketClients[serviceLocation]; has {
		return client
	}

	dialLock.Lock()
	// double check if someone already connected
	if client, has := ticketClients[serviceLocation]; has {
		dialLock.Unlock()
		return client
	}

	client := NewTicketMgrClient(DialGrpc(serviceLocation))
	newClients := map[string]TicketMgrClient{}
	for loc, oldClient := range ticketClients {
		newClients[loc] = oldClient
	}
	ticketClients = newClients
	dialLock.Unlock()
	return client
}
