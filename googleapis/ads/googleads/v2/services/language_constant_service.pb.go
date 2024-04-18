// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/language_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [LanguageConstantService.GetLanguageConstant][google.ads.googleads.v2.services.LanguageConstantService.GetLanguageConstant].
type GetLanguageConstantRequest struct {
	// Required. Resource name of the language constant to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLanguageConstantRequest) Reset()         { *m = GetLanguageConstantRequest{} }
func (m *GetLanguageConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetLanguageConstantRequest) ProtoMessage()    {}
func (*GetLanguageConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc548c57933bf456, []int{0}
}

func (m *GetLanguageConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLanguageConstantRequest.Unmarshal(m, b)
}
func (m *GetLanguageConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLanguageConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetLanguageConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLanguageConstantRequest.Merge(m, src)
}
func (m *GetLanguageConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetLanguageConstantRequest.Size(m)
}
func (m *GetLanguageConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLanguageConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLanguageConstantRequest proto.InternalMessageInfo

func (m *GetLanguageConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetLanguageConstantRequest)(nil), "google.ads.googleads.v2.services.GetLanguageConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/language_constant_service.proto", fileDescriptor_bc548c57933bf456)
}

var fileDescriptor_bc548c57933bf456 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0x26, 0x11, 0x04, 0x83, 0x5e, 0xe2, 0xa1, 0x35, 0x16, 0x2c, 0xa5, 0xa0, 0x55, 0x98, 0xc1,
	0xf4, 0xe4, 0xa8, 0xe8, 0xd4, 0x43, 0x3d, 0x88, 0x14, 0x85, 0x1c, 0x24, 0x10, 0xa6, 0xc9, 0x18,
	0x03, 0xc9, 0x4c, 0xcd, 0x4c, 0x73, 0x11, 0x2f, 0xf5, 0x27, 0xf8, 0x0f, 0x3c, 0xfa, 0x4f, 0xec,
	0xd5, 0xdb, 0x9e, 0xf6, 0xb0, 0xa7, 0xfd, 0x09, 0x7b, 0x5a, 0xd2, 0x99, 0x49, 0x3f, 0xb6, 0xa1,
	0xb7, 0x87, 0x3c, 0x1f, 0xef, 0x33, 0xef, 0x1b, 0xe7, 0x6d, 0xca, 0x79, 0x9a, 0x53, 0x48, 0x12,
	0x01, 0x15, 0xac, 0x51, 0xe5, 0x43, 0x41, 0xcb, 0x2a, 0x8b, 0xa9, 0x80, 0x39, 0x61, 0xe9, 0x92,
	0xa4, 0x34, 0x8a, 0x39, 0x13, 0x92, 0x30, 0x19, 0x69, 0x0a, 0x2c, 0x4a, 0x2e, 0xb9, 0xdb, 0x57,
	0x36, 0x40, 0x12, 0x01, 0x9a, 0x04, 0x50, 0xf9, 0xc0, 0x24, 0x78, 0x2f, 0xda, 0x66, 0x94, 0x54,
	0xf0, 0x65, 0x79, 0x74, 0x88, 0x0a, 0xf7, 0x7a, 0xc6, 0xba, 0xc8, 0x20, 0x61, 0x8c, 0x4b, 0x22,
	0x33, 0xce, 0x84, 0x66, 0x3b, 0x3b, 0x6c, 0x9c, 0x67, 0xb4, 0xb1, 0x3d, 0xda, 0x21, 0xbe, 0x66,
	0x34, 0x4f, 0xa2, 0x39, 0xfd, 0x46, 0xaa, 0x8c, 0x97, 0x5a, 0xf0, 0x60, 0x47, 0x60, 0x5a, 0x28,
	0x6a, 0x20, 0x1d, 0x6f, 0x4a, 0xe5, 0x07, 0x5d, 0xe8, 0x9d, 0xee, 0xf3, 0x89, 0x7e, 0x5f, 0x52,
	0x21, 0xdd, 0xc0, 0xb9, 0x67, 0xf4, 0x11, 0x23, 0x05, 0xed, 0x5a, 0x7d, 0xeb, 0xc9, 0x9d, 0xc9,
	0xf3, 0x73, 0x6c, 0x5f, 0xe1, 0x67, 0xce, 0x68, 0xbb, 0x01, 0x8d, 0x16, 0x99, 0x00, 0x31, 0x2f,
	0xe0, 0x8d, 0xc0, 0xbb, 0x26, 0xe7, 0x23, 0x29, 0xa8, 0xbf, 0xb2, 0x9d, 0xce, 0xa1, 0xe4, 0xb3,
	0x5a, 0xa0, 0xfb, 0xcf, 0x72, 0xee, 0x1f, 0xa9, 0xe4, 0xbe, 0x02, 0xa7, 0x56, 0x0f, 0xda, 0x5f,
	0xe2, 0x8d, 0x5b, 0xdd, 0xcd, 0x59, 0xc0, 0xa1, 0x77, 0xf0, 0xe6, 0x0c, 0xef, 0xbf, 0x7f, 0xf5,
	0xff, 0xe2, 0xb7, 0x3d, 0x72, 0x1f, 0xd7, 0xe7, 0xfc, 0xb1, 0xc7, 0xbc, 0xce, 0x0f, 0xcc, 0x02,
	0x3e, 0xfd, 0xe9, 0x3d, 0x5c, 0xe3, 0x6e, 0xdb, 0x8e, 0x26, 0xbf, 0x6c, 0x67, 0x18, 0xf3, 0xe2,
	0xe4, 0xb3, 0x26, 0xbd, 0x96, 0x55, 0xcd, 0xea, 0x0b, 0xce, 0xac, 0x2f, 0xef, 0x75, 0x42, 0xca,
	0xeb, 0x12, 0x80, 0x97, 0x29, 0x4c, 0x29, 0xdb, 0xdc, 0x17, 0x6e, 0x67, 0xb6, 0xff, 0xf4, 0x2f,
	0x0d, 0xf8, 0x63, 0xdf, 0x9a, 0x62, 0xfc, 0xd7, 0xee, 0x4f, 0x55, 0x20, 0x4e, 0x04, 0x50, 0xb0,
	0x46, 0x81, 0x0f, 0xf4, 0x60, 0xb1, 0x36, 0x92, 0x10, 0x27, 0x22, 0x6c, 0x24, 0x61, 0xe0, 0x87,
	0x46, 0x72, 0x69, 0x0f, 0xd5, 0x77, 0x84, 0x70, 0x22, 0x10, 0x6a, 0x44, 0x08, 0x05, 0x3e, 0x42,
	0x46, 0x36, 0xbf, 0xbd, 0xe9, 0x39, 0xbe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xc5, 0xae, 0x9e,
	0x9b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LanguageConstantServiceClient is the client API for LanguageConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LanguageConstantServiceClient interface {
	// Returns the requested language constant.
	GetLanguageConstant(ctx context.Context, in *GetLanguageConstantRequest, opts ...grpc.CallOption) (*resources.LanguageConstant, error)
}

type languageConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLanguageConstantServiceClient(cc grpc.ClientConnInterface) LanguageConstantServiceClient {
	return &languageConstantServiceClient{cc}
}

func (c *languageConstantServiceClient) GetLanguageConstant(ctx context.Context, in *GetLanguageConstantRequest, opts ...grpc.CallOption) (*resources.LanguageConstant, error) {
	out := new(resources.LanguageConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.LanguageConstantService/GetLanguageConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LanguageConstantServiceServer is the server API for LanguageConstantService service.
type LanguageConstantServiceServer interface {
	// Returns the requested language constant.
	GetLanguageConstant(context.Context, *GetLanguageConstantRequest) (*resources.LanguageConstant, error)
}

// UnimplementedLanguageConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLanguageConstantServiceServer struct {
}

func (*UnimplementedLanguageConstantServiceServer) GetLanguageConstant(ctx context.Context, req *GetLanguageConstantRequest) (*resources.LanguageConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLanguageConstant not implemented")
}

func RegisterLanguageConstantServiceServer(s *grpc.Server, srv LanguageConstantServiceServer) {
	s.RegisterService(&_LanguageConstantService_serviceDesc, srv)
}

func _LanguageConstantService_GetLanguageConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLanguageConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageConstantServiceServer).GetLanguageConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.LanguageConstantService/GetLanguageConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageConstantServiceServer).GetLanguageConstant(ctx, req.(*GetLanguageConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LanguageConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.LanguageConstantService",
	HandlerType: (*LanguageConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLanguageConstant",
			Handler:    _LanguageConstantService_GetLanguageConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/language_constant_service.proto",
}
