// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logan/logan.proto

package logan

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

func (Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d0d3176839f1211b, []int{0}
}

type Debug struct {
	StackTrace           []byte     `protobuf:"bytes,5,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	Hostname             string     `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Kafka                *KafkaInfo `protobuf:"bytes,23,opt,name=kafka,proto3" json:"kafka,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Debug) Reset()         { *m = Debug{} }
func (m *Debug) String() string { return proto.CompactTextString(m) }
func (*Debug) ProtoMessage()    {}
func (*Debug) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0d3176839f1211b, []int{0}
}

func (m *Debug) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Debug.Unmarshal(m, b)
}
func (m *Debug) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Debug.Marshal(b, m, deterministic)
}
func (m *Debug) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Debug.Merge(m, src)
}
func (m *Debug) XXX_Size() int {
	return xxx_messageInfo_Debug.Size(m)
}
func (m *Debug) XXX_DiscardUnknown() {
	xxx_messageInfo_Debug.DiscardUnknown(m)
}

var xxx_messageInfo_Debug proto.InternalMessageInfo

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
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Partition            int32    `protobuf:"varint,3,opt,name=partition,proto3" json:"partition,omitempty"`
	Offset               int64    `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaInfo) Reset()         { *m = KafkaInfo{} }
func (m *KafkaInfo) String() string { return proto.CompactTextString(m) }
func (*KafkaInfo) ProtoMessage()    {}
func (*KafkaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0d3176839f1211b, []int{1}
}

func (m *KafkaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaInfo.Unmarshal(m, b)
}
func (m *KafkaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaInfo.Marshal(b, m, deterministic)
}
func (m *KafkaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaInfo.Merge(m, src)
}
func (m *KafkaInfo) XXX_Size() int {
	return xxx_messageInfo_KafkaInfo.Size(m)
}
func (m *KafkaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaInfo proto.InternalMessageInfo

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
	Ctx                  *common.Context `protobuf:"bytes,1,opt,name=ctx,proto3" json:"ctx,omitempty"`
	TraceId              string          `protobuf:"bytes,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Created              int64           `protobuf:"varint,8,opt,name=created,proto3" json:"created,omitempty"`
	Level                string          `protobuf:"bytes,10,opt,name=level,proto3" json:"level,omitempty"`
	Tags                 []string        `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	Debug                *Debug          `protobuf:"bytes,22,opt,name=debug,proto3" json:"debug,omitempty"`
	Message              []byte          `protobuf:"bytes,24,opt,name=message,proto3" json:"message,omitempty"`
	ServiceName          string          `protobuf:"bytes,25,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0d3176839f1211b, []int{2}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

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
	proto.RegisterEnum("logan.Level", Level_name, Level_value)
	proto.RegisterType((*Debug)(nil), "logan.Debug")
	proto.RegisterType((*KafkaInfo)(nil), "logan.KafkaInfo")
	proto.RegisterType((*Log)(nil), "logan.Log")
}

func init() { proto.RegisterFile("logan/logan.proto", fileDescriptor_d0d3176839f1211b) }

var fileDescriptor_d0d3176839f1211b = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x14, 0x24, 0x9b, 0x8f, 0x26, 0xaf, 0x45, 0x98, 0x27, 0xb4, 0x78, 0x57, 0x48, 0x64, 0x7b, 0x40,
	0x11, 0x48, 0xad, 0x04, 0x3f, 0x01, 0x2e, 0x2b, 0x56, 0x1c, 0x22, 0x24, 0x8e, 0x95, 0xeb, 0xbc,
	0xa4, 0x56, 0x13, 0xbb, 0x72, 0xdc, 0x65, 0xe1, 0xc0, 0xef, 0xe5, 0x67, 0x20, 0x3b, 0xd9, 0xe5,
	0x92, 0x78, 0x66, 0x92, 0x37, 0xf3, 0x46, 0x86, 0x97, 0xbd, 0xe9, 0x84, 0xde, 0x86, 0xe7, 0xe6,
	0x64, 0x8d, 0x33, 0x98, 0x06, 0x70, 0xfd, 0xa1, 0x53, 0xee, 0x70, 0xde, 0x6f, 0xa4, 0x19, 0xb6,
	0xe3, 0x79, 0xaf, 0x7e, 0x6f, 0x0f, 0x24, 0x1a, 0xb2, 0x5b, 0x69, 0x86, 0xc1, 0xe8, 0xf9, 0x35,
	0xfd, 0xb3, 0xee, 0x21, 0xfd, 0x42, 0xfb, 0x73, 0x87, 0x6f, 0x61, 0x39, 0x3a, 0x21, 0x8f, 0x3b,
	0x67, 0x85, 0x24, 0x9e, 0x96, 0x51, 0xb5, 0xaa, 0x21, 0x50, 0xdf, 0x3d, 0x83, 0xd7, 0x90, 0x1f,
	0xcc, 0xe8, 0xb4, 0x18, 0x88, 0x67, 0x65, 0x54, 0x15, 0xf5, 0x13, 0xc6, 0x77, 0x90, 0x1e, 0x45,
	0x7b, 0x14, 0xfc, 0x75, 0x19, 0x55, 0xcb, 0x8f, 0x6c, 0x33, 0xc5, 0xfa, 0xea, 0xb9, 0x5b, 0xdd,
	0x9a, 0x7a, 0x92, 0xd7, 0x3f, 0xa0, 0x78, 0xe2, 0xf0, 0x15, 0xa4, 0xce, 0x9c, 0x94, 0xe4, 0x17,
	0x61, 0xda, 0x04, 0xf0, 0x0d, 0x14, 0x27, 0x61, 0x9d, 0x72, 0xca, 0x68, 0x1e, 0x97, 0x51, 0x95,
	0xd6, 0xff, 0x09, 0xbc, 0x84, 0xcc, 0xb4, 0xed, 0x48, 0x8e, 0x27, 0x65, 0x54, 0xc5, 0xf5, 0x8c,
	0xd6, 0x7f, 0x23, 0x88, 0xef, 0x4c, 0x87, 0x37, 0x10, 0x4b, 0xf7, 0xc0, 0xa3, 0x10, 0xe3, 0xc5,
	0x66, 0x5e, 0xf5, 0xb3, 0xd1, 0x8e, 0x1e, 0x5c, 0xed, 0x35, 0xbc, 0x82, 0x3c, 0xac, 0xb8, 0x53,
	0xcd, 0xec, 0xbc, 0x08, 0xf8, 0xb6, 0x41, 0x0e, 0x0b, 0x69, 0x49, 0x38, 0x6a, 0x78, 0x1e, 0xc6,
	0x3f, 0x42, 0x9f, 0xb5, 0xa7, 0x7b, 0xea, 0x39, 0x4c, 0x59, 0x03, 0x40, 0x84, 0xc4, 0x89, 0x6e,
	0xe4, 0x49, 0x19, 0x57, 0x45, 0x1d, 0xce, 0xb8, 0x86, 0xb4, 0xf1, 0x85, 0xf2, 0xcb, 0x90, 0x61,
	0x35, 0x57, 0x11, 0x4a, 0xae, 0x27, 0xc9, 0xfb, 0x0c, 0x34, 0x8e, 0xa2, 0x23, 0xce, 0x43, 0xcf,
	0x8f, 0x10, 0x6f, 0x60, 0x35, 0x92, 0xbd, 0x57, 0x92, 0x76, 0xa1, 0xe8, 0xab, 0x60, 0xb7, 0x9c,
	0xb9, 0x6f, 0x62, 0xa0, 0xf7, 0x7f, 0x20, 0xbd, 0x0b, 0xee, 0xc5, 0xec, 0xc4, 0x9e, 0x61, 0x0e,
	0x89, 0xd2, 0xad, 0x61, 0x11, 0x02, 0x64, 0xda, 0x38, 0x25, 0x89, 0x5d, 0xe0, 0x12, 0x16, 0x3f,
	0x85, 0xd5, 0x4a, 0x77, 0x2c, 0xf6, 0x5f, 0x93, 0xb5, 0xc6, 0xb2, 0x04, 0x57, 0x90, 0x4b, 0xab,
	0x9c, 0x92, 0xa2, 0x67, 0xa9, 0x17, 0x44, 0x4f, 0xd6, 0xb1, 0x0c, 0x9f, 0x43, 0x41, 0x03, 0xd9,
	0x8e, 0xb4, 0xfc, 0xc5, 0x16, 0x5e, 0x39, 0x09, 0xad, 0x24, 0xcb, 0xfd, 0xb1, 0x15, 0x4e, 0xf4,
	0xac, 0xd8, 0x67, 0xe1, 0xe2, 0x7c, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x77, 0x8e, 0xbd,
	0x81, 0x02, 0x00, 0x00,
}
