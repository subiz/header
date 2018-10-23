// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logan/logan.proto

/*
Package logan is a generated protocol buffer package.

It is generated from these files:
	logan/logan.proto

It has these top-level messages:
	Debug
	KafkaInfo
	Log
*/
package logan

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "git.subiz.net/header/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Level int32

const (
	Level_debug     Level = 0
	Level_info      Level = 1
	Level_notice    Level = 2
	Level_warning   Level = 3
	Level_error     Level = 4
	Level_critical  Level = 5
	Level_alert     Level = 6
	Level_emergency Level = 7
	Level_panic     Level = 8
	Level_fatal     Level = 9
)

var Level_name = map[int32]string{
	0: "debug",
	1: "info",
	2: "notice",
	3: "warning",
	4: "error",
	5: "critical",
	6: "alert",
	7: "emergency",
	8: "panic",
	9: "fatal",
}
var Level_value = map[string]int32{
	"debug":     0,
	"info":      1,
	"notice":    2,
	"warning":   3,
	"error":     4,
	"critical":  5,
	"alert":     6,
	"emergency": 7,
	"panic":     8,
	"fatal":     9,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}
func (Level) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Event int32

const (
	Event_LogLogRequested Event = 0
	Event_LogRequested    Event = 1000
	Event_LogSynced       Event = 1001
)

var Event_name = map[int32]string{
	0:    "LogLogRequested",
	1000: "LogRequested",
	1001: "LogSynced",
}
var Event_value = map[string]int32{
	"LogLogRequested": 0,
	"LogRequested":    1000,
	"LogSynced":       1001,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Debug struct {
	StackTrace []byte     `protobuf:"bytes,5,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	Hostname   string     `protobuf:"bytes,6,opt,name=hostname" json:"hostname,omitempty"`
	Kafka      *KafkaInfo `protobuf:"bytes,23,opt,name=kafka" json:"kafka,omitempty"`
}

func (m *Debug) Reset()                    { *m = Debug{} }
func (m *Debug) String() string            { return proto.CompactTextString(m) }
func (*Debug) ProtoMessage()               {}
func (*Debug) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Debug) GetStackTrace() []byte {
	if m != nil {
		return m.StackTrace
	}
	return nil
}

func (m *Debug) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Debug) GetKafka() *KafkaInfo {
	if m != nil {
		return m.Kafka
	}
	return nil
}

type KafkaInfo struct {
	Topic     string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Partition int32  `protobuf:"varint,3,opt,name=partition" json:"partition,omitempty"`
	Offset    int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *KafkaInfo) Reset()                    { *m = KafkaInfo{} }
func (m *KafkaInfo) String() string            { return proto.CompactTextString(m) }
func (*KafkaInfo) ProtoMessage()               {}
func (*KafkaInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KafkaInfo) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *KafkaInfo) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *KafkaInfo) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Log struct {
	Ctx         *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	TraceId     string          `protobuf:"bytes,2,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	Created     int64           `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	Level       string          `protobuf:"bytes,10,opt,name=level" json:"level,omitempty"`
	Tags        []string        `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	Debug       *Debug          `protobuf:"bytes,22,opt,name=debug" json:"debug,omitempty"`
	Message     []byte          `protobuf:"bytes,24,opt,name=message,proto3" json:"message,omitempty"`
	ServiceName string          `protobuf:"bytes,25,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Log) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Log) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Log) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Log) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *Log) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Log) GetDebug() *Debug {
	if m != nil {
		return m.Debug
	}
	return nil
}

func (m *Log) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Log) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func init() {
	proto.RegisterType((*Debug)(nil), "logan.Debug")
	proto.RegisterType((*KafkaInfo)(nil), "logan.KafkaInfo")
	proto.RegisterType((*Log)(nil), "logan.Log")
	proto.RegisterEnum("logan.Level", Level_name, Level_value)
	proto.RegisterEnum("logan.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("logan/logan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x41, 0x6f, 0xd3, 0x4c,
	0x14, 0xac, 0xeb, 0x38, 0xb6, 0x5f, 0xf2, 0x7d, 0xdd, 0x3e, 0x50, 0xd9, 0x56, 0x48, 0xb8, 0x39,
	0x20, 0xab, 0x87, 0x54, 0x82, 0x1f, 0xc0, 0x01, 0x38, 0x54, 0x44, 0x1c, 0x0c, 0x12, 0xc7, 0x68,
	0xb3, 0x7e, 0x76, 0x57, 0xb1, 0x77, 0xc3, 0x7a, 0x13, 0x5a, 0x0e, 0xfc, 0xde, 0xf2, 0x2f, 0x90,
	0xd7, 0x6e, 0x11, 0x17, 0x7b, 0x67, 0xc6, 0x7e, 0x33, 0x6f, 0xb4, 0x70, 0xda, 0x98, 0x5a, 0xe8,
	0x6b, 0xff, 0x5c, 0xee, 0xac, 0x71, 0x06, 0x23, 0x0f, 0x2e, 0xf2, 0x5a, 0xb9, 0x65, 0xb7, 0xdf,
	0xa8, 0x9f, 0x4b, 0x4d, 0xee, 0xfa, 0x96, 0x44, 0x49, 0xf6, 0x5a, 0x9a, 0xb6, 0x35, 0x7a, 0x7c,
	0x0d, 0x3f, 0x2c, 0x1a, 0x88, 0x3e, 0xd0, 0x66, 0x5f, 0xe3, 0x2b, 0x98, 0x75, 0x4e, 0xc8, 0xed,
	0xda, 0x59, 0x21, 0x89, 0x47, 0x59, 0x90, 0xcf, 0x0b, 0xf0, 0xd4, 0xd7, 0x9e, 0xc1, 0x0b, 0x48,
	0x6e, 0x4d, 0xe7, 0xb4, 0x68, 0x89, 0x4f, 0xb3, 0x20, 0x4f, 0x8b, 0x27, 0x8c, 0xaf, 0x21, 0xda,
	0x8a, 0x6a, 0x2b, 0xf8, 0x8b, 0x2c, 0xc8, 0x67, 0x6f, 0xd8, 0x72, 0xc8, 0xf4, 0xa9, 0xe7, 0x6e,
	0x74, 0x65, 0x8a, 0x41, 0x5e, 0x7c, 0x83, 0xf4, 0x89, 0xc3, 0xe7, 0x10, 0x39, 0xb3, 0x53, 0x92,
	0x1f, 0xfb, 0x69, 0x03, 0xc0, 0x97, 0x90, 0xee, 0x84, 0x75, 0xca, 0x29, 0xa3, 0x79, 0x98, 0x05,
	0x79, 0x54, 0xfc, 0x25, 0xf0, 0x0c, 0xa6, 0xa6, 0xaa, 0x3a, 0x72, 0x7c, 0x92, 0x05, 0x79, 0x58,
	0x8c, 0x68, 0xf1, 0x10, 0x40, 0xb8, 0x32, 0x35, 0x5e, 0x42, 0x28, 0xdd, 0x1d, 0x0f, 0x7c, 0x8c,
	0x93, 0xe5, 0xb8, 0xea, 0x7b, 0xa3, 0x1d, 0xdd, 0xb9, 0xa2, 0xd7, 0xf0, 0x1c, 0x12, 0xbf, 0xe2,
	0x5a, 0x95, 0xa3, 0x73, 0xec, 0xf1, 0x4d, 0x89, 0x1c, 0x62, 0x69, 0x49, 0x38, 0x2a, 0x79, 0xe2,
	0xc7, 0x3f, 0xc2, 0x3e, 0x6b, 0x43, 0x07, 0x6a, 0x38, 0x0c, 0x59, 0x3d, 0x40, 0x84, 0x89, 0x13,
	0x75, 0xc7, 0x27, 0x59, 0x98, 0xa7, 0x85, 0x3f, 0xe3, 0x02, 0xa2, 0xb2, 0x2f, 0x94, 0x9f, 0xf9,
	0x0c, 0xf3, 0xb1, 0x0a, 0x5f, 0x72, 0x31, 0x48, 0xbd, 0x4f, 0x4b, 0x5d, 0x27, 0x6a, 0xe2, 0xdc,
	0xf7, 0xfc, 0x08, 0xf1, 0x12, 0xe6, 0x1d, 0xd9, 0x83, 0x92, 0xb4, 0xf6, 0x45, 0x9f, 0x7b, 0xbb,
	0xd9, 0xc8, 0x7d, 0x16, 0x2d, 0x5d, 0xfd, 0x82, 0x68, 0xe5, 0xdd, 0xd3, 0xd1, 0x89, 0x1d, 0x61,
	0x02, 0x13, 0xa5, 0x2b, 0xc3, 0x02, 0x04, 0x98, 0x6a, 0xe3, 0x94, 0x24, 0x76, 0x8c, 0x33, 0x88,
	0x7f, 0x08, 0xab, 0x95, 0xae, 0x59, 0xd8, 0x7f, 0x4d, 0xd6, 0x1a, 0xcb, 0x26, 0x38, 0x87, 0x44,
	0x5a, 0xe5, 0x94, 0x14, 0x0d, 0x8b, 0x7a, 0x41, 0x34, 0x64, 0x1d, 0x9b, 0xe2, 0x7f, 0x90, 0x52,
	0x4b, 0xb6, 0x26, 0x2d, 0xef, 0x59, 0xdc, 0x2b, 0x3b, 0xa1, 0x95, 0x64, 0x49, 0x7f, 0xac, 0x84,
	0x13, 0x0d, 0x4b, 0xaf, 0xde, 0x41, 0xf4, 0xf1, 0x40, 0xda, 0xe1, 0x33, 0x38, 0x59, 0x99, 0x7a,
	0x65, 0xea, 0x82, 0xbe, 0xef, 0xa9, 0x73, 0x54, 0xb2, 0x23, 0x3c, 0x85, 0xf9, 0x3f, 0xcc, 0x43,
	0x8c, 0xff, 0x43, 0xba, 0x32, 0xf5, 0x97, 0x7b, 0x2d, 0xa9, 0x64, 0xbf, 0xe3, 0xcd, 0xd4, 0xdf,
	0xbc, 0xb7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x35, 0xf2, 0x8a, 0xbf, 0x02, 0x00, 0x00,
}
