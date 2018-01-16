// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/client.proto

/*
Package client is a generated protocol buffer package.

It is generated from these files:
	client/client.proto

It has these top-level messages:
	AllType
	Clients
	Client
	Authorization
*/
package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bitbucket.org/subiz/header/common"
import auth "bitbucket.org/subiz/header/auth"

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
	Event_ClientUpserted        Event = 25
	Event_ClientDeleted         Event = 26
	// ConnectorAuthorized = 13; // conversation send to auth to send ClientAuthorized
	Event_ClientAuthorized    Event = 12
	Event_ClientRequested     Event = 30
	Event_ClientSynced        Event = 31
	Event_ConnectorAuthorized Event = 13
)

var Event_name = map[int32]string{
	20: "ClientCreateRequested",
	21: "ClientUpdateRequested",
	22: "ClientDeleteRequested",
	23: "ClientReadRequested",
	24: "ClientListRequested",
	25: "ClientUpserted",
	26: "ClientDeleted",
	12: "ClientAuthorized",
	30: "ClientRequested",
	31: "ClientSynced",
	13: "ConnectorAuthorized",
}
var Event_value = map[string]int32{
	"ClientCreateRequested": 20,
	"ClientUpdateRequested": 21,
	"ClientDeleteRequested": 22,
	"ClientReadRequested":   23,
	"ClientListRequested":   24,
	"ClientUpserted":        25,
	"ClientDeleted":         26,
	"ClientAuthorized":      12,
	"ClientRequested":       30,
	"ClientSynced":          31,
	"ConnectorAuthorized":   13,
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

type Client_Type int32

const (
	Client_app       Client_Type = 0
	Client_connector Client_Type = 1
	Client_bot       Client_Type = 3
)

var Client_Type_name = map[int32]string{
	0: "app",
	1: "connector",
	3: "bot",
}
var Client_Type_value = map[string]int32{
	"app":       0,
	"connector": 1,
	"bot":       3,
}

func (x Client_Type) Enum() *Client_Type {
	p := new(Client_Type)
	*p = x
	return p
}
func (x Client_Type) String() string {
	return proto.EnumName(Client_Type_name, int32(x))
}
func (x *Client_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Client_Type_value, data, "Client_Type")
	if err != nil {
		return err
	}
	*x = Client_Type(value)
	return nil
}
func (Client_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Client_ChannelType int32

const (
	Client_subiz    Client_ChannelType = 0
	Client_email    Client_ChannelType = 1
	Client_facebook Client_ChannelType = 2
	Client_viber    Client_ChannelType = 3
)

var Client_ChannelType_name = map[int32]string{
	0: "subiz",
	1: "email",
	2: "facebook",
	3: "viber",
}
var Client_ChannelType_value = map[string]int32{
	"subiz":    0,
	"email":    1,
	"facebook": 2,
	"viber":    3,
}

func (x Client_ChannelType) Enum() *Client_ChannelType {
	p := new(Client_ChannelType)
	*p = x
	return p
}
func (x Client_ChannelType) String() string {
	return proto.EnumName(Client_ChannelType_name, int32(x))
}
func (x *Client_ChannelType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Client_ChannelType_value, data, "Client_ChannelType")
	if err != nil {
		return err
	}
	*x = Client_ChannelType(value)
	return nil
}
func (Client_ChannelType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

type AllType struct {
	Client           *Clients `protobuf:"bytes,2,opt,name=client" json:"client,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AllType) Reset()                    { *m = AllType{} }
func (m *AllType) String() string            { return proto.CompactTextString(m) }
func (*AllType) ProtoMessage()               {}
func (*AllType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AllType) GetClient() *Clients {
	if m != nil {
		return m.Client
	}
	return nil
}

type Clients struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Clients          []*Client       `protobuf:"bytes,2,rep,name=clients" json:"clients,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Clients) Reset()                    { *m = Clients{} }
func (m *Clients) String() string            { return proto.CompactTextString(m) }
func (*Clients) ProtoMessage()               {}
func (*Clients) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
	Id  *string         `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	// secre used to authorize client with oauth2 server
	Secret *string `protobuf:"bytes,4,opt,name=secret" json:"secret,omitempty"`
	// LogoUrl is url to logo of the client, should be 256x256 and lessthan 256KB
	LogoUrl   *string `protobuf:"bytes,5,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
	AccountId *string `protobuf:"bytes,6,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	// IsVerified will be true if client is verified by the authority
	IsVerified *bool  `protobuf:"varint,8,opt,name=is_verified,json=isVerified" json:"is_verified,omitempty"`
	Verified   *int64 `protobuf:"varint,9,opt,name=verified" json:"verified,omitempty"`
	// List of URLs which client must register for oauth redirection
	RedirectUrl *string `protobuf:"bytes,10,opt,name=redirect_url,json=redirectUrl" json:"redirect_url,omitempty"`
	Type        *string `protobuf:"bytes,11,opt,name=type" json:"type,omitempty"`
	Name        *string `protobuf:"bytes,12,opt,name=name" json:"name,omitempty"`
	// Version number of the client.
	Version    *string  `protobuf:"bytes,14,opt,name=version" json:"version,omitempty"`
	IsEnabled  *bool    `protobuf:"varint,15,opt,name=is_enabled,json=isEnabled" json:"is_enabled,omitempty"`
	Created    *int64   `protobuf:"varint,17,opt,name=created" json:"created,omitempty"`
	Modified   *int64   `protobuf:"varint,18,opt,name=modified" json:"modified,omitempty"`
	WebhookUrl *string  `protobuf:"bytes,20,opt,name=webhook_url,json=webhookUrl" json:"webhook_url,omitempty"`
	Events     []string `protobuf:"bytes,19,rep,name=events" json:"events,omitempty"`
	// for connector only
	ChannelType      *string `protobuf:"bytes,21,opt,name=channel_type,json=channelType" json:"channel_type,omitempty"`
	AvailabilityUrl  *string `protobuf:"bytes,22,opt,name=availability_url,json=availabilityUrl" json:"availability_url,omitempty"`
	PingUrl          *string `protobuf:"bytes,23,opt,name=ping_url,json=pingUrl" json:"ping_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

func (m *Client) GetVerified() int64 {
	if m != nil && m.Verified != nil {
		return *m.Verified
	}
	return 0
}

func (m *Client) GetRedirectUrl() string {
	if m != nil && m.RedirectUrl != nil {
		return *m.RedirectUrl
	}
	return ""
}

func (m *Client) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
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

func (m *Client) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Client) GetModified() int64 {
	if m != nil && m.Modified != nil {
		return *m.Modified
	}
	return 0
}

func (m *Client) GetWebhookUrl() string {
	if m != nil && m.WebhookUrl != nil {
		return *m.WebhookUrl
	}
	return ""
}

func (m *Client) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Client) GetChannelType() string {
	if m != nil && m.ChannelType != nil {
		return *m.ChannelType
	}
	return ""
}

func (m *Client) GetAvailabilityUrl() string {
	if m != nil && m.AvailabilityUrl != nil {
		return *m.AvailabilityUrl
	}
	return ""
}

func (m *Client) GetPingUrl() string {
	if m != nil && m.PingUrl != nil {
		return *m.PingUrl
	}
	return ""
}

type Authorization struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Issuer           *string         `protobuf:"bytes,3,opt,name=issuer" json:"issuer,omitempty"`
	Type             *auth.Type      `protobuf:"varint,4,opt,name=type,enum=auth.Type" json:"type,omitempty"`
	Method           *auth.Method    `protobuf:"bytes,5,opt,name=method" json:"method,omitempty"`
	ClientId         *string         `protobuf:"bytes,7,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	ClientType       *auth.Type      `protobuf:"varint,8,opt,name=client_type,json=clientType,enum=auth.Type" json:"client_type,omitempty"`
	ClientAccountId  *string         `protobuf:"bytes,10,opt,name=client_account_id,json=clientAccountId" json:"client_account_id,omitempty"`
	Scopes           []string        `protobuf:"bytes,9,rep,name=scopes" json:"scopes,omitempty"`
	IntegrationId    *string         `protobuf:"bytes,11,opt,name=integration_id,json=integrationId" json:"integration_id,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Authorization) Reset()                    { *m = Authorization{} }
func (m *Authorization) String() string            { return proto.CompactTextString(m) }
func (*Authorization) ProtoMessage()               {}
func (*Authorization) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Authorization) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Authorization) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Authorization) GetIssuer() string {
	if m != nil && m.Issuer != nil {
		return *m.Issuer
	}
	return ""
}

func (m *Authorization) GetType() auth.Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return auth.Type_unknown
}

func (m *Authorization) GetMethod() *auth.Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *Authorization) GetClientId() string {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return ""
}

func (m *Authorization) GetClientType() auth.Type {
	if m != nil && m.ClientType != nil {
		return *m.ClientType
	}
	return auth.Type_unknown
}

func (m *Authorization) GetClientAccountId() string {
	if m != nil && m.ClientAccountId != nil {
		return *m.ClientAccountId
	}
	return ""
}

func (m *Authorization) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *Authorization) GetIntegrationId() string {
	if m != nil && m.IntegrationId != nil {
		return *m.IntegrationId
	}
	return ""
}

func init() {
	proto.RegisterType((*AllType)(nil), "client.AllType")
	proto.RegisterType((*Clients)(nil), "client.Clients")
	proto.RegisterType((*Client)(nil), "client.Client")
	proto.RegisterType((*Authorization)(nil), "client.Authorization")
	proto.RegisterEnum("client.Event", Event_name, Event_value)
	proto.RegisterEnum("client.Client_Type", Client_Type_name, Client_Type_value)
	proto.RegisterEnum("client.Client_ChannelType", Client_ChannelType_name, Client_ChannelType_value)
}

func init() { proto.RegisterFile("client/client.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xe3, 0x36,
	0x10, 0x5d, 0xdb, 0x89, 0x2d, 0x8d, 0xfc, 0xa1, 0x30, 0x89, 0xc3, 0xa4, 0xe8, 0xae, 0xd7, 0x68,
	0xb1, 0x6e, 0x0a, 0x38, 0x80, 0xcf, 0xbd, 0x04, 0xe9, 0x1e, 0x02, 0xb4, 0x17, 0xb5, 0xd9, 0xab,
	0x21, 0x89, 0xb3, 0x31, 0x11, 0x59, 0x54, 0x49, 0xca, 0x5d, 0xef, 0x1f, 0xe8, 0xa9, 0xff, 0xb2,
	0x3f, 0xa4, 0xe0, 0x87, 0xd6, 0x4a, 0x0b, 0x2c, 0x72, 0x31, 0x39, 0xef, 0x0d, 0xe7, 0xcd, 0xd0,
	0x8f, 0x82, 0xd3, 0xbc, 0xe0, 0x58, 0xea, 0x1b, 0xb7, 0x2c, 0x2b, 0x29, 0xb4, 0x20, 0x7d, 0x17,
	0x5d, 0x2d, 0x33, 0xae, 0xb3, 0x3a, 0x7f, 0x42, 0xbd, 0x14, 0xf2, 0xf1, 0x46, 0xd5, 0x19, 0xff,
	0x7c, 0xb3, 0xc1, 0x94, 0xa1, 0xbc, 0xc9, 0xc5, 0x76, 0x2b, 0x4a, 0xbf, 0xb8, 0x73, 0x57, 0xd7,
	0x5f, 0xc9, 0x4f, 0x6b, 0xbd, 0xb1, 0x3f, 0x2e, 0x77, 0xbe, 0x82, 0xc1, 0x6d, 0x51, 0xfc, 0xbe,
	0xaf, 0x90, 0xbc, 0x03, 0x2f, 0x48, 0xbb, 0xb3, 0xce, 0x22, 0x5a, 0x4d, 0x96, 0xbe, 0x9b, 0x3b,
	0xbb, 0xa8, 0xc4, 0xd3, 0xf3, 0x0f, 0x30, 0xf0, 0x10, 0x79, 0x0b, 0xbd, 0x5c, 0x7f, 0xa2, 0x9d,
	0xe6, 0x80, 0x6b, 0xe3, 0x4e, 0x94, 0x1a, 0x3f, 0xe9, 0xc4, 0x70, 0x64, 0x01, 0x03, 0x77, 0x4e,
	0xd1, 0xee, 0xac, 0xb7, 0x88, 0x56, 0xe3, 0xe7, 0x75, 0x93, 0x86, 0x9e, 0xff, 0x7d, 0x0c, 0x7d,
	0x87, 0xbd, 0xa4, 0xee, 0x18, 0xba, 0x9c, 0xd1, 0xde, 0xac, 0xb3, 0x08, 0x93, 0x2e, 0x67, 0x64,
	0x0a, 0x7d, 0x85, 0xb9, 0x44, 0x4d, 0x8f, 0x2c, 0xe6, 0x23, 0x72, 0x09, 0x41, 0x21, 0x1e, 0xc5,
	0xba, 0x96, 0x05, 0x3d, 0xb6, 0xcc, 0xc0, 0xc4, 0x0f, 0xb2, 0x20, 0xdf, 0x02, 0xa4, 0x79, 0x2e,
	0xea, 0x52, 0xaf, 0x39, 0xa3, 0x7d, 0x4b, 0x86, 0x1e, 0xb9, 0x67, 0xe4, 0x0d, 0x44, 0x5c, 0xad,
	0x77, 0x28, 0xf9, 0x47, 0x8e, 0x8c, 0x06, 0xb3, 0xce, 0x22, 0x48, 0x80, 0xab, 0x0f, 0x1e, 0x21,
	0x57, 0x10, 0x7c, 0x61, 0xc3, 0x59, 0x67, 0xd1, 0x4b, 0xbe, 0xc4, 0xe4, 0x2d, 0x0c, 0x25, 0x32,
	0x2e, 0x31, 0xd7, 0x56, 0x1a, 0x6c, 0xf5, 0xa8, 0xc1, 0x8c, 0x3c, 0x81, 0x23, 0xbd, 0xaf, 0x90,
	0x46, 0x96, 0xb2, 0x7b, 0x83, 0x95, 0xe9, 0x16, 0xe9, 0xd0, 0x61, 0x66, 0x4f, 0x28, 0x0c, 0x76,
	0x28, 0x15, 0x17, 0x25, 0x1d, 0xbb, 0x01, 0x7c, 0x68, 0x06, 0xe0, 0x6a, 0x8d, 0x65, 0x9a, 0x15,
	0xc8, 0xe8, 0xc4, 0x36, 0x18, 0x72, 0xf5, 0xde, 0x01, 0xe6, 0x60, 0x2e, 0x31, 0xd5, 0xc8, 0xe8,
	0x89, 0x6d, 0xaf, 0x09, 0x4d, 0xe7, 0x5b, 0xc1, 0x5c, 0xe7, 0xc4, 0x75, 0xde, 0xc4, 0x66, 0xec,
	0x3f, 0x31, 0xdb, 0x08, 0xf1, 0x64, 0x1b, 0x3f, 0xb3, 0x92, 0xe0, 0x21, 0xd3, 0xf7, 0x14, 0xfa,
	0xb8, 0xb3, 0x7f, 0xe8, 0xe9, 0xac, 0x67, 0x6e, 0xda, 0x45, 0x66, 0xe4, 0x7c, 0x93, 0x96, 0x25,
	0x16, 0x6b, 0x3b, 0xd7, 0xb9, 0x1b, 0xd9, 0x63, 0xd6, 0x63, 0x3f, 0x40, 0x9c, 0xee, 0x52, 0x5e,
	0xa4, 0x19, 0x2f, 0xb8, 0xde, 0x5b, 0x81, 0xa9, 0x4d, 0x9b, 0xb4, 0x71, 0xa3, 0x72, 0x09, 0x41,
	0xc5, 0xcb, 0x47, 0x9b, 0x72, 0xe1, 0xc6, 0x36, 0xf1, 0x83, 0x2c, 0xe6, 0xef, 0xe0, 0xc8, 0x56,
	0x1b, 0x40, 0x2f, 0xad, 0xaa, 0xf8, 0x15, 0x19, 0x41, 0x98, 0x8b, 0xb2, 0xc4, 0x5c, 0x0b, 0x19,
	0x77, 0x0c, 0x9e, 0x09, 0x1d, 0xf7, 0xe6, 0x3f, 0x41, 0x74, 0xd7, 0x52, 0x0f, 0xe1, 0xd8, 0x3e,
	0x86, 0xf8, 0x95, 0xd9, 0xe2, 0x36, 0xe5, 0x45, 0xdc, 0x21, 0x43, 0x08, 0x3e, 0xa6, 0x39, 0x66,
	0x42, 0x3c, 0xc5, 0x5d, 0x43, 0xec, 0x78, 0x86, 0x32, 0xee, 0xcd, 0xff, 0xe9, 0xc2, 0xe8, 0xb6,
	0xd6, 0x1b, 0x21, 0xf9, 0xe7, 0x54, 0x9b, 0xfb, 0x7e, 0x81, 0x2d, 0x9f, 0x7b, 0xaa, 0xfb, 0x5f,
	0x4f, 0x4d, 0xa1, 0xcf, 0x95, 0xaa, 0x51, 0x7a, 0xe7, 0xfa, 0x88, 0xbc, 0xf6, 0x5e, 0x30, 0xde,
	0x1d, 0xaf, 0x60, 0x69, 0x9f, 0xa8, 0x69, 0xda, 0xfb, 0xe2, 0x3b, 0xe8, 0x6f, 0x51, 0x6f, 0x04,
	0xb3, 0x1e, 0x8e, 0x56, 0x43, 0x97, 0xf1, 0xab, 0xc5, 0x12, 0xcf, 0x91, 0x6f, 0x20, 0x74, 0x8f,
	0xc9, 0x68, 0x0f, 0xac, 0x40, 0xe0, 0x80, 0x7b, 0x46, 0x7e, 0x84, 0xc8, 0x93, 0x56, 0x29, 0xf8,
	0x9f, 0x12, 0x38, 0xda, 0x5e, 0xd5, 0x35, 0x9c, 0xf8, 0xe4, 0xd6, 0x34, 0xce, 0xc3, 0x13, 0x47,
	0xdc, 0xb6, 0x67, 0x52, 0xb9, 0xa8, 0x50, 0xd1, 0xd0, 0xf9, 0xc1, 0x45, 0xe4, 0x7b, 0x18, 0xf3,
	0x52, 0xe3, 0xa3, 0xb4, 0x97, 0x67, 0x0a, 0x38, 0xa7, 0x8f, 0x5a, 0xe8, 0x3d, 0xbb, 0xfe, 0xab,
	0x0b, 0xc7, 0xef, 0x8d, 0x83, 0xc8, 0x25, 0x9c, 0xbb, 0xf7, 0x7f, 0x67, 0x6d, 0x9a, 0xe0, 0x1f,
	0x35, 0x2a, 0x8d, 0x2c, 0x3e, 0x3b, 0x50, 0x0f, 0x15, 0x7b, 0x46, 0x9d, 0x1f, 0xa8, 0x9f, 0xb1,
	0xc0, 0x36, 0x35, 0x25, 0x17, 0x70, 0xea, 0x3f, 0x32, 0x98, 0xb2, 0x03, 0x71, 0x71, 0x20, 0x7e,
	0xe1, 0x4a, 0x1f, 0x08, 0x4a, 0x08, 0x8c, 0x1b, 0x1d, 0x85, 0xd2, 0x60, 0x97, 0xe4, 0x04, 0x46,
	0x6d, 0x01, 0x16, 0x5f, 0x91, 0x33, 0x88, 0x1d, 0xd4, 0xf8, 0x03, 0x59, 0x3c, 0x24, 0xa7, 0x30,
	0x69, 0xe4, 0x9a, 0x8a, 0xaf, 0x49, 0x0c, 0x43, 0x07, 0xfe, 0xb6, 0x2f, 0x73, 0x64, 0xf1, 0x1b,
	0x2b, 0xde, 0xb8, 0xb5, 0x75, 0x7e, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0xe7, 0x92,
	0x24, 0x06, 0x06, 0x00, 0x00,
}
