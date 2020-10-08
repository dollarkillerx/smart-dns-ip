// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: smart_dns_ip.proto

package smart_dns_ip

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

type IPSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *IPSearchRequest) Reset() {
	*x = IPSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smart_dns_ip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPSearchRequest) ProtoMessage() {}

func (x *IPSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smart_dns_ip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPSearchRequest.ProtoReflect.Descriptor instead.
func (*IPSearchRequest) Descriptor() ([]byte, []int) {
	return file_smart_dns_ip_proto_rawDescGZIP(), []int{0}
}

func (x *IPSearchRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type IPSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityId  int64  `protobuf:"varint,1,opt,name=cityId,proto3" json:"cityId,omitempty"`
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Region  string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	City    string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Isp     string `protobuf:"bytes,5,opt,name=isp,proto3" json:"isp,omitempty"`
}

func (x *IPSearchResponse) Reset() {
	*x = IPSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smart_dns_ip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPSearchResponse) ProtoMessage() {}

func (x *IPSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smart_dns_ip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPSearchResponse.ProtoReflect.Descriptor instead.
func (*IPSearchResponse) Descriptor() ([]byte, []int) {
	return file_smart_dns_ip_proto_rawDescGZIP(), []int{1}
}

func (x *IPSearchResponse) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *IPSearchResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *IPSearchResponse) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *IPSearchResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *IPSearchResponse) GetIsp() string {
	if x != nil {
		return x.Isp
	}
	return ""
}

var File_smart_dns_ip_proto protoreflect.FileDescriptor

var file_smart_dns_ip_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x69, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x6e, 0x73, 0x5f,
	0x69, 0x70, 0x22, 0x21, 0x0a, 0x0f, 0x49, 0x50, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x82, 0x01, 0x0a, 0x10, 0x49, 0x50, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x73, 0x70, 0x32, 0x55, 0x0a, 0x08, 0x49, 0x50,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x49, 0x0a, 0x08, 0x49, 0x50, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x1d, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x69,
	0x70, 0x2e, 0x49, 0x50, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x69, 0x70,
	0x2e, 0x49, 0x50, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smart_dns_ip_proto_rawDescOnce sync.Once
	file_smart_dns_ip_proto_rawDescData = file_smart_dns_ip_proto_rawDesc
)

func file_smart_dns_ip_proto_rawDescGZIP() []byte {
	file_smart_dns_ip_proto_rawDescOnce.Do(func() {
		file_smart_dns_ip_proto_rawDescData = protoimpl.X.CompressGZIP(file_smart_dns_ip_proto_rawDescData)
	})
	return file_smart_dns_ip_proto_rawDescData
}

var file_smart_dns_ip_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_smart_dns_ip_proto_goTypes = []interface{}{
	(*IPSearchRequest)(nil),  // 0: smart_dns_ip.IPSearchRequest
	(*IPSearchResponse)(nil), // 1: smart_dns_ip.IPSearchResponse
}
var file_smart_dns_ip_proto_depIdxs = []int32{
	0, // 0: smart_dns_ip.IPSearch.IPSearch:input_type -> smart_dns_ip.IPSearchRequest
	1, // 1: smart_dns_ip.IPSearch.IPSearch:output_type -> smart_dns_ip.IPSearchResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_smart_dns_ip_proto_init() }
func file_smart_dns_ip_proto_init() {
	if File_smart_dns_ip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smart_dns_ip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPSearchRequest); i {
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
		file_smart_dns_ip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPSearchResponse); i {
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
			RawDescriptor: file_smart_dns_ip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smart_dns_ip_proto_goTypes,
		DependencyIndexes: file_smart_dns_ip_proto_depIdxs,
		MessageInfos:      file_smart_dns_ip_proto_msgTypes,
	}.Build()
	File_smart_dns_ip_proto = out.File
	file_smart_dns_ip_proto_rawDesc = nil
	file_smart_dns_ip_proto_goTypes = nil
	file_smart_dns_ip_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IPSearchClient is the client API for IPSearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IPSearchClient interface {
	IPSearch(ctx context.Context, in *IPSearchRequest, opts ...grpc.CallOption) (*IPSearchResponse, error)
}

type iPSearchClient struct {
	cc grpc.ClientConnInterface
}

func NewIPSearchClient(cc grpc.ClientConnInterface) IPSearchClient {
	return &iPSearchClient{cc}
}

func (c *iPSearchClient) IPSearch(ctx context.Context, in *IPSearchRequest, opts ...grpc.CallOption) (*IPSearchResponse, error) {
	out := new(IPSearchResponse)
	err := c.cc.Invoke(ctx, "/smart_dns_ip.IPSearch/IPSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IPSearchServer is the server API for IPSearch service.
type IPSearchServer interface {
	IPSearch(context.Context, *IPSearchRequest) (*IPSearchResponse, error)
}

// UnimplementedIPSearchServer can be embedded to have forward compatible implementations.
type UnimplementedIPSearchServer struct {
}

func (*UnimplementedIPSearchServer) IPSearch(context.Context, *IPSearchRequest) (*IPSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IPSearch not implemented")
}

func RegisterIPSearchServer(s *grpc.Server, srv IPSearchServer) {
	s.RegisterService(&_IPSearch_serviceDesc, srv)
}

func _IPSearch_IPSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPSearchServer).IPSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smart_dns_ip.IPSearch/IPSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPSearchServer).IPSearch(ctx, req.(*IPSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IPSearch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smart_dns_ip.IPSearch",
	HandlerType: (*IPSearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IPSearch",
			Handler:    _IPSearch_IPSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smart_dns_ip.proto",
}
