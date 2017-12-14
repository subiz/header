// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mailkon/mailkon.proto

/*
Package mailkon is a generated protocol buffer package.

It is generated from these files:
	mailkon/mailkon.proto

It has these top-level messages:
	Address
	Message
*/
package mailkon

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Address struct {
	AccountId        *string `protobuf:"bytes,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Address          *string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Updated          *int64  `protobuf:"varint,3,opt,name=updated" json:"updated,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Address) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Address) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *Address) GetUpdated() int64 {
	if m != nil && m.Updated != nil {
		return *m.Updated
	}
	return 0
}

type Message struct {
	MessageId        *string  `protobuf:"bytes,2,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	AccountId        *string  `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ConversationId   *string  `protobuf:"bytes,4,opt,name=conversation_id,json=conversationId" json:"conversation_id,omitempty"`
	FromAddr         *string  `protobuf:"bytes,5,opt,name=from_addr,json=fromAddr" json:"from_addr,omitempty"`
	ToAddr           *string  `protobuf:"bytes,6,opt,name=to_addr,json=toAddr" json:"to_addr,omitempty"`
	Subject          *string  `protobuf:"bytes,7,opt,name=subject" json:"subject,omitempty"`
	Body             *string  `protobuf:"bytes,8,opt,name=body" json:"body,omitempty"`
	Header           *string  `protobuf:"bytes,9,opt,name=header" json:"header,omitempty"`
	Attachments      []string `protobuf:"bytes,10,rep,name=attachments" json:"attachments,omitempty"`
	Created          *int64   `protobuf:"varint,11,opt,name=created" json:"created,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Message) GetMessageId() string {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return ""
}

func (m *Message) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Message) GetConversationId() string {
	if m != nil && m.ConversationId != nil {
		return *m.ConversationId
	}
	return ""
}

func (m *Message) GetFromAddr() string {
	if m != nil && m.FromAddr != nil {
		return *m.FromAddr
	}
	return ""
}

func (m *Message) GetToAddr() string {
	if m != nil && m.ToAddr != nil {
		return *m.ToAddr
	}
	return ""
}

func (m *Message) GetSubject() string {
	if m != nil && m.Subject != nil {
		return *m.Subject
	}
	return ""
}

func (m *Message) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *Message) GetHeader() string {
	if m != nil && m.Header != nil {
		return *m.Header
	}
	return ""
}

func (m *Message) GetAttachments() []string {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Message) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func init() {
	proto.RegisterType((*Address)(nil), "mailkon.Address")
	proto.RegisterType((*Message)(nil), "mailkon.Message")
}

func init() { proto.RegisterFile("mailkon/mailkon.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0x66, 0xbb, 0x75, 0xd3, 0x4c, 0x41, 0x21, 0xa0, 0x06, 0x44, 0x58, 0x7a, 0x71, 0x4f, 0xfa,
	0x0c, 0x1e, 0xf7, 0xe0, 0x65, 0xcf, 0x42, 0x49, 0x93, 0x68, 0x57, 0xdd, 0x4c, 0x49, 0x66, 0x05,
	0xdf, 0xc8, 0xc7, 0x2c, 0xf9, 0x29, 0x94, 0x9e, 0x32, 0xdf, 0x0f, 0xdf, 0x97, 0x19, 0xb8, 0x9d,
	0xd4, 0xf8, 0xf3, 0x8d, 0xee, 0xa5, 0xbc, 0xcf, 0x07, 0x8f, 0x84, 0x82, 0x15, 0xb8, 0x79, 0x07,
	0xf6, 0x6a, 0x8c, 0xb7, 0x21, 0x88, 0x47, 0x00, 0xa5, 0x35, 0xce, 0x8e, 0xb6, 0xa3, 0x91, 0x55,
	0x5b, 0x75, 0x7c, 0xe0, 0x85, 0xe9, 0x8d, 0x90, 0xc0, 0x54, 0x76, 0xca, 0x45, 0xd2, 0x4e, 0x30,
	0x2a, 0xf3, 0xc1, 0x28, 0xb2, 0x46, 0xd6, 0x6d, 0xd5, 0xd5, 0xc3, 0x09, 0x6e, 0xfe, 0x17, 0xc0,
	0xde, 0x6c, 0x08, 0xea, 0xd3, 0xc6, 0xf8, 0x29, 0x8f, 0x31, 0x3e, 0x47, 0xf0, 0xc2, 0xf4, 0xe6,
	0xa2, 0xbd, 0xbe, 0x6c, 0x7f, 0x82, 0x1b, 0x8d, 0xee, 0xd7, 0xfa, 0xa0, 0x68, 0x44, 0x17, 0x3d,
	0xcb, 0xe4, 0xb9, 0x3e, 0xa7, 0x7b, 0x23, 0x1e, 0x80, 0x7f, 0x78, 0x9c, 0xb6, 0xf1, 0x73, 0xf2,
	0x2a, 0x59, 0x56, 0x91, 0x88, 0x5b, 0x8a, 0x7b, 0x60, 0x84, 0x59, 0x6a, 0x92, 0xd4, 0x10, 0x26,
	0x41, 0x02, 0x0b, 0xf3, 0xee, 0xcb, 0x6a, 0x92, 0x2c, 0x2f, 0x57, 0xa0, 0x10, 0xb0, 0xdc, 0xa1,
	0xf9, 0x93, 0xab, 0x44, 0xa7, 0x59, 0xdc, 0x41, 0xb3, 0xb7, 0xca, 0x58, 0x2f, 0x79, 0x4e, 0xc9,
	0x48, 0xb4, 0xb0, 0x56, 0x44, 0x4a, 0xef, 0x27, 0xeb, 0x28, 0x48, 0x68, 0xeb, 0x8e, 0x0f, 0xe7,
	0x54, 0xec, 0xd1, 0xde, 0xa6, 0x53, 0xad, 0xf3, 0xa9, 0x0a, 0x3c, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xda, 0x32, 0x12, 0xb4, 0xa9, 0x01, 0x00, 0x00,
}
