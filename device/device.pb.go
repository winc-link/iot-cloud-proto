// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: device/device.proto

package device

import (
	common "github.com/winc-link/iot-cloud-proto/common"
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

type ConnectIotCloudRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string            `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	DriverId string            `protobuf:"bytes,2,opt,name=driverId,proto3" json:"driverId,omitempty"`
	Extra    map[string]string `protobuf:"bytes,3,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConnectIotCloudRequest) Reset() {
	*x = ConnectIotCloudRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectIotCloudRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectIotCloudRequest) ProtoMessage() {}

func (x *ConnectIotCloudRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectIotCloudRequest.ProtoReflect.Descriptor instead.
func (*ConnectIotCloudRequest) Descriptor() ([]byte, []int) {
	return file_device_device_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectIotCloudRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ConnectIotCloudRequest) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *ConnectIotCloudRequest) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type QueryDeviceDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
}

func (x *QueryDeviceDetailRequest) Reset() {
	*x = QueryDeviceDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDeviceDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDeviceDetailRequest) ProtoMessage() {}

func (x *QueryDeviceDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDeviceDetailRequest.ProtoReflect.Descriptor instead.
func (*QueryDeviceDetailRequest) Descriptor() ([]byte, []int) {
	return file_device_device_proto_rawDescGZIP(), []int{1}
}

func (x *QueryDeviceDetailRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type QueryDeviceDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data         *QueryDeviceDetailResponse_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	RequestId    string                          `protobuf:"bytes,2,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	Code         string                          `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Success      bool                            `protobuf:"varint,4,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrorMessage string                          `protobuf:"bytes,5,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *QueryDeviceDetailResponse) Reset() {
	*x = QueryDeviceDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDeviceDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDeviceDetailResponse) ProtoMessage() {}

func (x *QueryDeviceDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_device_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDeviceDetailResponse.ProtoReflect.Descriptor instead.
func (*QueryDeviceDetailResponse) Descriptor() ([]byte, []int) {
	return file_device_device_proto_rawDescGZIP(), []int{2}
}

func (x *QueryDeviceDetailResponse) GetData() *QueryDeviceDetailResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryDeviceDetailResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *QueryDeviceDetailResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *QueryDeviceDetailResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *QueryDeviceDetailResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type GetDeviceStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
}

func (x *GetDeviceStatusRequest) Reset() {
	*x = GetDeviceStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceStatusRequest) ProtoMessage() {}

func (x *GetDeviceStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceStatusRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceStatusRequest) Descriptor() ([]byte, []int) {
	return file_device_device_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeviceStatusRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type GetDeviceStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data         *GetDeviceStatusResponse_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	RequestId    string                        `protobuf:"bytes,2,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	Code         string                        `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Success      bool                          `protobuf:"varint,4,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrorMessage string                        `protobuf:"bytes,5,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *GetDeviceStatusResponse) Reset() {
	*x = GetDeviceStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceStatusResponse) ProtoMessage() {}

func (x *GetDeviceStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_device_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceStatusResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceStatusResponse) Descriptor() ([]byte, []int) {
	return file_device_device_proto_rawDescGZIP(), []int{4}
}

func (x *GetDeviceStatusResponse) GetData() *GetDeviceStatusResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetDeviceStatusResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *GetDeviceStatusResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetDeviceStatusResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetDeviceStatusResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type RegisterDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceName string            `protobuf:"bytes,1,opt,name=DeviceName,proto3" json:"DeviceName,omitempty"`
	ProductKey string            `protobuf:"bytes,2,opt,name=ProductKey,proto3" json:"ProductKey,omitempty"`
	Extra      map[string]string `protobuf:"bytes,3,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RegisterDeviceRequest) Reset() {
	*x = RegisterDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDeviceRequest) ProtoMessage() {}

func (x *RegisterDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_device_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDeviceRequest.ProtoReflect.Descriptor instead.
func (*RegisterDeviceRequest) Descriptor() ([]byte, []int) {
	return file_device_device_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterDeviceRequest) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *RegisterDeviceRequest) GetProductKey() string {
	if x != nil {
		return x.ProductKey
	}
	return ""
}

func (x *RegisterDeviceRequest) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type RegisterDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId     string            `protobuf:"bytes,1,opt,name=DeviceId,proto3" json:"DeviceId,omitempty"`
	DeviceSecret string            `protobuf:"bytes,2,opt,name=DeviceSecret,proto3" json:"DeviceSecret,omitempty"`
	Extra        map[string]string `protobuf:"bytes,3,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RegisterDeviceResponse) Reset() {
	*x = RegisterDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDeviceResponse) ProtoMessage() {}

func (x *RegisterDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_device_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDeviceResponse.ProtoReflect.Descriptor instead.
func (*RegisterDeviceResponse) Descriptor() ([]byte, []int) {
	return file_device_device_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterDeviceResponse) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *RegisterDeviceResponse) GetDeviceSecret() string {
	if x != nil {
		return x.DeviceSecret
	}
	return ""
}

func (x *RegisterDeviceResponse) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

type QueryDeviceDetailResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      string                 `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	ProductName string                 `protobuf:"bytes,3,opt,name=ProductName,proto3" json:"ProductName,omitempty"`
	UtcOnline   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=UtcOnline,proto3" json:"UtcOnline,omitempty"`
	IotId       string                 `protobuf:"bytes,6,opt,name=IotId,proto3" json:"IotId,omitempty"`
	GmtCreate   string                 `protobuf:"bytes,7,opt,name=GmtCreate,proto3" json:"GmtCreate,omitempty"`
	UtcCreate   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=UtcCreate,proto3" json:"UtcCreate,omitempty"`
	UtcActive   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=UtcActive,proto3" json:"UtcActive,omitempty"`
	GmtActive   string                 `protobuf:"bytes,10,opt,name=GmtActive,proto3" json:"GmtActive,omitempty"`
	NodeType    uint32                 `protobuf:"varint,11,opt,name=NodeType,proto3" json:"NodeType,omitempty"`
	Region      string                 `protobuf:"bytes,12,opt,name=Region,proto3" json:"Region,omitempty"`
	IpAddress   string                 `protobuf:"bytes,13,opt,name=IpAddress,proto3" json:"IpAddress,omitempty"`
	GmtOnline   string                 `protobuf:"bytes,14,opt,name=GmtOnline,proto3" json:"GmtOnline,omitempty"`
	ProductKey  string                 `protobuf:"bytes,15,opt,name=ProductKey,proto3" json:"ProductKey,omitempty"`
	DeviceName  string                 `protobuf:"bytes,16,opt,name=DeviceName,proto3" json:"DeviceName,omitempty"`
	Nickname    string                 `protobuf:"bytes,17,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
}

func (x *QueryDeviceDetailResponse_Data) Reset() {
	*x = QueryDeviceDetailResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_device_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDeviceDetailResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDeviceDetailResponse_Data) ProtoMessage() {}

func (x *QueryDeviceDetailResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_device_device_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDeviceDetailResponse_Data.ProtoReflect.Descriptor instead.
func (*QueryDeviceDetailResponse_Data) Descriptor() ([]byte, []int) {
	return file_device_device_proto_rawDescGZIP(), []int{2, 0}
}

func (x *QueryDeviceDetailResponse_Data) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *QueryDeviceDetailResponse_Data) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *QueryDeviceDetailResponse_Data) GetUtcOnline() *timestamppb.Timestamp {
	if x != nil {
		return x.UtcOnline
	}
	return nil
}

func (x *QueryDeviceDetailResponse_Data) GetIotId() string {
	if x != nil {
		return x.IotId
	}
	return ""
}

func (x *QueryDeviceDetailResponse_Data) GetGmtCreate() string {
	if x != nil {
		return x.GmtCreate
	}
	return ""
}

func (x *QueryDeviceDetailResponse_Data) GetUtcCreate() *timestamppb.Timestamp {
	if x != nil {
		return x.UtcCreate
	}
	return nil
}

func (x *QueryDeviceDetailResponse_Data) GetUtcActive() *timestamppb.Timestamp {
	if x != nil {
		return x.UtcActive
	}
	return nil
}

func (x *QueryDeviceDetailResponse_Data) GetGmtActive() string {
	if x != nil {
		return x.GmtActive
	}
	return ""
}

func (x *QueryDeviceDetailResponse_Data) GetNodeType() uint32 {
	if x != nil {
		return x.NodeType
	}
	return 0
}

func (x *QueryDeviceDetailResponse_Data) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *QueryDeviceDetailResponse_Data) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *QueryDeviceDetailResponse_Data) GetGmtOnline() string {
	if x != nil {
		return x.GmtOnline
	}
	return ""
}

func (x *QueryDeviceDetailResponse_Data) GetProductKey() string {
	if x != nil {
		return x.ProductKey
	}
	return ""
}

func (x *QueryDeviceDetailResponse_Data) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *QueryDeviceDetailResponse_Data) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type GetDeviceStatusResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Timestamp uint64 `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *GetDeviceStatusResponse_Data) Reset() {
	*x = GetDeviceStatusResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_device_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceStatusResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceStatusResponse_Data) ProtoMessage() {}

func (x *GetDeviceStatusResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_device_device_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceStatusResponse_Data.ProtoReflect.Descriptor instead.
func (*GetDeviceStatusResponse_Data) Descriptor() ([]byte, []int) {
	return file_device_device_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetDeviceStatusResponse_Data) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetDeviceStatusResponse_Data) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_device_device_proto protoreflect.FileDescriptor

var file_device_device_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49,
	0x6f, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x6f, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x36, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0xd6, 0x05, 0x0a, 0x19, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x8c, 0x04, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x74, 0x63, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x74, 0x63, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x49, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x49, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x6d, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x6d, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x74, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x74, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x55, 0x74, 0x63, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55,
	0x74, 0x63, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x6d, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x6d, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x6d, 0x74, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x6d, 0x74,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4b, 0x65, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x34, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x81, 0x02, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c,
	0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3c,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xd1, 0x01, 0x0a,
	0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x3e, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xd3, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x05, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38, 0x0a, 0x0a,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xdd, 0x02, 0x0a, 0x09, 0x52, 0x70, 0x63, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49,
	0x6f, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x1e, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x6f, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5a, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6e, 0x63, 0x2d, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69,
	0x6f, 0x74, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_device_device_proto_rawDescOnce sync.Once
	file_device_device_proto_rawDescData = file_device_device_proto_rawDesc
)

func file_device_device_proto_rawDescGZIP() []byte {
	file_device_device_proto_rawDescOnce.Do(func() {
		file_device_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_device_proto_rawDescData)
	})
	return file_device_device_proto_rawDescData
}

var file_device_device_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_device_device_proto_goTypes = []interface{}{
	(*ConnectIotCloudRequest)(nil),         // 0: device.ConnectIotCloudRequest
	(*QueryDeviceDetailRequest)(nil),       // 1: device.QueryDeviceDetailRequest
	(*QueryDeviceDetailResponse)(nil),      // 2: device.QueryDeviceDetailResponse
	(*GetDeviceStatusRequest)(nil),         // 3: device.GetDeviceStatusRequest
	(*GetDeviceStatusResponse)(nil),        // 4: device.GetDeviceStatusResponse
	(*RegisterDeviceRequest)(nil),          // 5: device.RegisterDeviceRequest
	(*RegisterDeviceResponse)(nil),         // 6: device.RegisterDeviceResponse
	nil,                                    // 7: device.ConnectIotCloudRequest.ExtraEntry
	(*QueryDeviceDetailResponse_Data)(nil), // 8: device.QueryDeviceDetailResponse.Data
	(*GetDeviceStatusResponse_Data)(nil),   // 9: device.GetDeviceStatusResponse.Data
	nil,                                    // 10: device.RegisterDeviceRequest.ExtraEntry
	nil,                                    // 11: device.RegisterDeviceResponse.ExtraEntry
	(*timestamppb.Timestamp)(nil),          // 12: google.protobuf.Timestamp
	(*common.CommonResponse)(nil),          // 13: common.CommonResponse
}
var file_device_device_proto_depIdxs = []int32{
	7,  // 0: device.ConnectIotCloudRequest.extra:type_name -> device.ConnectIotCloudRequest.ExtraEntry
	8,  // 1: device.QueryDeviceDetailResponse.data:type_name -> device.QueryDeviceDetailResponse.Data
	9,  // 2: device.GetDeviceStatusResponse.data:type_name -> device.GetDeviceStatusResponse.Data
	10, // 3: device.RegisterDeviceRequest.extra:type_name -> device.RegisterDeviceRequest.ExtraEntry
	11, // 4: device.RegisterDeviceResponse.extra:type_name -> device.RegisterDeviceResponse.ExtraEntry
	12, // 5: device.QueryDeviceDetailResponse.Data.UtcOnline:type_name -> google.protobuf.Timestamp
	12, // 6: device.QueryDeviceDetailResponse.Data.UtcCreate:type_name -> google.protobuf.Timestamp
	12, // 7: device.QueryDeviceDetailResponse.Data.UtcActive:type_name -> google.protobuf.Timestamp
	0,  // 8: device.RpcDevice.ConnectIotCloud:input_type -> device.ConnectIotCloudRequest
	1,  // 9: device.RpcDevice.QueryDeviceDetail:input_type -> device.QueryDeviceDetailRequest
	3,  // 10: device.RpcDevice.GetDeviceStatus:input_type -> device.GetDeviceStatusRequest
	5,  // 11: device.RpcDevice.RegisterDevice:input_type -> device.RegisterDeviceRequest
	13, // 12: device.RpcDevice.ConnectIotCloud:output_type -> common.CommonResponse
	2,  // 13: device.RpcDevice.QueryDeviceDetail:output_type -> device.QueryDeviceDetailResponse
	4,  // 14: device.RpcDevice.GetDeviceStatus:output_type -> device.GetDeviceStatusResponse
	6,  // 15: device.RpcDevice.RegisterDevice:output_type -> device.RegisterDeviceResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_device_device_proto_init() }
func file_device_device_proto_init() {
	if File_device_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectIotCloudRequest); i {
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
		file_device_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDeviceDetailRequest); i {
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
		file_device_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDeviceDetailResponse); i {
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
		file_device_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceStatusRequest); i {
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
		file_device_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceStatusResponse); i {
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
		file_device_device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDeviceRequest); i {
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
		file_device_device_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDeviceResponse); i {
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
		file_device_device_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDeviceDetailResponse_Data); i {
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
		file_device_device_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceStatusResponse_Data); i {
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
			RawDescriptor: file_device_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_device_device_proto_goTypes,
		DependencyIndexes: file_device_device_proto_depIdxs,
		MessageInfos:      file_device_device_proto_msgTypes,
	}.Build()
	File_device_device_proto = out.File
	file_device_device_proto_rawDesc = nil
	file_device_device_proto_goTypes = nil
	file_device_device_proto_depIdxs = nil
}
