// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email/email.proto

/*
Package email is a generated protocol buffer package.

It is generated from these files:
	email/email.proto

It has these top-level messages:
	Email
*/
package email

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

type Email struct {
	From        string   `protobuf:"bytes,1,opt,name=From" json:"From,omitempty"`
	To          string   `protobuf:"bytes,2,opt,name=To" json:"To,omitempty"`
	Subject     string   `protobuf:"bytes,3,opt,name=Subject" json:"Subject,omitempty"`
	Body        string   `protobuf:"bytes,4,opt,name=Body" json:"Body,omitempty"`
	JSONData    string   `protobuf:"bytes,5,opt,name=JSONData" json:"JSONData,omitempty"`
	AttachLinks []string `protobuf:"bytes,6,rep,name=AttachLinks" json:"AttachLinks,omitempty"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Email) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Email) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Email) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Email) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Email) GetJSONData() string {
	if m != nil {
		return m.JSONData
	}
	return ""
}

func (m *Email) GetAttachLinks() []string {
	if m != nil {
		return m.AttachLinks
	}
	return nil
}

func init() {
	proto.RegisterType((*Email)(nil), "email.Email")
}

func init() { proto.RegisterFile("email/email.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xcd, 0x4d, 0xcc,
	0xcc, 0xd1, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0xd2, 0x74,
	0x46, 0x2e, 0x56, 0x57, 0x10, 0x4b, 0x48, 0x88, 0x8b, 0xc5, 0xad, 0x28, 0x3f, 0x57, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0xe2, 0xe3, 0x62, 0x0a, 0xc9, 0x97, 0x60, 0x02, 0x8b,
	0x30, 0x85, 0xe4, 0x0b, 0x49, 0x70, 0xb1, 0x07, 0x97, 0x26, 0x65, 0xa5, 0x26, 0x97, 0x48, 0x30,
	0x83, 0x05, 0x61, 0x5c, 0x90, 0x6e, 0xa7, 0xfc, 0x94, 0x4a, 0x09, 0x16, 0x88, 0x6e, 0x10, 0x5b,
	0x48, 0x8a, 0x8b, 0xc3, 0x2b, 0xd8, 0xdf, 0xcf, 0x25, 0xb1, 0x24, 0x51, 0x82, 0x15, 0x2c, 0x0e,
	0xe7, 0x0b, 0x29, 0x70, 0x71, 0x3b, 0x96, 0x94, 0x24, 0x26, 0x67, 0xf8, 0x64, 0xe6, 0x65, 0x17,
	0x4b, 0xb0, 0x29, 0x30, 0x6b, 0x70, 0x06, 0x21, 0x0b, 0x25, 0xb1, 0x81, 0xdd, 0x69, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x64, 0xf9, 0xe8, 0x7b, 0xbc, 0x00, 0x00, 0x00,
}
