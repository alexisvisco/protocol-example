// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protos/leads/v1/leads.proto

package leadspb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "protocol-example/gen/go/protos/common/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_leads_v1_leads_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_leads_v1_leads_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_protos_leads_v1_leads_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeadId   int64     `protobuf:"varint,1,opt,name=lead_id,json=leadId,proto3" json:"lead_id,omitempty"`
	Progress *Progress `protobuf:"bytes,2,opt,name=progress,proto3" json:"progress,omitempty"`
	Step     *Step     `protobuf:"bytes,3,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_leads_v1_leads_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_leads_v1_leads_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_protos_leads_v1_leads_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetLeadId() int64 {
	if x != nil {
		return x.LeadId
	}
	return 0
}

func (x *CreateResponse) GetProgress() *Progress {
	if x != nil {
		return x.Progress
	}
	return nil
}

func (x *CreateResponse) GetStep() *Step {
	if x != nil {
		return x.Step
	}
	return nil
}

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeadId int64 `protobuf:"varint,1,opt,name=lead_id,json=leadId,proto3" json:"lead_id,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_leads_v1_leads_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_leads_v1_leads_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_protos_leads_v1_leads_proto_rawDescGZIP(), []int{2}
}

func (x *FetchRequest) GetLeadId() int64 {
	if x != nil {
		return x.LeadId
	}
	return 0
}

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress *Progress `protobuf:"bytes,1,opt,name=progress,proto3" json:"progress,omitempty"`
	Step     *Step     `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_leads_v1_leads_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_leads_v1_leads_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_protos_leads_v1_leads_proto_rawDescGZIP(), []int{3}
}

func (x *FetchResponse) GetProgress() *Progress {
	if x != nil {
		return x.Progress
	}
	return nil
}

func (x *FetchResponse) GetStep() *Step {
	if x != nil {
		return x.Step
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeadId   int64         `protobuf:"varint,1,opt,name=lead_id,json=leadId,proto3" json:"lead_id,omitempty"`
	StepName string        `protobuf:"bytes,2,opt,name=step_name,json=stepName,proto3" json:"step_name,omitempty"`
	Answers  []*InputValue `protobuf:"bytes,3,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_leads_v1_leads_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_leads_v1_leads_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_protos_leads_v1_leads_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRequest) GetLeadId() int64 {
	if x != nil {
		return x.LeadId
	}
	return 0
}

func (x *UpdateRequest) GetStepName() string {
	if x != nil {
		return x.StepName
	}
	return ""
}

func (x *UpdateRequest) GetAnswers() []*InputValue {
	if x != nil {
		return x.Answers
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress *Progress `protobuf:"bytes,1,opt,name=progress,proto3" json:"progress,omitempty"`
	Step     *Step     `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_leads_v1_leads_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_leads_v1_leads_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_protos_leads_v1_leads_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateResponse) GetProgress() *Progress {
	if x != nil {
		return x.Progress
	}
	return nil
}

func (x *UpdateResponse) GetStep() *Step {
	if x != nil {
		return x.Step
	}
	return nil
}

type PreviousRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeadId int64 `protobuf:"varint,1,opt,name=lead_id,json=leadId,proto3" json:"lead_id,omitempty"`
}

func (x *PreviousRequest) Reset() {
	*x = PreviousRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_leads_v1_leads_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviousRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviousRequest) ProtoMessage() {}

func (x *PreviousRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_leads_v1_leads_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviousRequest.ProtoReflect.Descriptor instead.
func (*PreviousRequest) Descriptor() ([]byte, []int) {
	return file_protos_leads_v1_leads_proto_rawDescGZIP(), []int{6}
}

func (x *PreviousRequest) GetLeadId() int64 {
	if x != nil {
		return x.LeadId
	}
	return 0
}

type PreviousResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress *Progress `protobuf:"bytes,1,opt,name=progress,proto3" json:"progress,omitempty"`
	Step     *Step     `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *PreviousResponse) Reset() {
	*x = PreviousResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_leads_v1_leads_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviousResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviousResponse) ProtoMessage() {}

func (x *PreviousResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_leads_v1_leads_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviousResponse.ProtoReflect.Descriptor instead.
func (*PreviousResponse) Descriptor() ([]byte, []int) {
	return file_protos_leads_v1_leads_proto_rawDescGZIP(), []int{7}
}

func (x *PreviousResponse) GetProgress() *Progress {
	if x != nil {
		return x.Progress
	}
	return nil
}

func (x *PreviousResponse) GetStep() *Step {
	if x != nil {
		return x.Step
	}
	return nil
}

// returns the all needed for each step
// name is the step name
// list of inputs names for the front to build
// when filling inputs answers returns a map of inputs names with their values
type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PreviousName     string             `protobuf:"bytes,2,opt,name=previous_name,json=previousName,proto3" json:"previous_name,omitempty"`
	InputDescriptors []*InputDescriptor `protobuf:"bytes,3,rep,name=input_descriptors,json=inputDescriptors,proto3" json:"input_descriptors,omitempty"`
	Answers          []*InputValue      `protobuf:"bytes,5,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_leads_v1_leads_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_protos_leads_v1_leads_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_protos_leads_v1_leads_proto_rawDescGZIP(), []int{8}
}

func (x *Step) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Step) GetPreviousName() string {
	if x != nil {
		return x.PreviousName
	}
	return ""
}

func (x *Step) GetInputDescriptors() []*InputDescriptor {
	if x != nil {
		return x.InputDescriptors
	}
	return nil
}

func (x *Step) GetAnswers() []*InputValue {
	if x != nil {
		return x.Answers
	}
	return nil
}

// Represent the progression of a lead, primary usage is for the frontend
// **-- ---- ---- ---- = {step_count: 4, current_step: 1, sub_step_count: 4, current_sub_step: 2}
type Progress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StepCount      int32 `protobuf:"varint,1,opt,name=step_count,json=stepCount,proto3" json:"step_count,omitempty"`
	CurrentStep    int32 `protobuf:"varint,2,opt,name=current_step,json=currentStep,proto3" json:"current_step,omitempty"`
	SubStepCount   int32 `protobuf:"varint,3,opt,name=sub_step_count,json=subStepCount,proto3" json:"sub_step_count,omitempty"`
	CurrentSubStep int32 `protobuf:"varint,4,opt,name=current_sub_step,json=currentSubStep,proto3" json:"current_sub_step,omitempty"`
}

func (x *Progress) Reset() {
	*x = Progress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_leads_v1_leads_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Progress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Progress) ProtoMessage() {}

func (x *Progress) ProtoReflect() protoreflect.Message {
	mi := &file_protos_leads_v1_leads_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Progress.ProtoReflect.Descriptor instead.
func (*Progress) Descriptor() ([]byte, []int) {
	return file_protos_leads_v1_leads_proto_rawDescGZIP(), []int{9}
}

func (x *Progress) GetStepCount() int32 {
	if x != nil {
		return x.StepCount
	}
	return 0
}

func (x *Progress) GetCurrentStep() int32 {
	if x != nil {
		return x.CurrentStep
	}
	return 0
}

func (x *Progress) GetSubStepCount() int32 {
	if x != nil {
		return x.SubStepCount
	}
	return 0
}

func (x *Progress) GetCurrentSubStep() int32 {
	if x != nil {
		return x.CurrentSubStep
	}
	return 0
}

var File_protos_leads_v1_leads_proto protoreflect.FileDescriptor

var file_protos_leads_v1_leads_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x65, 0x70, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x49,
	0x64, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x22, 0x27, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x0d,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22,
	0x7c, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x65,
	0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74,
	0x65, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x72, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x22, 0x2a, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x49, 0x64, 0x22, 0x74, 0x0a,
	0x10, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x22, 0xc5, 0x01, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x11, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x52, 0x10, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x65, 0x70,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74,
	0x65, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75,
	0x62, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x53, 0x74, 0x65, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x28, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x5f,
	0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x53, 0x74, 0x65, 0x70, 0x32, 0x93, 0x04, 0x0a, 0x0c, 0x4c,
	0x65, 0x61, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x12, 0x66, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x66, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x32, 0x10, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x5f, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x12, 0x6b, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x12,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x42, 0x31, 0x5a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_leads_v1_leads_proto_rawDescOnce sync.Once
	file_protos_leads_v1_leads_proto_rawDescData = file_protos_leads_v1_leads_proto_rawDesc
)

func file_protos_leads_v1_leads_proto_rawDescGZIP() []byte {
	file_protos_leads_v1_leads_proto_rawDescOnce.Do(func() {
		file_protos_leads_v1_leads_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_leads_v1_leads_proto_rawDescData)
	})
	return file_protos_leads_v1_leads_proto_rawDescData
}

var file_protos_leads_v1_leads_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_leads_v1_leads_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),     // 0: protos.leads.v1.CreateRequest
	(*CreateResponse)(nil),    // 1: protos.leads.v1.CreateResponse
	(*FetchRequest)(nil),      // 2: protos.leads.v1.FetchRequest
	(*FetchResponse)(nil),     // 3: protos.leads.v1.FetchResponse
	(*UpdateRequest)(nil),     // 4: protos.leads.v1.UpdateRequest
	(*UpdateResponse)(nil),    // 5: protos.leads.v1.UpdateResponse
	(*PreviousRequest)(nil),   // 6: protos.leads.v1.PreviousRequest
	(*PreviousResponse)(nil),  // 7: protos.leads.v1.PreviousResponse
	(*Step)(nil),              // 8: protos.leads.v1.Step
	(*Progress)(nil),          // 9: protos.leads.v1.Progress
	(*InputValue)(nil),        // 10: protos.leads.v1.InputValue
	(*InputDescriptor)(nil),   // 11: protos.leads.v1.InputDescriptor
	(*v1.HealthRequest)(nil),  // 12: protos.common.v1.HealthRequest
	(*v1.HealthResponse)(nil), // 13: protos.common.v1.HealthResponse
}
var file_protos_leads_v1_leads_proto_depIdxs = []int32{
	9,  // 0: protos.leads.v1.CreateResponse.progress:type_name -> protos.leads.v1.Progress
	8,  // 1: protos.leads.v1.CreateResponse.step:type_name -> protos.leads.v1.Step
	9,  // 2: protos.leads.v1.FetchResponse.progress:type_name -> protos.leads.v1.Progress
	8,  // 3: protos.leads.v1.FetchResponse.step:type_name -> protos.leads.v1.Step
	10, // 4: protos.leads.v1.UpdateRequest.answers:type_name -> protos.leads.v1.InputValue
	9,  // 5: protos.leads.v1.UpdateResponse.progress:type_name -> protos.leads.v1.Progress
	8,  // 6: protos.leads.v1.UpdateResponse.step:type_name -> protos.leads.v1.Step
	9,  // 7: protos.leads.v1.PreviousResponse.progress:type_name -> protos.leads.v1.Progress
	8,  // 8: protos.leads.v1.PreviousResponse.step:type_name -> protos.leads.v1.Step
	11, // 9: protos.leads.v1.Step.input_descriptors:type_name -> protos.leads.v1.InputDescriptor
	10, // 10: protos.leads.v1.Step.answers:type_name -> protos.leads.v1.InputValue
	12, // 11: protos.leads.v1.LeadsService.Health:input_type -> protos.common.v1.HealthRequest
	0,  // 12: protos.leads.v1.LeadsService.Create:input_type -> protos.leads.v1.CreateRequest
	4,  // 13: protos.leads.v1.LeadsService.Update:input_type -> protos.leads.v1.UpdateRequest
	2,  // 14: protos.leads.v1.LeadsService.Fetch:input_type -> protos.leads.v1.FetchRequest
	6,  // 15: protos.leads.v1.LeadsService.Previous:input_type -> protos.leads.v1.PreviousRequest
	13, // 16: protos.leads.v1.LeadsService.Health:output_type -> protos.common.v1.HealthResponse
	1,  // 17: protos.leads.v1.LeadsService.Create:output_type -> protos.leads.v1.CreateResponse
	5,  // 18: protos.leads.v1.LeadsService.Update:output_type -> protos.leads.v1.UpdateResponse
	3,  // 19: protos.leads.v1.LeadsService.Fetch:output_type -> protos.leads.v1.FetchResponse
	7,  // 20: protos.leads.v1.LeadsService.Previous:output_type -> protos.leads.v1.PreviousResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_protos_leads_v1_leads_proto_init() }
func file_protos_leads_v1_leads_proto_init() {
	if File_protos_leads_v1_leads_proto != nil {
		return
	}
	file_protos_leads_v1_steps_names_proto_init()
	file_protos_leads_v1_input_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_leads_v1_leads_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_protos_leads_v1_leads_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_protos_leads_v1_leads_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
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
		file_protos_leads_v1_leads_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
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
		file_protos_leads_v1_leads_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_protos_leads_v1_leads_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_protos_leads_v1_leads_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviousRequest); i {
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
		file_protos_leads_v1_leads_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreviousResponse); i {
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
		file_protos_leads_v1_leads_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
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
		file_protos_leads_v1_leads_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Progress); i {
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
			RawDescriptor: file_protos_leads_v1_leads_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_leads_v1_leads_proto_goTypes,
		DependencyIndexes: file_protos_leads_v1_leads_proto_depIdxs,
		MessageInfos:      file_protos_leads_v1_leads_proto_msgTypes,
	}.Build()
	File_protos_leads_v1_leads_proto = out.File
	file_protos_leads_v1_leads_proto_rawDesc = nil
	file_protos_leads_v1_leads_proto_goTypes = nil
	file_protos_leads_v1_leads_proto_depIdxs = nil
}
