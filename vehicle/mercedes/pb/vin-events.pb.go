// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: vin-events.proto

package protos

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

type VINUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddedVINs   []string `protobuf:"bytes,1,rep,name=addedVINs,proto3" json:"addedVINs,omitempty"`
	DeletedVINs []string `protobuf:"bytes,2,rep,name=deletedVINs,proto3" json:"deletedVINs,omitempty"`
}

func (x *VINUpdate) Reset() {
	*x = VINUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vin_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VINUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VINUpdate) ProtoMessage() {}

func (x *VINUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_vin_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VINUpdate.ProtoReflect.Descriptor instead.
func (*VINUpdate) Descriptor() ([]byte, []int) {
	return file_vin_events_proto_rawDescGZIP(), []int{0}
}

func (x *VINUpdate) GetAddedVINs() []string {
	if x != nil {
		return x.AddedVINs
	}
	return nil
}

func (x *VINUpdate) GetDeletedVINs() []string {
	if x != nil {
		return x.DeletedVINs
	}
	return nil
}

var File_vin_events_proto protoreflect.FileDescriptor

var file_vin_events_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x69, 0x6e, 0x2d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x09, 0x56, 0x49, 0x4e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x65, 0x64, 0x56,
	0x49, 0x4e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x56, 0x49, 0x4e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x56,
	0x49, 0x4e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x56, 0x49, 0x4e, 0x73, 0x42, 0x1c, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61,
	0x69, 0x6d, 0x6c, 0x65, 0x72, 0x2e, 0x6d, 0x62, 0x63, 0x61, 0x72, 0x6b, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vin_events_proto_rawDescOnce sync.Once
	file_vin_events_proto_rawDescData = file_vin_events_proto_rawDesc
)

func file_vin_events_proto_rawDescGZIP() []byte {
	file_vin_events_proto_rawDescOnce.Do(func() {
		file_vin_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_vin_events_proto_rawDescData)
	})
	return file_vin_events_proto_rawDescData
}

var file_vin_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_vin_events_proto_goTypes = []interface{}{
	(*VINUpdate)(nil), // 0: proto.VINUpdate
}
var file_vin_events_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vin_events_proto_init() }
func file_vin_events_proto_init() {
	if File_vin_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vin_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VINUpdate); i {
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
			RawDescriptor: file_vin_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vin_events_proto_goTypes,
		DependencyIndexes: file_vin_events_proto_depIdxs,
		MessageInfos:      file_vin_events_proto_msgTypes,
	}.Build()
	File_vin_events_proto = out.File
	file_vin_events_proto_rawDesc = nil
	file_vin_events_proto_goTypes = nil
	file_vin_events_proto_depIdxs = nil
}
