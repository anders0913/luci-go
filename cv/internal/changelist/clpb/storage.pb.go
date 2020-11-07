// Copyright 2020 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/cv/internal/changelist/clpb/storage.proto

package clpb

import (
	gerrit "go.chromium.org/luci/common/proto/gerrit"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DepKind int32

const (
	DepKind_DEP_KIND_UNSPECIFIED DepKind = 0
	// Dep MUST be patched in / submitted before the dependent CL.
	DepKind_HARD DepKind = 1
	// Dep SHOULD be patched in / submitted before the dependent CL,
	// but doesn't have to be.
	DepKind_SOFT DepKind = 2
)

// Enum value maps for DepKind.
var (
	DepKind_name = map[int32]string{
		0: "DEP_KIND_UNSPECIFIED",
		1: "HARD",
		2: "SOFT",
	}
	DepKind_value = map[string]int32{
		"DEP_KIND_UNSPECIFIED": 0,
		"HARD":                 1,
		"SOFT":                 2,
	}
)

func (x DepKind) Enum() *DepKind {
	p := new(DepKind)
	*p = x
	return p
}

func (x DepKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DepKind) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_enumTypes[0].Descriptor()
}

func (DepKind) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_enumTypes[0]
}

func (x DepKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DepKind.Descriptor instead.
func (DepKind) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescGZIP(), []int{0}
}

// Snapshot stores a snapshot of CL info as seen by CV at a certain time.
//
// When stored in CL entity, represents latest known Gerrit data.
// When stored in RunCL entity, represents data pertaining to a fixed patchset.
type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The timestamp from external system.
	// Used to determine if re-querying external system is needed.
	ExternalUpdateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=external_update_time,json=externalUpdateTime,proto3" json:"external_update_time,omitempty"`
	// Resolved dependencies of a CL.
	Deps []*Dep `protobuf:"bytes,2,rep,name=deps,proto3" json:"deps,omitempty"`
	// CL-kind specific data.
	//
	// Types that are assignable to Kind:
	//	*Snapshot_Gerrit
	Kind isSnapshot_Kind `protobuf_oneof:"kind"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescGZIP(), []int{0}
}

func (x *Snapshot) GetExternalUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExternalUpdateTime
	}
	return nil
}

func (x *Snapshot) GetDeps() []*Dep {
	if x != nil {
		return x.Deps
	}
	return nil
}

func (m *Snapshot) GetKind() isSnapshot_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Snapshot) GetGerrit() *Gerrit {
	if x, ok := x.GetKind().(*Snapshot_Gerrit); ok {
		return x.Gerrit
	}
	return nil
}

type isSnapshot_Kind interface {
	isSnapshot_Kind()
}

type Snapshot_Gerrit struct {
	Gerrit *Gerrit `protobuf:"bytes,11,opt,name=gerrit,proto3,oneof"`
}

func (*Snapshot_Gerrit) isSnapshot_Kind() {}

type Dep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CLID is internal CV ID of a CL which is the dependency.
	Clid int64   `protobuf:"varint,1,opt,name=clid,proto3" json:"clid,omitempty"`
	Kind DepKind `protobuf:"varint,2,opt,name=kind,proto3,enum=cv.clpb.DepKind" json:"kind,omitempty"`
}

func (x *Dep) Reset() {
	*x = Dep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dep) ProtoMessage() {}

func (x *Dep) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dep.ProtoReflect.Descriptor instead.
func (*Dep) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescGZIP(), []int{1}
}

func (x *Dep) GetClid() int64 {
	if x != nil {
		return x.Clid
	}
	return 0
}

func (x *Dep) GetKind() DepKind {
	if x != nil {
		return x.Kind
	}
	return DepKind_DEP_KIND_UNSPECIFIED
}

type Gerrit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Info contains all revisions, but non-current revisions will not have all
	// the fields populated.
	//
	// Exact fields TODO.
	Info *gerrit.ChangeInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	// Files are filenames touched in the current revision.
	//
	// It's derived frm gerrit.ListFilesResponse, see
	// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-files.
	Files []string `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
	// Git dependencies of the current revision.
	GitDeps []*GerritGitDep `protobuf:"bytes,3,rep,name=git_deps,json=gitDeps,proto3" json:"git_deps,omitempty"`
	// Free-form dependencies. Currently, sourced from CQ-Depend footers.
	// In the future, this may be derived from Gerrit hashtags, topics, or other
	// mechanisms.
	SoftDeps []*GerritSoftDep `protobuf:"bytes,4,rep,name=soft_deps,json=softDeps,proto3" json:"soft_deps,omitempty"`
}

func (x *Gerrit) Reset() {
	*x = Gerrit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gerrit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gerrit) ProtoMessage() {}

func (x *Gerrit) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gerrit.ProtoReflect.Descriptor instead.
func (*Gerrit) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescGZIP(), []int{2}
}

func (x *Gerrit) GetInfo() *gerrit.ChangeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Gerrit) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Gerrit) GetGitDeps() []*GerritGitDep {
	if x != nil {
		return x.GitDeps
	}
	return nil
}

func (x *Gerrit) GetSoftDeps() []*GerritSoftDep {
	if x != nil {
		return x.SoftDeps
	}
	return nil
}

// GerritGitDep is a dependency discovered via Git child->parent chain for one Gerrit CL.
type GerritGitDep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Gerrit Change number.
	Change int32 `protobuf:"varint,1,opt,name=change,proto3" json:"change,omitempty"`
	// Immediate is set iff this dep is an immediate parent of the Gerrit CL.
	//
	// Immediate dep must be submitted before its child.
	// Non-immediate CLs don't necessarily have to be submitted before:
	//   for example, for a chain <base> <- A1 <- B1 <- C1 <- D1
	//   D1's deps are [A,B,C] but only C is immediate, and 1 stands for patchset.
	//   Developer may then swap B,C without re-uploading D (say, to avoid
	//   patchset churn), resulting in a new logical chain:
	//      <base> <- A1 <- C2 <- B2
	//                   \
	//                    <- B1 <- C1 <- D1
	//
	//   In this case, Gerrit's related changes for D1 will still return A1,B1,C1,
	//   which CV interprets as C must be landed before D, while B and A should
	//   be landed before D.
	//
	// TODO(tandrii): this is replicating existing CQDaemon logic. I think
	// it'd be reasonable to treat all (A,B,C) as MUST BE submitted before D.
	Immediate bool `protobuf:"varint,2,opt,name=immediate,proto3" json:"immediate,omitempty"`
}

func (x *GerritGitDep) Reset() {
	*x = GerritGitDep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GerritGitDep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GerritGitDep) ProtoMessage() {}

func (x *GerritGitDep) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GerritGitDep.ProtoReflect.Descriptor instead.
func (*GerritGitDep) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescGZIP(), []int{3}
}

func (x *GerritGitDep) GetChange() int32 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *GerritGitDep) GetImmediate() bool {
	if x != nil {
		return x.Immediate
	}
	return false
}

type GerritSoftDep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Gerrit host.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Gerrit change number.
	Change int32 `protobuf:"varint,2,opt,name=change,proto3" json:"change,omitempty"`
}

func (x *GerritSoftDep) Reset() {
	*x = GerritSoftDep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GerritSoftDep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GerritSoftDep) ProtoMessage() {}

func (x *GerritSoftDep) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GerritSoftDep.ProtoReflect.Descriptor instead.
func (*GerritSoftDep) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescGZIP(), []int{4}
}

func (x *GerritSoftDep) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GerritSoftDep) GetChange() int32 {
	if x != nil {
		return x.Change
	}
	return 0
}

var File_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x63, 0x6c,
	0x70, 0x62, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x63, 0x76, 0x2e, 0x63, 0x6c, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x72, 0x72, 0x69, 0x74, 0x2f, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xad, 0x01, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x4c,
	0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x04,
	0x64, 0x65, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x76, 0x2e,
	0x63, 0x6c, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x52, 0x04, 0x64, 0x65, 0x70, 0x73, 0x12, 0x29,
	0x0a, 0x06, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x76, 0x2e, 0x63, 0x6c, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x48,
	0x00, 0x52, 0x06, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x22, 0x3f, 0x0a, 0x03, 0x44, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6c, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6c, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x76, 0x2e,
	0x63, 0x6c, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x70, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x22, 0xad, 0x01, 0x0a, 0x06, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x12, 0x26, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65,
	0x72, 0x72, 0x69, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x67,
	0x69, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x63, 0x76, 0x2e, 0x63, 0x6c, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x47, 0x69,
	0x74, 0x44, 0x65, 0x70, 0x52, 0x07, 0x67, 0x69, 0x74, 0x44, 0x65, 0x70, 0x73, 0x12, 0x33, 0x0a,
	0x09, 0x73, 0x6f, 0x66, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x63, 0x76, 0x2e, 0x63, 0x6c, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x72, 0x72, 0x69,
	0x74, 0x53, 0x6f, 0x66, 0x74, 0x44, 0x65, 0x70, 0x52, 0x08, 0x73, 0x6f, 0x66, 0x74, 0x44, 0x65,
	0x70, 0x73, 0x22, 0x44, 0x0a, 0x0c, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x47, 0x69, 0x74, 0x44,
	0x65, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x0d, 0x47, 0x65, 0x72, 0x72,
	0x69, 0x74, 0x53, 0x6f, 0x66, 0x74, 0x44, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2a, 0x37, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x18, 0x0a, 0x14, 0x44, 0x45, 0x50, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x41,
	0x52, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4f, 0x46, 0x54, 0x10, 0x02, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x63, 0x6c,
	0x70, 0x62, 0x3b, 0x63, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescData = file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_goTypes = []interface{}{
	(DepKind)(0),                  // 0: cv.clpb.DepKind
	(*Snapshot)(nil),              // 1: cv.clpb.Snapshot
	(*Dep)(nil),                   // 2: cv.clpb.Dep
	(*Gerrit)(nil),                // 3: cv.clpb.Gerrit
	(*GerritGitDep)(nil),          // 4: cv.clpb.GerritGitDep
	(*GerritSoftDep)(nil),         // 5: cv.clpb.GerritSoftDep
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*gerrit.ChangeInfo)(nil),     // 7: gerrit.ChangeInfo
}
var file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_depIdxs = []int32{
	6, // 0: cv.clpb.Snapshot.external_update_time:type_name -> google.protobuf.Timestamp
	2, // 1: cv.clpb.Snapshot.deps:type_name -> cv.clpb.Dep
	3, // 2: cv.clpb.Snapshot.gerrit:type_name -> cv.clpb.Gerrit
	0, // 3: cv.clpb.Dep.kind:type_name -> cv.clpb.DepKind
	7, // 4: cv.clpb.Gerrit.info:type_name -> gerrit.ChangeInfo
	4, // 5: cv.clpb.Gerrit.git_deps:type_name -> cv.clpb.GerritGitDep
	5, // 6: cv.clpb.Gerrit.soft_deps:type_name -> cv.clpb.GerritSoftDep
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_init() }
func file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_init() {
	if File_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
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
		file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dep); i {
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
		file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gerrit); i {
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
		file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GerritGitDep); i {
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
		file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GerritSoftDep); i {
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
	file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Snapshot_Gerrit)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto = out.File
	file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_changelist_clpb_storage_proto_depIdxs = nil
}
