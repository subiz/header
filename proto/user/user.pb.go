// Code generated by protoc-gen-go.
// source: user/user.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user/user.proto

It has these top-level messages:
	User
	Users
	Device
	Trace
	Id
	Empty
	Ids
	ListRequest
	MergeRequest
	GreetingRequest
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type DeviceType int32

const (
	DeviceType_UNKNOWN DeviceType = 0
	DeviceType_DESKTOP DeviceType = 1
	DeviceType_MOBILE  DeviceType = 2
	DeviceType_TABLET  DeviceType = 3
)

var DeviceType_name = map[int32]string{
	0: "UNKNOWN",
	1: "DESKTOP",
	2: "MOBILE",
	3: "TABLET",
}
var DeviceType_value = map[string]int32{
	"UNKNOWN": 0,
	"DESKTOP": 1,
	"MOBILE":  2,
	"TABLET":  3,
}

func (x DeviceType) String() string {
	return proto.EnumName(DeviceType_name, int32(x))
}
func (DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type User struct {
	Id        string    `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
	Alias     []string  `protobuf:"bytes,2,rep,name=Alias,json=alias" json:"Alias,omitempty"`
	AccountId string    `protobuf:"bytes,3,opt,name=AccountId,json=accountId" json:"AccountId,omitempty"`
	FirstName string    `protobuf:"bytes,4,opt,name=FirstName,json=firstName" json:"FirstName,omitempty"`
	LastName  string    `protobuf:"bytes,5,opt,name=LastName,json=lastName" json:"LastName,omitempty"`
	Phones    []string  `protobuf:"bytes,6,rep,name=Phones,json=phones" json:"Phones,omitempty"`
	Emails    []string  `protobuf:"bytes,7,rep,name=Emails,json=emails" json:"Emails,omitempty"`
	Traces    []*Trace  `protobuf:"bytes,8,rep,name=Traces,json=traces" json:"Traces,omitempty"`
	Devices   []*Device `protobuf:"bytes,9,rep,name=Devices,json=devices" json:"Devices,omitempty"`
	IsBan     bool      `protobuf:"varint,10,opt,name=IsBan,json=isBan" json:"IsBan,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *User) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetPhones() []string {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *User) GetEmails() []string {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *User) GetTraces() []*Trace {
	if m != nil {
		return m.Traces
	}
	return nil
}

func (m *User) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *User) GetIsBan() bool {
	if m != nil {
		return m.IsBan
	}
	return false
}

type Users struct {
	Users []*User `protobuf:"bytes,1,rep,name=Users,json=users" json:"Users,omitempty"`
}

func (m *Users) Reset()                    { *m = Users{} }
func (m *Users) String() string            { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()               {}
func (*Users) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Device struct {
	Id              int32  `protobuf:"varint,1,opt,name=Id,json=id" json:"Id,omitempty"`
	UserAgentZipped int32  `protobuf:"varint,2,opt,name=UserAgentZipped,json=userAgentZipped" json:"UserAgentZipped,omitempty"`
	Resolution      string `protobuf:"bytes,3,opt,name=Resolution,json=resolution" json:"Resolution,omitempty"`
	LanguageZipped  int32  `protobuf:"varint,4,opt,name=LanguageZipped,json=languageZipped" json:"LanguageZipped,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Device) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Device) GetUserAgentZipped() int32 {
	if m != nil {
		return m.UserAgentZipped
	}
	return 0
}

func (m *Device) GetResolution() string {
	if m != nil {
		return m.Resolution
	}
	return ""
}

func (m *Device) GetLanguageZipped() int32 {
	if m != nil {
		return m.LanguageZipped
	}
	return 0
}

type Trace struct {
	Id             string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
	IP             string `protobuf:"bytes,2,opt,name=IP,json=iP" json:"IP,omitempty"`
	LocationZipped int32  `protobuf:"varint,3,opt,name=LocationZipped,json=locationZipped" json:"LocationZipped,omitempty"`
}

func (m *Trace) Reset()                    { *m = Trace{} }
func (m *Trace) String() string            { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()               {}
func (*Trace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Trace) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Trace) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Trace) GetLocationZipped() int32 {
	if m != nil {
		return m.LocationZipped
	}
	return 0
}

type Id struct {
	Id string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Ids struct {
	Ids []string `protobuf:"bytes,1,rep,name=Ids,json=ids" json:"Ids,omitempty"`
}

func (m *Ids) Reset()                    { *m = Ids{} }
func (m *Ids) String() string            { return proto.CompactTextString(m) }
func (*Ids) ProtoMessage()               {}
func (*Ids) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Ids) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ListRequest struct {
	AccountId string `protobuf:"bytes,1,opt,name=AccountId,json=accountId" json:"AccountId,omitempty"`
	StartId   string `protobuf:"bytes,2,opt,name=StartId,json=startId" json:"StartId,omitempty"`
	Limit     int32  `protobuf:"varint,3,opt,name=Limit,json=limit" json:"Limit,omitempty"`
	Keyword   string `protobuf:"bytes,4,opt,name=Keyword,json=keyword" json:"Keyword,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *ListRequest) GetStartId() string {
	if m != nil {
		return m.StartId
	}
	return ""
}

func (m *ListRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type MergeRequest struct {
	FormerUserId string `protobuf:"bytes,1,opt,name=FormerUserId,json=formerUserId" json:"FormerUserId,omitempty"`
	RecentUserId string `protobuf:"bytes,2,opt,name=RecentUserId,json=recentUserId" json:"RecentUserId,omitempty"`
}

func (m *MergeRequest) Reset()                    { *m = MergeRequest{} }
func (m *MergeRequest) String() string            { return proto.CompactTextString(m) }
func (*MergeRequest) ProtoMessage()               {}
func (*MergeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *MergeRequest) GetFormerUserId() string {
	if m != nil {
		return m.FormerUserId
	}
	return ""
}

func (m *MergeRequest) GetRecentUserId() string {
	if m != nil {
		return m.RecentUserId
	}
	return ""
}

type GreetingRequest struct {
	AccountId string `protobuf:"bytes,1,opt,name=AccountId,json=accountId" json:"AccountId,omitempty"`
	UserAgent string `protobuf:"bytes,2,opt,name=UserAgent,json=userAgent" json:"UserAgent,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=UserId,json=userId" json:"UserId,omitempty"`
	UUID      string `protobuf:"bytes,4,opt,name=UUID,json=uUID" json:"UUID,omitempty"`
}

func (m *GreetingRequest) Reset()                    { *m = GreetingRequest{} }
func (m *GreetingRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetingRequest) ProtoMessage()               {}
func (*GreetingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GreetingRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *GreetingRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *GreetingRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *GreetingRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Users)(nil), "user.Users")
	proto.RegisterType((*Device)(nil), "user.Device")
	proto.RegisterType((*Trace)(nil), "user.Trace")
	proto.RegisterType((*Id)(nil), "user.Id")
	proto.RegisterType((*Empty)(nil), "user.Empty")
	proto.RegisterType((*Ids)(nil), "user.Ids")
	proto.RegisterType((*ListRequest)(nil), "user.ListRequest")
	proto.RegisterType((*MergeRequest)(nil), "user.MergeRequest")
	proto.RegisterType((*GreetingRequest)(nil), "user.GreetingRequest")
	proto.RegisterEnum("user.DeviceType", DeviceType_name, DeviceType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserMgr service

type UserMgrClient interface {
	Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*Id, error)
	Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
	Ban(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	Unban(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	Merge(ctx context.Context, in *MergeRequest, opts ...grpc.CallOption) (*Empty, error)
	ReadBulk(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Users, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Users, error)
}

type userMgrClient struct {
	cc *grpc.ClientConn
}

func NewUserMgrClient(cc *grpc.ClientConn) UserMgrClient {
	return &userMgrClient{cc}
}

func (c *userMgrClient) Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := grpc.Invoke(ctx, "/user.UserMgr/Greeting", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.UserMgr/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Ban(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.UserMgr/Ban", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Unban(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.UserMgr/Unban", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) Merge(ctx context.Context, in *MergeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.UserMgr/Merge", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) ReadBulk(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := grpc.Invoke(ctx, "/user.UserMgr/ReadBulk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgrClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := grpc.Invoke(ctx, "/user.UserMgr/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserMgr service

type UserMgrServer interface {
	Greeting(context.Context, *GreetingRequest) (*Id, error)
	Update(context.Context, *User) (*Empty, error)
	Ban(context.Context, *Id) (*Empty, error)
	Unban(context.Context, *Id) (*Empty, error)
	Merge(context.Context, *MergeRequest) (*Empty, error)
	ReadBulk(context.Context, *Ids) (*Users, error)
	List(context.Context, *ListRequest) (*Users, error)
}

func RegisterUserMgrServer(s *grpc.Server, srv UserMgrServer) {
	s.RegisterService(&_UserMgr_serviceDesc, srv)
}

func _UserMgr_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/Greeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Greeting(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Update(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Ban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Ban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/Ban",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Ban(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Unban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Unban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/Unban",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Unban(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_Merge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).Merge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/Merge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).Merge(ctx, req.(*MergeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_ReadBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).ReadBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/ReadBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).ReadBulk(ctx, req.(*Ids))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgr_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgrServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserMgr/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgrServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserMgr",
	HandlerType: (*UserMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeting",
			Handler:    _UserMgr_Greeting_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserMgr_Update_Handler,
		},
		{
			MethodName: "Ban",
			Handler:    _UserMgr_Ban_Handler,
		},
		{
			MethodName: "Unban",
			Handler:    _UserMgr_Unban_Handler,
		},
		{
			MethodName: "Merge",
			Handler:    _UserMgr_Merge_Handler,
		},
		{
			MethodName: "ReadBulk",
			Handler:    _UserMgr_ReadBulk_Handler,
		},
		{
			MethodName: "List",
			Handler:    _UserMgr_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xda, 0x4a,
	0x10, 0x05, 0x8c, 0x6d, 0x3c, 0xa0, 0x84, 0xbb, 0xca, 0xbd, 0xd7, 0x8a, 0xa2, 0x08, 0xb9, 0x55,
	0x44, 0x23, 0x35, 0x95, 0xd2, 0xf7, 0x4a, 0x20, 0x48, 0x65, 0x85, 0x00, 0x72, 0xa0, 0x95, 0xfa,
	0xb6, 0xc1, 0x1b, 0x6a, 0xc5, 0xd8, 0xee, 0xee, 0xba, 0x15, 0x7f, 0xa1, 0x7f, 0xa0, 0x0f, 0xfd,
	0xb3, 0xd5, 0xec, 0x2e, 0xc8, 0xa1, 0x52, 0xd5, 0x17, 0xd8, 0x73, 0xe6, 0x70, 0x66, 0xe7, 0x63,
	0x81, 0xe3, 0x52, 0x30, 0xfe, 0x06, 0x3f, 0xae, 0x0a, 0x9e, 0xcb, 0x9c, 0x34, 0xf1, 0x1c, 0xfc,
	0x68, 0x40, 0x73, 0x29, 0x18, 0x27, 0x47, 0xd0, 0x08, 0x63, 0xbf, 0xde, 0xab, 0xf7, 0xbd, 0xa8,
	0x91, 0xc4, 0xe4, 0x04, 0xec, 0x41, 0x9a, 0x50, 0xe1, 0x37, 0x7a, 0x56, 0xdf, 0x8b, 0x6c, 0x8a,
	0x80, 0x9c, 0x81, 0x37, 0x58, 0xad, 0xf2, 0x32, 0x93, 0x61, 0xec, 0x5b, 0x4a, 0xec, 0xd1, 0x1d,
	0x81, 0xd1, 0x9b, 0x84, 0x0b, 0x39, 0xa5, 0x1b, 0xe6, 0x37, 0x75, 0xf4, 0x71, 0x47, 0x90, 0x53,
	0x68, 0x4d, 0xa8, 0x09, 0xda, 0x2a, 0xd8, 0x4a, 0x0d, 0x26, 0xff, 0x81, 0x33, 0xff, 0x9c, 0x67,
	0x4c, 0xf8, 0x8e, 0x4a, 0xe7, 0x14, 0x0a, 0x21, 0x3f, 0xde, 0xd0, 0x24, 0x15, 0xbe, 0xab, 0x79,
	0xa6, 0x10, 0x79, 0x01, 0xce, 0x82, 0xd3, 0x15, 0x13, 0x7e, 0xab, 0x67, 0xf5, 0xdb, 0xd7, 0xed,
	0x2b, 0x55, 0x99, 0xe2, 0x22, 0x47, 0xaa, 0x10, 0xb9, 0x00, 0x77, 0xc4, 0xbe, 0x26, 0xa8, 0xf2,
	0x94, 0xaa, 0xa3, 0x55, 0x9a, 0x8c, 0xdc, 0x58, 0x07, 0xb1, 0xd4, 0x50, 0x0c, 0x69, 0xe6, 0x43,
	0xaf, 0xde, 0x6f, 0x45, 0x76, 0x82, 0x20, 0x78, 0x05, 0x36, 0x36, 0x46, 0x90, 0x9e, 0x39, 0xf8,
	0x75, 0x65, 0x02, 0xda, 0x04, 0xa9, 0xc8, 0xc6, 0xa3, 0x08, 0xbe, 0xd7, 0xc1, 0xd1, 0xa6, 0x95,
	0x36, 0xda, 0xaa, 0x8d, 0x7d, 0x38, 0x46, 0xe5, 0x60, 0xcd, 0x32, 0xf9, 0x29, 0x29, 0x0a, 0x16,
	0xfb, 0x0d, 0x15, 0x54, 0xf3, 0xa8, 0xd0, 0xe4, 0x1c, 0x20, 0x62, 0x22, 0x4f, 0x4b, 0x99, 0xe4,
	0x99, 0xe9, 0x2d, 0xf0, 0x3d, 0x43, 0x2e, 0xe0, 0x68, 0x42, 0xb3, 0x75, 0x49, 0xd7, 0xcc, 0x18,
	0x35, 0x95, 0xd1, 0x51, 0xfa, 0x8c, 0x0d, 0x66, 0x60, 0xab, 0x36, 0xfc, 0x36, 0x51, 0xc4, 0x73,
	0x95, 0x1d, 0xf1, 0x5c, 0x19, 0xe6, 0x2b, 0x8a, 0xe6, 0xc6, 0xd0, 0x32, 0x86, 0xcf, 0xd8, 0xe0,
	0x04, 0x7d, 0x0e, 0xdd, 0x02, 0x17, 0xec, 0xf1, 0xa6, 0x90, 0xdb, 0xe0, 0x7f, 0xb0, 0xc2, 0x58,
	0x90, 0xae, 0xfa, 0x52, 0x3d, 0xf2, 0x22, 0x2b, 0x89, 0x45, 0x50, 0x42, 0x7b, 0x92, 0x08, 0x19,
	0xb1, 0x2f, 0x25, 0x13, 0xf2, 0xf9, 0xea, 0xd4, 0x0f, 0x57, 0xc7, 0x07, 0xf7, 0x5e, 0x52, 0x8e,
	0x31, 0x7d, 0x43, 0x57, 0x68, 0x88, 0xd3, 0x99, 0x24, 0x9b, 0x44, 0x9a, 0xdb, 0xd9, 0x29, 0x02,
	0xd4, 0xdf, 0xb2, 0xed, 0xb7, 0x9c, 0xc7, 0x66, 0xd1, 0xdc, 0x27, 0x0d, 0x83, 0x0f, 0xd0, 0xb9,
	0x63, 0x7c, 0xcd, 0x76, 0x79, 0x03, 0xe8, 0xdc, 0xe4, 0x7c, 0xc3, 0x38, 0xce, 0x61, 0x9f, 0xba,
	0xf3, 0x58, 0xe1, 0x50, 0x13, 0xb1, 0x15, 0xcb, 0xa4, 0xd1, 0xe8, 0x2b, 0x74, 0x78, 0x85, 0x0b,
	0xb6, 0x70, 0xfc, 0x9e, 0x33, 0x26, 0x93, 0x6c, 0xfd, 0x77, 0x25, 0x9d, 0x81, 0xb7, 0x1f, 0xbd,
	0x71, 0xf4, 0xf6, 0x43, 0xc7, 0xcd, 0x36, 0xc9, 0xf4, 0xa8, 0x9d, 0x52, 0x5f, 0x85, 0x40, 0x73,
	0xb9, 0x0c, 0x47, 0xa6, 0xaa, 0x66, 0xb9, 0x0c, 0x47, 0x97, 0xef, 0x00, 0xf4, 0x7a, 0x2d, 0xb6,
	0x05, 0x23, 0x6d, 0x70, 0x97, 0xd3, 0xdb, 0xe9, 0xec, 0xe3, 0xb4, 0x5b, 0x43, 0x30, 0x1a, 0xdf,
	0xdf, 0x2e, 0x66, 0xf3, 0x6e, 0x9d, 0x00, 0x38, 0x77, 0xb3, 0x61, 0x38, 0x19, 0x77, 0x1b, 0x78,
	0x5e, 0x0c, 0x86, 0x93, 0xf1, 0xa2, 0x6b, 0x5d, 0xff, 0x6c, 0x80, 0x8b, 0xc9, 0xee, 0xd6, 0x9c,
	0xbc, 0x86, 0xd6, 0xae, 0x0c, 0xf2, 0xaf, 0x5e, 0xe5, 0x83, 0xb2, 0x4e, 0x5b, 0x9a, 0x0e, 0xe3,
	0xa0, 0x86, 0x0f, 0x6d, 0x59, 0xc4, 0x54, 0x32, 0x52, 0xd9, 0xfb, 0x53, 0xf3, 0xdc, 0xf4, 0x02,
	0xd4, 0xc8, 0x39, 0x58, 0x43, 0x9a, 0x91, 0xfd, 0xef, 0x0e, 0xe3, 0xf8, 0x82, 0xb2, 0x87, 0x3f,
	0x29, 0x2e, 0xc1, 0x56, 0x43, 0x23, 0x44, 0xf3, 0xd5, 0x09, 0x1e, 0x6a, 0x5f, 0x42, 0x2b, 0x62,
	0x34, 0x1e, 0x96, 0xe9, 0x13, 0xf1, 0x76, 0x86, 0x62, 0xa7, 0x52, 0x4f, 0x35, 0xa8, 0x91, 0x3e,
	0x34, 0x71, 0xfb, 0xc8, 0x3f, 0x9a, 0xae, 0x6c, 0xe2, 0x81, 0xf2, 0xc1, 0x51, 0xff, 0x87, 0x6f,
	0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xca, 0xe4, 0x62, 0x1f, 0x22, 0x05, 0x00, 0x00,
}
