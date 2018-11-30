// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content/content.proto

package content

import (
	fmt "fmt"
	common "git.subiz.net/header/common"
	conversation "git.subiz.net/header/conversation"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Event int32

const (
	Event_ContentRequested Event = 2
	Event_ContentViewed    Event = 3
)

var Event_name = map[int32]string{
	2: "ContentRequested",
	3: "ContentViewed",
}

var Event_value = map[string]int32{
	"ContentRequested": 2,
	"ContentViewed":    3,
}

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}

func (x *Event) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Event_value, data, "Event")
	if err != nil {
		return err
	}
	*x = Event(value)
	return nil
}

func (Event) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{0}
}

type Content_Availability int32

const (
	Content_in_stock     Content_Availability = 1
	Content_out_of_stock Content_Availability = 2
	Content_preorder     Content_Availability = 3
	Content_discontinued Content_Availability = 4
)

var Content_Availability_name = map[int32]string{
	1: "in_stock",
	2: "out_of_stock",
	3: "preorder",
	4: "discontinued",
}

var Content_Availability_value = map[string]int32{
	"in_stock":     1,
	"out_of_stock": 2,
	"preorder":     3,
	"discontinued": 4,
}

func (x Content_Availability) Enum() *Content_Availability {
	p := new(Content_Availability)
	*p = x
	return p
}

func (x Content_Availability) String() string {
	return proto.EnumName(Content_Availability_name, int32(x))
}

func (x *Content_Availability) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Content_Availability_value, data, "Content_Availability")
	if err != nil {
		return err
	}
	*x = Content_Availability(value)
	return nil
}

func (Content_Availability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{3, 0}
}

type LinkRequest struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId            *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Url                  *string         `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Ids                  []string        `protobuf:"bytes,4,rep,name=ids" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LinkRequest) Reset()         { *m = LinkRequest{} }
func (m *LinkRequest) String() string { return proto.CompactTextString(m) }
func (*LinkRequest) ProtoMessage()    {}
func (*LinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{0}
}

func (m *LinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkRequest.Unmarshal(m, b)
}
func (m *LinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkRequest.Marshal(b, m, deterministic)
}
func (m *LinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkRequest.Merge(m, src)
}
func (m *LinkRequest) XXX_Size() int {
	return xxx_messageInfo_LinkRequest.Size(m)
}
func (m *LinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LinkRequest proto.InternalMessageInfo

func (m *LinkRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *LinkRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *LinkRequest) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *LinkRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type SearchContentRequest struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId            *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Id                   *string         `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Title                *string         `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Url                  *string         `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
	Labels               *string         `protobuf:"bytes,6,opt,name=labels" json:"labels,omitempty"`
	Categories           *string         `protobuf:"bytes,10,opt,name=categories" json:"categories,omitempty"`
	Relates              *string         `protobuf:"bytes,11,opt,name=relates" json:"relates,omitempty"`
	Fieldkeys            *string         `protobuf:"bytes,12,opt,name=fieldkeys" json:"fieldkeys,omitempty"`
	Query                *string         `protobuf:"bytes,13,opt,name=query" json:"query,omitempty"`
	Limit                *int32          `protobuf:"varint,14,opt,name=limit" json:"limit,omitempty"`
	Anchor               *string         `protobuf:"bytes,15,opt,name=anchor" json:"anchor,omitempty"`
	Stringify            *string         `protobuf:"bytes,16,opt,name=stringify" json:"stringify,omitempty"`
	Fieldvalues          *string         `protobuf:"bytes,17,opt,name=fieldvalues" json:"fieldvalues,omitempty"`
	Created              *int64          `protobuf:"varint,18,opt,name=created" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SearchContentRequest) Reset()         { *m = SearchContentRequest{} }
func (m *SearchContentRequest) String() string { return proto.CompactTextString(m) }
func (*SearchContentRequest) ProtoMessage()    {}
func (*SearchContentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{1}
}

func (m *SearchContentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchContentRequest.Unmarshal(m, b)
}
func (m *SearchContentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchContentRequest.Marshal(b, m, deterministic)
}
func (m *SearchContentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchContentRequest.Merge(m, src)
}
func (m *SearchContentRequest) XXX_Size() int {
	return xxx_messageInfo_SearchContentRequest.Size(m)
}
func (m *SearchContentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchContentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchContentRequest proto.InternalMessageInfo

func (m *SearchContentRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *SearchContentRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *SearchContentRequest) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *SearchContentRequest) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *SearchContentRequest) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *SearchContentRequest) GetLabels() string {
	if m != nil && m.Labels != nil {
		return *m.Labels
	}
	return ""
}

func (m *SearchContentRequest) GetCategories() string {
	if m != nil && m.Categories != nil {
		return *m.Categories
	}
	return ""
}

func (m *SearchContentRequest) GetRelates() string {
	if m != nil && m.Relates != nil {
		return *m.Relates
	}
	return ""
}

func (m *SearchContentRequest) GetFieldkeys() string {
	if m != nil && m.Fieldkeys != nil {
		return *m.Fieldkeys
	}
	return ""
}

func (m *SearchContentRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *SearchContentRequest) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *SearchContentRequest) GetAnchor() string {
	if m != nil && m.Anchor != nil {
		return *m.Anchor
	}
	return ""
}

func (m *SearchContentRequest) GetStringify() string {
	if m != nil && m.Stringify != nil {
		return *m.Stringify
	}
	return ""
}

func (m *SearchContentRequest) GetFieldvalues() string {
	if m != nil && m.Fieldvalues != nil {
		return *m.Fieldvalues
	}
	return ""
}

func (m *SearchContentRequest) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

type KeyValue struct {
	Key                  *string  `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Value                *string  `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{2}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Content struct {
	Ctx *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	// optional string sbid = 2;
	Id                   *string                    `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	AccountId            *string                    `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Description          *string                    `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Title                *string                    `protobuf:"bytes,6,opt,name=title" json:"title,omitempty"`
	Url                  *string                    `protobuf:"bytes,7,opt,name=url" json:"url,omitempty"`
	Labels               []string                   `protobuf:"bytes,8,rep,name=labels" json:"labels,omitempty"`
	Availability         *string                    `protobuf:"bytes,9,opt,name=availability" json:"availability,omitempty"`
	Price                *float32                   `protobuf:"fixed32,10,opt,name=price" json:"price,omitempty"`
	Currency             *string                    `protobuf:"bytes,11,opt,name=currency" json:"currency,omitempty"`
	SalePrice            *float32                   `protobuf:"fixed32,12,opt,name=sale_price,json=salePrice" json:"sale_price,omitempty"`
	Fields               []*KeyValue                `protobuf:"bytes,13,rep,name=fields" json:"fields,omitempty"`
	Categories           []string                   `protobuf:"bytes,14,rep,name=categories" json:"categories,omitempty"`
	Relates              []string                   `protobuf:"bytes,15,rep,name=relates" json:"relates,omitempty"`
	Attachments          []*conversation.Attachment `protobuf:"bytes,16,rep,name=attachments" json:"attachments,omitempty"`
	Created              *int64                     `protobuf:"varint,17,opt,name=created" json:"created,omitempty"`
	Updated              *int64                     `protobuf:"varint,18,opt,name=updated" json:"updated,omitempty"`
	AttachmentUrls       []string                   `protobuf:"bytes,20,rep,name=attachment_urls,json=attachmentUrls" json:"attachment_urls,omitempty"`
	RelatedIds           []string                   `protobuf:"bytes,21,rep,name=related_ids,json=relatedIds" json:"related_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{3}
}

func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Content) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Content) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Content) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Content) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Content) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *Content) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Content) GetAvailability() string {
	if m != nil && m.Availability != nil {
		return *m.Availability
	}
	return ""
}

func (m *Content) GetPrice() float32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *Content) GetCurrency() string {
	if m != nil && m.Currency != nil {
		return *m.Currency
	}
	return ""
}

func (m *Content) GetSalePrice() float32 {
	if m != nil && m.SalePrice != nil {
		return *m.SalePrice
	}
	return 0
}

func (m *Content) GetFields() []*KeyValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Content) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Content) GetRelates() []string {
	if m != nil {
		return m.Relates
	}
	return nil
}

func (m *Content) GetAttachments() []*conversation.Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Content) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Content) GetUpdated() int64 {
	if m != nil && m.Updated != nil {
		return *m.Updated
	}
	return 0
}

func (m *Content) GetAttachmentUrls() []string {
	if m != nil {
		return m.AttachmentUrls
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
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Contents             []*Content      `protobuf:"bytes,2,rep,name=contents" json:"contents,omitempty"`
	Anchor               *string         `protobuf:"bytes,3,opt,name=anchor" json:"anchor,omitempty"`
	Total                *int64          `protobuf:"varint,4,opt,name=total" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Contents) Reset()         { *m = Contents{} }
func (m *Contents) String() string { return proto.CompactTextString(m) }
func (*Contents) ProtoMessage()    {}
func (*Contents) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{4}
}

func (m *Contents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contents.Unmarshal(m, b)
}
func (m *Contents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contents.Marshal(b, m, deterministic)
}
func (m *Contents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contents.Merge(m, src)
}
func (m *Contents) XXX_Size() int {
	return xxx_messageInfo_Contents.Size(m)
}
func (m *Contents) XXX_DiscardUnknown() {
	xxx_messageInfo_Contents.DiscardUnknown(m)
}

var xxx_messageInfo_Contents proto.InternalMessageInfo

func (m *Contents) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Contents) GetContents() []*Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *Contents) GetAnchor() string {
	if m != nil && m.Anchor != nil {
		return *m.Anchor
	}
	return ""
}

func (m *Contents) GetTotal() int64 {
	if m != nil && m.Total != nil {
		return *m.Total
	}
	return 0
}

type ContentUrl struct {
	AccountId            *string  `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Par                  *int32   `protobuf:"varint,3,opt,name=par" json:"par,omitempty"`
	Url                  *string  `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
	Err                  *string  `protobuf:"bytes,5,opt,name=err" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentUrl) Reset()         { *m = ContentUrl{} }
func (m *ContentUrl) String() string { return proto.CompactTextString(m) }
func (*ContentUrl) ProtoMessage()    {}
func (*ContentUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ed60e503df5706, []int{5}
}

func (m *ContentUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentUrl.Unmarshal(m, b)
}
func (m *ContentUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentUrl.Marshal(b, m, deterministic)
}
func (m *ContentUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentUrl.Merge(m, src)
}
func (m *ContentUrl) XXX_Size() int {
	return xxx_messageInfo_ContentUrl.Size(m)
}
func (m *ContentUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentUrl.DiscardUnknown(m)
}

var xxx_messageInfo_ContentUrl proto.InternalMessageInfo

func (m *ContentUrl) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *ContentUrl) GetPar() int32 {
	if m != nil && m.Par != nil {
		return *m.Par
	}
	return 0
}

func (m *ContentUrl) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *ContentUrl) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

func init() {
	proto.RegisterEnum("content.Event", Event_name, Event_value)
	proto.RegisterEnum("content.Content_Availability", Content_Availability_name, Content_Availability_value)
	proto.RegisterType((*LinkRequest)(nil), "content.LinkRequest")
	proto.RegisterType((*SearchContentRequest)(nil), "content.SearchContentRequest")
	proto.RegisterType((*KeyValue)(nil), "content.KeyValue")
	proto.RegisterType((*Content)(nil), "content.Content")
	proto.RegisterType((*Contents)(nil), "content.Contents")
	proto.RegisterType((*ContentUrl)(nil), "content.ContentUrl")
}

func init() { proto.RegisterFile("content/content.proto", fileDescriptor_d1ed60e503df5706) }

var fileDescriptor_d1ed60e503df5706 = []byte{
	// 889 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xe3, 0x36,
	0x10, 0x85, 0x2d, 0x3b, 0xb1, 0xc7, 0x1f, 0x91, 0xb9, 0x4e, 0x41, 0x18, 0xdd, 0xd6, 0xf5, 0xa5,
	0xde, 0xa0, 0x70, 0x0a, 0xb7, 0x87, 0xa2, 0xb7, 0xdd, 0xed, 0xa2, 0x08, 0xba, 0x2d, 0x16, 0x4a,
	0xb3, 0x57, 0x83, 0x11, 0x27, 0x0e, 0x61, 0x59, 0xf2, 0x92, 0x94, 0xbb, 0xee, 0xad, 0xc7, 0xfe,
	0x89, 0x02, 0xfd, 0xa7, 0x05, 0x29, 0x52, 0x96, 0x9d, 0x2c, 0x90, 0x43, 0x4f, 0xd2, 0xbc, 0x19,
	0xce, 0x8c, 0xe6, 0x3d, 0x0d, 0xe1, 0x3c, 0xce, 0x52, 0x8d, 0xa9, 0xbe, 0x74, 0xcf, 0xd9, 0x46,
	0x66, 0x3a, 0x23, 0xa7, 0xce, 0x1c, 0x4d, 0x97, 0x42, 0xcf, 0x54, 0x7e, 0x2b, 0xfe, 0x9c, 0xa5,
	0xa8, 0x2f, 0xef, 0x91, 0x71, 0x94, 0x97, 0x71, 0xb6, 0x5e, 0x67, 0xa9, 0x7b, 0x14, 0x47, 0x46,
	0xdf, 0x7f, 0x22, 0x32, 0xdd, 0xa2, 0x54, 0x4c, 0x0b, 0x1b, 0xbf, 0x37, 0x8a, 0x53, 0x13, 0x05,
	0x9d, 0xb7, 0x22, 0x5d, 0x45, 0xf8, 0x21, 0x47, 0xa5, 0xc9, 0x57, 0x10, 0xc4, 0xfa, 0x23, 0xad,
	0x8d, 0x6b, 0xd3, 0xce, 0xfc, 0x6c, 0xe6, 0x0a, 0xbc, 0x36, 0xcd, 0x7c, 0xd4, 0x91, 0xf1, 0x91,
	0xe7, 0x00, 0x2c, 0x8e, 0xb3, 0x3c, 0xd5, 0x0b, 0xc1, 0x69, 0x7d, 0x5c, 0x9b, 0xb6, 0xa3, 0xb6,
	0x43, 0xae, 0x38, 0x09, 0x21, 0xc8, 0x65, 0x42, 0x03, 0x8b, 0x9b, 0x57, 0x83, 0x08, 0xae, 0x68,
	0x63, 0x1c, 0x18, 0x44, 0x70, 0x35, 0xf9, 0x27, 0x80, 0xe1, 0x35, 0x32, 0x19, 0xdf, 0xbf, 0x2e,
	0x3e, 0xf3, 0xff, 0x2b, 0xdf, 0x87, 0xba, 0xe0, 0xae, 0x7a, 0x5d, 0x70, 0x32, 0x84, 0xa6, 0x16,
	0x3a, 0x41, 0xda, 0xb0, 0x50, 0x61, 0xf8, 0x26, 0x9b, 0xfb, 0x26, 0x3f, 0x83, 0x93, 0x84, 0xdd,
	0x62, 0xa2, 0xe8, 0x89, 0x05, 0x9d, 0x45, 0xbe, 0x00, 0x88, 0x99, 0xc6, 0x65, 0x26, 0x05, 0x2a,
	0x0a, 0xd6, 0x57, 0x41, 0x08, 0x85, 0x53, 0x89, 0x09, 0xd3, 0xa8, 0x68, 0xc7, 0x3a, 0xbd, 0x49,
	0x3e, 0x87, 0xf6, 0x9d, 0xc0, 0x84, 0xaf, 0x70, 0xa7, 0x68, 0xb7, 0xe8, 0xb3, 0x04, 0x4c, 0x5f,
	0x1f, 0x72, 0x94, 0x3b, 0xda, 0x2b, 0xfa, 0xb2, 0x86, 0x41, 0x13, 0xb1, 0x16, 0x9a, 0xf6, 0xc7,
	0xb5, 0x69, 0x33, 0x2a, 0x0c, 0xd3, 0x1b, 0x4b, 0xe3, 0xfb, 0x4c, 0xd2, 0xb3, 0xa2, 0xb7, 0xc2,
	0x32, 0x15, 0x94, 0x96, 0x22, 0x5d, 0x8a, 0xbb, 0x1d, 0x0d, 0x8b, 0x0a, 0x25, 0x40, 0xc6, 0xd0,
	0xb1, 0xe5, 0xb6, 0x2c, 0xc9, 0x51, 0xd1, 0x81, 0xf5, 0x57, 0x21, 0xd3, 0x7b, 0x2c, 0x91, 0x69,
	0xe4, 0x94, 0x8c, 0x6b, 0xd3, 0x20, 0xf2, 0xe6, 0x64, 0x0e, 0xad, 0x5f, 0x70, 0xf7, 0xde, 0x84,
	0x99, 0x59, 0xad, 0x70, 0xe7, 0x09, 0x5d, 0xa1, 0xed, 0xd2, 0x66, 0xf0, 0x33, 0xb5, 0xc6, 0xe4,
	0xdf, 0x26, 0x9c, 0x3a, 0x3a, 0x9f, 0xc2, 0xe3, 0x31, 0x51, 0x87, 0xbc, 0x36, 0x8e, 0x79, 0x1d,
	0x43, 0x87, 0xa3, 0x8a, 0xa5, 0xd8, 0x18, 0xf1, 0x3a, 0xe6, 0xaa, 0xd0, 0x9e, 0xe9, 0x93, 0x47,
	0x98, 0x3e, 0x7d, 0x8c, 0xe9, 0x96, 0x55, 0xa4, 0x67, 0x7a, 0x02, 0x5d, 0xb6, 0x65, 0x22, 0x61,
	0xb7, 0x22, 0x11, 0x7a, 0x47, 0xdb, 0xf6, 0xc8, 0x01, 0x66, 0x6a, 0x6c, 0xa4, 0x88, 0xd1, 0x0a,
	0xa1, 0x1e, 0x15, 0x06, 0x19, 0x41, 0x2b, 0xce, 0xa5, 0xc4, 0x34, 0xde, 0x39, 0x11, 0x94, 0xb6,
	0xf9, 0x2c, 0xc5, 0x12, 0x5c, 0x14, 0xc7, 0xba, 0xf6, 0x58, 0xdb, 0x20, 0xef, 0xec, 0xd1, 0x17,
	0x70, 0x62, 0x19, 0x51, 0xb4, 0x37, 0x0e, 0xa6, 0x9d, 0xf9, 0x60, 0xe6, 0xf7, 0x80, 0x9f, 0x7f,
	0xe4, 0x02, 0x8e, 0x94, 0xd8, 0xb7, 0xbd, 0x7f, 0x42, 0x89, 0x67, 0xd6, 0x59, 0x2a, 0xf1, 0x47,
	0xe8, 0x30, 0xad, 0x59, 0x7c, 0xbf, 0xc6, 0x54, 0x2b, 0x1a, 0xda, 0x4a, 0x74, 0x76, 0xb0, 0x0d,
	0x5e, 0x96, 0x01, 0x51, 0x35, 0xb8, 0xaa, 0x91, 0xc1, 0x81, 0x46, 0x8c, 0x27, 0xdf, 0xf0, 0xaa,
	0x7a, 0x9c, 0x49, 0xbe, 0x86, 0xb3, 0x7d, 0x8a, 0x45, 0x2e, 0x13, 0x45, 0x87, 0xb6, 0xa3, 0xfe,
	0x1e, 0xbe, 0x91, 0x89, 0x22, 0x5f, 0x42, 0xa7, 0xe8, 0x91, 0x2f, 0xcc, 0x86, 0x38, 0x2f, 0xbe,
	0xc9, 0x41, 0x57, 0x5c, 0x4d, 0x7e, 0x83, 0xee, 0xcb, 0xea, 0xfc, 0xbb, 0xd0, 0x12, 0xe9, 0x42,
	0xe9, 0x2c, 0x5e, 0x85, 0x35, 0x12, 0x42, 0x37, 0xcb, 0xf5, 0x22, 0xbb, 0x73, 0x48, 0xdd, 0xf8,
	0x37, 0x12, 0x33, 0xc9, 0x51, 0x86, 0x81, 0xf1, 0x73, 0xa1, 0xcc, 0x40, 0x45, 0x9a, 0x23, 0x0f,
	0x1b, 0x93, 0xbf, 0x6b, 0xd0, 0x72, 0x1a, 0x55, 0x4f, 0x11, 0xe9, 0x37, 0xd0, 0x72, 0x7c, 0x28,
	0x5a, 0xb7, 0x63, 0x0b, 0x4b, 0x82, 0xfc, 0xea, 0x2a, 0x23, 0x2a, 0xff, 0x69, 0x70, 0xf0, 0x9f,
	0x1a, 0x65, 0x66, 0x9a, 0x25, 0x56, 0xd5, 0x41, 0x54, 0x18, 0x13, 0x06, 0xe0, 0x52, 0xdc, 0xc8,
	0xe4, 0x09, 0x5b, 0x75, 0xc3, 0x8a, 0xbc, 0xcd, 0xc8, 0xbc, 0x7a, 0x61, 0x37, 0x0e, 0xf6, 0x2c,
	0x4a, 0xe9, 0x97, 0x1a, 0x4a, 0x79, 0xf1, 0x2d, 0x34, 0xdf, 0x6c, 0xcd, 0xff, 0x38, 0x84, 0xf0,
	0x70, 0xd3, 0x22, 0x0f, 0xeb, 0x64, 0x00, 0x3d, 0x87, 0xbe, 0x17, 0xf8, 0x07, 0xf2, 0x30, 0x98,
	0xff, 0x15, 0xc0, 0xf9, 0x35, 0xca, 0xad, 0x88, 0xd1, 0xb9, 0x7e, 0x65, 0x29, 0x5b, 0xa2, 0x24,
	0x2f, 0xa0, 0xf1, 0x4e, 0xa4, 0x4b, 0xf2, 0xcc, 0x0f, 0xca, 0x58, 0x2e, 0xdb, 0xa8, 0x5b, 0x82,
	0x59, 0xba, 0x24, 0xdf, 0x41, 0xef, 0x66, 0xa3, 0x50, 0x6a, 0xbf, 0x0e, 0x1e, 0x0c, 0x6d, 0xf4,
	0x00, 0x21, 0x17, 0xd0, 0xfb, 0x09, 0x13, 0xd4, 0xbe, 0x2e, 0x01, 0x9f, 0xf3, 0x8a, 0x8f, 0x7a,
	0xfe, 0xfd, 0xcd, 0x7a, 0xa3, 0x77, 0xe4, 0x02, 0xe0, 0x67, 0xd4, 0x8f, 0x05, 0x3e, 0xcc, 0xfb,
	0x0a, 0xfa, 0x07, 0x57, 0x8d, 0x22, 0xcf, 0xcb, 0x98, 0xc7, 0xee, 0xa0, 0xd1, 0xe0, 0x38, 0x85,
	0x22, 0x73, 0x78, 0xf6, 0x36, 0xcb, 0x56, 0xf9, 0xc6, 0x23, 0xaf, 0x76, 0x86, 0xb3, 0xbe, 0x2f,
	0x7c, 0x6d, 0x97, 0xee, 0xa8, 0xb3, 0x6f, 0x44, 0x91, 0x1f, 0x60, 0x60, 0x2e, 0x56, 0x7f, 0xe2,
	0xf7, 0xcc, 0x9c, 0x18, 0x96, 0xb9, 0x2b, 0x97, 0xee, 0xd1, 0xd7, 0xfd, 0x17, 0x00, 0x00, 0xff,
	0xff, 0xe9, 0xf6, 0x8d, 0x09, 0x13, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceContentManagerClient is the client API for ServiceContentManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceContentManagerClient interface {
	Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error)
	UpsertContent(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Content, error)
	DeleteContent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*common.Empty, error)
	GetContent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*Content, error)
	SearchContents(ctx context.Context, in *SearchContentRequest, opts ...grpc.CallOption) (*Contents, error)
	LookupContentsByUrl(ctx context.Context, in *common.String, opts ...grpc.CallOption) (*common.Ids, error)
	LinkContentsToUrl(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*common.Empty, error)
}

type serviceContentManagerClient struct {
	cc *grpc.ClientConn
}

func NewServiceContentManagerClient(cc *grpc.ClientConn) ServiceContentManagerClient {
	return &serviceContentManagerClient{cc}
}

func (c *serviceContentManagerClient) Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error) {
	out := new(common.Pong)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) UpsertContent(ctx context.Context, in *Content, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/UpsertContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) DeleteContent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/DeleteContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) GetContent(ctx context.Context, in *common.Id, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/GetContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) SearchContents(ctx context.Context, in *SearchContentRequest, opts ...grpc.CallOption) (*Contents, error) {
	out := new(Contents)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/SearchContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) LookupContentsByUrl(ctx context.Context, in *common.String, opts ...grpc.CallOption) (*common.Ids, error) {
	out := new(common.Ids)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/LookupContentsByUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceContentManagerClient) LinkContentsToUrl(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/content.ServiceContentManager/LinkContentsToUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceContentManagerServer is the server API for ServiceContentManager service.
type ServiceContentManagerServer interface {
	Ping(context.Context, *common.PingRequest) (*common.Pong, error)
	UpsertContent(context.Context, *Content) (*Content, error)
	DeleteContent(context.Context, *common.Id) (*common.Empty, error)
	GetContent(context.Context, *common.Id) (*Content, error)
	SearchContents(context.Context, *SearchContentRequest) (*Contents, error)
	LookupContentsByUrl(context.Context, *common.String) (*common.Ids, error)
	LinkContentsToUrl(context.Context, *LinkRequest) (*common.Empty, error)
}

func RegisterServiceContentManagerServer(s *grpc.Server, srv ServiceContentManagerServer) {
	s.RegisterService(&_ServiceContentManager_serviceDesc, srv)
}

func _ServiceContentManager_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).Ping(ctx, req.(*common.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_UpsertContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).UpsertContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/UpsertContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).UpsertContent(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_DeleteContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).DeleteContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/DeleteContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).DeleteContent(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_GetContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).GetContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/GetContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).GetContent(ctx, req.(*common.Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_SearchContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).SearchContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/SearchContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).SearchContents(ctx, req.(*SearchContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_LookupContentsByUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).LookupContentsByUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/LookupContentsByUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).LookupContentsByUrl(ctx, req.(*common.String))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceContentManager_LinkContentsToUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceContentManagerServer).LinkContentsToUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/content.ServiceContentManager/LinkContentsToUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceContentManagerServer).LinkContentsToUrl(ctx, req.(*LinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceContentManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "content.ServiceContentManager",
	HandlerType: (*ServiceContentManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ServiceContentManager_Ping_Handler,
		},
		{
			MethodName: "UpsertContent",
			Handler:    _ServiceContentManager_UpsertContent_Handler,
		},
		{
			MethodName: "DeleteContent",
			Handler:    _ServiceContentManager_DeleteContent_Handler,
		},
		{
			MethodName: "GetContent",
			Handler:    _ServiceContentManager_GetContent_Handler,
		},
		{
			MethodName: "SearchContents",
			Handler:    _ServiceContentManager_SearchContents_Handler,
		},
		{
			MethodName: "LookupContentsByUrl",
			Handler:    _ServiceContentManager_LookupContentsByUrl_Handler,
		},
		{
			MethodName: "LinkContentsToUrl",
			Handler:    _ServiceContentManager_LinkContentsToUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content/content.proto",
}
