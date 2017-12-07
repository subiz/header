// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file/file.proto

/*
Package file is a generated protocol buffer package.

It is generated from these files:
	file/file.proto

It has these top-level messages:
	AllType
	FileHeader
	PresignResult
	File
*/
package file

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bitbucket.org/subiz/header/common"

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

type Event int32

const (
	Event_FilePresignRequested Event = 0
	Event_FileCreated          Event = 3
	Event_FileInfoRequested    Event = 4
	Event_FileUploaded         Event = 6
)

var Event_name = map[int32]string{
	0: "FilePresignRequested",
	3: "FileCreated",
	4: "FileInfoRequested",
	6: "FileUploaded",
}
var Event_value = map[string]int32{
	"FilePresignRequested": 0,
	"FileCreated":          3,
	"FileInfoRequested":    4,
	"FileUploaded":         6,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllType struct {
	Fh *FileHeader    `protobuf:"bytes,2,opt,name=fh" json:"fh,omitempty"`
	Ps *PresignResult `protobuf:"bytes,3,opt,name=ps" json:"ps,omitempty"`
	F  *File          `protobuf:"bytes,4,opt,name=f" json:"f,omitempty"`
}

func (m *AllType) Reset()                    { *m = AllType{} }
func (m *AllType) String() string            { return proto.CompactTextString(m) }
func (*AllType) ProtoMessage()               {}
func (*AllType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AllType) GetFh() *FileHeader {
	if m != nil {
		return m.Fh
	}
	return nil
}

func (m *AllType) GetPs() *PresignResult {
	if m != nil {
		return m.Ps
	}
	return nil
}

func (m *AllType) GetF() *File {
	if m != nil {
		return m.F
	}
	return nil
}

type FileHeader struct {
	Ctx         *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId   string          `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Name        string          `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Type        string          `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Size        int64           `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	Md5         string          `protobuf:"bytes,6,opt,name=md5" json:"md5,omitempty"`
	Description string          `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
}

func (m *FileHeader) Reset()                    { *m = FileHeader{} }
func (m *FileHeader) String() string            { return proto.CompactTextString(m) }
func (*FileHeader) ProtoMessage()               {}
func (*FileHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FileHeader) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *FileHeader) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *FileHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileHeader) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FileHeader) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileHeader) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *FileHeader) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type PresignResult struct {
	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId string          `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Url       string          `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
	Id        string          `protobuf:"bytes,6,opt,name=id" json:"id,omitempty"`
	SignedUrl string          `protobuf:"bytes,5,opt,name=signed_url,json=signedUrl" json:"signed_url,omitempty"`
}

func (m *PresignResult) Reset()                    { *m = PresignResult{} }
func (m *PresignResult) String() string            { return proto.CompactTextString(m) }
func (*PresignResult) ProtoMessage()               {}
func (*PresignResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PresignResult) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *PresignResult) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *PresignResult) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PresignResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PresignResult) GetSignedUrl() string {
	if m != nil {
		return m.SignedUrl
	}
	return ""
}

type File struct {
	Ctx         *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId   string          `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Name        string          `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Type        string          `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Size        int64           `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	Md5         string          `protobuf:"bytes,6,opt,name=md5" json:"md5,omitempty"`
	Description string          `protobuf:"bytes,10,opt,name=description" json:"description,omitempty"`
	Created     int64           `protobuf:"varint,7,opt,name=created" json:"created,omitempty"`
	Url         string          `protobuf:"bytes,8,opt,name=url" json:"url,omitempty"`
	Creator     string          `protobuf:"bytes,9,opt,name=creator" json:"creator,omitempty"`
	Id          string          `protobuf:"bytes,11,opt,name=id" json:"id,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *File) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *File) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *File) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *File) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *File) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *File) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *File) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *File) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *File) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*AllType)(nil), "file.AllType")
	proto.RegisterType((*FileHeader)(nil), "file.FileHeader")
	proto.RegisterType((*PresignResult)(nil), "file.PresignResult")
	proto.RegisterType((*File)(nil), "file.File")
	proto.RegisterEnum("file.Event", Event_name, Event_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MyServer service

type MyServerClient interface {
	Do(ctx context.Context, in *AllType, opts ...grpc.CallOption) (*AllType, error)
}

type myServerClient struct {
	cc *grpc.ClientConn
}

func NewMyServerClient(cc *grpc.ClientConn) MyServerClient {
	return &myServerClient{cc}
}

func (c *myServerClient) Do(ctx context.Context, in *AllType, opts ...grpc.CallOption) (*AllType, error) {
	out := new(AllType)
	err := grpc.Invoke(ctx, "/file.MyServer/Do", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MyServer service

type MyServerServer interface {
	Do(context.Context, *AllType) (*AllType, error)
}

func RegisterMyServerServer(s *grpc.Server, srv MyServerServer) {
	s.RegisterService(&_MyServer_serviceDesc, srv)
}

func _MyServer_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyServerServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.MyServer/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyServerServer).Do(ctx, req.(*AllType))
	}
	return interceptor(ctx, in, info, handler)
}

var _MyServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "file.MyServer",
	HandlerType: (*MyServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _MyServer_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file/file.proto",
}

func init() { proto.RegisterFile("file/file.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x6d, 0x9c, 0xfd, 0x68, 0x66, 0x29, 0x1b, 0x06, 0x90, 0xac, 0x4a, 0x48, 0x4b, 0xe0, 0x50,
	0x71, 0xc8, 0xa2, 0x22, 0x7e, 0x00, 0x2a, 0x20, 0x7a, 0x40, 0x42, 0x86, 0x9e, 0x38, 0x54, 0xd9,
	0x78, 0xd2, 0x35, 0x64, 0xe3, 0xe0, 0x38, 0x55, 0xb7, 0x3f, 0x82, 0x03, 0x3f, 0x88, 0xdf, 0x86,
	0xec, 0x24, 0x2c, 0x85, 0x13, 0xb7, 0x5e, 0x92, 0xe7, 0x37, 0xcf, 0xe3, 0x37, 0x2f, 0x31, 0xcc,
	0x0b, 0x55, 0xd2, 0xd2, 0x3d, 0xd2, 0xda, 0x68, 0xab, 0x71, 0xe4, 0xf0, 0x61, 0xba, 0x52, 0x76,
	0xd5, 0xe6, 0x5f, 0xc9, 0xa6, 0xda, 0x5c, 0x2c, 0x9b, 0x76, 0xa5, 0xae, 0x97, 0x6b, 0xca, 0x24,
	0x99, 0x65, 0xae, 0x37, 0x1b, 0x5d, 0xf5, 0xaf, 0x6e, 0x57, 0xf2, 0x05, 0xa6, 0xaf, 0xca, 0xf2,
	0xd3, 0xb6, 0x26, 0x5c, 0x00, 0x2b, 0xd6, 0x9c, 0x2d, 0x82, 0xa3, 0xd9, 0x71, 0x9c, 0xfa, 0xce,
	0x6f, 0x55, 0x49, 0xef, 0xfc, 0x66, 0xc1, 0x8a, 0x35, 0x3e, 0x01, 0x56, 0x37, 0x3c, 0xf4, 0x8a,
	0xfb, 0x9d, 0xe2, 0x83, 0xa1, 0x46, 0x5d, 0x54, 0x82, 0x9a, 0xb6, 0xb4, 0x82, 0xd5, 0x0d, 0x72,
	0x08, 0x0a, 0x3e, 0xf2, 0x1a, 0xd8, 0x75, 0x11, 0x41, 0x91, 0xfc, 0x0c, 0x00, 0x76, 0x1d, 0xf1,
	0x31, 0x84, 0xb9, 0xbd, 0xe2, 0x81, 0x97, 0xce, 0xd3, 0xde, 0xd6, 0x89, 0xae, 0x2c, 0x5d, 0x59,
	0xe1, 0x6a, 0xf8, 0x08, 0x20, 0xcb, 0x73, 0xdd, 0x56, 0xf6, 0x5c, 0x49, 0x6f, 0x2d, 0x12, 0x51,
	0xcf, 0x9c, 0x4a, 0x44, 0x18, 0x55, 0xd9, 0x86, 0xbc, 0xa3, 0x48, 0x78, 0xec, 0x38, 0xbb, 0xad,
	0xc9, 0x3b, 0x88, 0x84, 0xc7, 0x8e, 0x6b, 0xd4, 0x35, 0xf1, 0xf1, 0x22, 0x38, 0x0a, 0x85, 0xc7,
	0x18, 0x43, 0xb8, 0x91, 0x2f, 0xf9, 0xc4, 0xcb, 0x1c, 0xc4, 0x05, 0xcc, 0x24, 0x35, 0xb9, 0x51,
	0xb5, 0x55, 0xba, 0xe2, 0x53, 0x5f, 0xf9, 0x93, 0x4a, 0x7e, 0x04, 0x70, 0x70, 0x63, 0xe0, 0xff,
	0x9f, 0x21, 0xfc, 0x7b, 0x86, 0x18, 0xc2, 0xd6, 0x94, 0xbd, 0x5d, 0x07, 0xf1, 0x2e, 0x30, 0x25,
	0x7b, 0x63, 0x4c, 0x49, 0xd7, 0xc0, 0x9d, 0x48, 0xf2, 0xdc, 0x09, 0xc7, 0x5d, 0x83, 0x8e, 0x39,
	0x33, 0x65, 0xf2, 0x9d, 0xc1, 0xc8, 0xa5, 0x7a, 0xdb, 0xf3, 0x84, 0x7f, 0xf2, 0x44, 0x0e, 0xd3,
	0xdc, 0x50, 0x66, 0x49, 0xfa, 0xb4, 0x43, 0x31, 0x2c, 0x87, 0x54, 0xf6, 0x77, 0xa9, 0x0c, 0x5a,
	0x6d, 0x78, 0xe4, 0xd9, 0x61, 0xd9, 0xe7, 0x35, 0x1b, 0xf2, 0x7a, 0xf6, 0x19, 0xc6, 0x6f, 0x2e,
	0xa9, 0xb2, 0xc8, 0xe1, 0x81, 0x0b, 0xe6, 0xf7, 0x17, 0xfb, 0xd6, 0x52, 0x63, 0x49, 0xc6, 0x7b,
	0x38, 0x87, 0x99, 0xab, 0x9c, 0x74, 0xa7, 0xc5, 0x21, 0x3e, 0x84, 0x7b, 0x8e, 0x38, 0xad, 0x0a,
	0xbd, 0xd3, 0x8d, 0x30, 0x86, 0x3b, 0x8e, 0x3e, 0xab, 0x4b, 0x9d, 0x49, 0x92, 0xf1, 0xe4, 0xf8,
	0x39, 0xec, 0xbf, 0xdf, 0x7e, 0x24, 0x73, 0x49, 0x06, 0x9f, 0x02, 0x7b, 0xad, 0xf1, 0xa0, 0xfb,
	0xc9, 0xfb, 0x5b, 0x74, 0x78, 0x73, 0x99, 0xec, 0xad, 0x26, 0xfe, 0xa2, 0xbd, 0xf8, 0x15, 0x00,
	0x00, 0xff, 0xff, 0x53, 0xc7, 0x14, 0x8b, 0xb1, 0x03, 0x00, 0x00,
}
