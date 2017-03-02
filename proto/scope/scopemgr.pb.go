// Code generated by protoc-gen-go.
// source: scope/scopemgr.proto
// DO NOT EDIT!

/*
Package scope is a generated protocol buffer package.

It is generated from these files:
	scope/scopemgr.proto

It has these top-level messages:
	ScopeCreateResponse
	Scope
	ScopeIdRequest
	ScopeResponse
	ListScopesResponse
	Response
	Empty
*/
package scope

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

type ErrorCode int32

const (
	ErrorCode_no_error         ErrorCode = 0
	ErrorCode_resource_notfoud ErrorCode = 1
	ErrorCode_invalid_request  ErrorCode = 2
	ErrorCode_unauthorized     ErrorCode = 3
	ErrorCode_internal_error   ErrorCode = 4
)

var ErrorCode_name = map[int32]string{
	0: "no_error",
	1: "resource_notfoud",
	2: "invalid_request",
	3: "unauthorized",
	4: "internal_error",
}
var ErrorCode_value = map[string]int32{
	"no_error":         0,
	"resource_notfoud": 1,
	"invalid_request":  2,
	"unauthorized":     3,
	"internal_error":   4,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ScopeCreateResponse struct {
	Error            ErrorCode `protobuf:"varint,1,opt,name=error,enum=scope.ErrorCode" json:"error,omitempty"`
	ErrorDescription string    `protobuf:"bytes,2,opt,name=error_description,json=errorDescription" json:"error_description,omitempty"`
	Id               string    `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *ScopeCreateResponse) Reset()                    { *m = ScopeCreateResponse{} }
func (m *ScopeCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*ScopeCreateResponse) ProtoMessage()               {}
func (*ScopeCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ScopeCreateResponse) GetError() ErrorCode {
	if m != nil {
		return m.Error
	}
	return ErrorCode_no_error
}

func (m *ScopeCreateResponse) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

func (m *ScopeCreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Scope struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	LogoUrl     string   `protobuf:"bytes,3,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
	Title       string   `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Description string   `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Methods     []string `protobuf:"bytes,6,rep,name=methods" json:"methods,omitempty"`
}

func (m *Scope) Reset()                    { *m = Scope{} }
func (m *Scope) String() string            { return proto.CompactTextString(m) }
func (*Scope) ProtoMessage()               {}
func (*Scope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Scope) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Scope) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Scope) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func (m *Scope) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Scope) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Scope) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

type ScopeIdRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ScopeIdRequest) Reset()                    { *m = ScopeIdRequest{} }
func (m *ScopeIdRequest) String() string            { return proto.CompactTextString(m) }
func (*ScopeIdRequest) ProtoMessage()               {}
func (*ScopeIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ScopeIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ScopeResponse struct {
	Id               string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name             string    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	LogoUrl          string    `protobuf:"bytes,3,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
	Title            string    `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Description      string    `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Methods          []string  `protobuf:"bytes,6,rep,name=methods" json:"methods,omitempty"`
	Error            ErrorCode `protobuf:"varint,11,opt,name=error,enum=scope.ErrorCode" json:"error,omitempty"`
	ErrorDescription string    `protobuf:"bytes,12,opt,name=error_description,json=errorDescription" json:"error_description,omitempty"`
}

func (m *ScopeResponse) Reset()                    { *m = ScopeResponse{} }
func (m *ScopeResponse) String() string            { return proto.CompactTextString(m) }
func (*ScopeResponse) ProtoMessage()               {}
func (*ScopeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ScopeResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ScopeResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScopeResponse) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func (m *ScopeResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ScopeResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ScopeResponse) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *ScopeResponse) GetError() ErrorCode {
	if m != nil {
		return m.Error
	}
	return ErrorCode_no_error
}

func (m *ScopeResponse) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

type ListScopesResponse struct {
	Error            ErrorCode `protobuf:"varint,1,opt,name=error,enum=scope.ErrorCode" json:"error,omitempty"`
	ErrorDescription string    `protobuf:"bytes,2,opt,name=error_description,json=errorDescription" json:"error_description,omitempty"`
	Scopes           []*Scope  `protobuf:"bytes,3,rep,name=scopes" json:"scopes,omitempty"`
}

func (m *ListScopesResponse) Reset()                    { *m = ListScopesResponse{} }
func (m *ListScopesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListScopesResponse) ProtoMessage()               {}
func (*ListScopesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListScopesResponse) GetError() ErrorCode {
	if m != nil {
		return m.Error
	}
	return ErrorCode_no_error
}

func (m *ListScopesResponse) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

func (m *ListScopesResponse) GetScopes() []*Scope {
	if m != nil {
		return m.Scopes
	}
	return nil
}

type Response struct {
	Error            ErrorCode `protobuf:"varint,1,opt,name=error,enum=scope.ErrorCode" json:"error,omitempty"`
	ErrorDescription string    `protobuf:"bytes,2,opt,name=error_description,json=errorDescription" json:"error_description,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Response) GetError() ErrorCode {
	if m != nil {
		return m.Error
	}
	return ErrorCode_no_error
}

func (m *Response) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*ScopeCreateResponse)(nil), "scope.ScopeCreateResponse")
	proto.RegisterType((*Scope)(nil), "scope.Scope")
	proto.RegisterType((*ScopeIdRequest)(nil), "scope.ScopeIdRequest")
	proto.RegisterType((*ScopeResponse)(nil), "scope.ScopeResponse")
	proto.RegisterType((*ListScopesResponse)(nil), "scope.ListScopesResponse")
	proto.RegisterType((*Response)(nil), "scope.Response")
	proto.RegisterType((*Empty)(nil), "scope.Empty")
	proto.RegisterEnum("scope.ErrorCode", ErrorCode_name, ErrorCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ScopeMgt service

type ScopeMgtClient interface {
	Create(ctx context.Context, in *Scope, opts ...grpc.CallOption) (*ScopeCreateResponse, error)
	Read(ctx context.Context, in *ScopeIdRequest, opts ...grpc.CallOption) (*ScopeResponse, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListScopesResponse, error)
	Delete(ctx context.Context, in *ScopeIdRequest, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *Scope, opts ...grpc.CallOption) (*Response, error)
}

type scopeMgtClient struct {
	cc *grpc.ClientConn
}

func NewScopeMgtClient(cc *grpc.ClientConn) ScopeMgtClient {
	return &scopeMgtClient{cc}
}

func (c *scopeMgtClient) Create(ctx context.Context, in *Scope, opts ...grpc.CallOption) (*ScopeCreateResponse, error) {
	out := new(ScopeCreateResponse)
	err := grpc.Invoke(ctx, "/scope.ScopeMgt/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopeMgtClient) Read(ctx context.Context, in *ScopeIdRequest, opts ...grpc.CallOption) (*ScopeResponse, error) {
	out := new(ScopeResponse)
	err := grpc.Invoke(ctx, "/scope.ScopeMgt/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopeMgtClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListScopesResponse, error) {
	out := new(ListScopesResponse)
	err := grpc.Invoke(ctx, "/scope.ScopeMgt/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopeMgtClient) Delete(ctx context.Context, in *ScopeIdRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/scope.ScopeMgt/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scopeMgtClient) Update(ctx context.Context, in *Scope, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/scope.ScopeMgt/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ScopeMgt service

type ScopeMgtServer interface {
	Create(context.Context, *Scope) (*ScopeCreateResponse, error)
	Read(context.Context, *ScopeIdRequest) (*ScopeResponse, error)
	List(context.Context, *Empty) (*ListScopesResponse, error)
	Delete(context.Context, *ScopeIdRequest) (*Response, error)
	Update(context.Context, *Scope) (*Response, error)
}

func RegisterScopeMgtServer(s *grpc.Server, srv ScopeMgtServer) {
	s.RegisterService(&_ScopeMgt_serviceDesc, srv)
}

func _ScopeMgt_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Scope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopeMgtServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scope.ScopeMgt/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopeMgtServer).Create(ctx, req.(*Scope))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopeMgt_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScopeIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopeMgtServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scope.ScopeMgt/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopeMgtServer).Read(ctx, req.(*ScopeIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopeMgt_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopeMgtServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scope.ScopeMgt/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopeMgtServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopeMgt_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScopeIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopeMgtServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scope.ScopeMgt/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopeMgtServer).Delete(ctx, req.(*ScopeIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScopeMgt_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Scope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopeMgtServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scope.ScopeMgt/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopeMgtServer).Update(ctx, req.(*Scope))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScopeMgt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scope.ScopeMgt",
	HandlerType: (*ScopeMgtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ScopeMgt_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ScopeMgt_Read_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ScopeMgt_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ScopeMgt_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ScopeMgt_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scope/scopemgr.proto",
}

func init() { proto.RegisterFile("scope/scopemgr.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x7f, 0x93, 0x4e, 0x42, 0x6a, 0xa6, 0x41, 0x72, 0x73, 0xb2, 0x2c, 0x84, 0x02, 0x48,
	0x41, 0x04, 0x78, 0x82, 0xb6, 0x07, 0x24, 0xb8, 0x18, 0xf5, 0x6c, 0x99, 0xec, 0x90, 0x2e, 0x72,
	0xbc, 0x66, 0x77, 0x8d, 0x44, 0xdf, 0x80, 0x13, 0x57, 0x5e, 0x94, 0x3b, 0xf2, 0x7a, 0x49, 0x1c,
	0x02, 0x17, 0xa4, 0x4a, 0x5c, 0xac, 0xfd, 0xe6, 0x9b, 0xfd, 0xe6, 0x77, 0x0d, 0x33, 0xb5, 0x16,
	0x35, 0x3d, 0x33, 0xdf, 0xed, 0x46, 0x2e, 0x6b, 0x29, 0xb4, 0xc0, 0xc0, 0xe0, 0xf4, 0x16, 0xce,
	0xde, 0xb5, 0x87, 0x0b, 0x49, 0x85, 0xa6, 0x8c, 0x54, 0x2d, 0x2a, 0x45, 0xf8, 0x08, 0x02, 0x92,
	0x52, 0xc8, 0xd8, 0x49, 0x9c, 0xc5, 0x74, 0x15, 0x2d, 0x8d, 0xf7, 0xf2, 0xaa, 0xb5, 0x5d, 0x08,
	0x46, 0x59, 0x47, 0xe3, 0x53, 0xb8, 0x6f, 0x0e, 0x39, 0x23, 0xb5, 0x96, 0xbc, 0xd6, 0x5c, 0x54,
	0xb1, 0x9b, 0x38, 0x8b, 0x93, 0x2c, 0x32, 0xc4, 0xe5, 0xde, 0x8e, 0x53, 0x70, 0x39, 0x8b, 0x3d,
	0xc3, 0xba, 0x9c, 0xa5, 0xdf, 0x1d, 0x08, 0x4c, 0x70, 0xcb, 0x38, 0xbf, 0x18, 0x44, 0xf0, 0xab,
	0x62, 0x4b, 0x56, 0xc9, 0x9c, 0xf1, 0x1c, 0x46, 0xa5, 0xd8, 0x88, 0xbc, 0x91, 0xa5, 0xd5, 0x18,
	0xb6, 0xf8, 0x5a, 0x96, 0x38, 0x83, 0x40, 0x73, 0x5d, 0x52, 0xec, 0x1b, 0x7b, 0x07, 0x30, 0x81,
	0x71, 0x3f, 0xab, 0xc0, 0x70, 0x7d, 0x13, 0xc6, 0x30, 0xdc, 0x92, 0xbe, 0x11, 0x4c, 0xc5, 0x61,
	0xe2, 0xb5, 0x8a, 0x16, 0xa6, 0x09, 0x4c, 0x4d, 0x66, 0xaf, 0x59, 0x46, 0x9f, 0x1a, 0x52, 0xfa,
	0xf7, 0x14, 0xd3, 0x1f, 0x0e, 0xdc, 0x33, 0x2e, 0xbb, 0x9e, 0xfd, 0x6f, 0x45, 0xec, 0x87, 0x38,
	0xfe, 0x87, 0x21, 0x4e, 0xfe, 0x3c, 0xc4, 0xf4, 0x9b, 0x03, 0xf8, 0x86, 0x2b, 0x6d, 0x6a, 0x57,
	0x77, 0xbb, 0x30, 0x0f, 0x21, 0x34, 0x32, 0x2a, 0xf6, 0x12, 0x6f, 0x31, 0x5e, 0x4d, 0xac, 0x6a,
	0xd7, 0x77, 0xcb, 0xa5, 0x39, 0x8c, 0xee, 0x34, 0x8d, 0x74, 0x08, 0xc1, 0xd5, 0xb6, 0xd6, 0x5f,
	0x9e, 0x7c, 0x84, 0x93, 0x9d, 0x12, 0x4e, 0x60, 0x54, 0x89, 0xdc, 0x38, 0x47, 0x03, 0x9c, 0x41,
	0x24, 0x49, 0x89, 0x46, 0xae, 0x29, 0xaf, 0x84, 0xfe, 0x20, 0x1a, 0x16, 0x39, 0x78, 0x06, 0xa7,
	0xbc, 0xfa, 0x5c, 0x94, 0x9c, 0xe5, 0xb2, 0xdb, 0xa3, 0xc8, 0xc5, 0x08, 0x26, 0x4d, 0x55, 0x34,
	0xfa, 0x46, 0x48, 0x7e, 0x4b, 0x2c, 0xf2, 0x10, 0x61, 0xca, 0x2b, 0x4d, 0xb2, 0x2a, 0x4a, 0x2b,
	0xe8, 0xaf, 0xbe, 0xba, 0x30, 0x32, 0x75, 0xbe, 0xdd, 0x68, 0x7c, 0x09, 0x61, 0xf7, 0x40, 0xf1,
	0xa0, 0x05, 0xf3, 0x79, 0x1f, 0x1d, 0x3e, 0xe1, 0x74, 0x80, 0xaf, 0xc0, 0xcf, 0xa8, 0x60, 0xf8,
	0xa0, 0xef, 0xb5, 0xdb, 0xe8, 0xf9, 0xec, 0xa0, 0x9b, 0xfb, 0x6b, 0xcf, 0xc1, 0x6f, 0x07, 0xbc,
	0x0b, 0x65, 0x6a, 0x9f, 0x9f, 0x5b, 0x74, 0x3c, 0xfb, 0x74, 0x80, 0x2b, 0x08, 0x2f, 0xa9, 0x24,
	0x4d, 0x7f, 0x8b, 0x75, 0x6a, 0xcd, 0xbd, 0x3b, 0x8f, 0x21, 0xbc, 0xae, 0xd9, 0x71, 0x4d, 0xc7,
	0xae, 0xef, 0x43, 0xf3, 0xcb, 0x7a, 0xf1, 0x33, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xf6, 0x0d, 0x6d,
	0xca, 0x04, 0x00, 0x00,
}
