// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logan.proto

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
	return fileDescriptor_40a4996c0d11ee01, []int{0}
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
	return fileDescriptor_40a4996c0d11ee01, []int{0}
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
	return fileDescriptor_40a4996c0d11ee01, []int{1}
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
	return fileDescriptor_40a4996c0d11ee01, []int{2}
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

func init() { proto.RegisterFile("logan.proto", fileDescriptor_40a4996c0d11ee01) }

var fileDescriptor_40a4996c0d11ee01 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x92, 0x41, 0x6f, 0xd4, 0x30,
	0x10, 0x85, 0x49, 0xb3, 0xd9, 0x4d, 0x26, 0x41, 0x58, 0x23, 0x54, 0xdc, 0x0a, 0xd4, 0x74, 0x0f,
	0x28, 0xe2, 0xd0, 0x95, 0xe0, 0x1f, 0x00, 0x97, 0x8a, 0x8a, 0x43, 0x84, 0x84, 0xc4, 0x65, 0xe5,
	0x75, 0x26, 0x59, 0x6b, 0x13, 0x7b, 0xe5, 0x78, 0x4b, 0xe1, 0xc0, 0xef, 0xe5, 0x67, 0x20, 0x3b,
	0x69, 0xb9, 0xf9, 0x7b, 0x2f, 0x9a, 0x37, 0xf3, 0x14, 0xc8, 0x7b, 0xd3, 0x09, 0x7d, 0x73, 0xb4,
	0xc6, 0x19, 0x4c, 0x02, 0x5c, 0x16, 0xd2, 0x0c, 0x83, 0x99, 0xc5, 0x75, 0x0f, 0xc9, 0x67, 0xda,
	0x9d, 0x3a, 0xbc, 0x82, 0x7c, 0x74, 0x42, 0x1e, 0xb6, 0xce, 0x0a, 0x49, 0x3c, 0x29, 0xa3, 0xaa,
	0xa8, 0x21, 0x48, 0xdf, 0xbc, 0x82, 0x97, 0x90, 0xee, 0xcd, 0xe8, 0xb4, 0x18, 0x88, 0x2f, 0xcb,
	0xa8, 0xca, 0xea, 0x27, 0xc6, 0xb7, 0x90, 0x1c, 0x44, 0x7b, 0x10, 0xfc, 0x55, 0x19, 0x55, 0xf9,
	0x7b, 0x76, 0x33, 0xe5, 0x7e, 0xf1, 0xda, 0xad, 0x6e, 0x4d, 0x3d, 0xd9, 0xeb, 0xef, 0x90, 0x3d,
	0x69, 0xf8, 0x12, 0x12, 0x67, 0x8e, 0x4a, 0xf2, 0xb3, 0x30, 0x6d, 0x02, 0x7c, 0x0d, 0xd9, 0x51,
	0x58, 0xa7, 0x9c, 0x32, 0x9a, 0xc7, 0x65, 0x54, 0x25, 0xf5, 0x7f, 0x01, 0xcf, 0x61, 0x69, 0xda,
	0x76, 0x24, 0xc7, 0x17, 0x65, 0x54, 0xc5, 0xf5, 0x4c, 0xeb, 0xbf, 0x11, 0xc4, 0x77, 0xa6, 0xc3,
	0x6b, 0x88, 0xa5, 0x7b, 0xe0, 0x51, 0x58, 0xe3, 0xc5, 0xcd, 0x7c, 0xea, 0x27, 0xa3, 0x1d, 0x3d,
	0xb8, 0xda, 0x7b, 0x78, 0x01, 0x69, 0x38, 0x71, 0xab, 0x9a, 0x39, 0x79, 0x15, 0xf8, 0xb6, 0x41,
	0x0e, 0x2b, 0x69, 0x49, 0x38, 0x6a, 0x78, 0x1a, 0xc6, 0x3f, 0xa2, 0xdf, 0xb5, 0xa7, 0x7b, 0xea,
	0x39, 0x4c, 0xbb, 0x06, 0x40, 0x84, 0x85, 0x13, 0xdd, 0xc8, 0x17, 0x65, 0x5c, 0x65, 0x75, 0x78,
	0xe3, 0x1a, 0x92, 0xc6, 0x17, 0xca, 0xcf, 0xc3, 0x0e, 0xc5, 0x5c, 0x45, 0x28, 0xb9, 0x9e, 0x2c,
	0x9f, 0x33, 0xd0, 0x38, 0x8a, 0x8e, 0x38, 0x0f, 0x3d, 0x3f, 0x22, 0x5e, 0x43, 0x31, 0x92, 0xbd,
	0x57, 0x92, 0xb6, 0xa1, 0xe8, 0x8b, 0x10, 0x97, 0xcf, 0xda, 0x57, 0x31, 0xd0, 0xbb, 0x3f, 0x90,
	0xdc, 0x85, 0xf4, 0x6c, 0x4e, 0x62, 0xcf, 0x30, 0x85, 0x85, 0xd2, 0xad, 0x61, 0x11, 0x02, 0x2c,
	0xb5, 0x71, 0x4a, 0x12, 0x3b, 0xc3, 0x1c, 0x56, 0x3f, 0x85, 0xd5, 0x4a, 0x77, 0x2c, 0xf6, 0x5f,
	0x93, 0xb5, 0xc6, 0xb2, 0x05, 0x16, 0x90, 0x4a, 0xab, 0x9c, 0x92, 0xa2, 0x67, 0x89, 0x37, 0x44,
	0x4f, 0xd6, 0xb1, 0x25, 0x3e, 0x87, 0x8c, 0x06, 0xb2, 0x1d, 0x69, 0xf9, 0x8b, 0xad, 0xbc, 0x73,
	0x14, 0x5a, 0x49, 0x96, 0xfa, 0x67, 0x2b, 0x9c, 0xe8, 0x59, 0xf6, 0xf1, 0xea, 0xc7, 0x9b, 0x4e,
	0xb9, 0xfd, 0x69, 0xe7, 0xdb, 0xdd, 0x8c, 0xa7, 0x9d, 0xfa, 0xbd, 0xd9, 0x93, 0x68, 0xc8, 0x6e,
	0xc2, 0xa9, 0xbb, 0x65, 0xf8, 0xb3, 0x3e, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x40, 0x83,
	0xbf, 0x7d, 0x02, 0x00, 0x00,
}
