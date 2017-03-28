// Code generated by protoc-gen-go.
// source: scheduler/scheduler.proto
// DO NOT EDIT!

/*
Package scheduler is a generated protocol buffer package.

It is generated from these files:
	scheduler/scheduler.proto

It has these top-level messages:
	ScheduleItem
	Id
	Empty
	BrokerList
*/
package scheduler

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

type Error int32

const (
	Error_OK              Error = 0
	Error_NOTFOUND        Error = 404
	Error_INVALID_REQUEST Error = 400
	Error_UNAUTHORIZED    Error = 401
	Error_INTERNAL_ERROR  Error = 500
)

var Error_name = map[int32]string{
	0:   "OK",
	404: "NOTFOUND",
	400: "INVALID_REQUEST",
	401: "UNAUTHORIZED",
	500: "INTERNAL_ERROR",
}
var Error_value = map[string]int32{
	"OK":              0,
	"NOTFOUND":        404,
	"INVALID_REQUEST": 400,
	"UNAUTHORIZED":    401,
	"INTERNAL_ERROR":  500,
}

func (x Error) String() string {
	return proto.EnumName(Error_name, int32(x))
}
func (Error) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ScheduleItem struct {
	Id           string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
	ScheduleTime string `protobuf:"bytes,2,opt,name=ScheduleTime,json=scheduleTime" json:"ScheduleTime,omitempty"`
	RunTime      string `protobuf:"bytes,3,opt,name=RunTime,json=runTime" json:"RunTime,omitempty"`
	Topic        string `protobuf:"bytes,4,opt,name=Topic,json=topic" json:"Topic,omitempty"`
	Data         []byte `protobuf:"bytes,5,opt,name=Data,json=data,proto3" json:"Data,omitempty"`
	Text         string `protobuf:"bytes,6,opt,name=Text,json=text" json:"Text,omitempty"`
}

func (m *ScheduleItem) Reset()                    { *m = ScheduleItem{} }
func (m *ScheduleItem) String() string            { return proto.CompactTextString(m) }
func (*ScheduleItem) ProtoMessage()               {}
func (*ScheduleItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ScheduleItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ScheduleItem) GetScheduleTime() string {
	if m != nil {
		return m.ScheduleTime
	}
	return ""
}

func (m *ScheduleItem) GetRunTime() string {
	if m != nil {
		return m.RunTime
	}
	return ""
}

func (m *ScheduleItem) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ScheduleItem) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ScheduleItem) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Id struct {
	Id string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type BrokerList struct {
	Brokers []string `protobuf:"bytes,1,rep,name=Brokers,json=brokers" json:"Brokers,omitempty"`
}

func (m *BrokerList) Reset()                    { *m = BrokerList{} }
func (m *BrokerList) String() string            { return proto.CompactTextString(m) }
func (*BrokerList) ProtoMessage()               {}
func (*BrokerList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BrokerList) GetBrokers() []string {
	if m != nil {
		return m.Brokers
	}
	return nil
}

func init() {
	proto.RegisterType((*ScheduleItem)(nil), "scheduler.ScheduleItem")
	proto.RegisterType((*Id)(nil), "scheduler.Id")
	proto.RegisterType((*Empty)(nil), "scheduler.Empty")
	proto.RegisterType((*BrokerList)(nil), "scheduler.BrokerList")
	proto.RegisterEnum("scheduler.Error", Error_name, Error_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Scheduler service

type SchedulerClient interface {
	Register(ctx context.Context, in *ScheduleItem, opts ...grpc.CallOption) (*Id, error)
	Cancel(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) Register(ctx context.Context, in *ScheduleItem, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := grpc.Invoke(ctx, "/scheduler.Scheduler/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) Cancel(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/scheduler.Scheduler/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	Register(context.Context, *ScheduleItem) (*Id, error)
	Cancel(context.Context, *Id) (*Empty, error)
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).Register(ctx, req.(*ScheduleItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).Cancel(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Scheduler_Register_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Scheduler_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduler/scheduler.proto",
}

func init() { proto.RegisterFile("scheduler/scheduler.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x91, 0xcd, 0x4a, 0xf3, 0x40,
	0x18, 0x85, 0x9b, 0xdf, 0xb6, 0x2f, 0x69, 0xbf, 0x7c, 0xaf, 0x05, 0x63, 0x57, 0x25, 0x0b, 0x29,
	0x0a, 0x15, 0xd4, 0x1b, 0xa8, 0x26, 0x62, 0xb0, 0x24, 0x38, 0x9d, 0xb8, 0x10, 0xa1, 0xa4, 0xcd,
	0xa0, 0xc1, 0xb6, 0x29, 0x93, 0x29, 0xd4, 0xbb, 0x50, 0x70, 0xe5, 0x75, 0x7a, 0x01, 0x92, 0x44,
	0xdb, 0xe8, 0x6e, 0x9e, 0xf3, 0xcc, 0xe2, 0xcc, 0x19, 0x38, 0xc8, 0x66, 0x4f, 0x2c, 0x5e, 0xcf,
	0x19, 0x3f, 0xd9, 0x9e, 0x06, 0x2b, 0x9e, 0x8a, 0x14, 0x9b, 0xdb, 0xc0, 0xfe, 0x90, 0xc0, 0x18,
	0x7f, 0x93, 0x27, 0xd8, 0x02, 0xdb, 0x20, 0x7b, 0xb1, 0x25, 0xf5, 0xa4, 0x7e, 0x93, 0xc8, 0x49,
	0x8c, 0xf6, 0xce, 0xd3, 0x64, 0xc1, 0x2c, 0xb9, 0x30, 0x46, 0x56, 0xc9, 0xd0, 0x82, 0x3a, 0x59,
	0x2f, 0x0b, 0xad, 0x14, 0xba, 0xce, 0x4b, 0xc4, 0x0e, 0x68, 0x34, 0x5d, 0x25, 0x33, 0x4b, 0x2d,
	0x72, 0x4d, 0xe4, 0x80, 0x08, 0xaa, 0x13, 0x89, 0xc8, 0xd2, 0x7a, 0x52, 0xdf, 0x20, 0x6a, 0x1c,
	0x89, 0x28, 0xcf, 0x28, 0xdb, 0x08, 0x4b, 0x2f, 0x2e, 0xaa, 0x82, 0x6d, 0x84, 0xdd, 0xc9, 0xbb,
	0xfc, 0x6d, 0x64, 0xd7, 0x41, 0x73, 0x17, 0x2b, 0xf1, 0x62, 0x1f, 0x02, 0x5c, 0xf0, 0xf4, 0x99,
	0xf1, 0x51, 0x92, 0x89, 0xbc, 0x44, 0x49, 0x99, 0x25, 0xf5, 0x94, 0xbc, 0xc4, 0xb4, 0xc4, 0xa3,
	0x07, 0xd0, 0x5c, 0xce, 0x53, 0x8e, 0x3a, 0xc8, 0xc1, 0x8d, 0x59, 0xc3, 0x16, 0x34, 0xfc, 0x80,
	0x5e, 0x05, 0xa1, 0xef, 0x98, 0xef, 0x0a, 0x76, 0xe0, 0x9f, 0xe7, 0xdf, 0x0d, 0x47, 0x9e, 0x33,
	0x21, 0xee, 0x6d, 0xe8, 0x8e, 0xa9, 0xf9, 0xaa, 0xe0, 0x7f, 0x30, 0x42, 0x7f, 0x18, 0xd2, 0xeb,
	0x80, 0x78, 0xf7, 0xae, 0x63, 0xbe, 0x29, 0xb8, 0x07, 0x6d, 0xcf, 0xa7, 0x2e, 0xf1, 0x87, 0xa3,
	0x89, 0x4b, 0x48, 0x40, 0xcc, 0x4f, 0xe5, 0x74, 0x09, 0xcd, 0x9f, 0x81, 0x38, 0x9e, 0x43, 0x83,
	0xb0, 0xc7, 0x24, 0x13, 0x8c, 0xe3, 0xfe, 0x60, 0xb7, 0x7b, 0x75, 0xe2, 0x6e, 0xab, 0x22, 0xbc,
	0xd8, 0xae, 0xe1, 0x31, 0xe8, 0x97, 0xd1, 0x72, 0xc6, 0xe6, 0xf8, 0x5b, 0x75, 0xcd, 0x0a, 0x96,
	0x6f, 0xae, 0x4d, 0xf5, 0xe2, 0x0f, 0xcf, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x7e, 0x44,
	0x49, 0xe0, 0x01, 0x00, 0x00,
}
