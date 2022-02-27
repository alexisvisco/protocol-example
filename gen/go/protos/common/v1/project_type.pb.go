// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/common/v1/project_type.proto

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

type ProjectType int32

const (
	ProjectType_PROJECT_TYPE_UNSPECIFIED ProjectType = 0 // All
	ProjectType_PROJECT_TYPE_DONATION    ProjectType = 1 // Project::Collect::Donation
	ProjectType_PROJECT_TYPE_LENDING     ProjectType = 2 // Project::Collect::Lending
)

// Enum value maps for ProjectType.
var (
	ProjectType_name = map[int32]string{
		0: "PROJECT_TYPE_UNSPECIFIED",
		1: "PROJECT_TYPE_DONATION",
		2: "PROJECT_TYPE_LENDING",
	}
	ProjectType_value = map[string]int32{
		"PROJECT_TYPE_UNSPECIFIED": 0,
		"PROJECT_TYPE_DONATION":    1,
		"PROJECT_TYPE_LENDING":     2,
	}
)

func (x ProjectType) Enum() *ProjectType {
	p := new(ProjectType)
	*p = x
	return p
}

func (x ProjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_common_v1_project_type_proto_enumTypes[0].Descriptor()
}

func (ProjectType) Type() protoreflect.EnumType {
	return &file_protos_common_v1_project_type_proto_enumTypes[0]
}

func (x ProjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectType.Descriptor instead.
func (ProjectType) EnumDescriptor() ([]byte, []int) {
	return file_protos_common_v1_project_type_proto_rawDescGZIP(), []int{0}
}

var File_protos_common_v1_project_type_proto protoreflect.FileDescriptor

var file_protos_common_v1_project_type_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2a, 0x60, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12,
	0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4c, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x42, 0x33, 0x5a, 0x31, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_common_v1_project_type_proto_rawDescOnce sync.Once
	file_protos_common_v1_project_type_proto_rawDescData = file_protos_common_v1_project_type_proto_rawDesc
)

func file_protos_common_v1_project_type_proto_rawDescGZIP() []byte {
	file_protos_common_v1_project_type_proto_rawDescOnce.Do(func() {
		file_protos_common_v1_project_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_common_v1_project_type_proto_rawDescData)
	})
	return file_protos_common_v1_project_type_proto_rawDescData
}

var file_protos_common_v1_project_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_common_v1_project_type_proto_goTypes = []interface{}{
	(ProjectType)(0), // 0: protos.common.v1.ProjectType
}
var file_protos_common_v1_project_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_common_v1_project_type_proto_init() }
func file_protos_common_v1_project_type_proto_init() {
	if File_protos_common_v1_project_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_common_v1_project_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_common_v1_project_type_proto_goTypes,
		DependencyIndexes: file_protos_common_v1_project_type_proto_depIdxs,
		EnumInfos:         file_protos_common_v1_project_type_proto_enumTypes,
	}.Build()
	File_protos_common_v1_project_type_proto = out.File
	file_protos_common_v1_project_type_proto_rawDesc = nil
	file_protos_common_v1_project_type_proto_goTypes = nil
	file_protos_common_v1_project_type_proto_depIdxs = nil
}