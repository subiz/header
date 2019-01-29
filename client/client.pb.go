// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/client.proto

package client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/subiz/header/common"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
	Event_ClientRequested    Event = 30
	Event_ClientSynced       Event = 31
	Event_ClientAuthorized   Event = 33
	Event_ClientDeauthorized Event = 34
)

var Event_name = map[int32]string{
	20: "ClientCreateRequested",
	21: "ClientUpdateRequested",
	22: "ClientDeleteRequested",
	23: "ClientReadRequested",
	24: "ClientListRequested",
	25: "ClientUpserted",
	26: "ClientDeleted",
	30: "ClientRequested",
	31: "ClientSynced",
	33: "ClientAuthorized",
	34: "ClientDeauthorized",
}

var Event_value = map[string]int32{
	"ClientCreateRequested": 20,
	"ClientUpdateRequested": 21,
	"ClientDeleteRequested": 22,
	"ClientReadRequested":   23,
	"ClientListRequested":   24,
	"ClientUpserted":        25,
	"ClientDeleted":         26,
	"ClientRequested":       30,
	"ClientSynced":          31,
	"ClientAuthorized":      33,
	"ClientDeauthorized":    34,
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

func (Event) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{0}
}

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

func (Client_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{1, 0}
}

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

func (Client_ChannelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{1, 1}
}

type Clients struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Clients              []*Client       `protobuf:"bytes,2,rep,name=clients" json:"clients,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Clients) Reset()         { *m = Clients{} }
func (m *Clients) String() string { return proto.CompactTextString(m) }
func (*Clients) ProtoMessage()    {}
func (*Clients) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{0}
}

func (m *Clients) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clients.Unmarshal(m, b)
}
func (m *Clients) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clients.Marshal(b, m, deterministic)
}
func (m *Clients) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clients.Merge(m, src)
}
func (m *Clients) XXX_Size() int {
	return xxx_messageInfo_Clients.Size(m)
}
func (m *Clients) XXX_DiscardUnknown() {
	xxx_messageInfo_Clients.DiscardUnknown(m)
}

var xxx_messageInfo_Clients proto.InternalMessageInfo

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
	RedirectUri *string `protobuf:"bytes,10,opt,name=redirect_uri,json=redirectUri" json:"redirect_uri,omitempty"`
	Type        *string `protobuf:"bytes,11,opt,name=type" json:"type,omitempty"`
	Name        *string `protobuf:"bytes,12,opt,name=name" json:"name,omitempty"`
	// Version number of the client.
	Version    *string  `protobuf:"bytes,14,opt,name=version" json:"version,omitempty"`
	IsEnabled  *bool    `protobuf:"varint,15,opt,name=is_enabled,json=isEnabled" json:"is_enabled,omitempty"`
	Created    *int64   `protobuf:"varint,17,opt,name=created" json:"created,omitempty"`
	Modified   *int64   `protobuf:"varint,18,opt,name=modified" json:"modified,omitempty"`
	WebhookUri *string  `protobuf:"bytes,20,opt,name=webhook_uri,json=webhookUri" json:"webhook_uri,omitempty"`
	Events     []string `protobuf:"bytes,19,rep,name=events" json:"events,omitempty"`
	// for connector only
	ChannelType          *string  `protobuf:"bytes,21,opt,name=channel_type,json=channelType" json:"channel_type,omitempty"`
	AvailabilityUri      *string  `protobuf:"bytes,22,opt,name=availability_uri,json=availabilityUri" json:"availability_uri,omitempty"`
	PingUri              *string  `protobuf:"bytes,23,opt,name=ping_uri,json=pingUri" json:"ping_uri,omitempty"`
	IsInternal           *bool    `protobuf:"varint,24,opt,name=is_internal,json=isInternal" json:"is_internal,omitempty"`
	UnsubscribeUri       *string  `protobuf:"bytes,25,opt,name=unsubscribe_uri,json=unsubscribeUri" json:"unsubscribe_uri,omitempty"`
	Scopes               []string `protobuf:"bytes,26,rep,name=scopes" json:"scopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{1}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

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

func (m *Client) GetRedirectUri() string {
	if m != nil && m.RedirectUri != nil {
		return *m.RedirectUri
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

func (m *Client) GetWebhookUri() string {
	if m != nil && m.WebhookUri != nil {
		return *m.WebhookUri
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

func (m *Client) GetAvailabilityUri() string {
	if m != nil && m.AvailabilityUri != nil {
		return *m.AvailabilityUri
	}
	return ""
}

func (m *Client) GetPingUri() string {
	if m != nil && m.PingUri != nil {
		return *m.PingUri
	}
	return ""
}

func (m *Client) GetIsInternal() bool {
	if m != nil && m.IsInternal != nil {
		return *m.IsInternal
	}
	return false
}

func (m *Client) GetUnsubscribeUri() string {
	if m != nil && m.UnsubscribeUri != nil {
		return *m.UnsubscribeUri
	}
	return ""
}

func (m *Client) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

type AuthorizedClient struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Client               *Client         `protobuf:"bytes,2,opt,name=client" json:"client,omitempty"`
	IssueAccountId       *string         `protobuf:"bytes,3,opt,name=issue_account_id,json=issueAccountId" json:"issue_account_id,omitempty"`
	Issuer               *string         `protobuf:"bytes,4,opt,name=issuer" json:"issuer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AuthorizedClient) Reset()         { *m = AuthorizedClient{} }
func (m *AuthorizedClient) String() string { return proto.CompactTextString(m) }
func (*AuthorizedClient) ProtoMessage()    {}
func (*AuthorizedClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{2}
}

func (m *AuthorizedClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizedClient.Unmarshal(m, b)
}
func (m *AuthorizedClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizedClient.Marshal(b, m, deterministic)
}
func (m *AuthorizedClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizedClient.Merge(m, src)
}
func (m *AuthorizedClient) XXX_Size() int {
	return xxx_messageInfo_AuthorizedClient.Size(m)
}
func (m *AuthorizedClient) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizedClient.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizedClient proto.InternalMessageInfo

func (m *AuthorizedClient) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *AuthorizedClient) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *AuthorizedClient) GetIssueAccountId() string {
	if m != nil && m.IssueAccountId != nil {
		return *m.IssueAccountId
	}
	return ""
}

func (m *AuthorizedClient) GetIssuer() string {
	if m != nil && m.Issuer != nil {
		return *m.Issuer
	}
	return ""
}

func init() {
	proto.RegisterEnum("client.Event", Event_name, Event_value)
	proto.RegisterEnum("client.Client_Type", Client_Type_name, Client_Type_value)
	proto.RegisterEnum("client.Client_ChannelType", Client_ChannelType_name, Client_ChannelType_value)
	proto.RegisterType((*Clients)(nil), "client.Clients")
	proto.RegisterType((*Client)(nil), "client.Client")
	proto.RegisterType((*AuthorizedClient)(nil), "client.AuthorizedClient")
}

func init() { proto.RegisterFile("client/client.proto", fileDescriptor_4d3551c163a1d198) }

var fileDescriptor_4d3551c163a1d198 = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0x1b, 0x37,
	0x10, 0xce, 0x6a, 0x6d, 0xfd, 0x8c, 0x1c, 0x89, 0xa1, 0x6d, 0x99, 0x12, 0xd0, 0x46, 0xd6, 0x21,
	0xd9, 0xb6, 0x80, 0x0d, 0xf8, 0xdc, 0x1e, 0x02, 0x37, 0x07, 0x03, 0xed, 0x45, 0xad, 0x73, 0x15,
	0x76, 0xc9, 0x89, 0x44, 0x64, 0x45, 0xaa, 0x24, 0x57, 0x8d, 0x73, 0x6f, 0x5f, 0xa3, 0x8f, 0xd8,
	0x57, 0x28, 0x48, 0xee, 0x5a, 0xaa, 0x72, 0xb0, 0x4f, 0xda, 0xef, 0xfb, 0xe6, 0x9f, 0x33, 0x82,
	0x53, 0x5e, 0x4a, 0x54, 0xee, 0x3a, 0xfe, 0x5c, 0x6d, 0x8c, 0x76, 0x9a, 0xb6, 0x23, 0x9a, 0xfc,
	0xb0, 0x94, 0x6e, 0x55, 0x15, 0x57, 0x5c, 0xaf, 0xaf, 0x6d, 0x55, 0xc8, 0x2f, 0xd7, 0x2b, 0xcc,
	0x05, 0x9a, 0x6b, 0xae, 0xd7, 0x6b, 0xad, 0xea, 0x9f, 0xe8, 0x34, 0xfb, 0x00, 0x9d, 0xdb, 0xe0,
	0x66, 0xe9, 0x25, 0xa4, 0xdc, 0x7d, 0x66, 0xc9, 0x34, 0xc9, 0xfa, 0x37, 0xc3, 0xab, 0xda, 0xec,
	0x56, 0x2b, 0x87, 0x9f, 0xdd, 0xdc, 0x6b, 0x34, 0x83, 0x4e, 0x4c, 0x62, 0x59, 0x6b, 0x9a, 0x66,
	0xfd, 0x9b, 0xc1, 0x55, 0x5d, 0x42, 0x0c, 0x32, 0x6f, 0xe4, 0xd9, 0xbf, 0xc7, 0xd0, 0x8e, 0xdc,
	0x73, 0xe2, 0x0e, 0xa0, 0x25, 0x05, 0x4b, 0xa7, 0x49, 0xd6, 0x9b, 0xb7, 0xa4, 0xa0, 0x23, 0x68,
	0x5b, 0xe4, 0x06, 0x1d, 0x3b, 0x0a, 0x5c, 0x8d, 0xe8, 0x18, 0xba, 0xa5, 0x5e, 0xea, 0x45, 0x65,
	0x4a, 0x76, 0x1c, 0x94, 0x8e, 0xc7, 0xf7, 0xa6, 0xa4, 0xdf, 0x00, 0xe4, 0x9c, 0xeb, 0x4a, 0xb9,
	0x85, 0x14, 0xac, 0x1d, 0xc4, 0x5e, 0xcd, 0xdc, 0x09, 0xfa, 0x1a, 0xfa, 0xd2, 0x2e, 0xb6, 0x68,
	0xe4, 0x47, 0x89, 0x82, 0x75, 0xa7, 0x49, 0xd6, 0x9d, 0x83, 0xb4, 0x1f, 0x6a, 0x86, 0x4e, 0xa0,
	0xfb, 0xa8, 0xf6, 0xa6, 0x49, 0x96, 0xce, 0x1f, 0x31, 0xbd, 0x84, 0x13, 0x83, 0x42, 0x1a, 0xe4,
	0x6e, 0x51, 0x19, 0xc9, 0x20, 0x44, 0xef, 0x37, 0xdc, 0xbd, 0x91, 0x94, 0xc2, 0x91, 0x7b, 0xd8,
	0x20, 0xeb, 0x07, 0x29, 0x7c, 0x7b, 0x4e, 0xe5, 0x6b, 0x64, 0x27, 0x91, 0xf3, 0xdf, 0x94, 0x41,
	0x67, 0x8b, 0xc6, 0x4a, 0xad, 0xd8, 0x20, 0x36, 0x50, 0x43, 0xdf, 0x80, 0xb4, 0x0b, 0x54, 0x79,
	0x51, 0xa2, 0x60, 0xc3, 0x50, 0x60, 0x4f, 0xda, 0xf7, 0x91, 0xf0, 0x8e, 0xdc, 0x60, 0xee, 0x50,
	0xb0, 0x57, 0xa1, 0xbc, 0x06, 0xfa, 0xca, 0xd7, 0x5a, 0xc4, 0xca, 0x69, 0xac, 0xbc, 0xc1, 0xbe,
	0xed, 0x3f, 0xb1, 0x58, 0x69, 0xfd, 0x29, 0x14, 0x7e, 0x16, 0x52, 0x42, 0x4d, 0xf9, 0xba, 0x47,
	0xd0, 0xc6, 0x6d, 0x78, 0xd0, 0xd3, 0x69, 0xea, 0x27, 0x1d, 0x91, 0x6f, 0x99, 0xaf, 0x72, 0xa5,
	0xb0, 0x5c, 0x84, 0xbe, 0xce, 0x63, 0xcb, 0x35, 0xf7, 0xbb, 0x6f, 0xef, 0x3b, 0x20, 0xf9, 0x36,
	0x97, 0x65, 0x5e, 0xc8, 0x52, 0xba, 0x87, 0x90, 0x60, 0x14, 0xcc, 0x86, 0xfb, 0xbc, 0xcf, 0x32,
	0x86, 0xee, 0x46, 0xaa, 0x65, 0x30, 0xb9, 0x88, 0x6d, 0x7b, 0xec, 0xa5, 0xf8, 0x30, 0x52, 0x39,
	0x34, 0x2a, 0x2f, 0x19, 0x6b, 0x1e, 0xe6, 0xae, 0x66, 0xe8, 0x5b, 0x18, 0x56, 0xca, 0x56, 0x85,
	0xe5, 0x46, 0x16, 0x18, 0x42, 0x8c, 0x43, 0x88, 0xc1, 0x1e, 0x5d, 0xb7, 0x62, 0xb9, 0xde, 0xa0,
	0x65, 0x93, 0xd8, 0x4a, 0x44, 0xb3, 0xb7, 0x70, 0x14, 0xea, 0xed, 0x40, 0x9a, 0x6f, 0x36, 0xe4,
	0x05, 0x7d, 0x09, 0x3d, 0xae, 0x95, 0x42, 0xee, 0xb4, 0x21, 0x89, 0xe7, 0x0b, 0xed, 0x48, 0x3a,
	0xfb, 0x11, 0xfa, 0xb7, 0x7b, 0xfd, 0xf5, 0xe0, 0x38, 0x9c, 0x0f, 0x79, 0xe1, 0x3f, 0x71, 0x9d,
	0xcb, 0x92, 0x24, 0xf4, 0x04, 0xba, 0x1f, 0x73, 0x8e, 0x85, 0xd6, 0x9f, 0x48, 0xcb, 0x0b, 0x5b,
	0x59, 0xa0, 0x21, 0xe9, 0xec, 0x9f, 0x04, 0xc8, 0xbb, 0xca, 0xad, 0xb4, 0x91, 0x5f, 0x50, 0x3c,
	0x7f, 0xf7, 0xdf, 0x40, 0x7d, 0xb8, 0xac, 0x15, 0xac, 0x0e, 0x4f, 0xaa, 0x56, 0x69, 0x06, 0x44,
	0x5a, 0x5b, 0xe1, 0x62, 0x6f, 0xcd, 0xe3, 0xc5, 0x0c, 0x02, 0xff, 0xee, 0x71, 0xd7, 0x47, 0xd0,
	0x0e, 0x8c, 0x69, 0xae, 0x27, 0xa2, 0xef, 0xff, 0x6e, 0xc1, 0xf1, 0x7b, 0xff, 0xbc, 0x74, 0x0c,
	0xe7, 0x31, 0xfa, 0x6d, 0xd8, 0xa1, 0x39, 0xfe, 0x51, 0xa1, 0x75, 0x28, 0xc8, 0xd9, 0x4e, 0xba,
	0xdf, 0x88, 0xff, 0x49, 0xe7, 0x3b, 0xe9, 0x67, 0x2c, 0x71, 0x5f, 0x1a, 0xd1, 0x0b, 0x38, 0xad,
	0xcb, 0xc5, 0x5c, 0xec, 0x84, 0x8b, 0x9d, 0xf0, 0x8b, 0xb4, 0x6e, 0x27, 0x30, 0x4a, 0x61, 0xd0,
	0xe4, 0xb1, 0x68, 0x3c, 0x37, 0xa6, 0xaf, 0xe0, 0xe5, 0x7e, 0x02, 0x41, 0x26, 0xf4, 0x14, 0x86,
	0x4d, 0xe0, 0xc6, 0xf7, 0x5b, 0x4a, 0xe0, 0x24, 0x92, 0xbf, 0x3d, 0x28, 0x8e, 0x82, 0xbc, 0xa6,
	0x67, 0x40, 0x22, 0xb3, 0x7b, 0x01, 0x72, 0x49, 0x47, 0x40, 0x9b, 0x78, 0xf9, 0x8e, 0x9f, 0xdd,
	0xfc, 0xd5, 0x82, 0x5e, 0x14, 0x7e, 0x5d, 0x1a, 0x9a, 0x41, 0x3b, 0x8e, 0x81, 0x1e, 0x8c, 0x7e,
	0x72, 0x80, 0xbd, 0x65, 0x9c, 0xca, 0x93, 0x96, 0x6f, 0xe0, 0xc8, 0x4f, 0xe2, 0x49, 0xbb, 0x9f,
	0xe0, 0xdc, 0x0f, 0xc6, 0x2f, 0xfb, 0xd2, 0xf8, 0x6b, 0x6e, 0xfe, 0x8c, 0xa1, 0xd9, 0x95, 0x3b,
	0x31, 0x61, 0x8d, 0xd3, 0x57, 0xeb, 0x55, 0xbb, 0x1f, 0xf2, 0xcf, 0x74, 0xff, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x6b, 0x60, 0x8d, 0x24, 0x47, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientMgrClient is the client API for ClientMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientMgrClient interface {
	Create(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error)
	Update(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error)
	Read(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error)
	ListIntegratedClients(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*AuthorizedClient, error)
	ListAuthorizedClients(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*AuthorizedClient, error)
}

type clientMgrClient struct {
	cc *grpc.ClientConn
}

func NewClientMgrClient(cc *grpc.ClientConn) ClientMgrClient {
	return &clientMgrClient{cc}
}

func (c *clientMgrClient) Create(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/client.ClientMgr/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientMgrClient) Update(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/client.ClientMgr/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientMgrClient) Read(ctx context.Context, in *Client, opts ...grpc.CallOption) (*Client, error) {
	out := new(Client)
	err := c.cc.Invoke(ctx, "/client.ClientMgr/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientMgrClient) ListIntegratedClients(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*AuthorizedClient, error) {
	out := new(AuthorizedClient)
	err := c.cc.Invoke(ctx, "/client.ClientMgr/ListIntegratedClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientMgrClient) ListAuthorizedClients(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*AuthorizedClient, error) {
	out := new(AuthorizedClient)
	err := c.cc.Invoke(ctx, "/client.ClientMgr/ListAuthorizedClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientMgrServer is the server API for ClientMgr service.
type ClientMgrServer interface {
	Create(context.Context, *Client) (*Client, error)
	Update(context.Context, *Client) (*Client, error)
	Read(context.Context, *Client) (*Client, error)
	ListIntegratedClients(context.Context, *common.Id) (*AuthorizedClient, error)
	ListAuthorizedClients(context.Context, *common.Id) (*AuthorizedClient, error)
}

func RegisterClientMgrServer(s *grpc.Server, srv ClientMgrServer) {
	s.RegisterService(&_ClientMgr_serviceDesc, srv)
}

func _ClientMgr_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientMgrServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientMgr/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientMgrServer).Create(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientMgr_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientMgrServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientMgr/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientMgrServer).Update(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientMgr_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Client)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientMgrServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientMgr/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientMgrServer).Read(ctx, req.(*Client))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientMgr_ListIntegratedClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientMgrServer).ListIntegratedClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientMgr/ListIntegratedClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientMgrServer).ListIntegratedClients(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientMgr_ListAuthorizedClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientMgrServer).ListAuthorizedClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientMgr/ListAuthorizedClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientMgrServer).ListAuthorizedClients(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client.ClientMgr",
	HandlerType: (*ClientMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClientMgr_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClientMgr_Update_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ClientMgr_Read_Handler,
		},
		{
			MethodName: "ListIntegratedClients",
			Handler:    _ClientMgr_ListIntegratedClients_Handler,
		},
		{
			MethodName: "ListAuthorizedClients",
			Handler:    _ClientMgr_ListAuthorizedClients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/client.proto",
}
