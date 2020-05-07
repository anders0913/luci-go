// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/milo/api/config/settings.proto

package config

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Settings represents the format for the global (service) config for Milo.
type Settings struct {
	Buildbot    *Settings_Buildbot    `protobuf:"bytes,1,opt,name=buildbot,proto3" json:"buildbot,omitempty"`
	Buildbucket *Settings_Buildbucket `protobuf:"bytes,2,opt,name=buildbucket,proto3" json:"buildbucket,omitempty"`
	Swarming    *Settings_Swarming    `protobuf:"bytes,3,opt,name=swarming,proto3" json:"swarming,omitempty"`
	// source_acls instructs Milo to provide Git/Gerrit data
	// (e.g., blamelist) to some of its users on entire subdomains or individual
	// repositories (Gerrit "projects").
	//
	// Multiple records are allowed, but each host and project must appear only in
	// one record.
	SourceAcls           []*Settings_SourceAcls `protobuf:"bytes,4,rep,name=source_acls,json=sourceAcls,proto3" json:"source_acls,omitempty"`
	Resultdb             *Settings_ResultDB     `protobuf:"bytes,5,opt,name=resultdb,proto3" json:"resultdb,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dd5cb9562385c0, []int{0}
}

func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetBuildbot() *Settings_Buildbot {
	if m != nil {
		return m.Buildbot
	}
	return nil
}

func (m *Settings) GetBuildbucket() *Settings_Buildbucket {
	if m != nil {
		return m.Buildbucket
	}
	return nil
}

func (m *Settings) GetSwarming() *Settings_Swarming {
	if m != nil {
		return m.Swarming
	}
	return nil
}

func (m *Settings) GetSourceAcls() []*Settings_SourceAcls {
	if m != nil {
		return m.SourceAcls
	}
	return nil
}

func (m *Settings) GetResultdb() *Settings_ResultDB {
	if m != nil {
		return m.Resultdb
	}
	return nil
}

type Settings_Buildbot struct {
	// internal_reader is the infra-auth group that is allowed to read internal
	// buildbot data.
	InternalReader string `protobuf:"bytes,1,opt,name=internal_reader,json=internalReader,proto3" json:"internal_reader,omitempty"`
	// public_subscription is the name of the pubsub topic where public builds come in
	// from
	PublicSubscription string `protobuf:"bytes,2,opt,name=public_subscription,json=publicSubscription,proto3" json:"public_subscription,omitempty"`
	// internal_subscription is the name of the pubsub topic where internal builds
	// come in from
	InternalSubscription string   `protobuf:"bytes,3,opt,name=internal_subscription,json=internalSubscription,proto3" json:"internal_subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_Buildbot) Reset()         { *m = Settings_Buildbot{} }
func (m *Settings_Buildbot) String() string { return proto.CompactTextString(m) }
func (*Settings_Buildbot) ProtoMessage()    {}
func (*Settings_Buildbot) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dd5cb9562385c0, []int{0, 0}
}

func (m *Settings_Buildbot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_Buildbot.Unmarshal(m, b)
}
func (m *Settings_Buildbot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_Buildbot.Marshal(b, m, deterministic)
}
func (m *Settings_Buildbot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_Buildbot.Merge(m, src)
}
func (m *Settings_Buildbot) XXX_Size() int {
	return xxx_messageInfo_Settings_Buildbot.Size(m)
}
func (m *Settings_Buildbot) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_Buildbot.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_Buildbot proto.InternalMessageInfo

func (m *Settings_Buildbot) GetInternalReader() string {
	if m != nil {
		return m.InternalReader
	}
	return ""
}

func (m *Settings_Buildbot) GetPublicSubscription() string {
	if m != nil {
		return m.PublicSubscription
	}
	return ""
}

func (m *Settings_Buildbot) GetInternalSubscription() string {
	if m != nil {
		return m.InternalSubscription
	}
	return ""
}

type Settings_Buildbucket struct {
	// name is the user friendly name of the Buildbucket instance we're pointing to.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// host is the hostname of the buildbucket instance we're pointing to (sans scheme).
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// project is the name of the Google Cloud project that the pubsub topic
	// belongs to.
	Project              string   `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_Buildbucket) Reset()         { *m = Settings_Buildbucket{} }
func (m *Settings_Buildbucket) String() string { return proto.CompactTextString(m) }
func (*Settings_Buildbucket) ProtoMessage()    {}
func (*Settings_Buildbucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dd5cb9562385c0, []int{0, 1}
}

func (m *Settings_Buildbucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_Buildbucket.Unmarshal(m, b)
}
func (m *Settings_Buildbucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_Buildbucket.Marshal(b, m, deterministic)
}
func (m *Settings_Buildbucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_Buildbucket.Merge(m, src)
}
func (m *Settings_Buildbucket) XXX_Size() int {
	return xxx_messageInfo_Settings_Buildbucket.Size(m)
}
func (m *Settings_Buildbucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_Buildbucket.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_Buildbucket proto.InternalMessageInfo

func (m *Settings_Buildbucket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Settings_Buildbucket) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Settings_Buildbucket) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

type Settings_Swarming struct {
	// default_host is the hostname of the swarming host Milo defaults to, if
	// none is specified.  Default host is implicitly an allowed host.
	DefaultHost string `protobuf:"bytes,1,opt,name=default_host,json=defaultHost,proto3" json:"default_host,omitempty"`
	// allowed_hosts is a whitelist of hostnames of swarming instances
	// that Milo is allowed to talk to.  This is specified here for security
	// reasons, because Milo will hand out its oauth2 token to a swarming host.
	AllowedHosts         []string `protobuf:"bytes,2,rep,name=allowed_hosts,json=allowedHosts,proto3" json:"allowed_hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_Swarming) Reset()         { *m = Settings_Swarming{} }
func (m *Settings_Swarming) String() string { return proto.CompactTextString(m) }
func (*Settings_Swarming) ProtoMessage()    {}
func (*Settings_Swarming) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dd5cb9562385c0, []int{0, 2}
}

func (m *Settings_Swarming) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_Swarming.Unmarshal(m, b)
}
func (m *Settings_Swarming) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_Swarming.Marshal(b, m, deterministic)
}
func (m *Settings_Swarming) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_Swarming.Merge(m, src)
}
func (m *Settings_Swarming) XXX_Size() int {
	return xxx_messageInfo_Settings_Swarming.Size(m)
}
func (m *Settings_Swarming) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_Swarming.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_Swarming proto.InternalMessageInfo

func (m *Settings_Swarming) GetDefaultHost() string {
	if m != nil {
		return m.DefaultHost
	}
	return ""
}

func (m *Settings_Swarming) GetAllowedHosts() []string {
	if m != nil {
		return m.AllowedHosts
	}
	return nil
}

// SourceAcls grants read access on a set of Git/Gerrit hosts or projects.
type Settings_SourceAcls struct {
	// host grants read access on all project at this host.
	//
	// For more granularity, use the project field instead.
	//
	// For *.googlesource.com domains, host should not be a Gerrit host,
	// i.e.  it shouldn't be <subdomain>-review.googlesource.com.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// project is a URL to a Git repository.
	//
	// Read access is granted on both git data and Gerrit CLs of this project.
	//
	// For *.googlesource.com Git repositories:
	//   URL Path should not start with '/a/' (forced authentication).
	//   URL Path should not end with '.git' (redundant).
	Projects []string `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	// readers are allowed to read git/gerrit data from targets.
	//
	// Three types of identity strings are supported:
	//  * Emails.                   For example: "someuser@example.com"
	//  * Chrome-infra-auth Groups. For example: "group:committers"
	//  * Auth service identities.  For example: "kind:name"
	//
	// Required.
	Readers              []string `protobuf:"bytes,3,rep,name=readers,proto3" json:"readers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_SourceAcls) Reset()         { *m = Settings_SourceAcls{} }
func (m *Settings_SourceAcls) String() string { return proto.CompactTextString(m) }
func (*Settings_SourceAcls) ProtoMessage()    {}
func (*Settings_SourceAcls) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dd5cb9562385c0, []int{0, 3}
}

func (m *Settings_SourceAcls) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_SourceAcls.Unmarshal(m, b)
}
func (m *Settings_SourceAcls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_SourceAcls.Marshal(b, m, deterministic)
}
func (m *Settings_SourceAcls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_SourceAcls.Merge(m, src)
}
func (m *Settings_SourceAcls) XXX_Size() int {
	return xxx_messageInfo_Settings_SourceAcls.Size(m)
}
func (m *Settings_SourceAcls) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_SourceAcls.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_SourceAcls proto.InternalMessageInfo

func (m *Settings_SourceAcls) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Settings_SourceAcls) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

func (m *Settings_SourceAcls) GetReaders() []string {
	if m != nil {
		return m.Readers
	}
	return nil
}

type Settings_ResultDB struct {
	// host is the hostname of the ResultDB instance we're pointing to (sans scheme).
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_ResultDB) Reset()         { *m = Settings_ResultDB{} }
func (m *Settings_ResultDB) String() string { return proto.CompactTextString(m) }
func (*Settings_ResultDB) ProtoMessage()    {}
func (*Settings_ResultDB) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dd5cb9562385c0, []int{0, 4}
}

func (m *Settings_ResultDB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_ResultDB.Unmarshal(m, b)
}
func (m *Settings_ResultDB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_ResultDB.Marshal(b, m, deterministic)
}
func (m *Settings_ResultDB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_ResultDB.Merge(m, src)
}
func (m *Settings_ResultDB) XXX_Size() int {
	return xxx_messageInfo_Settings_ResultDB.Size(m)
}
func (m *Settings_ResultDB) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_ResultDB.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_ResultDB proto.InternalMessageInfo

func (m *Settings_ResultDB) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func init() {
	proto.RegisterType((*Settings)(nil), "milo.Settings")
	proto.RegisterType((*Settings_Buildbot)(nil), "milo.Settings.Buildbot")
	proto.RegisterType((*Settings_Buildbucket)(nil), "milo.Settings.Buildbucket")
	proto.RegisterType((*Settings_Swarming)(nil), "milo.Settings.Swarming")
	proto.RegisterType((*Settings_SourceAcls)(nil), "milo.Settings.SourceAcls")
	proto.RegisterType((*Settings_ResultDB)(nil), "milo.Settings.ResultDB")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/milo/api/config/settings.proto", fileDescriptor_98dd5cb9562385c0)
}

var fileDescriptor_98dd5cb9562385c0 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x95, 0x4d, 0xbb, 0xa4, 0x93, 0x05, 0x24, 0xb3, 0x88, 0x90, 0x03, 0x2a, 0x70, 0xa0,
	0xa7, 0x44, 0xda, 0xde, 0x10, 0x17, 0x2a, 0x0e, 0xdc, 0x90, 0xdc, 0x0b, 0xe2, 0x52, 0x39, 0x8e,
	0x37, 0x6b, 0x70, 0xe2, 0xc8, 0x7f, 0xb4, 0xcf, 0xc2, 0xcb, 0xf1, 0x2c, 0xc8, 0x8e, 0x9d, 0x06,
	0xd4, 0xdb, 0xcc, 0x7c, 0xbf, 0xf9, 0xc6, 0x1e, 0x1b, 0xf6, 0x9d, 0xac, 0xe8, 0x83, 0x92, 0x3d,
	0xb7, 0x7d, 0x25, 0x55, 0x57, 0x0b, 0x4b, 0x79, 0xdd, 0x73, 0x21, 0x6b, 0x32, 0xf2, 0x9a, 0xca,
	0xe1, 0x9e, 0x77, 0xb5, 0x66, 0xc6, 0xf0, 0xa1, 0xd3, 0xd5, 0xa8, 0xa4, 0x91, 0x68, 0xe5, 0xf4,
	0x77, 0x7f, 0xd6, 0x90, 0x1d, 0x83, 0x80, 0xf6, 0x90, 0x35, 0x96, 0x8b, 0xb6, 0x91, 0xa6, 0x48,
	0xb6, 0xc9, 0x2e, 0xbf, 0x7b, 0x55, 0x39, 0xaa, 0x8a, 0x44, 0x75, 0x08, 0x32, 0x9e, 0x41, 0xf4,
	0x09, 0xf2, 0x29, 0xb6, 0xf4, 0x17, 0x33, 0xc5, 0x95, 0xef, 0x2b, 0x2f, 0xf6, 0x79, 0x02, 0x2f,
	0x71, 0x37, 0x52, 0x3f, 0x12, 0xd5, 0xf3, 0xa1, 0x2b, 0xd2, 0x8b, 0x23, 0x8f, 0x41, 0xc6, 0x33,
	0x88, 0x3e, 0x42, 0xae, 0xa5, 0x55, 0x94, 0x9d, 0x08, 0x15, 0xba, 0x58, 0x6d, 0xd3, 0x5d, 0x7e,
	0xf7, 0xfa, 0xff, 0x3e, 0x4f, 0x7c, 0xa6, 0x42, 0x63, 0xd0, 0x73, 0xec, 0x06, 0x2a, 0xa6, 0xad,
	0x30, 0x6d, 0x53, 0xac, 0x2f, 0x0e, 0xc4, 0x5e, 0xfe, 0x72, 0xc0, 0x33, 0x58, 0xfe, 0x4e, 0x20,
	0x8b, 0x57, 0x47, 0x1f, 0xe0, 0x39, 0x1f, 0x0c, 0x53, 0x03, 0x11, 0x27, 0xc5, 0x48, 0xcb, 0x94,
	0x5f, 0xd6, 0x06, 0x3f, 0x8b, 0x65, 0xec, 0xab, 0xa8, 0x86, 0x17, 0xa3, 0x6d, 0x04, 0xa7, 0x27,
	0x6d, 0x1b, 0x4d, 0x15, 0x1f, 0x0d, 0x97, 0x83, 0xdf, 0xd0, 0x06, 0xa3, 0x49, 0x3a, 0x2e, 0x14,
	0xb4, 0x87, 0x97, 0xb3, 0xf3, 0x3f, 0x2d, 0xa9, 0x6f, 0xb9, 0x8d, 0xe2, 0xb2, 0xa9, 0xfc, 0x06,
	0xf9, 0x62, 0xbb, 0x08, 0xc1, 0x6a, 0x20, 0x3d, 0x0b, 0x47, 0xf2, 0xb1, 0xab, 0x3d, 0x48, 0x6d,
	0xc2, 0x64, 0x1f, 0xa3, 0x02, 0x9e, 0x8c, 0x4a, 0xfe, 0x64, 0xd4, 0x04, 0xf7, 0x98, 0x96, 0x18,
	0xb2, 0xb8, 0x73, 0xf4, 0x16, 0x6e, 0x5a, 0x76, 0x4f, 0xac, 0x30, 0x27, 0xef, 0x30, 0xb9, 0xe6,
	0xa1, 0xf6, 0xd5, 0x19, 0xbd, 0x87, 0xa7, 0x44, 0x08, 0xf9, 0xc8, 0x5a, 0x8f, 0xe8, 0xe2, 0x6a,
	0x9b, 0xee, 0x36, 0xf8, 0x26, 0x14, 0x1d, 0xa3, 0xcb, 0xef, 0x00, 0xe7, 0xf7, 0x40, 0xb7, 0xb0,
	0x9e, 0xd0, 0xc4, 0xa3, 0x53, 0x82, 0x4a, 0xc8, 0xc2, 0x11, 0xa2, 0xc7, 0x9c, 0xbb, 0xd3, 0x4e,
	0xab, 0xd6, 0x45, 0xea, 0xa5, 0x98, 0x96, 0x6f, 0x20, 0x8b, 0x0f, 0x36, 0xdf, 0x33, 0x39, 0xdf,
	0xf3, 0x90, 0xfd, 0xb8, 0x9e, 0xfe, 0x7f, 0x73, 0xed, 0xff, 0xfd, 0xfe, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x6f, 0x94, 0x7b, 0x4b, 0x2e, 0x03, 0x00, 0x00,
}
