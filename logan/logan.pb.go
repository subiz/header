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
	MemStats
*/
package logan

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
	MemStats     *MemStats  `protobuf:"bytes,2,opt,name=mem_stats,json=memStats" json:"mem_stats,omitempty"`
	NumCpu       int32      `protobuf:"varint,3,opt,name=num_cpu,json=numCpu" json:"num_cpu,omitempty"`
	NumGoroutine int32      `protobuf:"varint,4,opt,name=num_goroutine,json=numGoroutine" json:"num_goroutine,omitempty"`
	StackTrace   []byte     `protobuf:"bytes,5,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	Hostname     string     `protobuf:"bytes,6,opt,name=hostname" json:"hostname,omitempty"`
	Kafka        *KafkaInfo `protobuf:"bytes,23,opt,name=kafka" json:"kafka,omitempty"`
}

func (m *Debug) Reset()                    { *m = Debug{} }
func (m *Debug) String() string            { return proto.CompactTextString(m) }
func (*Debug) ProtoMessage()               {}
func (*Debug) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Debug) GetMemStats() *MemStats {
	if m != nil {
		return m.MemStats
	}
	return nil
}

func (m *Debug) GetNumCpu() int32 {
	if m != nil {
		return m.NumCpu
	}
	return 0
}

func (m *Debug) GetNumGoroutine() int32 {
	if m != nil {
		return m.NumGoroutine
	}
	return 0
}

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
	Ctx     *common.Context `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	TraceId string          `protobuf:"bytes,2,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	Created int64           `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	Level   string          `protobuf:"bytes,10,opt,name=level" json:"level,omitempty"`
	Tags    []string        `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	Debug   *Debug          `protobuf:"bytes,22,opt,name=debug" json:"debug,omitempty"`
	Message []byte          `protobuf:"bytes,24,opt,name=message,proto3" json:"message,omitempty"`
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

type MemStats struct {
	Alloc uint64 `protobuf:"varint,2,opt,name=alloc" json:"alloc,omitempty"`
	// TotalAlloc is cumulative bytes allocated for heap objects.
	//
	// TotalAlloc increases as heap objects are allocated, but
	// unlike Alloc and HeapAlloc, it does not decrease when
	// objects are freed.
	TotalAlloc uint64 `protobuf:"varint,3,opt,name=total_alloc,json=totalAlloc" json:"total_alloc,omitempty"`
	// Sys is the total bytes of memory obtained from the OS.
	//
	// Sys is the sum of the XSys fields below. Sys measures the
	// virtual address space reserved by the Go runtime for the
	// heap, stacks, and other internal data structures. It's
	// likely that not all of the virtual address space is backed
	// by physical memory at any given moment, though in general
	// it all was at some point.
	Sys uint64 `protobuf:"varint,4,opt,name=sys" json:"sys,omitempty"`
	// Lookups is the number of pointer lookups performed by the
	// runtime.
	//
	// This is primarily useful for debugging runtime internals.
	Lookups uint64 `protobuf:"varint,5,opt,name=lookups" json:"lookups,omitempty"`
	// Mallocs is the cumulative count of heap objects allocated.
	// The number of live objects is Mallocs - Frees.
	Mallocs uint64 `protobuf:"varint,6,opt,name=mallocs" json:"mallocs,omitempty"`
	// Frees is the cumulative count of heap objects freed.
	Frees uint64 `protobuf:"varint,7,opt,name=frees" json:"frees,omitempty"`
	// HeapAlloc is bytes of allocated heap objects.
	//
	// "Allocated" heap objects include all reachable objects, as
	// well as unreachable objects that the garbage collector has
	// not yet freed. Specifically, HeapAlloc increases as heap
	// objects are allocated and decreases as the heap is swept
	// and unreachable objects are freed. Sweeping occurs
	// incrementally between GC cycles, so these two processes
	// occur simultaneously, and as a result HeapAlloc tends to
	// change smoothly (in contrast with the sawtooth that is
	// typical of stop-the-world garbage collectors).
	HeapAlloc uint64 `protobuf:"varint,8,opt,name=heap_alloc,json=heapAlloc" json:"heap_alloc,omitempty"`
	// HeapSys is bytes of heap memory obtained from the OS.
	//
	// HeapSys measures the amount of virtual address space
	// reserved for the heap. This includes virtual address space
	// that has been reserved but not yet used, which consumes no
	// physical memory, but tends to be small, as well as virtual
	// address space for which the physical memory has been
	// returned to the OS after it became unused (see HeapReleased
	// for a measure of the latter).
	//
	// HeapSys estimates the largest size the heap has had.
	HeapSys uint64 `protobuf:"varint,9,opt,name=heap_sys,json=heapSys" json:"heap_sys,omitempty"`
	// HeapIdle is bytes in idle (unused) spans.
	//
	// Idle spans have no objects in them. These spans could be
	// (and may already have been) returned to the OS, or they can
	// be reused for heap allocations, or they can be reused as
	// stack memory.
	//
	// HeapIdle minus HeapReleased estimates the amount of memory
	// that could be returned to the OS, but is being retained by
	// the runtime so it can grow the heap without requesting more
	// memory from the OS. If this difference is significantly
	// larger than the heap size, it indicates there was a recent
	// transient spike in live heap size.
	HeapIdle uint64 `protobuf:"varint,10,opt,name=heap_idle,json=heapIdle" json:"heap_idle,omitempty"`
	// HeapInuse is bytes in in-use spans.
	//
	// In-use spans have at least one object in them. These spans
	// can only be used for other objects of roughly the same
	// size.
	//
	// HeapInuse minus HeapAlloc esimates the amount of memory
	// that has been dedicated to particular size classes, but is
	// not currently being used. This is an upper bound on
	// fragmentation, but in general this memory can be reused
	// efficiently.
	HeapInuse uint64 `protobuf:"varint,11,opt,name=heap_inuse,json=heapInuse" json:"heap_inuse,omitempty"`
	// HeapReleased is bytes of physical memory returned to the OS.
	//
	// This counts heap memory from idle spans that was returned
	// to the OS and has not yet been reacquired for the heap.
	HeapReleased uint64 `protobuf:"varint,12,opt,name=heap_released,json=heapReleased" json:"heap_released,omitempty"`
	// HeapObjects is the number of allocated heap objects.
	//
	// Like HeapAlloc, this increases as objects are allocated and
	// decreases as the heap is swept and unreachable objects are
	// freed.
	HeapObjects uint64 `protobuf:"varint,13,opt,name=heap_objects,json=heapObjects" json:"heap_objects,omitempty"`
	// StackInuse is bytes in stack spans.
	//
	// In-use stack spans have at least one stack in them. These
	// spans can only be used for other stacks of the same size.
	//
	// There is no StackIdle because unused stack spans are
	// returned to the heap (and hence counted toward HeapIdle).
	StackInuse uint64 `protobuf:"varint,14,opt,name=stack_inuse,json=stackInuse" json:"stack_inuse,omitempty"`
	// StackSys is bytes of stack memory obtained from the OS.
	//
	// StackSys is StackInuse, plus any memory obtained directly
	// from the OS for OS thread stacks (which should be minimal).
	StackSys uint64 `protobuf:"varint,15,opt,name=stack_sys,json=stackSys" json:"stack_sys,omitempty"`
	// MSpanInuse is bytes of allocated mspan structures.
	MSpanInuse uint64 `protobuf:"varint,16,opt,name=m_span_inuse,json=mSpanInuse" json:"m_span_inuse,omitempty"`
	// MSpanSys is bytes of memory obtained from the OS for mspan
	// structures.
	MSpanSys uint64 `protobuf:"varint,17,opt,name=m_span_sys,json=mSpanSys" json:"m_span_sys,omitempty"`
	// MCacheInuse is bytes of allocated mcache structures.
	MCacheInuse uint64 `protobuf:"varint,18,opt,name=m_cache_inuse,json=mCacheInuse" json:"m_cache_inuse,omitempty"`
	// MCacheSys is bytes of memory obtained from the OS for
	// mcache structures.
	MCacheSys uint64 `protobuf:"varint,19,opt,name=m_cache_sys,json=mCacheSys" json:"m_cache_sys,omitempty"`
	// BuckHashSys is bytes of memory in profiling bucket hash tables.
	BuckHashSys uint64 `protobuf:"varint,20,opt,name=buck_hash_sys,json=buckHashSys" json:"buck_hash_sys,omitempty"`
	// GCSys is bytes of memory in garbage collection metadata.
	GcSys uint64 `protobuf:"varint,21,opt,name=gc_sys,json=gcSys" json:"gc_sys,omitempty"`
	// OtherSys is bytes of memory in miscellaneous off-heap
	// runtime allocations.
	OtherSys uint64 `protobuf:"varint,22,opt,name=other_sys,json=otherSys" json:"other_sys,omitempty"`
	// NextGC is the target heap size of the next GC cycle.
	//
	// The garbage collector's goal is to keep HeapAlloc ≤ NextGC.
	// At the end of each GC cycle, the target for the next cycle
	// is computed based on the amount of reachable data and the
	// value of GOGC.
	NextGc uint64 `protobuf:"varint,23,opt,name=next_gc,json=nextGc" json:"next_gc,omitempty"`
	// LastGC is the time the last garbage collection finished, as
	// nanoseconds since 1970 (the UNIX epoch).
	LastGc uint64 `protobuf:"varint,24,opt,name=last_gc,json=lastGc" json:"last_gc,omitempty"`
	// PauseTotalNs is the cumulative nanoseconds in GC
	// stop-the-world pauses since the program started.
	//
	// During a stop-the-world pause, all goroutines are paused
	// and only the garbage collector can run.
	PauseTotalNs uint64 `protobuf:"varint,25,opt,name=pause_total_ns,json=pauseTotalNs" json:"pause_total_ns,omitempty"`
	// NumGC is the number of completed GC cycles.
	NumGc uint32 `protobuf:"varint,26,opt,name=num_gc,json=numGc" json:"num_gc,omitempty"`
	// NumForcedGC is the number of GC cycles that were forced by
	// the application calling the GC function.
	NumForcedGc uint32 `protobuf:"varint,27,opt,name=num_forced_gc,json=numForcedGc" json:"num_forced_gc,omitempty"`
	// GCCPUFraction is the fraction of this program's available
	// CPU time used by the GC since the program started.
	//
	// GCCPUFraction is expressed as a number between 0 and 1,
	// where 0 means GC has consumed none of this program's CPU. A
	// program's available CPU time is defined as the integral of
	// GOMAXPROCS since the program started. That is, if
	// GOMAXPROCS is 2 and a program has been running for 10
	// seconds, its "available CPU" is 20 seconds. GCCPUFraction
	// does not include CPU time used for write barrier activity.
	//
	// This is the same as the fraction of CPU reported by
	// GODEBUG=gctrace=1.
	GcCpuFraction float64 `protobuf:"fixed64,28,opt,name=gc_cpu_fraction,json=gcCpuFraction" json:"gc_cpu_fraction,omitempty"`
}

func (m *MemStats) Reset()                    { *m = MemStats{} }
func (m *MemStats) String() string            { return proto.CompactTextString(m) }
func (*MemStats) ProtoMessage()               {}
func (*MemStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MemStats) GetAlloc() uint64 {
	if m != nil {
		return m.Alloc
	}
	return 0
}

func (m *MemStats) GetTotalAlloc() uint64 {
	if m != nil {
		return m.TotalAlloc
	}
	return 0
}

func (m *MemStats) GetSys() uint64 {
	if m != nil {
		return m.Sys
	}
	return 0
}

func (m *MemStats) GetLookups() uint64 {
	if m != nil {
		return m.Lookups
	}
	return 0
}

func (m *MemStats) GetMallocs() uint64 {
	if m != nil {
		return m.Mallocs
	}
	return 0
}

func (m *MemStats) GetFrees() uint64 {
	if m != nil {
		return m.Frees
	}
	return 0
}

func (m *MemStats) GetHeapAlloc() uint64 {
	if m != nil {
		return m.HeapAlloc
	}
	return 0
}

func (m *MemStats) GetHeapSys() uint64 {
	if m != nil {
		return m.HeapSys
	}
	return 0
}

func (m *MemStats) GetHeapIdle() uint64 {
	if m != nil {
		return m.HeapIdle
	}
	return 0
}

func (m *MemStats) GetHeapInuse() uint64 {
	if m != nil {
		return m.HeapInuse
	}
	return 0
}

func (m *MemStats) GetHeapReleased() uint64 {
	if m != nil {
		return m.HeapReleased
	}
	return 0
}

func (m *MemStats) GetHeapObjects() uint64 {
	if m != nil {
		return m.HeapObjects
	}
	return 0
}

func (m *MemStats) GetStackInuse() uint64 {
	if m != nil {
		return m.StackInuse
	}
	return 0
}

func (m *MemStats) GetStackSys() uint64 {
	if m != nil {
		return m.StackSys
	}
	return 0
}

func (m *MemStats) GetMSpanInuse() uint64 {
	if m != nil {
		return m.MSpanInuse
	}
	return 0
}

func (m *MemStats) GetMSpanSys() uint64 {
	if m != nil {
		return m.MSpanSys
	}
	return 0
}

func (m *MemStats) GetMCacheInuse() uint64 {
	if m != nil {
		return m.MCacheInuse
	}
	return 0
}

func (m *MemStats) GetMCacheSys() uint64 {
	if m != nil {
		return m.MCacheSys
	}
	return 0
}

func (m *MemStats) GetBuckHashSys() uint64 {
	if m != nil {
		return m.BuckHashSys
	}
	return 0
}

func (m *MemStats) GetGcSys() uint64 {
	if m != nil {
		return m.GcSys
	}
	return 0
}

func (m *MemStats) GetOtherSys() uint64 {
	if m != nil {
		return m.OtherSys
	}
	return 0
}

func (m *MemStats) GetNextGc() uint64 {
	if m != nil {
		return m.NextGc
	}
	return 0
}

func (m *MemStats) GetLastGc() uint64 {
	if m != nil {
		return m.LastGc
	}
	return 0
}

func (m *MemStats) GetPauseTotalNs() uint64 {
	if m != nil {
		return m.PauseTotalNs
	}
	return 0
}

func (m *MemStats) GetNumGc() uint32 {
	if m != nil {
		return m.NumGc
	}
	return 0
}

func (m *MemStats) GetNumForcedGc() uint32 {
	if m != nil {
		return m.NumForcedGc
	}
	return 0
}

func (m *MemStats) GetGcCpuFraction() float64 {
	if m != nil {
		return m.GcCpuFraction
	}
	return 0
}

func init() {
	proto.RegisterType((*Debug)(nil), "logan.Debug")
	proto.RegisterType((*KafkaInfo)(nil), "logan.KafkaInfo")
	proto.RegisterType((*Log)(nil), "logan.Log")
	proto.RegisterType((*MemStats)(nil), "logan.MemStats")
	proto.RegisterEnum("logan.Level", Level_name, Level_value)
	proto.RegisterEnum("logan.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("logan/logan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 910 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x54, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x5e, 0x6f, 0xec, 0x24, 0x3e, 0x49, 0xb6, 0xde, 0xd9, 0x3f, 0xf7, 0x07, 0xc8, 0x06, 0xb4,
	0x8a, 0x56, 0xa8, 0x95, 0xe0, 0x01, 0x10, 0x2a, 0x6c, 0xa9, 0x28, 0x20, 0xb9, 0x2b, 0x71, 0x69,
	0x4d, 0xc6, 0x27, 0x8e, 0xa9, 0x3d, 0x63, 0x3c, 0xe3, 0xa5, 0xe5, 0x82, 0x67, 0xe2, 0x25, 0x78,
	0x07, 0x1e, 0x01, 0xde, 0x02, 0x9d, 0x33, 0x4e, 0x2b, 0x6e, 0x12, 0x7f, 0xdf, 0x77, 0xfe, 0x8f,
	0x7d, 0xe0, 0x69, 0x6d, 0x4a, 0xa9, 0xcf, 0xf8, 0xf7, 0xb4, 0xed, 0x8c, 0x33, 0x22, 0x62, 0x70,
	0x74, 0xba, 0xa9, 0xdc, 0xa6, 0x57, 0x37, 0xe8, 0x4e, 0x4d, 0x57, 0x9e, 0xd9, 0x7e, 0x53, 0xfd,
	0x7e, 0xb6, 0x43, 0x59, 0x60, 0x77, 0xa6, 0x4c, 0xd3, 0x18, 0x3d, 0xfc, 0x79, 0xb7, 0xd5, 0xdf,
	0x01, 0x44, 0xdf, 0xe0, 0xa6, 0x2f, 0xc5, 0xe7, 0x10, 0x37, 0xd8, 0xe4, 0xd6, 0x49, 0x67, 0xd3,
	0xc7, 0xcb, 0x60, 0x3d, 0xfb, 0xe2, 0xe0, 0xd4, 0x67, 0xf8, 0x01, 0x9b, 0x6b, 0xa2, 0xb3, 0x69,
	0x33, 0x3c, 0x89, 0x57, 0x30, 0xd1, 0x7d, 0x93, 0xab, 0xb6, 0x4f, 0x47, 0xcb, 0x60, 0x1d, 0x65,
	0x63, 0xdd, 0x37, 0xe7, 0x6d, 0x2f, 0x3e, 0x85, 0x05, 0x09, 0xa5, 0xe9, 0x4c, 0xef, 0x2a, 0x8d,
	0x69, 0xc8, 0xf2, 0x5c, 0xf7, 0xcd, 0xc5, 0x9e, 0x13, 0x9f, 0xc0, 0xcc, 0x3a, 0xa9, 0x6e, 0x72,
	0xd7, 0x49, 0x85, 0x69, 0xb4, 0x0c, 0xd6, 0xf3, 0x0c, 0x98, 0x7a, 0x4f, 0x8c, 0x38, 0x82, 0xe9,
	0xce, 0x58, 0xa7, 0x65, 0x83, 0xe9, 0x78, 0x19, 0xac, 0xe3, 0xec, 0x1e, 0x8b, 0x37, 0x10, 0xdd,
	0xc8, 0xed, 0x8d, 0x4c, 0x5f, 0x71, 0x91, 0xc9, 0x50, 0xe4, 0xf7, 0xc4, 0x5d, 0xea, 0xad, 0xc9,
	0xbc, 0xbc, 0xfa, 0x19, 0xe2, 0x7b, 0x4e, 0x3c, 0x87, 0xc8, 0x99, 0xb6, 0x52, 0xdc, 0x59, 0x9c,
	0x79, 0x20, 0x4e, 0x20, 0x6e, 0x65, 0xe7, 0x2a, 0x57, 0x19, 0x3d, 0xf4, 0xf1, 0x40, 0x88, 0x97,
	0x30, 0x36, 0xdb, 0xad, 0x45, 0xc7, 0x3d, 0x8c, 0xb2, 0x01, 0xad, 0xfe, 0x0a, 0x60, 0x74, 0x65,
	0x4a, 0xf1, 0x1a, 0x46, 0xca, 0xdd, 0xa6, 0xc1, 0x30, 0xab, 0x61, 0xae, 0xe7, 0x46, 0x3b, 0xbc,
	0x75, 0x19, 0x69, 0xe2, 0x10, 0xa6, 0xdc, 0x62, 0x5e, 0x15, 0x43, 0xe6, 0x09, 0xe3, 0xcb, 0x42,
	0xa4, 0x30, 0x51, 0x1d, 0x4a, 0x87, 0x45, 0x3a, 0xe5, 0xf0, 0x7b, 0x48, 0xb5, 0xd6, 0xf8, 0x01,
	0xeb, 0x14, 0x7c, 0xad, 0x0c, 0x84, 0x80, 0xd0, 0xc9, 0xd2, 0xa6, 0xe1, 0x72, 0xb4, 0x8e, 0x33,
	0x7e, 0x16, 0x2b, 0x88, 0x0a, 0x5a, 0x5e, 0xfa, 0x92, 0x6b, 0x98, 0x0f, 0xa3, 0xe0, 0x85, 0x66,
	0x5e, 0xa2, 0x3c, 0x0d, 0x5a, 0x2b, 0x4b, 0x4c, 0x53, 0x9e, 0xf3, 0x1e, 0xae, 0xfe, 0x1c, 0xc3,
	0x74, 0xbf, 0x5a, 0x4a, 0x2a, 0xeb, 0xda, 0xf8, 0x01, 0x85, 0x99, 0x07, 0xb4, 0x28, 0x67, 0x9c,
	0xac, 0x73, 0xaf, 0x8d, 0x58, 0x03, 0xa6, 0xbe, 0x66, 0x83, 0x04, 0x46, 0xf6, 0xce, 0xf2, 0x80,
	0xc2, 0x8c, 0x1e, 0x29, 0x5f, 0x6d, 0xcc, 0x4d, 0xdf, 0x5a, 0xde, 0x6b, 0x98, 0xed, 0x21, 0x57,
	0xc2, 0x71, 0x2c, 0xef, 0x34, 0xcc, 0xf6, 0x90, 0x92, 0x6f, 0x3b, 0x44, 0x9b, 0x4e, 0x7c, 0x72,
	0x06, 0xe2, 0x23, 0x80, 0x1d, 0xca, 0x76, 0xc8, 0x3d, 0x65, 0x29, 0x26, 0xc6, 0xa7, 0x3e, 0x84,
	0x29, 0xcb, 0x94, 0x3f, 0xf6, 0xf1, 0x08, 0x5f, 0xdf, 0x59, 0x71, 0x0c, 0x6c, 0x97, 0x57, 0x45,
	0x8d, 0x3c, 0xc5, 0x30, 0x63, 0xdb, 0xcb, 0xa2, 0xc6, 0xfb, 0xb0, 0x95, 0xee, 0x2d, 0xa6, 0xb3,
	0x87, 0xb0, 0x97, 0x44, 0xd0, 0x0b, 0xcc, 0x72, 0x87, 0x35, 0x4a, 0x8b, 0x45, 0x3a, 0x67, 0x8b,
	0x39, 0x91, 0xd9, 0xc0, 0x89, 0xd7, 0xc0, 0x38, 0x37, 0x9b, 0x5f, 0x50, 0x39, 0x9b, 0x2e, 0xd8,
	0x66, 0x46, 0xdc, 0x4f, 0x9e, 0x7a, 0x78, 0xc7, 0x7d, 0x9e, 0x27, 0x7e, 0x74, 0x4c, 0xf9, 0x44,
	0xc7, 0x10, 0x7b, 0x03, 0x6a, 0xe0, 0xc0, 0x17, 0xc9, 0x04, 0x75, 0xb0, 0x84, 0x79, 0x93, 0xdb,
	0x56, 0xea, 0xc1, 0x3d, 0xf1, 0xee, 0xcd, 0x75, 0x2b, 0xb5, 0x77, 0x3f, 0x01, 0x18, 0x2c, 0xc8,
	0xff, 0xa9, 0xf7, 0x67, 0x9d, 0xfc, 0x57, 0xb0, 0x68, 0x72, 0x25, 0xd5, 0x0e, 0x87, 0x00, 0xc2,
	0x57, 0xd8, 0x9c, 0x13, 0xe7, 0x23, 0x7c, 0x0c, 0xb3, 0xbd, 0x0d, 0x85, 0x78, 0xe6, 0x27, 0xe1,
	0x2d, 0x86, 0x18, 0x74, 0x4a, 0xf2, 0x9d, 0xb4, 0x3b, 0xb6, 0x78, 0xee, 0x63, 0x10, 0xf9, 0x9d,
	0xb4, 0x3b, 0xb2, 0x79, 0x01, 0xe3, 0x52, 0xb1, 0xf8, 0xc2, 0xaf, 0xae, 0x54, 0xc3, 0x02, 0x8c,
	0xdb, 0x61, 0xc7, 0xca, 0x4b, 0x5f, 0x1b, 0x13, 0x24, 0xd2, 0xed, 0xc0, 0x5b, 0x97, 0x97, 0x8a,
	0x3f, 0xe1, 0x30, 0x1b, 0x13, 0xbc, 0x50, 0x24, 0xd4, 0xd2, 0xb2, 0x90, 0x7a, 0x81, 0xe0, 0x85,
	0x12, 0x9f, 0xc1, 0x93, 0x56, 0xf6, 0x16, 0x73, 0xff, 0x32, 0x6a, 0x9b, 0x1e, 0xfa, 0xa5, 0x30,
	0xfb, 0x9e, 0xc8, 0x1f, 0xb9, 0x16, 0x3e, 0x3d, 0x2a, 0x3d, 0x5a, 0x06, 0xeb, 0x45, 0x16, 0xd1,
	0xcd, 0x51, 0xd4, 0x06, 0xd1, 0x5b, 0xd3, 0x29, 0x2c, 0x48, 0x3d, 0x66, 0x75, 0xa6, 0xfb, 0xe6,
	0x1d, 0x73, 0x17, 0x4a, 0xbc, 0x81, 0x83, 0x52, 0xd1, 0x35, 0xcb, 0xb7, 0x9d, 0x54, 0x7c, 0x0e,
	0x4e, 0x96, 0xc1, 0x3a, 0xc8, 0x16, 0xa5, 0x3a, 0x6f, 0xfb, 0x77, 0x03, 0xf9, 0xf6, 0x0f, 0x88,
	0xae, 0xf8, 0x6b, 0x8c, 0x87, 0x2f, 0x2f, 0x79, 0x24, 0xa6, 0x10, 0x56, 0x7a, 0x6b, 0x92, 0x40,
	0x00, 0x8c, 0xb5, 0x71, 0x95, 0xc2, 0xe4, 0xb1, 0x98, 0xc1, 0xe4, 0x37, 0xd9, 0xe9, 0x4a, 0x97,
	0xc9, 0x88, 0xac, 0xb1, 0xeb, 0x4c, 0x97, 0x84, 0x62, 0x0e, 0x53, 0xd5, 0x55, 0xae, 0x52, 0xb2,
	0x4e, 0x22, 0x12, 0x64, 0x8d, 0x9d, 0x4b, 0xc6, 0x62, 0x01, 0x31, 0x36, 0xd8, 0x95, 0xa8, 0xd5,
	0x5d, 0x32, 0x21, 0xa5, 0x95, 0xba, 0x52, 0xc9, 0x94, 0x1e, 0xb7, 0xd2, 0xc9, 0x3a, 0x89, 0xdf,
	0x7e, 0x05, 0xd1, 0xb7, 0x1f, 0x50, 0x3b, 0xf1, 0x0c, 0x0e, 0xae, 0x4c, 0x79, 0x65, 0xca, 0x0c,
	0x7f, 0xed, 0xd1, 0x3a, 0x2c, 0x92, 0x47, 0xe2, 0x29, 0xcc, 0xff, 0xc7, 0xfc, 0x33, 0x11, 0x4f,
	0x20, 0xbe, 0x32, 0xe5, 0xf5, 0x9d, 0x56, 0x58, 0x24, 0xff, 0x4e, 0x36, 0x63, 0x3e, 0xfb, 0x5f,
	0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xbb, 0xde, 0xb0, 0x42, 0x06, 0x00, 0x00,
}
