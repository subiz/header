// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/api.proto

It has these top-level messages:
	Empty
	SendMessageState
	State
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Method int32

const (
	Method_GET    Method = 0
	Method_POST   Method = 1
	Method_PUT    Method = 2
	Method_DELETE Method = 3
	Method_PATCH  Method = 4
)

var Method_name = map[int32]string{
	0: "GET",
	1: "POST",
	2: "PUT",
	3: "DELETE",
	4: "PATCH",
}
var Method_value = map[string]int32{
	"GET":    0,
	"POST":   1,
	"PUT":    2,
	"DELETE": 3,
	"PATCH":  4,
}

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}
func (x Method) String() string {
	return proto.EnumName(Method_name, int32(x))
}
func (x *Method) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Method_value, data, "Method")
	if err != nil {
		return err
	}
	*x = Method(value)
	return nil
}
func (Method) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Event int32

const (
	Event_ConversationStart Event = 2
	Event_ApiReply          Event = 4
)

var Event_name = map[int32]string{
	2: "ConversationStart",
	4: "ApiReply",
}
var Event_value = map[string]int32{
	"ConversationStart": 2,
	"ApiReply":          4,
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

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SendMessageState struct {
	Topic            *string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SendMessageState) Reset()                    { *m = SendMessageState{} }
func (m *SendMessageState) String() string            { return proto.CompactTextString(m) }
func (*SendMessageState) ProtoMessage()               {}
func (*SendMessageState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SendMessageState) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

type State struct {
	Event            *string `protobuf:"bytes,3,opt,name=event" json:"event,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *State) GetEvent() string {
	if m != nil && m.Event != nil {
		return *m.Event
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*SendMessageState)(nil), "api.SendMessageState")
	proto.RegisterType((*State)(nil), "api.State")
	proto.RegisterEnum("api.Method", Method_name, Method_value)
	proto.RegisterEnum("api.Event", Event_name, Event_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RestApi service

type RestApiClient interface {
}

type restApiClient struct {
	cc *grpc.ClientConn
}

func NewRestApiClient(cc *grpc.ClientConn) RestApiClient {
	return &restApiClient{cc}
}

// Server API for RestApi service

type RestApiServer interface {
}

func RegisterRestApiServer(s *grpc.Server, srv RestApiServer) {
	s.RegisterService(&_RestApi_serviceDesc, srv)
}

var _RestApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.RestApi",
	HandlerType: (*RestApiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "api/api.proto",
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xcc, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x71, 0xd3, 0x64, 0x9b, 0x66, 0x50, 0x18, 0x07, 0x85, 0x5e, 0x04, 0xc9, 0xa9, 0x14,
	0x51, 0xf0, 0xea, 0x29, 0xd4, 0x45, 0x0f, 0x16, 0x43, 0xb2, 0x3e, 0xc0, 0x62, 0x07, 0x5d, 0xd0,
	0xec, 0x90, 0x1d, 0x0a, 0x7d, 0x7b, 0xa9, 0x81, 0x1e, 0xff, 0xdf, 0x0f, 0x3e, 0xb8, 0xf0, 0x12,
	0x1e, 0xbc, 0x84, 0x7b, 0x19, 0xa3, 0x46, 0xca, 0xbd, 0x84, 0xba, 0x04, 0x63, 0x7f, 0x45, 0x0f,
	0xf5, 0x0a, 0xb0, 0xe7, 0x61, 0xb7, 0xe5, 0x94, 0xfc, 0x17, 0xf7, 0xea, 0x95, 0xe9, 0x0a, 0x8c,
	0x46, 0x09, 0x9f, 0xcb, 0xec, 0x36, 0x5b, 0x55, 0xdd, 0x14, 0xf5, 0x0d, 0x98, 0x13, 0xf3, 0x9e,
	0x07, 0x5d, 0xe6, 0x13, 0xff, 0xc7, 0xfa, 0x09, 0xe6, 0x5b, 0xd6, 0xef, 0xb8, 0xa3, 0x12, 0xf2,
	0x17, 0xeb, 0xf0, 0x8c, 0x16, 0x50, 0xb4, 0xef, 0xbd, 0xc3, 0xec, 0x38, 0xb5, 0x1f, 0x0e, 0x67,
	0x04, 0x30, 0x7f, 0xb6, 0x6f, 0xd6, 0x59, 0xcc, 0xa9, 0x02, 0xd3, 0x36, 0x6e, 0xf3, 0x8a, 0xc5,
	0xfa, 0x0e, 0x8c, 0x3d, 0xbe, 0xd0, 0x35, 0x5c, 0x6e, 0xe2, 0xb0, 0xe7, 0x31, 0x79, 0x0d, 0x71,
	0xe8, 0xd5, 0x8f, 0x8a, 0x33, 0x3a, 0x87, 0x45, 0x23, 0xa1, 0x63, 0xf9, 0x39, 0x60, 0xf1, 0x58,
	0x41, 0xd9, 0x71, 0xd2, 0x46, 0xc2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xb1, 0x55, 0xac,
	0xdc, 0x00, 0x00, 0x00,
}
