// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dashboard/dashboard.proto

/*
Package dashboard is a generated protocol buffer package.

It is generated from these files:
	dashboard/dashboard.proto

It has these top-level messages:
	SessionCookie
	Global
*/
package dashboard

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "git.subiz.net/header/common"

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
	Event_DashboardVersionUpdated          Event = 0
	Event_DashboardLanguageUpdated         Event = 5
	Event_DashboardLanguageUpdateRequested Event = 7
)

var Event_name = map[int32]string{
	0: "DashboardVersionUpdated",
	5: "DashboardLanguageUpdated",
	7: "DashboardLanguageUpdateRequested",
}
var Event_value = map[string]int32{
	"DashboardVersionUpdated":          0,
	"DashboardLanguageUpdated":         5,
	"DashboardLanguageUpdateRequested": 7,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

// global setting. eg: dashboard_4/4234098234
type Global struct {
	Ctx  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Name string          `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Data string          `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	Pk   string          `protobuf:"bytes,4,opt,name=pk" json:"pk,omitempty"`
}

func (m *Global) Reset()                    { *m = Global{} }
func (m *Global) String() string            { return proto.CompactTextString(m) }
func (*Global) ProtoMessage()               {}
func (*Global) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Global) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Global) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Global) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Global) GetPk() string {
	if m != nil {
		return m.Pk
	}
	return ""
}

func init() {
	proto.RegisterType((*SessionCookie)(nil), "dashboard.SessionCookie")
	proto.RegisterType((*Global)(nil), "dashboard.Global")
	proto.RegisterEnum("dashboard.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("dashboard/dashboard.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcb, 0x4e, 0xe3, 0x30,
	0x14, 0x86, 0x27, 0x6d, 0xd3, 0x36, 0xa7, 0xd3, 0x19, 0x64, 0x21, 0x61, 0x6e, 0x22, 0x14, 0x16,
	0x11, 0x8b, 0x54, 0x82, 0x27, 0xa8, 0x0a, 0x62, 0x03, 0x9b, 0x70, 0xd9, 0x56, 0x4e, 0x73, 0x68,
	0xa2, 0xd4, 0x76, 0xb0, 0x1d, 0x54, 0x78, 0x35, 0x5e, 0x0e, 0xd9, 0x49, 0xbb, 0x63, 0x95, 0x3f,
	0xdf, 0x97, 0xf8, 0xe8, 0x3f, 0x86, 0xc3, 0x8c, 0xe9, 0x3c, 0x95, 0x4c, 0x65, 0xd3, 0x5d, 0x8a,
	0x2b, 0x25, 0x8d, 0x24, 0xc1, 0x0e, 0x1c, 0x45, 0xab, 0xc2, 0xc4, 0xba, 0x4e, 0x8b, 0xaf, 0x58,
	0xa0, 0x99, 0xe6, 0xc8, 0x32, 0x54, 0xd3, 0xa5, 0xe4, 0x5c, 0x8a, 0xf6, 0xd1, 0xfc, 0x34, 0xf9,
	0xf6, 0x60, 0xfc, 0x84, 0x5a, 0x17, 0x52, 0xcc, 0xa5, 0x2c, 0x0b, 0x24, 0x17, 0x30, 0x56, 0xf8,
	0xa6, 0x50, 0xe7, 0x0b, 0x23, 0x4b, 0x14, 0xb4, 0x1b, 0x7a, 0x51, 0x90, 0xfc, 0x6d, 0xe1, 0xb3,
	0x65, 0xe4, 0x14, 0x00, 0x37, 0x55, 0xa1, 0x30, 0x5b, 0x30, 0x43, 0x7b, 0xa1, 0x17, 0x75, 0x93,
	0xa0, 0x25, 0x33, 0x43, 0x8e, 0x21, 0x28, 0xb4, 0xae, 0x1b, 0xeb, 0x3b, 0x3b, 0x6c, 0xc0, 0xcc,
	0x10, 0x02, 0x3d, 0xf3, 0x59, 0x21, 0xed, 0xbb, 0x73, 0x5d, 0x26, 0xfb, 0xe0, 0x23, 0x67, 0xc5,
	0x9a, 0x0e, 0x1c, 0x6c, 0x5e, 0xc8, 0x19, 0x8c, 0x14, 0x72, 0xe4, 0x29, 0xaa, 0x05, 0x47, 0x3a,
	0x0c, 0xbd, 0x68, 0x98, 0xc0, 0x16, 0x3d, 0xe2, 0x64, 0x09, 0xfd, 0xfb, 0xb5, 0x4c, 0xd9, 0x9a,
	0x9c, 0x43, 0x77, 0x69, 0x36, 0xd4, 0x0b, 0xbd, 0x68, 0x74, 0xfd, 0x3f, 0x6e, 0x3b, 0xce, 0xa5,
	0x30, 0xb8, 0x31, 0x89, 0x75, 0x76, 0xae, 0x60, 0x1c, 0x69, 0xa7, 0x99, 0x6b, 0xb3, 0x65, 0x19,
	0x33, 0xac, 0xed, 0xe8, 0x32, 0xf9, 0x07, 0x9d, 0xaa, 0x74, 0x9d, 0x82, 0xa4, 0x53, 0x95, 0x57,
	0x39, 0xf8, 0x77, 0x1f, 0x28, 0x6c, 0xab, 0x83, 0xdb, 0xed, 0x8a, 0x5f, 0x51, 0xd9, 0x9d, 0xbd,
	0x54, 0x19, 0x33, 0x98, 0xed, 0xfd, 0x21, 0x27, 0x40, 0x77, 0xf2, 0x81, 0x89, 0x55, 0xcd, 0x56,
	0xb8, 0xb5, 0x3e, 0xb9, 0x84, 0xf0, 0x17, 0x9b, 0xe0, 0x7b, 0x8d, 0xda, 0x7e, 0x35, 0x48, 0xfb,
	0xee, 0x4e, 0x6e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x2e, 0xf1, 0xa5, 0xe5, 0x01, 0x00,
	0x00,
}
