// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto

It has these top-level messages:
	Empty
	Id
	Ids
	Context
	Reply
	Device
	Error
	PingRequest
	Pong
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import lang "git.subiz.net/header/lang"
import auth "git.subiz.net/header/auth"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type E int32

const (
	E_no_error                                   E = 0
	E_unknown                                    E = 1
	E_missing_account_id                         E = 10
	E_proto_marshal_error                        E = 11
	E_database_error                             E = 12
	E_access_deny                                E = 13
	E_json_marshal_error                         E = 14
	E_idempotency_input_changed                  E = 15
	E_kv_key_not_found                           E = 1500
	E_subscription_not_found                     E = 1012
	E_payment_method_not_found                   E = 1013
	E_missing_primary_payment_method             E = 1014
	E_payment_method_state_is_failed             E = 1015
	E_primary_payment_method_is_not_credit_card  E = 1016
	E_missing_payment_method_id                  E = 1017
	E_plan_not_found                             E = 1018
	E_invalid_plan_cannot_buy                    E = 1019
	E_missing_next_billing_cycle_month           E = 1020
	E_trial_package_cannot_buy                   E = 1021
	E_missing_billing_cycle_month                E = 1022
	E_primary_payment_method_must_be_credit_card E = 1023
	E_bank_transfer_payment_method_not_found     E = 1024
	E_invoice_not_found                          E = 1025
	E_invoice_template_error                     E = 1026
	E_execute_invoice_template_error             E = 1027
	E_missing_stripe_token                       E = 1028
	E_missing_stripe_info                        E = 1029
	E_stripe_call_error                          E = 1030
	E_missing_stripe_customer_id                 E = 1031
	E_invalid_invoice_id                         E = 1032
	E_invalid_payment_log_id                     E = 1034
	E_invalid_bill_id                            E = 1035
	E_s3_call_error                              E = 1503
	E_pdf_generate_error                         E = 1504
)

var E_name = map[int32]string{
	0:    "no_error",
	1:    "unknown",
	10:   "missing_account_id",
	11:   "proto_marshal_error",
	12:   "database_error",
	13:   "access_deny",
	14:   "json_marshal_error",
	15:   "idempotency_input_changed",
	1500: "kv_key_not_found",
	1012: "subscription_not_found",
	1013: "payment_method_not_found",
	1014: "missing_primary_payment_method",
	1015: "payment_method_state_is_failed",
	1016: "primary_payment_method_is_not_credit_card",
	1017: "missing_payment_method_id",
	1018: "plan_not_found",
	1019: "invalid_plan_cannot_buy",
	1020: "missing_next_billing_cycle_month",
	1021: "trial_package_cannot_buy",
	1022: "missing_billing_cycle_month",
	1023: "primary_payment_method_must_be_credit_card",
	1024: "bank_transfer_payment_method_not_found",
	1025: "invoice_not_found",
	1026: "invoice_template_error",
	1027: "execute_invoice_template_error",
	1028: "missing_stripe_token",
	1029: "missing_stripe_info",
	1030: "stripe_call_error",
	1031: "missing_stripe_customer_id",
	1032: "invalid_invoice_id",
	1034: "invalid_payment_log_id",
	1035: "invalid_bill_id",
	1503: "s3_call_error",
	1504: "pdf_generate_error",
}
var E_value = map[string]int32{
	"no_error":                                   0,
	"unknown":                                    1,
	"missing_account_id":                         10,
	"proto_marshal_error":                        11,
	"database_error":                             12,
	"access_deny":                                13,
	"json_marshal_error":                         14,
	"idempotency_input_changed":                  15,
	"kv_key_not_found":                           1500,
	"subscription_not_found":                     1012,
	"payment_method_not_found":                   1013,
	"missing_primary_payment_method":             1014,
	"payment_method_state_is_failed":             1015,
	"primary_payment_method_is_not_credit_card":  1016,
	"missing_payment_method_id":                  1017,
	"plan_not_found":                             1018,
	"invalid_plan_cannot_buy":                    1019,
	"missing_next_billing_cycle_month":           1020,
	"trial_package_cannot_buy":                   1021,
	"missing_billing_cycle_month":                1022,
	"primary_payment_method_must_be_credit_card": 1023,
	"bank_transfer_payment_method_not_found":     1024,
	"invoice_not_found":                          1025,
	"invoice_template_error":                     1026,
	"execute_invoice_template_error":             1027,
	"missing_stripe_token":                       1028,
	"missing_stripe_info":                        1029,
	"stripe_call_error":                          1030,
	"missing_stripe_customer_id":                 1031,
	"invalid_invoice_id":                         1032,
	"invalid_payment_log_id":                     1034,
	"invalid_bill_id":                            1035,
	"s3_call_error":                              1503,
	"pdf_generate_error":                         1504,
}

func (x E) String() string {
	return proto.EnumName(E_name, int32(x))
}
func (E) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Event int32

const (
	Event_Send_    Event = 0
	Event_ApiReply Event = 2
)

var Event_name = map[int32]string{
	0: "Send_",
	2: "ApiReply",
}
var Event_value = map[string]int32{
	"Send_":    0,
	"ApiReply": 2,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Device_DeviceType int32

const (
	Device_unknown Device_DeviceType = 0
	Device_mobile  Device_DeviceType = 1
	Device_tablet  Device_DeviceType = 2
	Device_desktop Device_DeviceType = 3
)

var Device_DeviceType_name = map[int32]string{
	0: "unknown",
	1: "mobile",
	2: "tablet",
	3: "desktop",
}
var Device_DeviceType_value = map[string]int32{
	"unknown": 0,
	"mobile":  1,
	"tablet":  2,
	"desktop": 3,
}

func (x Device_DeviceType) String() string {
	return proto.EnumName(Device_DeviceType_name, int32(x))
}
func (Device_DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type Empty struct {
	Ctx *Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Empty) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

type Id struct {
	Ctx       *Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId string   `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Id        string   `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Id) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Id) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Ids struct {
	Ctx       *Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId string   `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Ids       []string `protobuf:"bytes,3,rep,name=ids" json:"ids,omitempty"`
}

func (m *Ids) Reset()                    { *m = Ids{} }
func (m *Ids) String() string            { return proto.CompactTextString(m) }
func (*Ids) ProtoMessage()               {}
func (*Ids) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ids) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Ids) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Ids) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Context struct {
	EventId    string `protobuf:"bytes,1,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	State      []byte `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Node       string `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
	ReplyTopic string `protobuf:"bytes,4,opt,name=reply_topic,json=replyTopic" json:"reply_topic,omitempty"`
	// 	optional int32 reply_partition = 5;
	Credential *auth.Credential `protobuf:"bytes,6,opt,name=credential" json:"credential,omitempty"`
	Tracing    []byte           `protobuf:"bytes,7,opt,name=tracing,proto3" json:"tracing,omitempty"`
	ReplyKey   string           `protobuf:"bytes,8,opt,name=reply_key,json=replyKey" json:"reply_key,omitempty"`
	ByDevice   *Device          `protobuf:"bytes,10,opt,name=by_device,json=byDevice" json:"by_device,omitempty"`
	// for kafka
	Topic         string `protobuf:"bytes,11,opt,name=topic" json:"topic,omitempty"`
	Partition     int32  `protobuf:"varint,13,opt,name=partition" json:"partition,omitempty"`
	Offset        int64  `protobuf:"varint,14,opt,name=offset" json:"offset,omitempty"`
	Term          uint64 `protobuf:"varint,15,opt,name=term" json:"term,omitempty"`
	RouterTopic   string `protobuf:"bytes,16,opt,name=router_topic,json=routerTopic" json:"router_topic,omitempty"`
	IdempotentKey string `protobuf:"bytes,17,opt,name=idempotent_key,json=idempotentKey" json:"idempotent_key,omitempty"`
}

func (m *Context) Reset()                    { *m = Context{} }
func (m *Context) String() string            { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()               {}
func (*Context) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Context) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Context) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Context) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Context) GetReplyTopic() string {
	if m != nil {
		return m.ReplyTopic
	}
	return ""
}

func (m *Context) GetCredential() *auth.Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *Context) GetTracing() []byte {
	if m != nil {
		return m.Tracing
	}
	return nil
}

func (m *Context) GetReplyKey() string {
	if m != nil {
		return m.ReplyKey
	}
	return ""
}

func (m *Context) GetByDevice() *Device {
	if m != nil {
		return m.ByDevice
	}
	return nil
}

func (m *Context) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Context) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *Context) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Context) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Context) GetRouterTopic() string {
	if m != nil {
		return m.RouterTopic
	}
	return ""
}

func (m *Context) GetIdempotentKey() string {
	if m != nil {
		return m.IdempotentKey
	}
	return ""
}

type Reply struct {
	Ctx            *Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	EventId        string   `protobuf:"bytes,3,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	State          []byte   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Service        string   `protobuf:"bytes,5,opt,name=service" json:"service,omitempty"`
	ServiceId      string   `protobuf:"bytes,6,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
	Err            bool     `protobuf:"varint,10,opt,name=err" json:"err,omitempty"`
	ErrDescription string   `protobuf:"bytes,12,opt,name=err_description,json=errDescription" json:"err_description,omitempty"`
	ErrCode        lang.T   `protobuf:"varint,13,opt,name=err_code,json=errCode,enum=lang.T" json:"err_code,omitempty"`
	ErrClass       int32    `protobuf:"varint,15,opt,name=err_class,json=errClass" json:"err_class,omitempty"`
	ErrHash        string   `protobuf:"bytes,16,opt,name=err_hash,json=errHash" json:"err_hash,omitempty"`
	Payload        []byte   `protobuf:"bytes,7,opt,name=payload,proto3" json:"payload,omitempty"`
	Published      int64    `protobuf:"varint,20,opt,name=published" json:"published,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Reply) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Reply) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Reply) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Reply) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Reply) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *Reply) GetErr() bool {
	if m != nil {
		return m.Err
	}
	return false
}

func (m *Reply) GetErrDescription() string {
	if m != nil {
		return m.ErrDescription
	}
	return ""
}

func (m *Reply) GetErrCode() lang.T {
	if m != nil {
		return m.ErrCode
	}
	return lang.T_undefined
}

func (m *Reply) GetErrClass() int32 {
	if m != nil {
		return m.ErrClass
	}
	return 0
}

func (m *Reply) GetErrHash() string {
	if m != nil {
		return m.ErrHash
	}
	return ""
}

func (m *Reply) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Reply) GetPublished() int64 {
	if m != nil {
		return m.Published
	}
	return 0
}

type Device struct {
	Ip               string `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	UserAgent        string `protobuf:"bytes,4,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	ScreenResolution string `protobuf:"bytes,5,opt,name=screen_resolution,json=screenResolution" json:"screen_resolution,omitempty"`
	Timezone         string `protobuf:"bytes,6,opt,name=timezone" json:"timezone,omitempty"`
	Language         string `protobuf:"bytes,7,opt,name=language" json:"language,omitempty"`
	Referrer         string `protobuf:"bytes,8,opt,name=referrer" json:"referrer,omitempty"`
	Type             string `protobuf:"bytes,9,opt,name=type" json:"type,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Device) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Device) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *Device) GetScreenResolution() string {
	if m != nil {
		return m.ScreenResolution
	}
	return ""
}

func (m *Device) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Device) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Device) GetReferrer() string {
	if m != nil {
		return m.Referrer
	}
	return ""
}

func (m *Device) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Error struct {
	Ctx         *Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Debug       string   `protobuf:"bytes,3,opt,name=debug" json:"debug,omitempty"`
	Hash        string   `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	Class       int32    `protobuf:"varint,6,opt,name=class" json:"class,omitempty"`
	Stack       string   `protobuf:"bytes,7,opt,name=stack" json:"stack,omitempty"`
	Created     int64    `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	Code        string   `protobuf:"bytes,4,opt,name=code" json:"code,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Error) GetCtx() *Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Error) GetDebug() string {
	if m != nil {
		return m.Debug
	}
	return ""
}

func (m *Error) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Error) GetClass() int32 {
	if m != nil {
		return m.Class
	}
	return 0
}

func (m *Error) GetStack() string {
	if m != nil {
		return m.Stack
	}
	return ""
}

func (m *Error) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type PingRequest struct {
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Pong struct {
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Pong) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "common.Empty")
	proto.RegisterType((*Id)(nil), "common.Id")
	proto.RegisterType((*Ids)(nil), "common.Ids")
	proto.RegisterType((*Context)(nil), "common.Context")
	proto.RegisterType((*Reply)(nil), "common.Reply")
	proto.RegisterType((*Device)(nil), "common.Device")
	proto.RegisterType((*Error)(nil), "common.Error")
	proto.RegisterType((*PingRequest)(nil), "common.PingRequest")
	proto.RegisterType((*Pong)(nil), "common.Pong")
	proto.RegisterEnum("common.E", E_name, E_value)
	proto.RegisterEnum("common.Event", Event_name, Event_value)
	proto.RegisterEnum("common.Device_DeviceType", Device_DeviceType_name, Device_DeviceType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Agent service

type AgentClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*Pong, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/common.Agent/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agent service

type AgentServer interface {
	Ping(context.Context, *PingRequest) (*Pong, error)
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.Agent/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Agent_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/common.proto",
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5b, 0x6f, 0xdb, 0xc6,
	0x12, 0x8e, 0x24, 0x4b, 0xa4, 0x46, 0xb6, 0xcc, 0xac, 0x7d, 0x1c, 0xda, 0xce, 0x45, 0xd1, 0x39,
	0x39, 0x71, 0x1c, 0xc0, 0x2e, 0x9c, 0xf7, 0x02, 0x81, 0x63, 0xa0, 0x46, 0x5e, 0x02, 0x36, 0x28,
	0xd0, 0x27, 0x62, 0x45, 0x8e, 0xa4, 0xad, 0xc8, 0x5d, 0x76, 0x77, 0xe9, 0x9a, 0x79, 0xea, 0xbd,
	0x45, 0xdb, 0xff, 0xd2, 0xbf, 0x52, 0xf4, 0xb9, 0x40, 0x0b, 0xf4, 0x2f, 0xf4, 0x7e, 0xc7, 0x2e,
	0x49, 0x49, 0x31, 0x12, 0xc0, 0x0f, 0x7d, 0x91, 0x76, 0xbe, 0x99, 0x9d, 0xdd, 0xf9, 0xbe, 0xd9,
	0x21, 0x6c, 0x44, 0x22, 0x4d, 0x05, 0x3f, 0x2c, 0xff, 0x0e, 0x32, 0x29, 0xb4, 0x20, 0x9d, 0xd2,
	0xda, 0xf9, 0xdf, 0x84, 0xe9, 0x03, 0x95, 0x8f, 0xd8, 0xb3, 0x03, 0x8e, 0xfa, 0x70, 0x8a, 0x34,
	0x46, 0x79, 0x98, 0x50, 0x3e, 0xb1, 0x3f, 0x65, 0xf4, 0x4b, 0xa2, 0x68, 0xae, 0xa7, 0xf6, 0xa7,
	0x8c, 0x1a, 0xee, 0x43, 0xfb, 0x24, 0xcd, 0x74, 0x41, 0x6e, 0x43, 0x2b, 0xd2, 0xe7, 0x7e, 0x63,
	0xd0, 0xd8, 0xeb, 0x1d, 0xad, 0x1f, 0x54, 0x07, 0x1f, 0x0b, 0xae, 0xf1, 0x5c, 0x07, 0xc6, 0x37,
	0x7c, 0x03, 0x9a, 0xa7, 0xf1, 0x25, 0x02, 0xc9, 0x0d, 0x00, 0x1a, 0x45, 0x22, 0xe7, 0x3a, 0x64,
	0xb1, 0xdf, 0x1c, 0x34, 0xf6, 0xba, 0x41, 0xb7, 0x42, 0x4e, 0x63, 0xd2, 0x87, 0x26, 0x8b, 0xfd,
	0x96, 0x85, 0x9b, 0x2c, 0x1e, 0xbe, 0x09, 0xad, 0xd3, 0x58, 0xfd, 0x0b, 0x89, 0x3d, 0x68, 0xb1,
	0x58, 0xf9, 0xad, 0x41, 0x6b, 0xaf, 0x1b, 0x98, 0xe5, 0xf0, 0xcb, 0x16, 0x38, 0x55, 0x06, 0xb2,
	0x0d, 0x2e, 0x9e, 0x61, 0xb9, 0xb5, 0x61, 0xb7, 0x3a, 0xd6, 0x3e, 0x8d, 0xc9, 0x26, 0xb4, 0x95,
	0xa6, 0x1a, 0x6d, 0xca, 0xd5, 0xa0, 0x34, 0x08, 0x81, 0x15, 0x2e, 0x62, 0xac, 0x6e, 0x6a, 0xd7,
	0xe4, 0x16, 0xf4, 0x24, 0x66, 0x49, 0x11, 0x6a, 0x91, 0xb1, 0xc8, 0x5f, 0xb1, 0x2e, 0xb0, 0xd0,
	0x53, 0x83, 0x90, 0x57, 0x00, 0x22, 0x89, 0x31, 0x72, 0xcd, 0x68, 0xe2, 0x77, 0x6c, 0x31, 0xde,
	0x81, 0x65, 0xfc, 0x78, 0x8e, 0x07, 0x4b, 0x31, 0xc4, 0x07, 0x47, 0x4b, 0x1a, 0x31, 0x3e, 0xf1,
	0x1d, 0x7b, 0x7c, 0x6d, 0x92, 0x5d, 0xe8, 0x96, 0x87, 0xcd, 0xb0, 0xf0, 0x5d, 0x7b, 0x94, 0x6b,
	0x81, 0xc7, 0x58, 0x90, 0xfb, 0xd0, 0x1d, 0x15, 0x61, 0x8c, 0x67, 0x2c, 0x42, 0x1f, 0xec, 0x39,
	0xfd, 0x9a, 0xb4, 0x47, 0x16, 0x0d, 0xdc, 0x51, 0x51, 0xae, 0x4c, 0x81, 0xe5, 0x85, 0x7b, 0x36,
	0x4b, 0x69, 0x90, 0xeb, 0xd0, 0xcd, 0xa8, 0xd4, 0x4c, 0x33, 0xc1, 0xfd, 0xb5, 0x41, 0x63, 0xaf,
	0x1d, 0x2c, 0x00, 0xb2, 0x05, 0x1d, 0x31, 0x1e, 0x2b, 0xd4, 0x7e, 0x7f, 0xd0, 0xd8, 0x6b, 0x05,
	0x95, 0x65, 0x68, 0xd1, 0x28, 0x53, 0x7f, 0x7d, 0xd0, 0xd8, 0x5b, 0x09, 0xec, 0x9a, 0xdc, 0x86,
	0x55, 0x29, 0x72, 0x8d, 0xb2, 0xe2, 0xc5, 0xb3, 0xc7, 0xf4, 0x4a, 0xac, 0x24, 0xe6, 0x0e, 0xf4,
	0x59, 0x8c, 0x69, 0x26, 0xb4, 0xd1, 0xc0, 0x54, 0x74, 0xd5, 0x06, 0xad, 0x2d, 0xd0, 0xc7, 0x58,
	0x0c, 0xbf, 0x6f, 0x42, 0x3b, 0x30, 0x35, 0x5e, 0xa6, 0x1f, 0x96, 0x25, 0x6d, 0x5d, 0x46, 0x52,
	0x1f, 0x1c, 0x85, 0xd2, 0x52, 0xd6, 0x2e, 0xe3, 0x2b, 0xd3, 0xb4, 0x56, 0xb5, 0x34, 0xc9, 0x3a,
	0x65, 0x6b, 0x55, 0x48, 0xd9, 0x5a, 0x28, 0xa5, 0xe5, 0xd9, 0x0d, 0xcc, 0x92, 0xdc, 0x85, 0x75,
	0x94, 0x32, 0x8c, 0x51, 0x45, 0x92, 0x65, 0x96, 0xc2, 0x55, 0xbb, 0xab, 0x8f, 0x52, 0x3e, 0x5a,
	0xa0, 0x64, 0x08, 0xae, 0x09, 0x8c, 0x4c, 0x2b, 0x19, 0x92, 0xfb, 0x47, 0xce, 0x81, 0x7d, 0xa7,
	0x4f, 0x03, 0x07, 0xa5, 0x3c, 0x36, 0x6d, 0xb5, 0x0b, 0x5d, 0x1b, 0x93, 0x50, 0xa5, 0x2c, 0xb1,
	0xed, 0xc0, 0x6c, 0x3a, 0x36, 0xb6, 0xad, 0x52, 0xca, 0x70, 0x4a, 0xd5, 0xb4, 0x22, 0xd6, 0xec,
	0x7b, 0x8d, 0xaa, 0xa9, 0xa9, 0x27, 0xa3, 0x45, 0x22, 0x68, 0x5c, 0xf7, 0x4e, 0x65, 0x5a, 0x6d,
	0xf3, 0x51, 0xc2, 0xd4, 0x14, 0x63, 0x7f, 0xd3, 0x0a, 0xb8, 0x00, 0x86, 0x5f, 0x34, 0xa1, 0x53,
	0xb5, 0x86, 0x79, 0x8d, 0xd9, 0xfc, 0x35, 0x66, 0x86, 0x88, 0x5c, 0xa1, 0x0c, 0xe9, 0x04, 0xb9,
	0xae, 0x1a, 0xbc, 0x6b, 0x90, 0x87, 0x06, 0x20, 0xf7, 0xe1, 0xaa, 0x8a, 0x24, 0x22, 0x0f, 0x25,
	0x2a, 0x91, 0xe4, 0xb6, 0xf0, 0x92, 0x4b, 0xaf, 0x74, 0x04, 0x73, 0x9c, 0xec, 0x80, 0xab, 0x59,
	0x8a, 0xcf, 0x04, 0xc7, 0x8a, 0xd2, 0xb9, 0x6d, 0x7c, 0x86, 0x85, 0x9c, 0x4e, 0xd0, 0xde, 0xbd,
	0x1b, 0xcc, 0x6d, 0xe3, 0x93, 0x38, 0x46, 0x29, 0x51, 0x2e, 0xfa, 0xbe, 0xb4, 0x6d, 0xfb, 0x15,
	0x19, 0xfa, 0xdd, 0xf2, 0x55, 0x9a, 0xf5, 0xf0, 0x55, 0x80, 0xb2, 0x9a, 0xa7, 0x45, 0x86, 0xa4,
	0x07, 0x4e, 0xce, 0x67, 0x5c, 0xbc, 0xc3, 0xbd, 0x2b, 0x04, 0xa0, 0x93, 0x8a, 0x11, 0x4b, 0xd0,
	0x6b, 0x98, 0xb5, 0xa6, 0xa3, 0x04, 0xb5, 0xd7, 0x34, 0x41, 0x31, 0xaa, 0x99, 0x16, 0x99, 0xd7,
	0x1a, 0x7e, 0xd5, 0x80, 0xf6, 0x89, 0x94, 0x42, 0x5e, 0xa6, 0xe9, 0x06, 0xd0, 0x5b, 0x16, 0xbd,
	0x9c, 0x42, 0xcb, 0x90, 0xe9, 0xbd, 0x18, 0x47, 0xf9, 0xa4, 0x62, 0xb5, 0x34, 0xcc, 0xc5, 0xad,
	0x84, 0x25, 0x59, 0x76, 0x6d, 0x22, 0x4b, 0xcd, 0x3b, 0x56, 0xf3, 0xd2, 0xa8, 0x7a, 0x37, 0x9a,
	0x55, 0xbc, 0x94, 0x86, 0xd1, 0x3a, 0x92, 0x48, 0x35, 0xc6, 0x96, 0x93, 0x56, 0x50, 0x9b, 0x26,
	0xb3, 0xed, 0xae, 0x52, 0x2c, 0xbb, 0x1e, 0xde, 0x85, 0xde, 0x13, 0xc6, 0x27, 0x01, 0xbe, 0x9d,
	0xa3, 0xd2, 0x66, 0x73, 0x8a, 0x4a, 0x19, 0xb2, 0xcb, 0x0b, 0xd7, 0xe6, 0x70, 0x00, 0x2b, 0x4f,
	0x04, 0x9f, 0x2c, 0x47, 0xb4, 0x9e, 0x8b, 0xd8, 0xff, 0xda, 0x81, 0xc6, 0x09, 0x59, 0x05, 0x97,
	0x8b, 0x10, 0x0d, 0x4b, 0xde, 0x95, 0x65, 0x8e, 0x1b, 0x64, 0x0b, 0x48, 0xca, 0x94, 0x62, 0x7c,
	0x12, 0x2e, 0xc6, 0xb3, 0x07, 0xe4, 0x1a, 0x6c, 0xd8, 0xaf, 0x4c, 0x98, 0x52, 0xa9, 0xa6, 0x34,
	0xa9, 0x76, 0xf7, 0x08, 0x81, 0x7e, 0x4c, 0x35, 0x1d, 0x51, 0x85, 0x15, 0xb6, 0x4a, 0xd6, 0xa1,
	0x47, 0xa3, 0x08, 0x95, 0x0a, 0x63, 0xe4, 0x85, 0xb7, 0x66, 0xb2, 0xbe, 0xa5, 0x04, 0xbf, 0xb0,
	0xb9, 0x4f, 0x6e, 0xc0, 0xf6, 0x7c, 0x64, 0x44, 0x45, 0xc8, 0x78, 0x96, 0xeb, 0x30, 0x9a, 0x52,
	0x3e, 0xc1, 0xd8, 0x5b, 0x27, 0xff, 0x01, 0x6f, 0x76, 0x66, 0xe6, 0x4b, 0xc8, 0x85, 0x0e, 0xc7,
	0x22, 0xe7, 0xb1, 0xf7, 0x4d, 0x8f, 0xec, 0xc2, 0x96, 0xca, 0x47, 0x73, 0x8d, 0x96, 0x9c, 0x3f,
	0x38, 0xe4, 0x06, 0xf8, 0x19, 0x2d, 0x52, 0x33, 0x49, 0x52, 0xd4, 0x53, 0x11, 0x2f, 0xb9, 0x7f,
	0x74, 0xc8, 0x7f, 0xe1, 0x66, 0x5d, 0x5f, 0x26, 0x59, 0x4a, 0x65, 0x11, 0x3e, 0x1f, 0xee, 0xfd,
	0x64, 0x83, 0x2e, 0xe4, 0xb0, 0x23, 0x27, 0x64, 0x2a, 0x1c, 0x53, 0x96, 0x60, 0xec, 0xfd, 0xec,
	0x90, 0x03, 0xb8, 0xf7, 0xe2, 0x0c, 0x26, 0xcc, 0x9c, 0x69, 0xbe, 0x0b, 0x4c, 0x87, 0x11, 0x95,
	0xb1, 0xf7, 0x8b, 0x43, 0x6e, 0xc2, 0xf6, 0xfc, 0xe4, 0x0b, 0xf1, 0xb1, 0xf7, 0xab, 0x43, 0x36,
	0xa0, 0x9f, 0x25, 0x74, 0xb9, 0x9a, 0xdf, 0x1c, 0x72, 0x1d, 0xae, 0x31, 0x7e, 0x46, 0x13, 0x16,
	0x87, 0xd6, 0x19, 0x51, 0x6e, 0xfc, 0xa3, 0xbc, 0xf0, 0x7e, 0x77, 0xc8, 0x1d, 0x18, 0xd4, 0x29,
	0x39, 0x9e, 0xeb, 0x70, 0xc4, 0x92, 0xc4, 0x18, 0x51, 0x11, 0x25, 0x18, 0xa6, 0x82, 0xeb, 0xa9,
	0xf7, 0x87, 0xa5, 0x44, 0x4b, 0x46, 0x93, 0x30, 0xa3, 0xd1, 0x8c, 0x4e, 0x70, 0x39, 0xcb, 0x9f,
	0x0e, 0x19, 0xc0, 0x6e, 0x9d, 0xe5, 0x45, 0x09, 0xfe, 0x72, 0xc8, 0x21, 0xec, 0xbf, 0xa4, 0xd4,
	0x34, 0x57, 0x3a, 0x1c, 0xe1, 0x73, 0xb5, 0xfe, 0xed, 0x90, 0xfb, 0xf0, 0xff, 0x11, 0xe5, 0xb3,
	0x50, 0x4b, 0xca, 0xd5, 0x18, 0x65, 0xf8, 0x52, 0x49, 0xde, 0x75, 0xc9, 0x16, 0x5c, 0x65, 0xfc,
	0x4c, 0x98, 0x71, 0xbd, 0xc0, 0xdf, 0x73, 0x8d, 0xcc, 0x35, 0xae, 0x31, 0xcd, 0x12, 0xa3, 0x40,
	0xd9, 0x38, 0xef, 0xbb, 0x46, 0x22, 0x3c, 0xc7, 0x28, 0x37, 0xaa, 0xbc, 0x38, 0xe8, 0x03, 0x97,
	0x6c, 0xc3, 0x66, 0x5d, 0x99, 0xd2, 0x92, 0x65, 0x18, 0x6a, 0x31, 0x43, 0xee, 0x7d, 0xe8, 0x12,
	0x1f, 0x36, 0x2e, 0xb8, 0x18, 0x1f, 0x0b, 0xef, 0x23, 0x7b, 0x9d, 0x0a, 0x89, 0x68, 0x52, 0xb7,
	0xea, 0xc7, 0x2e, 0xb9, 0x05, 0x3b, 0x17, 0x76, 0x44, 0xb9, 0xd2, 0x22, 0x45, 0x69, 0x04, 0xfc,
	0xc4, 0x25, 0xd7, 0x80, 0xd4, 0x5a, 0xd5, 0x57, 0x62, 0xb1, 0xf7, 0x69, 0x5d, 0x48, 0x29, 0x62,
	0xc5, 0x43, 0x22, 0x26, 0xc6, 0xf9, 0x99, 0x4b, 0x36, 0x61, 0xbd, 0x76, 0x1a, 0xf6, 0x0d, 0xfa,
	0xb9, 0x4b, 0x08, 0xac, 0xa9, 0x07, 0xcb, 0x17, 0xf8, 0xb6, 0x67, 0xf2, 0x67, 0xf1, 0x38, 0x9c,
	0x20, 0x47, 0xb9, 0x28, 0xf3, 0xbb, 0xde, 0xfe, 0x00, 0xda, 0x27, 0xe6, 0x53, 0x49, 0xba, 0xd0,
	0x7e, 0x1d, 0x79, 0x1c, 0x7a, 0x57, 0xcc, 0x13, 0x7f, 0x98, 0x31, 0xfb, 0xf5, 0xf5, 0x9a, 0x47,
	0x47, 0xd0, 0x2e, 0x47, 0xfe, 0x3d, 0x58, 0x31, 0xa3, 0x84, 0x6c, 0xd4, 0xe3, 0x70, 0x69, 0xb0,
	0xec, 0xac, 0xce, 0x41, 0xc1, 0x27, 0xa3, 0x8e, 0x7d, 0xef, 0x0f, 0xfe, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0x05, 0xee, 0xb2, 0xc0, 0x0a, 0x00, 0x00,
}
