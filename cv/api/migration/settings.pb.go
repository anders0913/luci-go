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
// source: go.chromium.org/luci/cv/api/migration/settings.proto

package migrationpb

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

// Settings is schema of service-wide commit-queue/migration-settings.cfg which
// is used only during migration. It applies to all LUCI projects and is read by
// CQDaemon and LUCI CV.
type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CQDaemon doesn't really have a -dev version, therefore to test -dev of CV,
	// production CQDaemon can connect to both prod and dev migration API.
	ApiHosts  []*Settings_ApiHost `protobuf:"bytes,1,rep,name=api_hosts,json=apiHosts,proto3" json:"api_hosts,omitempty"`
	UseCvRuns *Settings_UseCVRuns `protobuf:"bytes,3,opt,name=use_cv_runs,json=useCvRuns,proto3" json:"use_cv_runs,omitempty"`
	// TODO(tandrii): move this off migration-specific settings once CQDaemon is
	// shut down. This is located here only to avoid extra throw away code in
	// CQDaemon to read & refresh these from a different file.
	PssaMigration *PSSAMigration `protobuf:"bytes,2,opt,name=pssa_migration,json=pssaMigration,proto3" json:"pssa_migration,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDescGZIP(), []int{0}
}

func (x *Settings) GetApiHosts() []*Settings_ApiHost {
	if x != nil {
		return x.ApiHosts
	}
	return nil
}

func (x *Settings) GetUseCvRuns() *Settings_UseCVRuns {
	if x != nil {
		return x.UseCvRuns
	}
	return nil
}

func (x *Settings) GetPssaMigration() *PSSAMigration {
	if x != nil {
		return x.PssaMigration
	}
	return nil
}

type PSSAMigration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of LUCI Projects which must ues legacy ~/.netrc credentials,
	// because although they have registered project-scoped service account
	// (PSSA), their Gerrit ACLs aren't ready yet.
	ProjectsBlocklist []string `protobuf:"bytes,1,rep,name=projects_blocklist,json=projectsBlocklist,proto3" json:"projects_blocklist,omitempty"`
}

func (x *PSSAMigration) Reset() {
	*x = PSSAMigration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PSSAMigration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PSSAMigration) ProtoMessage() {}

func (x *PSSAMigration) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PSSAMigration.ProtoReflect.Descriptor instead.
func (*PSSAMigration) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDescGZIP(), []int{1}
}

func (x *PSSAMigration) GetProjectsBlocklist() []string {
	if x != nil {
		return x.ProjectsBlocklist
	}
	return nil
}

type Settings_ApiHost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// e.g. luci-change-verifier-dev.appspot.com.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// If a LUCI Project matches any of the regexps,
	// CQDaemon will connect to the above Migration API host.
	ProjectRegexp []string `protobuf:"bytes,2,rep,name=project_regexp,json=projectRegexp,proto3" json:"project_regexp,omitempty"`
	// If true and several hosts are configured, all other hosts' responses are
	// ignored.
	Prod bool `protobuf:"varint,3,opt,name=prod,proto3" json:"prod,omitempty"`
}

func (x *Settings_ApiHost) Reset() {
	*x = Settings_ApiHost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings_ApiHost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings_ApiHost) ProtoMessage() {}

func (x *Settings_ApiHost) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings_ApiHost.ProtoReflect.Descriptor instead.
func (*Settings_ApiHost) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Settings_ApiHost) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Settings_ApiHost) GetProjectRegexp() []string {
	if x != nil {
		return x.ProjectRegexp
	}
	return nil
}

func (x *Settings_ApiHost) GetProd() bool {
	if x != nil {
		return x.Prod
	}
	return false
}

// Determines which projects should start relying on CV for computing
// Runs to work on.
type Settings_UseCVRuns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectRegexp []string `protobuf:"bytes,1,rep,name=project_regexp,json=projectRegexp,proto3" json:"project_regexp,omitempty"`
}

func (x *Settings_UseCVRuns) Reset() {
	*x = Settings_UseCVRuns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings_UseCVRuns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings_UseCVRuns) ProtoMessage() {}

func (x *Settings_UseCVRuns) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings_UseCVRuns.ProtoReflect.Descriptor instead.
func (*Settings_UseCVRuns) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Settings_UseCVRuns) GetProjectRegexp() []string {
	if x != nil {
		return x.ProjectRegexp
	}
	return nil
}

var File_go_chromium_org_luci_cv_api_migration_settings_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDesc = []byte{
	0x0a, 0x34, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xd2, 0x02, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x38,
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x08,
	0x61, 0x70, 0x69, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x5f,
	0x63, 0x76, 0x5f, 0x72, 0x75, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x43, 0x56, 0x52, 0x75, 0x6e, 0x73, 0x52, 0x09, 0x75, 0x73,
	0x65, 0x43, 0x76, 0x52, 0x75, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x0e, 0x70, 0x73, 0x73, 0x61, 0x5f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x53, 0x53, 0x41,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x73, 0x73, 0x61, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x58, 0x0a, 0x07, 0x41, 0x70, 0x69, 0x48,
	0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x72, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x70, 0x72,
	0x6f, 0x64, 0x1a, 0x32, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x43, 0x56, 0x52, 0x75, 0x6e, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x22, 0x3e, 0x0a, 0x0d, 0x50, 0x53, 0x53, 0x41, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63,
	0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDescData = file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDesc
)

func file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDescData
}

var file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_chromium_org_luci_cv_api_migration_settings_proto_goTypes = []interface{}{
	(*Settings)(nil),           // 0: migration.Settings
	(*PSSAMigration)(nil),      // 1: migration.PSSAMigration
	(*Settings_ApiHost)(nil),   // 2: migration.Settings.ApiHost
	(*Settings_UseCVRuns)(nil), // 3: migration.Settings.UseCVRuns
}
var file_go_chromium_org_luci_cv_api_migration_settings_proto_depIdxs = []int32{
	2, // 0: migration.Settings.api_hosts:type_name -> migration.Settings.ApiHost
	3, // 1: migration.Settings.use_cv_runs:type_name -> migration.Settings.UseCVRuns
	1, // 2: migration.Settings.pssa_migration:type_name -> migration.PSSAMigration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_api_migration_settings_proto_init() }
func file_go_chromium_org_luci_cv_api_migration_settings_proto_init() {
	if File_go_chromium_org_luci_cv_api_migration_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
		file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PSSAMigration); i {
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
		file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings_ApiHost); i {
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
		file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings_UseCVRuns); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_api_migration_settings_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_api_migration_settings_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_api_migration_settings_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_api_migration_settings_proto = out.File
	file_go_chromium_org_luci_cv_api_migration_settings_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_api_migration_settings_proto_goTypes = nil
	file_go_chromium_org_luci_cv_api_migration_settings_proto_depIdxs = nil
}
