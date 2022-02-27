// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/common/v1/health.proto

package commonpb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type HealthStatus int32

const (
	HealthStatus_HEALTH_STATUS_UNSPECIFIED HealthStatus = 0
	// The application run normally and successfully
	HealthStatus_HEALTH_STATUS_SERVING HealthStatus = 1
	// An error is occurred recently and can produce more errors. One or more components in the stack might become overloaded soon.
	HealthStatus_HEALTH_STATUS_WARNING HealthStatus = 2
	// A partial anomaly of full outage of service. Non-security patches might be available.
	HealthStatus_HEALTH_STATUS_DEGRADED HealthStatus = 3
	// One or more components in the stack might be in an unusable status and the stack might become unresponsive soon. Security patches might be available.
	HealthStatus_HEALTH_STATUS_CRITICAL HealthStatus = 4
)

// Enum value maps for HealthStatus.
var (
	HealthStatus_name = map[int32]string{
		0: "HEALTH_STATUS_UNSPECIFIED",
		1: "HEALTH_STATUS_SERVING",
		2: "HEALTH_STATUS_WARNING",
		3: "HEALTH_STATUS_DEGRADED",
		4: "HEALTH_STATUS_CRITICAL",
	}
	HealthStatus_value = map[string]int32{
		"HEALTH_STATUS_UNSPECIFIED": 0,
		"HEALTH_STATUS_SERVING":     1,
		"HEALTH_STATUS_WARNING":     2,
		"HEALTH_STATUS_DEGRADED":    3,
		"HEALTH_STATUS_CRITICAL":    4,
	}
)

func (x HealthStatus) Enum() *HealthStatus {
	p := new(HealthStatus)
	*p = x
	return p
}

func (x HealthStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_common_v1_health_proto_enumTypes[0].Descriptor()
}

func (HealthStatus) Type() protoreflect.EnumType {
	return &file_protos_common_v1_health_proto_enumTypes[0]
}

func (x HealthStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthStatus.Descriptor instead.
func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return file_protos_common_v1_health_proto_rawDescGZIP(), []int{0}
}

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_v1_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_v1_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_protos_common_v1_health_proto_rawDescGZIP(), []int{0}
}

//
// HealthResponse Message
//
// This message is only used to respond to Health Service.
// This message return an enumerator of Status onf app
type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status represents the current status of the application
	Status HealthStatus `protobuf:"varint,1,opt,name=status,proto3,enum=protos.common.v1.HealthStatus" json:"status,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_v1_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_v1_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_protos_common_v1_health_proto_rawDescGZIP(), []int{1}
}

func (x *HealthResponse) GetStatus() HealthStatus {
	if x != nil {
		return x.Status
	}
	return HealthStatus_HEALTH_STATUS_UNSPECIFIED
}

var File_protos_common_v1_health_proto protoreflect.FileDescriptor

var file_protos_common_v1_health_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x69, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x1f, 0x92, 0x41,
	0x1c, 0x0a, 0x1a, 0x2a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x9b, 0x01,
	0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x19, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x48, 0x45, 0x41, 0x4c,
	0x54, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x47, 0x52, 0x41, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x1a, 0x0a, 0x16, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x04, 0x42, 0x33, 0x5a, 0x31, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_common_v1_health_proto_rawDescOnce sync.Once
	file_protos_common_v1_health_proto_rawDescData = file_protos_common_v1_health_proto_rawDesc
)

func file_protos_common_v1_health_proto_rawDescGZIP() []byte {
	file_protos_common_v1_health_proto_rawDescOnce.Do(func() {
		file_protos_common_v1_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_common_v1_health_proto_rawDescData)
	})
	return file_protos_common_v1_health_proto_rawDescData
}

var file_protos_common_v1_health_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_common_v1_health_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_common_v1_health_proto_goTypes = []interface{}{
	(HealthStatus)(0),      // 0: protos.common.v1.HealthStatus
	(*HealthRequest)(nil),  // 1: protos.common.v1.HealthRequest
	(*HealthResponse)(nil), // 2: protos.common.v1.HealthResponse
}
var file_protos_common_v1_health_proto_depIdxs = []int32{
	0, // 0: protos.common.v1.HealthResponse.status:type_name -> protos.common.v1.HealthStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_common_v1_health_proto_init() }
func file_protos_common_v1_health_proto_init() {
	if File_protos_common_v1_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_common_v1_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
		file_protos_common_v1_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
			RawDescriptor: file_protos_common_v1_health_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_common_v1_health_proto_goTypes,
		DependencyIndexes: file_protos_common_v1_health_proto_depIdxs,
		EnumInfos:         file_protos_common_v1_health_proto_enumTypes,
		MessageInfos:      file_protos_common_v1_health_proto_msgTypes,
	}.Build()
	File_protos_common_v1_health_proto = out.File
	file_protos_common_v1_health_proto_rawDesc = nil
	file_protos_common_v1_health_proto_goTypes = nil
	file_protos_common_v1_health_proto_depIdxs = nil
}
