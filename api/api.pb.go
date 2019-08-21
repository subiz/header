// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type Request struct {
	Header               map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Method               string            `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Path                 string            `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Response struct {
	Header               map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Code                 int32             `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type BatchOp struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Method               string          `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Url                  string          `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Body                 string          `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BatchOp) Reset()         { *m = BatchOp{} }
func (m *BatchOp) String() string { return proto.CompactTextString(m) }
func (*BatchOp) ProtoMessage()    {}
func (*BatchOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *BatchOp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchOp.Unmarshal(m, b)
}
func (m *BatchOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchOp.Marshal(b, m, deterministic)
}
func (m *BatchOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchOp.Merge(m, src)
}
func (m *BatchOp) XXX_Size() int {
	return xxx_messageInfo_BatchOp.Size(m)
}
func (m *BatchOp) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchOp.DiscardUnknown(m)
}

var xxx_messageInfo_BatchOp proto.InternalMessageInfo

func (m *BatchOp) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *BatchOp) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *BatchOp) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *BatchOp) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type BatchRequest struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	Paths                []string        `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
	Ops                  []*BatchOp      `protobuf:"bytes,3,rep,name=ops,proto3" json:"ops,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BatchRequest) Reset()         { *m = BatchRequest{} }
func (m *BatchRequest) String() string { return proto.CompactTextString(m) }
func (*BatchRequest) ProtoMessage()    {}
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *BatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchRequest.Unmarshal(m, b)
}
func (m *BatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchRequest.Marshal(b, m, deterministic)
}
func (m *BatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchRequest.Merge(m, src)
}
func (m *BatchRequest) XXX_Size() int {
	return xxx_messageInfo_BatchRequest.Size(m)
}
func (m *BatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchRequest proto.InternalMessageInfo

func (m *BatchRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *BatchRequest) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *BatchRequest) GetOps() []*BatchOp {
	if m != nil {
		return m.Ops
	}
	return nil
}

type WhitelistUrl struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Url                  string          `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Created              int64           `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	By                   string          `protobuf:"bytes,6,opt,name=by,proto3" json:"by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WhitelistUrl) Reset()         { *m = WhitelistUrl{} }
func (m *WhitelistUrl) String() string { return proto.CompactTextString(m) }
func (*WhitelistUrl) ProtoMessage()    {}
func (*WhitelistUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *WhitelistUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhitelistUrl.Unmarshal(m, b)
}
func (m *WhitelistUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhitelistUrl.Marshal(b, m, deterministic)
}
func (m *WhitelistUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhitelistUrl.Merge(m, src)
}
func (m *WhitelistUrl) XXX_Size() int {
	return xxx_messageInfo_WhitelistUrl.Size(m)
}
func (m *WhitelistUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_WhitelistUrl.DiscardUnknown(m)
}

var xxx_messageInfo_WhitelistUrl proto.InternalMessageInfo

func (m *WhitelistUrl) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *WhitelistUrl) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *WhitelistUrl) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *WhitelistUrl) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *WhitelistUrl) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

// deprecated
type BlacklistDomain struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Domain               string          `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	Created              int64           `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	By                   string          `protobuf:"bytes,6,opt,name=by,proto3" json:"by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BlacklistDomain) Reset()         { *m = BlacklistDomain{} }
func (m *BlacklistDomain) String() string { return proto.CompactTextString(m) }
func (*BlacklistDomain) ProtoMessage()    {}
func (*BlacklistDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *BlacklistDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlacklistDomain.Unmarshal(m, b)
}
func (m *BlacklistDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlacklistDomain.Marshal(b, m, deterministic)
}
func (m *BlacklistDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlacklistDomain.Merge(m, src)
}
func (m *BlacklistDomain) XXX_Size() int {
	return xxx_messageInfo_BlacklistDomain.Size(m)
}
func (m *BlacklistDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_BlacklistDomain.DiscardUnknown(m)
}

var xxx_messageInfo_BlacklistDomain proto.InternalMessageInfo

func (m *BlacklistDomain) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *BlacklistDomain) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *BlacklistDomain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *BlacklistDomain) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *BlacklistDomain) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

type WhitelistDomain struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Domain               string          `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	Created              int64           `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	By                   string          `protobuf:"bytes,6,opt,name=by,proto3" json:"by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WhitelistDomain) Reset()         { *m = WhitelistDomain{} }
func (m *WhitelistDomain) String() string { return proto.CompactTextString(m) }
func (*WhitelistDomain) ProtoMessage()    {}
func (*WhitelistDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *WhitelistDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhitelistDomain.Unmarshal(m, b)
}
func (m *WhitelistDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhitelistDomain.Marshal(b, m, deterministic)
}
func (m *WhitelistDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhitelistDomain.Merge(m, src)
}
func (m *WhitelistDomain) XXX_Size() int {
	return xxx_messageInfo_WhitelistDomain.Size(m)
}
func (m *WhitelistDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_WhitelistDomain.DiscardUnknown(m)
}

var xxx_messageInfo_WhitelistDomain proto.InternalMessageInfo

func (m *WhitelistDomain) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *WhitelistDomain) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *WhitelistDomain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *WhitelistDomain) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *WhitelistDomain) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

type BlacklistIP struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Ip                   string          `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Created              int64           `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	By                   string          `protobuf:"bytes,6,opt,name=by,proto3" json:"by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BlacklistIP) Reset()         { *m = BlacklistIP{} }
func (m *BlacklistIP) String() string { return proto.CompactTextString(m) }
func (*BlacklistIP) ProtoMessage()    {}
func (*BlacklistIP) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *BlacklistIP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlacklistIP.Unmarshal(m, b)
}
func (m *BlacklistIP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlacklistIP.Marshal(b, m, deterministic)
}
func (m *BlacklistIP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlacklistIP.Merge(m, src)
}
func (m *BlacklistIP) XXX_Size() int {
	return xxx_messageInfo_BlacklistIP.Size(m)
}
func (m *BlacklistIP) XXX_DiscardUnknown() {
	xxx_messageInfo_BlacklistIP.DiscardUnknown(m)
}

var xxx_messageInfo_BlacklistIP proto.InternalMessageInfo

func (m *BlacklistIP) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *BlacklistIP) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *BlacklistIP) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *BlacklistIP) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *BlacklistIP) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

type BannedUser struct {
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId            string          `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	UserId               string          `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Created              int64           `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	By                   string          `protobuf:"bytes,6,opt,name=by,proto3" json:"by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BannedUser) Reset()         { *m = BannedUser{} }
func (m *BannedUser) String() string { return proto.CompactTextString(m) }
func (*BannedUser) ProtoMessage()    {}
func (*BannedUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *BannedUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BannedUser.Unmarshal(m, b)
}
func (m *BannedUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BannedUser.Marshal(b, m, deterministic)
}
func (m *BannedUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannedUser.Merge(m, src)
}
func (m *BannedUser) XXX_Size() int {
	return xxx_messageInfo_BannedUser.Size(m)
}
func (m *BannedUser) XXX_DiscardUnknown() {
	xxx_messageInfo_BannedUser.DiscardUnknown(m)
}

var xxx_messageInfo_BannedUser proto.InternalMessageInfo

func (m *BannedUser) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *BannedUser) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *BannedUser) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *BannedUser) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *BannedUser) GetBy() string {
	if m != nil {
		return m.By
	}
	return ""
}

type ScryptChallenge struct {
	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	AccountId string          `protobuf:"bytes,9,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Id        string          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	N         int32           `protobuf:"varint,3,opt,name=N,proto3" json:"N,omitempty"`
	P         int32           `protobuf:"varint,4,opt,name=P,proto3" json:"P,omitempty"`
	R         int32           `protobuf:"varint,5,opt,name=r,proto3" json:"r,omitempty"`
	Salt      string          `protobuf:"bytes,6,opt,name=salt,proto3" json:"salt,omitempty"`
	Hash      string          `protobuf:"bytes,8,opt,name=hash,proto3" json:"hash,omitempty"`
	Domain    int32           `protobuf:"varint,14,opt,name=domain,proto3" json:"domain,omitempty"`
	Dklen     int32           `protobuf:"varint,15,opt,name=dklen,proto3" json:"dklen,omitempty"`
	// only for subiz use
	Answer               string   `protobuf:"bytes,10,opt,name=answer,proto3" json:"answer,omitempty"`
	Created              int64    `protobuf:"varint,11,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScryptChallenge) Reset()         { *m = ScryptChallenge{} }
func (m *ScryptChallenge) String() string { return proto.CompactTextString(m) }
func (*ScryptChallenge) ProtoMessage()    {}
func (*ScryptChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *ScryptChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScryptChallenge.Unmarshal(m, b)
}
func (m *ScryptChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScryptChallenge.Marshal(b, m, deterministic)
}
func (m *ScryptChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScryptChallenge.Merge(m, src)
}
func (m *ScryptChallenge) XXX_Size() int {
	return xxx_messageInfo_ScryptChallenge.Size(m)
}
func (m *ScryptChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_ScryptChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_ScryptChallenge proto.InternalMessageInfo

func (m *ScryptChallenge) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ScryptChallenge) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *ScryptChallenge) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ScryptChallenge) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *ScryptChallenge) GetP() int32 {
	if m != nil {
		return m.P
	}
	return 0
}

func (m *ScryptChallenge) GetR() int32 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *ScryptChallenge) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *ScryptChallenge) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ScryptChallenge) GetDomain() int32 {
	if m != nil {
		return m.Domain
	}
	return 0
}

func (m *ScryptChallenge) GetDklen() int32 {
	if m != nil {
		return m.Dklen
	}
	return 0
}

func (m *ScryptChallenge) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *ScryptChallenge) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type Apikey struct {
	Secret               string   `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	AccountId            string   `protobuf:"bytes,4,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ClientId             string   `protobuf:"bytes,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientType           string   `protobuf:"bytes,6,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	IsInternal           bool     `protobuf:"varint,7,opt,name=is_internal,json=isInternal,proto3" json:"is_internal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Apikey) Reset()         { *m = Apikey{} }
func (m *Apikey) String() string { return proto.CompactTextString(m) }
func (*Apikey) ProtoMessage()    {}
func (*Apikey) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *Apikey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Apikey.Unmarshal(m, b)
}
func (m *Apikey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Apikey.Marshal(b, m, deterministic)
}
func (m *Apikey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Apikey.Merge(m, src)
}
func (m *Apikey) XXX_Size() int {
	return xxx_messageInfo_Apikey.Size(m)
}
func (m *Apikey) XXX_DiscardUnknown() {
	xxx_messageInfo_Apikey.DiscardUnknown(m)
}

var xxx_messageInfo_Apikey proto.InternalMessageInfo

func (m *Apikey) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Apikey) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Apikey) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Apikey) GetClientType() string {
	if m != nil {
		return m.ClientType
	}
	return ""
}

func (m *Apikey) GetIsInternal() bool {
	if m != nil {
		return m.IsInternal
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "api.Request")
	proto.RegisterMapType((map[string]string)(nil), "api.Request.HeaderEntry")
	proto.RegisterType((*Response)(nil), "api.Response")
	proto.RegisterMapType((map[string]string)(nil), "api.Response.HeaderEntry")
	proto.RegisterType((*BatchOp)(nil), "api.BatchOp")
	proto.RegisterType((*BatchRequest)(nil), "api.BatchRequest")
	proto.RegisterType((*WhitelistUrl)(nil), "api.WhitelistUrl")
	proto.RegisterType((*BlacklistDomain)(nil), "api.BlacklistDomain")
	proto.RegisterType((*WhitelistDomain)(nil), "api.WhitelistDomain")
	proto.RegisterType((*BlacklistIP)(nil), "api.BlacklistIP")
	proto.RegisterType((*BannedUser)(nil), "api.BannedUser")
	proto.RegisterType((*ScryptChallenge)(nil), "api.ScryptChallenge")
	proto.RegisterType((*Apikey)(nil), "api.Apikey")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x4d, 0x6f, 0x13, 0x3b,
	0x14, 0x95, 0x67, 0x92, 0x49, 0xe6, 0x26, 0x6a, 0x2a, 0xab, 0xea, 0xf3, 0x6b, 0xd5, 0xf7, 0xf2,
	0xb2, 0xca, 0x2a, 0x7d, 0x94, 0x0d, 0xb0, 0x23, 0x05, 0x89, 0x6c, 0x4a, 0x65, 0xa8, 0x90, 0xd8,
	0x54, 0xce, 0x8c, 0xd5, 0x31, 0x9d, 0xcc, 0x98, 0xb1, 0x07, 0x3a, 0xfc, 0x03, 0x24, 0x36, 0x2c,
	0xd8, 0x22, 0x7e, 0x05, 0xbf, 0x0f, 0xf9, 0x23, 0x21, 0xad, 0x58, 0x34, 0xa2, 0x0b, 0x76, 0xf7,
	0x9c, 0x7b, 0x3d, 0x3e, 0xf7, 0x5c, 0x8f, 0x0d, 0x31, 0x93, 0x62, 0x22, 0xab, 0x52, 0x97, 0x38,
	0x64, 0x52, 0xec, 0xf5, 0x93, 0x72, 0xb1, 0x28, 0x0b, 0x47, 0x8d, 0xbe, 0x23, 0xe8, 0x50, 0xfe,
	0xb6, 0xe6, 0x4a, 0xe3, 0xff, 0x21, 0xca, 0x38, 0x4b, 0x79, 0x45, 0x82, 0x61, 0x38, 0xee, 0x1d,
	0x91, 0x89, 0x59, 0xea, 0xb3, 0x93, 0x67, 0x36, 0xf5, 0xb4, 0xd0, 0x55, 0x43, 0x7d, 0x1d, 0xc6,
	0xd0, 0x9a, 0x97, 0x69, 0x43, 0xc2, 0x21, 0x1a, 0xf7, 0xa9, 0x8d, 0xf1, 0x2e, 0x44, 0x0b, 0xae,
	0xb3, 0x32, 0x25, 0xad, 0x21, 0x1a, 0xc7, 0xd4, 0x23, 0x53, 0x2b, 0x99, 0xce, 0x48, 0xdb, 0xb2,
	0x36, 0xde, 0x7b, 0x08, 0xbd, 0xb5, 0xcf, 0xe2, 0x6d, 0x08, 0x2f, 0x79, 0x43, 0x90, 0xad, 0x30,
	0x21, 0xde, 0x81, 0xf6, 0x3b, 0x96, 0xd7, 0x9c, 0x04, 0x96, 0x73, 0xe0, 0x51, 0xf0, 0x00, 0x8d,
	0xbe, 0x21, 0xe8, 0x52, 0xae, 0x64, 0x59, 0x28, 0x8e, 0xef, 0xdd, 0x50, 0xfe, 0xb7, 0x57, 0xee,
	0xd2, 0xb7, 0x96, 0x8e, 0xa1, 0x95, 0x94, 0x29, 0xb7, 0xc2, 0xdb, 0xd4, 0xc6, 0xbf, 0x23, 0xf1,
	0x0d, 0x74, 0xa6, 0x4c, 0x27, 0xd9, 0x73, 0x89, 0xff, 0x83, 0x30, 0xd1, 0x57, 0x76, 0x59, 0xef,
	0x68, 0x30, 0xf1, 0x23, 0x38, 0x2e, 0x0b, 0xcd, 0xaf, 0x34, 0x35, 0xb9, 0x35, 0xdf, 0x82, 0x6b,
	0xbe, 0x6d, 0x43, 0x58, 0x57, 0xb9, 0xd5, 0x19, 0x53, 0x13, 0xae, 0xa4, 0x3b, 0x7f, 0x6d, 0x3c,
	0xba, 0x80, 0xbe, 0xdd, 0x6b, 0x39, 0xcb, 0x5b, 0x6c, 0xb8, 0x03, 0x6d, 0x33, 0x04, 0x65, 0x3d,
	0x8b, 0xa9, 0x03, 0xf8, 0x1f, 0x08, 0x4b, 0xa9, 0x48, 0x68, 0x7d, 0xec, 0x5b, 0x1f, 0x7d, 0x13,
	0xd4, 0x24, 0x46, 0x9f, 0x10, 0xf4, 0x5f, 0x65, 0x42, 0xf3, 0x5c, 0x28, 0x7d, 0x56, 0xe5, 0xb7,
	0xd9, 0xe9, 0x00, 0x80, 0x25, 0x49, 0x59, 0x17, 0xfa, 0x5c, 0xa4, 0xbe, 0x93, 0xd8, 0x33, 0xb3,
	0x55, 0x87, 0xad, 0x9f, 0x1d, 0x12, 0xe8, 0x24, 0x15, 0x67, 0x9a, 0xa7, 0xf6, 0xb8, 0x84, 0x74,
	0x09, 0xf1, 0x16, 0x04, 0xf3, 0x86, 0x44, 0xb6, 0x34, 0x98, 0x37, 0xa3, 0x2f, 0x08, 0x06, 0xd3,
	0x9c, 0x25, 0x97, 0x46, 0xce, 0x93, 0x72, 0xc1, 0x44, 0x71, 0x07, 0x8a, 0x76, 0x21, 0x4a, 0xed,
	0xb7, 0x96, 0x67, 0xd8, 0xa1, 0x0d, 0x75, 0xad, 0x6c, 0xfa, 0x83, 0x74, 0x7d, 0x44, 0xd0, 0x5b,
	0xf9, 0x35, 0x3b, 0xbd, 0x03, 0x4d, 0x5b, 0x10, 0x08, 0xe9, 0xf5, 0x04, 0x42, 0x6e, 0xa0, 0xe5,
	0x33, 0x02, 0x98, 0xb2, 0xa2, 0xe0, 0xe9, 0x99, 0xe2, 0xd5, 0x1d, 0x48, 0xf9, 0x0b, 0x3a, 0xb5,
	0xe2, 0x95, 0xc9, 0x79, 0x7f, 0x0c, 0x9c, 0xa5, 0x9b, 0xcc, 0x2d, 0x80, 0xc1, 0x8b, 0xa4, 0x6a,
	0xa4, 0x3e, 0xce, 0x58, 0x9e, 0xf3, 0xe2, 0x82, 0x6f, 0x2e, 0x2c, 0xfe, 0x95, 0x47, 0xcb, 0xff,
	0x3a, 0x10, 0x29, 0xee, 0x03, 0x3a, 0xb1, 0xf2, 0xdb, 0x14, 0x9d, 0x18, 0x74, 0xea, 0xef, 0x1c,
	0x74, 0x6a, 0x50, 0x65, 0x55, 0xb6, 0x29, 0xb2, 0xd7, 0x94, 0x62, 0xb9, 0xf6, 0x0a, 0x6d, 0x6c,
	0xb8, 0x8c, 0xa9, 0x8c, 0x74, 0x1d, 0x67, 0xe2, 0xb5, 0x93, 0xb1, 0x65, 0x97, 0x2e, 0x4f, 0xc6,
	0x0e, 0xb4, 0xd3, 0xcb, 0x9c, 0x17, 0x64, 0x60, 0x69, 0x07, 0x4c, 0x35, 0x2b, 0xd4, 0x7b, 0x5e,
	0x11, 0x70, 0x3e, 0x39, 0xb4, 0xee, 0x53, 0xef, 0x9a, 0x4f, 0xa3, 0xaf, 0x08, 0xa2, 0xc7, 0x52,
	0x98, 0x0b, 0x6f, 0x17, 0x22, 0xc5, 0x93, 0x8a, 0x6b, 0x3f, 0x00, 0x8f, 0x6e, 0x78, 0xd0, 0xba,
	0xe9, 0xc1, 0x3e, 0xc4, 0x49, 0x2e, 0xb8, 0xcb, 0xba, 0x47, 0xa0, 0xeb, 0x88, 0x59, 0x8a, 0xff,
	0x85, 0x9e, 0x4f, 0xea, 0x46, 0x72, 0xdf, 0x2d, 0x38, 0xea, 0x65, 0x23, 0xb9, 0x29, 0x10, 0xea,
	0x5c, 0x14, 0x9a, 0x57, 0x05, 0xcb, 0x49, 0x67, 0x88, 0xc6, 0x5d, 0x0a, 0x42, 0xcd, 0x3c, 0x33,
	0x3d, 0x78, 0xbd, 0x7f, 0x21, 0x74, 0x56, 0xcf, 0xcd, 0x7c, 0x0e, 0x55, 0x3d, 0x17, 0x1f, 0x0e,
	0xdd, 0x55, 0x7f, 0xc8, 0xa4, 0x98, 0x47, 0xf6, 0xb9, 0xbb, 0xff, 0x23, 0x00, 0x00, 0xff, 0xff,
	0x1b, 0xf5, 0xb2, 0x72, 0x0e, 0x07, 0x00, 0x00,
}
