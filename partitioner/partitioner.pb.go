// Code generated by protoc-gen-go. DO NOT EDIT.
// source: partitioner/partitioner.proto

package partitioner

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
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

type Configuration struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Term    int32  `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	// ID of the cluster
	Cluster string `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// 	repeated WorkerConfiguration workers = 5;
	TotalPartitions      int32                        `protobuf:"varint,6,opt,name=total_partitions,json=totalPartitions,proto3" json:"total_partitions,omitempty"`
	NextTerm             int32                        `protobuf:"varint,7,opt,name=next_term,json=nextTerm,proto3" json:"next_term,omitempty"`
	Partitions           map[string]*WorkerPartitions `protobuf:"bytes,9,rep,name=partitions,proto3" json:"partitions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Hosts                map[string]string            `protobuf:"bytes,10,rep,name=hosts,proto3" json:"hosts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{0}
}

func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Configuration) GetTerm() int32 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Configuration) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Configuration) GetTotalPartitions() int32 {
	if m != nil {
		return m.TotalPartitions
	}
	return 0
}

func (m *Configuration) GetNextTerm() int32 {
	if m != nil {
		return m.NextTerm
	}
	return 0
}

func (m *Configuration) GetPartitions() map[string]*WorkerPartitions {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func (m *Configuration) GetHosts() map[string]string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

type WorkerHost struct {
	Cluster              string   `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Host                 string   `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkerHost) Reset()         { *m = WorkerHost{} }
func (m *WorkerHost) String() string { return proto.CompactTextString(m) }
func (*WorkerHost) ProtoMessage()    {}
func (*WorkerHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{1}
}

func (m *WorkerHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerHost.Unmarshal(m, b)
}
func (m *WorkerHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerHost.Marshal(b, m, deterministic)
}
func (m *WorkerHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerHost.Merge(m, src)
}
func (m *WorkerHost) XXX_Size() int {
	return xxx_messageInfo_WorkerHost.Size(m)
}
func (m *WorkerHost) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerHost.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerHost proto.InternalMessageInfo

func (m *WorkerHost) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *WorkerHost) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WorkerHost) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type WorkerPartitions struct {
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Partitions           []int32  `protobuf:"varint,7,rep,packed,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkerPartitions) Reset()         { *m = WorkerPartitions{} }
func (m *WorkerPartitions) String() string { return proto.CompactTextString(m) }
func (*WorkerPartitions) ProtoMessage()    {}
func (*WorkerPartitions) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{2}
}

func (m *WorkerPartitions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerPartitions.Unmarshal(m, b)
}
func (m *WorkerPartitions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerPartitions.Marshal(b, m, deterministic)
}
func (m *WorkerPartitions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerPartitions.Merge(m, src)
}
func (m *WorkerPartitions) XXX_Size() int {
	return xxx_messageInfo_WorkerPartitions.Size(m)
}
func (m *WorkerPartitions) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerPartitions.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerPartitions proto.InternalMessageInfo

func (m *WorkerPartitions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WorkerPartitions) GetPartitions() []int32 {
	if m != nil {
		return m.Partitions
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Cluster struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_e94f51d1df456051, []int{4}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Configuration)(nil), "partitioner.Configuration")
	proto.RegisterMapType((map[string]string)(nil), "partitioner.Configuration.HostsEntry")
	proto.RegisterMapType((map[string]*WorkerPartitions)(nil), "partitioner.Configuration.PartitionsEntry")
	proto.RegisterType((*WorkerHost)(nil), "partitioner.WorkerHost")
	proto.RegisterType((*WorkerPartitions)(nil), "partitioner.WorkerPartitions")
	proto.RegisterType((*Empty)(nil), "partitioner.Empty")
	proto.RegisterType((*Cluster)(nil), "partitioner.Cluster")
}

func init() { proto.RegisterFile("partitioner/partitioner.proto", fileDescriptor_e94f51d1df456051) }

var fileDescriptor_e94f51d1df456051 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0xbb, 0x49, 0xd3, 0x98, 0x09, 0x7a, 0x65, 0x38, 0x70, 0x8d, 0x9c, 0x84, 0x80, 0x10,
	0x7d, 0xa8, 0x90, 0x43, 0x38, 0xee, 0x1e, 0x04, 0xcb, 0x71, 0xd2, 0xa7, 0x23, 0x08, 0xbe, 0x08,
	0x47, 0xb4, 0xab, 0x2e, 0xd7, 0xcb, 0x86, 0xcd, 0xf6, 0xb0, 0x1f, 0xc0, 0xef, 0x25, 0x7e, 0x32,
	0xd9, 0xdd, 0x6b, 0xbb, 0xa9, 0x69, 0x5f, 0x7c, 0x9b, 0xfc, 0x3b, 0xff, 0xdf, 0xcc, 0xce, 0x9f,
	0xc2, 0x49, 0x53, 0x49, 0xc5, 0x15, 0x17, 0x35, 0x93, 0x6f, 0x9c, 0x7a, 0xd2, 0x48, 0xa1, 0x04,
	0xc6, 0x8e, 0x94, 0xfd, 0xf6, 0xe1, 0xf1, 0x54, 0xd4, 0xdf, 0xf8, 0xf7, 0xa5, 0xac, 0xb4, 0x86,
	0x14, 0xc2, 0x7b, 0x26, 0x5b, 0x2e, 0x6a, 0x4a, 0x52, 0x92, 0x47, 0xe5, 0xfa, 0x13, 0x11, 0x86,
	0x8a, 0xc9, 0x3b, 0xea, 0xa5, 0x24, 0x0f, 0x4a, 0x53, 0xeb, 0xee, 0xaf, 0x8b, 0x65, 0xab, 0x98,
	0xa4, 0xbe, 0xed, 0x7e, 0xf8, 0xc4, 0x57, 0x30, 0x56, 0x42, 0x55, 0x8b, 0x9b, 0xcd, 0xb8, 0x96,
	0x8e, 0x8c, 0xf3, 0xc8, 0xe8, 0xd7, 0x1b, 0x19, 0x9f, 0x43, 0x54, 0xb3, 0x9f, 0xea, 0xc6, 0xd0,
	0x43, 0xd3, 0xf3, 0x48, 0x0b, 0x1f, 0xf5, 0x84, 0x19, 0x80, 0x43, 0x88, 0x52, 0x3f, 0x8f, 0x8b,
	0xd7, 0x13, 0xf7, 0x59, 0x9d, 0xfd, 0x27, 0x5b, 0xee, 0x65, 0xad, 0xe4, 0xaa, 0x74, 0xdc, 0x78,
	0x01, 0xc1, 0x0f, 0xd1, 0xaa, 0x96, 0x82, 0xc1, 0xbc, 0x3c, 0x80, 0xf9, 0xa0, 0xfb, 0x2c, 0xc1,
	0x7a, 0x92, 0xcf, 0x70, 0xb4, 0xc3, 0xc6, 0x31, 0xf8, 0xb7, 0x6c, 0xf5, 0x70, 0x27, 0x5d, 0xe2,
	0x29, 0x04, 0xf7, 0xd5, 0x62, 0xc9, 0xcc, 0x91, 0xe2, 0xe2, 0xa4, 0x33, 0xe1, 0x93, 0x90, 0xb7,
	0x4c, 0x6e, 0x21, 0xa5, 0xed, 0x3d, 0xf7, 0xce, 0x48, 0x72, 0x06, 0xb0, 0x1d, 0xd9, 0x03, 0x3e,
	0x76, 0xc1, 0x91, 0xe3, 0xcc, 0x66, 0x00, 0x16, 0xac, 0xfd, 0x6e, 0x20, 0xa4, 0x1b, 0xc8, 0x13,
	0xf0, 0xf8, 0x9c, 0x0e, 0x8d, 0xe8, 0xf1, 0xb9, 0x8e, 0x53, 0x3f, 0x8c, 0x06, 0x46, 0x31, 0x75,
	0xf6, 0x1e, 0xc6, 0xbb, 0x4b, 0xfe, 0xe3, 0x7b, 0xd1, 0x09, 0x24, 0x4c, 0xfd, 0x3c, 0x70, 0x8f,
	0x9c, 0x85, 0x10, 0x5c, 0xde, 0x35, 0x6a, 0x95, 0x3d, 0x83, 0x70, 0xda, 0x99, 0x4d, 0xd6, 0x8c,
	0xe2, 0x17, 0x81, 0x78, 0x2a, 0x84, 0x9c, 0xf3, 0xba, 0x52, 0x42, 0xe2, 0x3b, 0x88, 0xae, 0x98,
	0xb2, 0x09, 0xe0, 0x71, 0x37, 0x16, 0x8b, 0x48, 0x92, 0xfd, 0x61, 0x65, 0x03, 0x7c, 0x0b, 0xc3,
	0x99, 0xe0, 0x35, 0x3e, 0xed, 0x39, 0xb8, 0xbe, 0x4b, 0x82, 0x9d, 0x1f, 0xec, 0x82, 0x83, 0xe2,
	0x0f, 0x81, 0x91, 0x6d, 0xfa, 0xff, 0x15, 0xce, 0x21, 0xbc, 0x62, 0xca, 0x84, 0xd0, 0x6f, 0xdf,
	0xb7, 0x5b, 0x36, 0xc0, 0x0b, 0x08, 0xaf, 0x25, 0x6b, 0x2a, 0xc9, 0xf0, 0xc0, 0x90, 0xfe, 0x47,
	0x7c, 0x19, 0x99, 0xff, 0xf5, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0xbf, 0xc9, 0x0b,
	0xf8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoordinatorClient interface {
	GetConfig(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Configuration, error)
	Join(ctx context.Context, in *WorkerHost, opts ...grpc.CallOption) (*Empty, error)
}

type coordinatorClient struct {
	cc *grpc.ClientConn
}

func NewCoordinatorClient(cc *grpc.ClientConn) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) GetConfig(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/partitioner.Coordinator/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorClient) Join(ctx context.Context, in *WorkerHost, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/partitioner.Coordinator/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
type CoordinatorServer interface {
	GetConfig(context.Context, *Cluster) (*Configuration, error)
	Join(context.Context, *WorkerHost) (*Empty, error)
}

func RegisterCoordinatorServer(s *grpc.Server, srv CoordinatorServer) {
	s.RegisterService(&_Coordinator_serviceDesc, srv)
}

func _Coordinator_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partitioner.Coordinator/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).GetConfig(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coordinator_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerHost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partitioner.Coordinator/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).Join(ctx, req.(*WorkerHost))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coordinator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "partitioner.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Coordinator_GetConfig_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Coordinator_Join_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "partitioner/partitioner.proto",
}

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerClient interface {
	GetConfig(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Configuration, error)
	GetHost(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*WorkerHost, error)
	Prepare(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*Empty, error)
}

type workerClient struct {
	cc *grpc.ClientConn
}

func NewWorkerClient(cc *grpc.ClientConn) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) GetConfig(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/partitioner.Worker/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) GetHost(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*WorkerHost, error) {
	out := new(WorkerHost)
	err := c.cc.Invoke(ctx, "/partitioner.Worker/GetHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Prepare(ctx context.Context, in *Configuration, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/partitioner.Worker/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
type WorkerServer interface {
	GetConfig(context.Context, *Cluster) (*Configuration, error)
	GetHost(context.Context, *Cluster) (*WorkerHost, error)
	Prepare(context.Context, *Configuration) (*Empty, error)
}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partitioner.Worker/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetConfig(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_GetHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partitioner.Worker/GetHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetHost(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Configuration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/partitioner.Worker/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Prepare(ctx, req.(*Configuration))
	}
	return interceptor(ctx, in, info, handler)
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "partitioner.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Worker_GetConfig_Handler,
		},
		{
			MethodName: "GetHost",
			Handler:    _Worker_GetHost_Handler,
		},
		{
			MethodName: "Prepare",
			Handler:    _Worker_Prepare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "partitioner/partitioner.proto",
}
