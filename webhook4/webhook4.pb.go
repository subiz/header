// Code generated by protoc-gen-go. DO NOT EDIT.
// source: webhook4/webhook4.proto

/*
Package webhook4 is a generated protocol buffer package.

It is generated from these files:
	webhook4/webhook4.proto

It has these top-level messages:
	Retry
	Webhook
*/
package webhook4

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
	Event_Webhook4RetryCallback          Event = 0
	Event_Webhook4WebhookReadRequested   Event = 3
	Event_Webhook4WebhookUpserted        Event = 4
	Event_Webhook4WebhookVerifyRequested Event = 5
	Event_Webhook4Requested              Event = 6
	Event_Webhook4ClearBackoffRequested  Event = 7
	Event_Webhook4PurgeQueueRequested    Event = 8
	Event_Webhook4ShardSend              Event = 9
)

var Event_name = map[int32]string{
	0: "Webhook4RetryCallback",
	3: "Webhook4WebhookReadRequested",
	4: "Webhook4WebhookUpserted",
	5: "Webhook4WebhookVerifyRequested",
	6: "Webhook4Requested",
	7: "Webhook4ClearBackoffRequested",
	8: "Webhook4PurgeQueueRequested",
	9: "Webhook4ShardSend",
}
var Event_value = map[string]int32{
	"Webhook4RetryCallback":          0,
	"Webhook4WebhookReadRequested":   3,
	"Webhook4WebhookUpserted":        4,
	"Webhook4WebhookVerifyRequested": 5,
	"Webhook4Requested":              6,
	"Webhook4ClearBackoffRequested":  7,
	"Webhook4PurgeQueueRequested":    8,
	"Webhook4ShardSend":              9,
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

type Webhook_State int32

const (
	Webhook_normal   Webhook_State = 0
	Webhook_backoff  Webhook_State = 1
	Webhook_inactive Webhook_State = 2
)

var Webhook_State_name = map[int32]string{
	0: "normal",
	1: "backoff",
	2: "inactive",
}
var Webhook_State_value = map[string]int32{
	"normal":   0,
	"backoff":  1,
	"inactive": 2,
}

func (x Webhook_State) Enum() *Webhook_State {
	p := new(Webhook_State)
	*p = x
	return p
}
func (x Webhook_State) String() string {
	return proto.EnumName(Webhook_State_name, int32(x))
}
func (x *Webhook_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Webhook_State_value, data, "Webhook_State")
	if err != nil {
		return err
	}
	*x = Webhook_State(value)
	return nil
}
func (Webhook_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Retry struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Partition        *string         `protobuf:"bytes,2,opt,name=partition" json:"partition,omitempty"`
	Queue            *string         `protobuf:"bytes,3,opt,name=queue" json:"queue,omitempty"`
	JobId            *string         `protobuf:"bytes,4,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	NBackoff         *int32          `protobuf:"varint,5,opt,name=n_backoff,json=nBackoff" json:"n_backoff,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Retry) Reset()                    { *m = Retry{} }
func (m *Retry) String() string            { return proto.CompactTextString(m) }
func (*Retry) ProtoMessage()               {}
func (*Retry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Retry) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Retry) GetPartition() string {
	if m != nil && m.Partition != nil {
		return *m.Partition
	}
	return ""
}

func (m *Retry) GetQueue() string {
	if m != nil && m.Queue != nil {
		return *m.Queue
	}
	return ""
}

func (m *Retry) GetJobId() string {
	if m != nil && m.JobId != nil {
		return *m.JobId
	}
	return ""
}

func (m *Retry) GetNBackoff() int32 {
	if m != nil && m.NBackoff != nil {
		return *m.NBackoff
	}
	return 0
}

type Webhook struct {
	Ctx              *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string         `protobuf:"bytes,12,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	ClientId         *string         `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Url              *string         `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Secret           *string         `protobuf:"bytes,4,opt,name=secret" json:"secret,omitempty"`
	Events           []string        `protobuf:"bytes,6,rep,name=events" json:"events,omitempty"`
	State            *string         `protobuf:"bytes,7,opt,name=state" json:"state,omitempty"`
	NextHookAt       *int64          `protobuf:"varint,8,opt,name=next_hook_at,json=nextHookAt" json:"next_hook_at,omitempty"`
	LastHookAt       *int64          `protobuf:"varint,9,opt,name=last_hook_at,json=lastHookAt" json:"last_hook_at,omitempty"`
	LastHookId       *int64          `protobuf:"varint,10,opt,name=last_hook_id,json=lastHookId" json:"last_hook_id,omitempty"`
	LastHookResponse *string         `protobuf:"bytes,14,opt,name=last_hook_response,json=lastHookResponse" json:"last_hook_response,omitempty"`
	LastHookStatus   *int32          `protobuf:"varint,15,opt,name=last_hook_status,json=lastHookStatus" json:"last_hook_status,omitempty"`
	EventsCount      *int64          `protobuf:"varint,13,opt,name=events_count,json=eventsCount" json:"events_count,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Webhook) Reset()                    { *m = Webhook{} }
func (m *Webhook) String() string            { return proto.CompactTextString(m) }
func (*Webhook) ProtoMessage()               {}
func (*Webhook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Webhook) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Webhook) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Webhook) GetClientId() string {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return ""
}

func (m *Webhook) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *Webhook) GetSecret() string {
	if m != nil && m.Secret != nil {
		return *m.Secret
	}
	return ""
}

func (m *Webhook) GetEvents() []string {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Webhook) GetState() string {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ""
}

func (m *Webhook) GetNextHookAt() int64 {
	if m != nil && m.NextHookAt != nil {
		return *m.NextHookAt
	}
	return 0
}

func (m *Webhook) GetLastHookAt() int64 {
	if m != nil && m.LastHookAt != nil {
		return *m.LastHookAt
	}
	return 0
}

func (m *Webhook) GetLastHookId() int64 {
	if m != nil && m.LastHookId != nil {
		return *m.LastHookId
	}
	return 0
}

func (m *Webhook) GetLastHookResponse() string {
	if m != nil && m.LastHookResponse != nil {
		return *m.LastHookResponse
	}
	return ""
}

func (m *Webhook) GetLastHookStatus() int32 {
	if m != nil && m.LastHookStatus != nil {
		return *m.LastHookStatus
	}
	return 0
}

func (m *Webhook) GetEventsCount() int64 {
	if m != nil && m.EventsCount != nil {
		return *m.EventsCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Retry)(nil), "webhook4.Retry")
	proto.RegisterType((*Webhook)(nil), "webhook4.Webhook")
	proto.RegisterEnum("webhook4.Event", Event_name, Event_value)
	proto.RegisterEnum("webhook4.Webhook_State", Webhook_State_name, Webhook_State_value)
}

func init() { proto.RegisterFile("webhook4/webhook4.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x80, 0x97, 0x65, 0x69, 0xd3, 0x5b, 0xd9, 0xc2, 0x89, 0x31, 0x43, 0x37, 0x48, 0xfb, 0x54,
	0x21, 0xd4, 0x4a, 0x68, 0x7f, 0x00, 0x2a, 0x24, 0xfa, 0x06, 0xa9, 0x80, 0xc7, 0xca, 0x49, 0xae,
	0x6b, 0xd6, 0xd4, 0xee, 0x1c, 0x67, 0x74, 0xfc, 0x0c, 0xfe, 0x19, 0xbf, 0x86, 0x57, 0xe4, 0x38,
	0x51, 0x60, 0xe2, 0x81, 0xa7, 0xe4, 0xbe, 0xfb, 0xec, 0xbb, 0xb3, 0x0d, 0xe7, 0xdf, 0x28, 0x5e,
	0x4b, 0xb9, 0xb9, 0x9a, 0x36, 0x3f, 0x93, 0x9d, 0x92, 0x5a, 0xa2, 0xdf, 0xc4, 0xcf, 0x27, 0x71,
	0xa6, 0xe3, 0x32, 0xd9, 0x90, 0x9e, 0x48, 0x75, 0x3d, 0x2d, 0xca, 0x38, 0xfb, 0x3e, 0x5d, 0x13,
	0x4f, 0x49, 0x4d, 0x13, 0xb9, 0xdd, 0x4a, 0x51, 0x7f, 0xec, 0xca, 0xd1, 0x0f, 0x07, 0xbc, 0x88,
	0xb4, 0xba, 0xc7, 0x21, 0xb8, 0x89, 0xde, 0x33, 0x27, 0x74, 0xc6, 0xc7, 0x6f, 0x4e, 0x27, 0xb5,
	0x35, 0x93, 0x42, 0xd3, 0x5e, 0x47, 0x26, 0x87, 0x17, 0xd0, 0xdb, 0x71, 0xa5, 0x33, 0x9d, 0x49,
	0xc1, 0x0e, 0x43, 0x67, 0xdc, 0x8b, 0x5a, 0x80, 0x4f, 0xc0, 0xbb, 0x2d, 0xa9, 0x24, 0xe6, 0x56,
	0x19, 0x1b, 0xe0, 0x19, 0x74, 0x6e, 0x64, 0xbc, 0xcc, 0x52, 0x76, 0x64, 0xf1, 0x8d, 0x8c, 0xe7,
	0x29, 0x0e, 0xa0, 0x27, 0x96, 0x31, 0x4f, 0x36, 0x72, 0xb5, 0x62, 0x5e, 0xe8, 0x8c, 0xbd, 0xc8,
	0x17, 0xef, 0x6c, 0x3c, 0xfa, 0xe9, 0x42, 0xf7, 0xab, 0x9d, 0xe8, 0x7f, 0xda, 0xba, 0x04, 0xe0,
	0x49, 0x22, 0x4b, 0xa1, 0x4d, 0x99, 0xbe, 0xed, 0xab, 0x26, 0xb6, 0x54, 0x92, 0x67, 0x64, 0xb3,
	0xb6, 0x6b, 0xdf, 0x82, 0x79, 0x8a, 0x01, 0xb8, 0xa5, 0xca, 0xeb, 0x96, 0xcd, 0x2f, 0x3e, 0x85,
	0x4e, 0x41, 0x89, 0x22, 0x5d, 0x37, 0x5c, 0x47, 0x86, 0xd3, 0x1d, 0x09, 0x5d, 0xb0, 0x4e, 0xe8,
	0x1a, 0x6e, 0x23, 0x33, 0x76, 0xa1, 0xb9, 0x26, 0xd6, 0xb5, 0xf3, 0x55, 0x01, 0x86, 0xd0, 0x17,
	0xb4, 0xd7, 0x4b, 0x33, 0xc3, 0x92, 0x6b, 0xe6, 0x87, 0xce, 0xd8, 0x8d, 0xc0, 0xb0, 0x0f, 0x52,
	0x6e, 0xde, 0x6a, 0x63, 0xe4, 0xbc, 0x68, 0x8d, 0x9e, 0x35, 0x0c, 0xfb, 0x97, 0x91, 0xa5, 0x0c,
	0xfe, 0x36, 0xe6, 0x29, 0xbe, 0x06, 0x6c, 0x0d, 0x45, 0xc5, 0x4e, 0x8a, 0x82, 0xd8, 0x49, 0xd5,
	0x48, 0xd0, 0x78, 0x51, 0xcd, 0x71, 0x0c, 0x41, 0x6b, 0x9b, 0x36, 0xcb, 0x82, 0x9d, 0x56, 0x47,
	0x7f, 0xd2, 0xb8, 0x8b, 0x8a, 0xe2, 0x10, 0xfa, 0x76, 0xba, 0x65, 0x75, 0x88, 0xec, 0x51, 0x55,
	0xf9, 0xd8, 0xb2, 0x99, 0x41, 0xa3, 0x09, 0x78, 0x8b, 0x6a, 0x52, 0x80, 0x8e, 0x90, 0x6a, 0xcb,
	0xf3, 0xe0, 0x00, 0x8f, 0xa1, 0x5b, 0xdf, 0x69, 0xe0, 0x60, 0x1f, 0xfc, 0x4c, 0xf0, 0x44, 0x67,
	0x77, 0x14, 0x1c, 0xbe, 0xfa, 0xe5, 0x80, 0xf7, 0xde, 0xac, 0xc7, 0x67, 0x70, 0x56, 0x5f, 0xee,
	0x55, 0xf5, 0xf2, 0x66, 0x3c, 0xcf, 0xcd, 0xaa, 0xe0, 0x00, 0x43, 0xb8, 0x68, 0x52, 0xf5, 0x37,
	0x22, 0x9e, 0x46, 0x74, 0x5b, 0x52, 0xa1, 0x29, 0x0d, 0x5c, 0x1c, 0xc0, 0xf9, 0x03, 0xe3, 0xf3,
	0xae, 0x20, 0x65, 0x92, 0x47, 0x38, 0x82, 0x17, 0x0f, 0x92, 0x5f, 0x48, 0x65, 0xab, 0xfb, 0x76,
	0x03, 0x0f, 0xcf, 0xe0, 0x71, 0x5b, 0xbd, 0xc1, 0x1d, 0x1c, 0xc2, 0x65, 0x83, 0x67, 0x39, 0x71,
	0x55, 0x3f, 0xc5, 0x56, 0xe9, 0xe2, 0x4b, 0x18, 0x34, 0xca, 0xc7, 0x52, 0x5d, 0xd3, 0x27, 0xf3,
	0xbe, 0x5b, 0xc1, 0xff, 0x73, 0xeb, 0xc5, 0x9a, 0xab, 0x74, 0x41, 0x22, 0x0d, 0x7a, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xd4, 0x07, 0xa6, 0xe9, 0xb6, 0x03, 0x00, 0x00,
}
