// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/user_interest.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A user interest: a particular interest-based vertical to be targeted.
type UserInterest struct {
	// Output only. The resource name of the user interest.
	// User interest resource names have the form:
	//
	// `customers/{customer_id}/userInterests/{user_interest_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Taxonomy type of the user interest.
	TaxonomyType enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType `protobuf:"varint,2,opt,name=taxonomy_type,json=taxonomyType,proto3,enum=google.ads.googleads.v1.enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType" json:"taxonomy_type,omitempty"`
	// Output only. The ID of the user interest.
	UserInterestId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=user_interest_id,json=userInterestId,proto3" json:"user_interest_id,omitempty"`
	// Output only. The name of the user interest.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The parent of the user interest.
	UserInterestParent *wrappers.StringValue `protobuf:"bytes,5,opt,name=user_interest_parent,json=userInterestParent,proto3" json:"user_interest_parent,omitempty"`
	// Output only. True if the user interest is launched to all channels and locales.
	LaunchedToAll *wrappers.BoolValue `protobuf:"bytes,6,opt,name=launched_to_all,json=launchedToAll,proto3" json:"launched_to_all,omitempty"`
	// Output only. Availability information of the user interest.
	Availabilities       []*common.CriterionCategoryAvailability `protobuf:"bytes,7,rep,name=availabilities,proto3" json:"availabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *UserInterest) Reset()         { *m = UserInterest{} }
func (m *UserInterest) String() string { return proto.CompactTextString(m) }
func (*UserInterest) ProtoMessage()    {}
func (*UserInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01746905ff80465b, []int{0}
}

func (m *UserInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInterest.Unmarshal(m, b)
}
func (m *UserInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInterest.Marshal(b, m, deterministic)
}
func (m *UserInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInterest.Merge(m, src)
}
func (m *UserInterest) XXX_Size() int {
	return xxx_messageInfo_UserInterest.Size(m)
}
func (m *UserInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInterest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInterest proto.InternalMessageInfo

func (m *UserInterest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *UserInterest) GetTaxonomyType() enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType {
	if m != nil {
		return m.TaxonomyType
	}
	return enums.UserInterestTaxonomyTypeEnum_UNSPECIFIED
}

func (m *UserInterest) GetUserInterestId() *wrappers.Int64Value {
	if m != nil {
		return m.UserInterestId
	}
	return nil
}

func (m *UserInterest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserInterest) GetUserInterestParent() *wrappers.StringValue {
	if m != nil {
		return m.UserInterestParent
	}
	return nil
}

func (m *UserInterest) GetLaunchedToAll() *wrappers.BoolValue {
	if m != nil {
		return m.LaunchedToAll
	}
	return nil
}

func (m *UserInterest) GetAvailabilities() []*common.CriterionCategoryAvailability {
	if m != nil {
		return m.Availabilities
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInterest)(nil), "google.ads.googleads.v1.resources.UserInterest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/user_interest.proto", fileDescriptor_01746905ff80465b)
}

var fileDescriptor_01746905ff80465b = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0x9a, 0x6d, 0x88, 0xec, 0x03, 0x88, 0x78, 0x08, 0x63, 0x82, 0x0e, 0x69, 0xa2, 0x2f,
	0xd8, 0x6a, 0x61, 0x3c, 0x04, 0x21, 0x94, 0x0e, 0x34, 0x6d, 0x0f, 0x68, 0x2a, 0xa5, 0x0f, 0x50,
	0x11, 0xb9, 0x89, 0x97, 0x19, 0x39, 0x76, 0x64, 0x3b, 0x85, 0x6a, 0x2a, 0x3f, 0x86, 0x47, 0x7e,
	0x0a, 0xbf, 0x62, 0xcf, 0xfb, 0x07, 0xf0, 0x80, 0x50, 0x13, 0x27, 0x73, 0x99, 0x0a, 0xe3, 0xed,
	0xb6, 0xf7, 0x9c, 0x73, 0x4f, 0xce, 0xd5, 0xb5, 0xb3, 0x9b, 0x70, 0x9e, 0x50, 0x0c, 0x51, 0x2c,
	0x61, 0x59, 0xce, 0xaa, 0x71, 0x1b, 0x0a, 0x2c, 0x79, 0x2e, 0x22, 0x2c, 0x61, 0x2e, 0xb1, 0x08,
	0x09, 0x53, 0x58, 0x60, 0xa9, 0x40, 0x26, 0xb8, 0xe2, 0xee, 0x76, 0x89, 0x05, 0x28, 0x96, 0xa0,
	0xa6, 0x81, 0x71, 0x1b, 0xd4, 0xb4, 0xcd, 0x97, 0x8b, 0x94, 0x23, 0x9e, 0xa6, 0x9c, 0xc1, 0x48,
	0x10, 0x85, 0x05, 0xe1, 0x2c, 0x8c, 0x90, 0xc2, 0x09, 0x17, 0x93, 0x10, 0x8d, 0x11, 0xa1, 0x68,
	0x44, 0x28, 0x51, 0x93, 0x72, 0xd0, 0xe6, 0x8b, 0x45, 0x2a, 0x98, 0xe5, 0xe9, 0x1f, 0xde, 0x42,
	0x85, 0x3e, 0x73, 0xc6, 0xd3, 0x49, 0xa8, 0x26, 0x19, 0xd6, 0x02, 0xf7, 0x2b, 0x81, 0x8c, 0xc0,
	0x63, 0x82, 0x69, 0x1c, 0x8e, 0xf0, 0x09, 0x1a, 0x13, 0x2e, 0x34, 0xe0, 0x8e, 0x01, 0xa8, 0xdc,
	0xeb, 0xd6, 0x3d, 0xdd, 0x2a, 0x7e, 0x8d, 0xf2, 0x63, 0xf8, 0x49, 0xa0, 0x2c, 0xc3, 0x42, 0xea,
	0xfe, 0x96, 0x41, 0x45, 0x8c, 0x71, 0x85, 0x14, 0xe1, 0x4c, 0x77, 0x1f, 0xfc, 0x58, 0x76, 0xd6,
	0xde, 0x4a, 0x2c, 0x0e, 0xb4, 0x3d, 0xb7, 0xe7, 0xac, 0x57, 0x03, 0x42, 0x86, 0x52, 0xec, 0x59,
	0x4d, 0xab, 0x75, 0xbd, 0xfb, 0xe8, 0x2c, 0xb0, 0x7f, 0x06, 0x0f, 0x9d, 0x9d, 0x8b, 0x20, 0x75,
	0x95, 0x11, 0x09, 0x22, 0x9e, 0x42, 0x53, 0xa5, 0xb7, 0x56, 0x69, 0xbc, 0x46, 0x29, 0x76, 0xbf,
	0x38, 0xeb, 0x73, 0x5f, 0xed, 0x35, 0x9a, 0x56, 0x6b, 0xa3, 0xd3, 0x07, 0x8b, 0x16, 0x54, 0xe4,
	0x06, 0x4c, 0xc5, 0xbe, 0xe6, 0xf7, 0x27, 0x19, 0x7e, 0xc5, 0xf2, 0x74, 0x61, 0xb3, 0x6b, 0x9f,
	0x05, 0x76, 0x6f, 0x4d, 0x19, 0x7f, 0xb9, 0x87, 0xce, 0xcd, 0xf9, 0x1d, 0x90, 0xd8, 0xb3, 0x9b,
	0x56, 0x6b, 0xb5, 0x73, 0xb7, 0xb2, 0x50, 0xa5, 0x07, 0x0e, 0x98, 0x7a, 0xfa, 0x64, 0x80, 0x68,
	0xae, 0x95, 0x36, 0x72, 0x63, 0xd0, 0x41, 0xec, 0xee, 0x3a, 0x4b, 0x45, 0x2c, 0x4b, 0x05, 0x7f,
	0xeb, 0x12, 0xff, 0x8d, 0x12, 0x84, 0x25, 0x86, 0x40, 0x01, 0x77, 0xa7, 0xce, 0xed, 0x79, 0x0b,
	0x19, 0x12, 0x98, 0x29, 0x6f, 0xf9, 0x0a, 0x32, 0xff, 0x99, 0xbd, 0x6b, 0x3a, 0x3e, 0x2a, 0xc6,
	0xb8, 0xfb, 0xce, 0x0d, 0x8a, 0x72, 0x16, 0x9d, 0xe0, 0x38, 0x54, 0x3c, 0x44, 0x94, 0x7a, 0x2b,
	0xc5, 0xe4, 0xcd, 0x4b, 0x93, 0xbb, 0x9c, 0x53, 0xc3, 0xfe, 0x7a, 0xc5, 0xeb, 0xf3, 0x80, 0x52,
	0xf7, 0xa3, 0xb3, 0x61, 0x1c, 0x00, 0xc1, 0xd2, 0xbb, 0xd6, 0xb4, 0x5b, 0xab, 0x9d, 0xe7, 0x0b,
	0x77, 0x59, 0x5e, 0x12, 0xd8, 0xab, 0x2e, 0x69, 0x4f, 0x1f, 0x52, 0x60, 0xdc, 0x91, 0x8e, 0x7a,
	0x5e, 0xd9, 0xff, 0x70, 0x1e, 0xbc, 0xbf, 0xe2, 0x47, 0xbb, 0x9d, 0x28, 0x97, 0x8a, 0xa7, 0x58,
	0x48, 0x78, 0x5a, 0x95, 0x53, 0x68, 0x26, 0x21, 0xe1, 0xe9, 0xdc, 0x06, 0xa6, 0xdd, 0x5f, 0x96,
	0xb3, 0x13, 0xf1, 0x14, 0xfc, 0xf3, 0x99, 0xe8, 0xde, 0x32, 0x67, 0x1d, 0xcd, 0x92, 0x3a, 0xb2,
	0xde, 0x1d, 0x6a, 0x5e, 0xc2, 0x29, 0x62, 0x09, 0xe0, 0x22, 0x81, 0x09, 0x66, 0x45, 0x8e, 0xf0,
	0xc2, 0xea, 0x5f, 0x1e, 0xad, 0x67, 0x75, 0xf5, 0xb5, 0x61, 0xef, 0x07, 0xc1, 0xb7, 0xc6, 0xf6,
	0x7e, 0x29, 0x19, 0xc4, 0x12, 0x94, 0xe5, 0xac, 0x1a, 0xb4, 0x41, 0xaf, 0x42, 0x7e, 0xaf, 0x30,
	0xc3, 0x20, 0x96, 0xc3, 0x1a, 0x33, 0x1c, 0xb4, 0x87, 0x35, 0xe6, 0xbc, 0xb1, 0x53, 0x36, 0x7c,
	0x3f, 0x88, 0xa5, 0xef, 0xd7, 0x28, 0xdf, 0x1f, 0xb4, 0x7d, 0xbf, 0xc6, 0x8d, 0x56, 0x0a, 0xb3,
	0x8f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x45, 0xb6, 0xad, 0x60, 0x05, 0x00, 0x00,
}