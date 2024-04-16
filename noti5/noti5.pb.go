// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.3
// source: noti5.proto

package noti5

import (
	common "github.com/subiz/header/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Token_Platform int32

const (
	Token_desktop  Token_Platform = 0
	Token_mobile   Token_Platform = 1
	Token_ios_voip Token_Platform = 2
)

// Enum value maps for Token_Platform.
var (
	Token_Platform_name = map[int32]string{
		0: "desktop",
		1: "mobile",
		2: "ios_voip",
	}
	Token_Platform_value = map[string]int32{
		"desktop":  0,
		"mobile":   1,
		"ios_voip": 2,
	}
)

func (x Token_Platform) Enum() *Token_Platform {
	p := new(Token_Platform)
	*p = x
	return p
}

func (x Token_Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Token_Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_noti5_proto_enumTypes[0].Descriptor()
}

func (Token_Platform) Type() protoreflect.EnumType {
	return &file_noti5_proto_enumTypes[0]
}

func (x Token_Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Token_Platform) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Token_Platform(num)
	return nil
}

// Deprecated: Use Token_Platform.Descriptor instead.
func (Token_Platform) EnumDescriptor() ([]byte, []int) {
	return file_noti5_proto_rawDescGZIP(), []int{3, 0}
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewMessage            *bool  `protobuf:"varint,3,opt,name=new_message,json=newMessage" json:"new_message,omitempty"`
	UserCreated           *bool  `protobuf:"varint,6,opt,name=user_created,json=userCreated" json:"user_created,omitempty"`
	UserReturned          *bool  `protobuf:"varint,7,opt,name=user_returned,json=userReturned" json:"user_returned,omitempty"`
	CampaignUserConverted *bool  `protobuf:"varint,8,opt,name=campaign_user_converted,json=campaignUserConverted" json:"campaign_user_converted,omitempty"`
	UserOpenedEmail       *bool  `protobuf:"varint,9,opt,name=user_opened_email,json=userOpenedEmail" json:"user_opened_email,omitempty"`
	Delay                 *int32 `protobuf:"varint,10,opt,name=delay" json:"delay,omitempty"`
	PrivateMessage        *bool  `protobuf:"varint,11,opt,name=private_message,json=privateMessage" json:"private_message,omitempty"`
	ZaloOaFollowed        *bool  `protobuf:"varint,12,opt,name=zalo_oa_followed,json=zaloOaFollowed" json:"zalo_oa_followed,omitempty"`
	LeadCreated           *bool  `protobuf:"varint,13,opt,name=lead_created,json=leadCreated" json:"lead_created,omitempty"`
	TaskUpdated           *int64 `protobuf:"varint,14,opt,name=task_updated,json=taskUpdated" json:"task_updated,omitempty"`
	NewAssignedOrder      *int64 `protobuf:"varint,15,opt,name=new_assigned_order,json=newAssignedOrder" json:"new_assigned_order,omitempty"`
	LeadAssigned          *int64 `protobuf:"varint,16,opt,name=lead_assigned,json=leadAssigned" json:"lead_assigned,omitempty"`
	SegmentJoined         *int64 `protobuf:"varint,17,opt,name=segment_joined,json=segmentJoined" json:"segment_joined,omitempty"`
	TicketUpdated         *int64 `protobuf:"varint,18,opt,name=ticket_updated,json=ticketUpdated" json:"ticket_updated,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noti5_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_noti5_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_noti5_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetNewMessage() bool {
	if x != nil && x.NewMessage != nil {
		return *x.NewMessage
	}
	return false
}

func (x *Subscription) GetUserCreated() bool {
	if x != nil && x.UserCreated != nil {
		return *x.UserCreated
	}
	return false
}

func (x *Subscription) GetUserReturned() bool {
	if x != nil && x.UserReturned != nil {
		return *x.UserReturned
	}
	return false
}

func (x *Subscription) GetCampaignUserConverted() bool {
	if x != nil && x.CampaignUserConverted != nil {
		return *x.CampaignUserConverted
	}
	return false
}

func (x *Subscription) GetUserOpenedEmail() bool {
	if x != nil && x.UserOpenedEmail != nil {
		return *x.UserOpenedEmail
	}
	return false
}

func (x *Subscription) GetDelay() int32 {
	if x != nil && x.Delay != nil {
		return *x.Delay
	}
	return 0
}

func (x *Subscription) GetPrivateMessage() bool {
	if x != nil && x.PrivateMessage != nil {
		return *x.PrivateMessage
	}
	return false
}

func (x *Subscription) GetZaloOaFollowed() bool {
	if x != nil && x.ZaloOaFollowed != nil {
		return *x.ZaloOaFollowed
	}
	return false
}

func (x *Subscription) GetLeadCreated() bool {
	if x != nil && x.LeadCreated != nil {
		return *x.LeadCreated
	}
	return false
}

func (x *Subscription) GetTaskUpdated() int64 {
	if x != nil && x.TaskUpdated != nil {
		return *x.TaskUpdated
	}
	return 0
}

func (x *Subscription) GetNewAssignedOrder() int64 {
	if x != nil && x.NewAssignedOrder != nil {
		return *x.NewAssignedOrder
	}
	return 0
}

func (x *Subscription) GetLeadAssigned() int64 {
	if x != nil && x.LeadAssigned != nil {
		return *x.LeadAssigned
	}
	return 0
}

func (x *Subscription) GetSegmentJoined() int64 {
	if x != nil && x.SegmentJoined != nil {
		return *x.SegmentJoined
	}
	return 0
}

func (x *Subscription) GetTicketUpdated() int64 {
	if x != nil && x.TicketUpdated != nil {
		return *x.TicketUpdated
	}
	return 0
}

type Setting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	AgentId          *string         `protobuf:"bytes,3,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Updated          *int64          `protobuf:"varint,8,opt,name=updated" json:"updated,omitempty"`
	Web              *Subscription   `protobuf:"bytes,9,opt,name=web" json:"web,omitempty"` // desktop
	Mobile           *Subscription   `protobuf:"bytes,10,opt,name=mobile" json:"mobile,omitempty"`
	Email            *Subscription   `protobuf:"bytes,11,opt,name=email" json:"email,omitempty"`
	Instant          *Subscription   `protobuf:"bytes,15,opt,name=instant" json:"instant,omitempty"` // web
	DoNotDisturb     *DoNotDisturb   `protobuf:"bytes,12,opt,name=do_not_disturb,json=doNotDisturb" json:"do_not_disturb,omitempty"`
	InstantMuteUntil *int64          `protobuf:"varint,14,opt,name=instant_mute_until,json=instantMuteUntil" json:"instant_mute_until,omitempty"`
}

func (x *Setting) Reset() {
	*x = Setting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noti5_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Setting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Setting) ProtoMessage() {}

func (x *Setting) ProtoReflect() protoreflect.Message {
	mi := &file_noti5_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Setting.ProtoReflect.Descriptor instead.
func (*Setting) Descriptor() ([]byte, []int) {
	return file_noti5_proto_rawDescGZIP(), []int{1}
}

func (x *Setting) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *Setting) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *Setting) GetAgentId() string {
	if x != nil && x.AgentId != nil {
		return *x.AgentId
	}
	return ""
}

func (x *Setting) GetUpdated() int64 {
	if x != nil && x.Updated != nil {
		return *x.Updated
	}
	return 0
}

func (x *Setting) GetWeb() *Subscription {
	if x != nil {
		return x.Web
	}
	return nil
}

func (x *Setting) GetMobile() *Subscription {
	if x != nil {
		return x.Mobile
	}
	return nil
}

func (x *Setting) GetEmail() *Subscription {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *Setting) GetInstant() *Subscription {
	if x != nil {
		return x.Instant
	}
	return nil
}

func (x *Setting) GetDoNotDisturb() *DoNotDisturb {
	if x != nil {
		return x.DoNotDisturb
	}
	return nil
}

func (x *Setting) GetInstantMuteUntil() int64 {
	if x != nil && x.InstantMuteUntil != nil {
		return *x.InstantMuteUntil
	}
	return 0
}

type DoNotDisturb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DailyFrom   *int64 `protobuf:"varint,2,opt,name=daily_from,json=dailyFrom" json:"daily_from,omitempty"`       // min
	DailyTo     *int64 `protobuf:"varint,3,opt,name=daily_to,json=dailyTo" json:"daily_to,omitempty"`             // min
	EnableDaily *bool  `protobuf:"varint,5,opt,name=enable_daily,json=enableDaily" json:"enable_daily,omitempty"` // enable daily
	Until       *int64 `protobuf:"varint,4,opt,name=until" json:"until,omitempty"`                                // ms
}

func (x *DoNotDisturb) Reset() {
	*x = DoNotDisturb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noti5_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoNotDisturb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoNotDisturb) ProtoMessage() {}

func (x *DoNotDisturb) ProtoReflect() protoreflect.Message {
	mi := &file_noti5_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoNotDisturb.ProtoReflect.Descriptor instead.
func (*DoNotDisturb) Descriptor() ([]byte, []int) {
	return file_noti5_proto_rawDescGZIP(), []int{2}
}

func (x *DoNotDisturb) GetDailyFrom() int64 {
	if x != nil && x.DailyFrom != nil {
		return *x.DailyFrom
	}
	return 0
}

func (x *DoNotDisturb) GetDailyTo() int64 {
	if x != nil && x.DailyTo != nil {
		return *x.DailyTo
	}
	return 0
}

func (x *DoNotDisturb) GetEnableDaily() bool {
	if x != nil && x.EnableDaily != nil {
		return *x.EnableDaily
	}
	return false
}

func (x *DoNotDisturb) GetUntil() int64 {
	if x != nil && x.Until != nil {
		return *x.Until
	}
	return 0
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ctx         *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId   *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId      *string         `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserType    *string         `protobuf:"bytes,4,opt,name=user_type,json=userType" json:"user_type,omitempty"`
	FcmToken    *string         `protobuf:"bytes,5,opt,name=fcm_token,json=fcmToken" json:"fcm_token,omitempty"`
	DeviceId    *string         `protobuf:"bytes,6,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Platform    *string         `protobuf:"bytes,7,opt,name=platform" json:"platform,omitempty"` // desktop, mobile, ios_voip
	Created     *int64          `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	Topics      []string        `protobuf:"bytes,9,rep,name=topics" json:"topics,omitempty"`
	AccessToken *string         `protobuf:"bytes,10,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noti5_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_noti5_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_noti5_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetCtx() *common.Context {
	if x != nil {
		return x.Ctx
	}
	return nil
}

func (x *Token) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *Token) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *Token) GetUserType() string {
	if x != nil && x.UserType != nil {
		return *x.UserType
	}
	return ""
}

func (x *Token) GetFcmToken() string {
	if x != nil && x.FcmToken != nil {
		return *x.FcmToken
	}
	return ""
}

func (x *Token) GetDeviceId() string {
	if x != nil && x.DeviceId != nil {
		return *x.DeviceId
	}
	return ""
}

func (x *Token) GetPlatform() string {
	if x != nil && x.Platform != nil {
		return *x.Platform
	}
	return ""
}

func (x *Token) GetCreated() int64 {
	if x != nil && x.Created != nil {
		return *x.Created
	}
	return 0
}

func (x *Token) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *Token) GetAccessToken() string {
	if x != nil && x.AccessToken != nil {
		return *x.AccessToken
	}
	return ""
}

var File_noti5_proto protoreflect.FileDescriptor

var file_noti5_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x69, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e,
	0x6f, 0x74, 0x69, 0x35, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xab, 0x04, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x17,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x63,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x65,
	0x6e, 0x65, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x7a, 0x61, 0x6c, 0x6f, 0x5f, 0x6f, 0x61, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x7a, 0x61, 0x6c, 0x6f, 0x4f,
	0x61, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x65, 0x61,
	0x64, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x6c, 0x65, 0x61, 0x64, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x2c, 0x0a, 0x12, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6e, 0x65, 0x77,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x0d, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6a, 0x6f,
	0x69, 0x6e, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x22, 0x97, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x03,
	0x63, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x03, 0x77, 0x65, 0x62, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x35, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x77, 0x65, 0x62, 0x12, 0x2b, 0x0a, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x35, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x35, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x2d, 0x0a, 0x07, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x35, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x12, 0x39, 0x0a, 0x0e, 0x64, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x74,
	0x75, 0x72, 0x62, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x35, 0x2e, 0x44, 0x6f, 0x4e, 0x6f, 0x74, 0x44, 0x69, 0x73, 0x74, 0x75, 0x72, 0x62, 0x52, 0x0c,
	0x64, 0x6f, 0x4e, 0x6f, 0x74, 0x44, 0x69, 0x73, 0x74, 0x75, 0x72, 0x62, 0x12, 0x2c, 0x0a, 0x12,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x75, 0x6e, 0x74,
	0x69, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x4d, 0x75, 0x74, 0x65, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x22, 0x81, 0x01, 0x0a, 0x0c, 0x44,
	0x6f, 0x4e, 0x6f, 0x74, 0x44, 0x69, 0x73, 0x74, 0x75, 0x72, 0x62, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x61, 0x69, 0x6c, 0x79, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x54, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x64, 0x61, 0x69, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x74, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x22, 0xdd,
	0x02, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x74, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x63, 0x74, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x63, 0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x63, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x31, 0x0a, 0x08, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x6b, 0x74,
	0x6f, 0x70, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x69, 0x6f, 0x73, 0x5f, 0x76, 0x6f, 0x69, 0x70, 0x10, 0x02, 0x42, 0x1f,
	0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x62,
	0x69, 0x7a, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x35,
}

var (
	file_noti5_proto_rawDescOnce sync.Once
	file_noti5_proto_rawDescData = file_noti5_proto_rawDesc
)

func file_noti5_proto_rawDescGZIP() []byte {
	file_noti5_proto_rawDescOnce.Do(func() {
		file_noti5_proto_rawDescData = protoimpl.X.CompressGZIP(file_noti5_proto_rawDescData)
	})
	return file_noti5_proto_rawDescData
}

var file_noti5_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_noti5_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_noti5_proto_goTypes = []interface{}{
	(Token_Platform)(0),    // 0: noti5.Token.Platform
	(*Subscription)(nil),   // 1: noti5.Subscription
	(*Setting)(nil),        // 2: noti5.Setting
	(*DoNotDisturb)(nil),   // 3: noti5.DoNotDisturb
	(*Token)(nil),          // 4: noti5.Token
	(*common.Context)(nil), // 5: common.Context
}
var file_noti5_proto_depIdxs = []int32{
	5, // 0: noti5.Setting.ctx:type_name -> common.Context
	1, // 1: noti5.Setting.web:type_name -> noti5.Subscription
	1, // 2: noti5.Setting.mobile:type_name -> noti5.Subscription
	1, // 3: noti5.Setting.email:type_name -> noti5.Subscription
	1, // 4: noti5.Setting.instant:type_name -> noti5.Subscription
	3, // 5: noti5.Setting.do_not_disturb:type_name -> noti5.DoNotDisturb
	5, // 6: noti5.Token.ctx:type_name -> common.Context
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_noti5_proto_init() }
func file_noti5_proto_init() {
	if File_noti5_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_noti5_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_noti5_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Setting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_noti5_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoNotDisturb); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_noti5_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_noti5_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_noti5_proto_goTypes,
		DependencyIndexes: file_noti5_proto_depIdxs,
		EnumInfos:         file_noti5_proto_enumTypes,
		MessageInfos:      file_noti5_proto_msgTypes,
	}.Build()
	File_noti5_proto = out.File
	file_noti5_proto_rawDesc = nil
	file_noti5_proto_goTypes = nil
	file_noti5_proto_depIdxs = nil
}
