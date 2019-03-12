// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file/file.proto

package file

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/subiz/header/common"
	math "math"
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

type AllType struct {
	Fh                   *FileHeader    `protobuf:"bytes,2,opt,name=fh,proto3" json:"fh,omitempty"`
	Ps                   *PresignResult `protobuf:"bytes,3,opt,name=ps,proto3" json:"ps,omitempty"`
	F                    *File          `protobuf:"bytes,4,opt,name=f,proto3" json:"f,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AllType) Reset()         { *m = AllType{} }
func (m *AllType) String() string { return proto.CompactTextString(m) }
func (*AllType) ProtoMessage()    {}
func (*AllType) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{0}
}

func (m *AllType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllType.Unmarshal(m, b)
}
func (m *AllType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllType.Marshal(b, m, deterministic)
}
func (m *AllType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllType.Merge(m, src)
}
func (m *AllType) XXX_Size() int {
	return xxx_messageInfo_AllType.Size(m)
}
func (m *AllType) XXX_DiscardUnknown() {
	xxx_messageInfo_AllType.DiscardUnknown(m)
}

var xxx_messageInfo_AllType proto.InternalMessageInfo

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
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name                 string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string          `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Size                 int64           `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Md5                  string          `protobuf:"bytes,6,opt,name=md5,proto3" json:"md5,omitempty"`
	Description          string          `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FileHeader) Reset()         { *m = FileHeader{} }
func (m *FileHeader) String() string { return proto.CompactTextString(m) }
func (*FileHeader) ProtoMessage()    {}
func (*FileHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{1}
}

func (m *FileHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileHeader.Unmarshal(m, b)
}
func (m *FileHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileHeader.Marshal(b, m, deterministic)
}
func (m *FileHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileHeader.Merge(m, src)
}
func (m *FileHeader) XXX_Size() int {
	return xxx_messageInfo_FileHeader.Size(m)
}
func (m *FileHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_FileHeader.DiscardUnknown(m)
}

var xxx_messageInfo_FileHeader proto.InternalMessageInfo

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
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Url                  string          `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Id                   string          `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	SignedUrl            string          `protobuf:"bytes,5,opt,name=signed_url,json=signedUrl,proto3" json:"signed_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PresignResult) Reset()         { *m = PresignResult{} }
func (m *PresignResult) String() string { return proto.CompactTextString(m) }
func (*PresignResult) ProtoMessage()    {}
func (*PresignResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{2}
}

func (m *PresignResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PresignResult.Unmarshal(m, b)
}
func (m *PresignResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PresignResult.Marshal(b, m, deterministic)
}
func (m *PresignResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresignResult.Merge(m, src)
}
func (m *PresignResult) XXX_Size() int {
	return xxx_messageInfo_PresignResult.Size(m)
}
func (m *PresignResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PresignResult.DiscardUnknown(m)
}

var xxx_messageInfo_PresignResult proto.InternalMessageInfo

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

type FileRequest struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Id                   string          `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{3}
}

func (m *FileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileRequest.Unmarshal(m, b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
}
func (m *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(m, src)
}
func (m *FileRequest) XXX_Size() int {
	return xxx_messageInfo_FileRequest.Size(m)
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

func (m *FileRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *FileRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *FileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type File struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name                 string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string          `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Size                 int64           `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Md5                  string          `protobuf:"bytes,6,opt,name=md5,proto3" json:"md5,omitempty"`
	Description          string          `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	Created              int64           `protobuf:"varint,7,opt,name=created,proto3" json:"created,omitempty"`
	Url                  string          `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	Creator              string          `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
	Id                   string          `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{4}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

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

type FileMeta struct {
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Uploaded             int64    `protobuf:"varint,4,opt,name=uploaded,proto3" json:"uploaded,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64    `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	FullUrl              string   `protobuf:"bytes,7,opt,name=full_url,json=fullUrl,proto3" json:"full_url,omitempty"`
	MimeType             string   `protobuf:"bytes,8,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Downloaded           int64    `protobuf:"varint,10,opt,name=downloaded,proto3" json:"downloaded,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileMeta) Reset()         { *m = FileMeta{} }
func (m *FileMeta) String() string { return proto.CompactTextString(m) }
func (*FileMeta) ProtoMessage()    {}
func (*FileMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad806f8986a0c3f6, []int{5}
}

func (m *FileMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileMeta.Unmarshal(m, b)
}
func (m *FileMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileMeta.Marshal(b, m, deterministic)
}
func (m *FileMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileMeta.Merge(m, src)
}
func (m *FileMeta) XXX_Size() int {
	return xxx_messageInfo_FileMeta.Size(m)
}
func (m *FileMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_FileMeta.DiscardUnknown(m)
}

var xxx_messageInfo_FileMeta proto.InternalMessageInfo

func (m *FileMeta) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *FileMeta) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FileMeta) GetUploaded() int64 {
	if m != nil {
		return m.Uploaded
	}
	return 0
}

func (m *FileMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FileMeta) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileMeta) GetFullUrl() string {
	if m != nil {
		return m.FullUrl
	}
	return ""
}

func (m *FileMeta) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *FileMeta) GetDownloaded() int64 {
	if m != nil {
		return m.Downloaded
	}
	return 0
}

func init() {
	proto.RegisterType((*AllType)(nil), "file.AllType")
	proto.RegisterType((*FileHeader)(nil), "file.FileHeader")
	proto.RegisterType((*PresignResult)(nil), "file.PresignResult")
	proto.RegisterType((*FileRequest)(nil), "file.FileRequest")
	proto.RegisterType((*File)(nil), "file.File")
	proto.RegisterType((*FileMeta)(nil), "file.FileMeta")
}

func init() { proto.RegisterFile("file/file.proto", fileDescriptor_ad806f8986a0c3f6) }

var fileDescriptor_ad806f8986a0c3f6 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcf, 0x8e, 0xd3, 0x3e,
	0x10, 0xc7, 0x95, 0xb8, 0x7f, 0x92, 0xa9, 0x7e, 0xbf, 0xad, 0xcc, 0xc5, 0x2c, 0x02, 0x95, 0x70,
	0x59, 0x09, 0xa9, 0x95, 0x40, 0x3c, 0x00, 0x42, 0x42, 0x70, 0x40, 0x42, 0x16, 0x7b, 0xae, 0xd2,
	0x78, 0xb2, 0x35, 0x72, 0xe2, 0xe0, 0x38, 0x62, 0x77, 0x1f, 0x82, 0x03, 0x0f, 0xc4, 0x5b, 0xf0,
	0x3e, 0xc8, 0x93, 0xa4, 0xed, 0xc2, 0x69, 0x25, 0x0e, 0x5c, 0x92, 0x99, 0xaf, 0xc7, 0xf3, 0xe7,
	0x63, 0x27, 0x70, 0x56, 0x6a, 0x83, 0x9b, 0xf0, 0x58, 0x37, 0xce, 0x7a, 0xcb, 0x27, 0xc1, 0x3e,
	0x7f, 0x7e, 0xa5, 0xfd, 0xbe, 0xdb, 0xad, 0x0b, 0x5b, 0x6d, 0xda, 0x6e, 0xa7, 0x6f, 0x37, 0x7b,
	0xcc, 0x15, 0xba, 0x4d, 0x61, 0xab, 0xca, 0xd6, 0xc3, 0xab, 0xdf, 0x92, 0x7d, 0x86, 0xf9, 0x6b,
	0x63, 0x3e, 0xdd, 0x34, 0xc8, 0x57, 0x10, 0x97, 0x7b, 0x11, 0xaf, 0xa2, 0x8b, 0xc5, 0x8b, 0xe5,
	0x9a, 0xd2, 0xbe, 0xd5, 0x06, 0xdf, 0xd1, 0x66, 0x19, 0x97, 0x7b, 0xfe, 0x0c, 0xe2, 0xa6, 0x15,
	0x8c, 0x22, 0x1e, 0xf4, 0x11, 0x1f, 0x1d, 0xb6, 0xfa, 0xaa, 0x96, 0xd8, 0x76, 0xc6, 0xcb, 0xb8,
	0x69, 0xb9, 0x80, 0xa8, 0x14, 0x13, 0x8a, 0x81, 0x63, 0x16, 0x19, 0x95, 0xd9, 0x8f, 0x08, 0xe0,
	0x98, 0x91, 0x3f, 0x05, 0x56, 0xf8, 0x6b, 0x11, 0x51, 0xe8, 0xd9, 0x7a, 0x68, 0xeb, 0x8d, 0xad,
	0x3d, 0x5e, 0x7b, 0x19, 0xd6, 0xf8, 0x63, 0x80, 0xbc, 0x28, 0x6c, 0x57, 0xfb, 0xad, 0x56, 0xd4,
	0x5a, 0x2a, 0xd3, 0x41, 0x79, 0xaf, 0x38, 0x87, 0x49, 0x9d, 0x57, 0x48, 0x1d, 0xa5, 0x92, 0xec,
	0xa0, 0xf9, 0x9b, 0x06, 0xa9, 0x83, 0x54, 0x92, 0x1d, 0xb4, 0x56, 0xdf, 0xa2, 0x98, 0xae, 0xa2,
	0x0b, 0x26, 0xc9, 0xe6, 0x4b, 0x60, 0x95, 0x7a, 0x25, 0x66, 0x14, 0x16, 0x4c, 0xbe, 0x82, 0x85,
	0xc2, 0xb6, 0x70, 0xba, 0xf1, 0xda, 0xd6, 0x62, 0x4e, 0x2b, 0xa7, 0x52, 0xf6, 0x3d, 0x82, 0xff,
	0xee, 0x0c, 0x7c, 0xff, 0x19, 0xd8, 0xef, 0x33, 0x2c, 0x81, 0x75, 0xce, 0x0c, 0xed, 0x06, 0x93,
	0xff, 0x0f, 0xb1, 0x56, 0x43, 0x63, 0xb1, 0x56, 0x21, 0x41, 0xa8, 0x88, 0x6a, 0x1b, 0x02, 0xa7,
	0x7d, 0x82, 0x5e, 0xb9, 0x74, 0x26, 0xdb, 0xc2, 0x82, 0x00, 0xe3, 0x97, 0x0e, 0x5b, 0xff, 0x17,
	0xa8, 0xf6, 0xf5, 0xd9, 0x58, 0x3f, 0xfb, 0x16, 0xc3, 0x24, 0x54, 0xf8, 0xd7, 0x0f, 0x0c, 0xfe,
	0x38, 0x30, 0x2e, 0x60, 0x5e, 0x38, 0xcc, 0x3d, 0x2a, 0x3a, 0x4e, 0x26, 0x47, 0x77, 0xc4, 0x9e,
	0x1c, 0xb1, 0x8f, 0xb1, 0xd6, 0x89, 0x94, 0xd4, 0xd1, 0x1d, 0x80, 0x2c, 0x0e, 0x40, 0x7e, 0x46,
	0x90, 0x04, 0x20, 0x1f, 0xd0, 0xe7, 0xf7, 0x84, 0xc9, 0xcf, 0x21, 0xe9, 0x1a, 0x63, 0x73, 0x85,
	0x8a, 0x26, 0x66, 0xf2, 0xe0, 0x1f, 0xe8, 0x4c, 0xef, 0xd2, 0x21, 0x12, 0xb3, 0x13, 0x12, 0x0f,
	0x21, 0x29, 0x3b, 0x63, 0xe8, 0x3a, 0xf4, 0xb7, 0x74, 0x1e, 0xfc, 0x4b, 0x67, 0xf8, 0x23, 0x48,
	0x2b, 0x5d, 0xe1, 0x96, 0x88, 0xf6, 0xc3, 0x25, 0x41, 0xa0, 0x0f, 0xfc, 0x09, 0x80, 0xb2, 0x5f,
	0xeb, 0xa1, 0x3a, 0x50, 0xc6, 0x13, 0x65, 0x37, 0xa3, 0x5f, 0xc2, 0xcb, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x97, 0xa5, 0xfe, 0x05, 0x58, 0x04, 0x00, 0x00,
}
