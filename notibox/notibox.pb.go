// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notibox/notibox.proto

package notibox

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

type Type int32

const (
	Type_account_created                 Type = 0
	Type_trial_almost_expired            Type = 2
	Type_trial_expired                   Type = 3
	Type_system_maintainance_scheduled_1 Type = 4
	Type_system_maintainance_scheduled_2 Type = 5
	Type_system_maintainance_completed   Type = 6
	Type_agent_activated                 Type = 7
	Type_conversation_unassigned         Type = 8
	Type_agent_permission_updated        Type = 9
)

var Type_name = map[int32]string{
	0: "account_created",
	2: "trial_almost_expired",
	3: "trial_expired",
	4: "system_maintainance_scheduled_1",
	5: "system_maintainance_scheduled_2",
	6: "system_maintainance_completed",
	7: "agent_activated",
	8: "conversation_unassigned",
	9: "agent_permission_updated",
}

var Type_value = map[string]int32{
	"account_created":                 0,
	"trial_almost_expired":            2,
	"trial_expired":                   3,
	"system_maintainance_scheduled_1": 4,
	"system_maintainance_scheduled_2": 5,
	"system_maintainance_completed":   6,
	"agent_activated":                 7,
	"conversation_unassigned":         8,
	"agent_permission_updated":        9,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dfff48b43247ada2, []int{0}
}

type Event int32

const (
	Event_NotiboxRequested         Event = 0
	Event_NotiUpsertRequested      Event = 1
	Event_NotiReadRequested        Event = 2
	Event_NotiUnreadRequested      Event = 4
	Event_NotiListRequested        Event = 5
	Event_NotiDecreaseNewRequested Event = 6
	Event_NotiNewReadRequested     Event = 7
	Event_NotiboxSynced            Event = 101
)

var Event_name = map[int32]string{
	0:   "NotiboxRequested",
	1:   "NotiUpsertRequested",
	2:   "NotiReadRequested",
	4:   "NotiUnreadRequested",
	5:   "NotiListRequested",
	6:   "NotiDecreaseNewRequested",
	7:   "NotiNewReadRequested",
	101: "NotiboxSynced",
}

var Event_value = map[string]int32{
	"NotiboxRequested":         0,
	"NotiUpsertRequested":      1,
	"NotiReadRequested":        2,
	"NotiUnreadRequested":      4,
	"NotiListRequested":        5,
	"NotiDecreaseNewRequested": 6,
	"NotiNewReadRequested":     7,
	"NotiboxSynced":            101,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}

func (Event) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dfff48b43247ada2, []int{1}
}

// Notification represent an user's notification
type Notification struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AgentId              string          `protobuf:"bytes,3,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Topic                string          `protobuf:"bytes,10,opt,name=topic,proto3" json:"topic,omitempty"`
	Type                 string          `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Data                 string          `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	Created              int64           `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"`
	Read                 int64           `protobuf:"varint,9,opt,name=read,proto3" json:"read,omitempty"`
	Seen                 bool            `protobuf:"varint,12,opt,name=seen,proto3" json:"seen,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfff48b43247ada2, []int{0}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Notification) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Notification) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *Notification) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Notification) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Notification) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Notification) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Notification) GetRead() int64 {
	if m != nil {
		return m.Read
	}
	return 0
}

func (m *Notification) GetSeen() bool {
	if m != nil {
		return m.Seen
	}
	return false
}

type Notibox struct {
	Ctx                           *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId                     string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AgentId                       string          `protobuf:"bytes,3,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	NewCount                      int64           `protobuf:"varint,5,opt,name=new_count,json=newCount,proto3" json:"new_count,omitempty"`
	StartAnchor                   string          `protobuf:"bytes,6,opt,name=start_anchor,json=startAnchor,proto3" json:"start_anchor,omitempty"`
	TrialExpired                  bool            `protobuf:"varint,10,opt,name=trial_expired,json=trialExpired,proto3" json:"trial_expired,omitempty"`
	SystemMaintainanceScheduled_1 bool            `protobuf:"varint,11,opt,name=system_maintainance_scheduled_1,json=systemMaintainanceScheduled1,proto3" json:"system_maintainance_scheduled_1,omitempty"`
	SystemMaintainanceScheduled_2 bool            `protobuf:"varint,12,opt,name=system_maintainance_scheduled_2,json=systemMaintainanceScheduled2,proto3" json:"system_maintainance_scheduled_2,omitempty"`
	SystemMaintainanceCompleted   bool            `protobuf:"varint,13,opt,name=system_maintainance_completed,json=systemMaintainanceCompleted,proto3" json:"system_maintainance_completed,omitempty"`
	ConversationUnassigned        bool            `protobuf:"varint,14,opt,name=conversation_unassigned,json=conversationUnassigned,proto3" json:"conversation_unassigned,omitempty"`
	AgentPermissionUpdated        bool            `protobuf:"varint,15,opt,name=agent_permission_updated,json=agentPermissionUpdated,proto3" json:"agent_permission_updated,omitempty"`
	AccountCreated                bool            `protobuf:"varint,16,opt,name=account_created,json=accountCreated,proto3" json:"account_created,omitempty"`
	TrialAlmostExpired            bool            `protobuf:"varint,17,opt,name=trial_almost_expired,json=trialAlmostExpired,proto3" json:"trial_almost_expired,omitempty"`
	AgentActivated                bool            `protobuf:"varint,18,opt,name=agent_activated,json=agentActivated,proto3" json:"agent_activated,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}        `json:"-"`
	XXX_unrecognized              []byte          `json:"-"`
	XXX_sizecache                 int32           `json:"-"`
}

func (m *Notibox) Reset()         { *m = Notibox{} }
func (m *Notibox) String() string { return proto.CompactTextString(m) }
func (*Notibox) ProtoMessage()    {}
func (*Notibox) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfff48b43247ada2, []int{1}
}

func (m *Notibox) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notibox.Unmarshal(m, b)
}
func (m *Notibox) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notibox.Marshal(b, m, deterministic)
}
func (m *Notibox) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notibox.Merge(m, src)
}
func (m *Notibox) XXX_Size() int {
	return xxx_messageInfo_Notibox.Size(m)
}
func (m *Notibox) XXX_DiscardUnknown() {
	xxx_messageInfo_Notibox.DiscardUnknown(m)
}

var xxx_messageInfo_Notibox proto.InternalMessageInfo

func (m *Notibox) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Notibox) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Notibox) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *Notibox) GetNewCount() int64 {
	if m != nil {
		return m.NewCount
	}
	return 0
}

func (m *Notibox) GetStartAnchor() string {
	if m != nil {
		return m.StartAnchor
	}
	return ""
}

func (m *Notibox) GetTrialExpired() bool {
	if m != nil {
		return m.TrialExpired
	}
	return false
}

func (m *Notibox) GetSystemMaintainanceScheduled_1() bool {
	if m != nil {
		return m.SystemMaintainanceScheduled_1
	}
	return false
}

func (m *Notibox) GetSystemMaintainanceScheduled_2() bool {
	if m != nil {
		return m.SystemMaintainanceScheduled_2
	}
	return false
}

func (m *Notibox) GetSystemMaintainanceCompleted() bool {
	if m != nil {
		return m.SystemMaintainanceCompleted
	}
	return false
}

func (m *Notibox) GetConversationUnassigned() bool {
	if m != nil {
		return m.ConversationUnassigned
	}
	return false
}

func (m *Notibox) GetAgentPermissionUpdated() bool {
	if m != nil {
		return m.AgentPermissionUpdated
	}
	return false
}

func (m *Notibox) GetAccountCreated() bool {
	if m != nil {
		return m.AccountCreated
	}
	return false
}

func (m *Notibox) GetTrialAlmostExpired() bool {
	if m != nil {
		return m.TrialAlmostExpired
	}
	return false
}

func (m *Notibox) GetAgentActivated() bool {
	if m != nil {
		return m.AgentActivated
	}
	return false
}

type AddNotificationRequest struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountIds           []string        `protobuf:"bytes,2,rep,name=account_ids,json=accountIds,proto3" json:"account_ids,omitempty"`
	AgentIds             []string        `protobuf:"bytes,4,rep,name=agent_ids,json=agentIds,proto3" json:"agent_ids,omitempty"`
	Notification         *Notification   `protobuf:"bytes,3,opt,name=notification,proto3" json:"notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AddNotificationRequest) Reset()         { *m = AddNotificationRequest{} }
func (m *AddNotificationRequest) String() string { return proto.CompactTextString(m) }
func (*AddNotificationRequest) ProtoMessage()    {}
func (*AddNotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfff48b43247ada2, []int{2}
}

func (m *AddNotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddNotificationRequest.Unmarshal(m, b)
}
func (m *AddNotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddNotificationRequest.Marshal(b, m, deterministic)
}
func (m *AddNotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNotificationRequest.Merge(m, src)
}
func (m *AddNotificationRequest) XXX_Size() int {
	return xxx_messageInfo_AddNotificationRequest.Size(m)
}
func (m *AddNotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddNotificationRequest proto.InternalMessageInfo

func (m *AddNotificationRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *AddNotificationRequest) GetAccountIds() []string {
	if m != nil {
		return m.AccountIds
	}
	return nil
}

func (m *AddNotificationRequest) GetAgentIds() []string {
	if m != nil {
		return m.AgentIds
	}
	return nil
}

func (m *AddNotificationRequest) GetNotification() *Notification {
	if m != nil {
		return m.Notification
	}
	return nil
}

type ListRequest struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AgentId              string          `protobuf:"bytes,3,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Anchor               string          `protobuf:"bytes,4,opt,name=anchor,proto3" json:"anchor,omitempty"`
	Limit                int32           `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Seen                 bool            `protobuf:"varint,6,opt,name=seen,proto3" json:"seen,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfff48b43247ada2, []int{3}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ListRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *ListRequest) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *ListRequest) GetAnchor() string {
	if m != nil {
		return m.Anchor
	}
	return ""
}

func (m *ListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetSeen() bool {
	if m != nil {
		return m.Seen
	}
	return false
}

type TopicRequest struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AgentId              string          `protobuf:"bytes,3,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Topics               []string        `protobuf:"bytes,4,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TopicRequest) Reset()         { *m = TopicRequest{} }
func (m *TopicRequest) String() string { return proto.CompactTextString(m) }
func (*TopicRequest) ProtoMessage()    {}
func (*TopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfff48b43247ada2, []int{4}
}

func (m *TopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicRequest.Unmarshal(m, b)
}
func (m *TopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicRequest.Marshal(b, m, deterministic)
}
func (m *TopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicRequest.Merge(m, src)
}
func (m *TopicRequest) XXX_Size() int {
	return xxx_messageInfo_TopicRequest.Size(m)
}
func (m *TopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TopicRequest proto.InternalMessageInfo

func (m *TopicRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *TopicRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *TopicRequest) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *TopicRequest) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type Notifications struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Notifications        []*Notification `protobuf:"bytes,2,rep,name=notifications,proto3" json:"notifications,omitempty"`
	Anchor               string          `protobuf:"bytes,3,opt,name=anchor,proto3" json:"anchor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Notifications) Reset()         { *m = Notifications{} }
func (m *Notifications) String() string { return proto.CompactTextString(m) }
func (*Notifications) ProtoMessage()    {}
func (*Notifications) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfff48b43247ada2, []int{5}
}

func (m *Notifications) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notifications.Unmarshal(m, b)
}
func (m *Notifications) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notifications.Marshal(b, m, deterministic)
}
func (m *Notifications) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notifications.Merge(m, src)
}
func (m *Notifications) XXX_Size() int {
	return xxx_messageInfo_Notifications.Size(m)
}
func (m *Notifications) XXX_DiscardUnknown() {
	xxx_messageInfo_Notifications.DiscardUnknown(m)
}

var xxx_messageInfo_Notifications proto.InternalMessageInfo

func (m *Notifications) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Notifications) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

func (m *Notifications) GetAnchor() string {
	if m != nil {
		return m.Anchor
	}
	return ""
}

func init() {
	proto.RegisterEnum("notibox.Type", Type_name, Type_value)
	proto.RegisterEnum("notibox.Event", Event_name, Event_value)
	proto.RegisterType((*Notification)(nil), "notibox.Notification")
	proto.RegisterType((*Notibox)(nil), "notibox.Notibox")
	proto.RegisterType((*AddNotificationRequest)(nil), "notibox.AddNotificationRequest")
	proto.RegisterType((*ListRequest)(nil), "notibox.ListRequest")
	proto.RegisterType((*TopicRequest)(nil), "notibox.TopicRequest")
	proto.RegisterType((*Notifications)(nil), "notibox.Notifications")
}

func init() { proto.RegisterFile("notibox/notibox.proto", fileDescriptor_dfff48b43247ada2) }

var fileDescriptor_dfff48b43247ada2 = []byte{
	// 803 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x66, 0xfc, 0xef, 0xb2, 0xb3, 0x99, 0xf4, 0x26, 0xd9, 0x81, 0xec, 0x6a, 0x1d, 0xef, 0x01,
	0x6b, 0x91, 0x12, 0xd6, 0x1c, 0x00, 0x71, 0x0a, 0x21, 0x87, 0x95, 0x20, 0x42, 0xb3, 0x9b, 0xf3,
	0xa8, 0x33, 0x5d, 0x24, 0x2d, 0x79, 0xba, 0x87, 0xe9, 0x76, 0xe2, 0x70, 0x45, 0x88, 0x17, 0xe0,
	0x31, 0x38, 0xf2, 0x10, 0xbc, 0x0b, 0x2f, 0x81, 0xba, 0xa6, 0xc7, 0x1e, 0xb3, 0x1b, 0xbc, 0x20,
	0xc1, 0xc9, 0x5d, 0xdf, 0x57, 0x5d, 0xae, 0xae, 0xfa, 0xaa, 0x06, 0xf6, 0x94, 0xb6, 0xf2, 0x52,
	0x2f, 0x8e, 0xfd, 0xef, 0x51, 0x5e, 0x68, 0xab, 0x59, 0xd7, 0x9b, 0x1f, 0x7c, 0x74, 0x25, 0xed,
	0xf5, 0xfc, 0xf2, 0x28, 0xd5, 0xd9, 0xb1, 0x99, 0x5f, 0xca, 0x1f, 0x8e, 0xaf, 0x91, 0x0b, 0x2c,
	0x8e, 0x53, 0x9d, 0x65, 0x5a, 0xf9, 0x9f, 0xf2, 0xd6, 0xf8, 0x8f, 0x00, 0x86, 0xe7, 0xda, 0xca,
	0xef, 0x64, 0xca, 0xad, 0xd4, 0x8a, 0x1d, 0x42, 0x33, 0xb5, 0x8b, 0x28, 0x18, 0x05, 0x93, 0xc1,
	0x74, 0xfb, 0xc8, 0x3b, 0x9f, 0x6a, 0x65, 0x71, 0x61, 0x63, 0xc7, 0xb1, 0x27, 0x00, 0x3c, 0x4d,
	0xf5, 0x5c, 0xd9, 0x44, 0x8a, 0xa8, 0x31, 0x0a, 0x26, 0xfd, 0xb8, 0xef, 0x91, 0x97, 0x82, 0xbd,
	0x0f, 0x3d, 0x7e, 0x85, 0x25, 0xd9, 0x24, 0xb2, 0x4b, 0xf6, 0x4b, 0xc1, 0x76, 0xa1, 0x6d, 0x75,
	0x2e, 0xd3, 0x08, 0x08, 0x2f, 0x0d, 0xc6, 0xa0, 0x65, 0xef, 0x72, 0x8c, 0xda, 0x04, 0xd2, 0xd9,
	0x61, 0x82, 0x5b, 0x1e, 0xf5, 0x4a, 0xcc, 0x9d, 0x59, 0x04, 0xdd, 0xb4, 0x40, 0x6e, 0x51, 0x44,
	0x9d, 0x51, 0x30, 0x69, 0xc6, 0x95, 0xe9, 0xbc, 0x0b, 0xe4, 0x22, 0xea, 0x13, 0x4c, 0x67, 0x87,
	0x19, 0x44, 0x15, 0x0d, 0x47, 0xc1, 0xa4, 0x17, 0xd3, 0x79, 0xfc, 0x4b, 0x1b, 0xba, 0xe7, 0x65,
	0x99, 0xfe, 0xdb, 0x87, 0x1e, 0x40, 0x5f, 0xe1, 0x6d, 0x42, 0x9e, 0xf4, 0xae, 0x66, 0xdc, 0x53,
	0x78, 0x7b, 0xea, 0x6c, 0x76, 0x08, 0x43, 0x63, 0x79, 0x61, 0x13, 0xae, 0xd2, 0x6b, 0x5d, 0xd0,
	0x63, 0xfa, 0xf1, 0x80, 0xb0, 0x13, 0x82, 0xd8, 0x33, 0xd8, 0xb2, 0x85, 0xe4, 0xb3, 0x04, 0x17,
	0xb9, 0x2c, 0x50, 0x50, 0xc1, 0x7a, 0xf1, 0x90, 0xc0, 0xb3, 0x12, 0x63, 0x67, 0xf0, 0xd4, 0xdc,
	0x19, 0x8b, 0x59, 0x92, 0x71, 0xa9, 0x2c, 0x97, 0x8a, 0xab, 0x14, 0x13, 0x93, 0x5e, 0xa3, 0x98,
	0xcf, 0x50, 0x24, 0x2f, 0xa2, 0x01, 0x5d, 0x7b, 0x5c, 0xba, 0x7d, 0x53, 0xf3, 0x7a, 0x55, 0x39,
	0xbd, 0xd8, 0x1c, 0x66, 0xea, 0x6b, 0xf8, 0x77, 0x61, 0xa6, 0xec, 0x4b, 0x78, 0xf2, 0xb6, 0x30,
	0xa9, 0xce, 0xf2, 0x19, 0xba, 0x9e, 0x6d, 0x51, 0x90, 0x83, 0x37, 0x83, 0x9c, 0x56, 0x2e, 0xec,
	0x53, 0x78, 0x94, 0x6a, 0x75, 0x83, 0x85, 0x21, 0x31, 0x26, 0x73, 0xc5, 0x8d, 0x91, 0x57, 0x0a,
	0x45, 0xf4, 0x80, 0x6e, 0xef, 0xd7, 0xe9, 0x8b, 0x25, 0xcb, 0x3e, 0x83, 0xa8, 0x6c, 0x45, 0x8e,
	0x45, 0x26, 0x8d, 0xa1, 0xcb, 0xb9, 0x20, 0xad, 0x6c, 0x97, 0x37, 0x89, 0xff, 0x76, 0x49, 0x5f,
	0x94, 0x2c, 0xfb, 0x10, 0xb6, 0xab, 0x1e, 0x57, 0xe2, 0x0a, 0xe9, 0xc2, 0x03, 0x0f, 0x9f, 0x7a,
	0x8d, 0x7d, 0x0c, 0xbb, 0x65, 0x4b, 0xf8, 0x2c, 0xd3, 0xc6, 0x2e, 0x3b, 0xb3, 0x43, 0xde, 0x8c,
	0xb8, 0x13, 0xa2, 0xaa, 0xfe, 0xb8, 0xd0, 0x94, 0x14, 0x4f, 0xad, 0xbc, 0xa1, 0xd0, 0xcc, 0x87,
	0x76, 0xf0, 0x49, 0x85, 0x8e, 0x7f, 0x0b, 0x60, 0xff, 0x44, 0x88, 0xfa, 0x1c, 0xc6, 0xf8, 0xfd,
	0x1c, 0x8d, 0x7d, 0x17, 0x95, 0x3e, 0x85, 0xc1, 0x4a, 0xa5, 0x26, 0x6a, 0x8c, 0x9a, 0x93, 0x7e,
	0x0c, 0x4b, 0x99, 0x1a, 0x27, 0xc6, 0x4a, 0xa7, 0x26, 0x6a, 0x11, 0xdd, 0xf3, 0x42, 0x35, 0xec,
	0x73, 0x18, 0xaa, 0xda, 0xff, 0x92, 0x90, 0x07, 0xd3, 0xbd, 0xa3, 0x6a, 0xb9, 0xac, 0x25, 0xb5,
	0xe6, 0x3a, 0xfe, 0x35, 0x80, 0xc1, 0xd7, 0xd2, 0xd8, 0x7f, 0x90, 0xeb, 0xbf, 0x9f, 0xa8, 0x7d,
	0xe8, 0xf8, 0x71, 0x69, 0x11, 0xe1, 0x2d, 0xb7, 0x52, 0x66, 0x32, 0x93, 0xe5, 0x94, 0xb5, 0xe3,
	0xd2, 0x58, 0x0e, 0x7f, 0xa7, 0x36, 0xfc, 0x3f, 0x06, 0x30, 0x7c, 0xed, 0x16, 0xce, 0xff, 0x95,
	0x2f, 0x6d, 0xb7, 0xaa, 0xe2, 0xde, 0x1a, 0xff, 0x1c, 0xc0, 0x56, 0xbd, 0xa6, 0xe6, 0x5d, 0xd2,
	0xf8, 0x02, 0xb6, 0xea, 0x95, 0x2f, 0x9b, 0x7c, 0x6f, 0x97, 0xd6, 0x7d, 0x6b, 0x95, 0x6b, 0xd6,
	0x2b, 0xf7, 0xfc, 0xa7, 0x06, 0xb4, 0x5e, 0xbb, 0x5d, 0xfb, 0xf0, 0x8d, 0x11, 0x08, 0xdf, 0x63,
	0xd1, 0xdb, 0xe5, 0x1e, 0x36, 0xd8, 0xce, 0x5f, 0x76, 0x53, 0xd8, 0x64, 0xcf, 0x36, 0x6e, 0xa2,
	0xb0, 0xb5, 0xd9, 0x69, 0x1a, 0xb6, 0xd9, 0xe1, 0x86, 0x2d, 0x12, 0x76, 0x28, 0xdd, 0xf5, 0xb1,
	0x0a, 0xbb, 0xec, 0xe0, 0xde, 0xcd, 0x11, 0xf6, 0xd8, 0xe3, 0xfb, 0xb7, 0x43, 0xd8, 0x7f, 0xfe,
	0x7b, 0x00, 0xed, 0xb3, 0x1b, 0x54, 0x96, 0xed, 0x42, 0xe8, 0xbf, 0x0e, 0x5e, 0x22, 0x54, 0x89,
	0x47, 0xf0, 0xd0, 0xa1, 0x17, 0xb9, 0xc1, 0xc2, 0xae, 0x88, 0x80, 0xed, 0xc1, 0x8e, 0x23, 0x62,
	0xe4, 0x62, 0x05, 0x37, 0x96, 0xfe, 0xaa, 0x58, 0x23, 0x5a, 0x95, 0x7f, 0x6d, 0x64, 0x50, 0x84,
	0x6d, 0x97, 0x9d, 0x83, 0xbf, 0x42, 0x57, 0x7c, 0x83, 0xe7, 0x78, 0xbb, 0x62, 0x3b, 0xae, 0x0f,
	0x8e, 0x25, 0xb4, 0x1e, 0xae, 0xeb, 0xfa, 0xe0, 0xb3, 0x7d, 0x75, 0xa7, 0x52, 0x14, 0x21, 0x5e,
	0x76, 0xe8, 0xa3, 0xfe, 0xc9, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x17, 0x3d, 0xaa, 0x3b, 0x23,
	0x08, 0x00, 0x00,
}
