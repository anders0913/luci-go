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
// source: go.chromium.org/luci/led/job/job.proto

package job

import (
	proto "go.chromium.org/luci/buildbucket/proto"
	api "go.chromium.org/luci/swarming/proto/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Buildbucket is, ideally, just a BBAgentArgs, but there are bits of data that
// led needs to track which aren't currently contained in BBAgentArgs.
//
// Where it makes sense, this additional data should be moved from this
// Buildbucket message into BBAgentArgs, but for now we store it separately to
// get led v2 up and running.
type Buildbucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BbagentArgs      *proto.BBAgentArgs    `protobuf:"bytes,1,opt,name=bbagent_args,json=bbagentArgs,proto3" json:"bbagent_args,omitempty"`
	CipdPackages     []*api.CIPDPackage    `protobuf:"bytes,2,rep,name=cipd_packages,json=cipdPackages,proto3" json:"cipd_packages,omitempty"`
	EnvVars          []*api.StringPair     `protobuf:"bytes,3,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty"`
	EnvPrefixes      []*api.StringListPair `protobuf:"bytes,4,rep,name=env_prefixes,json=envPrefixes,proto3" json:"env_prefixes,omitempty"`
	ExtraTags        []string              `protobuf:"bytes,5,rep,name=extra_tags,json=extraTags,proto3" json:"extra_tags,omitempty"`
	GracePeriod      *durationpb.Duration  `protobuf:"bytes,6,opt,name=grace_period,json=gracePeriod,proto3" json:"grace_period,omitempty"`
	BotPingTolerance *durationpb.Duration  `protobuf:"bytes,7,opt,name=bot_ping_tolerance,json=botPingTolerance,proto3" json:"bot_ping_tolerance,omitempty"`
	Containment      *api.Containment      `protobuf:"bytes,8,opt,name=containment,proto3" json:"containment,omitempty"`
	// Indicates that this build should be generated as a legacy kitchen task when
	// launched.
	LegacyKitchen bool `protobuf:"varint,9,opt,name=legacy_kitchen,json=legacyKitchen,proto3" json:"legacy_kitchen,omitempty"`
	// Eventually becomes the name of the launched swarming task.
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// This field contains the path relative to ${ISOLATED_OUTDIR} for the final
	// build.proto result. If blank, will cause the job not to emit any build
	// proto to the output directory.
	//
	// For bbagent-based jobs this must have the file extension ".pb", ".textpb"
	// or ".json", to get the respective encoding.
	//
	// For legacy kitchen jobs this must have the file extension ".json".
	//
	// By default, led will populate this with "build.proto.json".
	FinalBuildProtoPath string `protobuf:"bytes,11,opt,name=final_build_proto_path,json=finalBuildProtoPath,proto3" json:"final_build_proto_path,omitempty"`
}

func (x *Buildbucket) Reset() {
	*x = Buildbucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_led_job_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Buildbucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Buildbucket) ProtoMessage() {}

func (x *Buildbucket) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_led_job_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Buildbucket.ProtoReflect.Descriptor instead.
func (*Buildbucket) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_led_job_job_proto_rawDescGZIP(), []int{0}
}

func (x *Buildbucket) GetBbagentArgs() *proto.BBAgentArgs {
	if x != nil {
		return x.BbagentArgs
	}
	return nil
}

func (x *Buildbucket) GetCipdPackages() []*api.CIPDPackage {
	if x != nil {
		return x.CipdPackages
	}
	return nil
}

func (x *Buildbucket) GetEnvVars() []*api.StringPair {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

func (x *Buildbucket) GetEnvPrefixes() []*api.StringListPair {
	if x != nil {
		return x.EnvPrefixes
	}
	return nil
}

func (x *Buildbucket) GetExtraTags() []string {
	if x != nil {
		return x.ExtraTags
	}
	return nil
}

func (x *Buildbucket) GetGracePeriod() *durationpb.Duration {
	if x != nil {
		return x.GracePeriod
	}
	return nil
}

func (x *Buildbucket) GetBotPingTolerance() *durationpb.Duration {
	if x != nil {
		return x.BotPingTolerance
	}
	return nil
}

func (x *Buildbucket) GetContainment() *api.Containment {
	if x != nil {
		return x.Containment
	}
	return nil
}

func (x *Buildbucket) GetLegacyKitchen() bool {
	if x != nil {
		return x.LegacyKitchen
	}
	return false
}

func (x *Buildbucket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Buildbucket) GetFinalBuildProtoPath() string {
	if x != nil {
		return x.FinalBuildProtoPath
	}
	return ""
}

// Swarming is the raw TaskRequest. When a Definition is in this form, the
// user's ability to manipulate it via `led` subcommands is extremely limited.
type Swarming struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task     *api.TaskRequest `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Hostname string           `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *Swarming) Reset() {
	*x = Swarming{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_led_job_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Swarming) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Swarming) ProtoMessage() {}

func (x *Swarming) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_led_job_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Swarming.ProtoReflect.Descriptor instead.
func (*Swarming) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_led_job_job_proto_rawDescGZIP(), []int{1}
}

func (x *Swarming) GetTask() *api.TaskRequest {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *Swarming) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type Definition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to JobType:
	//	*Definition_Buildbucket
	//	*Definition_Swarming
	JobType isDefinition_JobType `protobuf_oneof:"job_type"`
	// WILL BE DEPRECATED AFTER MIGRATION TO RBE-CAS (crbug.com/1145959).
	//
	// If set, this holds the CASTree to use with the build, when launched.
	//
	// At the time of launch, this will be merged with
	// swarming.task_slice[*].properties.inputs_ref, if any.
	//
	// The 'server' and 'namespace' fields here are used as the defaults for any
	// digests specified without server/namespace.
	UserPayload *api.CASTree `protobuf:"bytes,3,opt,name=user_payload,json=userPayload,proto3" json:"user_payload,omitempty"`
	// [In Experiment]
	// If set, this holds the CASReference to use with the job, when launched.
	//
	// At the time of launch, this will be merged with
	// swarming.task_slice[*].properties.cas_input_root, if any.
	//
	// can only use either cas_user_payload or user_payload.
	CasUserPayload *api.CASReference `protobuf:"bytes,4,opt,name=cas_user_payload,json=casUserPayload,proto3" json:"cas_user_payload,omitempty"`
}

func (x *Definition) Reset() {
	*x = Definition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_led_job_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Definition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Definition) ProtoMessage() {}

func (x *Definition) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_led_job_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Definition.ProtoReflect.Descriptor instead.
func (*Definition) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_led_job_job_proto_rawDescGZIP(), []int{2}
}

func (m *Definition) GetJobType() isDefinition_JobType {
	if m != nil {
		return m.JobType
	}
	return nil
}

func (x *Definition) GetBuildbucket() *Buildbucket {
	if x, ok := x.GetJobType().(*Definition_Buildbucket); ok {
		return x.Buildbucket
	}
	return nil
}

func (x *Definition) GetSwarming() *Swarming {
	if x, ok := x.GetJobType().(*Definition_Swarming); ok {
		return x.Swarming
	}
	return nil
}

func (x *Definition) GetUserPayload() *api.CASTree {
	if x != nil {
		return x.UserPayload
	}
	return nil
}

func (x *Definition) GetCasUserPayload() *api.CASReference {
	if x != nil {
		return x.CasUserPayload
	}
	return nil
}

type isDefinition_JobType interface {
	isDefinition_JobType()
}

type Definition_Buildbucket struct {
	// Represents a buildbucket-native task; May be recovered from a swarming
	// task, or provided directly via buildbucket.
	Buildbucket *Buildbucket `protobuf:"bytes,1,opt,name=buildbucket,proto3,oneof"`
}

type Definition_Swarming struct {
	// Represents a swarming task. This will be filled for jobs sourced directly
	// from swarming which weren't recognized as a buildbucket task.
	//
	// A limited subset of the edit and info functionality is available for
	// raw swarming jobs.
	Swarming *Swarming `protobuf:"bytes,2,opt,name=swarming,proto3,oneof"`
}

func (*Definition_Buildbucket) isDefinition_JobType() {}

func (*Definition_Swarming) isDefinition_JobType() {}

var File_go_chromium_org_luci_led_job_job_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_led_job_job_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x65, 0x64, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x6a,
	0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6a, 0x6f, 0x62, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x67,
	0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c,
	0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75,
	0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x73, 0x77, 0x61, 0x72, 0x6d,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x77,
	0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x04, 0x0a,
	0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x3e, 0x0a, 0x0c,
	0x62, 0x62, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x42, 0x42, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x73, 0x52,
	0x0b, 0x62, 0x62, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x3d, 0x0a, 0x0d,
	0x63, 0x69, 0x70, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x49, 0x50, 0x44, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x63,
	0x69, 0x70, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x65,
	0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x12,
	0x3e, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x54, 0x61, 0x67, 0x73, 0x12, 0x3c,
	0x0a, 0x0c, 0x67, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x67, 0x72, 0x61, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x47, 0x0a, 0x12,
	0x62, 0x6f, 0x74, 0x5f, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x10, 0x62, 0x6f, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6c, 0x65,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x77, 0x61,
	0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x6b, 0x69, 0x74, 0x63,
	0x68, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x74,
	0x68, 0x22, 0x54, 0x0a, 0x08, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x77,
	0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf9, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x00, 0x52,
	0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x08,
	0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52,
	0x08, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x37, 0x0a, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x41,
	0x53, 0x54, 0x72, 0x65, 0x65, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x43, 0x0a, 0x10, 0x63, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x41, 0x53, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x63, 0x61, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x6c, 0x65, 0x64, 0x2f,
	0x6a, 0x6f, 0x62, 0x3b, 0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_led_job_job_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_led_job_job_proto_rawDescData = file_go_chromium_org_luci_led_job_job_proto_rawDesc
)

func file_go_chromium_org_luci_led_job_job_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_led_job_job_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_led_job_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_led_job_job_proto_rawDescData)
	})
	return file_go_chromium_org_luci_led_job_job_proto_rawDescData
}

var file_go_chromium_org_luci_led_job_job_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_go_chromium_org_luci_led_job_job_proto_goTypes = []interface{}{
	(*Buildbucket)(nil),         // 0: job.Buildbucket
	(*Swarming)(nil),            // 1: job.Swarming
	(*Definition)(nil),          // 2: job.Definition
	(*proto.BBAgentArgs)(nil),   // 3: buildbucket.v2.BBAgentArgs
	(*api.CIPDPackage)(nil),     // 4: swarming.v1.CIPDPackage
	(*api.StringPair)(nil),      // 5: swarming.v1.StringPair
	(*api.StringListPair)(nil),  // 6: swarming.v1.StringListPair
	(*durationpb.Duration)(nil), // 7: google.protobuf.Duration
	(*api.Containment)(nil),     // 8: swarming.v1.Containment
	(*api.TaskRequest)(nil),     // 9: swarming.v1.TaskRequest
	(*api.CASTree)(nil),         // 10: swarming.v1.CASTree
	(*api.CASReference)(nil),    // 11: swarming.v1.CASReference
}
var file_go_chromium_org_luci_led_job_job_proto_depIdxs = []int32{
	3,  // 0: job.Buildbucket.bbagent_args:type_name -> buildbucket.v2.BBAgentArgs
	4,  // 1: job.Buildbucket.cipd_packages:type_name -> swarming.v1.CIPDPackage
	5,  // 2: job.Buildbucket.env_vars:type_name -> swarming.v1.StringPair
	6,  // 3: job.Buildbucket.env_prefixes:type_name -> swarming.v1.StringListPair
	7,  // 4: job.Buildbucket.grace_period:type_name -> google.protobuf.Duration
	7,  // 5: job.Buildbucket.bot_ping_tolerance:type_name -> google.protobuf.Duration
	8,  // 6: job.Buildbucket.containment:type_name -> swarming.v1.Containment
	9,  // 7: job.Swarming.task:type_name -> swarming.v1.TaskRequest
	0,  // 8: job.Definition.buildbucket:type_name -> job.Buildbucket
	1,  // 9: job.Definition.swarming:type_name -> job.Swarming
	10, // 10: job.Definition.user_payload:type_name -> swarming.v1.CASTree
	11, // 11: job.Definition.cas_user_payload:type_name -> swarming.v1.CASReference
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_led_job_job_proto_init() }
func file_go_chromium_org_luci_led_job_job_proto_init() {
	if File_go_chromium_org_luci_led_job_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_led_job_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Buildbucket); i {
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
		file_go_chromium_org_luci_led_job_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Swarming); i {
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
		file_go_chromium_org_luci_led_job_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Definition); i {
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
	file_go_chromium_org_luci_led_job_job_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Definition_Buildbucket)(nil),
		(*Definition_Swarming)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_led_job_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_led_job_job_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_led_job_job_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_led_job_job_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_led_job_job_proto = out.File
	file_go_chromium_org_luci_led_job_job_proto_rawDesc = nil
	file_go_chromium_org_luci_led_job_job_proto_goTypes = nil
	file_go_chromium_org_luci_led_job_job_proto_depIdxs = nil
}
