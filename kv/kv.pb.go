// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv/kv.proto

/*
Package kv is a generated protocol buffer package.

It is generated from these files:
	kv/kv.proto

It has these top-level messages:
	Key
	Value
	Bool
*/
package kv

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

type Key struct {
	Partition string `protobuf:"bytes,2,opt,name=partition" json:"partition,omitempty"`
	Key       string `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Key) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func (m *Key) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Value struct {
	Partition string `protobuf:"bytes,2,opt,name=partition" json:"partition,omitempty"`
	Key       string `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Blob      []byte `protobuf:"bytes,4,opt,name=blob,proto3" json:"blob,omitempty"`
	Value     string `protobuf:"bytes,5,opt,name=value" json:"value,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Value) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func (m *Value) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Value) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

func (m *Value) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Bool struct {
	Value bool `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *Bool) Reset()                    { *m = Bool{} }
func (m *Bool) String() string            { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()               {}
func (*Bool) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*Key)(nil), "kv.Key")
	proto.RegisterType((*Value)(nil), "kv.Value")
	proto.RegisterType((*Bool)(nil), "kv.Bool")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KV service

type KVClient interface {
	Set(ctx context.Context, in *Value, opts ...grpc.CallOption) (*Value, error)
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	Has(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Bool, error)
}

type kVClient struct {
	cc *grpc.ClientConn
}

func NewKVClient(cc *grpc.ClientConn) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) Set(ctx context.Context, in *Value, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := grpc.Invoke(ctx, "/kv.KV/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := grpc.Invoke(ctx, "/kv.KV/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Has(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := grpc.Invoke(ctx, "/kv.KV/Has", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KV service

type KVServer interface {
	Set(context.Context, *Value) (*Value, error)
	Get(context.Context, *Key) (*Value, error)
	Has(context.Context, *Key) (*Bool, error)
}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	s.RegisterService(&_KV_serviceDesc, srv)
}

func _KV_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Set(ctx, req.(*Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Has_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Has(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kv.KV/Has",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Has(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

var _KV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kv.KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _KV_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KV_Get_Handler,
		},
		{
			MethodName: "Has",
			Handler:    _KV_Has_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kv/kv.proto",
}

func init() { proto.RegisterFile("kv/kv.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x31, 0xcb, 0xc2, 0x30,
	0x10, 0x86, 0x69, 0xd3, 0x7e, 0x5f, 0x7b, 0x3a, 0xc8, 0xe1, 0x10, 0xb5, 0x43, 0xe9, 0xd4, 0xa9,
	0x82, 0xe2, 0x1f, 0x70, 0x51, 0xe8, 0x56, 0xa1, 0x7b, 0x0a, 0x19, 0x4a, 0x8a, 0x29, 0x35, 0x06,
	0xfa, 0xef, 0xe5, 0xa2, 0x10, 0x57, 0xb7, 0x27, 0xef, 0x93, 0xbc, 0x39, 0x0e, 0x16, 0xca, 0xee,
	0x95, 0xad, 0xc6, 0x49, 0x1b, 0x8d, 0xa1, 0xb2, 0xc5, 0x09, 0x58, 0x2d, 0x67, 0xcc, 0x20, 0x1d,
	0xc5, 0x64, 0x7a, 0xd3, 0xeb, 0x3b, 0x0f, 0xf3, 0xa0, 0x4c, 0x1b, 0x1f, 0xe0, 0x0a, 0x98, 0x92,
	0x33, 0x67, 0x2e, 0x27, 0x2c, 0x04, 0xc4, 0xad, 0x18, 0x9e, 0xf2, 0xd7, 0x87, 0x88, 0x10, 0x75,
	0x83, 0xee, 0x78, 0x94, 0x07, 0xe5, 0xb2, 0x71, 0x8c, 0x6b, 0x88, 0x2d, 0x95, 0xf1, 0xd8, 0xdd,
	0x7b, 0x1f, 0x8a, 0x0c, 0xa2, 0xb3, 0xd6, 0x83, 0xb7, 0xd4, 0x9e, 0x7c, 0xec, 0xa1, 0x85, 0xb0,
	0x6e, 0x71, 0x07, 0xec, 0x26, 0x0d, 0xa6, 0x95, 0xb2, 0x95, 0x9b, 0x67, 0xeb, 0x11, 0x37, 0xc0,
	0x2e, 0xd2, 0xe0, 0x3f, 0x25, 0xb5, 0x9c, 0xbf, 0x15, 0x07, 0x76, 0x15, 0x0f, 0xaf, 0x12, 0x02,
	0xfa, 0xad, 0xfb, 0x73, 0xab, 0x39, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xda, 0x14, 0xd3, 0xee,
	0x29, 0x01, 0x00, 0x00,
}
