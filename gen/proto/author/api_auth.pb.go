// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: proto/author/api_auth.proto

package grpc_author

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type ApiAuthRes_Code int32

const (
	ApiAuthRes_VALID                ApiAuthRes_Code = 0
	ApiAuthRes_INTERNAL_EXCEPTION   ApiAuthRes_Code = -1
	ApiAuthRes_PARAMETER_EXCEPTION  ApiAuthRes_Code = -2
	ApiAuthRes_UNREGISTERED_SERVICE ApiAuthRes_Code = -3
	ApiAuthRes_UNREGISTERED_TOKEN   ApiAuthRes_Code = -4
	ApiAuthRes_TERMINATED_SERVICE   ApiAuthRes_Code = -9
	ApiAuthRes_LIMIT_EXCEEDED       ApiAuthRes_Code = -10
	ApiAuthRes_UNAUTHORIZED         ApiAuthRes_Code = -401
	ApiAuthRes_UNKNOWN              ApiAuthRes_Code = -999
)

// Enum value maps for ApiAuthRes_Code.
var (
	ApiAuthRes_Code_name = map[int32]string{
		0:    "VALID",
		-1:   "INTERNAL_EXCEPTION",
		-2:   "PARAMETER_EXCEPTION",
		-3:   "UNREGISTERED_SERVICE",
		-4:   "UNREGISTERED_TOKEN",
		-9:   "TERMINATED_SERVICE",
		-10:  "LIMIT_EXCEEDED",
		-401: "UNAUTHORIZED",
		-999: "UNKNOWN",
	}
	ApiAuthRes_Code_value = map[string]int32{
		"VALID":                0,
		"INTERNAL_EXCEPTION":   -1,
		"PARAMETER_EXCEPTION":  -2,
		"UNREGISTERED_SERVICE": -3,
		"UNREGISTERED_TOKEN":   -4,
		"TERMINATED_SERVICE":   -9,
		"LIMIT_EXCEEDED":       -10,
		"UNAUTHORIZED":         -401,
		"UNKNOWN":              -999,
	}
)

func (x ApiAuthRes_Code) Enum() *ApiAuthRes_Code {
	p := new(ApiAuthRes_Code)
	*p = x
	return p
}

func (x ApiAuthRes_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiAuthRes_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_author_api_auth_proto_enumTypes[0].Descriptor()
}

func (ApiAuthRes_Code) Type() protoreflect.EnumType {
	return &file_proto_author_api_auth_proto_enumTypes[0]
}

func (x ApiAuthRes_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiAuthRes_Code.Descriptor instead.
func (ApiAuthRes_Code) EnumDescriptor() ([]byte, []int) {
	return file_proto_author_api_auth_proto_rawDescGZIP(), []int{1, 0}
}

type ApiAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameSpace    string `protobuf:"bytes,1,opt,name=name_space,json=nameSpace,proto3" json:"name_space,omitempty"`
	OperationUrl string `protobuf:"bytes,2,opt,name=operation_url,json=operationUrl,proto3" json:"operation_url,omitempty"`
	Token        string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ApiAuthReq) Reset() {
	*x = ApiAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_author_api_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiAuthReq) ProtoMessage() {}

func (x *ApiAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_author_api_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiAuthReq.ProtoReflect.Descriptor instead.
func (*ApiAuthReq) Descriptor() ([]byte, []int) {
	return file_proto_author_api_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ApiAuthReq) GetNameSpace() string {
	if x != nil {
		return x.NameSpace
	}
	return ""
}

func (x *ApiAuthReq) GetOperationUrl() string {
	if x != nil {
		return x.OperationUrl
	}
	return ""
}

func (x *ApiAuthReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ApiAuthRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ApiAuthRes_Code `protobuf:"varint,1,opt,name=code,proto3,enum=grpc_author.ApiAuthRes_Code" json:"code,omitempty"`
}

func (x *ApiAuthRes) Reset() {
	*x = ApiAuthRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_author_api_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiAuthRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiAuthRes) ProtoMessage() {}

func (x *ApiAuthRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_author_api_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiAuthRes.ProtoReflect.Descriptor instead.
func (*ApiAuthRes) Descriptor() ([]byte, []int) {
	return file_proto_author_api_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ApiAuthRes) GetCode() ApiAuthRes_Code {
	if x != nil {
		return x.Code
	}
	return ApiAuthRes_VALID
}

var File_proto_author_api_auth_proto protoreflect.FileDescriptor

var file_proto_author_api_auth_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x66, 0x0a, 0x0a, 0x41, 0x70,
	0x69, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0xc8, 0x02, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x12, 0x30, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x70,
	0x69, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x87, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x09, 0x0a, 0x05,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x12, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x20, 0x0a, 0x13, 0x50, 0x41, 0x52, 0x41,
	0x4d, 0x45, 0x54, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x21, 0x0a, 0x14, 0x55, 0x4e,
	0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x10, 0xfd, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1f, 0x0a,
	0x12, 0x55, 0x4e, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x54, 0x4f,
	0x4b, 0x45, 0x4e, 0x10, 0xfc, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1f,
	0x0a, 0x12, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0xf7, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12,
	0x1b, 0x0a, 0x0e, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45,
	0x44, 0x10, 0xf6, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x19, 0x0a, 0x0c,
	0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0xef, 0xfc, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x14, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x99, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x32, 0x4a, 0x0a,
	0x0e, 0x41, 0x70, 0x69, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x69, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41,
	0x70, 0x69, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x42, 0x1a, 0x5a, 0x18, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_author_api_auth_proto_rawDescOnce sync.Once
	file_proto_author_api_auth_proto_rawDescData = file_proto_author_api_auth_proto_rawDesc
)

func file_proto_author_api_auth_proto_rawDescGZIP() []byte {
	file_proto_author_api_auth_proto_rawDescOnce.Do(func() {
		file_proto_author_api_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_author_api_auth_proto_rawDescData)
	})
	return file_proto_author_api_auth_proto_rawDescData
}

var file_proto_author_api_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_author_api_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_author_api_auth_proto_goTypes = []interface{}{
	(ApiAuthRes_Code)(0), // 0: grpc_author.ApiAuthRes.Code
	(*ApiAuthReq)(nil),   // 1: grpc_author.ApiAuthReq
	(*ApiAuthRes)(nil),   // 2: grpc_author.ApiAuthRes
}
var file_proto_author_api_auth_proto_depIdxs = []int32{
	0, // 0: grpc_author.ApiAuthRes.code:type_name -> grpc_author.ApiAuthRes.Code
	1, // 1: grpc_author.ApiAuthService.Auth:input_type -> grpc_author.ApiAuthReq
	2, // 2: grpc_author.ApiAuthService.Auth:output_type -> grpc_author.ApiAuthRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_author_api_auth_proto_init() }
func file_proto_author_api_auth_proto_init() {
	if File_proto_author_api_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_author_api_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiAuthReq); i {
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
		file_proto_author_api_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiAuthRes); i {
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
			RawDescriptor: file_proto_author_api_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_author_api_auth_proto_goTypes,
		DependencyIndexes: file_proto_author_api_auth_proto_depIdxs,
		EnumInfos:         file_proto_author_api_auth_proto_enumTypes,
		MessageInfos:      file_proto_author_api_auth_proto_msgTypes,
	}.Build()
	File_proto_author_api_auth_proto = out.File
	file_proto_author_api_auth_proto_rawDesc = nil
	file_proto_author_api_auth_proto_goTypes = nil
	file_proto_author_api_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApiAuthServiceClient is the client API for ApiAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiAuthServiceClient interface {
	Auth(ctx context.Context, in *ApiAuthReq, opts ...grpc.CallOption) (*ApiAuthRes, error)
}

type apiAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiAuthServiceClient(cc grpc.ClientConnInterface) ApiAuthServiceClient {
	return &apiAuthServiceClient{cc}
}

func (c *apiAuthServiceClient) Auth(ctx context.Context, in *ApiAuthReq, opts ...grpc.CallOption) (*ApiAuthRes, error) {
	out := new(ApiAuthRes)
	err := c.cc.Invoke(ctx, "/grpc_author.ApiAuthService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiAuthServiceServer is the server API for ApiAuthService service.
type ApiAuthServiceServer interface {
	Auth(context.Context, *ApiAuthReq) (*ApiAuthRes, error)
}

// UnimplementedApiAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApiAuthServiceServer struct {
}

func (*UnimplementedApiAuthServiceServer) Auth(context.Context, *ApiAuthReq) (*ApiAuthRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}

func RegisterApiAuthServiceServer(s *grpc.Server, srv ApiAuthServiceServer) {
	s.RegisterService(&_ApiAuthService_serviceDesc, srv)
}

func _ApiAuthService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApiAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiAuthServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_author.ApiAuthService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiAuthServiceServer).Auth(ctx, req.(*ApiAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiAuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_author.ApiAuthService",
	HandlerType: (*ApiAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _ApiAuthService_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/author/api_auth.proto",
}
