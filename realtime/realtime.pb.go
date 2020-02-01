// Code generated by protoc-gen-go. DO NOT EDIT.
// source: realtime.proto

package realtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/subiz/header/common"
	event "github.com/subiz/header/event"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Meta_Type int32

const (
	Meta_tombstone Meta_Type = 0
)

var Meta_Type_name = map[int32]string{
	0: "tombstone",
}

var Meta_Type_value = map[string]int32{
	"tombstone": 0,
}

func (x Meta_Type) String() string {
	return proto.EnumName(Meta_Type_name, int32(x))
}

func (Meta_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{2, 0}
}

type Subscription struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	InitialToken         string          `protobuf:"bytes,2,opt,name=initial_token,json=initialToken,proto3" json:"initial_token,omitempty"`
	Events               []string        `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{0}
}

func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Subscription) GetInitialToken() string {
	if m != nil {
		return m.InitialToken
	}
	return ""
}

func (m *Subscription) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

type PollResult struct {
	Events               []*event.Event `protobuf:"bytes,4,rep,name=events,proto3" json:"events,omitempty"`
	SequentialToken      string         `protobuf:"bytes,6,opt,name=sequential_token,json=sequentialToken,proto3" json:"sequential_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PollResult) Reset()         { *m = PollResult{} }
func (m *PollResult) String() string { return proto.CompactTextString(m) }
func (*PollResult) ProtoMessage()    {}
func (*PollResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{1}
}

func (m *PollResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollResult.Unmarshal(m, b)
}
func (m *PollResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollResult.Marshal(b, m, deterministic)
}
func (m *PollResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollResult.Merge(m, src)
}
func (m *PollResult) XXX_Size() int {
	return xxx_messageInfo_PollResult.Size(m)
}
func (m *PollResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PollResult.DiscardUnknown(m)
}

var xxx_messageInfo_PollResult proto.InternalMessageInfo

func (m *PollResult) GetEvents() []*event.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *PollResult) GetSequentialToken() string {
	if m != nil {
		return m.SequentialToken
	}
	return ""
}

type Meta struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Type                 string          `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64           `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Expired              int64           `protobuf:"varint,4,opt,name=expired,proto3" json:"expired,omitempty"`
	ConnectionId         string          `protobuf:"bytes,10,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{2}
}

func (m *Meta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meta.Unmarshal(m, b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meta.Marshal(b, m, deterministic)
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return xxx_messageInfo_Meta.Size(m)
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

func (m *Meta) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Meta) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Meta) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Meta) GetExpired() int64 {
	if m != nil {
		return m.Expired
	}
	return 0
}

func (m *Meta) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

type Token struct {
	ConnectionId         string   `protobuf:"bytes,4,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	LastOffset           int64    `protobuf:"varint,5,opt,name=last_offset,json=lastOffset,proto3" json:"last_offset,omitempty"`
	Created              int64    `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"`
	Expired              int64    `protobuf:"varint,7,opt,name=expired,proto3" json:"expired,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Token) GetLastOffset() int64 {
	if m != nil {
		return m.LastOffset
	}
	return 0
}

func (m *Token) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Token) GetExpired() int64 {
	if m != nil {
		return m.Expired
	}
	return 0
}

func init() {
	proto.RegisterEnum("realtime.Meta_Type", Meta_Type_name, Meta_Type_value)
	proto.RegisterType((*Subscription)(nil), "realtime.Subscription")
	proto.RegisterType((*PollResult)(nil), "realtime.PollResult")
	proto.RegisterType((*Meta)(nil), "realtime.Meta")
	proto.RegisterType((*Token)(nil), "realtime.Token")
}

func init() { proto.RegisterFile("realtime.proto", fileDescriptor_dcbdca058206953b) }

var fileDescriptor_dcbdca058206953b = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xad, 0x2a, 0x59, 0xae, 0x47, 0x72, 0x6d, 0x16, 0x5a, 0x16, 0x5f, 0xaa, 0xaa, 0x3d, 0xa8,
	0x17, 0x1b, 0xdc, 0x7f, 0xd0, 0xd2, 0x43, 0x0f, 0xa5, 0x45, 0xf5, 0xa9, 0x50, 0x8c, 0x3e, 0xc6,
	0xf1, 0x12, 0x69, 0x57, 0xd1, 0x8e, 0x82, 0x93, 0x5b, 0x7e, 0x50, 0xfe, 0x63, 0x58, 0x7d, 0xc4,
	0x0e, 0x26, 0x90, 0xcb, 0xb2, 0xef, 0xcd, 0xce, 0xbc, 0xb7, 0x8f, 0x81, 0xb7, 0x35, 0x26, 0x05,
	0x89, 0x12, 0x97, 0x55, 0xad, 0x48, 0xb1, 0x37, 0x03, 0x5e, 0xf8, 0x99, 0x2a, 0x4b, 0x25, 0x3b,
	0x7e, 0xe1, 0xe1, 0x35, 0x4a, 0xea, 0x40, 0x28, 0xc1, 0xff, 0xdb, 0xa4, 0x3a, 0xab, 0x45, 0x45,
	0x42, 0x49, 0xf6, 0x11, 0xec, 0x8c, 0x0e, 0xdc, 0x0a, 0xac, 0xc8, 0x5b, 0xcf, 0x96, 0x7d, 0xe3,
	0x77, 0x25, 0x09, 0x0f, 0x14, 0x9b, 0x1a, 0xfb, 0x04, 0x53, 0x21, 0x05, 0x89, 0xa4, 0xd8, 0x92,
	0xba, 0x44, 0xc9, 0x5f, 0x07, 0x56, 0x34, 0x89, 0xfd, 0x9e, 0xdc, 0x18, 0x8e, 0xbd, 0x07, 0xb7,
	0x95, 0xd1, 0xdc, 0x0e, 0xec, 0x68, 0x12, 0xf7, 0x28, 0xfc, 0x0f, 0xf0, 0x47, 0x15, 0x45, 0x8c,
	0xba, 0x29, 0x88, 0x7d, 0x7e, 0x7c, 0xe5, 0x04, 0x76, 0xe4, 0xad, 0xfd, 0x65, 0xe7, 0xed, 0x87,
	0x39, 0x87, 0x1e, 0xf6, 0x05, 0xe6, 0x1a, 0xaf, 0x1a, 0x94, 0x27, 0x9a, 0x6e, 0xab, 0x39, 0x3b,
	0xf2, 0xad, 0x6c, 0x78, 0x6f, 0x81, 0xf3, 0x0b, 0x29, 0x79, 0xc9, 0x3f, 0x18, 0x38, 0x74, 0x53,
	0x61, 0x6f, 0xbf, 0xbd, 0x33, 0x0e, 0xe3, 0xac, 0xc6, 0x84, 0x30, 0xe7, 0x76, 0x60, 0x45, 0x76,
	0x3c, 0x40, 0x53, 0xc1, 0x43, 0x25, 0x6a, 0xcc, 0xb9, 0xd3, 0x55, 0x7a, 0x68, 0xf2, 0xc8, 0x94,
	0x94, 0x98, 0x99, 0x00, 0xb7, 0x22, 0xe7, 0xd0, 0xe5, 0x71, 0x24, 0x7f, 0xe6, 0xe1, 0x3b, 0x70,
	0x36, 0x46, 0x60, 0x0a, 0x13, 0x52, 0x65, 0xaa, 0x49, 0x49, 0x9c, 0xbf, 0x0a, 0xef, 0x2c, 0x18,
	0x75, 0x81, 0x9d, 0x4d, 0x71, 0xce, 0xa7, 0xb0, 0x0f, 0xe0, 0x15, 0x89, 0xa6, 0xad, 0xda, 0xed,
	0x34, 0x12, 0x1f, 0xb5, 0x46, 0xc0, 0x50, 0xbf, 0x5b, 0xe6, 0xd4, 0xbf, 0xfb, 0xac, 0xff, 0xf1,
	0x13, 0xff, 0xdf, 0xc2, 0x7f, 0xc1, 0x85, 0xa0, 0x7d, 0x93, 0x9a, 0x94, 0x56, 0xba, 0x49, 0xc5,
	0xed, 0x6a, 0x8f, 0x49, 0x8e, 0xf5, 0x6a, 0xd8, 0xa0, 0xd4, 0x6d, 0xb7, 0xe5, 0xeb, 0x43, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xcd, 0xb4, 0x5c, 0x5b, 0x64, 0x02, 0x00, 0x00,
}
