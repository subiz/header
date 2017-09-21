// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/client.proto

/*
Package client is a generated protocol buffer package.

It is generated from these files:
	client/client.proto

It has these top-level messages:
	Clients
	Client
	Id
*/
package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import auth "bitbucket.org/subiz/servicespec/proto/auth"
import common "bitbucket.org/subiz/servicespec/proto/common"

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
	Event_ClientCreateRequested Event = 20
	Event_ClientUpdateRequested Event = 21
	Event_ClientDeleteRequested Event = 22
	Event_ClientReadRequested   Event = 23
	Event_ClientListRequested   Event = 24
)

var Event_name = map[int32]string{
	20: "ClientCreateRequested",
	21: "ClientUpdateRequested",
	22: "ClientDeleteRequested",
	23: "ClientReadRequested",
	24: "ClientListRequested",
}
var Event_value = map[string]int32{
	"ClientCreateRequested": 20,
	"ClientUpdateRequested": 21,
	"ClientDeleteRequested": 22,
	"ClientReadRequested":   23,
	"ClientListRequested":   24,
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
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Clients struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Clients          []*Client       `protobuf:"bytes,2,rep,name=clients" json:"clients,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Clients) Reset()                    { *m = Clients{} }
func (m *Clients) String() string            { return proto.CompactTextString(m) }
func (*Clients) ProtoMessage()               {}
func (*Clients) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Clients) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Clients) GetClients() []*Client {
	if m != nil {
		return m.Clients
	}
	return nil
}

type Client struct {
	Ctx *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	// 128 bit string identification of the client.
	Id *string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	// secre used to authorize client with oauth2 server
	Secret *string `protobuf:"bytes,4,opt,name=secret" json:"secret,omitempty"`
	// LogoUrl is url to logo of the client, should be 256x256 and lessthan 256KB
	LogoUrl   *string `protobuf:"bytes,5,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
	AccountId *string `protobuf:"bytes,6,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	// IsVerified will be true if client is verified by the authority
	IsVerified *bool `protobuf:"varint,8,opt,name=is_verified,json=isVerified" json:"is_verified,omitempty"`
	// List of URLs which client must register for oauth redirection
	RedirectUrl *string          `protobuf:"bytes,9,opt,name=redirect_url,json=redirectUrl" json:"redirect_url,omitempty"`
	Type        *auth.ClientType `protobuf:"varint,11,opt,name=type,enum=auth.ClientType" json:"type,omitempty"`
	Name        *string          `protobuf:"bytes,12,opt,name=name" json:"name,omitempty"`
	// Version number of the client.
	Version          *string `protobuf:"bytes,14,opt,name=version" json:"version,omitempty"`
	IsEnabled        *bool   `protobuf:"varint,15,opt,name=is_enabled,json=isEnabled" json:"is_enabled,omitempty"`
	CreatedTime      *int64  `protobuf:"varint,17,opt,name=created_time,json=createdTime" json:"created_time,omitempty"`
	LastModifiedTime *int64  `protobuf:"varint,18,opt,name=last_modified_time,json=lastModifiedTime" json:"last_modified_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Client) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Client) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Client) GetSecret() string {
	if m != nil && m.Secret != nil {
		return *m.Secret
	}
	return ""
}

func (m *Client) GetLogoUrl() string {
	if m != nil && m.LogoUrl != nil {
		return *m.LogoUrl
	}
	return ""
}

func (m *Client) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Client) GetIsVerified() bool {
	if m != nil && m.IsVerified != nil {
		return *m.IsVerified
	}
	return false
}

func (m *Client) GetRedirectUrl() string {
	if m != nil && m.RedirectUrl != nil {
		return *m.RedirectUrl
	}
	return ""
}

func (m *Client) GetType() auth.ClientType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return auth.ClientType_STONE
}

func (m *Client) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Client) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *Client) GetIsEnabled() bool {
	if m != nil && m.IsEnabled != nil {
		return *m.IsEnabled
	}
	return false
}

func (m *Client) GetCreatedTime() int64 {
	if m != nil && m.CreatedTime != nil {
		return *m.CreatedTime
	}
	return 0
}

func (m *Client) GetLastModifiedTime() int64 {
	if m != nil && m.LastModifiedTime != nil {
		return *m.LastModifiedTime
	}
	return 0
}

type Id struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Id               *string         `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	AccountId        *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Id) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Id) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Id) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func init() {
	proto.RegisterType((*Clients)(nil), "client.Clients")
	proto.RegisterType((*Client)(nil), "client.Client")
	proto.RegisterType((*Id)(nil), "client.Id")
	proto.RegisterEnum("client.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("client/client.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x95, 0xcd, 0x76, 0xff, 0x4c, 0xaa, 0xed, 0xe2, 0xd2, 0xd6, 0xad, 0x84, 0x48, 0x57,
	0x1c, 0x22, 0x84, 0x12, 0x69, 0x25, 0x0e, 0x9c, 0x97, 0x1e, 0x2a, 0xc1, 0x25, 0x6a, 0xf7, 0x1a,
	0x65, 0xed, 0xa1, 0x58, 0x24, 0x71, 0xb0, 0x9d, 0x55, 0xcb, 0x23, 0x70, 0xe7, 0x7d, 0x51, 0x6c,
	0x57, 0xa1, 0x9c, 0xaa, 0x5e, 0xd6, 0x3b, 0xdf, 0x6f, 0x3c, 0xfe, 0x66, 0x26, 0x70, 0xcc, 0x2a,
	0x81, 0x8d, 0xc9, 0xdc, 0x91, 0xb6, 0x4a, 0x1a, 0x49, 0x26, 0x2e, 0xba, 0xf8, 0xb8, 0x13, 0x66,
	0xd7, 0xb1, 0x1f, 0x68, 0x52, 0xa9, 0xee, 0x32, 0xdd, 0xed, 0xc4, 0xaf, 0x4c, 0xa3, 0xda, 0x0b,
	0x86, 0xba, 0x45, 0x96, 0xd9, 0xf4, 0xac, 0xec, 0xcc, 0x77, 0xfb, 0xe3, 0xae, 0x5f, 0x7c, 0x7a,
	0xde, 0x35, 0x26, 0xeb, 0x5a, 0x36, 0xfe, 0x70, 0x57, 0x57, 0x5b, 0x98, 0x6e, 0xec, 0xdb, 0x9a,
	0x5c, 0x42, 0xc8, 0xcc, 0x3d, 0x0d, 0xe2, 0x20, 0x89, 0xd6, 0x47, 0xa9, 0x4f, 0xdb, 0xc8, 0xc6,
	0xe0, 0xbd, 0xc9, 0x7b, 0x46, 0x12, 0x98, 0x3a, 0xa7, 0x9a, 0x8e, 0xe2, 0x30, 0x89, 0xd6, 0x8b,
	0xd4, 0xf7, 0xe1, 0x8a, 0xe4, 0x8f, 0x78, 0xf5, 0x27, 0x84, 0x89, 0xd3, 0x9e, 0x53, 0x77, 0x01,
	0x23, 0xc1, 0x69, 0x18, 0x07, 0xc9, 0x3c, 0x1f, 0x09, 0x4e, 0x4e, 0x61, 0xa2, 0x91, 0x29, 0x34,
	0x74, 0x6c, 0x35, 0x1f, 0x91, 0x73, 0x98, 0x55, 0xf2, 0x4e, 0x16, 0x9d, 0xaa, 0xe8, 0x81, 0x25,
	0xd3, 0x3e, 0xbe, 0x55, 0x15, 0x79, 0x03, 0x50, 0x32, 0x26, 0xbb, 0xc6, 0x14, 0x82, 0xd3, 0x89,
	0x85, 0x73, 0xaf, 0x5c, 0x73, 0xf2, 0x16, 0x22, 0xa1, 0x8b, 0x3d, 0x2a, 0xf1, 0x4d, 0x20, 0xa7,
	0xb3, 0x38, 0x48, 0x66, 0x39, 0x08, 0xbd, 0xf5, 0x0a, 0xb9, 0x84, 0x43, 0x85, 0x5c, 0x28, 0x64,
	0xc6, 0x96, 0x9f, 0xdb, 0x0a, 0xd1, 0xa3, 0xd6, 0x3f, 0xf1, 0x0e, 0xc6, 0xe6, 0xa1, 0x45, 0x1a,
	0xc5, 0x41, 0xb2, 0x58, 0x2f, 0x53, 0xbb, 0x01, 0xd7, 0xe4, 0xcd, 0x43, 0x8b, 0xb9, 0xa5, 0x84,
	0xc0, 0xb8, 0x29, 0x6b, 0xa4, 0x87, 0xb6, 0x80, 0xfd, 0x4f, 0x28, 0x4c, 0xf7, 0xa8, 0xb4, 0x90,
	0x0d, 0x5d, 0x38, 0xdb, 0x3e, 0xec, 0x6d, 0x0b, 0x5d, 0x60, 0x53, 0xee, 0x2a, 0xe4, 0xf4, 0xc8,
	0xda, 0x9a, 0x0b, 0x7d, 0xe5, 0x84, 0xde, 0x15, 0x53, 0x58, 0x1a, 0xe4, 0x85, 0x11, 0x35, 0xd2,
	0x57, 0x71, 0x90, 0x84, 0x79, 0xe4, 0xb5, 0x1b, 0x51, 0x23, 0xf9, 0x00, 0xa4, 0x2a, 0xb5, 0x29,
	0x6a, 0xc9, 0x6d, 0x27, 0x2e, 0x91, 0xd8, 0xc4, 0x65, 0x4f, 0xbe, 0x7a, 0xd0, 0x67, 0xaf, 0xb6,
	0x30, 0xba, 0xe6, 0x2f, 0x59, 0xc9, 0xd3, 0xf9, 0x8e, 0xfe, 0x9b, 0xef, 0xfb, 0xdf, 0x01, 0x1c,
	0x5c, 0xed, 0xfb, 0x75, 0x9f, 0xc3, 0x89, 0x9b, 0xc9, 0xc6, 0x9a, 0xcc, 0xf1, 0x67, 0x87, 0xda,
	0x20, 0x5f, 0xbe, 0x1e, 0xd0, 0x6d, 0xcb, 0x9f, 0xa0, 0x93, 0x01, 0x7d, 0xc6, 0x0a, 0xff, 0x45,
	0xa7, 0xe4, 0x0c, 0x8e, 0xfd, 0xd7, 0x85, 0x25, 0x1f, 0xc0, 0xd9, 0x00, 0xbe, 0x08, 0x6d, 0x06,
	0x40, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x51, 0x84, 0xf3, 0x57, 0x64, 0x03, 0x00, 0x00,
}
