// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api/api.proto

It has these top-level messages:
	Empty
	AllType
	SendMessageState
	State
	WlDomain
	WlIP
	WlUser
	ListRequest
	ScryptChallenge
	Apikey
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bitbucket.org/subiz/header/common"

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
	Event_ApiRequested           Event = 0
	Event_ConversationStart      Event = 2
	Event_ApiReply               Event = 4
	Event_ApiWhiteListSynced     Event = 5
	Event_DomainWhiteListCreated Event = 6
	Event_DomainWhiteListDeleted Event = 7
	Event_UserWhiteListCreated   Event = 8
	Event_UserWhiteListDeleted   Event = 9
	Event_IpWhiteListCreated     Event = 10
	Event_IpWhiteListDeleted     Event = 11
)

var Event_name = map[int32]string{
	0:  "ApiRequested",
	2:  "ConversationStart",
	4:  "ApiReply",
	5:  "ApiWhiteListSynced",
	6:  "DomainWhiteListCreated",
	7:  "DomainWhiteListDeleted",
	8:  "UserWhiteListCreated",
	9:  "UserWhiteListDeleted",
	10: "IpWhiteListCreated",
	11: "IpWhiteListDeleted",
}
var Event_value = map[string]int32{
	"ApiRequested":           0,
	"ConversationStart":      2,
	"ApiReply":               4,
	"ApiWhiteListSynced":     5,
	"DomainWhiteListCreated": 6,
	"DomainWhiteListDeleted": 7,
	"UserWhiteListCreated":   8,
	"UserWhiteListDeleted":   9,
	"IpWhiteListCreated":     10,
	"IpWhiteListDeleted":     11,
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

type AllType struct {
	Domain           *WlDomain        `protobuf:"bytes,2,opt,name=domain" json:"domain,omitempty"`
	Wlip             *WlIP            `protobuf:"bytes,3,opt,name=wlip" json:"wlip,omitempty"`
	Wluse            *WlUser          `protobuf:"bytes,4,opt,name=wluse" json:"wluse,omitempty"`
	Lr               *ListRequest     `protobuf:"bytes,5,opt,name=lr" json:"lr,omitempty"`
	Sc               *ScryptChallenge `protobuf:"bytes,6,opt,name=sc" json:"sc,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *AllType) Reset()                    { *m = AllType{} }
func (m *AllType) String() string            { return proto.CompactTextString(m) }
func (*AllType) ProtoMessage()               {}
func (*AllType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AllType) GetDomain() *WlDomain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *AllType) GetWlip() *WlIP {
	if m != nil {
		return m.Wlip
	}
	return nil
}

func (m *AllType) GetWluse() *WlUser {
	if m != nil {
		return m.Wluse
	}
	return nil
}

func (m *AllType) GetLr() *ListRequest {
	if m != nil {
		return m.Lr
	}
	return nil
}

func (m *AllType) GetSc() *ScryptChallenge {
	if m != nil {
		return m.Sc
	}
	return nil
}

type SendMessageState struct {
	Topic            *string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SendMessageState) Reset()                    { *m = SendMessageState{} }
func (m *SendMessageState) String() string            { return proto.CompactTextString(m) }
func (*SendMessageState) ProtoMessage()               {}
func (*SendMessageState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
func (*WlDomain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
func (*WlIP) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
func (*WlUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

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
	N         *int32          `protobuf:"varint,3,opt,name=N" json:"N,omitempty"`
	P         *int32          `protobuf:"varint,4,opt,name=P" json:"P,omitempty"`
	R         *int32          `protobuf:"varint,5,opt,name=r" json:"r,omitempty"`
	Salt      *string         `protobuf:"bytes,6,opt,name=salt" json:"salt,omitempty"`
	Hash      *string         `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	Domain    *int32          `protobuf:"varint,14,opt,name=domain" json:"domain,omitempty"`
	Dklen     *int32          `protobuf:"varint,15,opt,name=dklen" json:"dklen,omitempty"`
	// only for subiz use
	Answer           *string `protobuf:"bytes,10,opt,name=answer" json:"answer,omitempty"`
	Created          *int64  `protobuf:"varint,11,opt,name=created" json:"created,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ScryptChallenge) Reset()                    { *m = ScryptChallenge{} }
func (m *ScryptChallenge) String() string            { return proto.CompactTextString(m) }
func (*ScryptChallenge) ProtoMessage()               {}
func (*ScryptChallenge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

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

func (m *ScryptChallenge) GetDomain() int32 {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return 0
}

func (m *ScryptChallenge) GetDklen() int32 {
	if m != nil && m.Dklen != nil {
		return *m.Dklen
	}
	return 0
}

func (m *ScryptChallenge) GetAnswer() string {
	if m != nil && m.Answer != nil {
		return *m.Answer
	}
	return ""
}

func (m *ScryptChallenge) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

type Apikey struct {
	Secret           *string `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
	AccountId        *string `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ClientId         *string `protobuf:"bytes,5,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	ClientType       *string `protobuf:"bytes,6,opt,name=client_type,json=clientType" json:"client_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Apikey) Reset()                    { *m = Apikey{} }
func (m *Apikey) String() string            { return proto.CompactTextString(m) }
func (*Apikey) ProtoMessage()               {}
func (*Apikey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Apikey) GetSecret() string {
	if m != nil && m.Secret != nil {
		return *m.Secret
	}
	return ""
}

func (m *Apikey) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Apikey) GetClientId() string {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return ""
}

func (m *Apikey) GetClientType() string {
	if m != nil && m.ClientType != nil {
		return *m.ClientType
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*AllType)(nil), "api.AllType")
	proto.RegisterType((*SendMessageState)(nil), "api.SendMessageState")
	proto.RegisterType((*State)(nil), "api.State")
	proto.RegisterType((*WlDomain)(nil), "api.WlDomain")
	proto.RegisterType((*WlIP)(nil), "api.WlIP")
	proto.RegisterType((*WlUser)(nil), "api.WlUser")
	proto.RegisterType((*ListRequest)(nil), "api.ListRequest")
	proto.RegisterType((*ScryptChallenge)(nil), "api.ScryptChallenge")
	proto.RegisterType((*Apikey)(nil), "api.Apikey")
	proto.RegisterEnum("api.Method", Method_name, Method_value)
	proto.RegisterEnum("api.Event", Event_name, Event_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for X service

type XClient interface {
	Do(ctx context.Context, in *AllType, opts ...grpc.CallOption) (*AllType, error)
}

type xClient struct {
	cc *grpc.ClientConn
}

func NewXClient(cc *grpc.ClientConn) XClient {
	return &xClient{cc}
}

func (c *xClient) Do(ctx context.Context, in *AllType, opts ...grpc.CallOption) (*AllType, error) {
	out := new(AllType)
	err := grpc.Invoke(ctx, "/api.X/Do", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for X service

type XServer interface {
	Do(context.Context, *AllType) (*AllType, error)
}

func RegisterXServer(s *grpc.Server, srv XServer) {
	s.RegisterService(&_X_serviceDesc, srv)
}

func _X_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.X/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XServer).Do(ctx, req.(*AllType))
	}
	return interceptor(ctx, in, info, handler)
}

var _X_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.X",
	HandlerType: (*XServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _X_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xea, 0x46,
	0x14, 0x8d, 0x8d, 0x0d, 0xf8, 0xc2, 0x4b, 0xdc, 0x11, 0x4d, 0xad, 0x54, 0x51, 0xf3, 0xac, 0x56,
	0x8d, 0xde, 0x02, 0xa4, 0x6c, 0xbb, 0x42, 0x80, 0x5a, 0xa4, 0x24, 0x45, 0x86, 0x28, 0xdd, 0x55,
	0xc6, 0xbe, 0x0a, 0xa3, 0x18, 0x7b, 0xea, 0x19, 0x92, 0xb8, 0x55, 0x37, 0xdd, 0xb7, 0xbb, 0xfe,
	0x99, 0xfe, 0xac, 0xfe, 0x82, 0xa7, 0xf9, 0x30, 0x4a, 0x48, 0x36, 0x48, 0x59, 0xe1, 0x73, 0xce,
	0xbd, 0xf8, 0xcc, 0x99, 0xeb, 0x0b, 0x1f, 0x62, 0x46, 0x07, 0x31, 0xa3, 0x7d, 0x56, 0x16, 0xa2,
	0x20, 0x8d, 0x98, 0xd1, 0x93, 0xfe, 0x92, 0x8a, 0xe5, 0x26, 0xb9, 0x47, 0xd1, 0x2f, 0xca, 0xbb,
	0x01, 0xdf, 0x2c, 0xe9, 0xef, 0x83, 0x15, 0xc6, 0x29, 0x96, 0x83, 0xa4, 0x58, 0xaf, 0x8b, 0xdc,
	0xfc, 0xe8, 0xa6, 0xb0, 0x05, 0xee, 0x64, 0xcd, 0x44, 0x15, 0xfe, 0x67, 0x41, 0x6b, 0x98, 0x65,
	0x8b, 0x8a, 0x21, 0xf9, 0x0e, 0x9a, 0x69, 0xb1, 0x8e, 0x69, 0x1e, 0xd8, 0x67, 0xd6, 0x79, 0xe7,
	0xe2, 0x43, 0x5f, 0xbe, 0xe5, 0x36, 0x1b, 0x2b, 0x32, 0x32, 0x22, 0x39, 0x05, 0xe7, 0x31, 0xa3,
	0x2c, 0x68, 0xa8, 0x22, 0xcf, 0x14, 0x4d, 0x67, 0x91, 0xa2, 0xc9, 0x47, 0x70, 0x1f, 0xb3, 0x0d,
	0xc7, 0xc0, 0x51, 0x7a, 0xc7, 0xe8, 0x37, 0x1c, 0xcb, 0x48, 0x2b, 0xe4, 0x0c, 0xec, 0xac, 0x0c,
	0x5c, 0xa5, 0xfb, 0x4a, 0xbf, 0xa4, 0x5c, 0x44, 0xf8, 0xdb, 0x06, 0xb9, 0x88, 0xec, 0xac, 0x24,
	0xdf, 0x82, 0xcd, 0x93, 0xa0, 0xa9, 0x2a, 0x7a, 0xaa, 0x62, 0x9e, 0x94, 0x15, 0x13, 0xa3, 0x55,
	0x9c, 0x65, 0x98, 0xdf, 0x61, 0x64, 0xf3, 0x24, 0x3c, 0x07, 0x7f, 0x8e, 0x79, 0x7a, 0x85, 0x9c,
	0xc7, 0x77, 0x38, 0x17, 0xb1, 0x40, 0xd2, 0x03, 0x57, 0x14, 0x8c, 0x26, 0x81, 0x75, 0x66, 0x9d,
	0x7b, 0x91, 0x06, 0xe1, 0x29, 0xb8, 0x5b, 0x19, 0x1f, 0x30, 0x17, 0xca, 0xbd, 0x17, 0x69, 0x10,
	0xfe, 0x63, 0x41, 0xbb, 0x3e, 0x27, 0xf9, 0x08, 0x8d, 0x44, 0x3c, 0xa9, 0xfe, 0xce, 0xc5, 0x51,
	0xdf, 0xe4, 0x36, 0x2a, 0x72, 0x81, 0x4f, 0x22, 0x92, 0x1a, 0x39, 0x05, 0x88, 0x93, 0xa4, 0xd8,
	0xe4, 0xe2, 0x57, 0x9a, 0x9a, 0xbf, 0xf2, 0x0c, 0x33, 0x4d, 0xc9, 0xf1, 0x36, 0x48, 0x47, 0x49,
	0x75, 0x72, 0x01, 0xb4, 0x92, 0x12, 0x63, 0x81, 0xa9, 0x3a, 0x7c, 0x23, 0xaa, 0x21, 0x39, 0x04,
	0x7b, 0x59, 0xa9, 0xf3, 0x7a, 0x91, 0xbd, 0xac, 0xc2, 0xbf, 0x2c, 0x70, 0x64, 0xa6, 0xef, 0x60,
	0xe6, 0x10, 0x6c, 0xca, 0x8c, 0x11, 0x9b, 0xb2, 0x3d, 0x4c, 0xfc, 0x6d, 0x41, 0x53, 0x5f, 0xdc,
	0x3b, 0xd8, 0xf8, 0x0a, 0x5a, 0x1b, 0x8e, 0xa5, 0xd4, 0x4c, 0x28, 0x12, 0x4e, 0xd3, 0x3d, 0xfc,
	0xfc, 0x01, 0x9d, 0x67, 0x73, 0xb2, 0xbf, 0x27, 0x7b, 0xd7, 0x53, 0x0f, 0x5c, 0x2e, 0xe2, 0x72,
	0x3b, 0x0c, 0x0a, 0x48, 0x36, 0xa3, 0x6b, 0x2a, 0x94, 0x4f, 0x37, 0xd2, 0x20, 0xfc, 0xd7, 0x86,
	0xa3, 0x9d, 0x19, 0xdc, 0xdf, 0x81, 0xf7, 0xd6, 0xe5, 0xd4, 0xc6, 0x6c, 0x9a, 0x92, 0x2e, 0x58,
	0xd7, 0xca, 0x8d, 0x1b, 0x59, 0xd7, 0x12, 0xcd, 0x8c, 0x0b, 0x6b, 0x26, 0x91, 0xfe, 0x68, 0xdc,
	0xc8, 0x2a, 0x09, 0x01, 0x87, 0xc7, 0x99, 0x30, 0xf1, 0xa8, 0x67, 0xc9, 0xad, 0x62, 0xbe, 0x0a,
	0xda, 0x9a, 0x93, 0xcf, 0xcf, 0x66, 0xf1, 0x50, 0xb5, 0xd6, 0xb3, 0xd8, 0x03, 0x37, 0xbd, 0xcf,
	0x30, 0x0f, 0x8e, 0xf4, 0x29, 0x15, 0x90, 0xd5, 0x71, 0xce, 0x1f, 0xb1, 0x0c, 0x40, 0x5f, 0x92,
	0x46, 0xcf, 0x2f, 0xa9, 0xf3, 0xe2, 0x92, 0xc2, 0x3f, 0xa1, 0x39, 0x64, 0xf4, 0x1e, 0x2b, 0xd9,
	0xcb, 0x31, 0x29, 0xb1, 0x8e, 0xd3, 0xa0, 0x9d, 0x08, 0x9c, 0xdd, 0x08, 0xbe, 0x06, 0x2f, 0xc9,
	0x28, 0x6a, 0xd5, 0x55, 0x6a, 0x5b, 0x13, 0xd3, 0x94, 0x7c, 0x03, 0x1d, 0x23, 0x8a, 0x8a, 0xa1,
	0x39, 0x2c, 0x68, 0x4a, 0xee, 0xac, 0x4f, 0x3f, 0x40, 0xf3, 0x0a, 0xc5, 0xaa, 0x48, 0x49, 0x0b,
	0x1a, 0x3f, 0x4e, 0x16, 0xfe, 0x01, 0x69, 0x83, 0x33, 0xfb, 0x79, 0xbe, 0xf0, 0x2d, 0x49, 0xcd,
	0x6e, 0x16, 0xbe, 0x4d, 0x00, 0x9a, 0xe3, 0xc9, 0xe5, 0x64, 0x31, 0xf1, 0x1b, 0xc4, 0x03, 0x77,
	0x36, 0x5c, 0x8c, 0x7e, 0xf2, 0x9d, 0x4f, 0xff, 0x5b, 0xe0, 0x4e, 0xe4, 0x02, 0x20, 0x3e, 0x74,
	0x87, 0x8c, 0x9a, 0xc9, 0xc2, 0xd4, 0x3f, 0x20, 0x5f, 0xc2, 0x17, 0xa3, 0x22, 0x7f, 0xc0, 0x92,
	0xc7, 0x82, 0x16, 0xf9, 0x5c, 0x8e, 0x86, 0x6f, 0x93, 0x2e, 0xb4, 0x55, 0x21, 0xcb, 0x2a, 0xdf,
	0x21, 0xc7, 0x40, 0x86, 0x8c, 0xde, 0xae, 0xa8, 0x40, 0x39, 0x99, 0xf3, 0x2a, 0x4f, 0x30, 0xf5,
	0x5d, 0x72, 0x02, 0xc7, 0x7a, 0x99, 0x6c, 0xa5, 0x91, 0x8e, 0xcb, 0x6f, 0xbe, 0xa1, 0x8d, 0x31,
	0x43, 0xa9, 0xb5, 0x48, 0x00, 0x3d, 0xf9, 0xb9, 0xbd, 0xea, 0x6a, 0xbf, 0x52, 0xea, 0x1e, 0x4f,
	0x7a, 0x98, 0xb2, 0x57, 0x1d, 0xb0, 0xc3, 0xd7, 0xf5, 0x9d, 0x8b, 0xef, 0xc1, 0xfa, 0x85, 0x84,
	0x60, 0x8f, 0x0b, 0xd2, 0x55, 0x9b, 0xd5, 0xac, 0xff, 0x93, 0x17, 0x28, 0x3c, 0xf8, 0x1c, 0x00,
	0x00, 0xff, 0xff, 0x24, 0x40, 0xd3, 0x5c, 0x68, 0x06, 0x00, 0x00,
}
