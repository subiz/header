// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user/user.proto

It has these top-level messages:
	User
	Users
	Location
	Device
	Trace
	ListRequest
	MergeRequest
	GreetingRequest
	CreateRequest
	UserTopic
	Topic
	ListEventsRequest
	ReadEventRequest
*/
package user

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
	Event_UserReadRequested        Event = 0
	Event_UserUpdateRequested      Event = 2
	Event_UserCreateRequested      Event = 3
	Event_UserReadBulkRequested    Event = 4
	Event_UserEventCreateRequested Event = 5
	Event_UserListEventsRequested  Event = 7
	Event_UserTopicSumaryRequested Event = 6
)

var Event_name = map[int32]string{
	0: "UserReadRequested",
	2: "UserUpdateRequested",
	3: "UserCreateRequested",
	4: "UserReadBulkRequested",
	5: "UserEventCreateRequested",
	7: "UserListEventsRequested",
	6: "UserTopicSumaryRequested",
}
var Event_value = map[string]int32{
	"UserReadRequested":        0,
	"UserUpdateRequested":      2,
	"UserCreateRequested":      3,
	"UserReadBulkRequested":    4,
	"UserEventCreateRequested": 5,
	"UserListEventsRequested":  7,
	"UserTopicSumaryRequested": 6,
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
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type User struct {
	Ctx       *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Id        *string         `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	AccountId *string         `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Fullname  *string         `protobuf:"bytes,5,opt,name=fullname" json:"fullname,omitempty"`
	Phones    []string        `protobuf:"bytes,7,rep,name=phones" json:"phones,omitempty"`
	Emails    []string        `protobuf:"bytes,10,rep,name=emails" json:"emails,omitempty"`
	Traces    []*Trace        `protobuf:"bytes,11,rep,name=traces" json:"traces,omitempty"`
	// 	repeated string alias = 12;
	Devices          []*Device `protobuf:"bytes,13,rep,name=devices" json:"devices,omitempty"`
	IsBan            *bool     `protobuf:"varint,14,opt,name=is_ban,json=isBan" json:"is_ban,omitempty"`
	AvatarUrl        *string   `protobuf:"bytes,15,opt,name=avatar_url,json=avatarUrl" json:"avatar_url,omitempty"`
	Segments         []string  `protobuf:"bytes,19,rep,name=segments" json:"segments,omitempty"`
	Labels           []string  `protobuf:"bytes,20,rep,name=labels" json:"labels,omitempty"`
	Unsubscribed     *bool     `protobuf:"varint,21,opt,name=unsubscribed" json:"unsubscribed,omitempty"`
	MarkedSpam       *bool     `protobuf:"varint,22,opt,name=marked_spam,json=markedSpam" json:"marked_spam,omitempty"`
	HardBounced      *bool     `protobuf:"varint,23,opt,name=hard_bounced,json=hardBounced" json:"hard_bounced,omitempty"`
	TotalSessions    *int32    `protobuf:"varint,24,opt,name=total_sessions,json=totalSessions" json:"total_sessions,omitempty"`
	SubizId          *string   `protobuf:"bytes,25,opt,name=subiz_id,json=subizId" json:"subiz_id,omitempty"`
	Timezone         *string   `protobuf:"bytes,26,opt,name=timezone" json:"timezone,omitempty"`
	CountryCode      *string   `protobuf:"bytes,27,opt,name=country_code,json=countryCode" json:"country_code,omitempty"`
	Language         *string   `protobuf:"bytes,28,opt,name=language" json:"language,omitempty"`
	Aliases          []string  `protobuf:"bytes,30,rep,name=aliases" json:"aliases,omitempty"`
	LastSeen         *int64    `protobuf:"varint,31,opt,name=last_seen,json=lastSeen" json:"last_seen,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *User) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *User) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *User) GetFullname() string {
	if m != nil && m.Fullname != nil {
		return *m.Fullname
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
	if m != nil && m.IsBan != nil {
		return *m.IsBan
	}
	return false
}

func (m *User) GetAvatarUrl() string {
	if m != nil && m.AvatarUrl != nil {
		return *m.AvatarUrl
	}
	return ""
}

func (m *User) GetSegments() []string {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *User) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *User) GetUnsubscribed() bool {
	if m != nil && m.Unsubscribed != nil {
		return *m.Unsubscribed
	}
	return false
}

func (m *User) GetMarkedSpam() bool {
	if m != nil && m.MarkedSpam != nil {
		return *m.MarkedSpam
	}
	return false
}

func (m *User) GetHardBounced() bool {
	if m != nil && m.HardBounced != nil {
		return *m.HardBounced
	}
	return false
}

func (m *User) GetTotalSessions() int32 {
	if m != nil && m.TotalSessions != nil {
		return *m.TotalSessions
	}
	return 0
}

func (m *User) GetSubizId() string {
	if m != nil && m.SubizId != nil {
		return *m.SubizId
	}
	return ""
}

func (m *User) GetTimezone() string {
	if m != nil && m.Timezone != nil {
		return *m.Timezone
	}
	return ""
}

func (m *User) GetCountryCode() string {
	if m != nil && m.CountryCode != nil {
		return *m.CountryCode
	}
	return ""
}

func (m *User) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *User) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *User) GetLastSeen() int64 {
	if m != nil && m.LastSeen != nil {
		return *m.LastSeen
	}
	return 0
}

type Users struct {
	Users            []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
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

type Location struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Device struct {
	Id               *int32  `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	UseragentId      *int32  `protobuf:"varint,4,opt,name=useragent_id,json=useragentId" json:"useragent_id,omitempty"`
	Useragent        *string `protobuf:"bytes,5,opt,name=useragent" json:"useragent,omitempty"`
	Resolution       *string `protobuf:"bytes,6,opt,name=resolution" json:"resolution,omitempty"`
	LanguageId       *int32  `protobuf:"varint,7,opt,name=language_id,json=languageId" json:"language_id,omitempty"`
	Language         *string `protobuf:"bytes,8,opt,name=language" json:"language,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Device) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Device) GetUseragentId() int32 {
	if m != nil && m.UseragentId != nil {
		return *m.UseragentId
	}
	return 0
}

func (m *Device) GetUseragent() string {
	if m != nil && m.Useragent != nil {
		return *m.Useragent
	}
	return ""
}

func (m *Device) GetResolution() string {
	if m != nil && m.Resolution != nil {
		return *m.Resolution
	}
	return ""
}

func (m *Device) GetLanguageId() int32 {
	if m != nil && m.LanguageId != nil {
		return *m.LanguageId
	}
	return 0
}

func (m *Device) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

type Trace struct {
	Id               *string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Ip               *string `protobuf:"bytes,4,opt,name=ip" json:"ip,omitempty"`
	LocationId       *int32  `protobuf:"varint,5,opt,name=location_id,json=locationId" json:"location_id,omitempty"`
	Location         *string `protobuf:"bytes,6,opt,name=location" json:"location,omitempty"`
	CityCode         *string `protobuf:"bytes,7,opt,name=city_code,json=cityCode" json:"city_code,omitempty"`
	CityName         *string `protobuf:"bytes,8,opt,name=city_name,json=cityName" json:"city_name,omitempty"`
	CountryName      *string `protobuf:"bytes,9,opt,name=country_name,json=countryName" json:"country_name,omitempty"`
	ContinentCode    *string `protobuf:"bytes,10,opt,name=continent_code,json=continentCode" json:"continent_code,omitempty"`
	CoutryCode       *string `protobuf:"bytes,11,opt,name=coutry_code,json=coutryCode" json:"coutry_code,omitempty"`
	Latitue          *int32  `protobuf:"varint,12,opt,name=latitue" json:"latitue,omitempty"`
	Longitude        *int32  `protobuf:"varint,13,opt,name=longitude" json:"longitude,omitempty"`
	PostalCode       *string `protobuf:"bytes,14,opt,name=postal_code,json=postalCode" json:"postal_code,omitempty"`
	RegionName       *string `protobuf:"bytes,15,opt,name=region_name,json=regionName" json:"region_name,omitempty"`
	Timezone         *string `protobuf:"bytes,16,opt,name=timezone" json:"timezone,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Trace) Reset()                    { *m = Trace{} }
func (m *Trace) String() string            { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()               {}
func (*Trace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Trace) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Trace) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Trace) GetLocationId() int32 {
	if m != nil && m.LocationId != nil {
		return *m.LocationId
	}
	return 0
}

func (m *Trace) GetLocation() string {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return ""
}

func (m *Trace) GetCityCode() string {
	if m != nil && m.CityCode != nil {
		return *m.CityCode
	}
	return ""
}

func (m *Trace) GetCityName() string {
	if m != nil && m.CityName != nil {
		return *m.CityName
	}
	return ""
}

func (m *Trace) GetCountryName() string {
	if m != nil && m.CountryName != nil {
		return *m.CountryName
	}
	return ""
}

func (m *Trace) GetContinentCode() string {
	if m != nil && m.ContinentCode != nil {
		return *m.ContinentCode
	}
	return ""
}

func (m *Trace) GetCoutryCode() string {
	if m != nil && m.CoutryCode != nil {
		return *m.CoutryCode
	}
	return ""
}

func (m *Trace) GetLatitue() int32 {
	if m != nil && m.Latitue != nil {
		return *m.Latitue
	}
	return 0
}

func (m *Trace) GetLongitude() int32 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *Trace) GetPostalCode() string {
	if m != nil && m.PostalCode != nil {
		return *m.PostalCode
	}
	return ""
}

func (m *Trace) GetRegionName() string {
	if m != nil && m.RegionName != nil {
		return *m.RegionName
	}
	return ""
}

func (m *Trace) GetTimezone() string {
	if m != nil && m.Timezone != nil {
		return *m.Timezone
	}
	return ""
}

type ListRequest struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,5,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	StartId          *string         `protobuf:"bytes,6,opt,name=start_id,json=startId" json:"start_id,omitempty"`
	Limit            *int32          `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Keyword          *string         `protobuf:"bytes,4,opt,name=keyword" json:"keyword,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ListRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *ListRequest) GetStartId() string {
	if m != nil && m.StartId != nil {
		return *m.StartId
	}
	return ""
}

func (m *ListRequest) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *ListRequest) GetKeyword() string {
	if m != nil && m.Keyword != nil {
		return *m.Keyword
	}
	return ""
}

type MergeRequest struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,3,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Id               *string         `protobuf:"bytes,5,opt,name=id" json:"id,omitempty"`
	RecentId         *string         `protobuf:"bytes,4,opt,name=recent_id,json=recentId" json:"recent_id,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *MergeRequest) Reset()                    { *m = MergeRequest{} }
func (m *MergeRequest) String() string            { return proto.CompactTextString(m) }
func (*MergeRequest) ProtoMessage()               {}
func (*MergeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MergeRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *MergeRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *MergeRequest) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *MergeRequest) GetRecentId() string {
	if m != nil && m.RecentId != nil {
		return *m.RecentId
	}
	return ""
}

type GreetingRequest struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,5,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Useragent        *string         `protobuf:"bytes,6,opt,name=useragent" json:"useragent,omitempty"`
	UserId           *string         `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Uuid             *string         `protobuf:"bytes,4,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GreetingRequest) Reset()                    { *m = GreetingRequest{} }
func (m *GreetingRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetingRequest) ProtoMessage()               {}
func (*GreetingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GreetingRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *GreetingRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *GreetingRequest) GetUseragent() string {
	if m != nil && m.Useragent != nil {
		return *m.Useragent
	}
	return ""
}

func (m *GreetingRequest) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *GreetingRequest) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

type CreateRequest struct {
	AccountId        *string `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ChallengeId      *string `protobuf:"bytes,3,opt,name=challenge_id,json=challengeId" json:"challenge_id,omitempty"`
	Answer           *string `protobuf:"bytes,4,opt,name=answer" json:"answer,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreateRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *CreateRequest) GetChallengeId() string {
	if m != nil && m.ChallengeId != nil {
		return *m.ChallengeId
	}
	return ""
}

func (m *CreateRequest) GetAnswer() string {
	if m != nil && m.Answer != nil {
		return *m.Answer
	}
	return ""
}

type UserTopic struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *UserTopic) Reset()                    { *m = UserTopic{} }
func (m *UserTopic) String() string            { return proto.CompactTextString(m) }
func (*UserTopic) ProtoMessage()               {}
func (*UserTopic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type Topic struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Topic            *string         `protobuf:"bytes,3,opt,name=topic" json:"topic,omitempty"`
	Type             *string         `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	TotalEvents      *int32          `protobuf:"varint,4,opt,name=total_events,json=totalEvents" json:"total_events,omitempty"`
	UserId           *string         `protobuf:"bytes,7,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Topic) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Topic) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Topic) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *Topic) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Topic) GetTotalEvents() int32 {
	if m != nil && m.TotalEvents != nil {
		return *m.TotalEvents
	}
	return 0
}

func (m *Topic) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

type ListEventsRequest struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	UserId           *string         `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	After            *int64          `protobuf:"varint,4,opt,name=after" json:"after,omitempty"`
	Before           *int64          `protobuf:"varint,5,opt,name=before" json:"before,omitempty"`
	Type             *string         `protobuf:"bytes,6,opt,name=type" json:"type,omitempty"`
	Topic            *string         `protobuf:"bytes,7,opt,name=topic" json:"topic,omitempty"`
	Keyword          *string         `protobuf:"bytes,10,opt,name=keyword" json:"keyword,omitempty"`
	StartId          *string         `protobuf:"bytes,11,opt,name=start_id,json=startId" json:"start_id,omitempty"`
	Limit            *int32          `protobuf:"varint,12,opt,name=limit" json:"limit,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ListEventsRequest) Reset()                    { *m = ListEventsRequest{} }
func (m *ListEventsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEventsRequest) ProtoMessage()               {}
func (*ListEventsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListEventsRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ListEventsRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *ListEventsRequest) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *ListEventsRequest) GetAfter() int64 {
	if m != nil && m.After != nil {
		return *m.After
	}
	return 0
}

func (m *ListEventsRequest) GetBefore() int64 {
	if m != nil && m.Before != nil {
		return *m.Before
	}
	return 0
}

func (m *ListEventsRequest) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *ListEventsRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *ListEventsRequest) GetKeyword() string {
	if m != nil && m.Keyword != nil {
		return *m.Keyword
	}
	return ""
}

func (m *ListEventsRequest) GetStartId() string {
	if m != nil && m.StartId != nil {
		return *m.StartId
	}
	return ""
}

func (m *ListEventsRequest) GetLimit() int32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

type ReadEventRequest struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Topic            *string         `protobuf:"bytes,3,opt,name=topic" json:"topic,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ReadEventRequest) Reset()                    { *m = ReadEventRequest{} }
func (m *ReadEventRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadEventRequest) ProtoMessage()               {}
func (*ReadEventRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ReadEventRequest) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *ReadEventRequest) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *ReadEventRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Users)(nil), "user.Users")
	proto.RegisterType((*Location)(nil), "user.Location")
	proto.RegisterType((*Device)(nil), "user.Device")
	proto.RegisterType((*Trace)(nil), "user.Trace")
	proto.RegisterType((*ListRequest)(nil), "user.ListRequest")
	proto.RegisterType((*MergeRequest)(nil), "user.MergeRequest")
	proto.RegisterType((*GreetingRequest)(nil), "user.GreetingRequest")
	proto.RegisterType((*CreateRequest)(nil), "user.CreateRequest")
	proto.RegisterType((*UserTopic)(nil), "user.UserTopic")
	proto.RegisterType((*Topic)(nil), "user.Topic")
	proto.RegisterType((*ListEventsRequest)(nil), "user.ListEventsRequest")
	proto.RegisterType((*ReadEventRequest)(nil), "user.ReadEventRequest")
	proto.RegisterEnum("user.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0xae, 0x2c, 0x53, 0x3f, 0x43, 0xf9, 0x27, 0x1b, 0x3b, 0xde, 0xc4, 0x69, 0xa2, 0xb0, 0x48,
	0xa1, 0xf6, 0x60, 0x03, 0x79, 0x84, 0xa4, 0x45, 0x61, 0x20, 0xed, 0x81, 0x4e, 0xce, 0xc2, 0x8a,
	0x3b, 0x91, 0xb7, 0xa6, 0x48, 0x95, 0xbb, 0x74, 0xe2, 0x00, 0x7d, 0x88, 0x9e, 0x0a, 0xf4, 0xd0,
	0x3e, 0x42, 0x6f, 0x7d, 0x82, 0x3e, 0x58, 0x31, 0xb3, 0x4b, 0x4a, 0xb2, 0x1b, 0xa0, 0x0d, 0x72,
	0xb1, 0xf5, 0x7d, 0x33, 0xbb, 0xf3, 0xf7, 0xcd, 0x12, 0xf6, 0x6a, 0x8b, 0xd5, 0x29, 0xfd, 0x39,
	0x59, 0x56, 0xa5, 0x2b, 0xc5, 0x36, 0xfd, 0x7e, 0x70, 0x32, 0x33, 0x6e, 0x56, 0x67, 0x97, 0xe8,
	0x4e, 0xca, 0x6a, 0x7e, 0x6a, 0xeb, 0x99, 0x79, 0x7f, 0x7a, 0x81, 0x4a, 0x63, 0x75, 0x9a, 0x95,
	0x8b, 0x45, 0x59, 0x84, 0x7f, 0xfe, 0x54, 0xf2, 0x6b, 0x04, 0xdb, 0xaf, 0x2d, 0x56, 0xe2, 0x09,
	0x74, 0x33, 0xf7, 0x4e, 0x76, 0xc6, 0x9d, 0x49, 0xfc, 0x6c, 0xef, 0x24, 0x38, 0xbd, 0x28, 0x0b,
	0x87, 0xef, 0x5c, 0x4a, 0x36, 0xb1, 0x0b, 0x5b, 0x46, 0xcb, 0xee, 0xb8, 0x33, 0x19, 0xa6, 0x5b,
	0x46, 0x8b, 0xcf, 0x01, 0x54, 0x96, 0x95, 0x75, 0xe1, 0xa6, 0x46, 0xcb, 0x6d, 0xe6, 0x87, 0x81,
	0x39, 0xd3, 0xe2, 0x01, 0x0c, 0xde, 0xd4, 0x79, 0x5e, 0xa8, 0x05, 0xca, 0x88, 0x8d, 0x2d, 0x16,
	0xf7, 0xa0, 0xb7, 0xbc, 0x28, 0x0b, 0xb4, 0xb2, 0x3f, 0xee, 0x4e, 0x86, 0x69, 0x40, 0xc4, 0xe3,
	0x42, 0x99, 0xdc, 0x4a, 0xf0, 0xbc, 0x47, 0xe2, 0x0b, 0xe8, 0xb9, 0x4a, 0x65, 0x68, 0x65, 0x3c,
	0xee, 0x4e, 0xe2, 0x67, 0xf1, 0x09, 0x57, 0xfe, 0x8a, 0xb8, 0x34, 0x98, 0xc4, 0x97, 0xd0, 0xd7,
	0x78, 0x65, 0xc8, 0x6b, 0x87, 0xbd, 0x46, 0xde, 0xeb, 0x1b, 0x26, 0xd3, 0xc6, 0x28, 0x0e, 0xa1,
	0x67, 0xec, 0x74, 0xa6, 0x0a, 0xb9, 0x3b, 0xee, 0x4c, 0x06, 0x69, 0x64, 0xec, 0x73, 0x55, 0x70,
	0x39, 0x57, 0xca, 0xa9, 0x6a, 0x5a, 0x57, 0xb9, 0xdc, 0x0b, 0xe5, 0x30, 0xf3, 0xba, 0xca, 0xa9,
	0x1c, 0x8b, 0xf3, 0x05, 0x16, 0xce, 0xca, 0xbb, 0x9c, 0x5c, 0x8b, 0x29, 0xed, 0x5c, 0xcd, 0x30,
	0xb7, 0xf2, 0xc0, 0xa7, 0xed, 0x91, 0x48, 0x60, 0x54, 0x17, 0xb6, 0x9e, 0xd9, 0xac, 0x32, 0x33,
	0xd4, 0xf2, 0x90, 0xe3, 0x6d, 0x70, 0xe2, 0x31, 0xc4, 0x0b, 0x55, 0x5d, 0xa2, 0x9e, 0xda, 0xa5,
	0x5a, 0xc8, 0x7b, 0xec, 0x02, 0x9e, 0x3a, 0x5f, 0xaa, 0x85, 0x78, 0x02, 0xa3, 0x0b, 0x55, 0xe9,
	0xe9, 0xac, 0xac, 0x8b, 0x0c, 0xb5, 0x3c, 0x62, 0x8f, 0x98, 0xb8, 0xe7, 0x9e, 0x12, 0x4f, 0x61,
	0xd7, 0x95, 0x4e, 0xe5, 0x53, 0x8b, 0xd6, 0x9a, 0xb2, 0xb0, 0x52, 0x8e, 0x3b, 0x93, 0x28, 0xdd,
	0x61, 0xf6, 0x3c, 0x90, 0xe2, 0x3e, 0x0c, 0x58, 0x10, 0x34, 0xae, 0xfb, 0x5c, 0x5f, 0x9f, 0xb1,
	0x1f, 0x96, 0x33, 0x0b, 0x7c, 0x5f, 0x16, 0x28, 0x1f, 0xf8, 0x61, 0x35, 0x98, 0x12, 0xe0, 0x99,
	0x56, 0xd7, 0xd3, 0xac, 0xd4, 0x28, 0x8f, 0xd9, 0x1e, 0x07, 0xee, 0x45, 0xa9, 0x91, 0x8e, 0xe7,
	0xaa, 0x98, 0xd7, 0x6a, 0x8e, 0xf2, 0xa1, 0x3f, 0xde, 0x60, 0x21, 0xa1, 0xaf, 0x72, 0xa3, 0x2c,
	0x5a, 0xf9, 0x88, 0xbb, 0xd3, 0x40, 0x71, 0x0c, 0xc3, 0x5c, 0x59, 0x37, 0xb5, 0x88, 0x85, 0x7c,
	0x3c, 0xee, 0x4c, 0xba, 0x74, 0xcc, 0xba, 0x73, 0xc4, 0x22, 0xf9, 0x0a, 0x22, 0x12, 0xa6, 0x15,
	0x63, 0x88, 0x68, 0x8c, 0x56, 0x76, 0x78, 0xa8, 0xe0, 0x87, 0x4a, 0xb6, 0xd4, 0x1b, 0x12, 0x80,
	0xc1, 0xcb, 0x32, 0x53, 0xce, 0x94, 0x45, 0xf2, 0x57, 0x07, 0x7a, 0x7e, 0xe0, 0x6b, 0x7a, 0x8d,
	0x58, 0xaf, 0x4f, 0x60, 0x44, 0xfe, 0x6a, 0x8e, 0x2b, 0xc5, 0x46, 0x69, 0xdc, 0x72, 0x67, 0x5a,
	0x3c, 0x84, 0x61, 0x0b, 0x83, 0x68, 0x57, 0x84, 0x78, 0x04, 0x50, 0xa1, 0x2d, 0xf3, 0x9a, 0x22,
	0xc9, 0x1e, 0x9b, 0xd7, 0x18, 0x1a, 0x65, 0x53, 0x35, 0xdd, 0xdf, 0xe7, 0xfb, 0xa1, 0xa1, 0x7c,
	0x97, 0xdb, 0x36, 0x0d, 0x36, 0xdb, 0x94, 0xfc, 0xd1, 0x85, 0x88, 0xf5, 0x7c, 0x6b, 0xcf, 0x08,
	0x2f, 0xc3, 0x7e, 0x6d, 0x99, 0x25, 0x87, 0x09, 0xe5, 0x52, 0x98, 0x28, 0x84, 0x09, 0x54, 0x08,
	0x13, 0x50, 0xc8, 0xb2, 0xc5, 0xd4, 0xf3, 0xcc, 0xb8, 0x30, 0xc9, 0xbe, 0x37, 0x12, 0xc1, 0x63,
	0x6c, 0x8c, 0xbc, 0xb3, 0x83, 0x95, 0xf1, 0x07, 0xda, 0xd9, 0x35, 0x19, 0xb0, 0x7d, 0xb8, 0x21,
	0x03, 0x76, 0x79, 0x0a, 0xbb, 0x59, 0x59, 0x38, 0x53, 0x50, 0x87, 0x39, 0x02, 0xb0, 0xd3, 0x4e,
	0xcb, 0x72, 0x98, 0xc7, 0x40, 0xa7, 0x5a, 0x3d, 0xc5, 0xbe, 0x91, 0x9e, 0x62, 0x07, 0x09, 0xfd,
	0x5c, 0x39, 0xe3, 0x6a, 0x94, 0x23, 0xae, 0xae, 0x81, 0x34, 0xa0, 0xbc, 0x2c, 0xe6, 0xc6, 0xd5,
	0x1a, 0xe5, 0x0e, 0xdb, 0x56, 0x04, 0x5d, 0xbc, 0x2c, 0x2d, 0x2d, 0x02, 0x5f, 0xbc, 0xeb, 0x2f,
	0xf6, 0x54, 0x13, 0xb9, 0xc2, 0x39, 0x35, 0x8e, 0x4b, 0xd8, 0x6b, 0x46, 0x48, 0x14, 0x57, 0xb0,
	0xbe, 0x07, 0xfb, 0x9b, 0x7b, 0x90, 0xfc, 0xd6, 0x81, 0xf8, 0xa5, 0xb1, 0x2e, 0xc5, 0x9f, 0x6a,
	0xb4, 0xee, 0xbf, 0x3c, 0x99, 0x9b, 0x4f, 0x64, 0x74, 0xf3, 0x89, 0xa4, 0x85, 0x74, 0xaa, 0x62,
	0x63, 0x2f, 0x2c, 0x24, 0xe1, 0x33, 0x2d, 0x0e, 0x20, 0xca, 0xcd, 0xc2, 0xb8, 0xa0, 0x5f, 0x0f,
	0xa8, 0x31, 0x97, 0x78, 0xfd, 0xb6, 0xac, 0x9a, 0xf7, 0xb6, 0x81, 0xc9, 0xcf, 0x30, 0xfa, 0x1e,
	0xab, 0x39, 0x7e, 0x74, 0x72, 0xdd, 0x9b, 0xc9, 0x79, 0x19, 0x46, 0xad, 0x0c, 0x8f, 0x61, 0x58,
	0x61, 0x86, 0xeb, 0xaf, 0xfd, 0xc0, 0x13, 0x67, 0x3a, 0xf9, 0xbd, 0x03, 0x7b, 0xdf, 0x55, 0x88,
	0xce, 0x14, 0xf3, 0x4f, 0xd7, 0x9f, 0x8d, 0x75, 0xec, 0xdd, 0x5c, 0xc7, 0x23, 0xe8, 0x13, 0x58,
	0x25, 0xdf, 0x23, 0x78, 0xa6, 0x85, 0x80, 0xed, 0xba, 0x6e, 0x93, 0xe4, 0xdf, 0x89, 0x81, 0x9d,
	0x17, 0x15, 0x2a, 0xd7, 0x36, 0x68, 0x33, 0xf4, 0xd6, 0xcd, 0xd0, 0xa4, 0xf6, 0x0b, 0x95, 0xe7,
	0x58, 0xf8, 0x65, 0xee, 0x06, 0xb5, 0x37, 0xdc, 0x99, 0xa6, 0x57, 0x5f, 0x15, 0xf6, 0x2d, 0x56,
	0x21, 0x50, 0x40, 0x49, 0x0c, 0x43, 0x7a, 0x9d, 0x5e, 0x95, 0x4b, 0x93, 0x25, 0x7f, 0x76, 0x20,
	0xe2, 0x5f, 0xff, 0xbf, 0x1d, 0xb7, 0x72, 0x3a, 0x80, 0xc8, 0xd1, 0x55, 0x21, 0x19, 0x0f, 0xa8,
	0x5a, 0x77, 0xbd, 0x6c, 0xbe, 0xb1, 0xfc, 0x9b, 0xb2, 0xf7, 0x1f, 0x04, 0xbc, 0xe2, 0x0f, 0x56,
	0x78, 0xea, 0x98, 0xfb, 0x96, 0xa9, 0xf5, 0xee, 0xf5, 0xd7, 0xbb, 0x97, 0xfc, 0xb2, 0x05, 0x77,
	0x48, 0xe6, 0xde, 0xef, 0xa3, 0x87, 0x79, 0x2b, 0xfb, 0x0f, 0x8e, 0xeb, 0x00, 0x22, 0xf5, 0xc6,
	0x85, 0x36, 0x76, 0x53, 0x0f, 0xa8, 0xbb, 0x33, 0x7c, 0x53, 0x56, 0xbe, 0xb0, 0x6e, 0x1a, 0x50,
	0x5b, 0x6e, 0x6f, 0xad, 0xdc, 0xb6, 0x31, 0xfd, 0xf5, 0xc6, 0xac, 0x2d, 0x0b, 0x6c, 0x2c, 0xcb,
	0xc6, 0xde, 0xc5, 0x1f, 0xd8, 0xbb, 0xd1, 0xda, 0xde, 0x25, 0x3f, 0xc2, 0x7e, 0x8a, 0x4a, 0x73,
	0x4b, 0x3e, 0x5d, 0x47, 0xfe, 0x75, 0x9e, 0x5f, 0xff, 0xdd, 0x81, 0x88, 0x03, 0x89, 0x43, 0xb8,
	0xc3, 0x9f, 0x39, 0x54, 0x3a, 0x04, 0x45, 0xbd, 0xff, 0x99, 0x38, 0x82, 0xbb, 0x44, 0xbf, 0x5e,
	0xea, 0x95, 0x9c, 0x51, 0xef, 0x6f, 0x35, 0x86, 0x0d, 0x9d, 0xa3, 0xde, 0xef, 0x8a, 0xfb, 0x70,
	0xd8, 0x5c, 0xf4, 0xbc, 0xce, 0x2f, 0x57, 0xa6, 0x6d, 0xf1, 0x10, 0x24, 0x99, 0x38, 0xe0, 0xcd,
	0x83, 0x91, 0x38, 0x86, 0x23, 0xb2, 0xde, 0x92, 0x03, 0xea, 0xfd, 0x7e, 0x73, 0x94, 0xd5, 0x7d,
	0x5e, 0x2f, 0x54, 0x75, 0xbd, 0xb2, 0xf6, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x20, 0xb0, 0xd8,
	0xc8, 0xa1, 0x0a, 0x00, 0x00,
}
