// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: property/property.proto

package property

import (
	common "github.com/winc-link/iot-cloud-proto/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DevicePropertyReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId string     `protobuf:"bytes,1,opt,name=DriverId,proto3" json:"DriverId,omitempty"`
	DeviceId string     `protobuf:"bytes,2,opt,name=DeviceId,proto3" json:"DeviceId,omitempty"`
	Data     *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DevicePropertyReportRequest) Reset() {
	*x = DevicePropertyReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevicePropertyReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevicePropertyReportRequest) ProtoMessage() {}

func (x *DevicePropertyReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevicePropertyReportRequest.ProtoReflect.Descriptor instead.
func (*DevicePropertyReportRequest) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{0}
}

func (x *DevicePropertyReportRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *DevicePropertyReportRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *DevicePropertyReportRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type DevicePropertySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string     `protobuf:"bytes,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	DeviceId  string     `protobuf:"bytes,2,opt,name=DeviceId,proto3" json:"DeviceId,omitempty"`
	Data      *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DevicePropertySetRequest) Reset() {
	*x = DevicePropertySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_property_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevicePropertySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevicePropertySetRequest) ProtoMessage() {}

func (x *DevicePropertySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_property_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevicePropertySetRequest.ProtoReflect.Descriptor instead.
func (*DevicePropertySetRequest) Descriptor() ([]byte, []int) {
	return file_property_property_proto_rawDescGZIP(), []int{1}
}

func (x *DevicePropertySetRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *DevicePropertySetRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *DevicePropertySetRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_property_property_proto protoreflect.FileDescriptor

var file_property_property_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x1b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x7e, 0x0a, 0x18, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0xb9, 0x01, 0x0a, 0x0b, 0x52, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x12, 0x57, 0x0a, 0x14, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x25, 0x2e, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a,
	0x11, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x53,
	0x65, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x69, 0x6e, 0x63, 0x2d, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6f, 0x74, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_property_property_proto_rawDescOnce sync.Once
	file_property_property_proto_rawDescData = file_property_property_proto_rawDesc
)

func file_property_property_proto_rawDescGZIP() []byte {
	file_property_property_proto_rawDescOnce.Do(func() {
		file_property_property_proto_rawDescData = protoimpl.X.CompressGZIP(file_property_property_proto_rawDescData)
	})
	return file_property_property_proto_rawDescData
}

var file_property_property_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_property_property_proto_goTypes = []interface{}{
	(*DevicePropertyReportRequest)(nil), // 0: property.DevicePropertyReportRequest
	(*DevicePropertySetRequest)(nil),    // 1: property.DevicePropertySetRequest
	(*anypb.Any)(nil),                   // 2: google.protobuf.Any
	(*common.CommonResponse)(nil),       // 3: common.CommonResponse
}
var file_property_property_proto_depIdxs = []int32{
	2, // 0: property.DevicePropertyReportRequest.data:type_name -> google.protobuf.Any
	2, // 1: property.DevicePropertySetRequest.data:type_name -> google.protobuf.Any
	0, // 2: property.RpcProperty.DevicePropertyReport:input_type -> property.DevicePropertyReportRequest
	1, // 3: property.RpcProperty.DevicePropertySet:input_type -> property.DevicePropertySetRequest
	3, // 4: property.RpcProperty.DevicePropertyReport:output_type -> common.CommonResponse
	3, // 5: property.RpcProperty.DevicePropertySet:output_type -> common.CommonResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_property_property_proto_init() }
func file_property_property_proto_init() {
	if File_property_property_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_property_property_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevicePropertyReportRequest); i {
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
		file_property_property_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevicePropertySetRequest); i {
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
			RawDescriptor: file_property_property_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_property_property_proto_goTypes,
		DependencyIndexes: file_property_property_proto_depIdxs,
		MessageInfos:      file_property_property_proto_msgTypes,
	}.Build()
	File_property_property_proto = out.File
	file_property_property_proto_rawDesc = nil
	file_property_property_proto_goTypes = nil
	file_property_property_proto_depIdxs = nil
}
