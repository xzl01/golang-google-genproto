// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/keyword_plan_keyword_error.proto

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

// Enum describing possible errors from applying a keyword plan keyword.
type KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError int32

const (
	// Enum unspecified.
	KeywordPlanKeywordErrorEnum_UNSPECIFIED KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 0
	// The received error code is not known in this version.
	KeywordPlanKeywordErrorEnum_UNKNOWN KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 1
	// A keyword or negative keyword has invalid match type.
	KeywordPlanKeywordErrorEnum_INVALID_KEYWORD_MATCH_TYPE KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 2
	// A keyword or negative keyword with same text and match type already
	// exists.
	KeywordPlanKeywordErrorEnum_DUPLICATE_KEYWORD KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 3
	// Keyword or negative keyword text exceeds the allowed limit.
	KeywordPlanKeywordErrorEnum_KEYWORD_TEXT_TOO_LONG KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 4
	// Keyword or negative keyword text has invalid characters or symbols.
	KeywordPlanKeywordErrorEnum_KEYWORD_HAS_INVALID_CHARS KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 5
	// Keyword or negative keyword text has too many words.
	KeywordPlanKeywordErrorEnum_KEYWORD_HAS_TOO_MANY_WORDS KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 6
	// Keyword or negative keyword has invalid text.
	KeywordPlanKeywordErrorEnum_INVALID_KEYWORD_TEXT KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError = 7
)

var KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_KEYWORD_MATCH_TYPE",
	3: "DUPLICATE_KEYWORD",
	4: "KEYWORD_TEXT_TOO_LONG",
	5: "KEYWORD_HAS_INVALID_CHARS",
	6: "KEYWORD_HAS_TOO_MANY_WORDS",
	7: "INVALID_KEYWORD_TEXT",
}

var KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"INVALID_KEYWORD_MATCH_TYPE": 2,
	"DUPLICATE_KEYWORD":          3,
	"KEYWORD_TEXT_TOO_LONG":      4,
	"KEYWORD_HAS_INVALID_CHARS":  5,
	"KEYWORD_HAS_TOO_MANY_WORDS": 6,
	"INVALID_KEYWORD_TEXT":       7,
}

func (x KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError) String() string {
	return proto.EnumName(KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError_name, int32(x))
}

func (KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0ba4e4719ffcf21f, []int{0, 0}
}

// Container for enum describing possible errors from applying a keyword or a
// negative keyword from a keyword plan.
type KeywordPlanKeywordErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanKeywordErrorEnum) Reset()         { *m = KeywordPlanKeywordErrorEnum{} }
func (m *KeywordPlanKeywordErrorEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanKeywordErrorEnum) ProtoMessage()    {}
func (*KeywordPlanKeywordErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ba4e4719ffcf21f, []int{0}
}

func (m *KeywordPlanKeywordErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanKeywordErrorEnum.Unmarshal(m, b)
}
func (m *KeywordPlanKeywordErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanKeywordErrorEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanKeywordErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanKeywordErrorEnum.Merge(m, src)
}
func (m *KeywordPlanKeywordErrorEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanKeywordErrorEnum.Size(m)
}
func (m *KeywordPlanKeywordErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanKeywordErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanKeywordErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError", KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError_name, KeywordPlanKeywordErrorEnum_KeywordPlanKeywordError_value)
	proto.RegisterType((*KeywordPlanKeywordErrorEnum)(nil), "google.ads.googleads.v3.errors.KeywordPlanKeywordErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/keyword_plan_keyword_error.proto", fileDescriptor_0ba4e4719ffcf21f)
}

var fileDescriptor_0ba4e4719ffcf21f = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcd, 0x6e, 0xd3, 0x30,
	0x00, 0xa6, 0x19, 0x6c, 0x92, 0x77, 0xc0, 0x58, 0x4c, 0xb0, 0x31, 0x7a, 0xc8, 0x03, 0x38, 0x87,
	0xdc, 0xcc, 0x01, 0x79, 0x89, 0x69, 0xa3, 0x76, 0x49, 0xb4, 0xa4, 0x19, 0x45, 0x91, 0xac, 0x40,
	0xa2, 0xa8, 0x22, 0xb3, 0xa3, 0xb8, 0x0c, 0x71, 0xe5, 0x51, 0x38, 0xf2, 0x28, 0x3c, 0x06, 0x47,
	0x5e, 0x80, 0x2b, 0x72, 0x5c, 0x57, 0x08, 0xa9, 0x9c, 0xf2, 0xc5, 0xdf, 0x9f, 0xf5, 0x19, 0xbc,
	0x6e, 0xa5, 0x6c, 0xbb, 0xc6, 0xab, 0x6a, 0xe5, 0x19, 0xa8, 0xd1, 0xbd, 0xef, 0x35, 0xc3, 0x20,
	0x07, 0xe5, 0x7d, 0x6c, 0xbe, 0x7c, 0x96, 0x43, 0xcd, 0xfb, 0xae, 0x12, 0xdc, 0xfe, 0x8c, 0x1c,
	0xee, 0x07, 0xb9, 0x95, 0x68, 0x6a, 0x5c, 0xb8, 0xaa, 0x15, 0xde, 0x07, 0xe0, 0x7b, 0x1f, 0x9b,
	0x80, 0x8b, 0x4b, 0x5b, 0xd0, 0x6f, 0xbc, 0x4a, 0x08, 0xb9, 0xad, 0xb6, 0x1b, 0x29, 0x94, 0x71,
	0xbb, 0x5f, 0x1d, 0xf0, 0x62, 0x61, 0x52, 0xd3, 0xae, 0x12, 0x3b, 0xc8, 0xb4, 0x95, 0x89, 0x4f,
	0x77, 0xee, 0xcf, 0x09, 0x78, 0x76, 0x80, 0x47, 0x8f, 0xc1, 0xe9, 0x2a, 0xce, 0x52, 0x16, 0x44,
	0x6f, 0x22, 0x16, 0xc2, 0x07, 0xe8, 0x14, 0x9c, 0xac, 0xe2, 0x45, 0x9c, 0xdc, 0xc6, 0x70, 0x82,
	0xa6, 0xe0, 0x22, 0x8a, 0x0b, 0xba, 0x8c, 0x42, 0xbe, 0x60, 0xeb, 0xdb, 0xe4, 0x26, 0xe4, 0xd7,
	0x34, 0x0f, 0xe6, 0x3c, 0x5f, 0xa7, 0x0c, 0x3a, 0xe8, 0x0c, 0x3c, 0x09, 0x57, 0xe9, 0x32, 0x0a,
	0x68, 0xce, 0xac, 0x02, 0x1e, 0xa1, 0x73, 0x70, 0x66, 0xe5, 0x39, 0x7b, 0x9b, 0xf3, 0x3c, 0x49,
	0xf8, 0x32, 0x89, 0x67, 0xf0, 0x21, 0x7a, 0x09, 0xce, 0x2d, 0x35, 0xa7, 0x19, 0xb7, 0xe9, 0xc1,
	0x9c, 0xde, 0x64, 0xf0, 0x91, 0x2e, 0xfc, 0x9b, 0xd6, 0xc6, 0x6b, 0x1a, 0xaf, 0xb9, 0x3e, 0xc9,
	0xe0, 0x31, 0x7a, 0x0e, 0x9e, 0xfe, 0x7b, 0x21, 0xdd, 0x00, 0x4f, 0xae, 0x7e, 0x4f, 0x80, 0xfb,
	0x41, 0xde, 0xe1, 0xff, 0x2f, 0x79, 0x75, 0x79, 0x60, 0x88, 0x54, 0x2f, 0x99, 0x4e, 0xde, 0x85,
	0x3b, 0x7f, 0x2b, 0xbb, 0x4a, 0xb4, 0x58, 0x0e, 0xad, 0xd7, 0x36, 0x62, 0xdc, 0xd9, 0x3e, 0x6d,
	0xbf, 0x51, 0x87, 0x5e, 0xfa, 0x95, 0xf9, 0x7c, 0x73, 0x8e, 0x66, 0x94, 0x7e, 0x77, 0xa6, 0x33,
	0x13, 0x46, 0x6b, 0x85, 0x0d, 0xd4, 0xa8, 0xf0, 0xf1, 0x58, 0xa9, 0x7e, 0x58, 0x41, 0x49, 0x6b,
	0x55, 0xee, 0x05, 0x65, 0xe1, 0x97, 0x46, 0xf0, 0xcb, 0x71, 0xcd, 0x29, 0x21, 0xb4, 0x56, 0x84,
	0xec, 0x25, 0x84, 0x14, 0x3e, 0x21, 0x46, 0xf4, 0xfe, 0x78, 0xbc, 0x9d, 0xff, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x49, 0xc0, 0xcb, 0xc7, 0x86, 0x02, 0x00, 0x00,
}
