// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.3
// source: goes/projection/service.proto

package projectionpb

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

// TriggerReq represents a request to trigger a projection. It contains two
// fields: Schedule, which is a string that specifies the schedule for the
// projection, and Reset_, which is a boolean that indicates whether the
// projection should be reset before being triggered.
type TriggerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schedule string `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Reset_   bool   `protobuf:"varint,2,opt,name=reset,proto3" json:"reset,omitempty"`
}

// Reset resets the TriggerReq to its zero value. It sets Schedule to empty
// string and Reset_ to false. This method is automatically called by protobuf
// library when a TriggerReq object is created.
func (x *TriggerReq) Reset() {
	*x = TriggerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goes_projection_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

// String returns a string representation of the TriggerReq message. It is
// generated by the protoimpl library.
func (x *TriggerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

// ProtoMessage is a method implemented by *TriggerReq that marks it as a
// protobuf message type. This method is used to identify a struct as a protobuf
// message and allows it to be encoded and decoded using the protobuf encoding
// format.
func (*TriggerReq) ProtoMessage() {}

// ProtoReflect returns the message's reflection interface. This method is used
// internally by the protocol buffer library and should not be called by users.
// [protoreflect.Message]
func (x *TriggerReq) ProtoReflect() protoreflect.Message {
	mi := &file_goes_projection_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerReq.ProtoReflect.Descriptor instead.
func (*TriggerReq) Descriptor() ([]byte, []int) {
	return file_goes_projection_service_proto_rawDescGZIP(), []int{0}
}

// GetSchedule returns the schedule string of a TriggerReq message. It returns
// an empty string if the TriggerReq is nil.
func (x *TriggerReq) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

// GetReset_ returns a boolean value indicating whether the reset flag is set in
// a TriggerReq message.
func (x *TriggerReq) GetReset_() bool {
	if x != nil {
		return x.Reset_
	}
	return false
}

// TriggerResp represents the response message of the Trigger RPC call in the
// ProjectionService. It contains a single field, Accepted, which is a boolean
// value indicating whether the projection was accepted or not.
type TriggerResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
}

// Reset resets the TriggerResp to its zero value. It sets the Accepted field to
// false.
func (x *TriggerResp) Reset() {
	*x = TriggerResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goes_projection_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

// String returns a string representation of the TriggerResp message. The
// returned string is in protobuf text format.
func (x *TriggerResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

// ProtoMessage is an interface implemented by both TriggerReq and TriggerResp.
func (*TriggerResp) ProtoMessage() {}

// ProtoReflect returns the message's reflection interface. This method is used
// internally by the protocol buffer library and should not be called directly
// by user code.
func (x *TriggerResp) ProtoReflect() protoreflect.Message {
	mi := &file_goes_projection_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerResp.ProtoReflect.Descriptor instead.
func (*TriggerResp) Descriptor() ([]byte, []int) {
	return file_goes_projection_service_proto_rawDescGZIP(), []int{1}
}

// GetAccepted returns a boolean value indicating whether the trigger request
// was accepted or not. It is a method of the TriggerResp struct.
func (x *TriggerResp) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

// File_goes_projection_service_proto defines the protocol buffer messages
// TriggerReq and TriggerResp, used for the ProjectionService's Trigger method.
// The TriggerReq message contains a schedule string and a reset boolean, while
// the TriggerResp message contains an accepted boolean.
var File_goes_projection_service_proto protoreflect.FileDescriptor

var file_goes_projection_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x6f, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x67, 0x6f, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x3e, 0x0a, 0x0a, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x22, 0x29, 0x0a, 0x0b, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x32, 0x59, 0x0a, 0x11, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x44, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x67, 0x6f,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x72, 0x6e, 0x69, 0x63, 0x65, 0x2f, 0x67,
	0x6f, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_goes_projection_service_proto_rawDescOnce sync.Once
	file_goes_projection_service_proto_rawDescData = file_goes_projection_service_proto_rawDesc
)

func file_goes_projection_service_proto_rawDescGZIP() []byte {
	file_goes_projection_service_proto_rawDescOnce.Do(func() {
		file_goes_projection_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_goes_projection_service_proto_rawDescData)
	})
	return file_goes_projection_service_proto_rawDescData
}

var file_goes_projection_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_goes_projection_service_proto_goTypes = []interface{}{
	(*TriggerReq)(nil),  // 0: goes.projection.TriggerReq
	(*TriggerResp)(nil), // 1: goes.projection.TriggerResp
}
var file_goes_projection_service_proto_depIdxs = []int32{
	0, // 0: goes.projection.ProjectionService.Trigger:input_type -> goes.projection.TriggerReq
	1, // 1: goes.projection.ProjectionService.Trigger:output_type -> goes.projection.TriggerResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goes_projection_service_proto_init() }
func file_goes_projection_service_proto_init() {
	if File_goes_projection_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goes_projection_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerReq); i {
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
		file_goes_projection_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerResp); i {
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
			RawDescriptor: file_goes_projection_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goes_projection_service_proto_goTypes,
		DependencyIndexes: file_goes_projection_service_proto_depIdxs,
		MessageInfos:      file_goes_projection_service_proto_msgTypes,
	}.Build()
	File_goes_projection_service_proto = out.File
	file_goes_projection_service_proto_rawDesc = nil
	file_goes_projection_service_proto_goTypes = nil
	file_goes_projection_service_proto_depIdxs = nil
}
