// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: batch/batch.proto

package batch

import (
	common "github.com/winc-link/iot-cloud-proto/common"
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

type DeviceBatchReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	DeviceId  string `protobuf:"bytes,3,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	ProductId string `protobuf:"bytes,4,opt,name=productId,proto3" json:"productId,omitempty"`
}

func (x *DeviceBatchReportRequest) Reset() {
	*x = DeviceBatchReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_batch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceBatchReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceBatchReportRequest) ProtoMessage() {}

func (x *DeviceBatchReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_batch_batch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceBatchReportRequest.ProtoReflect.Descriptor instead.
func (*DeviceBatchReportRequest) Descriptor() ([]byte, []int) {
	return file_batch_batch_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceBatchReportRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *DeviceBatchReportRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *DeviceBatchReportRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

var File_batch_batch_proto protoreflect.FileDescriptor

var file_batch_batch_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x68, 0x0a, 0x18, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x32, 0x5d, 0x0a, 0x0b, 0x52, 0x70, 0x63,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x11, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x2e,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6e, 0x63, 0x2d, 0x6c, 0x69, 0x6e, 0x6b,
	0x2f, 0x69, 0x6f, 0x74, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_batch_batch_proto_rawDescOnce sync.Once
	file_batch_batch_proto_rawDescData = file_batch_batch_proto_rawDesc
)

func file_batch_batch_proto_rawDescGZIP() []byte {
	file_batch_batch_proto_rawDescOnce.Do(func() {
		file_batch_batch_proto_rawDescData = protoimpl.X.CompressGZIP(file_batch_batch_proto_rawDescData)
	})
	return file_batch_batch_proto_rawDescData
}

var file_batch_batch_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_batch_batch_proto_goTypes = []interface{}{
	(*DeviceBatchReportRequest)(nil), // 0: batch.DeviceBatchReportRequest
	(*common.CommonResponse)(nil),    // 1: common.CommonResponse
}
var file_batch_batch_proto_depIdxs = []int32{
	0, // 0: batch.RpcProperty.DeviceBatchReport:input_type -> batch.DeviceBatchReportRequest
	1, // 1: batch.RpcProperty.DeviceBatchReport:output_type -> common.CommonResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_batch_batch_proto_init() }
func file_batch_batch_proto_init() {
	if File_batch_batch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_batch_batch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceBatchReportRequest); i {
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
			RawDescriptor: file_batch_batch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_batch_batch_proto_goTypes,
		DependencyIndexes: file_batch_batch_proto_depIdxs,
		MessageInfos:      file_batch_batch_proto_msgTypes,
	}.Build()
	File_batch_batch_proto = out.File
	file_batch_batch_proto_rawDesc = nil
	file_batch_batch_proto_goTypes = nil
	file_batch_batch_proto_depIdxs = nil
}
