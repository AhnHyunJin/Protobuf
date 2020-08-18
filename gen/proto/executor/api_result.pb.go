// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: proto/executor/api_result.proto

package grpc_executor

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Col string   `protobuf:"bytes,1,opt,name=col,proto3" json:"col,omitempty"`
	Val *any.Any `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_executor_api_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_proto_executor_api_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_proto_executor_api_result_proto_rawDescGZIP(), []int{0}
}

func (x *Condition) GetCol() string {
	if x != nil {
		return x.Col
	}
	return ""
}

func (x *Condition) GetVal() *any.Any {
	if x != nil {
		return x.Val
	}
	return nil
}

type ApiResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page         int32               `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage      int32               `protobuf:"varint,2,opt,name=perPage,proto3" json:"perPage,omitempty"`
	CurrentCount int32               `protobuf:"varint,3,opt,name=currentCount,proto3" json:"currentCount,omitempty"`
	MatchCount   int32               `protobuf:"varint,4,opt,name=matchCount,proto3" json:"matchCount,omitempty"`
	TotalCount   int32               `protobuf:"varint,5,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	Data         map[string]*any.Any `protobuf:"bytes,6,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ApiResult) Reset() {
	*x = ApiResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_executor_api_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResult) ProtoMessage() {}

func (x *ApiResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_executor_api_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResult.ProtoReflect.Descriptor instead.
func (*ApiResult) Descriptor() ([]byte, []int) {
	return file_proto_executor_api_result_proto_rawDescGZIP(), []int{1}
}

func (x *ApiResult) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ApiResult) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ApiResult) GetCurrentCount() int32 {
	if x != nil {
		return x.CurrentCount
	}
	return 0
}

func (x *ApiResult) GetMatchCount() int32 {
	if x != nil {
		return x.MatchCount
	}
	return 0
}

func (x *ApiResult) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ApiResult) GetData() map[string]*any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type ApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationId int32                 `protobuf:"varint,1,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
	ApiId         int32                 `protobuf:"varint,2,opt,name=apiId,proto3" json:"apiId,omitempty"`
	Page          int32                 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PerPage       int32                 `protobuf:"varint,4,opt,name=perPage,proto3" json:"perPage,omitempty"`
	Cond          map[string]*Condition `protobuf:"bytes,5,rep,name=cond,proto3" json:"cond,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ApiRequest) Reset() {
	*x = ApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_executor_api_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiRequest) ProtoMessage() {}

func (x *ApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_executor_api_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiRequest.ProtoReflect.Descriptor instead.
func (*ApiRequest) Descriptor() ([]byte, []int) {
	return file_proto_executor_api_result_proto_rawDescGZIP(), []int{2}
}

func (x *ApiRequest) GetApplicationId() int32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *ApiRequest) GetApiId() int32 {
	if x != nil {
		return x.ApiId
	}
	return 0
}

func (x *ApiRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ApiRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ApiRequest) GetCond() map[string]*Condition {
	if x != nil {
		return x.Cond
	}
	return nil
}

var File_proto_executor_api_result_proto protoreflect.FileDescriptor

var file_proto_executor_api_result_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6f, 0x6c, 0x12, 0x26, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x76,
	0x61, 0x6c, 0x22, 0xa4, 0x02, 0x0a, 0x09, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x4d, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x82, 0x02, 0x0a, 0x0a, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61,
	0x70, 0x69, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x50,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x63, 0x6f, 0x6e, 0x64, 0x1a, 0x51, 0x0a, 0x09, 0x43,
	0x6f, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x59,
	0x0a, 0x10, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x70,
	0x69, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x3b, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_executor_api_result_proto_rawDescOnce sync.Once
	file_proto_executor_api_result_proto_rawDescData = file_proto_executor_api_result_proto_rawDesc
)

func file_proto_executor_api_result_proto_rawDescGZIP() []byte {
	file_proto_executor_api_result_proto_rawDescOnce.Do(func() {
		file_proto_executor_api_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_executor_api_result_proto_rawDescData)
	})
	return file_proto_executor_api_result_proto_rawDescData
}

var file_proto_executor_api_result_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_executor_api_result_proto_goTypes = []interface{}{
	(*Condition)(nil),  // 0: grpc_executor.condition
	(*ApiResult)(nil),  // 1: grpc_executor.ApiResult
	(*ApiRequest)(nil), // 2: grpc_executor.ApiRequest
	nil,                // 3: grpc_executor.ApiResult.DataEntry
	nil,                // 4: grpc_executor.ApiRequest.CondEntry
	(*any.Any)(nil),    // 5: google.protobuf.Any
}
var file_proto_executor_api_result_proto_depIdxs = []int32{
	5, // 0: grpc_executor.condition.val:type_name -> google.protobuf.Any
	3, // 1: grpc_executor.ApiResult.data:type_name -> grpc_executor.ApiResult.DataEntry
	4, // 2: grpc_executor.ApiRequest.cond:type_name -> grpc_executor.ApiRequest.CondEntry
	5, // 3: grpc_executor.ApiResult.DataEntry.value:type_name -> google.protobuf.Any
	0, // 4: grpc_executor.ApiRequest.CondEntry.value:type_name -> grpc_executor.condition
	2, // 5: grpc_executor.ApiResultService.getApiResult:input_type -> grpc_executor.ApiRequest
	1, // 6: grpc_executor.ApiResultService.getApiResult:output_type -> grpc_executor.ApiResult
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_executor_api_result_proto_init() }
func file_proto_executor_api_result_proto_init() {
	if File_proto_executor_api_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_executor_api_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
		file_proto_executor_api_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResult); i {
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
		file_proto_executor_api_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiRequest); i {
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
			RawDescriptor: file_proto_executor_api_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_executor_api_result_proto_goTypes,
		DependencyIndexes: file_proto_executor_api_result_proto_depIdxs,
		MessageInfos:      file_proto_executor_api_result_proto_msgTypes,
	}.Build()
	File_proto_executor_api_result_proto = out.File
	file_proto_executor_api_result_proto_rawDesc = nil
	file_proto_executor_api_result_proto_goTypes = nil
	file_proto_executor_api_result_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApiResultServiceClient is the client API for ApiResultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiResultServiceClient interface {
	GetApiResult(ctx context.Context, in *ApiRequest, opts ...grpc.CallOption) (*ApiResult, error)
}

type apiResultServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiResultServiceClient(cc grpc.ClientConnInterface) ApiResultServiceClient {
	return &apiResultServiceClient{cc}
}

func (c *apiResultServiceClient) GetApiResult(ctx context.Context, in *ApiRequest, opts ...grpc.CallOption) (*ApiResult, error) {
	out := new(ApiResult)
	err := c.cc.Invoke(ctx, "/grpc_executor.ApiResultService/getApiResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiResultServiceServer is the server API for ApiResultService service.
type ApiResultServiceServer interface {
	GetApiResult(context.Context, *ApiRequest) (*ApiResult, error)
}

// UnimplementedApiResultServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApiResultServiceServer struct {
}

func (*UnimplementedApiResultServiceServer) GetApiResult(context.Context, *ApiRequest) (*ApiResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiResult not implemented")
}

func RegisterApiResultServiceServer(s *grpc.Server, srv ApiResultServiceServer) {
	s.RegisterService(&_ApiResultService_serviceDesc, srv)
}

func _ApiResultService_GetApiResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiResultServiceServer).GetApiResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_executor.ApiResultService/GetApiResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiResultServiceServer).GetApiResult(ctx, req.(*ApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiResultService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_executor.ApiResultService",
	HandlerType: (*ApiResultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getApiResult",
			Handler:    _ApiResultService_GetApiResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/executor/api_result.proto",
}
