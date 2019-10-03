// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/v1/invocation.proto

package resultspb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Invocation_State int32

const (
	// The default value. This value is used if the state is omitted.
	Invocation_STATE_UNSPECIFIED Invocation_State = 0
	// The invocation was created and accepts new results.
	Invocation_ACTIVE Invocation_State = 1
	// The invocation is finalized and contains all the results that the
	// associated computation was expected to compute; unlike INTERRUPTED state.
	//
	// The invocation is immutable and no longer accepts new results.
	Invocation_COMPLETED Invocation_State = 2
	// The invocation is finalized and does NOT contain all the results that the
	// associated computation was expected to compute.
	// The computation was interrupted prematurely.
	//
	// Such invocation should be discarded.
	// Often the associated computation is retried.
	//
	// The invocation is immutable and no longer accepts new results.
	Invocation_INTERRUPTED Invocation_State = 3
)

var Invocation_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "ACTIVE",
	2: "COMPLETED",
	3: "INTERRUPTED",
}

var Invocation_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"ACTIVE":            1,
	"COMPLETED":         2,
	"INTERRUPTED":       3,
}

func (x Invocation_State) String() string {
	return proto.EnumName(Invocation_State_name, int32(x))
}

func (Invocation_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f10bfdd46378b04d, []int{0, 0}
}

// A conceptual container of results. Immutable once finalized.
// It represents all results of some computation; examples: swarming task,
// buildbucket build, CQ attempt.
// Composable: can include other invocations, see inclusion.proto.
type Invocation struct {
	// The resource name of this invocation.
	// Format: invocations/{INVOCATION_ID}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Current state of the invocation.
	State Invocation_State `protobuf:"varint,2,opt,name=state,proto3,enum=luci.resultdb.Invocation_State" json:"state,omitempty"`
	// When the invocation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Invocation-level string key-value pairs.
	// A key can be repeated.
	Tags []*StringPair `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	// When the invocation was finalized, i.e. transitioned to COMPLETED or
	// INTERRUPTED state.
	// If this field is set, implies that the invocation is finalized.
	FinalizeTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=finalize_time,json=finalizeTime,proto3" json:"finalize_time,omitempty"`
	// Timestamp when the invocation will be forcefully finalized.
	// Can be extended with UpdateInvocation until finalized.
	Deadline *timestamp.Timestamp `protobuf:"bytes,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Base variant definition for test results in this invocation.
	// A particular test result can have additional key-value pairs.
	BaseTestVariantDef   *VariantDef `protobuf:"bytes,7,opt,name=base_test_variant_def,json=baseTestVariantDef,proto3" json:"base_test_variant_def,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Invocation) Reset()         { *m = Invocation{} }
func (m *Invocation) String() string { return proto.CompactTextString(m) }
func (*Invocation) ProtoMessage()    {}
func (*Invocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f10bfdd46378b04d, []int{0}
}

func (m *Invocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation.Unmarshal(m, b)
}
func (m *Invocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation.Marshal(b, m, deterministic)
}
func (m *Invocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation.Merge(m, src)
}
func (m *Invocation) XXX_Size() int {
	return xxx_messageInfo_Invocation.Size(m)
}
func (m *Invocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation proto.InternalMessageInfo

func (m *Invocation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Invocation) GetState() Invocation_State {
	if m != nil {
		return m.State
	}
	return Invocation_STATE_UNSPECIFIED
}

func (m *Invocation) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Invocation) GetTags() []*StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Invocation) GetFinalizeTime() *timestamp.Timestamp {
	if m != nil {
		return m.FinalizeTime
	}
	return nil
}

func (m *Invocation) GetDeadline() *timestamp.Timestamp {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *Invocation) GetBaseTestVariantDef() *VariantDef {
	if m != nil {
		return m.BaseTestVariantDef
	}
	return nil
}

// One inclusion edge in the invocation DAG.
//
// Invocations are composable: one invocation can include zero or more other
// invocations, representing a cumulative result. For example, a Buildbucket
// build invocation can include invocations of all child swarming tasks and
// represent overall result of the build, encapsulating the internal structure
// of the build from the client that just needs to load test results scoped
// to the build.
//
// The graph is directed and acyclic. There can be at most one edge between a
// given pair of invocations.
// Including invocation MUST NOT be finalized.
// Included invocation MAY be finalized.
type Inclusion struct {
	// Resource name, identifier of the inclusion.
	// Format:
	// invocations/{INCLUDING_INVOCATION_ID}/inclusions/{INCLUDED_INVOCATION_ID}
	// This implies that there can be only one direct edge between a given pair of
	// invocations
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Name of the included invocation.
	// FORMAT: invocations/{INCLUDED_INVOCATION_ID}.
	IncludedInvocation string `protobuf:"bytes,2,opt,name=included_invocation,json=includedInvocation,proto3" json:"included_invocation,omitempty"`
	// Name of the another inclusion that overrides this one. OUTPUT_ONLY.
	// If set, the invocation by this inclusion no longer influences the final
	// outcome of the including invocation. A typical example is a retry: the
	// new attempt overrides the previous one.
	//
	// Use recorder.OverrideInclusion to set this field.
	OverriddenBy string `protobuf:"bytes,3,opt,name=overridden_by,json=overriddenBy,proto3" json:"overridden_by,omitempty"`
	// Whether the included invocation is finalized before the including
	// invocation.
	// The formula for the field is
	//   included_inv.finalize_time < including_inv.finalize_time
	// If the included invocation is finalized, but the including invocation is
	// not yet, the edge is ready. If both are not finalized yet, the edge is not
	// ready *yet*, but its value may change over time, until the including
	// invocation is finalized.
	//
	// In practice, either
	// - an edge is ready because the including is expected to wait for its
	//   children to conclude its own result, OR
	// - it does not matter e.g. if the including was canceled and finalized
	//   prematurely.
	//
	// By default, QueryTestResults ignores un-ready inclusions.
	Ready                bool     `protobuf:"varint,4,opt,name=ready,proto3" json:"ready,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Inclusion) Reset()         { *m = Inclusion{} }
func (m *Inclusion) String() string { return proto.CompactTextString(m) }
func (*Inclusion) ProtoMessage()    {}
func (*Inclusion) Descriptor() ([]byte, []int) {
	return fileDescriptor_f10bfdd46378b04d, []int{1}
}

func (m *Inclusion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inclusion.Unmarshal(m, b)
}
func (m *Inclusion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inclusion.Marshal(b, m, deterministic)
}
func (m *Inclusion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inclusion.Merge(m, src)
}
func (m *Inclusion) XXX_Size() int {
	return xxx_messageInfo_Inclusion.Size(m)
}
func (m *Inclusion) XXX_DiscardUnknown() {
	xxx_messageInfo_Inclusion.DiscardUnknown(m)
}

var xxx_messageInfo_Inclusion proto.InternalMessageInfo

func (m *Inclusion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Inclusion) GetIncludedInvocation() string {
	if m != nil {
		return m.IncludedInvocation
	}
	return ""
}

func (m *Inclusion) GetOverriddenBy() string {
	if m != nil {
		return m.OverriddenBy
	}
	return ""
}

func (m *Inclusion) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

func init() {
	proto.RegisterEnum("luci.resultdb.Invocation_State", Invocation_State_name, Invocation_State_value)
	proto.RegisterType((*Invocation)(nil), "luci.resultdb.Invocation")
	proto.RegisterType((*Inclusion)(nil), "luci.resultdb.Inclusion")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/v1/invocation.proto", fileDescriptor_f10bfdd46378b04d)
}

var fileDescriptor_f10bfdd46378b04d = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x9d, 0x84, 0x66, 0x42, 0x20, 0x2c, 0xaa, 0xe4, 0xe6, 0xd2, 0x28, 0x27, 0x5f,
	0xb0, 0x21, 0x95, 0x40, 0x82, 0x53, 0xfe, 0x18, 0xc9, 0xa8, 0x94, 0xc8, 0x71, 0x7b, 0xe0, 0x62,
	0xad, 0xed, 0x89, 0xbb, 0x92, 0xed, 0x8d, 0xd6, 0x9b, 0x48, 0xe1, 0x5d, 0x78, 0xb7, 0x3c, 0x02,
	0x8f, 0x80, 0xbc, 0x8e, 0x13, 0xda, 0x4b, 0x7a, 0xdc, 0x6f, 0x7e, 0x33, 0xf3, 0xcd, 0xce, 0xc0,
	0xe7, 0x84, 0x5b, 0xd1, 0x83, 0xe0, 0x19, 0x5b, 0x67, 0x16, 0x17, 0x89, 0x9d, 0xae, 0x23, 0x66,
	0x0b, 0x2c, 0xd6, 0xa9, 0x8c, 0x43, 0x7b, 0x25, 0xb8, 0xe4, 0xf6, 0xe6, 0xa3, 0xcd, 0xf2, 0x0d,
	0x8f, 0xa8, 0x64, 0x3c, 0xb7, 0x94, 0x46, 0xba, 0x25, 0x68, 0xd5, 0x60, 0xff, 0x2a, 0xe1, 0x3c,
	0x49, 0xd1, 0xa6, 0x2b, 0x66, 0x2f, 0x19, 0xa6, 0x71, 0x10, 0xe2, 0x03, 0xdd, 0x30, 0x2e, 0x2a,
	0xfe, 0x00, 0xa8, 0x57, 0xb8, 0x5e, 0xda, 0x92, 0x65, 0x58, 0x48, 0x9a, 0xad, 0xf6, 0xc0, 0xf5,
	0x33, 0x9d, 0x44, 0x3c, 0xcb, 0x6a, 0x17, 0xc3, 0xbf, 0x3a, 0x80, 0x7b, 0xb0, 0x46, 0xfa, 0xd0,
	0xc8, 0x69, 0x86, 0x86, 0x36, 0xd0, 0xcc, 0xf6, 0xa4, 0xb5, 0x1b, 0xeb, 0xbb, 0x71, 0xd3, 0x53,
	0x1a, 0xf9, 0x02, 0xcd, 0x42, 0x52, 0x89, 0xc6, 0xd9, 0x40, 0x33, 0x5f, 0x8f, 0xae, 0xac, 0x47,
	0x03, 0x58, 0xc7, 0x2a, 0xd6, 0xa2, 0xc4, 0x26, 0xfa, 0x6e, 0xac, 0x7b, 0x55, 0x0a, 0x99, 0x42,
	0x27, 0x12, 0x48, 0x25, 0x06, 0xa5, 0x6b, 0x43, 0x1f, 0x68, 0x66, 0x67, 0xd4, 0xb7, 0xaa, 0x91,
	0xac, 0x7a, 0x24, 0xcb, 0xaf, 0x47, 0x3a, 0xb4, 0x86, 0x2a, 0xad, 0x0c, 0x90, 0xf7, 0xd0, 0x90,
	0x34, 0x29, 0x8c, 0xc6, 0x40, 0x37, 0x3b, 0xa3, 0xcb, 0x27, 0xfd, 0x17, 0x52, 0xb0, 0x3c, 0x99,
	0x53, 0x26, 0x3c, 0x85, 0x91, 0x19, 0x74, 0x97, 0x2c, 0xa7, 0x29, 0xfb, 0xbd, 0xef, 0xda, 0x3c,
	0xd9, 0x55, 0x59, 0x7e, 0x55, 0x67, 0xa9, 0xa6, 0x9f, 0xe0, 0x3c, 0x46, 0x1a, 0xa7, 0x2c, 0x47,
	0xa3, 0x75, 0xaa, 0x80, 0x77, 0x60, 0xc9, 0x0d, 0x5c, 0x84, 0xb4, 0xc0, 0x40, 0x62, 0x21, 0x83,
	0x0d, 0x15, 0x8c, 0xe6, 0x32, 0x88, 0x71, 0x69, 0xbc, 0x54, 0x45, 0x9e, 0xba, 0xbf, 0xaf, 0x88,
	0x19, 0x2e, 0x3d, 0x52, 0xe6, 0xf9, 0x58, 0xc8, 0xa3, 0x36, 0xfc, 0x0e, 0x4d, 0xf5, 0xa9, 0xe4,
	0x02, 0xde, 0x2e, 0xfc, 0xb1, 0xef, 0x04, 0x77, 0xb7, 0x8b, 0xb9, 0x33, 0x75, 0xbf, 0xb9, 0xce,
	0xac, 0xf7, 0x82, 0x00, 0xb4, 0xc6, 0x53, 0xdf, 0xbd, 0x77, 0x7a, 0x1a, 0xe9, 0x42, 0x7b, 0xfa,
	0xf3, 0xc7, 0xfc, 0xc6, 0xf1, 0x9d, 0x59, 0xef, 0x8c, 0xbc, 0x81, 0x8e, 0x7b, 0xeb, 0x3b, 0x9e,
	0x77, 0x37, 0x2f, 0x05, 0x7d, 0xf8, 0x47, 0x83, 0xb6, 0x9b, 0x47, 0xe9, 0xba, 0x38, 0xb5, 0x71,
	0x1b, 0xde, 0xb1, 0x12, 0x8c, 0x31, 0x0e, 0x8e, 0xf7, 0xab, 0xf6, 0xdf, 0xf6, 0x48, 0x1d, 0xfa,
	0xef, 0x7c, 0x4c, 0xe8, 0xf2, 0x0d, 0x0a, 0xc1, 0xe2, 0x18, 0xf3, 0x20, 0xdc, 0xaa, 0x45, 0xb7,
	0xf7, 0xdf, 0x7a, 0x8c, 0x4c, 0xb6, 0xe4, 0x12, 0x9a, 0x02, 0x69, 0xbc, 0x35, 0x1a, 0x03, 0xcd,
	0x3c, 0xdf, 0xdf, 0x8a, 0x52, 0x26, 0xa3, 0x5f, 0x1f, 0x9e, 0x77, 0xc9, 0x5f, 0x2b, 0xa5, 0x58,
	0x85, 0x61, 0x4b, 0x69, 0xd7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xae, 0xdc, 0x21, 0x03, 0x8e,
	0x03, 0x00, 0x00,
}