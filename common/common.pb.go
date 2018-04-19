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
	Error
	Device
	PingRequest
	Pong
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import lang "bitbucket.org/subiz/header/lang"
import auth "bitbucket.org/subiz/header/auth"

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
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
func (Device_DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

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
	Topic      string           `protobuf:"bytes,11,opt,name=topic" json:"topic,omitempty"`
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

type Error struct {
	Error       string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	RequestId   string `protobuf:"bytes,3,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Hash        string `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	Debug       string `protobuf:"bytes,6,opt,name=debug" json:"debug,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Error) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Error) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Error) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Error) GetDebug() string {
	if m != nil {
		return m.Debug
	}
	return ""
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
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
	proto.RegisterType((*Error)(nil), "common.Error")
	proto.RegisterType((*Device)(nil), "common.Device")
	proto.RegisterType((*PingRequest)(nil), "common.PingRequest")
	proto.RegisterType((*Pong)(nil), "common.Pong")
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
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0xde, 0xc4, 0x71, 0x62, 0x57, 0x86, 0x19, 0xd3, 0xbb, 0x07, 0x33, 0x0b, 0xc2, 0xf8, 0xb2,
	0x61, 0x56, 0xca, 0xa0, 0xe1, 0x8e, 0xb4, 0x9a, 0x1d, 0x89, 0xc0, 0x65, 0x65, 0x46, 0x48, 0x9c,
	0x22, 0xdb, 0x5d, 0x38, 0xad, 0x71, 0xba, 0x4d, 0x77, 0x7b, 0x59, 0xef, 0x13, 0x70, 0xe0, 0x69,
	0x78, 0x25, 0x5e, 0x04, 0x55, 0xb7, 0x93, 0x09, 0x17, 0x34, 0x07, 0x2e, 0x4e, 0x7d, 0x5f, 0x57,
	0xd7, 0xcf, 0x57, 0x5d, 0x81, 0xe7, 0xb5, 0xda, 0xef, 0x95, 0xbc, 0xf6, 0x3f, 0xeb, 0x4e, 0x2b,
	0xab, 0xd8, 0xdc, 0xa3, 0xcb, 0xab, 0x4a, 0xd8, 0xaa, 0xaf, 0x1f, 0xd0, 0xae, 0x95, 0x6e, 0xae,
	0x4d, 0x5f, 0x89, 0x8f, 0xd7, 0x3b, 0x2c, 0x39, 0xea, 0xeb, 0xb6, 0x94, 0x8d, 0xfb, 0xf8, 0x3b,
	0xff, 0xe9, 0x5b, 0xf6, 0x76, 0xe7, 0x3e, 0xde, 0x37, 0xbf, 0x82, 0xf0, 0x6e, 0xdf, 0xd9, 0x81,
	0x7d, 0x05, 0x41, 0x6d, 0x3f, 0xa4, 0x93, 0x6c, 0xb2, 0x5a, 0xde, 0x5c, 0xac, 0xc7, 0x22, 0x6e,
	0x95, 0xb4, 0xf8, 0xc1, 0x16, 0x74, 0x96, 0xff, 0x0c, 0xd3, 0x0d, 0x7f, 0x82, 0x23, 0xfb, 0x02,
	0xa0, 0xac, 0x6b, 0xd5, 0x4b, 0xbb, 0x15, 0x3c, 0x9d, 0x66, 0x93, 0x55, 0x5c, 0xc4, 0x23, 0xb3,
	0xe1, 0xec, 0x1c, 0xa6, 0x82, 0xa7, 0x81, 0xa3, 0xa7, 0x82, 0xe7, 0xbf, 0x40, 0xb0, 0xe1, 0xe6,
	0x7f, 0x08, 0x9c, 0x40, 0x20, 0xb8, 0x49, 0x83, 0x2c, 0x58, 0xc5, 0x05, 0x99, 0xf9, 0x5f, 0x53,
	0x58, 0x8c, 0x11, 0xd8, 0x67, 0x10, 0xe1, 0x7b, 0xf4, 0x57, 0x27, 0xee, 0xea, 0xc2, 0xe1, 0x0d,
	0x67, 0x2f, 0x20, 0x34, 0xb6, 0xb4, 0xe8, 0x42, 0x9e, 0x15, 0x1e, 0x30, 0x06, 0x33, 0xa9, 0x38,
	0x8e, 0x95, 0x3a, 0x9b, 0x7d, 0x09, 0x4b, 0x8d, 0x5d, 0x3b, 0x6c, 0xad, 0xea, 0x44, 0x9d, 0xce,
	0xdc, 0x11, 0x38, 0xea, 0x9e, 0x18, 0xf6, 0x0d, 0x40, 0xad, 0x91, 0xa3, 0xb4, 0xa2, 0x6c, 0xd3,
	0xb9, 0x6b, 0x26, 0x59, 0x3b, 0xc5, 0x6f, 0x8f, 0x7c, 0x71, 0xe2, 0xc3, 0x52, 0x58, 0x58, 0x5d,
	0xd6, 0x42, 0x36, 0xe9, 0xc2, 0xa5, 0x3f, 0x40, 0xf6, 0x12, 0x62, 0x9f, 0xec, 0x01, 0x87, 0x34,
	0x72, 0xa9, 0x22, 0x47, 0xfc, 0x88, 0x03, 0x7b, 0x0d, 0x71, 0x35, 0x6c, 0x39, 0xbe, 0x17, 0x35,
	0xa6, 0xe0, 0xf2, 0x9c, 0x1f, 0x44, 0x7b, 0xeb, 0xd8, 0x22, 0xaa, 0x06, 0x6f, 0x51, 0x83, 0xbe,
	0xe0, 0xa5, 0x8b, 0xe2, 0xc1, 0x0f, 0xb3, 0x28, 0x4c, 0xe6, 0xc5, 0x85, 0xcf, 0xd1, 0x95, 0xda,
	0x0a, 0x2b, 0x94, 0xcc, 0xff, 0x9e, 0x42, 0x58, 0x10, 0xf7, 0x94, 0x91, 0x9c, 0xaa, 0x1a, 0x3c,
	0x45, 0xd5, 0x14, 0x16, 0x06, 0xb5, 0xab, 0x3a, 0xf4, 0xfe, 0x23, 0xa4, 0xe9, 0x8e, 0x26, 0x05,
	0x9b, 0xfb, 0xe9, 0x8e, 0x8c, 0x9f, 0x2e, 0x6a, 0xed, 0x5a, 0x8d, 0x0a, 0x32, 0xd9, 0x2b, 0xb8,
	0x40, 0xad, 0xb7, 0x1c, 0x4d, 0xad, 0x45, 0x47, 0xb5, 0xa7, 0x67, 0xee, 0xd6, 0x39, 0x6a, 0xfd,
	0xf6, 0x91, 0x65, 0x39, 0x44, 0xe4, 0x58, 0xd3, 0x34, 0x3f, 0xc9, 0x26, 0xab, 0xf3, 0x9b, 0xc5,
	0xda, 0x2d, 0xcc, 0x7d, 0xb1, 0x40, 0xad, 0x6f, 0x69, 0xb2, 0x2f, 0x21, 0x76, 0x3e, 0x6d, 0x69,
	0x4c, 0x7a, 0x91, 0x4d, 0x56, 0x61, 0x41, 0x97, 0x6e, 0x09, 0xbb, 0x2e, 0xb5, 0xde, 0xee, 0x4a,
	0xb3, 0x4b, 0x93, 0xb1, 0x4b, 0xad, 0xbf, 0x2f, 0xcd, 0x8e, 0xfa, 0xe9, 0xca, 0xa1, 0x55, 0x25,
	0x3f, 0x8c, 0x6f, 0x84, 0xec, 0x73, 0x88, 0xbb, 0xbe, 0x6a, 0x85, 0xd9, 0x21, 0x4f, 0x5f, 0x64,
	0x93, 0x55, 0x50, 0x3c, 0x12, 0xf9, 0x1f, 0x13, 0x08, 0xef, 0xb4, 0x56, 0x9a, 0x74, 0x42, 0x32,
	0xc6, 0x07, 0xed, 0x01, 0xa9, 0xa1, 0xf1, 0xb7, 0x1e, 0xcd, 0x89, 0xb4, 0xf1, 0xc8, 0x6c, 0x38,
	0xcb, 0x60, 0x79, 0xda, 0xb7, 0x7f, 0x88, 0xa7, 0x14, 0x3d, 0x5f, 0x57, 0xaf, 0x57, 0xd9, 0xd9,
	0x94, 0x8a, 0x63, 0xd5, 0x37, 0xa3, 0xba, 0x1e, 0xe4, 0x7f, 0x4e, 0x61, 0x3e, 0x3e, 0x14, 0xda,
	0xcd, 0xee, 0xb8, 0x9b, 0x1d, 0x55, 0xd1, 0x1b, 0xd4, 0xdb, 0xb2, 0x41, 0x69, 0xc7, 0x2c, 0x31,
	0x31, 0x6f, 0x88, 0x60, 0xaf, 0xe1, 0x53, 0x53, 0x6b, 0x44, 0xb9, 0xd5, 0x68, 0x54, 0xdb, 0xbb,
	0x5a, 0x7c, 0xc2, 0xc4, 0x1f, 0x14, 0x47, 0x9e, 0x5d, 0x42, 0x64, 0xc5, 0x1e, 0x3f, 0x2a, 0x89,
	0x63, 0xfe, 0x23, 0xa6, 0x33, 0x1a, 0x48, 0x5f, 0x36, 0xe8, 0x64, 0x8c, 0x8b, 0x23, 0xa6, 0x33,
	0x8d, 0xbf, 0xa2, 0xd6, 0xa8, 0x1f, 0xb7, 0xc0, 0x63, 0x6a, 0xd2, 0x0e, 0x1d, 0xa6, 0xb1, 0x6f,
	0x92, 0xec, 0xfc, 0x3b, 0x00, 0xdf, 0xcd, 0xfd, 0xd0, 0x21, 0x5b, 0xc2, 0xa2, 0x97, 0x0f, 0x52,
	0xfd, 0x2e, 0x93, 0x67, 0x0c, 0x60, 0xbe, 0x57, 0x95, 0x68, 0x31, 0x99, 0x90, 0x6d, 0xcb, 0xaa,
	0x45, 0x9b, 0x4c, 0xc9, 0x89, 0xa3, 0x79, 0xb0, 0xaa, 0x4b, 0x82, 0xfc, 0x15, 0x2c, 0xdf, 0x09,
	0xd9, 0x14, 0x5e, 0x6b, 0x1a, 0xf0, 0x1e, 0x8d, 0xa1, 0xca, 0xfc, 0x80, 0x0e, 0x30, 0xcf, 0x60,
	0xf6, 0x4e, 0xc9, 0xe6, 0xd4, 0x23, 0xf8, 0x97, 0xc7, 0x55, 0x06, 0xe1, 0x1d, 0x6d, 0x03, 0x8b,
	0x21, 0xfc, 0x09, 0x25, 0xdf, 0x26, 0xcf, 0xd8, 0x19, 0x44, 0x6f, 0x3a, 0xe1, 0x16, 0x2c, 0x99,
	0xde, 0xdc, 0x40, 0xe8, 0xa5, 0xfc, 0x1a, 0x66, 0x94, 0x95, 0x3d, 0x3f, 0xac, 0xd9, 0x49, 0x0d,
	0x97, 0x67, 0x47, 0x52, 0xc9, 0xa6, 0x9a, 0xbb, 0xff, 0xee, 0x6f, 0xff, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x3b, 0x97, 0x65, 0xd8, 0x32, 0x06, 0x00, 0x00,
}
