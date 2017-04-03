// Code generated by protoc-gen-go.
// source: router/router.proto
// DO NOT EDIT!

/*
Package router is a generated protocol buffer package.

It is generated from these files:
	router/router.proto

It has these top-level messages:
	Id
	Empty
	Route
	Condition
	Routes
	RouteResult
	UserInfo
	AssignRequest
	GroupDeleteEvent
	GroupUpdateEvent
	AgentDeletedEvent
	AgentUpdateEvent
*/
package router

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

type Error int32

const (
	Error_OK           Error = 0
	Error_NoRouteMatch Error = 1
	Error_NotFound     Error = 404
	Error_AccessDeny   Error = 400
)

var Error_name = map[int32]string{
	0:   "OK",
	1:   "NoRouteMatch",
	404: "NotFound",
	400: "AccessDeny",
}
var Error_value = map[string]int32{
	"OK":           0,
	"NoRouteMatch": 1,
	"NotFound":     404,
	"AccessDeny":   400,
}

func (x Error) String() string {
	return proto.EnumName(Error_name, int32(x))
}
func (Error) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type JoinOperator int32

const (
	JoinOperator_None JoinOperator = 0
	JoinOperator_And  JoinOperator = 1
	JoinOperator_Or   JoinOperator = 2
)

var JoinOperator_name = map[int32]string{
	0: "None",
	1: "And",
	2: "Or",
}
var JoinOperator_value = map[string]int32{
	"None": 0,
	"And":  1,
	"Or":   2,
}

func (x JoinOperator) String() string {
	return proto.EnumName(JoinOperator_name, int32(x))
}
func (JoinOperator) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DeviceType int32

const (
	DeviceType_Unknown DeviceType = 0
	DeviceType_Mobile  DeviceType = 1
	DeviceType_Tablet  DeviceType = 2
	DeviceType_Desktop DeviceType = 3
)

var DeviceType_name = map[int32]string{
	0: "Unknown",
	1: "Mobile",
	2: "Tablet",
	3: "Desktop",
}
var DeviceType_value = map[string]int32{
	"Unknown": 0,
	"Mobile":  1,
	"Tablet":  2,
	"Desktop": 3,
}

func (x DeviceType) String() string {
	return proto.EnumName(DeviceType_name, int32(x))
}
func (DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type AssignTo int32

const (
	AssignTo_Unsetted            AssignTo = 0
	AssignTo_AllAgents           AssignTo = 1
	AssignTo_AgentGroup          AssignTo = 2
	AssignTo_SpecificAgent       AssignTo = 3
	AssignTo_MostRecent          AssignTo = 4
	AssignTo_RoundRobinAllAgents AssignTo = 5
	AssignTo_RoundRobinAgents    AssignTo = 6
	AssignTo_BalancedAllAgents   AssignTo = 7
	AssignTo_BalancedAgents      AssignTo = 8
)

var AssignTo_name = map[int32]string{
	0: "Unsetted",
	1: "AllAgents",
	2: "AgentGroup",
	3: "SpecificAgent",
	4: "MostRecent",
	5: "RoundRobinAllAgents",
	6: "RoundRobinAgents",
	7: "BalancedAllAgents",
	8: "BalancedAgents",
}
var AssignTo_value = map[string]int32{
	"Unsetted":            0,
	"AllAgents":           1,
	"AgentGroup":          2,
	"SpecificAgent":       3,
	"MostRecent":          4,
	"RoundRobinAllAgents": 5,
	"RoundRobinAgents":    6,
	"BalancedAllAgents":   7,
	"BalancedAgents":      8,
}

func (x AssignTo) String() string {
	return proto.EnumName(AssignTo_name, int32(x))
}
func (AssignTo) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Id struct {
	Id string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Route struct {
	Priority         string       `protobuf:"bytes,1,opt,name=Priority,json=priority" json:"Priority,omitempty"`
	AssignTo         AssignTo     `protobuf:"varint,2,opt,name=AssignTo,json=assignTo,enum=router.AssignTo" json:"AssignTo,omitempty"`
	AssignToAgentId  string       `protobuf:"bytes,3,opt,name=AssignToAgentId,json=assignToAgentId" json:"AssignToAgentId,omitempty"`
	AssignToGroupId  string       `protobuf:"bytes,4,opt,name=AssignToGroupId,json=assignToGroupId" json:"AssignToGroupId,omitempty"`
	AssignToAgentIds []string     `protobuf:"bytes,6,rep,name=AssignToAgentIds,json=assignToAgentIds" json:"AssignToAgentIds,omitempty"`
	Conditions       []*Condition `protobuf:"bytes,7,rep,name=Conditions,json=conditions" json:"Conditions,omitempty"`
	IsDisabled       bool         `protobuf:"varint,8,opt,name=IsDisabled,json=isDisabled" json:"IsDisabled,omitempty"`
	CTime            string       `protobuf:"bytes,9,opt,name=CTime,json=cTime" json:"CTime,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Route) GetPriority() string {
	if m != nil {
		return m.Priority
	}
	return ""
}

func (m *Route) GetAssignTo() AssignTo {
	if m != nil {
		return m.AssignTo
	}
	return AssignTo_Unsetted
}

func (m *Route) GetAssignToAgentId() string {
	if m != nil {
		return m.AssignToAgentId
	}
	return ""
}

func (m *Route) GetAssignToGroupId() string {
	if m != nil {
		return m.AssignToGroupId
	}
	return ""
}

func (m *Route) GetAssignToAgentIds() []string {
	if m != nil {
		return m.AssignToAgentIds
	}
	return nil
}

func (m *Route) GetConditions() []*Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *Route) GetIsDisabled() bool {
	if m != nil {
		return m.IsDisabled
	}
	return false
}

func (m *Route) GetCTime() string {
	if m != nil {
		return m.CTime
	}
	return ""
}

type Condition struct {
	Join     JoinOperator `protobuf:"varint,1,opt,name=Join,json=join,enum=router.JoinOperator" json:"Join,omitempty"`
	Key      string       `protobuf:"bytes,2,opt,name=Key,json=key" json:"Key,omitempty"`
	Operator string       `protobuf:"bytes,3,opt,name=Operator,json=operator" json:"Operator,omitempty"`
	Value    string       `protobuf:"bytes,4,opt,name=Value,json=value" json:"Value,omitempty"`
}

func (m *Condition) Reset()                    { *m = Condition{} }
func (m *Condition) String() string            { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()               {}
func (*Condition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Condition) GetJoin() JoinOperator {
	if m != nil {
		return m.Join
	}
	return JoinOperator_None
}

func (m *Condition) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Condition) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *Condition) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Routes struct {
	Routes []*Route `protobuf:"bytes,1,rep,name=Routes,json=routes" json:"Routes,omitempty"`
}

func (m *Routes) Reset()                    { *m = Routes{} }
func (m *Routes) String() string            { return proto.CompactTextString(m) }
func (*Routes) ProtoMessage()               {}
func (*Routes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Routes) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type RouteResult struct {
	AgentId []string `protobuf:"bytes,1,rep,name=AgentId,json=agentId" json:"AgentId,omitempty"`
}

func (m *RouteResult) Reset()                    { *m = RouteResult{} }
func (m *RouteResult) String() string            { return proto.CompactTextString(m) }
func (*RouteResult) ProtoMessage()               {}
func (*RouteResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RouteResult) GetAgentId() []string {
	if m != nil {
		return m.AgentId
	}
	return nil
}

type UserInfo struct {
	Name         string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	Email        string `protobuf:"bytes,2,opt,name=Email,json=email" json:"Email,omitempty"`
	Phone        string `protobuf:"bytes,3,opt,name=Phone,json=phone" json:"Phone,omitempty"`
	Sessions     string `protobuf:"bytes,4,opt,name=Sessions,json=sessions" json:"Sessions,omitempty"`
	Country      string `protobuf:"bytes,5,opt,name=Country,json=country" json:"Country,omitempty"`
	CountryCode  string `protobuf:"bytes,6,opt,name=CountryCode,json=countryCode" json:"CountryCode,omitempty"`
	City         string `protobuf:"bytes,7,opt,name=City,json=city" json:"City,omitempty"`
	TimeZone     string `protobuf:"bytes,8,opt,name=TimeZone,json=timeZone" json:"TimeZone,omitempty"`
	Segment      string `protobuf:"bytes,9,opt,name=Segment,json=segment" json:"Segment,omitempty"`
	Label        string `protobuf:"bytes,10,opt,name=Label,json=label" json:"Label,omitempty"`
	Unsubscribed bool   `protobuf:"varint,11,opt,name=Unsubscribed,json=unsubscribed" json:"Unsubscribed,omitempty"`
	MarkedSpam   bool   `protobuf:"varint,12,opt,name=MarkedSpam,json=markedSpam" json:"MarkedSpam,omitempty"`
	HardBounced  bool   `protobuf:"varint,13,opt,name=HardBounced,json=hardBounced" json:"HardBounced,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserInfo) GetSessions() string {
	if m != nil {
		return m.Sessions
	}
	return ""
}

func (m *UserInfo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *UserInfo) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

func (m *UserInfo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *UserInfo) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

func (m *UserInfo) GetSegment() string {
	if m != nil {
		return m.Segment
	}
	return ""
}

func (m *UserInfo) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *UserInfo) GetUnsubscribed() bool {
	if m != nil {
		return m.Unsubscribed
	}
	return false
}

func (m *UserInfo) GetMarkedSpam() bool {
	if m != nil {
		return m.MarkedSpam
	}
	return false
}

func (m *UserInfo) GetHardBounced() bool {
	if m != nil {
		return m.HardBounced
	}
	return false
}

type AssignRequest struct {
	ConvoId         string     `protobuf:"bytes,1,opt,name=ConvoId,json=convoId" json:"ConvoId,omitempty"`
	Channel         string     `protobuf:"bytes,2,opt,name=Channel,json=channel" json:"Channel,omitempty"`
	MessageTo       string     `protobuf:"bytes,3,opt,name=MessageTo,json=messageTo" json:"MessageTo,omitempty"`
	PageUrl         string     `protobuf:"bytes,4,opt,name=PageUrl,json=pageUrl" json:"PageUrl,omitempty"`
	PageTitle       string     `protobuf:"bytes,5,opt,name=PageTitle,json=pageTitle" json:"PageTitle,omitempty"`
	MessageContent  string     `protobuf:"bytes,6,opt,name=MessageContent,json=messageContent" json:"MessageContent,omitempty"`
	BrowserLanguage string     `protobuf:"bytes,7,opt,name=BrowserLanguage,json=browserLanguage" json:"BrowserLanguage,omitempty"`
	Language        string     `protobuf:"bytes,8,opt,name=Language,json=language" json:"Language,omitempty"`
	DeviceType      DeviceType `protobuf:"varint,9,opt,name=DeviceType,json=deviceType,enum=router.DeviceType" json:"DeviceType,omitempty"`
	User            *UserInfo  `protobuf:"bytes,10,opt,name=User,json=user" json:"User,omitempty"`
	AccountId       string     `protobuf:"bytes,11,opt,name=AccountId,json=accountId" json:"AccountId,omitempty"`
}

func (m *AssignRequest) Reset()                    { *m = AssignRequest{} }
func (m *AssignRequest) String() string            { return proto.CompactTextString(m) }
func (*AssignRequest) ProtoMessage()               {}
func (*AssignRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AssignRequest) GetConvoId() string {
	if m != nil {
		return m.ConvoId
	}
	return ""
}

func (m *AssignRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *AssignRequest) GetMessageTo() string {
	if m != nil {
		return m.MessageTo
	}
	return ""
}

func (m *AssignRequest) GetPageUrl() string {
	if m != nil {
		return m.PageUrl
	}
	return ""
}

func (m *AssignRequest) GetPageTitle() string {
	if m != nil {
		return m.PageTitle
	}
	return ""
}

func (m *AssignRequest) GetMessageContent() string {
	if m != nil {
		return m.MessageContent
	}
	return ""
}

func (m *AssignRequest) GetBrowserLanguage() string {
	if m != nil {
		return m.BrowserLanguage
	}
	return ""
}

func (m *AssignRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *AssignRequest) GetDeviceType() DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return DeviceType_Unknown
}

func (m *AssignRequest) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *AssignRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type GroupDeleteEvent struct {
	GroupId string `protobuf:"bytes,1,opt,name=GroupId,json=groupId" json:"GroupId,omitempty"`
}

func (m *GroupDeleteEvent) Reset()                    { *m = GroupDeleteEvent{} }
func (m *GroupDeleteEvent) String() string            { return proto.CompactTextString(m) }
func (*GroupDeleteEvent) ProtoMessage()               {}
func (*GroupDeleteEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GroupDeleteEvent) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

type GroupUpdateEvent struct {
	GroupId   string   `protobuf:"bytes,1,opt,name=GroupId,json=groupId" json:"GroupId,omitempty"`
	MemberIds []string `protobuf:"bytes,2,rep,name=MemberIds,json=memberIds" json:"MemberIds,omitempty"`
}

func (m *GroupUpdateEvent) Reset()                    { *m = GroupUpdateEvent{} }
func (m *GroupUpdateEvent) String() string            { return proto.CompactTextString(m) }
func (*GroupUpdateEvent) ProtoMessage()               {}
func (*GroupUpdateEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GroupUpdateEvent) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *GroupUpdateEvent) GetMemberIds() []string {
	if m != nil {
		return m.MemberIds
	}
	return nil
}

type AgentDeletedEvent struct {
	AgentId string `protobuf:"bytes,1,opt,name=AgentId,json=agentId" json:"AgentId,omitempty"`
}

func (m *AgentDeletedEvent) Reset()                    { *m = AgentDeletedEvent{} }
func (m *AgentDeletedEvent) String() string            { return proto.CompactTextString(m) }
func (*AgentDeletedEvent) ProtoMessage()               {}
func (*AgentDeletedEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AgentDeletedEvent) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

type AgentUpdateEvent struct {
	AgentId     string `protobuf:"bytes,1,opt,name=AgentId,json=agentId" json:"AgentId,omitempty"`
	IsActive    bool   `protobuf:"varint,2,opt,name=IsActive,json=isActive" json:"IsActive,omitempty"`
	IsConfirmed bool   `protobuf:"varint,3,opt,name=IsConfirmed,json=isConfirmed" json:"IsConfirmed,omitempty"`
}

func (m *AgentUpdateEvent) Reset()                    { *m = AgentUpdateEvent{} }
func (m *AgentUpdateEvent) String() string            { return proto.CompactTextString(m) }
func (*AgentUpdateEvent) ProtoMessage()               {}
func (*AgentUpdateEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AgentUpdateEvent) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *AgentUpdateEvent) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *AgentUpdateEvent) GetIsConfirmed() bool {
	if m != nil {
		return m.IsConfirmed
	}
	return false
}

func init() {
	proto.RegisterType((*Id)(nil), "router.Id")
	proto.RegisterType((*Empty)(nil), "router.Empty")
	proto.RegisterType((*Route)(nil), "router.Route")
	proto.RegisterType((*Condition)(nil), "router.Condition")
	proto.RegisterType((*Routes)(nil), "router.Routes")
	proto.RegisterType((*RouteResult)(nil), "router.RouteResult")
	proto.RegisterType((*UserInfo)(nil), "router.UserInfo")
	proto.RegisterType((*AssignRequest)(nil), "router.AssignRequest")
	proto.RegisterType((*GroupDeleteEvent)(nil), "router.GroupDeleteEvent")
	proto.RegisterType((*GroupUpdateEvent)(nil), "router.GroupUpdateEvent")
	proto.RegisterType((*AgentDeletedEvent)(nil), "router.AgentDeletedEvent")
	proto.RegisterType((*AgentUpdateEvent)(nil), "router.AgentUpdateEvent")
	proto.RegisterEnum("router.Error", Error_name, Error_value)
	proto.RegisterEnum("router.JoinOperator", JoinOperator_name, JoinOperator_value)
	proto.RegisterEnum("router.DeviceType", DeviceType_name, DeviceType_value)
	proto.RegisterEnum("router.AssignTo", AssignTo_name, AssignTo_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Router service

type RouterClient interface {
	Assign(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*RouteResult, error)
	Update(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Empty, error)
	Create(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Id, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	// List all routes by account id
	List(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Routes, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Assign(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*RouteResult, error) {
	out := new(RouteResult)
	err := grpc.Invoke(ctx, "/router.Router/Assign", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Update(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/router.Router/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Create(ctx context.Context, in *Route, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := grpc.Invoke(ctx, "/router.Router/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/router.Router/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) List(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Routes, error) {
	out := new(Routes)
	err := grpc.Invoke(ctx, "/router.Router/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Router service

type RouterServer interface {
	Assign(context.Context, *AssignRequest) (*RouteResult, error)
	Update(context.Context, *Route) (*Empty, error)
	Create(context.Context, *Route) (*Id, error)
	Delete(context.Context, *Id) (*Empty, error)
	// List all routes by account id
	List(context.Context, *Id) (*Routes, error)
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_Assign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Assign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/router.Router/Assign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Assign(ctx, req.(*AssignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Route)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/router.Router/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Update(ctx, req.(*Route))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Route)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/router.Router/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Create(ctx, req.(*Route))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/router.Router/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/router.Router/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).List(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "router.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Assign",
			Handler:    _Router_Assign_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Router_Update_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Router_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Router_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Router_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router/router.proto",
}

func init() { proto.RegisterFile("router/router.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x16, 0xf5, 0x47, 0x6a, 0x64, 0xcb, 0xf4, 0xc6, 0x41, 0x09, 0xa3, 0x28, 0x0c, 0x22, 0x6d,
	0x55, 0x23, 0x4d, 0x50, 0x17, 0xe8, 0xb1, 0x80, 0x22, 0xbb, 0xad, 0x92, 0xd8, 0x09, 0x68, 0xab,
	0x87, 0xde, 0x28, 0x72, 0x22, 0x6f, 0x4c, 0xee, 0xb2, 0xbb, 0x4b, 0x07, 0x7a, 0x8b, 0x1e, 0xfa,
	0x28, 0x3d, 0xf5, 0x61, 0x7a, 0xee, 0x5b, 0xb4, 0xd8, 0x1f, 0xea, 0xc7, 0x0d, 0xd0, 0x8b, 0xbd,
	0xdf, 0x37, 0xc3, 0xd9, 0xd9, 0xf9, 0x66, 0x67, 0x05, 0x8f, 0x04, 0xaf, 0x15, 0x8a, 0xe7, 0xf6,
	0xdf, 0xb3, 0x4a, 0x70, 0xc5, 0x49, 0xdf, 0xa2, 0xf8, 0x08, 0xda, 0xb3, 0x9c, 0x8c, 0xf4, 0xdf,
	0xc8, 0x3b, 0xf1, 0xc6, 0x83, 0xa4, 0x4d, 0xf3, 0xd8, 0x87, 0xde, 0x45, 0x59, 0xa9, 0x55, 0xfc,
	0x67, 0x1b, 0x7a, 0x89, 0xf6, 0x24, 0xc7, 0x10, 0xbc, 0x15, 0x94, 0x0b, 0xaa, 0x56, 0xce, 0x31,
	0xa8, 0x1c, 0x26, 0x4f, 0x21, 0x98, 0x48, 0x49, 0x97, 0xec, 0x86, 0x47, 0xed, 0x13, 0x6f, 0x3c,
	0x3a, 0x0b, 0x9f, 0xb9, 0xdd, 0x1a, 0x3e, 0x09, 0x52, 0xb7, 0x22, 0x63, 0x38, 0x68, 0xd8, 0xc9,
	0x12, 0x99, 0x9a, 0xe5, 0x51, 0xc7, 0x04, 0x3c, 0x48, 0x77, 0xe9, 0x6d, 0xcf, 0x1f, 0x05, 0xaf,
	0xab, 0x59, 0x1e, 0x75, 0x77, 0x3d, 0x1d, 0x4d, 0x4e, 0x21, 0x7c, 0x10, 0x53, 0x46, 0xfd, 0x93,
	0xce, 0x78, 0x90, 0x84, 0x0f, 0x82, 0x4a, 0xf2, 0x0d, 0xc0, 0x94, 0xb3, 0x9c, 0x2a, 0xca, 0x99,
	0x8c, 0xfc, 0x93, 0xce, 0x78, 0x78, 0x76, 0xd8, 0xe4, 0xbb, 0xb6, 0x24, 0x90, 0xad, 0x9d, 0xc8,
	0x67, 0x00, 0x33, 0x79, 0x4e, 0x65, 0xba, 0x28, 0x30, 0x8f, 0x82, 0x13, 0x6f, 0x1c, 0x24, 0x40,
	0xd7, 0x0c, 0x39, 0x82, 0xde, 0xf4, 0x86, 0x96, 0x18, 0x0d, 0x4c, 0x7a, 0xbd, 0x4c, 0x83, 0x78,
	0x05, 0x83, 0x75, 0x38, 0x32, 0x86, 0xee, 0x4b, 0x4e, 0x99, 0xa9, 0xdd, 0xe8, 0xec, 0xa8, 0xd9,
	0x4f, 0x73, 0x6f, 0x2a, 0x14, 0xa9, 0xe2, 0x22, 0xe9, 0xbe, 0xe7, 0x94, 0x91, 0x10, 0x3a, 0xaf,
	0x70, 0x65, 0x0a, 0x39, 0x48, 0x3a, 0x77, 0xb8, 0xd2, 0xb5, 0x6f, 0x7c, 0x5c, 0xa9, 0x02, 0xee,
	0xb0, 0xde, 0xfa, 0xe7, 0xb4, 0xa8, 0xd1, 0x55, 0xa6, 0x77, 0xaf, 0x41, 0xfc, 0x1c, 0xfa, 0x46,
	0x36, 0x49, 0x3e, 0x6f, 0x56, 0x91, 0x67, 0x4e, 0xba, 0xdf, 0xec, 0x6c, 0xd8, 0xc4, 0xf6, 0x81,
	0x8c, 0xbf, 0x84, 0xa1, 0x25, 0x50, 0xd6, 0x85, 0x22, 0x11, 0xf8, 0x8d, 0x36, 0x9e, 0x29, 0xa3,
	0x9f, 0x5a, 0x18, 0xff, 0xdd, 0x86, 0x60, 0x2e, 0x51, 0xcc, 0xd8, 0x3b, 0x4e, 0x08, 0x74, 0xaf,
	0xd2, 0x12, 0x5d, 0x43, 0x74, 0x59, 0x5a, 0xa2, 0x4e, 0xe8, 0xa2, 0x4c, 0x69, 0xe1, 0x0e, 0xd0,
	0x43, 0x0d, 0x34, 0xfb, 0xf6, 0x96, 0x33, 0x74, 0xf9, 0xf7, 0x2a, 0x0d, 0xf4, 0xc1, 0xae, 0x51,
	0x4a, 0x23, 0x84, 0xcd, 0x3f, 0x90, 0x0e, 0xeb, 0x14, 0xa6, 0xbc, 0x66, 0x4a, 0xac, 0xa2, 0x9e,
	0x31, 0xf9, 0x99, 0x85, 0xe4, 0x04, 0x86, 0xce, 0x32, 0xe5, 0x39, 0x46, 0x7d, 0x63, 0x1d, 0x66,
	0x1b, 0x4a, 0xe7, 0x35, 0xd5, 0x8d, 0xea, 0xdb, 0xbc, 0x32, 0xdd, 0xa4, 0xc7, 0x10, 0x68, 0x55,
	0x7e, 0xd1, 0x49, 0x04, 0x76, 0x2f, 0xe5, 0xb0, 0xde, 0xeb, 0x1a, 0x97, 0x25, 0x32, 0xe5, 0x14,
	0xf4, 0xa5, 0x85, 0x3a, 0xef, 0xd7, 0xe9, 0x02, 0x8b, 0x08, 0x6c, 0xde, 0x85, 0x06, 0x24, 0x86,
	0xbd, 0x39, 0x93, 0xf5, 0x42, 0x66, 0x82, 0x2e, 0x30, 0x8f, 0x86, 0xa6, 0x23, 0xf6, 0xea, 0x2d,
	0x4e, 0xf7, 0xcc, 0x65, 0x2a, 0xee, 0x30, 0xbf, 0xae, 0xd2, 0x32, 0xda, 0xb3, 0x3d, 0x53, 0xae,
	0x19, 0x7d, 0x8a, 0x9f, 0x52, 0x91, 0xbf, 0xe0, 0x35, 0xcb, 0x30, 0x8f, 0xf6, 0x8d, 0xc3, 0xf0,
	0x76, 0x43, 0xc5, 0xff, 0xb4, 0x61, 0xdf, 0x76, 0x75, 0x82, 0xbf, 0xd6, 0x28, 0x95, 0xad, 0x09,
	0xbb, 0xe7, 0xeb, 0xcb, 0xea, 0x67, 0x16, 0x1a, 0xcb, 0x6d, 0xca, 0x18, 0x36, 0x75, 0xf7, 0x33,
	0x0b, 0xc9, 0xa7, 0x30, 0xb8, 0x44, 0x29, 0xd3, 0x25, 0xde, 0x70, 0x57, 0xfd, 0x41, 0xd9, 0x10,
	0xfa, 0xbb, 0xb7, 0xe9, 0x12, 0xe7, 0xa2, 0x70, 0x02, 0xf8, 0x95, 0x85, 0xfa, 0x3b, 0x6d, 0xb9,
	0xa1, 0xaa, 0x40, 0xa7, 0xc0, 0xa0, 0x6a, 0x08, 0xf2, 0x05, 0x8c, 0x5c, 0xd4, 0x29, 0x67, 0x4a,
	0x17, 0xce, 0xca, 0x30, 0x2a, 0x77, 0x58, 0x7d, 0x85, 0x5f, 0x08, 0xfe, 0x41, 0xa2, 0x78, 0x9d,
	0xb2, 0x65, 0x9d, 0x2e, 0xd1, 0x89, 0x72, 0xb0, 0xd8, 0xa5, 0xb5, 0x3e, 0x6b, 0x17, 0xa7, 0x4f,
	0xd1, 0xd8, 0xce, 0x00, 0xce, 0xf1, 0x9e, 0x66, 0x78, 0xb3, 0xaa, 0xec, 0x25, 0x1b, 0x9d, 0x91,
	0xa6, 0x91, 0x37, 0x96, 0x04, 0xf2, 0xf5, 0x9a, 0x3c, 0x81, 0xae, 0xee, 0x53, 0x23, 0xdc, 0x70,
	0x33, 0x90, 0x9a, 0xde, 0x4d, 0xba, 0xb5, 0x44, 0xa1, 0x4f, 0x39, 0xc9, 0x4c, 0xeb, 0xcc, 0xac,
	0x8c, 0x83, 0x64, 0x90, 0x36, 0x44, 0xfc, 0x14, 0x42, 0x33, 0x61, 0xce, 0xb1, 0x40, 0x85, 0x17,
	0xf7, 0xfa, 0x44, 0x11, 0xf8, 0xcd, 0x30, 0x72, 0x1a, 0x2c, 0x2d, 0x8c, 0x5f, 0x3a, 0xef, 0x79,
	0x95, 0xa7, 0xff, 0xeb, 0x6d, 0x75, 0x29, 0x17, 0x28, 0xf4, 0xac, 0x6a, 0x9b, 0x4b, 0x36, 0x28,
	0x1b, 0x22, 0xfe, 0x1a, 0x0e, 0xcd, 0x05, 0xb4, 0x3b, 0xe7, 0xeb, 0x60, 0x9b, 0x5b, 0xe9, 0x6d,
	0xdf, 0xca, 0xf7, 0x10, 0x1a, 0xcb, 0x83, 0xad, 0x3f, 0xee, 0xad, 0x4b, 0x3d, 0x93, 0x93, 0x4c,
	0xd1, 0x7b, 0x34, 0xdd, 0x12, 0x24, 0x01, 0x75, 0x58, 0xb7, 0xe5, 0x4c, 0x4e, 0x39, 0x7b, 0x47,
	0x45, 0x89, 0x76, 0x32, 0x07, 0xc9, 0x90, 0x6e, 0xa8, 0xd3, 0x09, 0xf4, 0x2e, 0x84, 0xe0, 0x82,
	0xf4, 0xa1, 0xfd, 0xe6, 0x55, 0xd8, 0x22, 0x21, 0xec, 0x5d, 0x71, 0x33, 0x3d, 0x2e, 0x53, 0x95,
	0xdd, 0x86, 0x1e, 0xd9, 0x87, 0xe0, 0x8a, 0xab, 0x1f, 0x78, 0xcd, 0xf2, 0xf0, 0xf7, 0x0e, 0x39,
	0x00, 0x98, 0x64, 0x19, 0x4a, 0x79, 0x8e, 0x6c, 0x15, 0xfe, 0xd6, 0x39, 0xfd, 0x0a, 0xf6, 0xb6,
	0x07, 0x1f, 0x09, 0xa0, 0x7b, 0xc5, 0x19, 0x86, 0x2d, 0xe2, 0x43, 0x67, 0xc2, 0xf2, 0xd0, 0x33,
	0xc1, 0x45, 0xd8, 0x3e, 0xfd, 0x7e, 0x5b, 0x7a, 0x32, 0x04, 0x7f, 0xce, 0xee, 0x18, 0xff, 0xc0,
	0xc2, 0x16, 0x01, 0xe8, 0x5f, 0xf2, 0x05, 0x2d, 0x30, 0xf4, 0xf4, 0xfa, 0x46, 0xcf, 0x62, 0x15,
	0xb6, 0xb5, 0xd3, 0x39, 0xca, 0x3b, 0xc5, 0xab, 0xb0, 0x73, 0xfa, 0x87, 0xb7, 0x79, 0x9c, 0xc8,
	0x1e, 0x04, 0x73, 0x26, 0x51, 0x29, 0xcc, 0xc3, 0x16, 0xd9, 0x87, 0xc1, 0xa4, 0x28, 0x4c, 0x8d,
	0x64, 0xe8, 0x91, 0x11, 0x80, 0x59, 0x1b, 0xbd, 0xc2, 0x36, 0x39, 0x84, 0xfd, 0xeb, 0x0a, 0x33,
	0xfa, 0x8e, 0x66, 0x86, 0x0f, 0x3b, 0xda, 0xe5, 0x92, 0x4b, 0x95, 0x60, 0xa6, 0x71, 0x97, 0x7c,
	0x02, 0x8f, 0x12, 0x7d, 0xc8, 0x84, 0x2f, 0x28, 0xdb, 0xc4, 0xea, 0x91, 0x23, 0x08, 0xb7, 0x0c,
	0x96, 0xed, 0x93, 0xc7, 0x70, 0xf8, 0x22, 0x2d, 0x52, 0x7d, 0xb9, 0x37, 0xce, 0x3e, 0x21, 0x30,
	0x5a, 0xd3, 0x96, 0x0b, 0xce, 0xfe, 0xf2, 0xdc, 0xdc, 0x16, 0xe4, 0x3b, 0xe8, 0xdb, 0x03, 0x90,
	0xc7, 0xbb, 0xaf, 0xaa, 0x9b, 0x0a, 0xc7, 0x8f, 0x76, 0x47, 0xba, 0x99, 0xe0, 0x71, 0x8b, 0x8c,
	0xa1, 0x6f, 0xdb, 0x81, 0xec, 0xce, 0xfc, 0xe3, 0x35, 0xb4, 0x6f, 0x7c, 0x4b, 0xbf, 0x11, 0x53,
	0x81, 0x1f, 0xf1, 0x84, 0x06, 0xce, 0x72, 0xeb, 0x66, 0xdb, 0x91, 0x6c, 0xf1, 0xff, 0x8d, 0xf6,
	0x04, 0xba, 0xaf, 0xa9, 0x54, 0x3b, 0x4e, 0xa3, 0x9d, 0xb8, 0x32, 0x6e, 0x2d, 0xfa, 0xe6, 0x77,
	0xc8, 0xb7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x84, 0x67, 0xbd, 0x17, 0x9e, 0x08, 0x00, 0x00,
}
