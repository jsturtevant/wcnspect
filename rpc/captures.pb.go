// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.0
// source: captures.proto

package rpc

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

type PacketType int32

const (
	PacketType_all  PacketType = 0
	PacketType_flow PacketType = 1
	PacketType_drop PacketType = 2
)

// Enum value maps for PacketType.
var (
	PacketType_name = map[int32]string{
		0: "all",
		1: "flow",
		2: "drop",
	}
	PacketType_value = map[string]int32{
		"all":  0,
		"flow": 1,
		"drop": 2,
	}
)

func (x PacketType) Enum() *PacketType {
	p := new(PacketType)
	*p = x
	return p
}

func (x PacketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PacketType) Descriptor() protoreflect.EnumDescriptor {
	return file_captures_proto_enumTypes[0].Descriptor()
}

func (PacketType) Type() protoreflect.EnumType {
	return &file_captures_proto_enumTypes[0]
}

func (x PacketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PacketType.Descriptor instead.
func (PacketType) EnumDescriptor() ([]byte, []int) {
	return file_captures_proto_rawDescGZIP(), []int{0}
}

// models
type Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ips       []string `protobuf:"bytes,2,rep,name=ips,proto3" json:"ips,omitempty"`
	Protocols []string `protobuf:"bytes,3,rep,name=protocols,proto3" json:"protocols,omitempty"`
	Ports     []string `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	Macs      []string `protobuf:"bytes,5,rep,name=macs,proto3" json:"macs,omitempty"`
}

func (x *Filters) Reset() {
	*x = Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captures_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filters) ProtoMessage() {}

func (x *Filters) ProtoReflect() protoreflect.Message {
	mi := &file_captures_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filters.ProtoReflect.Descriptor instead.
func (*Filters) Descriptor() ([]byte, []int) {
	return file_captures_proto_rawDescGZIP(), []int{0}
}

func (x *Filters) GetIps() []string {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *Filters) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *Filters) GetPorts() []string {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *Filters) GetMacs() []string {
	if x != nil {
		return x.Macs
	}
	return nil
}

type Modifiers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pods         []string   `protobuf:"bytes,1,rep,name=pods,proto3" json:"pods,omitempty"`
	PacketType   PacketType `protobuf:"varint,4,opt,name=packet_type,json=packetType,proto3,enum=wcnspect.captures.PacketType" json:"packet_type,omitempty"`
	CountersOnly bool       `protobuf:"varint,5,opt,name=counters_only,json=countersOnly,proto3" json:"counters_only,omitempty"`
}

func (x *Modifiers) Reset() {
	*x = Modifiers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captures_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Modifiers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Modifiers) ProtoMessage() {}

func (x *Modifiers) ProtoReflect() protoreflect.Message {
	mi := &file_captures_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Modifiers.ProtoReflect.Descriptor instead.
func (*Modifiers) Descriptor() ([]byte, []int) {
	return file_captures_proto_rawDescGZIP(), []int{1}
}

func (x *Modifiers) GetPods() []string {
	if x != nil {
		return x.Pods
	}
	return nil
}

func (x *Modifiers) GetPacketType() PacketType {
	if x != nil {
		return x.PacketType
	}
	return PacketType_all
}

func (x *Modifiers) GetCountersOnly() bool {
	if x != nil {
		return x.CountersOnly
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captures_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_captures_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_captures_proto_rawDescGZIP(), []int{2}
}

// requests
type CaptureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration  int32                  `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Modifier  *Modifiers             `protobuf:"bytes,3,opt,name=modifier,proto3" json:"modifier,omitempty"`
	Filter    *Filters               `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *CaptureRequest) Reset() {
	*x = CaptureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captures_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptureRequest) ProtoMessage() {}

func (x *CaptureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captures_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptureRequest.ProtoReflect.Descriptor instead.
func (*CaptureRequest) Descriptor() ([]byte, []int) {
	return file_captures_proto_rawDescGZIP(), []int{3}
}

func (x *CaptureRequest) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *CaptureRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *CaptureRequest) GetModifier() *Modifiers {
	if x != nil {
		return x.Modifier
	}
	return nil
}

func (x *CaptureRequest) GetFilter() *Filters {
	if x != nil {
		return x.Filter
	}
	return nil
}

type CountersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeHidden bool `protobuf:"varint,1,opt,name=include_hidden,json=includeHidden,proto3" json:"include_hidden,omitempty"`
}

func (x *CountersRequest) Reset() {
	*x = CountersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captures_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountersRequest) ProtoMessage() {}

func (x *CountersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captures_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountersRequest.ProtoReflect.Descriptor instead.
func (*CountersRequest) Descriptor() ([]byte, []int) {
	return file_captures_proto_rawDescGZIP(), []int{4}
}

func (x *CountersRequest) GetIncludeHidden() bool {
	if x != nil {
		return x.IncludeHidden
	}
	return false
}

type VFPCountersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pod     string `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	Verbose bool   `protobuf:"varint,2,opt,name=verbose,proto3" json:"verbose,omitempty"`
}

func (x *VFPCountersRequest) Reset() {
	*x = VFPCountersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captures_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VFPCountersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VFPCountersRequest) ProtoMessage() {}

func (x *VFPCountersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_captures_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VFPCountersRequest.ProtoReflect.Descriptor instead.
func (*VFPCountersRequest) Descriptor() ([]byte, []int) {
	return file_captures_proto_rawDescGZIP(), []int{5}
}

func (x *VFPCountersRequest) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

func (x *VFPCountersRequest) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

// responses
type CaptureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    string                 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *CaptureResponse) Reset() {
	*x = CaptureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captures_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptureResponse) ProtoMessage() {}

func (x *CaptureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captures_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptureResponse.ProtoReflect.Descriptor instead.
func (*CaptureResponse) Descriptor() ([]byte, []int) {
	return file_captures_proto_rawDescGZIP(), []int{6}
}

func (x *CaptureResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *CaptureResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type StopCaptureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    string                 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *StopCaptureResponse) Reset() {
	*x = StopCaptureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captures_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopCaptureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopCaptureResponse) ProtoMessage() {}

func (x *StopCaptureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captures_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopCaptureResponse.ProtoReflect.Descriptor instead.
func (*StopCaptureResponse) Descriptor() ([]byte, []int) {
	return file_captures_proto_rawDescGZIP(), []int{7}
}

func (x *StopCaptureResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *StopCaptureResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type CountersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    string                 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *CountersResponse) Reset() {
	*x = CountersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captures_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountersResponse) ProtoMessage() {}

func (x *CountersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captures_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountersResponse.ProtoReflect.Descriptor instead.
func (*CountersResponse) Descriptor() ([]byte, []int) {
	return file_captures_proto_rawDescGZIP(), []int{8}
}

func (x *CountersResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *CountersResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type VFPCountersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    string                 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *VFPCountersResponse) Reset() {
	*x = VFPCountersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_captures_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VFPCountersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VFPCountersResponse) ProtoMessage() {}

func (x *VFPCountersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_captures_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VFPCountersResponse.ProtoReflect.Descriptor instead.
func (*VFPCountersResponse) Descriptor() ([]byte, []int) {
	return file_captures_proto_rawDescGZIP(), []int{9}
}

func (x *VFPCountersResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *VFPCountersResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_captures_proto protoreflect.FileDescriptor

var file_captures_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x77, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x70,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x63, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x63, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x09, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x64, 0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x77, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x63, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x4f, 0x6e, 0x6c, 0x79,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xd4, 0x01, 0x0a, 0x0e, 0x43, 0x61,
	0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e,
	0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x77,
	0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x38, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x68,
	0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x22, 0x40, 0x0a, 0x12, 0x56, 0x46,
	0x50, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70,
	0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x22, 0x63, 0x0a, 0x0f,
	0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x67, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x70, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x64, 0x0a, 0x10, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x67, 0x0a, 0x13, 0x56, 0x46, 0x50, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a, 0x29, 0x0a, 0x0a, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x66, 0x6c, 0x6f, 0x77, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x64, 0x72,
	0x6f, 0x70, 0x10, 0x02, 0x32, 0xfb, 0x02, 0x0a, 0x0e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x12, 0x21, 0x2e, 0x77, 0x69, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x69, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x51, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x18, 0x2e, 0x77, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x63, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26, 0x2e, 0x77, 0x69,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e,
	0x53, 0x74, 0x6f, 0x70, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x22, 0x2e, 0x77, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e,
	0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77, 0x69, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x61, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x46, 0x50, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x25, 0x2e, 0x77, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2e, 0x63, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x56, 0x46, 0x50, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x77, 0x69, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x2e, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x56, 0x46, 0x50,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74, 0x2f, 0x77, 0x69, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_captures_proto_rawDescOnce sync.Once
	file_captures_proto_rawDescData = file_captures_proto_rawDesc
)

func file_captures_proto_rawDescGZIP() []byte {
	file_captures_proto_rawDescOnce.Do(func() {
		file_captures_proto_rawDescData = protoimpl.X.CompressGZIP(file_captures_proto_rawDescData)
	})
	return file_captures_proto_rawDescData
}

var file_captures_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_captures_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_captures_proto_goTypes = []interface{}{
	(PacketType)(0),               // 0: wcnspect.captures.PacketType
	(*Filters)(nil),               // 1: wcnspect.captures.Filters
	(*Modifiers)(nil),             // 2: wcnspect.captures.Modifiers
	(*Empty)(nil),                 // 3: wcnspect.captures.Empty
	(*CaptureRequest)(nil),        // 4: wcnspect.captures.CaptureRequest
	(*CountersRequest)(nil),       // 5: wcnspect.captures.CountersRequest
	(*VFPCountersRequest)(nil),    // 6: wcnspect.captures.VFPCountersRequest
	(*CaptureResponse)(nil),       // 7: wcnspect.captures.CaptureResponse
	(*StopCaptureResponse)(nil),   // 8: wcnspect.captures.StopCaptureResponse
	(*CountersResponse)(nil),      // 9: wcnspect.captures.CountersResponse
	(*VFPCountersResponse)(nil),   // 10: wcnspect.captures.VFPCountersResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_captures_proto_depIdxs = []int32{
	0,  // 0: wcnspect.captures.Modifiers.packet_type:type_name -> wcnspect.captures.PacketType
	11, // 1: wcnspect.captures.CaptureRequest.timestamp:type_name -> google.protobuf.Timestamp
	2,  // 2: wcnspect.captures.CaptureRequest.modifier:type_name -> wcnspect.captures.Modifiers
	1,  // 3: wcnspect.captures.CaptureRequest.filter:type_name -> wcnspect.captures.Filters
	11, // 4: wcnspect.captures.CaptureResponse.timestamp:type_name -> google.protobuf.Timestamp
	11, // 5: wcnspect.captures.StopCaptureResponse.timestamp:type_name -> google.protobuf.Timestamp
	11, // 6: wcnspect.captures.CountersResponse.timestamp:type_name -> google.protobuf.Timestamp
	11, // 7: wcnspect.captures.VFPCountersResponse.timestamp:type_name -> google.protobuf.Timestamp
	4,  // 8: wcnspect.captures.CaptureService.StartCapture:input_type -> wcnspect.captures.CaptureRequest
	3,  // 9: wcnspect.captures.CaptureService.StopCapture:input_type -> wcnspect.captures.Empty
	5,  // 10: wcnspect.captures.CaptureService.GetCounters:input_type -> wcnspect.captures.CountersRequest
	6,  // 11: wcnspect.captures.CaptureService.GetVFPCounters:input_type -> wcnspect.captures.VFPCountersRequest
	7,  // 12: wcnspect.captures.CaptureService.StartCapture:output_type -> wcnspect.captures.CaptureResponse
	8,  // 13: wcnspect.captures.CaptureService.StopCapture:output_type -> wcnspect.captures.StopCaptureResponse
	9,  // 14: wcnspect.captures.CaptureService.GetCounters:output_type -> wcnspect.captures.CountersResponse
	10, // 15: wcnspect.captures.CaptureService.GetVFPCounters:output_type -> wcnspect.captures.VFPCountersResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_captures_proto_init() }
func file_captures_proto_init() {
	if File_captures_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_captures_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filters); i {
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
		file_captures_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Modifiers); i {
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
		file_captures_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_captures_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptureRequest); i {
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
		file_captures_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountersRequest); i {
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
		file_captures_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VFPCountersRequest); i {
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
		file_captures_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptureResponse); i {
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
		file_captures_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopCaptureResponse); i {
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
		file_captures_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountersResponse); i {
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
		file_captures_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VFPCountersResponse); i {
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
			RawDescriptor: file_captures_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_captures_proto_goTypes,
		DependencyIndexes: file_captures_proto_depIdxs,
		EnumInfos:         file_captures_proto_enumTypes,
		MessageInfos:      file_captures_proto_msgTypes,
	}.Build()
	File_captures_proto = out.File
	file_captures_proto_rawDesc = nil
	file_captures_proto_goTypes = nil
	file_captures_proto_depIdxs = nil
}
