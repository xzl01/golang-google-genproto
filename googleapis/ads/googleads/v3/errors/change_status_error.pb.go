// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/change_status_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Enum describing possible change status errors.
type ChangeStatusErrorEnum_ChangeStatusError int32

const (
	// Enum unspecified.
	ChangeStatusErrorEnum_UNSPECIFIED ChangeStatusErrorEnum_ChangeStatusError = 0
	// The received error code is not known in this version.
	ChangeStatusErrorEnum_UNKNOWN ChangeStatusErrorEnum_ChangeStatusError = 1
	// The requested start date is too old.
	ChangeStatusErrorEnum_START_DATE_TOO_OLD ChangeStatusErrorEnum_ChangeStatusError = 3
)

var ChangeStatusErrorEnum_ChangeStatusError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "START_DATE_TOO_OLD",
}

var ChangeStatusErrorEnum_ChangeStatusError_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"START_DATE_TOO_OLD": 3,
}

func (x ChangeStatusErrorEnum_ChangeStatusError) String() string {
	return proto.EnumName(ChangeStatusErrorEnum_ChangeStatusError_name, int32(x))
}

func (ChangeStatusErrorEnum_ChangeStatusError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75754e84c034b454, []int{0, 0}
}

// Container for enum describing possible change status errors.
type ChangeStatusErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeStatusErrorEnum) Reset()         { *m = ChangeStatusErrorEnum{} }
func (m *ChangeStatusErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusErrorEnum) ProtoMessage()    {}
func (*ChangeStatusErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_75754e84c034b454, []int{0}
}

func (m *ChangeStatusErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusErrorEnum.Unmarshal(m, b)
}
func (m *ChangeStatusErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusErrorEnum.Marshal(b, m, deterministic)
}
func (m *ChangeStatusErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusErrorEnum.Merge(m, src)
}
func (m *ChangeStatusErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusErrorEnum.Size(m)
}
func (m *ChangeStatusErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.ChangeStatusErrorEnum_ChangeStatusError", ChangeStatusErrorEnum_ChangeStatusError_name, ChangeStatusErrorEnum_ChangeStatusError_value)
	proto.RegisterType((*ChangeStatusErrorEnum)(nil), "google.ads.googleads.v3.errors.ChangeStatusErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/change_status_error.proto", fileDescriptor_75754e84c034b454)
}

var fileDescriptor_75754e84c034b454 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0xdd, 0x06, 0x0a, 0xd9, 0xc1, 0x5a, 0x70, 0x07, 0x91, 0x1d, 0xfa, 0x00, 0xc9, 0x21,
	0x17, 0x89, 0xa7, 0x6c, 0xad, 0xa3, 0x28, 0x6d, 0xb1, 0x5d, 0x05, 0x29, 0x94, 0x6c, 0x2d, 0xb1,
	0xb0, 0x25, 0xa5, 0xe9, 0xf6, 0x40, 0x1e, 0x7d, 0x14, 0x1f, 0xc5, 0x9b, 0x6f, 0x20, 0x4d, 0x6c,
	0x2f, 0x43, 0x4f, 0xfd, 0xf8, 0xf7, 0xf7, 0x7d, 0xff, 0x2f, 0x7f, 0x70, 0xc7, 0xa5, 0xe4, 0xbb,
	0x12, 0xb1, 0x42, 0x21, 0x23, 0x3b, 0x75, 0xc4, 0xa8, 0x6c, 0x1a, 0xd9, 0x28, 0xb4, 0x7d, 0x63,
	0x82, 0x97, 0xb9, 0x6a, 0x59, 0x7b, 0x50, 0xb9, 0x1e, 0xc2, 0xba, 0x91, 0xad, 0xb4, 0xe7, 0x06,
	0x87, 0xac, 0x50, 0x70, 0x70, 0xc2, 0x23, 0x86, 0xc6, 0x79, 0x73, 0xdb, 0x27, 0xd7, 0x15, 0x62,
	0x42, 0xc8, 0x96, 0xb5, 0x95, 0x14, 0xca, 0xb8, 0x9d, 0x0d, 0xb8, 0x5e, 0xea, 0xe8, 0x58, 0x27,
	0x7b, 0x9d, 0xc7, 0x13, 0x87, 0xbd, 0xe3, 0x83, 0xab, 0x93, 0x1f, 0xf6, 0x25, 0x98, 0xae, 0x83,
	0x38, 0xf2, 0x96, 0xfe, 0x83, 0xef, 0xb9, 0xd6, 0x99, 0x3d, 0x05, 0x17, 0xeb, 0xe0, 0x31, 0x08,
	0x5f, 0x02, 0x6b, 0x64, 0xcf, 0x80, 0x1d, 0x27, 0xf4, 0x39, 0xc9, 0x5d, 0x9a, 0x78, 0x79, 0x12,
	0x86, 0x79, 0xf8, 0xe4, 0x5a, 0x93, 0xc5, 0xf7, 0x08, 0x38, 0x5b, 0xb9, 0x87, 0xff, 0x17, 0x5d,
	0xcc, 0x4e, 0xf6, 0x45, 0x5d, 0xc5, 0x68, 0xf4, 0xea, 0xfe, 0x3a, 0xb9, 0xdc, 0x31, 0xc1, 0xa1,
	0x6c, 0x38, 0xe2, 0xa5, 0xd0, 0x0f, 0xe8, 0x8f, 0x55, 0x57, 0xea, 0xaf, 0xdb, 0xdd, 0x9b, 0xcf,
	0xfb, 0x78, 0xb2, 0xa2, 0xf4, 0x63, 0x3c, 0x5f, 0x99, 0x30, 0x5a, 0x28, 0x68, 0x64, 0xa7, 0x52,
	0x0c, 0xf5, 0x4a, 0xf5, 0xd9, 0x03, 0x19, 0x2d, 0x54, 0x36, 0x00, 0x59, 0x8a, 0x33, 0x03, 0x7c,
	0x8d, 0x1d, 0x33, 0x25, 0x84, 0x16, 0x8a, 0x90, 0x01, 0x21, 0x24, 0xc5, 0x84, 0x18, 0x68, 0x73,
	0xae, 0xdb, 0xe1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x0a, 0xc5, 0x15, 0xd8, 0x01, 0x00,
	0x00,
}
