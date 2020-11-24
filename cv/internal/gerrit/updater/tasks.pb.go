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
// source: go.chromium.org/luci/cv/internal/gerrit/updater/tasks.proto

package updater

import (
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

// RefreshGerritCL fetches latest Gerrit data and saves it to a CL snapshot.
//
// Queue: "refresh-gerrit-cl".
type RefreshGerritCL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LuciProject string `protobuf:"bytes,1,opt,name=luci_project,json=luciProject,proto3" json:"luci_project,omitempty"`
	Host        string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Change      int64  `protobuf:"varint,3,opt,name=change,proto3" json:"change,omitempty"`
	// Optional fields.
	UpdatedHint *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_hint,json=updatedHint,proto3" json:"updated_hint,omitempty"`
	ClidHint    int64                  `protobuf:"varint,5,opt,name=clid_hint,json=clidHint,proto3" json:"clid_hint,omitempty"`
}

func (x *RefreshGerritCL) Reset() {
	*x = RefreshGerritCL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshGerritCL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshGerritCL) ProtoMessage() {}

func (x *RefreshGerritCL) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshGerritCL.ProtoReflect.Descriptor instead.
func (*RefreshGerritCL) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *RefreshGerritCL) GetLuciProject() string {
	if x != nil {
		return x.LuciProject
	}
	return ""
}

func (x *RefreshGerritCL) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RefreshGerritCL) GetChange() int64 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *RefreshGerritCL) GetUpdatedHint() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedHint
	}
	return nil
}

func (x *RefreshGerritCL) GetClidHint() int64 {
	if x != nil {
		return x.ClidHint
	}
	return 0
}

var File_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63, 0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x72, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x2e, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x43, 0x4c, 0x12, 0x21, 0x0a, 0x0c, 0x6c,
	0x75, 0x63, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6c, 0x75, 0x63, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69,
	0x64, 0x5f, 0x68, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x64, 0x48, 0x69, 0x6e, 0x74, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x63,
	0x76, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x72, 0x72, 0x69,
	0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x3b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDescData = file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDesc
)

func file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDescData)
	})
	return file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDescData
}

var file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_goTypes = []interface{}{
	(*RefreshGerritCL)(nil),       // 0: internal.gerrit.updater.RefreshGerritCL
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_depIdxs = []int32{
	1, // 0: internal.gerrit.updater.RefreshGerritCL.updated_hint:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_init() }
func file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_init() {
	if File_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshGerritCL); i {
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
			RawDescriptor: file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto = out.File
	file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_rawDesc = nil
	file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_goTypes = nil
	file_go_chromium_org_luci_cv_internal_gerrit_updater_tasks_proto_depIdxs = nil
}
