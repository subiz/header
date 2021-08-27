// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.2
// source: type.proto

package header

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

type TextTransform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // trim, lower_case, upper_case
}

func (x *TextTransform) Reset() {
	*x = TextTransform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextTransform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextTransform) ProtoMessage() {}

func (x *TextTransform) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextTransform.ProtoReflect.Descriptor instead.
func (*TextTransform) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{0}
}

func (x *TextTransform) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FloatTransform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // round, ceil, floor
}

func (x *FloatTransform) Reset() {
	*x = FloatTransform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloatTransform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatTransform) ProtoMessage() {}

func (x *FloatTransform) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatTransform.ProtoReflect.Descriptor instead.
func (*FloatTransform) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{1}
}

func (x *FloatTransform) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TextCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op              string           `protobuf:"bytes,2,opt,name=op,proto3" json:"op,omitempty"`
	Transforms      []*TextTransform `protobuf:"bytes,3,rep,name=transforms,proto3" json:"transforms,omitempty"`
	HasValue        bool             `protobuf:"varint,4,opt,name=has_value,json=hasValue,proto3" json:"has_value,omitempty"` // true or false
	Empty           bool             `protobuf:"varint,5,opt,name=empty,proto3" json:"empty,omitempty"`                       // true or false
	Eq              []string         `protobuf:"bytes,6,rep,name=eq,proto3" json:"eq,omitempty"`
	Neq             []string         `protobuf:"bytes,7,rep,name=neq,proto3" json:"neq,omitempty"`
	StartWith       []string         `protobuf:"bytes,8,rep,name=start_with,json=startWith,proto3" json:"start_with,omitempty"`
	EndWith         []string         `protobuf:"bytes,9,rep,name=end_with,json=endWith,proto3" json:"end_with,omitempty"`
	Contain         []string         `protobuf:"bytes,10,rep,name=contain,proto3" json:"contain,omitempty"`
	Regex           string           `protobuf:"bytes,12,opt,name=regex,proto3" json:"regex,omitempty"`
	NotContain      []string         `protobuf:"bytes,13,rep,name=not_contain,json=notContain,proto3" json:"not_contain,omitempty"`
	NotStartWith    []string         `protobuf:"bytes,15,rep,name=not_start_with,json=notStartWith,proto3" json:"not_start_with,omitempty"`
	NotEndWith      []string         `protobuf:"bytes,16,rep,name=not_end_with,json=notEndWith,proto3" json:"not_end_with,omitempty"`
	ContainAll      []string         `protobuf:"bytes,18,rep,name=contain_all,json=containAll,proto3" json:"contain_all,omitempty"`
	CaseSensitive   string           `protobuf:"bytes,30,opt,name=case_sensitive,json=caseSensitive,proto3" json:"case_sensitive,omitempty"`
	AccentSensitive string           `protobuf:"bytes,31,opt,name=accent_sensitive,json=accentSensitive,proto3" json:"accent_sensitive,omitempty"`
}

func (x *TextCondition) Reset() {
	*x = TextCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextCondition) ProtoMessage() {}

func (x *TextCondition) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextCondition.ProtoReflect.Descriptor instead.
func (*TextCondition) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{2}
}

func (x *TextCondition) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *TextCondition) GetTransforms() []*TextTransform {
	if x != nil {
		return x.Transforms
	}
	return nil
}

func (x *TextCondition) GetHasValue() bool {
	if x != nil {
		return x.HasValue
	}
	return false
}

func (x *TextCondition) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

func (x *TextCondition) GetEq() []string {
	if x != nil {
		return x.Eq
	}
	return nil
}

func (x *TextCondition) GetNeq() []string {
	if x != nil {
		return x.Neq
	}
	return nil
}

func (x *TextCondition) GetStartWith() []string {
	if x != nil {
		return x.StartWith
	}
	return nil
}

func (x *TextCondition) GetEndWith() []string {
	if x != nil {
		return x.EndWith
	}
	return nil
}

func (x *TextCondition) GetContain() []string {
	if x != nil {
		return x.Contain
	}
	return nil
}

func (x *TextCondition) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

func (x *TextCondition) GetNotContain() []string {
	if x != nil {
		return x.NotContain
	}
	return nil
}

func (x *TextCondition) GetNotStartWith() []string {
	if x != nil {
		return x.NotStartWith
	}
	return nil
}

func (x *TextCondition) GetNotEndWith() []string {
	if x != nil {
		return x.NotEndWith
	}
	return nil
}

func (x *TextCondition) GetContainAll() []string {
	if x != nil {
		return x.ContainAll
	}
	return nil
}

func (x *TextCondition) GetCaseSensitive() string {
	if x != nil {
		return x.CaseSensitive
	}
	return ""
}

func (x *TextCondition) GetAccentSensitive() string {
	if x != nil {
		return x.AccentSensitive
	}
	return ""
}

type BoolCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op string `protobuf:"bytes,2,opt,name=op,proto3" json:"op,omitempty"` // has_value, true, false
}

func (x *BoolCondition) Reset() {
	*x = BoolCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolCondition) ProtoMessage() {}

func (x *BoolCondition) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolCondition.ProtoReflect.Descriptor instead.
func (*BoolCondition) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{3}
}

func (x *BoolCondition) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

type FloatCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op         string            `protobuf:"bytes,2,opt,name=op,proto3" json:"op,omitempty"`
	Transforms []*FloatTransform `protobuf:"bytes,3,rep,name=transforms,proto3" json:"transforms,omitempty"`
	HasValue   bool              `protobuf:"varint,4,opt,name=has_value,json=hasValue,proto3" json:"has_value,omitempty"`
	Gt         float64           `protobuf:"fixed64,5,opt,name=gt,proto3" json:"gt,omitempty"`
	Gte        float64           `protobuf:"fixed64,6,opt,name=gte,proto3" json:"gte,omitempty"`
	Lt         float64           `protobuf:"fixed64,7,opt,name=lt,proto3" json:"lt,omitempty"`
	Lte        float64           `protobuf:"fixed64,8,opt,name=lte,proto3" json:"lte,omitempty"`
	In         []float64         `protobuf:"fixed64,9,rep,packed,name=in,proto3" json:"in,omitempty"`
	Eq         float64           `protobuf:"fixed64,10,opt,name=eq,proto3" json:"eq,omitempty"`
	Neq        float64           `protobuf:"fixed64,11,opt,name=neq,proto3" json:"neq,omitempty"`
	NotIn      []float64         `protobuf:"fixed64,12,rep,packed,name=not_in,json=notIn,proto3" json:"not_in,omitempty"`
	InRange    []float64         `protobuf:"fixed64,13,rep,packed,name=in_range,json=inRange,proto3" json:"in_range,omitempty"`
	NotInRange []float64         `protobuf:"fixed64,14,rep,packed,name=not_in_range,json=notInRange,proto3" json:"not_in_range,omitempty"`
}

func (x *FloatCondition) Reset() {
	*x = FloatCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloatCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatCondition) ProtoMessage() {}

func (x *FloatCondition) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatCondition.ProtoReflect.Descriptor instead.
func (*FloatCondition) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{4}
}

func (x *FloatCondition) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *FloatCondition) GetTransforms() []*FloatTransform {
	if x != nil {
		return x.Transforms
	}
	return nil
}

func (x *FloatCondition) GetHasValue() bool {
	if x != nil {
		return x.HasValue
	}
	return false
}

func (x *FloatCondition) GetGt() float64 {
	if x != nil {
		return x.Gt
	}
	return 0
}

func (x *FloatCondition) GetGte() float64 {
	if x != nil {
		return x.Gte
	}
	return 0
}

func (x *FloatCondition) GetLt() float64 {
	if x != nil {
		return x.Lt
	}
	return 0
}

func (x *FloatCondition) GetLte() float64 {
	if x != nil {
		return x.Lte
	}
	return 0
}

func (x *FloatCondition) GetIn() []float64 {
	if x != nil {
		return x.In
	}
	return nil
}

func (x *FloatCondition) GetEq() float64 {
	if x != nil {
		return x.Eq
	}
	return 0
}

func (x *FloatCondition) GetNeq() float64 {
	if x != nil {
		return x.Neq
	}
	return 0
}

func (x *FloatCondition) GetNotIn() []float64 {
	if x != nil {
		return x.NotIn
	}
	return nil
}

func (x *FloatCondition) GetInRange() []float64 {
	if x != nil {
		return x.InRange
	}
	return nil
}

func (x *FloatCondition) GetNotInRange() []float64 {
	if x != nil {
		return x.NotInRange
	}
	return nil
}

type DateCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op string `protobuf:"bytes,2,opt,name=op,proto3" json:"op,omitempty"` // in_business_hour, non_business_hour, relative, today, yesterday, last_week, this_week, last_month, this_month, last, before_ago
	// repeated FloatTransform transforms = 3;
	// repeated int64 hour_of_day
	// bool in_business_hour = 9;
	// bool non_business_hour = 10;
	DaysOfWeek []string `protobuf:"bytes,10,rep,name=days_of_week,json=daysOfWeek,proto3" json:"days_of_week,omitempty"` // monday
	After      int64    `protobuf:"varint,11,opt,name=after,proto3" json:"after,omitempty"`
	Before     int64    `protobuf:"varint,12,opt,name=before,proto3" json:"before,omitempty"`
	Between    []int64  `protobuf:"varint,13,rep,packed,name=between,proto3" json:"between,omitempty"`
	Outsite    []int64  `protobuf:"varint,14,rep,packed,name=outsite,proto3" json:"outsite,omitempty"`
	// relative, minute
	Last      int64 `protobuf:"varint,17,opt,name=last,proto3" json:"last,omitempty"`                            // 1440 => last 24h sec
	BeforeAgo int64 `protobuf:"varint,18,opt,name=before_ago,json=beforeAgo,proto3" json:"before_ago,omitempty"` // 365*86400 sec
}

func (x *DateCondition) Reset() {
	*x = DateCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateCondition) ProtoMessage() {}

func (x *DateCondition) ProtoReflect() protoreflect.Message {
	mi := &file_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateCondition.ProtoReflect.Descriptor instead.
func (*DateCondition) Descriptor() ([]byte, []int) {
	return file_type_proto_rawDescGZIP(), []int{5}
}

func (x *DateCondition) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *DateCondition) GetDaysOfWeek() []string {
	if x != nil {
		return x.DaysOfWeek
	}
	return nil
}

func (x *DateCondition) GetAfter() int64 {
	if x != nil {
		return x.After
	}
	return 0
}

func (x *DateCondition) GetBefore() int64 {
	if x != nil {
		return x.Before
	}
	return 0
}

func (x *DateCondition) GetBetween() []int64 {
	if x != nil {
		return x.Between
	}
	return nil
}

func (x *DateCondition) GetOutsite() []int64 {
	if x != nil {
		return x.Outsite
	}
	return nil
}

func (x *DateCondition) GetLast() int64 {
	if x != nil {
		return x.Last
	}
	return 0
}

func (x *DateCondition) GetBeforeAgo() int64 {
	if x != nil {
		return x.BeforeAgo
	}
	return 0
}

var File_type_proto protoreflect.FileDescriptor

var file_type_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x0d, 0x54, 0x65, 0x78, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xf1, 0x03, 0x0a, 0x0d, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f,
	0x70, 0x12, 0x35, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x54,
	0x65, 0x78, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x0a, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x65,
	0x71, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x65, 0x71, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x65, 0x71, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x57, 0x69, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x57, 0x69, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x6f,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x57, 0x69, 0x74, 0x68, 0x12, 0x20,
	0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x18, 0x10,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x45, 0x6e, 0x64, 0x57, 0x69, 0x74, 0x68,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x6c, 0x6c, 0x18,
	0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x41, 0x6c,
	0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x61, 0x73, 0x65, 0x53,
	0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x6f, 0x70, 0x22, 0xbf, 0x02, 0x0a, 0x0e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x36, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x68, 0x61, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x67, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x67, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x67, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x67, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x6c, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x74, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x01, 0x52, 0x02, 0x69, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x65, 0x71, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x65, 0x71, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6e,
	0x65, 0x71, 0x12, 0x15, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x01, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x49, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x01, 0x52, 0x07, 0x69, 0x6e, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x49,
	0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0xd6, 0x01, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x61, 0x79, 0x73,
	0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x65, 0x74, 0x77,
	0x65, 0x65, 0x6e, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x62, 0x65, 0x74, 0x77, 0x65,
	0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x73, 0x69, 0x74, 0x65, 0x18, 0x0e, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x73, 0x69, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x61, 0x73, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x67, 0x6f, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x41, 0x67, 0x6f, 0x42,
	0x19, 0x5a, 0x17, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75,
	0x62, 0x69, 0x7a, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_type_proto_rawDescOnce sync.Once
	file_type_proto_rawDescData = file_type_proto_rawDesc
)

func file_type_proto_rawDescGZIP() []byte {
	file_type_proto_rawDescOnce.Do(func() {
		file_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_type_proto_rawDescData)
	})
	return file_type_proto_rawDescData
}

var file_type_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_type_proto_goTypes = []interface{}{
	(*TextTransform)(nil),  // 0: header.TextTransform
	(*FloatTransform)(nil), // 1: header.FloatTransform
	(*TextCondition)(nil),  // 2: header.TextCondition
	(*BoolCondition)(nil),  // 3: header.BoolCondition
	(*FloatCondition)(nil), // 4: header.FloatCondition
	(*DateCondition)(nil),  // 5: header.DateCondition
}
var file_type_proto_depIdxs = []int32{
	0, // 0: header.TextCondition.transforms:type_name -> header.TextTransform
	1, // 1: header.FloatCondition.transforms:type_name -> header.FloatTransform
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_type_proto_init() }
func file_type_proto_init() {
	if File_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextTransform); i {
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
		file_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FloatTransform); i {
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
		file_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextCondition); i {
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
		file_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolCondition); i {
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
		file_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FloatCondition); i {
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
		file_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateCondition); i {
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
			RawDescriptor: file_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_type_proto_goTypes,
		DependencyIndexes: file_type_proto_depIdxs,
		MessageInfos:      file_type_proto_msgTypes,
	}.Build()
	File_type_proto = out.File
	file_type_proto_rawDesc = nil
	file_type_proto_goTypes = nil
	file_type_proto_depIdxs = nil
}
