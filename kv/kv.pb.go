// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv/kv.proto

package kv

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Key struct {
	Partition            string   `protobuf:"bytes,2,opt,name=partition,proto3" json:"partition,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_5993384bbd56d931, []int{0}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

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
	Partition            string   `protobuf:"bytes,2,opt,name=partition,proto3" json:"partition,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Bytes                []byte   `protobuf:"bytes,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Value                string   `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Ttl                  int64    `protobuf:"varint,6,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_5993384bbd56d931, []int{1}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

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

func (m *Value) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *Value) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Value) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type Bool struct {
	Value                bool     `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bool) Reset()         { *m = Bool{} }
func (m *Bool) String() string { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()    {}
func (*Bool) Descriptor() ([]byte, []int) {
	return fileDescriptor_5993384bbd56d931, []int{2}
}

func (m *Bool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bool.Unmarshal(m, b)
}
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bool.Marshal(b, m, deterministic)
}
func (m *Bool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bool.Merge(m, src)
}
func (m *Bool) XXX_Size() int {
	return xxx_messageInfo_Bool.Size(m)
}
func (m *Bool) XXX_DiscardUnknown() {
	xxx_messageInfo_Bool.DiscardUnknown(m)
}

var xxx_messageInfo_Bool proto.InternalMessageInfo

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

func init() { proto.RegisterFile("kv/kv.proto", fileDescriptor_5993384bbd56d931) }

var fileDescriptor_5993384bbd56d931 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2e, 0xd3, 0xcf,
	0x2e, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xca, 0x2e, 0x53, 0x32, 0xe5, 0x62, 0xf6,
	0x4e, 0xad, 0x14, 0x92, 0xe1, 0xe2, 0x2c, 0x48, 0x2c, 0x2a, 0xc9, 0x2c, 0xc9, 0xcc, 0xcf, 0x93,
	0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x08, 0x08, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x4a,
	0x30, 0x83, 0xc5, 0x41, 0x4c, 0xa5, 0x52, 0x2e, 0xd6, 0xb0, 0xc4, 0x9c, 0xd2, 0x54, 0x52, 0x35,
	0x0a, 0x89, 0x70, 0xb1, 0x26, 0x55, 0x96, 0xa4, 0x16, 0x4b, 0xb0, 0x28, 0x30, 0x6a, 0xf0, 0x04,
	0x41, 0x38, 0x20, 0xd1, 0x32, 0x90, 0x71, 0x12, 0xac, 0x60, 0x95, 0x10, 0x0e, 0x48, 0x77, 0x49,
	0x49, 0x8e, 0x04, 0x9b, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x88, 0xa9, 0x24, 0xc3, 0xc5, 0xe2, 0x94,
	0x9f, 0x9f, 0x83, 0x50, 0x0f, 0xb2, 0x91, 0x03, 0xaa, 0x3e, 0x89, 0x0d, 0xec, 0x2d, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xf0, 0xc7, 0x7e, 0xe5, 0x00, 0x00, 0x00,
}
