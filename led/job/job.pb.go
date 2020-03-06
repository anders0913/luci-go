// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/led/job/job.proto

package job

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	proto1 "go.chromium.org/luci/buildbucket/proto"
	api "go.chromium.org/luci/swarming/proto/api"
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

// Buildbucket is, ideally, just a BBAgentArgs, but there are bits of data that
// led needs to track which aren't currently contained in BBAgentArgs.
//
// Where it makes sense, this additional data should be moved from this
// Buildbucket message into BBAgentArgs, but for now we store it separately to
// get led v2 up and running.
type Buildbucket struct {
	// NOTE: for kitchen builds, the "executable_path" field is used to hold the
	// argument "$checkout_dir/luciexe". This is a bit bogus, as for kitchen we
	// only will use the "$checkout_dir" bit; but it makes it easy to upgrade
	// a task from kitchen to luciexe by just flipping the 'legacy_kitchen'
	// boolean.
	BbagentArgs      *proto1.BBAgentArgs   `protobuf:"bytes,1,opt,name=bbagent_args,json=bbagentArgs,proto3" json:"bbagent_args,omitempty"`
	CipdPackages     []*api.CIPDPackage    `protobuf:"bytes,2,rep,name=cipd_packages,json=cipdPackages,proto3" json:"cipd_packages,omitempty"`
	EnvVars          []*api.StringPair     `protobuf:"bytes,3,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty"`
	EnvPrefixes      []*api.StringListPair `protobuf:"bytes,4,rep,name=env_prefixes,json=envPrefixes,proto3" json:"env_prefixes,omitempty"`
	ExtraTags        []string              `protobuf:"bytes,5,rep,name=extra_tags,json=extraTags,proto3" json:"extra_tags,omitempty"`
	GracePeriod      *duration.Duration    `protobuf:"bytes,6,opt,name=grace_period,json=gracePeriod,proto3" json:"grace_period,omitempty"`
	BotPingTolerance *duration.Duration    `protobuf:"bytes,7,opt,name=bot_ping_tolerance,json=botPingTolerance,proto3" json:"bot_ping_tolerance,omitempty"`
	Containment      *api.Containment      `protobuf:"bytes,8,opt,name=containment,proto3" json:"containment,omitempty"`
	// Indicates that this build should be generated as a legacy kitchen task when
	// launched.
	LegacyKitchen bool `protobuf:"varint,9,opt,name=legacy_kitchen,json=legacyKitchen,proto3" json:"legacy_kitchen,omitempty"`
	// Eventually becomes the name of the launched swarming task.
	Name                 string   `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Buildbucket) Reset()         { *m = Buildbucket{} }
func (m *Buildbucket) String() string { return proto.CompactTextString(m) }
func (*Buildbucket) ProtoMessage()    {}
func (*Buildbucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_3163e69eaed6833c, []int{0}
}

func (m *Buildbucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Buildbucket.Unmarshal(m, b)
}
func (m *Buildbucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Buildbucket.Marshal(b, m, deterministic)
}
func (m *Buildbucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buildbucket.Merge(m, src)
}
func (m *Buildbucket) XXX_Size() int {
	return xxx_messageInfo_Buildbucket.Size(m)
}
func (m *Buildbucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Buildbucket.DiscardUnknown(m)
}

var xxx_messageInfo_Buildbucket proto.InternalMessageInfo

func (m *Buildbucket) GetBbagentArgs() *proto1.BBAgentArgs {
	if m != nil {
		return m.BbagentArgs
	}
	return nil
}

func (m *Buildbucket) GetCipdPackages() []*api.CIPDPackage {
	if m != nil {
		return m.CipdPackages
	}
	return nil
}

func (m *Buildbucket) GetEnvVars() []*api.StringPair {
	if m != nil {
		return m.EnvVars
	}
	return nil
}

func (m *Buildbucket) GetEnvPrefixes() []*api.StringListPair {
	if m != nil {
		return m.EnvPrefixes
	}
	return nil
}

func (m *Buildbucket) GetExtraTags() []string {
	if m != nil {
		return m.ExtraTags
	}
	return nil
}

func (m *Buildbucket) GetGracePeriod() *duration.Duration {
	if m != nil {
		return m.GracePeriod
	}
	return nil
}

func (m *Buildbucket) GetBotPingTolerance() *duration.Duration {
	if m != nil {
		return m.BotPingTolerance
	}
	return nil
}

func (m *Buildbucket) GetContainment() *api.Containment {
	if m != nil {
		return m.Containment
	}
	return nil
}

func (m *Buildbucket) GetLegacyKitchen() bool {
	if m != nil {
		return m.LegacyKitchen
	}
	return false
}

func (m *Buildbucket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Swarming is the raw TaskRequest. When a Definition is in this form, the
// user's ability to manipulate it via `led` subcommands is extremely limited.
type Swarming struct {
	Task                 *api.TaskRequest `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Hostname             string           `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Swarming) Reset()         { *m = Swarming{} }
func (m *Swarming) String() string { return proto.CompactTextString(m) }
func (*Swarming) ProtoMessage()    {}
func (*Swarming) Descriptor() ([]byte, []int) {
	return fileDescriptor_3163e69eaed6833c, []int{1}
}

func (m *Swarming) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Swarming.Unmarshal(m, b)
}
func (m *Swarming) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Swarming.Marshal(b, m, deterministic)
}
func (m *Swarming) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Swarming.Merge(m, src)
}
func (m *Swarming) XXX_Size() int {
	return xxx_messageInfo_Swarming.Size(m)
}
func (m *Swarming) XXX_DiscardUnknown() {
	xxx_messageInfo_Swarming.DiscardUnknown(m)
}

var xxx_messageInfo_Swarming proto.InternalMessageInfo

func (m *Swarming) GetTask() *api.TaskRequest {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *Swarming) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type Definition struct {
	// Types that are valid to be assigned to JobType:
	//	*Definition_Buildbucket
	//	*Definition_Swarming
	JobType isDefinition_JobType `protobuf_oneof:"job_type"`
	// If set, this holds the CASTree to use with the build, when launched.
	//
	// At the time of launch, this will be merged with
	// swarming.task_slice[*].properties.cas_inputs, if any.
	//
	// The 'server' and 'namespace' fields here are used as the defaults for any
	// digests specified without server/namespace.
	UserPayload          *api.CASTree `protobuf:"bytes,3,opt,name=user_payload,json=userPayload,proto3" json:"user_payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Definition) Reset()         { *m = Definition{} }
func (m *Definition) String() string { return proto.CompactTextString(m) }
func (*Definition) ProtoMessage()    {}
func (*Definition) Descriptor() ([]byte, []int) {
	return fileDescriptor_3163e69eaed6833c, []int{2}
}

func (m *Definition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Definition.Unmarshal(m, b)
}
func (m *Definition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Definition.Marshal(b, m, deterministic)
}
func (m *Definition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Definition.Merge(m, src)
}
func (m *Definition) XXX_Size() int {
	return xxx_messageInfo_Definition.Size(m)
}
func (m *Definition) XXX_DiscardUnknown() {
	xxx_messageInfo_Definition.DiscardUnknown(m)
}

var xxx_messageInfo_Definition proto.InternalMessageInfo

type isDefinition_JobType interface {
	isDefinition_JobType()
}

type Definition_Buildbucket struct {
	Buildbucket *Buildbucket `protobuf:"bytes,1,opt,name=buildbucket,proto3,oneof"`
}

type Definition_Swarming struct {
	Swarming *Swarming `protobuf:"bytes,2,opt,name=swarming,proto3,oneof"`
}

func (*Definition_Buildbucket) isDefinition_JobType() {}

func (*Definition_Swarming) isDefinition_JobType() {}

func (m *Definition) GetJobType() isDefinition_JobType {
	if m != nil {
		return m.JobType
	}
	return nil
}

func (m *Definition) GetBuildbucket() *Buildbucket {
	if x, ok := m.GetJobType().(*Definition_Buildbucket); ok {
		return x.Buildbucket
	}
	return nil
}

func (m *Definition) GetSwarming() *Swarming {
	if x, ok := m.GetJobType().(*Definition_Swarming); ok {
		return x.Swarming
	}
	return nil
}

func (m *Definition) GetUserPayload() *api.CASTree {
	if m != nil {
		return m.UserPayload
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Definition) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Definition_Buildbucket)(nil),
		(*Definition_Swarming)(nil),
	}
}

func init() {
	proto.RegisterType((*Buildbucket)(nil), "job.Buildbucket")
	proto.RegisterType((*Swarming)(nil), "job.Swarming")
	proto.RegisterType((*Definition)(nil), "job.Definition")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/led/job/job.proto", fileDescriptor_3163e69eaed6833c)
}

var fileDescriptor_3163e69eaed6833c = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x5d, 0x6b, 0xdb, 0x3c,
	0x14, 0x80, 0x9b, 0x26, 0x6f, 0xeb, 0xc8, 0xe9, 0x4b, 0x11, 0x83, 0x79, 0x2d, 0x1b, 0x21, 0xb0,
	0x11, 0xd8, 0x50, 0x58, 0xf6, 0x05, 0xfb, 0x82, 0x66, 0x85, 0x75, 0x6c, 0x17, 0x46, 0x0d, 0xbb,
	0xd8, 0x8d, 0x91, 0x9c, 0x53, 0x45, 0x89, 0x23, 0x79, 0x92, 0xec, 0xb5, 0x3f, 0x66, 0xff, 0x62,
	0x3f, 0x70, 0x58, 0x76, 0xda, 0x74, 0x84, 0xed, 0xc2, 0x60, 0x1d, 0x3f, 0xcf, 0x39, 0xc7, 0x3a,
	0x12, 0x7a, 0x24, 0x34, 0x49, 0xe7, 0x46, 0xaf, 0x64, 0xb1, 0x22, 0xda, 0x88, 0x51, 0x56, 0xa4,
	0x72, 0x94, 0xc1, 0x6c, 0xb4, 0xd0, 0xbc, 0x7a, 0x48, 0x6e, 0xb4, 0xd3, 0xb8, 0xbd, 0xd0, 0xfc,
	0xe8, 0x81, 0xd0, 0x5a, 0x64, 0x30, 0xf2, 0x21, 0x5e, 0x5c, 0x8c, 0x66, 0x85, 0x61, 0x4e, 0x6a,
	0x55, 0x43, 0x47, 0x2f, 0xb6, 0x26, 0xe3, 0x85, 0xcc, 0x66, 0xbc, 0x48, 0x97, 0xe0, 0x6a, 0x73,
	0x94, 0xb1, 0x42, 0xa5, 0x73, 0x30, 0x8d, 0xf6, 0x72, 0xab, 0x66, 0x7f, 0x30, 0xb3, 0x92, 0x4a,
	0x34, 0x0e, 0xcb, 0x6f, 0x42, 0xb5, 0x37, 0xf8, 0xd9, 0x41, 0xe1, 0xe4, 0x26, 0x39, 0x7e, 0x8f,
	0x7a, 0x9c, 0x33, 0x01, 0xca, 0x25, 0xcc, 0x08, 0x1b, 0xb5, 0xfa, 0xad, 0x61, 0x38, 0x3e, 0x26,
	0x1b, 0x0d, 0x90, 0x72, 0x4c, 0x26, 0x93, 0x93, 0x8a, 0x39, 0x31, 0xc2, 0xd2, 0xb0, 0x11, 0xaa,
	0x05, 0x7e, 0x87, 0x0e, 0x52, 0x99, 0xcf, 0x92, 0x9c, 0xa5, 0x4b, 0x26, 0xc0, 0x46, 0xbb, 0xfd,
	0xf6, 0x30, 0x1c, 0x47, 0xe4, 0xba, 0x6e, 0xf9, 0x94, 0x7c, 0xf8, 0x14, 0x9f, 0xc6, 0x35, 0x40,
	0x7b, 0x15, 0xde, 0x2c, 0x2c, 0x1e, 0xa3, 0x00, 0x54, 0x99, 0x94, 0xcc, 0xd8, 0xa8, 0xed, 0xcd,
	0xbb, 0xb7, 0xcc, 0x73, 0x67, 0xa4, 0x12, 0x31, 0x93, 0x86, 0xee, 0x83, 0x2a, 0xbf, 0x32, 0x63,
	0xab, 0x96, 0x2b, 0x27, 0x37, 0x70, 0x21, 0x2f, 0xc1, 0x46, 0x1d, 0xef, 0x1d, 0x6f, 0xf1, 0xbe,
	0x48, 0xeb, 0xbc, 0x1b, 0x82, 0x2a, 0xe3, 0x86, 0xc7, 0xf7, 0x11, 0x82, 0x4b, 0x67, 0x58, 0xe2,
	0x98, 0xb0, 0xd1, 0x7f, 0xfd, 0xf6, 0xb0, 0x4b, 0xbb, 0x3e, 0x32, 0x65, 0xc2, 0xe2, 0xb7, 0xa8,
	0x27, 0x0c, 0x4b, 0x21, 0xc9, 0xc1, 0x48, 0x3d, 0x8b, 0xf6, 0xfc, 0x8e, 0xdc, 0x23, 0xf5, 0x1c,
	0xc9, 0x7a, 0x8e, 0xe4, 0xb4, 0x99, 0x23, 0x0d, 0x3d, 0x1e, 0x7b, 0x1a, 0x7f, 0x44, 0x98, 0x6b,
	0x97, 0xe4, 0x52, 0x89, 0xc4, 0xe9, 0x0c, 0x0c, 0x53, 0x29, 0x44, 0xfb, 0xff, 0xca, 0x71, 0xc8,
	0xb5, 0x8b, 0xa5, 0x12, 0xd3, 0xb5, 0x82, 0x5f, 0xa3, 0x30, 0xd5, 0xca, 0x31, 0xa9, 0x56, 0xa0,
	0x5c, 0x14, 0xf8, 0x0c, 0x7f, 0x6c, 0xeb, 0xcd, 0x77, 0xba, 0x09, 0xe3, 0x87, 0xe8, 0xff, 0x0c,
	0x04, 0x4b, 0xaf, 0x92, 0xa5, 0x74, 0xe9, 0x1c, 0x54, 0xd4, 0xed, 0xb7, 0x86, 0x01, 0x3d, 0xa8,
	0xa3, 0x9f, 0xeb, 0x20, 0xc6, 0xa8, 0xa3, 0xd8, 0x0a, 0x22, 0xd4, 0x6f, 0x0d, 0xbb, 0xd4, 0xbf,
	0x0f, 0xa6, 0x28, 0x38, 0x6f, 0x4a, 0xe0, 0x27, 0xa8, 0xe3, 0x98, 0x5d, 0x36, 0x67, 0xe2, 0x76,
	0xed, 0x29, 0xb3, 0x4b, 0x0a, 0xdf, 0x0b, 0xb0, 0x8e, 0x7a, 0x0a, 0x1f, 0xa1, 0x60, 0xae, 0xad,
	0xf3, 0x19, 0x77, 0x7d, 0xc6, 0xeb, 0xf5, 0xe0, 0x57, 0x0b, 0xa1, 0x53, 0xb8, 0x90, 0x4a, 0x56,
	0x7f, 0x8b, 0x9f, 0xa3, 0x70, 0xe3, 0x7c, 0x35, 0xf9, 0x0f, 0x49, 0x75, 0x73, 0x36, 0xce, 0xe6,
	0xd9, 0x0e, 0xdd, 0xc4, 0xf0, 0x63, 0x14, 0xac, 0x3b, 0xf0, 0x05, 0xc2, 0xf1, 0x81, 0x57, 0xd6,
	0xfd, 0x9e, 0xed, 0xd0, 0x6b, 0x00, 0xbf, 0x42, 0xbd, 0xc2, 0x82, 0x49, 0x72, 0x76, 0x95, 0x69,
	0x36, 0x8b, 0xda, 0x5e, 0xb8, 0x73, 0x7b, 0xff, 0x4e, 0xce, 0xa7, 0x06, 0x80, 0x86, 0x15, 0x19,
	0xd7, 0xe0, 0x04, 0xa1, 0x60, 0xa1, 0x79, 0xe2, 0xae, 0x72, 0x98, 0x0c, 0xbe, 0xf5, 0xff, 0x76,
	0xd5, 0xdf, 0x2c, 0x34, 0xe7, 0x7b, 0x7e, 0x98, 0xcf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc1,
	0x82, 0xb5, 0xd7, 0x15, 0x04, 0x00, 0x00,
}
