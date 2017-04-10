// Code generated by protoc-gen-go.
// source: content/content.proto
// DO NOT EDIT!

/*
Package content is a generated protocol buffer package.

It is generated from these files:
	content/content.proto

It has these top-level messages:
	KeyValue
	Content
	Contents
	Ids
	Id
	Empty
	ListRequest
*/
package content

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

type KeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Content struct {
	Id             string      `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	AccountId      string      `protobuf:"bytes,14,opt,name=AccountId" json:"AccountId,omitempty"`
	Description    string      `protobuf:"bytes,2,opt,name=Description" json:"Description,omitempty"`
	Title          string      `protobuf:"bytes,3,opt,name=Title" json:"Title,omitempty"`
	Link           string      `protobuf:"bytes,4,opt,name=Link" json:"Link,omitempty"`
	Type           string      `protobuf:"bytes,5,opt,name=Type" json:"Type,omitempty"`
	Category       string      `protobuf:"bytes,6,opt,name=Category" json:"Category,omitempty"`
	AttachmentLink string      `protobuf:"bytes,7,opt,name=AttachmentLink" json:"AttachmentLink,omitempty"`
	Labels         string      `protobuf:"bytes,8,opt,name=Labels" json:"Labels,omitempty"`
	Availability   bool        `protobuf:"varint,9,opt,name=Availability" json:"Availability,omitempty"`
	Price          string      `protobuf:"bytes,10,opt,name=Price" json:"Price,omitempty"`
	Currency       string      `protobuf:"bytes,11,opt,name=Currency" json:"Currency,omitempty"`
	SalePrice      string      `protobuf:"bytes,12,opt,name=SalePrice" json:"SalePrice,omitempty"`
	Fields         []*KeyValue `protobuf:"bytes,13,rep,name=Fields" json:"Fields,omitempty"`
	RelatedIds     []string    `protobuf:"bytes,15,rep,name=RelatedIds" json:"RelatedIds,omitempty"`
}

func (m *Content) Reset()                    { *m = Content{} }
func (m *Content) String() string            { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()               {}
func (*Content) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Content) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Content) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Content) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Content) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Content) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Content) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Content) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Content) GetAttachmentLink() string {
	if m != nil {
		return m.AttachmentLink
	}
	return ""
}

func (m *Content) GetLabels() string {
	if m != nil {
		return m.Labels
	}
	return ""
}

func (m *Content) GetAvailability() bool {
	if m != nil {
		return m.Availability
	}
	return false
}

func (m *Content) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Content) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Content) GetSalePrice() string {
	if m != nil {
		return m.SalePrice
	}
	return ""
}

func (m *Content) GetFields() []*KeyValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Content) GetRelatedIds() []string {
	if m != nil {
		return m.RelatedIds
	}
	return nil
}

type Contents struct {
	Contents []*Content `protobuf:"bytes,1,rep,name=Contents" json:"Contents,omitempty"`
}

func (m *Contents) Reset()                    { *m = Contents{} }
func (m *Contents) String() string            { return proto.CompactTextString(m) }
func (*Contents) ProtoMessage()               {}
func (*Contents) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Contents) GetContents() []*Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

type Ids struct {
	Id []string `protobuf:"bytes,1,rep,name=Id" json:"Id,omitempty"`
}

func (m *Ids) Reset()                    { *m = Ids{} }
func (m *Ids) String() string            { return proto.CompactTextString(m) }
func (*Ids) ProtoMessage()               {}
func (*Ids) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Ids) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

type Id struct {
	Id string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
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

type ListRequest struct {
	AccountId string `protobuf:"bytes,1,opt,name=AccountId" json:"AccountId,omitempty"`
	StartId   string `protobuf:"bytes,2,opt,name=StartId" json:"StartId,omitempty"`
	Limit     int32  `protobuf:"varint,3,opt,name=Limit" json:"Limit,omitempty"`
	Keyword   string `protobuf:"bytes,4,opt,name=Keyword" json:"Keyword,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

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

func init() {
	proto.RegisterType((*KeyValue)(nil), "content.KeyValue")
	proto.RegisterType((*Content)(nil), "content.Content")
	proto.RegisterType((*Contents)(nil), "content.Contents")
	proto.RegisterType((*Ids)(nil), "content.Ids")
	proto.RegisterType((*Id)(nil), "content.Id")
	proto.RegisterType((*Empty)(nil), "content.Empty")
	proto.RegisterType((*ListRequest)(nil), "content.ListRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContentMgr service

type ContentMgrClient interface {
	Insert(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Empty, error)
	InsertBulk(ctx context.Context, in *Contents, opts ...grpc.CallOption) (*Empty, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	DeleteBulk(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error)
	Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Content, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Contents, error)
}

type contentMgrClient struct {
	cc *grpc.ClientConn
}

func NewContentMgrClient(cc *grpc.ClientConn) ContentMgrClient {
	return &contentMgrClient{cc}
}

func (c *contentMgrClient) Insert(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/content.ContentMgr/Insert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentMgrClient) InsertBulk(ctx context.Context, in *Contents, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/content.ContentMgr/InsertBulk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentMgrClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/content.ContentMgr/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentMgrClient) DeleteBulk(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/content.ContentMgr/DeleteBulk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentMgrClient) Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := grpc.Invoke(ctx, "/content.ContentMgr/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentMgrClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Contents, error) {
	out := new(Contents)
	err := grpc.Invoke(ctx, "/content.ContentMgr/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContentMgr service

type ContentMgrServer interface {
	Insert(context.Context, *Content) (*Empty, error)
	InsertBulk(context.Context, *Contents) (*Empty, error)
	Delete(context.Context, *Id) (*Empty, error)
	DeleteBulk(context.Context, *Ids) (*Empty, error)
	Read(context.Context, *Id) (*Content, error)
	List(context.Context, *ListRequest) (*Contents, error)
}

func RegisterContentMgrServer(s *grpc.Server, srv ContentMgrServer) {
	s.RegisterService(&_ContentMgr_serviceDesc, srv)
}

func _ContentMgr_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).Insert(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentMgr_InsertBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Contents)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).InsertBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/InsertBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).InsertBulk(ctx, req.(*Contents))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentMgr_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentMgr_DeleteBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).DeleteBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/DeleteBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).DeleteBulk(ctx, req.(*Ids))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentMgr_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).Read(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentMgr_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentMgrServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ContentMgr/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentMgrServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "content.ContentMgr",
	HandlerType: (*ContentMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _ContentMgr_Insert_Handler,
		},
		{
			MethodName: "InsertBulk",
			Handler:    _ContentMgr_InsertBulk_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ContentMgr_Delete_Handler,
		},
		{
			MethodName: "DeleteBulk",
			Handler:    _ContentMgr_DeleteBulk_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ContentMgr_Read_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ContentMgr_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content/content.proto",
}

func init() { proto.RegisterFile("content/content.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x54, 0xdd, 0x6e, 0xda, 0x4c,
	0x10, 0x0d, 0x18, 0x0c, 0x0c, 0x7c, 0x7c, 0xc9, 0x88, 0x54, 0x2b, 0x54, 0x55, 0xc8, 0x17, 0x2d,
	0x95, 0x50, 0xaa, 0xd2, 0x9b, 0xde, 0xd2, 0xa4, 0x95, 0x2c, 0xa8, 0x54, 0x39, 0x51, 0xef, 0x8d,
	0x3d, 0x4a, 0x57, 0x31, 0x36, 0xdd, 0x5d, 0xb7, 0xf2, 0xdb, 0xf4, 0x19, 0xfa, 0x84, 0xd5, 0xfe,
	0xe0, 0xf0, 0x77, 0xc5, 0x9c, 0x33, 0x67, 0xc6, 0x67, 0x67, 0x46, 0xc0, 0x75, 0x52, 0xe4, 0x8a,
	0x72, 0xf5, 0xce, 0xfd, 0xde, 0x6c, 0x45, 0xa1, 0x0a, 0xec, 0x38, 0x18, 0xcc, 0xa1, 0xbb, 0xa4,
	0xea, 0x7b, 0x9c, 0x95, 0x84, 0x97, 0xe0, 0x2d, 0xa9, 0x62, 0x8d, 0x49, 0x63, 0xda, 0x8b, 0x74,
	0x88, 0x23, 0x68, 0x9b, 0x14, 0x6b, 0x1a, 0xce, 0x82, 0xe0, 0xaf, 0x07, 0x9d, 0x5b, 0x5b, 0x8f,
	0x43, 0x68, 0x86, 0xa9, 0x2b, 0x69, 0x86, 0x29, 0xbe, 0x84, 0xde, 0x22, 0x49, 0x8a, 0x32, 0x57,
	0x61, 0xca, 0x86, 0x86, 0x7e, 0x26, 0x70, 0x02, 0xfd, 0x3b, 0x92, 0x89, 0xe0, 0x5b, 0xc5, 0x8b,
	0xdc, 0x75, 0xdd, 0xa7, 0xf4, 0x17, 0x1f, 0xb8, 0xca, 0x88, 0x79, 0xf6, 0x8b, 0x06, 0x20, 0x42,
	0x6b, 0xc5, 0xf3, 0x27, 0xd6, 0x32, 0xa4, 0x89, 0x35, 0xf7, 0x50, 0x6d, 0x89, 0xb5, 0x2d, 0xa7,
	0x63, 0x1c, 0x43, 0xf7, 0x36, 0x56, 0xf4, 0x58, 0x88, 0x8a, 0xf9, 0x86, 0xaf, 0x31, 0xbe, 0x86,
	0xe1, 0x42, 0xa9, 0x38, 0xf9, 0xb1, 0xa1, 0x5c, 0x99, 0x6e, 0x1d, 0xa3, 0x38, 0x62, 0xf1, 0x05,
	0xf8, 0xab, 0x78, 0x4d, 0x99, 0x64, 0x5d, 0x93, 0x77, 0x08, 0x03, 0x18, 0x2c, 0x7e, 0xc5, 0x3c,
	0x8b, 0xd7, 0x3c, 0xe3, 0xaa, 0x62, 0xbd, 0x49, 0x63, 0xda, 0x8d, 0x0e, 0x38, 0xed, 0xfe, 0x9b,
	0xe0, 0x09, 0x31, 0xb0, 0xee, 0x0d, 0x30, 0xae, 0x4a, 0x21, 0x28, 0x4f, 0x2a, 0xd6, 0x77, 0xae,
	0x1c, 0xd6, 0xf3, 0xba, 0x8f, 0x33, 0xb2, 0x55, 0x03, 0x3b, 0xaf, 0x9a, 0xc0, 0xb7, 0xe0, 0x7f,
	0xe1, 0x94, 0xa5, 0x92, 0xfd, 0x37, 0xf1, 0xa6, 0xfd, 0xf9, 0xd5, 0xcd, 0x6e, 0x8d, 0xbb, 0xa5,
	0x45, 0x4e, 0x80, 0xaf, 0x00, 0x22, 0xca, 0x62, 0x45, 0x69, 0x98, 0x4a, 0xf6, 0xff, 0xc4, 0x9b,
	0xf6, 0xa2, 0x3d, 0x26, 0xf8, 0x08, 0x5d, 0xb7, 0x33, 0x89, 0xb3, 0xe7, 0x98, 0x35, 0x4c, 0xe3,
	0xcb, 0xba, 0xb1, 0x4b, 0x44, 0xb5, 0x22, 0xb8, 0x06, 0x2f, 0x4c, 0x65, 0xbd, 0x69, 0xcf, 0x6e,
	0x3a, 0x18, 0x69, 0x7c, 0xbc, 0xff, 0xa0, 0x03, 0xed, 0xcf, 0x9b, 0xad, 0xaa, 0x82, 0x12, 0xfa,
	0x2b, 0x2e, 0x55, 0x44, 0x3f, 0x4b, 0x92, 0xea, 0xf0, 0x2e, 0x1a, 0xc7, 0x77, 0xc1, 0xa0, 0x73,
	0xaf, 0x62, 0xa1, 0x73, 0xf6, 0x26, 0x76, 0x50, 0x4f, 0x74, 0xc5, 0x37, 0x5c, 0x99, 0x7b, 0x68,
	0x47, 0x16, 0x68, 0xfd, 0x92, 0xaa, 0xdf, 0x85, 0x48, 0xdd, 0x49, 0xec, 0xe0, 0xfc, 0x4f, 0x13,
	0xc0, 0x39, 0xff, 0xfa, 0x28, 0x70, 0x06, 0x7e, 0x98, 0x4b, 0x12, 0x0a, 0x4f, 0x5e, 0x38, 0x1e,
	0xd6, 0x8c, 0x75, 0x7c, 0x81, 0xef, 0x01, 0xac, 0xfa, 0x53, 0x99, 0x3d, 0xe1, 0xd5, 0x71, 0x85,
	0x3c, 0x53, 0xf2, 0x06, 0xfc, 0x3b, 0xca, 0x48, 0x11, 0xf6, 0xeb, 0x5c, 0x98, 0x9e, 0x11, 0xce,
	0x00, 0xac, 0xd0, 0xf4, 0x1e, 0xec, 0x89, 0xcf, 0xb7, 0x6d, 0x45, 0x14, 0xa7, 0x87, 0x4d, 0x4f,
	0x9e, 0x60, 0x2c, 0xb7, 0xf4, 0x98, 0x71, 0x54, 0xe7, 0xf6, 0xa6, 0x3e, 0x3e, 0x7d, 0x42, 0x70,
	0xb1, 0xf6, 0xcd, 0x5f, 0xc0, 0x87, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x29, 0xed, 0xb3, 0xf5,
	0x1b, 0x04, 0x00, 0x00,
}
