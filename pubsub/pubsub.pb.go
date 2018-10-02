// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pubsub/pubsub.proto

/*
Package pubsub is a generated protocol buffer package.

It is generated from these files:
	pubsub/pubsub.proto

It has these top-level messages:
	Subscription
	PublishMessage
*/
package pubsub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "git.subiz.net/header/common"

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

type Event int32

const (
	Event_PubsubSynced    Event = 0
	Event_PubsubRequested Event = 1
	Event_PubsubPublish   Event = 2
)

var Event_name = map[int32]string{
	0: "PubsubSynced",
	1: "PubsubRequested",
	2: "PubsubPublish",
}
var Event_value = map[string]int32{
	"PubsubSynced":    0,
	"PubsubRequested": 1,
	"PubsubPublish":   2,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Subscription struct {
	Ctx    *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	UserId string          `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Topic  string          `protobuf:"bytes,3,opt,name=topic" json:"topic,omitempty"`
	// optional SubscribeScope scope = 5;
	// optional string account_id = 6;
	SubId           string   `protobuf:"bytes,7,opt,name=sub_id,json=subId" json:"sub_id,omitempty"`
	TargetTopic     string   `protobuf:"bytes,10,opt,name=target_topic,json=targetTopic" json:"target_topic,omitempty"`
	TargetKey       string   `protobuf:"bytes,11,opt,name=target_key,json=targetKey" json:"target_key,omitempty"`
	Ttls            int64    `protobuf:"varint,12,opt,name=ttls" json:"ttls,omitempty"`
	RouterTopic     string   `protobuf:"bytes,13,opt,name=router_topic,json=routerTopic" json:"router_topic,omitempty"`
	TargetPartition int32    `protobuf:"varint,14,opt,name=target_partition,json=targetPartition" json:"target_partition,omitempty"`
	Topics          []string `protobuf:"bytes,4,rep,name=topics" json:"topics,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Subscription) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Subscription) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Subscription) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Subscription) GetSubId() string {
	if m != nil {
		return m.SubId
	}
	return ""
}

func (m *Subscription) GetTargetTopic() string {
	if m != nil {
		return m.TargetTopic
	}
	return ""
}

func (m *Subscription) GetTargetKey() string {
	if m != nil {
		return m.TargetKey
	}
	return ""
}

func (m *Subscription) GetTtls() int64 {
	if m != nil {
		return m.Ttls
	}
	return 0
}

func (m *Subscription) GetRouterTopic() string {
	if m != nil {
		return m.RouterTopic
	}
	return ""
}

func (m *Subscription) GetTargetPartition() int32 {
	if m != nil {
		return m.TargetPartition
	}
	return 0
}

func (m *Subscription) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type PublishMessage struct {
	Ctx     *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Topics  []string        `protobuf:"bytes,2,rep,name=topics" json:"topics,omitempty"`
	Payload []byte          `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	// user_ids bloom filter: only those users can get message
	UserIdsFilter     []byte `protobuf:"bytes,5,opt,name=user_ids_filter,json=userIdsFilter,proto3" json:"user_ids_filter,omitempty"`
	FilterItemBitSize int32  `protobuf:"varint,7,opt,name=filter_item_bit_size,json=filterItemBitSize" json:"filter_item_bit_size,omitempty"`
	FilterHashCount   int32  `protobuf:"varint,8,opt,name=filter_hash_count,json=filterHashCount" json:"filter_hash_count,omitempty"`
	// negative user_ids bloom filter: those users wont get the message
	NegUserIdsFilter []byte `protobuf:"bytes,6,opt,name=neg_user_ids_filter,json=negUserIdsFilter,proto3" json:"neg_user_ids_filter,omitempty"`
}

func (m *PublishMessage) Reset()                    { *m = PublishMessage{} }
func (m *PublishMessage) String() string            { return proto.CompactTextString(m) }
func (*PublishMessage) ProtoMessage()               {}
func (*PublishMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PublishMessage) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *PublishMessage) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *PublishMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PublishMessage) GetUserIdsFilter() []byte {
	if m != nil {
		return m.UserIdsFilter
	}
	return nil
}

func (m *PublishMessage) GetFilterItemBitSize() int32 {
	if m != nil {
		return m.FilterItemBitSize
	}
	return 0
}

func (m *PublishMessage) GetFilterHashCount() int32 {
	if m != nil {
		return m.FilterHashCount
	}
	return 0
}

func (m *PublishMessage) GetNegUserIdsFilter() []byte {
	if m != nil {
		return m.NegUserIdsFilter
	}
	return nil
}

func init() {
	proto.RegisterType((*Subscription)(nil), "pubsub.Subscription")
	proto.RegisterType((*PublishMessage)(nil), "pubsub.PublishMessage")
	proto.RegisterEnum("pubsub.Event", Event_name, Event_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pubsub service

type PubsubClient interface {
	Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*common.Empty, error)
	Unsubscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*common.Empty, error)
	Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error)
}

type pubsubClient struct {
	cc *grpc.ClientConn
}

func NewPubsubClient(cc *grpc.ClientConn) PubsubClient {
	return &pubsubClient{cc}
}

func (c *pubsubClient) Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/pubsub.Pubsub/Subscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubsubClient) Unsubscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/pubsub.Pubsub/Unsubscribe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubsubClient) Ping(ctx context.Context, in *common.PingRequest, opts ...grpc.CallOption) (*common.Pong, error) {
	out := new(common.Pong)
	err := grpc.Invoke(ctx, "/pubsub.Pubsub/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pubsub service

type PubsubServer interface {
	Subscribe(context.Context, *Subscription) (*common.Empty, error)
	Unsubscribe(context.Context, *Subscription) (*common.Empty, error)
	Ping(context.Context, *common.PingRequest) (*common.Pong, error)
}

func RegisterPubsubServer(s *grpc.Server, srv PubsubServer) {
	s.RegisterService(&_Pubsub_serviceDesc, srv)
}

func _Pubsub_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pubsub.Pubsub/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServer).Subscribe(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pubsub_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pubsub.Pubsub/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServer).Unsubscribe(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pubsub_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pubsub.Pubsub/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServer).Ping(ctx, req.(*common.PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pubsub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pubsub.Pubsub",
	HandlerType: (*PubsubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Subscribe",
			Handler:    _Pubsub_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _Pubsub_Unsubscribe_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Pubsub_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pubsub/pubsub.proto",
}

func init() { proto.RegisterFile("pubsub/pubsub.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdd, 0x6e, 0xd3, 0x40,
	0x10, 0x85, 0x71, 0x7e, 0x5c, 0x3a, 0x71, 0x1a, 0x77, 0x53, 0xc0, 0x8a, 0x84, 0xe4, 0xf6, 0x02,
	0xb9, 0x95, 0x48, 0x50, 0x78, 0x02, 0xa8, 0x8a, 0x88, 0x10, 0x52, 0xe4, 0xd0, 0x6b, 0xcb, 0x3f,
	0x83, 0xb3, 0x22, 0x59, 0x1b, 0xef, 0x2c, 0x6a, 0xf2, 0x26, 0x7d, 0x0e, 0x5e, 0x10, 0x79, 0x77,
	0x03, 0x11, 0x57, 0x70, 0xe5, 0x9c, 0x6f, 0x66, 0x8e, 0x66, 0x8f, 0x26, 0x30, 0xae, 0x55, 0x26,
	0x55, 0x36, 0x33, 0x9f, 0x69, 0xdd, 0x54, 0x54, 0x31, 0xd7, 0xa8, 0x49, 0x54, 0x72, 0x9a, 0x4a,
	0x95, 0xf1, 0xfd, 0x54, 0x20, 0xcd, 0xd6, 0x98, 0x16, 0xd8, 0xcc, 0xf2, 0x6a, 0xbb, 0xad, 0x84,
	0xfd, 0x98, 0x89, 0xab, 0x9f, 0x1d, 0xf0, 0x56, 0x2a, 0x93, 0x79, 0xc3, 0x6b, 0xe2, 0x95, 0x60,
	0x97, 0xd0, 0xcd, 0xe9, 0x21, 0x70, 0x42, 0x27, 0x1a, 0xcc, 0x47, 0x53, 0xdb, 0x7c, 0x5b, 0x09,
	0xc2, 0x07, 0x8a, 0xdb, 0x1a, 0x7b, 0x01, 0x27, 0x4a, 0x62, 0x93, 0xf0, 0x22, 0xe8, 0x84, 0x4e,
	0x74, 0x1a, 0xbb, 0xad, 0x5c, 0x14, 0xec, 0x02, 0xfa, 0x54, 0xd5, 0x3c, 0x0f, 0xba, 0x1a, 0x1b,
	0xc1, 0x9e, 0x81, 0x2b, 0x55, 0xd6, 0x76, 0x9f, 0x18, 0x2c, 0x55, 0xb6, 0x28, 0xd8, 0x25, 0x78,
	0x94, 0x36, 0x25, 0x52, 0x62, 0x66, 0x40, 0x17, 0x07, 0x86, 0x7d, 0xd1, 0x93, 0x2f, 0x01, 0x6c,
	0xcb, 0x37, 0xdc, 0x05, 0x03, 0xdd, 0x70, 0x6a, 0xc8, 0x27, 0xdc, 0x31, 0x06, 0x3d, 0xa2, 0x8d,
	0x0c, 0xbc, 0xd0, 0x89, 0xba, 0xb1, 0xfe, 0xdd, 0xba, 0x36, 0x95, 0x22, 0x6c, 0xac, 0xeb, 0xd0,
	0xb8, 0x1a, 0x66, 0x5c, 0xaf, 0xc1, 0xb7, 0xae, 0x75, 0xda, 0x10, 0x6f, 0x5f, 0x1d, 0x9c, 0x85,
	0x4e, 0xd4, 0x8f, 0x47, 0x86, 0x2f, 0x0f, 0x98, 0x3d, 0x07, 0x57, 0xdb, 0xc8, 0xa0, 0x17, 0x76,
	0xdb, 0x87, 0x1a, 0x75, 0xf5, 0xd8, 0x81, 0xb3, 0xa5, 0xca, 0x36, 0x5c, 0xae, 0x3f, 0xa3, 0x94,
	0x69, 0x89, 0xff, 0x92, 0xdb, 0x1f, 0xb7, 0xce, 0xb1, 0x1b, 0x0b, 0xe0, 0xa4, 0x4e, 0x77, 0x9b,
	0x2a, 0x2d, 0x82, 0x5e, 0xe8, 0x44, 0x5e, 0x7c, 0x90, 0xec, 0x15, 0x8c, 0x6c, 0xd2, 0x32, 0xf9,
	0xca, 0x37, 0x84, 0x4d, 0xd0, 0xd7, 0x1d, 0x43, 0x93, 0xb8, 0xfc, 0xa0, 0x21, 0x9b, 0xc1, 0x85,
	0x29, 0x27, 0x9c, 0x70, 0x9b, 0x64, 0x9c, 0x12, 0xc9, 0xf7, 0xa8, 0x03, 0xef, 0xc7, 0xe7, 0xa6,
	0xb6, 0x20, 0xdc, 0xbe, 0xe7, 0xb4, 0xe2, 0x7b, 0x64, 0x37, 0x60, 0x61, 0xb2, 0x4e, 0xe5, 0x3a,
	0xc9, 0x2b, 0x25, 0x28, 0x78, 0x6a, 0x42, 0x30, 0x85, 0x8f, 0xa9, 0x5c, 0xdf, 0xb6, 0x98, 0xbd,
	0x86, 0xb1, 0xc0, 0x32, 0xf9, 0x7b, 0x11, 0x57, 0x2f, 0xe2, 0x0b, 0x2c, 0xef, 0x8f, 0x77, 0xb9,
	0x79, 0x07, 0xfd, 0xbb, 0x1f, 0x28, 0x88, 0xf9, 0xe0, 0x2d, 0xf5, 0x39, 0xae, 0x76, 0x22, 0xc7,
	0xc2, 0x7f, 0xc2, 0xc6, 0x30, 0x32, 0x24, 0xc6, 0xef, 0x0a, 0x25, 0x61, 0xe1, 0x3b, 0xec, 0x1c,
	0x86, 0x06, 0xda, 0x40, 0xfd, 0xce, 0xfc, 0xd1, 0x01, 0xd7, 0x30, 0xf6, 0x06, 0x4e, 0xed, 0x79,
	0x66, 0xc8, 0x2e, 0xa6, 0xf6, 0xda, 0x8f, 0x2f, 0x76, 0x32, 0x3c, 0x84, 0x7d, 0xb7, 0xad, 0x69,
	0xc7, 0xe6, 0x30, 0xb8, 0x17, 0xf2, 0xff, 0x66, 0xae, 0xa1, 0xb7, 0xe4, 0xa2, 0x64, 0xe3, 0x03,
	0x6e, 0x95, 0x5d, 0x72, 0xe2, 0xfd, 0x86, 0x95, 0x28, 0x33, 0x57, 0xff, 0x6f, 0xde, 0xfe, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x34, 0x04, 0xfd, 0xec, 0x80, 0x03, 0x00, 0x00,
}
