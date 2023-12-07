// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/click_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for [ClickViewService.GetClickView][google.ads.googleads.v1.services.ClickViewService.GetClickView].
type GetClickViewRequest struct {
	// Required. The resource name of the click view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClickViewRequest) Reset()         { *m = GetClickViewRequest{} }
func (m *GetClickViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetClickViewRequest) ProtoMessage()    {}
func (*GetClickViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3c35caeb519e64a, []int{0}
}

func (m *GetClickViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClickViewRequest.Unmarshal(m, b)
}
func (m *GetClickViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClickViewRequest.Marshal(b, m, deterministic)
}
func (m *GetClickViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClickViewRequest.Merge(m, src)
}
func (m *GetClickViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetClickViewRequest.Size(m)
}
func (m *GetClickViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClickViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClickViewRequest proto.InternalMessageInfo

func (m *GetClickViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetClickViewRequest)(nil), "google.ads.googleads.v1.services.GetClickViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/click_view_service.proto", fileDescriptor_a3c35caeb519e64a)
}

var fileDescriptor_a3c35caeb519e64a = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcb, 0xaa, 0xd3, 0x40,
	0x18, 0x26, 0x11, 0x04, 0x43, 0x05, 0x89, 0x88, 0x35, 0x0a, 0x96, 0xd2, 0x85, 0x94, 0x32, 0x43,
	0x2a, 0x2e, 0x1c, 0x71, 0x31, 0x55, 0xa8, 0x2b, 0x2d, 0x0a, 0x59, 0x48, 0x20, 0x4c, 0x93, 0xbf,
	0x71, 0x30, 0xc9, 0xd4, 0x4c, 0x9a, 0x2e, 0xc4, 0x8d, 0xaf, 0xe0, 0x1b, 0xb8, 0xf4, 0x0d, 0x7c,
	0x85, 0x6e, 0xdd, 0xb9, 0x72, 0xe1, 0x4a, 0xdf, 0xe0, 0x70, 0x16, 0x87, 0x74, 0x32, 0x69, 0x7a,
	0x38, 0xa1, 0xbb, 0x8f, 0xf9, 0x2e, 0xff, 0x6d, 0xac, 0xa7, 0xb1, 0x10, 0x71, 0x02, 0x98, 0x45,
	0x12, 0x2b, 0x58, 0xa1, 0xd2, 0xc5, 0x12, 0xf2, 0x92, 0x87, 0x20, 0x71, 0x98, 0xf0, 0xf0, 0x63,
	0x50, 0x72, 0xd8, 0x06, 0xf5, 0x1b, 0x5a, 0xe7, 0xa2, 0x10, 0xf6, 0x40, 0xe9, 0x11, 0x8b, 0x24,
	0x6a, 0xac, 0xa8, 0x74, 0x91, 0xb6, 0x3a, 0xd3, 0xae, 0xf0, 0x1c, 0xa4, 0xd8, 0xe4, 0xc7, 0xe9,
	0x2a, 0xd5, 0x79, 0xa0, 0x3d, 0x6b, 0x8e, 0x59, 0x96, 0x89, 0x82, 0x15, 0x5c, 0x64, 0xb2, 0x66,
	0xef, 0xb6, 0xd8, 0x30, 0xe1, 0x90, 0x15, 0x35, 0xf1, 0xb0, 0x45, 0xac, 0x38, 0x24, 0x51, 0xb0,
	0x84, 0x0f, 0xac, 0xe4, 0x22, 0xaf, 0x05, 0xf7, 0x5a, 0x02, 0x5d, 0x5e, 0x51, 0xc3, 0x95, 0x75,
	0x7b, 0x0e, 0xc5, 0x8b, 0xaa, 0x13, 0x8f, 0xc3, 0xf6, 0x2d, 0x7c, 0xda, 0x80, 0x2c, 0xec, 0x37,
	0xd6, 0x4d, 0x2d, 0x0c, 0x32, 0x96, 0x42, 0xdf, 0x18, 0x18, 0x8f, 0x6e, 0xcc, 0xc6, 0x7f, 0xa8,
	0x79, 0x46, 0x47, 0xd6, 0xf0, 0x30, 0x73, 0x8d, 0xd6, 0x5c, 0xa2, 0x50, 0xa4, 0xf8, 0x90, 0xd4,
	0xd3, 0x01, 0xaf, 0x59, 0x0a, 0xd3, 0xff, 0x86, 0x75, 0xab, 0xe1, 0xde, 0xa9, 0x25, 0xd9, 0x3f,
	0x0d, 0xab, 0xd7, 0xae, 0x6e, 0x3f, 0x41, 0xa7, 0xf6, 0x8a, 0xae, 0xe8, 0xd6, 0x99, 0x74, 0xda,
	0x9a, 0x65, 0xa3, 0xc6, 0x34, 0x7c, 0xf9, 0x9b, 0x1e, 0x0f, 0xf7, 0xf5, 0xd7, 0xdf, 0x6f, 0x26,
	0xb2, 0x27, 0xd5, 0x75, 0x3e, 0x1f, 0x31, 0xcf, 0xc3, 0x8d, 0x2c, 0x44, 0x0a, 0xb9, 0xc4, 0x63,
	0x75, 0xae, 0x2a, 0x41, 0xe2, 0xf1, 0x17, 0xe7, 0xfe, 0x8e, 0xf6, 0xbb, 0xb6, 0x30, 0x3b, 0x37,
	0xac, 0x51, 0x28, 0xd2, 0x93, 0xd3, 0xcc, 0xee, 0x5c, 0xde, 0xc9, 0xa2, 0xba, 0xca, 0xc2, 0x78,
	0xff, 0xaa, 0xb6, 0xc6, 0x22, 0x61, 0x59, 0x8c, 0x44, 0x1e, 0xe3, 0x18, 0xb2, 0xfd, 0xcd, 0xf0,
	0xa1, 0x58, 0xf7, 0xd7, 0x7d, 0xa6, 0xc1, 0x77, 0xf3, 0xda, 0x9c, 0xd2, 0x1f, 0xe6, 0x60, 0xae,
	0x02, 0x69, 0x24, 0x91, 0x82, 0x15, 0xf2, 0x5c, 0x54, 0x17, 0x96, 0x3b, 0x2d, 0xf1, 0x69, 0x24,
	0xfd, 0x46, 0xe2, 0x7b, 0xae, 0xaf, 0x25, 0xff, 0xcc, 0x91, 0x7a, 0x27, 0x84, 0x46, 0x92, 0x90,
	0x46, 0x44, 0x88, 0xe7, 0x12, 0xa2, 0x65, 0xcb, 0xeb, 0xfb, 0x3e, 0x1f, 0x5f, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x17, 0x09, 0xee, 0xb5, 0x61, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClickViewServiceClient is the client API for ClickViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClickViewServiceClient interface {
	// Returns the requested click view in full detail.
	GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error)
}

type clickViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClickViewServiceClient(cc grpc.ClientConnInterface) ClickViewServiceClient {
	return &clickViewServiceClient{cc}
}

func (c *clickViewServiceClient) GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error) {
	out := new(resources.ClickView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ClickViewService/GetClickView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClickViewServiceServer is the server API for ClickViewService service.
type ClickViewServiceServer interface {
	// Returns the requested click view in full detail.
	GetClickView(context.Context, *GetClickViewRequest) (*resources.ClickView, error)
}

// UnimplementedClickViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClickViewServiceServer struct {
}

func (*UnimplementedClickViewServiceServer) GetClickView(ctx context.Context, req *GetClickViewRequest) (*resources.ClickView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClickView not implemented")
}

func RegisterClickViewServiceServer(s *grpc.Server, srv ClickViewServiceServer) {
	s.RegisterService(&_ClickViewService_serviceDesc, srv)
}

func _ClickViewService_GetClickView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClickViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickViewServiceServer).GetClickView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ClickViewService/GetClickView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickViewServiceServer).GetClickView(ctx, req.(*GetClickViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClickViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ClickViewService",
	HandlerType: (*ClickViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClickView",
			Handler:    _ClickViewService_GetClickView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/click_view_service.proto",
}
