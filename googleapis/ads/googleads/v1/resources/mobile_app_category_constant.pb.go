// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/mobile_app_category_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A mobile application category constant.
type MobileAppCategoryConstant struct {
	// Output only. The resource name of the mobile app category constant.
	// Mobile app category constant resource names have the form:
	//
	// `mobileAppCategoryConstants/{mobile_app_category_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the mobile app category constant.
	Id *wrappers.Int32Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. Mobile app category name.
	Name                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MobileAppCategoryConstant) Reset()         { *m = MobileAppCategoryConstant{} }
func (m *MobileAppCategoryConstant) String() string { return proto.CompactTextString(m) }
func (*MobileAppCategoryConstant) ProtoMessage()    {}
func (*MobileAppCategoryConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d66e9d0b240fe19, []int{0}
}

func (m *MobileAppCategoryConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileAppCategoryConstant.Unmarshal(m, b)
}
func (m *MobileAppCategoryConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileAppCategoryConstant.Marshal(b, m, deterministic)
}
func (m *MobileAppCategoryConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileAppCategoryConstant.Merge(m, src)
}
func (m *MobileAppCategoryConstant) XXX_Size() int {
	return xxx_messageInfo_MobileAppCategoryConstant.Size(m)
}
func (m *MobileAppCategoryConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileAppCategoryConstant.DiscardUnknown(m)
}

var xxx_messageInfo_MobileAppCategoryConstant proto.InternalMessageInfo

func (m *MobileAppCategoryConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MobileAppCategoryConstant) GetId() *wrappers.Int32Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MobileAppCategoryConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*MobileAppCategoryConstant)(nil), "google.ads.googleads.v1.resources.MobileAppCategoryConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/mobile_app_category_constant.proto", fileDescriptor_1d66e9d0b240fe19)
}

var fileDescriptor_1d66e9d0b240fe19 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcf, 0x6a, 0x14, 0x31,
	0x18, 0x67, 0x33, 0x22, 0x38, 0xea, 0x65, 0x4e, 0x6d, 0x2d, 0xb5, 0x15, 0x0a, 0x3d, 0x25, 0xce,
	0x56, 0x0f, 0xc6, 0x53, 0xb6, 0x42, 0x51, 0x50, 0xca, 0x0a, 0x7b, 0x90, 0x81, 0x21, 0x33, 0x49,
	0x63, 0x60, 0x26, 0x5f, 0x48, 0xb2, 0x2b, 0x22, 0x3e, 0x81, 0x6f, 0xe1, 0xd1, 0x47, 0xf1, 0x0d,
	0xbc, 0xf5, 0xdc, 0x47, 0xf0, 0x24, 0x3b, 0x93, 0x99, 0x2e, 0xc8, 0x56, 0xf0, 0xf6, 0x0b, 0xbf,
	0x3f, 0xdf, 0xf7, 0xe3, 0x4b, 0xfa, 0x4a, 0x01, 0xa8, 0x46, 0x12, 0x2e, 0x3c, 0xe9, 0xe1, 0x1a,
	0xad, 0x72, 0xe2, 0xa4, 0x87, 0xa5, 0xab, 0xa5, 0x27, 0x2d, 0x54, 0xba, 0x91, 0x25, 0xb7, 0xb6,
	0xac, 0x79, 0x90, 0x0a, 0xdc, 0xe7, 0xb2, 0x06, 0xe3, 0x03, 0x37, 0x01, 0x5b, 0x07, 0x01, 0xb2,
	0xa3, 0xde, 0x8a, 0xb9, 0xf0, 0x78, 0x4c, 0xc1, 0xab, 0x1c, 0x8f, 0x29, 0x7b, 0x8f, 0x87, 0x41,
	0x56, 0x93, 0x4b, 0x2d, 0x1b, 0x51, 0x56, 0xf2, 0x23, 0x5f, 0x69, 0x70, 0x7d, 0xc6, 0xde, 0xee,
	0x86, 0x60, 0xb0, 0x45, 0xea, 0x20, 0x52, 0xdd, 0xab, 0x5a, 0x5e, 0x92, 0x4f, 0x8e, 0x5b, 0x2b,
	0x9d, 0x8f, 0xfc, 0xfe, 0x86, 0x95, 0x1b, 0x03, 0x81, 0x07, 0x0d, 0x26, 0xb2, 0x4f, 0x7e, 0xa1,
	0x74, 0xf7, 0x6d, 0xd7, 0x81, 0x59, 0x7b, 0x16, 0x1b, 0x9c, 0xc5, 0x02, 0x59, 0x99, 0x3e, 0x1c,
	0xa6, 0x95, 0x86, 0xb7, 0x72, 0x67, 0x72, 0x38, 0x39, 0xb9, 0x37, 0xa3, 0x57, 0x2c, 0xf9, 0xcd,
	0x9e, 0xa5, 0xd3, 0x9b, 0x3a, 0x11, 0x59, 0xed, 0x71, 0x0d, 0x2d, 0xd9, 0x1a, 0x39, 0x7f, 0x30,
	0x04, 0xbe, 0xe3, 0xad, 0xcc, 0x9e, 0xa6, 0x48, 0x8b, 0x1d, 0x74, 0x38, 0x39, 0xb9, 0x3f, 0x7d,
	0x14, 0x43, 0xf0, 0xd0, 0x04, 0xbf, 0x36, 0xe1, 0x74, 0xba, 0xe0, 0xcd, 0x52, 0xce, 0x92, 0x2b,
	0x96, 0xcc, 0x91, 0x16, 0xd9, 0xf3, 0xf4, 0x4e, 0xb7, 0x49, 0xd2, 0x79, 0xf6, 0xff, 0xf2, 0xbc,
	0x0f, 0x4e, 0x1b, 0xb5, 0x61, 0xea, 0xe4, 0xd4, 0x5d, 0x33, 0xf8, 0x9f, 0x7d, 0xb3, 0x17, 0xed,
	0x36, 0xca, 0x93, 0x2f, 0xb7, 0x5d, 0xff, 0xeb, 0xec, 0x1b, 0x4a, 0x8f, 0x6b, 0x68, 0xf1, 0x3f,
	0xef, 0x3f, 0x3b, 0xd8, 0x3a, 0xff, 0x62, 0xdd, 0xeb, 0x62, 0xf2, 0xe1, 0x4d, 0x0c, 0x51, 0xd0,
	0x70, 0xa3, 0x30, 0x38, 0x45, 0x94, 0x34, 0x5d, 0x6b, 0x72, 0xd3, 0xe5, 0x96, 0x9f, 0xfa, 0x72,
	0x44, 0xdf, 0x51, 0x72, 0xce, 0xd8, 0x0f, 0x74, 0x74, 0xde, 0x47, 0x32, 0xe1, 0x71, 0x0f, 0xd7,
	0x68, 0x91, 0xe3, 0xf9, 0xa0, 0xfc, 0x39, 0x68, 0x0a, 0x26, 0x7c, 0x31, 0x6a, 0x8a, 0x45, 0x5e,
	0x8c, 0x9a, 0x6b, 0x74, 0xdc, 0x13, 0x94, 0x32, 0xe1, 0x29, 0x1d, 0x55, 0x94, 0x2e, 0x72, 0x4a,
	0x47, 0x5d, 0x75, 0xb7, 0x5b, 0xf6, 0xf4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x33, 0x57,
	0xa8, 0x55, 0x03, 0x00, 0x00,
}