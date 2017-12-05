// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dashboard/dashboard.proto

/*
Package dashboard is a generated protocol buffer package.

It is generated from these files:
	dashboard/dashboard.proto

It has these top-level messages:
	SessionCookie
*/
package dashboard

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

type SessionCookie struct {
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	ExpiredAt    int64  `protobuf:"varint,4,opt,name=expired_at,json=expiredAt" json:"expired_at,omitempty"`
	IssuedAt     int64  `protobuf:"varint,5,opt,name=issued_at,json=issuedAt" json:"issued_at,omitempty"`
	Type         string `protobuf:"bytes,6,opt,name=type" json:"type,omitempty"`
	Email        string `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	RememberMe   bool   `protobuf:"varint,8,opt,name=remember_me,json=rememberMe" json:"remember_me,omitempty"`
}

func (m *SessionCookie) Reset()                    { *m = SessionCookie{} }
func (m *SessionCookie) String() string            { return proto.CompactTextString(m) }
func (*SessionCookie) ProtoMessage()               {}
func (*SessionCookie) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SessionCookie) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *SessionCookie) GetExpiredAt() int64 {
	if m != nil {
		return m.ExpiredAt
	}
	return 0
}

func (m *SessionCookie) GetIssuedAt() int64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *SessionCookie) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SessionCookie) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SessionCookie) GetRememberMe() bool {
	if m != nil {
		return m.RememberMe
	}
	return false
}

func init() {
	proto.RegisterType((*SessionCookie)(nil), "dashboard.SessionCookie")
}

func init() { proto.RegisterFile("dashboard/dashboard.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xcf, 0xbf, 0x6a, 0xc3, 0x30,
	0x10, 0x06, 0x70, 0x84, 0xff, 0xd4, 0xba, 0xd6, 0xcb, 0xd1, 0x41, 0xa5, 0x94, 0x9a, 0x76, 0xf1,
	0xd4, 0x0e, 0x79, 0x02, 0x93, 0x39, 0x8b, 0x93, 0xdd, 0xc8, 0xf8, 0x82, 0x85, 0x23, 0xcb, 0x48,
	0x0a, 0x24, 0xcf, 0x96, 0x97, 0x0b, 0x91, 0x1d, 0x6f, 0xdf, 0xfd, 0x3e, 0x38, 0xf8, 0xe0, 0xa3,
	0x93, 0xae, 0x6f, 0x8d, 0xb4, 0xdd, 0xff, 0x9a, 0xfe, 0x26, 0x6b, 0xbc, 0x41, 0xbe, 0xc2, 0xcf,
	0x8d, 0x41, 0xbe, 0x27, 0xe7, 0x94, 0x19, 0xb7, 0xc6, 0x0c, 0x8a, 0xf0, 0x17, 0x72, 0x4b, 0x47,
	0x4b, 0xae, 0x6f, 0xbc, 0x19, 0x68, 0x14, 0x51, 0xc1, 0x4a, 0x5e, 0xbf, 0x2d, 0x78, 0x78, 0x18,
	0x7e, 0x01, 0xd0, 0x65, 0x52, 0x96, 0xba, 0x46, 0x7a, 0x11, 0x17, 0xac, 0x8c, 0x6a, 0xbe, 0x48,
	0xe5, 0xf1, 0x13, 0xb8, 0x72, 0xee, 0x3c, 0xb7, 0x49, 0x68, 0xb3, 0x19, 0x2a, 0x8f, 0x08, 0xb1,
	0xbf, 0x4e, 0x24, 0xd2, 0xf0, 0x37, 0x64, 0x7c, 0x87, 0x84, 0xb4, 0x54, 0x27, 0xf1, 0x12, 0x70,
	0x3e, 0xf0, 0x1b, 0x5e, 0x2d, 0x69, 0xd2, 0x2d, 0xd9, 0x46, 0x93, 0xc8, 0x0a, 0x56, 0x66, 0x35,
	0x3c, 0x69, 0x47, 0x6d, 0x1a, 0xf6, 0x6c, 0xee, 0x01, 0x00, 0x00, 0xff, 0xff, 0x89, 0xc4, 0xee,
	0x44, 0xec, 0x00, 0x00, 0x00,
}
