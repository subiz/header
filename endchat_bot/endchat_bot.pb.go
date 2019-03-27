// Code generated by protoc-gen-go. DO NOT EDIT.
// source: endchat_bot/endchat_bot.proto

package endchat_bot

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/subiz/header/common"
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

type ConnectorSetting struct {
	ConnectorId          string   `protobuf:"bytes,3,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	AtMidnight           bool     `protobuf:"varint,4,opt,name=at_midnight,json=atMidnight,proto3" json:"at_midnight,omitempty"`
	AfterAgentMessage    int64    `protobuf:"varint,6,opt,name=after_agent_message,json=afterAgentMessage,proto3" json:"after_agent_message,omitempty"`
	AfterAnyMessage      int64    `protobuf:"varint,7,opt,name=after_any_message,json=afterAnyMessage,proto3" json:"after_any_message,omitempty"`
	AfterUserMessage     int64    `protobuf:"varint,8,opt,name=after_user_message,json=afterUserMessage,proto3" json:"after_user_message,omitempty"`
	Age                  int64    `protobuf:"varint,10,opt,name=age,proto3" json:"age,omitempty"`
	Enabled              bool     `protobuf:"varint,11,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectorSetting) Reset()         { *m = ConnectorSetting{} }
func (m *ConnectorSetting) String() string { return proto.CompactTextString(m) }
func (*ConnectorSetting) ProtoMessage()    {}
func (*ConnectorSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b426c48535cb43, []int{0}
}

func (m *ConnectorSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectorSetting.Unmarshal(m, b)
}
func (m *ConnectorSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectorSetting.Marshal(b, m, deterministic)
}
func (m *ConnectorSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectorSetting.Merge(m, src)
}
func (m *ConnectorSetting) XXX_Size() int {
	return xxx_messageInfo_ConnectorSetting.Size(m)
}
func (m *ConnectorSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectorSetting.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectorSetting proto.InternalMessageInfo

func (m *ConnectorSetting) GetConnectorId() string {
	if m != nil {
		return m.ConnectorId
	}
	return ""
}

func (m *ConnectorSetting) GetAtMidnight() bool {
	if m != nil {
		return m.AtMidnight
	}
	return false
}

func (m *ConnectorSetting) GetAfterAgentMessage() int64 {
	if m != nil {
		return m.AfterAgentMessage
	}
	return 0
}

func (m *ConnectorSetting) GetAfterAnyMessage() int64 {
	if m != nil {
		return m.AfterAnyMessage
	}
	return 0
}

func (m *ConnectorSetting) GetAfterUserMessage() int64 {
	if m != nil {
		return m.AfterUserMessage
	}
	return 0
}

func (m *ConnectorSetting) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *ConnectorSetting) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type Setting struct {
	AccountId            string              `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ConnectorSettings    []*ConnectorSetting `protobuf:"bytes,3,rep,name=connector_settings,json=connectorSettings,proto3" json:"connector_settings,omitempty"`
	GlobalSetting        *ConnectorSetting   `protobuf:"bytes,4,opt,name=global_setting,json=globalSetting,proto3" json:"global_setting,omitempty"`
	Updated              int64               `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Setting) Reset()         { *m = Setting{} }
func (m *Setting) String() string { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()    {}
func (*Setting) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b426c48535cb43, []int{1}
}

func (m *Setting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Setting.Unmarshal(m, b)
}
func (m *Setting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Setting.Marshal(b, m, deterministic)
}
func (m *Setting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Setting.Merge(m, src)
}
func (m *Setting) XXX_Size() int {
	return xxx_messageInfo_Setting.Size(m)
}
func (m *Setting) XXX_DiscardUnknown() {
	xxx_messageInfo_Setting.DiscardUnknown(m)
}

var xxx_messageInfo_Setting proto.InternalMessageInfo

func (m *Setting) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Setting) GetConnectorSettings() []*ConnectorSetting {
	if m != nil {
		return m.ConnectorSettings
	}
	return nil
}

func (m *Setting) GetGlobalSetting() *ConnectorSetting {
	if m != nil {
		return m.GlobalSetting
	}
	return nil
}

func (m *Setting) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

// (create_min, account_id, id)
// (account_id, id)
type Conversation struct {
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	LastMessageSent      int64    `protobuf:"varint,5,opt,name=last_message_sent,json=lastMessageSent,proto3" json:"last_message_sent,omitempty"`
	LastUserMessageSent  int64    `protobuf:"varint,6,opt,name=last_user_message_sent,json=lastUserMessageSent,proto3" json:"last_user_message_sent,omitempty"`
	LastAgentMessageSent int64    `protobuf:"varint,7,opt,name=last_agent_message_sent,json=lastAgentMessageSent,proto3" json:"last_agent_message_sent,omitempty"`
	Ended                int64    `protobuf:"varint,8,opt,name=ended,proto3" json:"ended,omitempty"`
	EndRequested         int64    `protobuf:"varint,9,opt,name=end_requested,json=endRequested,proto3" json:"end_requested,omitempty"`
	State                string   `protobuf:"bytes,10,opt,name=state,proto3" json:"state,omitempty"`
	IntegrationId        string   `protobuf:"bytes,11,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	ConnectorId          string   `protobuf:"bytes,12,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Conversation) Reset()         { *m = Conversation{} }
func (m *Conversation) String() string { return proto.CompactTextString(m) }
func (*Conversation) ProtoMessage()    {}
func (*Conversation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b426c48535cb43, []int{2}
}

func (m *Conversation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conversation.Unmarshal(m, b)
}
func (m *Conversation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conversation.Marshal(b, m, deterministic)
}
func (m *Conversation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conversation.Merge(m, src)
}
func (m *Conversation) XXX_Size() int {
	return xxx_messageInfo_Conversation.Size(m)
}
func (m *Conversation) XXX_DiscardUnknown() {
	xxx_messageInfo_Conversation.DiscardUnknown(m)
}

var xxx_messageInfo_Conversation proto.InternalMessageInfo

func (m *Conversation) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Conversation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Conversation) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Conversation) GetLastMessageSent() int64 {
	if m != nil {
		return m.LastMessageSent
	}
	return 0
}

func (m *Conversation) GetLastUserMessageSent() int64 {
	if m != nil {
		return m.LastUserMessageSent
	}
	return 0
}

func (m *Conversation) GetLastAgentMessageSent() int64 {
	if m != nil {
		return m.LastAgentMessageSent
	}
	return 0
}

func (m *Conversation) GetEnded() int64 {
	if m != nil {
		return m.Ended
	}
	return 0
}

func (m *Conversation) GetEndRequested() int64 {
	if m != nil {
		return m.EndRequested
	}
	return 0
}

func (m *Conversation) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Conversation) GetIntegrationId() string {
	if m != nil {
		return m.IntegrationId
	}
	return ""
}

func (m *Conversation) GetConnectorId() string {
	if m != nil {
		return m.ConnectorId
	}
	return ""
}

type EndCallback struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ConversationId       string          `protobuf:"bytes,3,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	ConnectorId          string          `protobuf:"bytes,4,opt,name=connector_id,json=connectorId,proto3" json:"connector_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EndCallback) Reset()         { *m = EndCallback{} }
func (m *EndCallback) String() string { return proto.CompactTextString(m) }
func (*EndCallback) ProtoMessage()    {}
func (*EndCallback) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b426c48535cb43, []int{3}
}

func (m *EndCallback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndCallback.Unmarshal(m, b)
}
func (m *EndCallback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndCallback.Marshal(b, m, deterministic)
}
func (m *EndCallback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndCallback.Merge(m, src)
}
func (m *EndCallback) XXX_Size() int {
	return xxx_messageInfo_EndCallback.Size(m)
}
func (m *EndCallback) XXX_DiscardUnknown() {
	xxx_messageInfo_EndCallback.DiscardUnknown(m)
}

var xxx_messageInfo_EndCallback proto.InternalMessageInfo

func (m *EndCallback) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *EndCallback) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *EndCallback) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

func (m *EndCallback) GetConnectorId() string {
	if m != nil {
		return m.ConnectorId
	}
	return ""
}

type MidnightCallback struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	MidnightSec          int64           `protobuf:"varint,3,opt,name=midnight_sec,json=midnightSec,proto3" json:"midnight_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MidnightCallback) Reset()         { *m = MidnightCallback{} }
func (m *MidnightCallback) String() string { return proto.CompactTextString(m) }
func (*MidnightCallback) ProtoMessage()    {}
func (*MidnightCallback) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b426c48535cb43, []int{4}
}

func (m *MidnightCallback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MidnightCallback.Unmarshal(m, b)
}
func (m *MidnightCallback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MidnightCallback.Marshal(b, m, deterministic)
}
func (m *MidnightCallback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MidnightCallback.Merge(m, src)
}
func (m *MidnightCallback) XXX_Size() int {
	return xxx_messageInfo_MidnightCallback.Size(m)
}
func (m *MidnightCallback) XXX_DiscardUnknown() {
	xxx_messageInfo_MidnightCallback.DiscardUnknown(m)
}

var xxx_messageInfo_MidnightCallback proto.InternalMessageInfo

func (m *MidnightCallback) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *MidnightCallback) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *MidnightCallback) GetMidnightSec() int64 {
	if m != nil {
		return m.MidnightSec
	}
	return 0
}

func init() {
	proto.RegisterType((*ConnectorSetting)(nil), "endchat_bot.ConnectorSetting")
	proto.RegisterType((*Setting)(nil), "endchat_bot.Setting")
	proto.RegisterType((*Conversation)(nil), "endchat_bot.Conversation")
	proto.RegisterType((*EndCallback)(nil), "endchat_bot.EndCallback")
	proto.RegisterType((*MidnightCallback)(nil), "endchat_bot.MidnightCallback")
}

func init() { proto.RegisterFile("endchat_bot/endchat_bot.proto", fileDescriptor_d3b426c48535cb43) }

var fileDescriptor_d3b426c48535cb43 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xd1, 0x6e, 0xd3, 0x4a,
	0x10, 0x95, 0xe3, 0xb6, 0x69, 0xc6, 0x69, 0x92, 0x6e, 0xab, 0x7b, 0x2d, 0xa4, 0x8a, 0x34, 0x08,
	0x11, 0x01, 0x4a, 0x24, 0x2a, 0x3e, 0x00, 0x05, 0x1e, 0x22, 0xd1, 0x17, 0x47, 0x3c, 0x5b, 0xeb,
	0xdd, 0xc1, 0xb1, 0x70, 0x76, 0x8b, 0x77, 0x0d, 0x2d, 0x9f, 0x80, 0xf8, 0x03, 0xfe, 0x8b, 0xef,
	0x41, 0xbb, 0xeb, 0x4d, 0x4d, 0x83, 0x94, 0x17, 0x9e, 0xe2, 0x39, 0xe7, 0xcc, 0x64, 0x3c, 0x67,
	0xc6, 0x70, 0x81, 0x82, 0xb3, 0x35, 0xd5, 0x69, 0x26, 0xf5, 0xbc, 0xf5, 0x3c, 0xbb, 0xa9, 0xa4,
	0x96, 0x24, 0x6a, 0x41, 0x8f, 0x5e, 0xe4, 0x85, 0x5e, 0xd7, 0xd9, 0x8c, 0xc9, 0xcd, 0x5c, 0xd5,
	0x59, 0xf1, 0x6d, 0xbe, 0x46, 0xca, 0xb1, 0x9a, 0x33, 0xb9, 0xd9, 0x48, 0xd1, 0xfc, 0xb8, 0xcc,
	0xc9, 0x8f, 0x0e, 0x8c, 0x16, 0x52, 0x08, 0x64, 0x5a, 0x56, 0x2b, 0xd4, 0xba, 0x10, 0x39, 0xb9,
	0x84, 0x3e, 0xf3, 0x58, 0x5a, 0xf0, 0x38, 0x1c, 0x07, 0xd3, 0x5e, 0x12, 0x6d, 0xb1, 0x25, 0x27,
	0x8f, 0x21, 0xa2, 0x3a, 0xdd, 0x14, 0x5c, 0x14, 0xf9, 0x5a, 0xc7, 0x07, 0xe3, 0x60, 0x7a, 0x9c,
	0x00, 0xd5, 0xd7, 0x0d, 0x42, 0x66, 0x70, 0x46, 0x3f, 0x6a, 0xac, 0x52, 0x9a, 0xa3, 0xd0, 0xe9,
	0x06, 0x95, 0xa2, 0x39, 0xc6, 0x47, 0xe3, 0x60, 0x1a, 0x26, 0xa7, 0x96, 0x7a, 0x63, 0x98, 0x6b,
	0x47, 0x90, 0xe7, 0x70, 0xda, 0xe8, 0xc5, 0xdd, 0x56, 0xdd, 0xb5, 0xea, 0xa1, 0x53, 0x8b, 0x3b,
	0xaf, 0x7d, 0x09, 0xc4, 0x69, 0x6b, 0x85, 0xd5, 0x56, 0x7c, 0x6c, 0xc5, 0x23, 0xcb, 0x7c, 0x50,
	0x58, 0x79, 0xf5, 0x08, 0x42, 0x43, 0x83, 0xa5, 0xcd, 0x23, 0x89, 0xa1, 0x8b, 0x82, 0x66, 0x25,
	0xf2, 0x38, 0xb2, 0x8d, 0xfb, 0x70, 0xf2, 0x2b, 0x80, 0xae, 0x9f, 0xc2, 0x05, 0x00, 0x65, 0x4c,
	0xd6, 0x42, 0x9b, 0x19, 0x74, 0xec, 0x0c, 0x7a, 0x0d, 0xb2, 0xe4, 0xe4, 0x3d, 0x90, 0xfb, 0x21,
	0x29, 0x97, 0xa3, 0xe2, 0x70, 0x1c, 0x4e, 0xa3, 0x57, 0x17, 0xb3, 0xb6, 0x47, 0x0f, 0xe7, 0x9b,
	0x9c, 0xb2, 0x07, 0x88, 0x22, 0x6f, 0x61, 0x90, 0x97, 0x32, 0xa3, 0xa5, 0x2f, 0x65, 0x47, 0xba,
	0xb7, 0xd2, 0x89, 0x4b, 0xf2, 0x2d, 0xc7, 0xd0, 0xad, 0x6f, 0x38, 0xd5, 0xc8, 0xe3, 0x43, 0xfb,
	0xba, 0x3e, 0x9c, 0x7c, 0x0f, 0xa1, 0xbf, 0x90, 0xe2, 0x0b, 0x56, 0x8a, 0xea, 0x42, 0x8a, 0x7d,
	0x6f, 0x37, 0x80, 0xce, 0xd6, 0xf8, 0x4e, 0xc1, 0x4d, 0x65, 0x56, 0xa1, 0xad, 0x7c, 0xe0, 0x2a,
	0x37, 0xa1, 0x31, 0xae, 0xa4, 0x6a, 0xeb, 0x70, 0xaa, 0x50, 0xe8, 0xe6, 0xdf, 0x87, 0x86, 0x68,
	0x6c, 0x58, 0xa1, 0xd0, 0xe4, 0x0a, 0xfe, 0xb3, 0xda, 0xb6, 0x6f, 0x2e, 0xc1, 0xed, 0xc5, 0x99,
	0x61, 0x5b, 0xde, 0xd9, 0xa4, 0xd7, 0xf0, 0xbf, 0x4d, 0xfa, 0x63, 0x91, 0x5c, 0x96, 0xdb, 0x8f,
	0x73, 0x43, 0xb7, 0x97, 0xc9, 0xa6, 0x9d, 0xc3, 0x21, 0x0a, 0x8e, 0xbc, 0xd9, 0x0b, 0x17, 0x90,
	0x27, 0x70, 0x82, 0x82, 0xa7, 0x15, 0x7e, 0xae, 0x51, 0x99, 0xb7, 0xe9, 0x59, 0xb6, 0x8f, 0x82,
	0x27, 0x1e, 0x33, 0xa9, 0x4a, 0x53, 0xed, 0x76, 0xa6, 0x97, 0xb8, 0x80, 0x3c, 0x85, 0x41, 0x21,
	0x34, 0xe6, 0x95, 0x1d, 0xa0, 0x99, 0x5a, 0x64, 0xe9, 0x93, 0x16, 0xba, 0xe4, 0x3b, 0xc7, 0xd3,
	0xdf, 0x39, 0x9e, 0xc9, 0xcf, 0x00, 0xa2, 0x77, 0x82, 0x2f, 0x68, 0x59, 0x66, 0x94, 0x7d, 0x22,
	0x97, 0x10, 0x32, 0x7d, 0x1b, 0x07, 0xd6, 0xf1, 0xe1, 0xac, 0x39, 0xd0, 0x85, 0x14, 0x1a, 0x6f,
	0x75, 0x62, 0xb8, 0x7d, 0x76, 0x3d, 0x83, 0x21, 0x6b, 0xb9, 0x7b, 0x7f, 0xb4, 0x83, 0x36, 0xfc,
	0x97, 0xee, 0x0e, 0x76, 0xbb, 0xfb, 0x0a, 0x23, 0x7f, 0xc5, 0xff, 0xb0, 0xc3, 0x4b, 0xe8, 0xfb,
	0xaf, 0x45, 0xaa, 0x90, 0xd9, 0xf6, 0xc2, 0x24, 0xf2, 0xd8, 0x0a, 0x59, 0x76, 0x64, 0x3f, 0x49,
	0x57, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xf7, 0x3d, 0x57, 0xed, 0x04, 0x00, 0x00,
}
