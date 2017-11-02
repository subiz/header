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
	WlDomain
	WlIP
	WlUser
	ListRequest
	ScryptChallenge
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bitbucket.org/subiz/header/common"

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

type WlDomain struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Domain           *string         `protobuf:"bytes,4,opt,name=domain" json:"domain,omitempty"`
	Created          *int64          `protobuf:"varint,5,opt,name=created" json:"created,omitempty"`
	By               *string         `protobuf:"bytes,6,opt,name=by" json:"by,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *WlDomain) Reset()                    { *m = WlDomain{} }
func (m *WlDomain) String() string            { return proto.CompactTextString(m) }
func (*WlDomain) ProtoMessage()               {}
func (*WlDomain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WlDomain) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *WlDomain) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *WlDomain) GetDomain() string {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return ""
}

func (m *WlDomain) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *WlDomain) GetBy() string {
	if m != nil && m.By != nil {
		return *m.By
	}
	return ""
}

type WlIP struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Ip               *string         `protobuf:"bytes,4,opt,name=ip" json:"ip,omitempty"`
	Created          *int64          `protobuf:"varint,5,opt,name=created" json:"created,omitempty"`
	By               *string         `protobuf:"bytes,6,opt,name=by" json:"by,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *WlIP) Reset()                    { *m = WlIP{} }
func (m *WlIP) String() string            { return proto.CompactTextString(m) }
func (*WlIP) ProtoMessage()               {}
func (*WlIP) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WlIP) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *WlIP) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *WlIP) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *WlIP) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *WlIP) GetBy() string {
	if m != nil && m.By != nil {
		return *m.By
	}
	return ""
}

type WlUser struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId           *string         `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Created          *int64          `protobuf:"varint,5,opt,name=created" json:"created,omitempty"`
	By               *string         `protobuf:"bytes,6,opt,name=by" json:"by,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *WlUser) Reset()                    { *m = WlUser{} }
func (m *WlUser) String() string            { return proto.CompactTextString(m) }
func (*WlUser) ProtoMessage()               {}
func (*WlUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WlUser) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *WlUser) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *WlUser) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *WlUser) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *WlUser) GetBy() string {
	if m != nil && m.By != nil {
		return *m.By
	}
	return ""
}

type ListRequest struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Start            *string         `protobuf:"bytes,3,opt,name=start" json:"start,omitempty"`
	Limit            *int32          `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ListRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *ListRequest) GetStart() string {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return ""
}

func (m *ListRequest) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

type ScryptChallenge struct {
	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId *string         `protobuf:"bytes,9,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Id        *string         `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	N         *int32          `protobuf:"varint,3,opt,name=n" json:"n,omitempty"`
	P         *int32          `protobuf:"varint,4,opt,name=p" json:"p,omitempty"`
	R         *int32          `protobuf:"varint,5,opt,name=r" json:"r,omitempty"`
	Salt      *string         `protobuf:"bytes,6,opt,name=salt" json:"salt,omitempty"`
	Hash      *string         `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	// only for subiz use
	Answer           *int32 `protobuf:"varint,10,opt,name=answer" json:"answer,omitempty"`
	Created          *int64 `protobuf:"varint,11,opt,name=created" json:"created,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ScryptChallenge) Reset()                    { *m = ScryptChallenge{} }
func (m *ScryptChallenge) String() string            { return proto.CompactTextString(m) }
func (*ScryptChallenge) ProtoMessage()               {}
func (*ScryptChallenge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ScryptChallenge) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ScryptChallenge) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *ScryptChallenge) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *ScryptChallenge) GetN() int32 {
	if m != nil && m.N != nil {
		return *m.N
	}
	return 0
}

func (m *ScryptChallenge) GetP() int32 {
	if m != nil && m.P != nil {
		return *m.P
	}
	return 0
}

func (m *ScryptChallenge) GetR() int32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *ScryptChallenge) GetSalt() string {
	if m != nil && m.Salt != nil {
		return *m.Salt
	}
	return ""
}

func (m *ScryptChallenge) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func (m *ScryptChallenge) GetAnswer() int32 {
	if m != nil && m.Answer != nil {
		return *m.Answer
	}
	return 0
}

func (m *ScryptChallenge) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*SendMessageState)(nil), "api.SendMessageState")
	proto.RegisterType((*State)(nil), "api.State")
	proto.RegisterType((*WlDomain)(nil), "api.WlDomain")
	proto.RegisterType((*WlIP)(nil), "api.WlIP")
	proto.RegisterType((*WlUser)(nil), "api.WlUser")
	proto.RegisterType((*ListRequest)(nil), "api.ListRequest")
	proto.RegisterType((*ScryptChallenge)(nil), "api.ScryptChallenge")
	proto.RegisterEnum("api.Method", Method_name, Method_value)
	proto.RegisterEnum("api.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x71, 0x7e, 0xf4, 0xc7, 0xeb, 0xd8, 0x82, 0x35, 0x20, 0x42, 0x9a, 0x34, 0x72, 0xaa,
	0x26, 0xd4, 0x4a, 0x5c, 0x39, 0x4d, 0x5d, 0x04, 0x95, 0x36, 0x51, 0xa5, 0x9d, 0x7a, 0x44, 0x6e,
	0xf2, 0xd4, 0x5a, 0xa4, 0xb6, 0xb1, 0xdd, 0xb1, 0xc2, 0x8d, 0x3b, 0xfc, 0x9b, 0xfc, 0x1b, 0xc8,
	0x6e, 0x36, 0x01, 0xda, 0xa5, 0xd2, 0x4e, 0xf1, 0xe7, 0xbd, 0x3c, 0xe7, 0xa3, 0x6f, 0x6c, 0x78,
	0xca, 0x14, 0x1f, 0x32, 0xc5, 0x07, 0x4a, 0x4b, 0x2b, 0x69, 0xc8, 0x14, 0x7f, 0x35, 0x58, 0x70,
	0xbb, 0xd8, 0x94, 0x9f, 0xd1, 0x0e, 0xa4, 0x5e, 0x0e, 0xcd, 0x66, 0xc1, 0xbf, 0x0d, 0x57, 0xc8,
	0x2a, 0xd4, 0xc3, 0x52, 0xae, 0xd7, 0x52, 0x34, 0x8f, 0xdd, 0x50, 0xd6, 0x86, 0x38, 0x5f, 0x2b,
	0xbb, 0xcd, 0xfa, 0x90, 0x4c, 0x51, 0x54, 0x57, 0x68, 0x0c, 0x5b, 0xe2, 0xd4, 0x32, 0x8b, 0xf4,
	0x18, 0x62, 0x2b, 0x15, 0x2f, 0x53, 0x72, 0x4a, 0xfa, 0xdd, 0x62, 0x07, 0xd9, 0x09, 0xc4, 0xf7,
	0x6d, 0xbc, 0x41, 0x61, 0xd3, 0x70, 0xd7, 0xf6, 0x90, 0xfd, 0x22, 0xd0, 0x99, 0xd7, 0x17, 0x72,
	0xcd, 0xb8, 0xa0, 0xaf, 0x21, 0x2c, 0xed, 0xad, 0x9f, 0xef, 0xbd, 0x3d, 0x1a, 0x34, 0x9f, 0x1e,
	0x49, 0x61, 0xf1, 0xd6, 0x16, 0xae, 0x47, 0x4f, 0x00, 0x58, 0x59, 0xca, 0x8d, 0xb0, 0x9f, 0x78,
	0xd5, 0x6c, 0xd5, 0x6d, 0x2a, 0xe3, 0x8a, 0xbe, 0x80, 0x56, 0xe5, 0xf7, 0x4a, 0x23, 0xdf, 0x6a,
	0x88, 0xa6, 0xd0, 0x2e, 0x35, 0x32, 0x8b, 0x55, 0x1a, 0x9f, 0x92, 0x7e, 0x58, 0xdc, 0x21, 0x3d,
	0x84, 0x60, 0xb1, 0x4d, 0x5b, 0xfe, 0xed, 0x60, 0xb1, 0xcd, 0x7e, 0x10, 0x88, 0xe6, 0xf5, 0x78,
	0xf2, 0x08, 0x32, 0x87, 0x10, 0x70, 0xd5, 0x88, 0x04, 0x5c, 0xed, 0x21, 0xf1, 0x93, 0x40, 0x6b,
	0x5e, 0x5f, 0x1b, 0xd4, 0x8f, 0xa0, 0xf1, 0x12, 0xda, 0x1b, 0x83, 0xda, 0xf5, 0x9a, 0x50, 0x1c,
	0x8e, 0xab, 0x3d, 0x7c, 0xbe, 0x43, 0xef, 0x92, 0x1b, 0x5b, 0xe0, 0x97, 0x0d, 0x1a, 0xbb, 0xbf,
	0x53, 0xf0, 0xbf, 0xd3, 0x31, 0xc4, 0xc6, 0x32, 0x7d, 0x7f, 0x18, 0x3c, 0xb8, 0x6a, 0xcd, 0xd7,
	0xdc, 0x7a, 0xcf, 0xb8, 0xd8, 0x41, 0xf6, 0x9b, 0xc0, 0xd1, 0xb4, 0xd4, 0x5b, 0x65, 0x47, 0x2b,
	0x56, 0xd7, 0x28, 0x96, 0xb8, 0xbf, 0x41, 0xf7, 0xa1, 0x9f, 0x73, 0x27, 0x16, 0xf0, 0x8a, 0x1e,
	0x00, 0x11, 0xde, 0x26, 0x2e, 0x88, 0x70, 0xa4, 0x1a, 0x0b, 0xa2, 0x1c, 0x69, 0x1f, 0x51, 0x5c,
	0x10, 0x4d, 0x29, 0x44, 0x86, 0xd5, 0xb6, 0x89, 0xc7, 0xaf, 0x5d, 0x6d, 0xc5, 0xcc, 0x2a, 0xed,
	0xec, 0x6a, 0x6e, 0xed, 0xce, 0x22, 0x13, 0xe6, 0x2b, 0xea, 0x14, 0xfc, 0x68, 0x43, 0x7f, 0xc7,
	0xde, 0xfb, 0x27, 0xf6, 0xb3, 0x77, 0xd0, 0xba, 0x42, 0xbb, 0x92, 0x15, 0x6d, 0x43, 0xf8, 0x3e,
	0x9f, 0x25, 0x4f, 0x68, 0x07, 0xa2, 0xc9, 0xc7, 0xe9, 0x2c, 0x21, 0xae, 0x34, 0xb9, 0x9e, 0x25,
	0x01, 0x05, 0x68, 0x5d, 0xe4, 0x97, 0xf9, 0x2c, 0x4f, 0x42, 0xda, 0x85, 0x78, 0x72, 0x3e, 0x1b,
	0x7d, 0x48, 0xa2, 0xb3, 0x37, 0x10, 0xe7, 0xee, 0x4a, 0xd1, 0xe7, 0xf0, 0x6c, 0x24, 0xc5, 0x0d,
	0x6a, 0xc3, 0x2c, 0x97, 0x62, 0xea, 0xa2, 0x4d, 0x02, 0x7a, 0x00, 0x9d, 0x73, 0xc5, 0x0b, 0x54,
	0xf5, 0x36, 0x89, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x61, 0x9c, 0x17, 0x63, 0x0e, 0x04, 0x00,
	0x00,
}
