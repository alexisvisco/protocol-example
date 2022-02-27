// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/common/v1/project_state.proto

package commonpb

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

type ProjectState int32

const (
	ProjectState_PROJECT_STATE_UNSPECIFIED            ProjectState = 0
	ProjectState_PROJECT_STATE_PENDING                ProjectState = 1
	ProjectState_PROJECT_STATE_FUNDED                 ProjectState = 2
	ProjectState_PROJECT_STATE_WAITING_FOR_VALIDATION ProjectState = 3
	ProjectState_PROJECT_STATE_PUBLISHED              ProjectState = 4
	ProjectState_PROJECT_STATE_CREATED                ProjectState = 5
	ProjectState_PROJECT_STATE_WAITING_FOR_APPROVAL   ProjectState = 6
	ProjectState_PROJECT_STATE_FAILED                 ProjectState = 7
)

// Enum value maps for ProjectState.
var (
	ProjectState_name = map[int32]string{
		0: "PROJECT_STATE_UNSPECIFIED",
		1: "PROJECT_STATE_PENDING",
		2: "PROJECT_STATE_FUNDED",
		3: "PROJECT_STATE_WAITING_FOR_VALIDATION",
		4: "PROJECT_STATE_PUBLISHED",
		5: "PROJECT_STATE_CREATED",
		6: "PROJECT_STATE_WAITING_FOR_APPROVAL",
		7: "PROJECT_STATE_FAILED",
	}
	ProjectState_value = map[string]int32{
		"PROJECT_STATE_UNSPECIFIED":            0,
		"PROJECT_STATE_PENDING":                1,
		"PROJECT_STATE_FUNDED":                 2,
		"PROJECT_STATE_WAITING_FOR_VALIDATION": 3,
		"PROJECT_STATE_PUBLISHED":              4,
		"PROJECT_STATE_CREATED":                5,
		"PROJECT_STATE_WAITING_FOR_APPROVAL":   6,
		"PROJECT_STATE_FAILED":                 7,
	}
)

func (x ProjectState) Enum() *ProjectState {
	p := new(ProjectState)
	*p = x
	return p
}

func (x ProjectState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectState) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_common_v1_project_state_proto_enumTypes[0].Descriptor()
}

func (ProjectState) Type() protoreflect.EnumType {
	return &file_protos_common_v1_project_state_proto_enumTypes[0]
}

func (x ProjectState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectState.Descriptor instead.
func (ProjectState) EnumDescriptor() ([]byte, []int) {
	return file_protos_common_v1_project_state_proto_rawDescGZIP(), []int{0}
}

var File_protos_common_v1_project_state_proto protoreflect.FileDescriptor

var file_protos_common_v1_project_state_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2a, 0x86, 0x02, 0x0a, 0x0c, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x52, 0x4f,
	0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x4a,
	0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x55, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x28, 0x0a,
	0x24, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x57,
	0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52, 0x4f, 0x4a, 0x45,
	0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48,
	0x45, 0x44, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12,
	0x26, 0x0a, 0x22, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41, 0x50, 0x50,
	0x52, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x4a, 0x45,
	0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x07, 0x42, 0x33, 0x5a, 0x31, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_common_v1_project_state_proto_rawDescOnce sync.Once
	file_protos_common_v1_project_state_proto_rawDescData = file_protos_common_v1_project_state_proto_rawDesc
)

func file_protos_common_v1_project_state_proto_rawDescGZIP() []byte {
	file_protos_common_v1_project_state_proto_rawDescOnce.Do(func() {
		file_protos_common_v1_project_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_common_v1_project_state_proto_rawDescData)
	})
	return file_protos_common_v1_project_state_proto_rawDescData
}

var file_protos_common_v1_project_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_common_v1_project_state_proto_goTypes = []interface{}{
	(ProjectState)(0), // 0: protos.common.v1.ProjectState
}
var file_protos_common_v1_project_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_common_v1_project_state_proto_init() }
func file_protos_common_v1_project_state_proto_init() {
	if File_protos_common_v1_project_state_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_common_v1_project_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_common_v1_project_state_proto_goTypes,
		DependencyIndexes: file_protos_common_v1_project_state_proto_depIdxs,
		EnumInfos:         file_protos_common_v1_project_state_proto_enumTypes,
	}.Build()
	File_protos_common_v1_project_state_proto = out.File
	file_protos_common_v1_project_state_proto_rawDesc = nil
	file_protos_common_v1_project_state_proto_goTypes = nil
	file_protos_common_v1_project_state_proto_depIdxs = nil
}
