// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.2
// source: mailkon.proto

package mailkon

import (
	header "github.com/subiz/header"
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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId *string `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Address   *string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Updated   *int64  `protobuf:"varint,3,opt,name=updated" json:"updated,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailkon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_mailkon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_mailkon_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *Address) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *Address) GetUpdated() int64 {
	if x != nil && x.Updated != nil {
		return *x.Updated
	}
	return 0
}

type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId *string `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Domain    *string `protobuf:"bytes,2,opt,name=domain" json:"domain,omitempty"`
	Updated   *int64  `protobuf:"varint,3,opt,name=updated" json:"updated,omitempty"`
	Dnstype   *string `protobuf:"bytes,4,opt,name=dnstype" json:"dnstype,omitempty"`
	Data      *string `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailkon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_mailkon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_mailkon_proto_rawDescGZIP(), []int{1}
}

func (x *Domain) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *Domain) GetDomain() string {
	if x != nil && x.Domain != nil {
		return *x.Domain
	}
	return ""
}

func (x *Domain) GetUpdated() int64 {
	if x != nil && x.Updated != nil {
		return *x.Updated
	}
	return 0
}

func (x *Domain) GetDnstype() string {
	if x != nil && x.Dnstype != nil {
		return *x.Dnstype
	}
	return ""
}

func (x *Domain) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId      *string  `protobuf:"bytes,2,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	AccountId      *string  `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ConversationId *string  `protobuf:"bytes,4,opt,name=conversation_id,json=conversationId" json:"conversation_id,omitempty"`
	FromAddr       *string  `protobuf:"bytes,5,opt,name=from_addr,json=fromAddr" json:"from_addr,omitempty"`
	ToAddr         *string  `protobuf:"bytes,6,opt,name=to_addr,json=toAddr" json:"to_addr,omitempty"`
	Subject        *string  `protobuf:"bytes,7,opt,name=subject" json:"subject,omitempty"`
	Body           *string  `protobuf:"bytes,8,opt,name=body" json:"body,omitempty"`
	Header         *string  `protobuf:"bytes,9,opt,name=header" json:"header,omitempty"`
	Attachments    []string `protobuf:"bytes,10,rep,name=attachments" json:"attachments,omitempty"`
	Created        *int64   `protobuf:"varint,11,opt,name=created" json:"created,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailkon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_mailkon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_mailkon_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetMessageId() string {
	if x != nil && x.MessageId != nil {
		return *x.MessageId
	}
	return ""
}

func (x *Message) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *Message) GetConversationId() string {
	if x != nil && x.ConversationId != nil {
		return *x.ConversationId
	}
	return ""
}

func (x *Message) GetFromAddr() string {
	if x != nil && x.FromAddr != nil {
		return *x.FromAddr
	}
	return ""
}

func (x *Message) GetToAddr() string {
	if x != nil && x.ToAddr != nil {
		return *x.ToAddr
	}
	return ""
}

func (x *Message) GetSubject() string {
	if x != nil && x.Subject != nil {
		return *x.Subject
	}
	return ""
}

func (x *Message) GetBody() string {
	if x != nil && x.Body != nil {
		return *x.Body
	}
	return ""
}

func (x *Message) GetHeader() string {
	if x != nil && x.Header != nil {
		return *x.Header
	}
	return ""
}

func (x *Message) GetAttachments() []string {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *Message) GetCreated() int64 {
	if x != nil && x.Created != nil {
		return *x.Created
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId    *string `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId       *string `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	EmailAddress *string `protobuf:"bytes,4,opt,name=email_address,json=emailAddress" json:"email_address,omitempty"`
	Seen         *int64  `protobuf:"varint,5,opt,name=seen" json:"seen,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailkon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_mailkon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_mailkon_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *User) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *User) GetEmailAddress() string {
	if x != nil && x.EmailAddress != nil {
		return *x.EmailAddress
	}
	return ""
}

func (x *User) GetSeen() int64 {
	if x != nil && x.Seen != nil {
		return *x.Seen
	}
	return 0
}

type SendgridEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId        *string       `protobuf:"bytes,3,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	AccountId      *string       `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ConversationId *string       `protobuf:"bytes,6,opt,name=conversation_id,json=conversationId" json:"conversation_id,omitempty"`
	MessageId      *string       `protobuf:"bytes,7,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	Data           []string      `protobuf:"bytes,8,rep,name=data" json:"data,omitempty"`
	FullMessageId  *string       `protobuf:"bytes,9,opt,name=full_message_id,json=fullMessageId" json:"full_message_id,omitempty"`
	Subject        *string       `protobuf:"bytes,10,opt,name=subject" json:"subject,omitempty"`
	Event          *header.Event `protobuf:"bytes,11,opt,name=event" json:"event,omitempty"`
	RootMessageId  *string       `protobuf:"bytes,12,opt,name=root_message_id,json=rootMessageId" json:"root_message_id,omitempty"`
}

func (x *SendgridEvent) Reset() {
	*x = SendgridEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailkon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendgridEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendgridEvent) ProtoMessage() {}

func (x *SendgridEvent) ProtoReflect() protoreflect.Message {
	mi := &file_mailkon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendgridEvent.ProtoReflect.Descriptor instead.
func (*SendgridEvent) Descriptor() ([]byte, []int) {
	return file_mailkon_proto_rawDescGZIP(), []int{4}
}

func (x *SendgridEvent) GetEventId() string {
	if x != nil && x.EventId != nil {
		return *x.EventId
	}
	return ""
}

func (x *SendgridEvent) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *SendgridEvent) GetConversationId() string {
	if x != nil && x.ConversationId != nil {
		return *x.ConversationId
	}
	return ""
}

func (x *SendgridEvent) GetMessageId() string {
	if x != nil && x.MessageId != nil {
		return *x.MessageId
	}
	return ""
}

func (x *SendgridEvent) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SendgridEvent) GetFullMessageId() string {
	if x != nil && x.FullMessageId != nil {
		return *x.FullMessageId
	}
	return ""
}

func (x *SendgridEvent) GetSubject() string {
	if x != nil && x.Subject != nil {
		return *x.Subject
	}
	return ""
}

func (x *SendgridEvent) GetEvent() *header.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *SendgridEvent) GetRootMessageId() string {
	if x != nil && x.RootMessageId != nil {
		return *x.RootMessageId
	}
	return ""
}

type SendgridTrackingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       *string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Timestamp   *int64  `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	SmtpId      *string `protobuf:"bytes,5,opt,name=smtp_id,json=smtpId" json:"smtp_id,omitempty"`
	Event       *string `protobuf:"bytes,6,opt,name=event" json:"event,omitempty"`
	Category    *string `protobuf:"bytes,7,opt,name=category" json:"category,omitempty"`
	SgEventId   *string `protobuf:"bytes,8,opt,name=sg_event_id,json=sgEventId" json:"sg_event_id,omitempty"`
	SgMessageId *string `protobuf:"bytes,9,opt,name=sg_message_id,json=sgMessageId" json:"sg_message_id,omitempty"`
	Reason      *string `protobuf:"bytes,10,opt,name=reason" json:"reason,omitempty"`
	Status      *string `protobuf:"bytes,11,opt,name=status" json:"status,omitempty"`
	Response    *string `protobuf:"bytes,12,opt,name=response" json:"response,omitempty"`
	Attempt     *string `protobuf:"bytes,13,opt,name=attempt" json:"attempt,omitempty"`
	Type        *string `protobuf:"bytes,14,opt,name=type" json:"type,omitempty"`
	Ip          *string `protobuf:"bytes,16,opt,name=ip" json:"ip,omitempty"`
	Useragent   *string `protobuf:"bytes,17,opt,name=useragent" json:"useragent,omitempty"`
	Url         *string `protobuf:"bytes,18,opt,name=url" json:"url,omitempty"`
	AsmGroupId  *int64  `protobuf:"varint,19,opt,name=asm_group_id,json=asmGroupId" json:"asm_group_id,omitempty"`
}

func (x *SendgridTrackingEvent) Reset() {
	*x = SendgridTrackingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailkon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendgridTrackingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendgridTrackingEvent) ProtoMessage() {}

func (x *SendgridTrackingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_mailkon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendgridTrackingEvent.ProtoReflect.Descriptor instead.
func (*SendgridTrackingEvent) Descriptor() ([]byte, []int) {
	return file_mailkon_proto_rawDescGZIP(), []int{5}
}

func (x *SendgridTrackingEvent) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *SendgridTrackingEvent) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *SendgridTrackingEvent) GetSmtpId() string {
	if x != nil && x.SmtpId != nil {
		return *x.SmtpId
	}
	return ""
}

func (x *SendgridTrackingEvent) GetEvent() string {
	if x != nil && x.Event != nil {
		return *x.Event
	}
	return ""
}

func (x *SendgridTrackingEvent) GetCategory() string {
	if x != nil && x.Category != nil {
		return *x.Category
	}
	return ""
}

func (x *SendgridTrackingEvent) GetSgEventId() string {
	if x != nil && x.SgEventId != nil {
		return *x.SgEventId
	}
	return ""
}

func (x *SendgridTrackingEvent) GetSgMessageId() string {
	if x != nil && x.SgMessageId != nil {
		return *x.SgMessageId
	}
	return ""
}

func (x *SendgridTrackingEvent) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *SendgridTrackingEvent) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *SendgridTrackingEvent) GetResponse() string {
	if x != nil && x.Response != nil {
		return *x.Response
	}
	return ""
}

func (x *SendgridTrackingEvent) GetAttempt() string {
	if x != nil && x.Attempt != nil {
		return *x.Attempt
	}
	return ""
}

func (x *SendgridTrackingEvent) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *SendgridTrackingEvent) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

func (x *SendgridTrackingEvent) GetUseragent() string {
	if x != nil && x.Useragent != nil {
		return *x.Useragent
	}
	return ""
}

func (x *SendgridTrackingEvent) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *SendgridTrackingEvent) GetAsmGroupId() int64 {
	if x != nil && x.AsmGroupId != nil {
		return *x.AsmGroupId
	}
	return 0
}

var File_mailkon_proto protoreflect.FileDescriptor

var file_mailkon_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x61, 0x69, 0x6c, 0x6b, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x61, 0x69, 0x6c, 0x6b, 0x6f, 0x6e, 0x1a, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x6e, 0x73, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x6e, 0x73, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa8,
	0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x77, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x65, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x65,
	0x65, 0x6e, 0x22, 0xb4, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x67, 0x72, 0x69, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x75, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x6f, 0x6f, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0xb6, 0x03, 0x0a, 0x15, 0x53, 0x65,
	0x6e, 0x64, 0x67, 0x72, 0x69, 0x64, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6d, 0x74, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6d, 0x74, 0x70, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x20, 0x0a, 0x0c, 0x61, 0x73, 0x6d, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x73, 0x6d, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x75, 0x62, 0x69, 0x7a, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x61,
	0x69, 0x6c, 0x6b, 0x6f, 0x6e,
}

var (
	file_mailkon_proto_rawDescOnce sync.Once
	file_mailkon_proto_rawDescData = file_mailkon_proto_rawDesc
)

func file_mailkon_proto_rawDescGZIP() []byte {
	file_mailkon_proto_rawDescOnce.Do(func() {
		file_mailkon_proto_rawDescData = protoimpl.X.CompressGZIP(file_mailkon_proto_rawDescData)
	})
	return file_mailkon_proto_rawDescData
}

var file_mailkon_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mailkon_proto_goTypes = []interface{}{
	(*Address)(nil),               // 0: mailkon.Address
	(*Domain)(nil),                // 1: mailkon.Domain
	(*Message)(nil),               // 2: mailkon.Message
	(*User)(nil),                  // 3: mailkon.User
	(*SendgridEvent)(nil),         // 4: mailkon.SendgridEvent
	(*SendgridTrackingEvent)(nil), // 5: mailkon.SendgridTrackingEvent
	(*header.Event)(nil),          // 6: header.Event
}
var file_mailkon_proto_depIdxs = []int32{
	6, // 0: mailkon.SendgridEvent.event:type_name -> header.Event
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mailkon_proto_init() }
func file_mailkon_proto_init() {
	if File_mailkon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mailkon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_mailkon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain); i {
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
		file_mailkon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_mailkon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_mailkon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendgridEvent); i {
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
		file_mailkon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendgridTrackingEvent); i {
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
			RawDescriptor: file_mailkon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mailkon_proto_goTypes,
		DependencyIndexes: file_mailkon_proto_depIdxs,
		MessageInfos:      file_mailkon_proto_msgTypes,
	}.Build()
	File_mailkon_proto = out.File
	file_mailkon_proto_rawDesc = nil
	file_mailkon_proto_goTypes = nil
	file_mailkon_proto_depIdxs = nil
}
