// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file/file.proto

/*
Package file is a generated protocol buffer package.

It is generated from these files:
	file/file.proto

It has these top-level messages:
	FileHeader
	PresignResult
	File
*/
package file

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bitbucket.org/subiz/header/common"

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
func (*FileHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
func (*PresignResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
	proto.RegisterType((*FileHeader)(nil), "file.FileHeader")
	proto.RegisterType((*PresignResult)(nil), "file.PresignResult")
	proto.RegisterType((*File)(nil), "file.File")
	proto.RegisterEnum("file.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("file/file.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0xcf, 0xae, 0x93, 0x40,
	0x18, 0xc5, 0x85, 0xa1, 0xbd, 0xf6, 0x43, 0x2d, 0x4e, 0x34, 0x99, 0x98, 0x98, 0x60, 0x57, 0x8d,
	0x0b, 0x48, 0x34, 0x3e, 0xc1, 0x8d, 0xc6, 0xee, 0x0c, 0x49, 0x57, 0x2e, 0x1a, 0x60, 0xbe, 0xd6,
	0x89, 0x30, 0x83, 0xc3, 0x60, 0x6a, 0x1f, 0xc2, 0x85, 0x0f, 0xe4, 0xb3, 0x99, 0x6f, 0x00, 0x6b,
	0x74, 0xe5, 0xee, 0x6e, 0xe0, 0xcc, 0x39, 0x87, 0xf9, 0xf3, 0x63, 0x60, 0x7d, 0x54, 0x0d, 0xe6,
	0xf4, 0xc8, 0x3a, 0x6b, 0x9c, 0xe1, 0x11, 0xe9, 0x67, 0x59, 0xa5, 0x5c, 0x35, 0xd4, 0x9f, 0xd1,
	0x65, 0xc6, 0x9e, 0xf2, 0x7e, 0xa8, 0xd4, 0x25, 0xff, 0x84, 0xa5, 0x44, 0x9b, 0xd7, 0xa6, 0x6d,
	0x8d, 0x9e, 0x5e, 0xe3, 0x57, 0x9b, 0x9f, 0x01, 0xc0, 0x3b, 0xd5, 0xe0, 0x7b, 0x5f, 0xe1, 0x2f,
	0x80, 0xd5, 0xee, 0x2c, 0x82, 0x34, 0xd8, 0xc6, 0xaf, 0xd6, 0xd9, 0x54, 0xbd, 0x35, 0xda, 0xe1,
	0xd9, 0x15, 0x94, 0xf1, 0xe7, 0x00, 0x65, 0x5d, 0x9b, 0x41, 0xbb, 0x83, 0x92, 0x22, 0x4c, 0x83,
	0xed, 0xaa, 0x58, 0x4d, 0xce, 0x4e, 0x72, 0x0e, 0x91, 0x2e, 0x5b, 0x14, 0xcc, 0x07, 0x5e, 0x93,
	0xe7, 0xbe, 0x75, 0x28, 0xa2, 0xd1, 0x23, 0x4d, 0x5e, 0xaf, 0x2e, 0x28, 0x16, 0x69, 0xb0, 0x65,
	0x85, 0xd7, 0x3c, 0x01, 0xd6, 0xca, 0x37, 0x62, 0xe9, 0x6b, 0x24, 0x79, 0x0a, 0xb1, 0xc4, 0xbe,
	0xb6, 0xaa, 0x73, 0xca, 0x68, 0x71, 0xe3, 0x93, 0x3f, 0xad, 0xcd, 0x8f, 0x00, 0x1e, 0x7e, 0xb0,
	0xd8, 0xab, 0x93, 0x2e, 0xb0, 0x1f, 0x1a, 0xf7, 0xff, 0x67, 0x60, 0x7f, 0x9f, 0x21, 0x01, 0x36,
	0xd8, 0x66, 0xda, 0x2e, 0x49, 0xfe, 0x08, 0x42, 0x25, 0xa7, 0x8d, 0x85, 0x4a, 0xd2, 0x04, 0xb4,
	0x22, 0xca, 0x03, 0x15, 0x17, 0xe3, 0x04, 0xa3, 0xb3, 0xb7, 0xcd, 0xe6, 0x7b, 0x08, 0x11, 0x51,
	0xbd, 0xeb, 0x3c, 0xe1, 0x1f, 0x9e, 0x5c, 0xc0, 0x4d, 0x6d, 0xb1, 0x74, 0x28, 0x3d, 0x6d, 0x56,
	0xcc, 0xc3, 0x99, 0xca, 0xfd, 0x2b, 0x95, 0xb9, 0x6b, 0xac, 0x58, 0x79, 0x77, 0x1e, 0x4e, 0xbc,
	0xe2, 0x99, 0xd7, 0xcb, 0x8f, 0xb0, 0x78, 0xfb, 0x15, 0xb5, 0xe3, 0x02, 0x9e, 0x10, 0x98, 0xdf,
	0x7f, 0xec, 0xcb, 0x80, 0xbd, 0x43, 0x99, 0xdc, 0xe3, 0x6b, 0x88, 0x29, 0xb9, 0x1d, 0x57, 0x4b,
	0x18, 0x7f, 0x0a, 0x8f, 0xc9, 0xd8, 0xe9, 0xa3, 0xb9, 0xf6, 0x22, 0x9e, 0xc0, 0x03, 0xb2, 0xf7,
	0x5d, 0x63, 0x4a, 0x89, 0x32, 0x59, 0x56, 0x4b, 0x7f, 0x95, 0x5f, 0xff, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x8b, 0x69, 0x1b, 0xea, 0x13, 0x03, 0x00, 0x00,
}
