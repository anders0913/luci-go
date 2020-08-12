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
// source: go.chromium.org/luci/cv/api/bigquery/v1/attempt.proto

package bigquery

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Mode int32

const (
	// Default, never set.
	Mode_MODE_UNSPECIFIED Mode = 0
	// Run all tests but do not submit.
	Mode_DRY_RUN Mode = 1
	// Run all tests and potentially submit.
	Mode_FULL_RUN Mode = 2
)

// Enum value maps for Mode.
var (
	Mode_name = map[int32]string{
		0: "MODE_UNSPECIFIED",
		1: "DRY_RUN",
		2: "FULL_RUN",
	}
	Mode_value = map[string]int32{
		"MODE_UNSPECIFIED": 0,
		"DRY_RUN":          1,
		"FULL_RUN":         2,
	}
)

func (x Mode) Enum() *Mode {
	p := new(Mode)
	*p = x
	return p
}

func (x Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes[0].Descriptor()
}

func (Mode) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes[0]
}

func (x Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Mode.Descriptor instead.
func (Mode) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescGZIP(), []int{0}
}

type AttemptStatus int32

const (
	// Default, never set.
	AttemptStatus_ATTEMPT_STATUS_UNSPECIFIED AttemptStatus = 0
	// Started but not completed. Used by CQ API, TBD.
	AttemptStatus_STARTED AttemptStatus = 1
	// Ready to submit, all checks passed.
	AttemptStatus_SUCCESS AttemptStatus = 2
	// Attempt stopped before completion, due to some external event and not
	// a failure of the CLs to pass all tests. For example, this may happen
	// when a new patchset is uploaded, a CL is deleted, etc.
	AttemptStatus_ABORTED AttemptStatus = 3
	// Completed and failed some check. This may happen when a build failed,
	// footer syntax was incorrect, or CL was not approved.
	AttemptStatus_FAILURE AttemptStatus = 4
	// Failure in CQ itself caused the Attempt to be dropped.
	AttemptStatus_INFRA_FAILURE AttemptStatus = 5
)

// Enum value maps for AttemptStatus.
var (
	AttemptStatus_name = map[int32]string{
		0: "ATTEMPT_STATUS_UNSPECIFIED",
		1: "STARTED",
		2: "SUCCESS",
		3: "ABORTED",
		4: "FAILURE",
		5: "INFRA_FAILURE",
	}
	AttemptStatus_value = map[string]int32{
		"ATTEMPT_STATUS_UNSPECIFIED": 0,
		"STARTED":                    1,
		"SUCCESS":                    2,
		"ABORTED":                    3,
		"FAILURE":                    4,
		"INFRA_FAILURE":              5,
	}
)

func (x AttemptStatus) Enum() *AttemptStatus {
	p := new(AttemptStatus)
	*p = x
	return p
}

func (x AttemptStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttemptStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes[1].Descriptor()
}

func (AttemptStatus) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes[1]
}

func (x AttemptStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttemptStatus.Descriptor instead.
func (AttemptStatus) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescGZIP(), []int{1}
}

type AttemptSubstatus int32

const (
	// Default, never set.
	AttemptSubstatus_ATTEMPT_SUBSTATUS_UNSPECIFIED AttemptSubstatus = 0
	// There is no more detailed status set.
	AttemptSubstatus_NO_SUBSTATUS AttemptSubstatus = 1
	// Failed at least one critical tryjob.
	AttemptSubstatus_FAILED_TRYJOBS AttemptSubstatus = 2
	// Failed an initial quick check of CL and CL description state.
	AttemptSubstatus_FAILED_LINT AttemptSubstatus = 3
	// A CL didn't get sufficient approval for submitting via CQ.
	AttemptSubstatus_UNAPPROVED AttemptSubstatus = 4
	// A CQ triggerer doesn't have permission to trigger CQ.
	AttemptSubstatus_PERMISSION_DENIED AttemptSubstatus = 5
	// There was a problem with a dependency CL, e.g. some dependencies
	// were not submitted or not grouped together in this attempt.
	AttemptSubstatus_UNSATISFIED_DEPENDENCY AttemptSubstatus = 6
	// Aborted because of a manual cancelation.
	AttemptSubstatus_MANUAL_CANCEL AttemptSubstatus = 7
	// A request to buildbucket failed because CQ didn't have permission to
	// trigger builds.
	AttemptSubstatus_BUILDBUCKET_MISCONFIGURATION AttemptSubstatus = 8
)

// Enum value maps for AttemptSubstatus.
var (
	AttemptSubstatus_name = map[int32]string{
		0: "ATTEMPT_SUBSTATUS_UNSPECIFIED",
		1: "NO_SUBSTATUS",
		2: "FAILED_TRYJOBS",
		3: "FAILED_LINT",
		4: "UNAPPROVED",
		5: "PERMISSION_DENIED",
		6: "UNSATISFIED_DEPENDENCY",
		7: "MANUAL_CANCEL",
		8: "BUILDBUCKET_MISCONFIGURATION",
	}
	AttemptSubstatus_value = map[string]int32{
		"ATTEMPT_SUBSTATUS_UNSPECIFIED": 0,
		"NO_SUBSTATUS":                  1,
		"FAILED_TRYJOBS":                2,
		"FAILED_LINT":                   3,
		"UNAPPROVED":                    4,
		"PERMISSION_DENIED":             5,
		"UNSATISFIED_DEPENDENCY":        6,
		"MANUAL_CANCEL":                 7,
		"BUILDBUCKET_MISCONFIGURATION":  8,
	}
)

func (x AttemptSubstatus) Enum() *AttemptSubstatus {
	p := new(AttemptSubstatus)
	*p = x
	return p
}

func (x AttemptSubstatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttemptSubstatus) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes[2].Descriptor()
}

func (AttemptSubstatus) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes[2]
}

func (x AttemptSubstatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttemptSubstatus.Descriptor instead.
func (AttemptSubstatus) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescGZIP(), []int{2}
}

type GerritChange_SubmitStatus int32

const (
	// Default. Never set.
	GerritChange_SUBMIT_STATUS_UNSPECIFIED GerritChange_SubmitStatus = 0
	// CQ didn't try submitting this CL.
	//
	// Includes a case where CQ tried submitting the CL, but submission failed
	// due to transient error leaving CL as is, and CQ didn't try again.
	GerritChange_PENDING GerritChange_SubmitStatus = 1
	// CQ tried to submit, but got presumably transient errors and couldn't
	// ascertain whether submission was successful.
	//
	// It's possible that change was actually submitted, but CQ didn't receive
	// a confirmation from Gerrit and follow up checks of the change status
	// failed, too.
	GerritChange_UNKNOWN GerritChange_SubmitStatus = 2
	// CQ tried to submit, but Gerrit rejected the submission because this
	// Change can't be submitted.
	// Typically, this is because a rebase conflict needs to be resolved,
	// or rarely because the change needs some kind of approval.
	GerritChange_FAILURE GerritChange_SubmitStatus = 3
	// CQ submitted this change (aka "merged" in Gerrit jargon).
	//
	// Submission of Gerrit CLs in an Attempt is not an atomic operation,
	// so it's possible that only some of the GerritChanges are submitted.
	GerritChange_SUCCESS GerritChange_SubmitStatus = 4
)

// Enum value maps for GerritChange_SubmitStatus.
var (
	GerritChange_SubmitStatus_name = map[int32]string{
		0: "SUBMIT_STATUS_UNSPECIFIED",
		1: "PENDING",
		2: "UNKNOWN",
		3: "FAILURE",
		4: "SUCCESS",
	}
	GerritChange_SubmitStatus_value = map[string]int32{
		"SUBMIT_STATUS_UNSPECIFIED": 0,
		"PENDING":                   1,
		"UNKNOWN":                   2,
		"FAILURE":                   3,
		"SUCCESS":                   4,
	}
)

func (x GerritChange_SubmitStatus) Enum() *GerritChange_SubmitStatus {
	p := new(GerritChange_SubmitStatus)
	*p = x
	return p
}

func (x GerritChange_SubmitStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GerritChange_SubmitStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes[3].Descriptor()
}

func (GerritChange_SubmitStatus) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes[3]
}

func (x GerritChange_SubmitStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GerritChange_SubmitStatus.Descriptor instead.
func (GerritChange_SubmitStatus) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescGZIP(), []int{1, 0}
}

type Build_Origin int32

const (
	// Default. Never set.
	Build_ORIGIN_UNSPECIFIED Build_Origin = 0
	// Build was triggered as part of this attempt
	// because reuse was disabled for its builder.
	Build_NOT_REUSABLE Build_Origin = 1
	// Build was triggered as part of this attempt,
	// but if there was an already existing build it would have been reused.
	Build_NOT_REUSED Build_Origin = 2
	// Build was reused.
	Build_REUSED Build_Origin = 3
)

// Enum value maps for Build_Origin.
var (
	Build_Origin_name = map[int32]string{
		0: "ORIGIN_UNSPECIFIED",
		1: "NOT_REUSABLE",
		2: "NOT_REUSED",
		3: "REUSED",
	}
	Build_Origin_value = map[string]int32{
		"ORIGIN_UNSPECIFIED": 0,
		"NOT_REUSABLE":       1,
		"NOT_REUSED":         2,
		"REUSED":             3,
	}
)

func (x Build_Origin) Enum() *Build_Origin {
	p := new(Build_Origin)
	*p = x
	return p
}

func (x Build_Origin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Build_Origin) Descriptor() protoreflect.EnumDescriptor {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes[4].Descriptor()
}

func (Build_Origin) Type() protoreflect.EnumType {
	return &file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes[4]
}

func (x Build_Origin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Build_Origin.Descriptor instead.
func (Build_Origin) EnumDescriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescGZIP(), []int{2, 0}
}

// Attempt includes the state of one CQ attempt.
//
// An attempt involves doing checks for one or more CLs that could
// potentially be submitted together.
//
// Next ID: 13.
type Attempt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The opaque key unique to this Attempt.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The LUCI project that this Attempt belongs to.
	LuciProject string `protobuf:"bytes,2,opt,name=luci_project,json=luciProject,proto3" json:"luci_project,omitempty"`
	// The name of the config group that this Attempt belongs to.
	ConfigGroup string `protobuf:"bytes,11,opt,name=config_group,json=configGroup,proto3" json:"config_group,omitempty"`
	// An opaque key that is unique for a given set of Gerrit change patchsets.
	// (or, equivalently, buildsets). The same cl_group_key will be used if
	// another Attempt is made for the same set of changes at a different time.
	ClGroupKey string `protobuf:"bytes,3,opt,name=cl_group_key,json=clGroupKey,proto3" json:"cl_group_key,omitempty"`
	// Similar to cl_group_key, except the key will be the same when the
	// earliest_equivalent_patchset values are the same, even if the patchset
	// values are different.
	//
	// For example, when a new "trivial" patchset is uploaded, then the
	// cl_group_key will change but the equivalent_cl_group_key will stay the
	// same.
	EquivalentClGroupKey string `protobuf:"bytes,4,opt,name=equivalent_cl_group_key,json=equivalentClGroupKey,proto3" json:"equivalent_cl_group_key,omitempty"`
	// The time when the Attempt started (trigger time of the last CL triggered).
	StartTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time when the Attempt ended (released by CQ).
	EndTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Gerrit changes, with specific patchsets, in this Attempt.
	// There should be one or more.
	GerritChanges []*GerritChange `protobuf:"bytes,7,rep,name=gerrit_changes,json=gerritChanges,proto3" json:"gerrit_changes,omitempty"`
	// Relevant builds as of this Attempt's end time.
	//
	// While Attempt is processed, CQ may consider more builds than included here.
	//
	// For example, the following builds will be not be included:
	//   * builds triggered before this Attempt started, considered temporarily by
	//     CQ, but then ignored because they ultimately failed such that CQ had to
	//     trigger new builds instead.
	//   * successful builds which were fresh enough at the Attempt start time,
	//     but which were ignored after they became too old for consideration such
	//     that CQ had to trigger new builds instead.
	//   * builds triggered as part of this Attempt, which were later removed from
	//     project CQ config and hence were no longer required by CQ by Attempt
	//     end time.
	Builds []*Build `protobuf:"bytes,8,rep,name=builds,proto3" json:"builds,omitempty"`
	// Final status of the Attempt.
	Status AttemptStatus `protobuf:"varint,9,opt,name=status,proto3,enum=bigquery.AttemptStatus" json:"status,omitempty"`
	// A more fine-grained status the explains more details about the status.
	Substatus AttemptSubstatus `protobuf:"varint,10,opt,name=substatus,proto3,enum=bigquery.AttemptSubstatus" json:"substatus,omitempty"`
	// Whether or not the required builds for this attempt include additional
	// "opted-in" builders by the user via the `Cq-Include-Trybots` footer.
	HasCustomRequirement bool `protobuf:"varint,12,opt,name=has_custom_requirement,json=hasCustomRequirement,proto3" json:"has_custom_requirement,omitempty"`
}

func (x *Attempt) Reset() {
	*x = Attempt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attempt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attempt) ProtoMessage() {}

func (x *Attempt) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attempt.ProtoReflect.Descriptor instead.
func (*Attempt) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescGZIP(), []int{0}
}

func (x *Attempt) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Attempt) GetLuciProject() string {
	if x != nil {
		return x.LuciProject
	}
	return ""
}

func (x *Attempt) GetConfigGroup() string {
	if x != nil {
		return x.ConfigGroup
	}
	return ""
}

func (x *Attempt) GetClGroupKey() string {
	if x != nil {
		return x.ClGroupKey
	}
	return ""
}

func (x *Attempt) GetEquivalentClGroupKey() string {
	if x != nil {
		return x.EquivalentClGroupKey
	}
	return ""
}

func (x *Attempt) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Attempt) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Attempt) GetGerritChanges() []*GerritChange {
	if x != nil {
		return x.GerritChanges
	}
	return nil
}

func (x *Attempt) GetBuilds() []*Build {
	if x != nil {
		return x.Builds
	}
	return nil
}

func (x *Attempt) GetStatus() AttemptStatus {
	if x != nil {
		return x.Status
	}
	return AttemptStatus_ATTEMPT_STATUS_UNSPECIFIED
}

func (x *Attempt) GetSubstatus() AttemptSubstatus {
	if x != nil {
		return x.Substatus
	}
	return AttemptSubstatus_ATTEMPT_SUBSTATUS_UNSPECIFIED
}

func (x *Attempt) GetHasCustomRequirement() bool {
	if x != nil {
		return x.HasCustomRequirement
	}
	return false
}

// GerritChange represents one revision (patchset) of one Gerrit change
// in an Attempt.
//
// See also: GerritChange in buildbucket/proto/common.proto.
type GerritChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Gerrit hostname, e.g. "chromium-review.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Gerrit project, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Change number, e.g. 12345.
	Change int64 `protobuf:"varint,3,opt,name=change,proto3" json:"change,omitempty"`
	// Patch set number, e.g. 1.
	Patchset int64 `protobuf:"varint,4,opt,name=patchset,proto3" json:"patchset,omitempty"`
	// The earliest patchset of the CL that is considered equivalent to the
	// patchset above.
	EarliestEquivalentPatchset int64 `protobuf:"varint,5,opt,name=earliest_equivalent_patchset,json=earliestEquivalentPatchset,proto3" json:"earliest_equivalent_patchset,omitempty"`
	// The time that the CQ was triggered for this CL in this Attempt.
	TriggerTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=trigger_time,json=triggerTime,proto3" json:"trigger_time,omitempty"`
	// CQ Mode for this CL, e.g. dry run or full run.
	Mode Mode `protobuf:"varint,7,opt,name=mode,proto3,enum=bigquery.Mode" json:"mode,omitempty"`
	// Whether CQ tried to submit this change and the result of the operation.
	SubmitStatus GerritChange_SubmitStatus `protobuf:"varint,8,opt,name=submit_status,json=submitStatus,proto3,enum=bigquery.GerritChange_SubmitStatus" json:"submit_status,omitempty"`
}

func (x *GerritChange) Reset() {
	*x = GerritChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GerritChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GerritChange) ProtoMessage() {}

func (x *GerritChange) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GerritChange.ProtoReflect.Descriptor instead.
func (*GerritChange) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescGZIP(), []int{1}
}

func (x *GerritChange) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GerritChange) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GerritChange) GetChange() int64 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *GerritChange) GetPatchset() int64 {
	if x != nil {
		return x.Patchset
	}
	return 0
}

func (x *GerritChange) GetEarliestEquivalentPatchset() int64 {
	if x != nil {
		return x.EarliestEquivalentPatchset
	}
	return 0
}

func (x *GerritChange) GetTriggerTime() *timestamp.Timestamp {
	if x != nil {
		return x.TriggerTime
	}
	return nil
}

func (x *GerritChange) GetMode() Mode {
	if x != nil {
		return x.Mode
	}
	return Mode_MODE_UNSPECIFIED
}

func (x *GerritChange) GetSubmitStatus() GerritChange_SubmitStatus {
	if x != nil {
		return x.SubmitStatus
	}
	return GerritChange_SUBMIT_STATUS_UNSPECIFIED
}

// Build represents one tryjob Buildbucket build.
//
// See also: Build in buildbucket/proto/build.proto.
type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Buildbucket build ID, unique per Buildbucket instance.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Buildbucket host, e.g. "cr-buildbucket.appspot.com".
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// Information about whether this build was triggered previously and reused,
	// or triggered because there was no reusable build, or because builds by
	// this builder are all not reusable.
	Origin Build_Origin `protobuf:"varint,3,opt,name=origin,proto3,enum=bigquery.Build_Origin" json:"origin,omitempty"`
	// Whether the CQ must wait for this build to pass in order for the CLs to be
	// considered ready to submit. True means this builder must pass, false means
	// this builder is "optional", and so this build should not be used to assess
	// the correctness of the CLs in the Attempt. For example, builds added
	// because of the Cq-Include-Trybots footer are still critical; experimental
	// builders are not.
	//
	// Tip: join this with the Buildbucket BigQuery table to figure out which
	// builder this build belongs to.
	Critical bool `protobuf:"varint,4,opt,name=critical,proto3" json:"critical,omitempty"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescGZIP(), []int{2}
}

func (x *Build) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Build) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Build) GetOrigin() Build_Origin {
	if x != nil {
		return x.Origin
	}
	return Build_ORIGIN_UNSPECIFIED
}

func (x *Build) GetCritical() bool {
	if x != nil {
		return x.Critical
	}
	return false
}

var File_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb5, 0x04, 0x0a, 0x07, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x75, 0x63, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x75, 0x63, 0x69, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x6c, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x17, 0x65, 0x71, 0x75, 0x69,
	0x76, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65, 0x71, 0x75, 0x69, 0x76,
	0x61, 0x6c, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4b, 0x65, 0x79, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x0d, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x12, 0x27, 0x0a, 0x06, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x52, 0x06, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x53, 0x75, 0x62, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x68, 0x61, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x68, 0x61, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xc2, 0x03, 0x0a, 0x0c, 0x47,
	0x65, 0x72, 0x72, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x74, 0x63, 0x68, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x74, 0x63, 0x68, 0x73, 0x65, 0x74, 0x12, 0x40, 0x0a,
	0x1c, 0x65, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x71, 0x75, 0x69, 0x76, 0x61,
	0x6c, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x1a, 0x65, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x45, 0x71, 0x75,
	0x69, 0x76, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x73, 0x65, 0x74, 0x12,
	0x3d, 0x0a, 0x0c, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x62, 0x69, 0x67, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x61, 0x0a, 0x0c,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19,
	0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x04, 0x22,
	0xc7, 0x01, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x2e, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x22, 0x4e, 0x0a, 0x06, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4e,
	0x4f, 0x54, 0x5f, 0x52, 0x45, 0x55, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x55, 0x53, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x52, 0x45, 0x55, 0x53, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x37, 0x0a, 0x04, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x52, 0x59, 0x5f, 0x52,
	0x55, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x55, 0x4c, 0x4c, 0x5f, 0x52, 0x55, 0x4e,
	0x10, 0x02, 0x2a, 0x76, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x54, 0x54, 0x45, 0x4d, 0x50, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41,
	0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x46, 0x52, 0x41,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x05, 0x2a, 0xe4, 0x01, 0x0a, 0x10, 0x41,
	0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53, 0x75, 0x62, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x0a, 0x1d, 0x41, 0x54, 0x54, 0x45, 0x4d, 0x50, 0x54, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x54,
	0x52, 0x59, 0x4a, 0x4f, 0x42, 0x53, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x5f, 0x4c, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x41,
	0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x45, 0x52,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x05,
	0x12, 0x1a, 0x0a, 0x16, 0x55, 0x4e, 0x53, 0x41, 0x54, 0x49, 0x53, 0x46, 0x49, 0x45, 0x44, 0x5f,
	0x44, 0x45, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x4e, 0x43, 0x59, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d,
	0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x07, 0x12,
	0x20, 0x0a, 0x1c, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x5f, 0x4d,
	0x49, 0x53, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x08, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescData = file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDesc
)

func file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDescData
}

var file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_goTypes = []interface{}{
	(Mode)(0),                      // 0: bigquery.Mode
	(AttemptStatus)(0),             // 1: bigquery.AttemptStatus
	(AttemptSubstatus)(0),          // 2: bigquery.AttemptSubstatus
	(GerritChange_SubmitStatus)(0), // 3: bigquery.GerritChange.SubmitStatus
	(Build_Origin)(0),              // 4: bigquery.Build.Origin
	(*Attempt)(nil),                // 5: bigquery.Attempt
	(*GerritChange)(nil),           // 6: bigquery.GerritChange
	(*Build)(nil),                  // 7: bigquery.Build
	(*timestamp.Timestamp)(nil),    // 8: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_depIdxs = []int32{
	8,  // 0: bigquery.Attempt.start_time:type_name -> google.protobuf.Timestamp
	8,  // 1: bigquery.Attempt.end_time:type_name -> google.protobuf.Timestamp
	6,  // 2: bigquery.Attempt.gerrit_changes:type_name -> bigquery.GerritChange
	7,  // 3: bigquery.Attempt.builds:type_name -> bigquery.Build
	1,  // 4: bigquery.Attempt.status:type_name -> bigquery.AttemptStatus
	2,  // 5: bigquery.Attempt.substatus:type_name -> bigquery.AttemptSubstatus
	8,  // 6: bigquery.GerritChange.trigger_time:type_name -> google.protobuf.Timestamp
	0,  // 7: bigquery.GerritChange.mode:type_name -> bigquery.Mode
	3,  // 8: bigquery.GerritChange.submit_status:type_name -> bigquery.GerritChange.SubmitStatus
	4,  // 9: bigquery.Build.origin:type_name -> bigquery.Build.Origin
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_init() }
func file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_init() {
	if File_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attempt); i {
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
		file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GerritChange); i {
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
		file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_depIdxs,
		EnumInfos:         file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_enumTypes,
		MessageInfos:      file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto = out.File
	file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_goTypes = nil
	file_go_chromium_org_luci_cv_api_bigquery_v1_attempt_proto_depIdxs = nil
}