// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: proto/designer/application.proto

package grpc_designer

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_designer_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designer_application_proto_msgTypes[0]
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
	return file_proto_designer_application_proto_rawDescGZIP(), []int{0}
}

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NameSpace   string `protobuf:"bytes,2,opt,name=nameSpace,proto3" json:"nameSpace,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	UserId      int32  `protobuf:"varint,6,opt,name=userId,proto3" json:"userId,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_designer_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designer_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_proto_designer_application_proto_rawDescGZIP(), []int{1}
}

func (x *Application) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Application) GetNameSpace() string {
	if x != nil {
		return x.NameSpace
	}
	return ""
}

func (x *Application) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Application) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Application) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Application) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Application) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Application) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ApplicationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount   int32          `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	Page         int32          `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage      int32          `protobuf:"varint,3,opt,name=perPage,proto3" json:"perPage,omitempty"`
	Applications []*Application `protobuf:"bytes,4,rep,name=Applications,proto3" json:"Applications,omitempty"`
}

func (x *ApplicationList) Reset() {
	*x = ApplicationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_designer_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationList) ProtoMessage() {}

func (x *ApplicationList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designer_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationList.ProtoReflect.Descriptor instead.
func (*ApplicationList) Descriptor() ([]byte, []int) {
	return file_proto_designer_application_proto_rawDescGZIP(), []int{2}
}

func (x *ApplicationList) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ApplicationList) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ApplicationList) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ApplicationList) GetApplications() []*Application {
	if x != nil {
		return x.Applications
	}
	return nil
}

type ListApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PerPage   int32  `protobuf:"varint,1,opt,name=perPage,proto3" json:"perPage,omitempty"`
	Page      int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	NameSpace string `protobuf:"bytes,3,opt,name=nameSpace,proto3" json:"nameSpace,omitempty"`
}

func (x *ListApplicationRequest) Reset() {
	*x = ListApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_designer_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApplicationRequest) ProtoMessage() {}

func (x *ListApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_designer_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApplicationRequest.ProtoReflect.Descriptor instead.
func (*ListApplicationRequest) Descriptor() ([]byte, []int) {
	return file_proto_designer_application_proto_rawDescGZIP(), []int{3}
}

func (x *ListApplicationRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ListApplicationRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListApplicationRequest) GetNameSpace() string {
	if x != nil {
		return x.NameSpace
	}
	return ""
}

var File_proto_designer_application_proto protoreflect.FileDescriptor

var file_proto_designer_application_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xdf, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x64, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x32, 0x66, 0x0a, 0x12, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x50, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x72, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_designer_application_proto_rawDescOnce sync.Once
	file_proto_designer_application_proto_rawDescData = file_proto_designer_application_proto_rawDesc
)

func file_proto_designer_application_proto_rawDescGZIP() []byte {
	file_proto_designer_application_proto_rawDescOnce.Do(func() {
		file_proto_designer_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_designer_application_proto_rawDescData)
	})
	return file_proto_designer_application_proto_rawDescData
}

var file_proto_designer_application_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_designer_application_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: designer.Empty
	(*Application)(nil),            // 1: designer.Application
	(*ApplicationList)(nil),        // 2: designer.ApplicationList
	(*ListApplicationRequest)(nil), // 3: designer.ListApplicationRequest
}
var file_proto_designer_application_proto_depIdxs = []int32{
	1, // 0: designer.ApplicationList.Applications:type_name -> designer.Application
	3, // 1: designer.ApplicationService.ListApplication:input_type -> designer.ListApplicationRequest
	2, // 2: designer.ApplicationService.ListApplication:output_type -> designer.ApplicationList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_designer_application_proto_init() }
func file_proto_designer_application_proto_init() {
	if File_proto_designer_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_designer_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_designer_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_proto_designer_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationList); i {
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
		file_proto_designer_application_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListApplicationRequest); i {
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
			RawDescriptor: file_proto_designer_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_designer_application_proto_goTypes,
		DependencyIndexes: file_proto_designer_application_proto_depIdxs,
		MessageInfos:      file_proto_designer_application_proto_msgTypes,
	}.Build()
	File_proto_designer_application_proto = out.File
	file_proto_designer_application_proto_rawDesc = nil
	file_proto_designer_application_proto_goTypes = nil
	file_proto_designer_application_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApplicationServiceClient is the client API for ApplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationServiceClient interface {
	ListApplication(ctx context.Context, in *ListApplicationRequest, opts ...grpc.CallOption) (*ApplicationList, error)
}

type applicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApplicationServiceClient(cc grpc.ClientConnInterface) ApplicationServiceClient {
	return &applicationServiceClient{cc}
}

func (c *applicationServiceClient) ListApplication(ctx context.Context, in *ListApplicationRequest, opts ...grpc.CallOption) (*ApplicationList, error) {
	out := new(ApplicationList)
	err := c.cc.Invoke(ctx, "/designer.ApplicationService/ListApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationServiceServer is the server API for ApplicationService service.
type ApplicationServiceServer interface {
	ListApplication(context.Context, *ListApplicationRequest) (*ApplicationList, error)
}

// UnimplementedApplicationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationServiceServer struct {
}

func (*UnimplementedApplicationServiceServer) ListApplication(context.Context, *ListApplicationRequest) (*ApplicationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplication not implemented")
}

func RegisterApplicationServiceServer(s *grpc.Server, srv ApplicationServiceServer) {
	s.RegisterService(&_ApplicationService_serviceDesc, srv)
}

func _ApplicationService_ListApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServiceServer).ListApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/designer.ApplicationService/ListApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServiceServer).ListApplication(ctx, req.(*ListApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "designer.ApplicationService",
	HandlerType: (*ApplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListApplication",
			Handler:    _ApplicationService_ListApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/designer/application.proto",
}