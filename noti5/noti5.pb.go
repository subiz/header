// Code generated by protoc-gen-go. DO NOT EDIT.
// source: noti5/noti5.proto

/*
Package noti5 is a generated protocol buffer package.

It is generated from these files:
	noti5/noti5.proto

It has these top-level messages:
	Setting
	Token
	PushNoti
	Message
*/
package noti5

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "git.subiz.net/header/common"

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

type Type int32

const (
	Type_new_message      Type = 0
	Type_new_conversation Type = 1
	Type_nothing          Type = 2
)

var Type_name = map[int32]string{
	0: "new_message",
	1: "new_conversation",
	2: "nothing",
}
var Type_value = map[string]int32{
	"new_message":      0,
	"new_conversation": 1,
	"nothing":          2,
}

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}
func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (x *Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Type_value, data, "Type")
	if err != nil {
		return err
	}
	*x = Type(value)
	return nil
}
func (Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Event int32

const (
	Event_Noti5Requested     Event = 0
	Event_Noti5PushRequested Event = 1
)

var Event_name = map[int32]string{
	0: "Noti5Requested",
	1: "Noti5PushRequested",
}
var Event_value = map[string]int32{
	"Noti5Requested":     0,
	"Noti5PushRequested": 1,
}

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}
func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (x *Event) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Event_value, data, "Event")
	if err != nil {
		return err
	}
	*x = Event(value)
	return nil
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Token_Platform int32

const (
	Token_desktop Token_Platform = 0
	Token_mobile  Token_Platform = 1
)

var Token_Platform_name = map[int32]string{
	0: "desktop",
	1: "mobile",
}
var Token_Platform_value = map[string]int32{
	"desktop": 0,
	"mobile":  1,
}

func (x Token_Platform) Enum() *Token_Platform {
	p := new(Token_Platform)
	*p = x
	return p
}
func (x Token_Platform) String() string {
	return proto.EnumName(Token_Platform_name, int32(x))
}
func (x *Token_Platform) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Token_Platform_value, data, "Token_Platform")
	if err != nil {
		return err
	}
	*x = Token_Platform(value)
	return nil
}
func (Token_Platform) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Setting struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	AgentId          *string         `protobuf:"bytes,3,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	WebType          *string         `protobuf:"bytes,4,opt,name=web_type,json=webType" json:"web_type,omitempty"`
	MobileType       *string         `protobuf:"bytes,5,opt,name=mobile_type,json=mobileType" json:"mobile_type,omitempty"`
	EmailType        *string         `protobuf:"bytes,6,opt,name=email_type,json=emailType" json:"email_type,omitempty"`
	EmailAfter       *int32          `protobuf:"varint,7,opt,name=email_after,json=emailAfter" json:"email_after,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Setting) Reset()                    { *m = Setting{} }
func (m *Setting) String() string            { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()               {}
func (*Setting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Setting) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Setting) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Setting) GetAgentId() string {
	if m != nil && m.AgentId != nil {
		return *m.AgentId
	}
	return ""
}

func (m *Setting) GetWebType() string {
	if m != nil && m.WebType != nil {
		return *m.WebType
	}
	return ""
}

func (m *Setting) GetMobileType() string {
	if m != nil && m.MobileType != nil {
		return *m.MobileType
	}
	return ""
}

func (m *Setting) GetEmailType() string {
	if m != nil && m.EmailType != nil {
		return *m.EmailType
	}
	return ""
}

func (m *Setting) GetEmailAfter() int32 {
	if m != nil && m.EmailAfter != nil {
		return *m.EmailAfter
	}
	return 0
}

type Token struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId           *string         `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserType         *string         `protobuf:"bytes,4,opt,name=user_type,json=userType" json:"user_type,omitempty"`
	Token            *string         `protobuf:"bytes,5,opt,name=token" json:"token,omitempty"`
	DeviceId         *string         `protobuf:"bytes,6,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Platform         *string         `protobuf:"bytes,7,opt,name=platform" json:"platform,omitempty"`
	Created          *int64          `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	Topics           []string        `protobuf:"bytes,9,rep,name=topics" json:"topics,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Token) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Token) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Token) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *Token) GetUserType() string {
	if m != nil && m.UserType != nil {
		return *m.UserType
	}
	return ""
}

func (m *Token) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Token) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

func (m *Token) GetPlatform() string {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return ""
}

func (m *Token) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Token) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type PushNoti struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId           *string         `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Platform         *string         `protobuf:"bytes,5,opt,name=platform" json:"platform,omitempty"`
	Title            *string         `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	Body             *string         `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PushNoti) Reset()                    { *m = PushNoti{} }
func (m *PushNoti) String() string            { return proto.CompactTextString(m) }
func (*PushNoti) ProtoMessage()               {}
func (*PushNoti) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PushNoti) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *PushNoti) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *PushNoti) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *PushNoti) GetPlatform() string {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return ""
}

func (m *PushNoti) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *PushNoti) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

type Message struct {
	Token            *string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Title            *string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Body             *string `protobuf:"bytes,4,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Message) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *Message) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Message) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Setting)(nil), "noti5.Setting")
	proto.RegisterType((*Token)(nil), "noti5.Token")
	proto.RegisterType((*PushNoti)(nil), "noti5.PushNoti")
	proto.RegisterType((*Message)(nil), "noti5.Message")
	proto.RegisterEnum("noti5.Type", Type_name, Type_value)
	proto.RegisterEnum("noti5.Event", Event_name, Event_value)
	proto.RegisterEnum("noti5.Token_Platform", Token_Platform_name, Token_Platform_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Noti5Service service

type Noti5ServiceClient interface {
	// rpc ReadNotificationSetting(common.Empty) returns (Setting);
	// rpc UpdateNotificationSetting(Setting) returns (common.Empty);
	AddToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error)
	RemoveToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error)
	Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error)
}

type noti5ServiceClient struct {
	cc *grpc.ClientConn
}

func NewNoti5ServiceClient(cc *grpc.ClientConn) Noti5ServiceClient {
	return &noti5ServiceClient{cc}
}

func (c *noti5ServiceClient) AddToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/noti5.Noti5Service/AddToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noti5ServiceClient) RemoveToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/noti5.Noti5Service/RemoveToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noti5ServiceClient) Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error) {
	out := new(common.Pong)
	err := grpc.Invoke(ctx, "/noti5.Noti5Service/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Noti5Service service

type Noti5ServiceServer interface {
	// rpc ReadNotificationSetting(common.Empty) returns (Setting);
	// rpc UpdateNotificationSetting(Setting) returns (common.Empty);
	AddToken(context.Context, *Token) (*common.Empty, error)
	RemoveToken(context.Context, *Token) (*common.Empty, error)
	Ping(context.Context, *common.PingRequest) (*common.Pong, error)
}

func RegisterNoti5ServiceServer(s *grpc.Server, srv Noti5ServiceServer) {
	s.RegisterService(&_Noti5Service_serviceDesc, srv)
}

func _Noti5Service_AddToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Noti5ServiceServer).AddToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noti5.Noti5Service/AddToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Noti5ServiceServer).AddToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noti5Service_RemoveToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Noti5ServiceServer).RemoveToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noti5.Noti5Service/RemoveToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Noti5ServiceServer).RemoveToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Noti5Service_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Noti5ServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noti5.Noti5Service/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Noti5ServiceServer).Ping(ctx, req.(*common.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Noti5Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "noti5.Noti5Service",
	HandlerType: (*Noti5ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToken",
			Handler:    _Noti5Service_AddToken_Handler,
		},
		{
			MethodName: "RemoveToken",
			Handler:    _Noti5Service_RemoveToken_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Noti5Service_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noti5/noti5.proto",
}

func init() { proto.RegisterFile("noti5/noti5.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xad, 0x63, 0x3b, 0x76, 0x26, 0xfd, 0xbe, 0x9a, 0xa5, 0x2a, 0x26, 0xa8, 0x22, 0x84, 0x03,
	0xa1, 0x87, 0x54, 0x2a, 0xea, 0x85, 0x5b, 0x85, 0x7a, 0xc8, 0x01, 0x54, 0xb9, 0xbd, 0x57, 0x8e,
	0x77, 0xea, 0xae, 0x1a, 0xef, 0x1a, 0x7b, 0x92, 0x36, 0xfc, 0x08, 0xfe, 0x05, 0x27, 0xfe, 0x15,
	0xbf, 0x04, 0xed, 0xae, 0x9d, 0x04, 0x84, 0x10, 0x48, 0x5c, 0x12, 0xbf, 0xf7, 0x76, 0xc6, 0xf3,
	0xde, 0x78, 0xe1, 0x91, 0x54, 0x24, 0x4e, 0x8f, 0xcd, 0xef, 0xa4, 0xac, 0x14, 0x29, 0xe6, 0x1b,
	0x30, 0x18, 0xe7, 0x82, 0x26, 0xf5, 0x62, 0x26, 0x3e, 0x4d, 0x24, 0xd2, 0xf1, 0x2d, 0xa6, 0x1c,
	0xab, 0xe3, 0x4c, 0x15, 0x85, 0x92, 0xcd, 0x9f, 0x2d, 0x18, 0x7d, 0x73, 0x20, 0xb8, 0x44, 0x22,
	0x21, 0x73, 0xf6, 0x02, 0xdc, 0x8c, 0x1e, 0x62, 0x67, 0xe8, 0x8c, 0xfb, 0x27, 0x7b, 0x93, 0xe6,
	0xdc, 0x3b, 0x25, 0x09, 0x1f, 0x28, 0xd1, 0x1a, 0x3b, 0x04, 0x48, 0xb3, 0x4c, 0x2d, 0x24, 0x5d,
	0x0b, 0x1e, 0x77, 0x86, 0xce, 0xb8, 0x97, 0xf4, 0x1a, 0x66, 0xca, 0xd9, 0x53, 0x08, 0xd3, 0x1c,
	0xad, 0xe8, 0x1a, 0x31, 0x30, 0xd8, 0x4a, 0xf7, 0x38, 0xbb, 0xa6, 0x55, 0x89, 0xb1, 0x67, 0xa5,
	0x7b, 0x9c, 0x5d, 0xad, 0x4a, 0x64, 0xcf, 0xa1, 0x5f, 0xa8, 0x99, 0x98, 0xa3, 0x55, 0x7d, 0xa3,
	0x82, 0xa5, 0xcc, 0x81, 0x43, 0x00, 0x2c, 0x52, 0x31, 0xb7, 0x7a, 0xd7, 0xbe, 0xd5, 0x30, 0x6d,
	0xbd, 0x95, 0xd3, 0x1b, 0xc2, 0x2a, 0x0e, 0x86, 0xce, 0xd8, 0x4f, 0x6c, 0xc5, 0x99, 0x66, 0x46,
	0x5f, 0x3a, 0xe0, 0x5f, 0xa9, 0x3b, 0x94, 0xff, 0xc0, 0xe2, 0x13, 0x08, 0x16, 0x35, 0x56, 0x1b,
	0x87, 0x5d, 0x0d, 0xa7, 0x9c, 0x3d, 0x83, 0x9e, 0x11, 0xb6, 0x1c, 0x86, 0x9a, 0x30, 0x23, 0xee,
	0x83, 0x4f, 0x7a, 0x80, 0xc6, 0x9c, 0x05, 0xba, 0x84, 0xe3, 0x52, 0x64, 0xa8, 0xbb, 0x59, 0x5b,
	0xa1, 0x25, 0xa6, 0x9c, 0x0d, 0x20, 0x2c, 0xe7, 0x29, 0xdd, 0xa8, 0xaa, 0x30, 0x96, 0x7a, 0xc9,
	0x1a, 0xb3, 0x18, 0x82, 0xac, 0xc2, 0x94, 0x90, 0xc7, 0xe1, 0xd0, 0x19, 0xbb, 0x49, 0x0b, 0xd9,
	0x01, 0x74, 0x49, 0x95, 0x22, 0xab, 0xe3, 0xde, 0xd0, 0xd5, 0xd3, 0x59, 0x34, 0x7a, 0x09, 0xe1,
	0x45, 0x5b, 0xdd, 0x87, 0x80, 0x63, 0x7d, 0x47, 0xaa, 0x8c, 0x76, 0x18, 0x40, 0xd7, 0x26, 0x1d,
	0x39, 0xa3, 0xaf, 0x0e, 0x84, 0x17, 0x8b, 0xfa, 0xf6, 0x83, 0x22, 0xf1, 0xf7, 0x51, 0xb9, 0xbf,
	0x89, 0xca, 0xfb, 0x21, 0xaa, 0x6d, 0x6b, 0xfe, 0x4f, 0xd6, 0x74, 0x52, 0x82, 0xe6, 0xed, 0x9a,
	0x2d, 0x60, 0x0c, 0xbc, 0x99, 0xe2, 0xab, 0x26, 0x08, 0xf3, 0x3c, 0x9a, 0x42, 0xf0, 0x1e, 0xeb,
	0x3a, 0xcd, 0xb7, 0xe2, 0xed, 0x6c, 0xc7, 0xbb, 0x6e, 0xe5, 0xfe, 0xaa, 0x95, 0xb7, 0x69, 0x75,
	0xf4, 0x16, 0x3c, 0xb3, 0xa6, 0x3d, 0xe8, 0x4b, 0xbc, 0xbf, 0x2e, 0x6c, 0xdb, 0x68, 0x87, 0xed,
	0x43, 0xa4, 0x89, 0x4c, 0xc9, 0x25, 0x56, 0x75, 0x4a, 0x42, 0xc9, 0xc8, 0xd1, 0x01, 0x4a, 0x45,
	0xb7, 0x42, 0xe6, 0x51, 0xe7, 0xe8, 0x0d, 0xf8, 0xe7, 0x4b, 0x94, 0xc4, 0x18, 0xfc, 0xaf, 0x83,
	0x3b, 0x4d, 0xf0, 0xe3, 0x02, 0x6b, 0x42, 0x1e, 0xed, 0xb0, 0x03, 0x60, 0x86, 0xd3, 0xa9, 0x6e,
	0x78, 0xe7, 0xe4, 0xb3, 0x03, 0xbb, 0x46, 0xb8, 0xc4, 0x4a, 0xef, 0x9b, 0xbd, 0x82, 0xf0, 0x8c,
	0x73, 0xfb, 0x91, 0xee, 0x4e, 0xec, 0x95, 0x36, 0x68, 0xf0, 0x5f, 0x1b, 0xfd, 0x79, 0x51, 0xd2,
	0x8a, 0x1d, 0x41, 0x3f, 0xc1, 0x42, 0x2d, 0xf1, 0x0f, 0xce, 0xbe, 0x06, 0xef, 0x42, 0x5f, 0xec,
	0xc7, 0x2d, 0xad, 0x51, 0x33, 0xc6, 0x60, 0x77, 0x4d, 0x2a, 0x99, 0x7f, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0x64, 0xbb, 0x7f, 0x84, 0x4c, 0x04, 0x00, 0x00,
}
