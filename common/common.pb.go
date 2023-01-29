// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: common.proto

package common

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

type Type int32

const (
	Type_unknown   Type = 0
	Type_user      Type = 1
	Type_agent     Type = 2
	Type_subiz     Type = 3
	Type_app       Type = 4 // dashboard, accmgr
	Type_connector Type = 6 // fabikon, mailkon, subiz-driver, stringee-driver
	Type_bot       Type = 7 // subiz bot
	Type_dummy     Type = 8 // agent that cannot handle action
	Type_workflow  Type = 9
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "unknown",
		1: "user",
		2: "agent",
		3: "subiz",
		4: "app",
		6: "connector",
		7: "bot",
		8: "dummy",
		9: "workflow",
	}
	Type_value = map[string]int32{
		"unknown":   0,
		"user":      1,
		"agent":     2,
		"subiz":     3,
		"app":       4,
		"connector": 6,
		"bot":       7,
		"dummy":     8,
		"workflow":  9,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type Device_DeviceType int32

const (
	Device_unknown Device_DeviceType = 0
	Device_mobile  Device_DeviceType = 1
	Device_tablet  Device_DeviceType = 2
	Device_desktop Device_DeviceType = 3
)

// Enum value maps for Device_DeviceType.
var (
	Device_DeviceType_name = map[int32]string{
		0: "unknown",
		1: "mobile",
		2: "tablet",
		3: "desktop",
	}
	Device_DeviceType_value = map[string]int32{
		"unknown": 0,
		"mobile":  1,
		"tablet":  2,
		"desktop": 3,
	}
)

func (x Device_DeviceType) Enum() *Device_DeviceType {
	p := new(Device_DeviceType)
	*p = x
	return p
}

func (x Device_DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Device_DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (Device_DeviceType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x Device_DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Device_DeviceType.Descriptor instead.
func (Device_DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1, 0}
}

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credential *Credential `protobuf:"bytes,6,opt,name=credential,proto3" json:"credential,omitempty"`
	ByDevice   *Device     `protobuf:"bytes,10,opt,name=by_device,json=byDevice,proto3" json:"by_device,omitempty"`
	// for kafka
	SubTopic string `protobuf:"bytes,11,opt,name=sub_topic,json=subTopic,proto3" json:"sub_topic,omitempty"`
	// list of used fields in the object
	Fields []string `protobuf:"bytes,21,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *Context) GetCredential() *Credential {
	if x != nil {
		return x.Credential
	}
	return nil
}

func (x *Context) GetByDevice() *Device {
	if x != nil {
		return x.ByDevice
	}
	return nil
}

func (x *Context) GetSubTopic() string {
	if x != nil {
		return x.SubTopic
	}
	return ""
}

func (x *Context) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip               string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	UserAgent        string `protobuf:"bytes,4,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"` // hash
	ScreenResolution string `protobuf:"bytes,5,opt,name=screen_resolution,json=screenResolution,proto3" json:"screen_resolution,omitempty"`
	Timezone         string `protobuf:"bytes,6,opt,name=timezone,proto3" json:"timezone,omitempty"` // hash
	Language         string `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	Referrer         string `protobuf:"bytes,8,opt,name=referrer,proto3" json:"referrer,omitempty"` // hash, referer of api.subiz.com.vn (the current page that user is viewing)
	Type             string `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`         // 0 - 1 - 2 -3
	Platform         string `protobuf:"bytes,10,opt,name=platform,proto3" json:"platform,omitempty"`
	// referrer of referrer
	// for example: user go to google.com then click a link to company.com, a request
	// will be sent to api.subiz.com.vn
	// referer will be 'company.com' and source_referer will be 'google.com'
	SourceReferrer string             `protobuf:"bytes,11,opt,name=source_referrer,json=sourceReferrer,proto3" json:"source_referrer,omitempty"`
	GaTrackingIds  []string           `protobuf:"bytes,12,rep,name=ga_tracking_ids,json=gaTrackingIds,proto3" json:"ga_tracking_ids,omitempty"`
	PageUrl        string             `protobuf:"bytes,13,opt,name=page_url,json=pageUrl,proto3" json:"page_url,omitempty"`
	PageTitle      string             `protobuf:"bytes,14,opt,name=page_title,json=pageTitle,proto3" json:"page_title,omitempty"`
	GaClientId     string             `protobuf:"bytes,15,opt,name=ga_client_id,json=gaClientId,proto3" json:"ga_client_id,omitempty"`
	CityName       string             `protobuf:"bytes,18,opt,name=city_name,json=cityName,proto3" json:"city_name,omitempty"`                // derived from ip
	CountryName    string             `protobuf:"bytes,19,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`       // derived from ip
	ContinentCode  string             `protobuf:"bytes,20,opt,name=continent_code,json=continentCode,proto3" json:"continent_code,omitempty"` // derived from ip
	Latitude       float32            `protobuf:"fixed32,22,opt,name=latitude,proto3" json:"latitude,omitempty"`                              // derived from ip
	Longitude      float32            `protobuf:"fixed32,23,opt,name=longitude,proto3" json:"longitude,omitempty"`                            // derived from ip
	PostalCode     string             `protobuf:"bytes,24,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`          // derived from ip
	Isp            string             `protobuf:"bytes,27,opt,name=isp,proto3" json:"isp,omitempty"`                                          // derived from ip
	ContinentName  string             `protobuf:"bytes,28,opt,name=continent_name,json=continentName,proto3" json:"continent_name,omitempty"` // derived from ip
	CountryCode    string             `protobuf:"bytes,29,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`       // derived from ip
	BrowserName    string             `protobuf:"bytes,32,opt,name=browser_name,json=browserName,proto3" json:"browser_name,omitempty"`       // filled
	AdsNetwork     string             `protobuf:"bytes,33,opt,name=ads_network,json=adsNetwork,proto3" json:"ads_network,omitempty"`          // google_adwords
	Source         string             `protobuf:"bytes,34,opt,name=source,proto3" json:"source,omitempty"`                                    // direct, referring, organic, social, advertising
	Campaigns      []*SessionCampaign `protobuf:"bytes,35,rep,name=campaigns,proto3" json:"campaigns,omitempty"`
	Utm            *SessionCampaign   `protobuf:"bytes,36,opt,name=utm,proto3" json:"utm,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Device) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *Device) GetScreenResolution() string {
	if x != nil {
		return x.ScreenResolution
	}
	return ""
}

func (x *Device) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Device) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Device) GetReferrer() string {
	if x != nil {
		return x.Referrer
	}
	return ""
}

func (x *Device) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Device) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Device) GetSourceReferrer() string {
	if x != nil {
		return x.SourceReferrer
	}
	return ""
}

func (x *Device) GetGaTrackingIds() []string {
	if x != nil {
		return x.GaTrackingIds
	}
	return nil
}

func (x *Device) GetPageUrl() string {
	if x != nil {
		return x.PageUrl
	}
	return ""
}

func (x *Device) GetPageTitle() string {
	if x != nil {
		return x.PageTitle
	}
	return ""
}

func (x *Device) GetGaClientId() string {
	if x != nil {
		return x.GaClientId
	}
	return ""
}

func (x *Device) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *Device) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *Device) GetContinentCode() string {
	if x != nil {
		return x.ContinentCode
	}
	return ""
}

func (x *Device) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Device) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Device) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Device) GetIsp() string {
	if x != nil {
		return x.Isp
	}
	return ""
}

func (x *Device) GetContinentName() string {
	if x != nil {
		return x.ContinentName
	}
	return ""
}

func (x *Device) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *Device) GetBrowserName() string {
	if x != nil {
		return x.BrowserName
	}
	return ""
}

func (x *Device) GetAdsNetwork() string {
	if x != nil {
		return x.AdsNetwork
	}
	return ""
}

func (x *Device) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Device) GetCampaigns() []*SessionCampaign {
	if x != nil {
		return x.Campaigns
	}
	return nil
}

func (x *Device) GetUtm() *SessionCampaign {
	if x != nil {
		return x.Utm
	}
	return nil
}

type By struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device      *Device `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
	Id          string  `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	ClientId    string  `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Type        string  `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"` // agent, user, bot, dummy
	DisplayId   string  `protobuf:"bytes,7,opt,name=display_id,json=displayId,proto3" json:"display_id,omitempty"`
	DisplayType string  `protobuf:"bytes,8,opt,name=display_type,json=displayType,proto3" json:"display_type,omitempty"` // agent, user, bot
}

func (x *By) Reset() {
	*x = By{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *By) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*By) ProtoMessage() {}

func (x *By) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use By.ProtoReflect.Descriptor instead.
func (*By) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *By) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *By) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *By) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *By) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *By) GetDisplayId() string {
	if x != nil {
		return x.DisplayId
	}
	return ""
}

func (x *By) GetDisplayType() string {
	if x != nil {
		return x.DisplayType
	}
	return ""
}

type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Issuer    string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Type      Type   `protobuf:"varint,4,opt,name=type,proto3,enum=common.Type" json:"type,omitempty"` // should be agent, user or connector
	// Permission perm = 6;
	ClientId string   `protobuf:"bytes,7,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Scopes   []string `protobuf:"bytes,8,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// for x-signed-key
	SignedObjects []string `protobuf:"bytes,9,rep,name=signed_objects,json=signedObjects,proto3" json:"signed_objects,omitempty"`
	SignedType    string   `protobuf:"bytes,10,opt,name=signed_type,json=signedType,proto3" json:"signed_type,omitempty"`
	// for user-access token
	ProfileId          string `protobuf:"bytes,11,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	AgentAccessTokenId string `protobuf:"bytes,12,opt,name=agent_access_token_id,json=agentAccessTokenId,proto3" json:"agent_access_token_id,omitempty"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *Credential) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Credential) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *Credential) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_unknown
}

func (x *Credential) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Credential) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *Credential) GetSignedObjects() []string {
	if x != nil {
		return x.SignedObjects
	}
	return nil
}

func (x *Credential) GetSignedType() string {
	if x != nil {
		return x.SignedType
	}
	return ""
}

func (x *Credential) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *Credential) GetAgentAccessTokenId() string {
	if x != nil {
		return x.AgentAccessTokenId
	}
	return ""
}

type SessionCampaign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Source  string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Medium  string `protobuf:"bytes,4,opt,name=medium,proto3" json:"medium,omitempty"`
	Term    string `protobuf:"bytes,5,opt,name=term,proto3" json:"term,omitempty"`
	Content string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SessionCampaign) Reset() {
	*x = SessionCampaign{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionCampaign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionCampaign) ProtoMessage() {}

func (x *SessionCampaign) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionCampaign.ProtoReflect.Descriptor instead.
func (*SessionCampaign) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *SessionCampaign) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SessionCampaign) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *SessionCampaign) GetMedium() string {
	if x != nil {
		return x.Medium
	}
	return ""
}

func (x *SessionCampaign) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *SessionCampaign) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x09, 0x62, 0x79, 0x5f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x62, 0x79, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0xb1, 0x07, 0x0a, 0x06, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f,
	0x67, 0x61, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x61, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0c, 0x67, 0x61, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x70, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x20,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x73, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x73, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x22, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x63, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x18, 0x23, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x09, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x73, 0x12, 0x29, 0x0a, 0x03, 0x75, 0x74, 0x6d, 0x18, 0x24, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x03, 0x75, 0x74, 0x6d, 0x22, 0x3e, 0x0a, 0x0a,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x6b, 0x74, 0x6f, 0x70, 0x10, 0x03, 0x22, 0xaf, 0x01, 0x0a,
	0x02, 0x42, 0x79, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0xb4,
	0x02, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x72,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x6d, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x73, 0x75, 0x62, 0x69, 0x7a, 0x10, 0x03,
	0x12, 0x07, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x62, 0x6f, 0x74, 0x10,
	0x07, 0x12, 0x09, 0x0a, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x10, 0x09, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x62, 0x69, 0x7a, 0x2f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_proto_goTypes = []interface{}{
	(Type)(0),               // 0: common.Type
	(Device_DeviceType)(0),  // 1: common.Device.DeviceType
	(*Context)(nil),         // 2: common.Context
	(*Device)(nil),          // 3: common.Device
	(*By)(nil),              // 4: common.By
	(*Credential)(nil),      // 5: common.Credential
	(*SessionCampaign)(nil), // 6: common.SessionCampaign
}
var file_common_proto_depIdxs = []int32{
	5, // 0: common.Context.credential:type_name -> common.Credential
	3, // 1: common.Context.by_device:type_name -> common.Device
	6, // 2: common.Device.campaigns:type_name -> common.SessionCampaign
	6, // 3: common.Device.utm:type_name -> common.SessionCampaign
	3, // 4: common.By.device:type_name -> common.Device
	0, // 5: common.Credential.type:type_name -> common.Type
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*By); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credential); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionCampaign); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
