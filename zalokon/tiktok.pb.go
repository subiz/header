// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.2
// source: tiktok.proto

package zalokon

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OfficialAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OaId                    string `protobuf:"bytes,2,opt,name=oa_id,json=oaId,proto3" json:"oa_id,omitempty"`
	Description             string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Name                    string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Avatar                  string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Cover                   string `protobuf:"bytes,6,opt,name=cover,proto3" json:"cover,omitempty"`
	AccessToken             string `protobuf:"bytes,7,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Authorized              int64  `protobuf:"varint,8,opt,name=authorized,proto3" json:"authorized,omitempty"`
	AccountId               string `protobuf:"bytes,9,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	StrOaId                 string `protobuf:"bytes,10,opt,name=str_oa_id,json=strOaId,proto3" json:"str_oa_id,omitempty"`
	LastZaloHook            int64  `protobuf:"varint,11,opt,name=last_zalo_hook,json=lastZaloHook,proto3" json:"last_zalo_hook,omitempty"`
	State                   string `protobuf:"bytes,12,opt,name=state,proto3" json:"state,omitempty"` // activated || deleted || failed
	Version                 int32  `protobuf:"varint,13,opt,name=version,proto3" json:"version,omitempty"`
	LastRefreshTokenAt      int64  `protobuf:"varint,14,opt,name=last_refresh_token_at,json=lastRefreshTokenAt,proto3" json:"last_refresh_token_at,omitempty"`
	RefreshToken            string `protobuf:"bytes,15,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	LastCalled              int64  `protobuf:"varint,17,opt,name=last_called,json=lastCalled,proto3" json:"last_called,omitempty"`
	LastRefreshTokenErrorAt int64  `protobuf:"varint,18,opt,name=last_refresh_token_error_at,json=lastRefreshTokenErrorAt,proto3" json:"last_refresh_token_error_at,omitempty"`
	AccessTokenExpiredAt    int64  `protobuf:"varint,19,opt,name=access_token_expired_at,json=accessTokenExpiredAt,proto3" json:"access_token_expired_at,omitempty"`
}

func (x *OfficialAccount) Reset() {
	*x = OfficialAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficialAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficialAccount) ProtoMessage() {}

func (x *OfficialAccount) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficialAccount.ProtoReflect.Descriptor instead.
func (*OfficialAccount) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{0}
}

func (x *OfficialAccount) GetOaId() string {
	if x != nil {
		return x.OaId
	}
	return ""
}

func (x *OfficialAccount) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *OfficialAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OfficialAccount) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *OfficialAccount) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *OfficialAccount) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *OfficialAccount) GetAuthorized() int64 {
	if x != nil {
		return x.Authorized
	}
	return 0
}

func (x *OfficialAccount) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *OfficialAccount) GetStrOaId() string {
	if x != nil {
		return x.StrOaId
	}
	return ""
}

func (x *OfficialAccount) GetLastZaloHook() int64 {
	if x != nil {
		return x.LastZaloHook
	}
	return 0
}

func (x *OfficialAccount) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *OfficialAccount) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *OfficialAccount) GetLastRefreshTokenAt() int64 {
	if x != nil {
		return x.LastRefreshTokenAt
	}
	return 0
}

func (x *OfficialAccount) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *OfficialAccount) GetLastCalled() int64 {
	if x != nil {
		return x.LastCalled
	}
	return 0
}

func (x *OfficialAccount) GetLastRefreshTokenErrorAt() int64 {
	if x != nil {
		return x.LastRefreshTokenErrorAt
	}
	return 0
}

func (x *OfficialAccount) GetAccessTokenExpiredAt() int64 {
	if x != nil {
		return x.AccessTokenExpiredAt
	}
	return 0
}

type AccessTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int64                    `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Message   string                   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	RequestId string                   `protobuf:"bytes,5,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Data      *AccessTokenResponseData `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AccessTokenResponse) Reset() {
	*x = AccessTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessTokenResponse) ProtoMessage() {}

func (x *AccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessTokenResponse.ProtoReflect.Descriptor instead.
func (*AccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{1}
}

func (x *AccessTokenResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AccessTokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AccessTokenResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AccessTokenResponse) GetData() *AccessTokenResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type AccessTokenResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken   string  `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	AdvertiserIds []int64 `protobuf:"varint,3,rep,packed,name=advertiser_ids,json=advertiserIds,proto3" json:"advertiser_ids,omitempty"`
	Scope         []int64 `protobuf:"varint,4,rep,packed,name=scope,proto3" json:"scope,omitempty"`
}

func (x *AccessTokenResponseData) Reset() {
	*x = AccessTokenResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tiktok_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessTokenResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessTokenResponseData) ProtoMessage() {}

func (x *AccessTokenResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_tiktok_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessTokenResponseData.ProtoReflect.Descriptor instead.
func (*AccessTokenResponseData) Descriptor() ([]byte, []int) {
	return file_tiktok_proto_rawDescGZIP(), []int{2}
}

func (x *AccessTokenResponseData) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AccessTokenResponseData) GetAdvertiserIds() []int64 {
	if x != nil {
		return x.AdvertiserIds
	}
	return nil
}

func (x *AccessTokenResponseData) GetScope() []int64 {
	if x != nil {
		return x.Scope
	}
	return nil
}

var File_tiktok_proto protoreflect.FileDescriptor

var file_tiktok_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x7a, 0x61, 0x6c, 0x6f, 0x6b, 0x6f, 0x6e, 0x22, 0xcc, 0x04, 0x0a, 0x0f, 0x4f, 0x66, 0x66, 0x69,
	0x63, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x6f,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x61, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x5f, 0x6f, 0x61,
	0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x4f, 0x61,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x7a, 0x61, 0x6c, 0x6f, 0x5f,
	0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x5a, 0x61, 0x6c, 0x6f, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x65,
	0x64, 0x12, 0x3c, 0x0a, 0x1b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x61, 0x74,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x41, 0x74, 0x12,
	0x35, 0x0a, 0x17, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x14, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x22, 0x98, 0x01, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x7a, 0x61, 0x6c, 0x6f,
	0x6b, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x79, 0x0a, 0x17, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x42, 0x21, 0x5a, 0x1f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x62, 0x69, 0x7a,
	0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x7a, 0x61, 0x6c, 0x6f, 0x6b, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tiktok_proto_rawDescOnce sync.Once
	file_tiktok_proto_rawDescData = file_tiktok_proto_rawDesc
)

func file_tiktok_proto_rawDescGZIP() []byte {
	file_tiktok_proto_rawDescOnce.Do(func() {
		file_tiktok_proto_rawDescData = protoimpl.X.CompressGZIP(file_tiktok_proto_rawDescData)
	})
	return file_tiktok_proto_rawDescData
}

var file_tiktok_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tiktok_proto_goTypes = []interface{}{
	(*OfficialAccount)(nil),         // 0: zalokon.OfficialAccount
	(*AccessTokenResponse)(nil),     // 1: zalokon.AccessTokenResponse
	(*AccessTokenResponseData)(nil), // 2: zalokon.AccessTokenResponseData
}
var file_tiktok_proto_depIdxs = []int32{
	2, // 0: zalokon.AccessTokenResponse.data:type_name -> zalokon.AccessTokenResponseData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tiktok_proto_init() }
func file_tiktok_proto_init() {
	if File_tiktok_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tiktok_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficialAccount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tiktok_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tiktok_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessTokenResponseData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tiktok_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tiktok_proto_goTypes,
		DependencyIndexes: file_tiktok_proto_depIdxs,
		MessageInfos:      file_tiktok_proto_msgTypes,
	}.Build()
	File_tiktok_proto = out.File
	file_tiktok_proto_rawDesc = nil
	file_tiktok_proto_goTypes = nil
	file_tiktok_proto_depIdxs = nil
}
