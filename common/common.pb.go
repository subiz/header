// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto

It has these top-level messages:
	Empty
	Context
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Context struct {
	EventId    *string `protobuf:"bytes,1,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	State      []byte  `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Node       *string `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
	ReplyTopic *string `protobuf:"bytes,4,opt,name=reply_topic,json=replyTopic" json:"reply_topic,omitempty"`
	// 	optional int32 reply_partition = 5;
	Credential       *auth.Credential `protobuf:"bytes,6,opt,name=credential" json:"credential,omitempty"`
	Tracing          []byte           `protobuf:"bytes,7,opt,name=tracing" json:"tracing,omitempty"`
	ReplyKey         *string          `protobuf:"bytes,8,opt,name=reply_key,json=replyKey" json:"reply_key,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Context) Reset()                    { *m = Context{} }
func (m *Context) String() string            { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()               {}
func (*Context) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Context) GetEventId() string {
	if m != nil && m.EventId != nil {
		return *m.EventId
	}
	return ""
}

func (m *Context) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Context) GetNode() string {
	if m != nil && m.Node != nil {
		return *m.Node
	}
	return ""
}

func (m *Context) GetReplyTopic() string {
	if m != nil && m.ReplyTopic != nil {
		return *m.ReplyTopic
	}
	return ""
}

func (m *Context) GetCredential() *auth.Credential {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *Context) GetTracing() []byte {
	if m != nil {
		return m.Tracing
	}
	return nil
}

func (m *Context) GetReplyKey() string {
	if m != nil && m.ReplyKey != nil {
		return *m.ReplyKey
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "common.Empty")
	proto.RegisterType((*Context)(nil), "common.Context")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xcd, 0x4e, 0xeb, 0x30,
	0x10, 0x85, 0x95, 0x7b, 0xf3, 0xd7, 0x29, 0x12, 0x95, 0x61, 0x61, 0x60, 0x41, 0xd4, 0x55, 0xc4,
	0x22, 0x41, 0xbc, 0x42, 0xc5, 0x02, 0xd8, 0x45, 0xec, 0x2b, 0x27, 0x1e, 0xb5, 0x56, 0x1b, 0xdb,
	0x72, 0x27, 0x88, 0xf0, 0xd2, 0xbc, 0x02, 0x62, 0x2c, 0x10, 0x1b, 0xdb, 0xdf, 0xe7, 0x33, 0x3a,
	0x03, 0x17, 0x83, 0x1b, 0x47, 0x67, 0xdb, 0x78, 0x35, 0x3e, 0x38, 0x72, 0x22, 0x8f, 0x74, 0x7d,
	0xd7, 0x1b, 0xea, 0xa7, 0xe1, 0x80, 0xd4, 0xb8, 0xb0, 0x6b, 0x4f, 0x53, 0x6f, 0x3e, 0xda, 0x3d,
	0x2a, 0x8d, 0xa1, 0x55, 0x13, 0xed, 0xf9, 0x88, 0x33, 0xeb, 0x02, 0xb2, 0xc7, 0xd1, 0xd3, 0xbc,
	0xfe, 0x4c, 0xa0, 0xd8, 0x38, 0x4b, 0xf8, 0x4e, 0xe2, 0x0a, 0x4a, 0x7c, 0x43, 0x4b, 0x5b, 0xa3,
	0x65, 0x52, 0x25, 0xf5, 0xa2, 0x2b, 0x98, 0x9f, 0xb4, 0xb8, 0x84, 0xec, 0x44, 0x8a, 0x50, 0xfe,
	0xab, 0x92, 0xfa, 0xac, 0x8b, 0x20, 0x04, 0xa4, 0xd6, 0x69, 0x94, 0xff, 0x39, 0xcc, 0x6f, 0x71,
	0x0b, 0xcb, 0x80, 0xfe, 0x38, 0x6f, 0xc9, 0x79, 0x33, 0xc8, 0x94, 0xbf, 0x80, 0xd5, 0xeb, 0xb7,
	0x11, 0xf7, 0x00, 0x43, 0x40, 0x8d, 0x96, 0x8c, 0x3a, 0xca, 0xbc, 0x4a, 0xea, 0xe5, 0xc3, 0xaa,
	0xe1, 0xdd, 0x36, 0xbf, 0xbe, 0xfb, 0x93, 0x11, 0x12, 0x0a, 0x0a, 0x6a, 0x30, 0x76, 0x27, 0x0b,
	0xae, 0xff, 0x41, 0x71, 0x03, 0x8b, 0x58, 0x76, 0xc0, 0x59, 0x96, 0x5c, 0x55, 0xb2, 0x78, 0xc1,
	0xf9, 0x39, 0x2d, 0xb3, 0x55, 0xde, 0x9d, 0xc7, 0x80, 0x57, 0x81, 0x0c, 0x19, 0x67, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x54, 0x7f, 0x6f, 0x22, 0x44, 0x01, 0x00, 0x00,
}
