// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/appengine/v1/domain_mapping.proto

package appengine

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The SSL management type for this domain.
type SslSettings_SslManagementType int32

const (
	// Defaults to `AUTOMATIC`.
	SslSettings_SSL_MANAGEMENT_TYPE_UNSPECIFIED SslSettings_SslManagementType = 0
	// SSL support for this domain is configured automatically. The mapped SSL
	// certificate will be automatically renewed.
	SslSettings_AUTOMATIC SslSettings_SslManagementType = 1
	// SSL support for this domain is configured manually by the user. Either
	// the domain has no SSL support or a user-obtained SSL certificate has been
	// explictly mapped to this domain.
	SslSettings_MANUAL SslSettings_SslManagementType = 2
)

// Enum value maps for SslSettings_SslManagementType.
var (
	SslSettings_SslManagementType_name = map[int32]string{
		0: "SSL_MANAGEMENT_TYPE_UNSPECIFIED",
		1: "AUTOMATIC",
		2: "MANUAL",
	}
	SslSettings_SslManagementType_value = map[string]int32{
		"SSL_MANAGEMENT_TYPE_UNSPECIFIED": 0,
		"AUTOMATIC":                       1,
		"MANUAL":                          2,
	}
)

func (x SslSettings_SslManagementType) Enum() *SslSettings_SslManagementType {
	p := new(SslSettings_SslManagementType)
	*p = x
	return p
}

func (x SslSettings_SslManagementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SslSettings_SslManagementType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_appengine_v1_domain_mapping_proto_enumTypes[0].Descriptor()
}

func (SslSettings_SslManagementType) Type() protoreflect.EnumType {
	return &file_google_appengine_v1_domain_mapping_proto_enumTypes[0]
}

func (x SslSettings_SslManagementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SslSettings_SslManagementType.Descriptor instead.
func (SslSettings_SslManagementType) EnumDescriptor() ([]byte, []int) {
	return file_google_appengine_v1_domain_mapping_proto_rawDescGZIP(), []int{1, 0}
}

// A resource record type.
type ResourceRecord_RecordType int32

const (
	// An unknown resource record.
	ResourceRecord_RECORD_TYPE_UNSPECIFIED ResourceRecord_RecordType = 0
	// An A resource record. Data is an IPv4 address.
	ResourceRecord_A ResourceRecord_RecordType = 1
	// An AAAA resource record. Data is an IPv6 address.
	ResourceRecord_AAAA ResourceRecord_RecordType = 2
	// A CNAME resource record. Data is a domain name to be aliased.
	ResourceRecord_CNAME ResourceRecord_RecordType = 3
)

// Enum value maps for ResourceRecord_RecordType.
var (
	ResourceRecord_RecordType_name = map[int32]string{
		0: "RECORD_TYPE_UNSPECIFIED",
		1: "A",
		2: "AAAA",
		3: "CNAME",
	}
	ResourceRecord_RecordType_value = map[string]int32{
		"RECORD_TYPE_UNSPECIFIED": 0,
		"A":                       1,
		"AAAA":                    2,
		"CNAME":                   3,
	}
)

func (x ResourceRecord_RecordType) Enum() *ResourceRecord_RecordType {
	p := new(ResourceRecord_RecordType)
	*p = x
	return p
}

func (x ResourceRecord_RecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceRecord_RecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_appengine_v1_domain_mapping_proto_enumTypes[1].Descriptor()
}

func (ResourceRecord_RecordType) Type() protoreflect.EnumType {
	return &file_google_appengine_v1_domain_mapping_proto_enumTypes[1]
}

func (x ResourceRecord_RecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceRecord_RecordType.Descriptor instead.
func (ResourceRecord_RecordType) EnumDescriptor() ([]byte, []int) {
	return file_google_appengine_v1_domain_mapping_proto_rawDescGZIP(), []int{2, 0}
}

// A domain serving an App Engine application.
type DomainMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full path to the `DomainMapping` resource in the API. Example:
	// `apps/myapp/domainMapping/example.com`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Relative name of the domain serving the application. Example:
	// `example.com`.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// SSL configuration for this domain. If unconfigured, this domain will not
	// serve with SSL.
	SslSettings *SslSettings `protobuf:"bytes,3,opt,name=ssl_settings,json=sslSettings,proto3" json:"ssl_settings,omitempty"`
	// The resource records required to configure this domain mapping. These
	// records must be added to the domain's DNS configuration in order to
	// serve the application via this domain mapping.
	//
	// @OutputOnly
	ResourceRecords []*ResourceRecord `protobuf:"bytes,4,rep,name=resource_records,json=resourceRecords,proto3" json:"resource_records,omitempty"`
}

func (x *DomainMapping) Reset() {
	*x = DomainMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_domain_mapping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainMapping) ProtoMessage() {}

func (x *DomainMapping) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_domain_mapping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainMapping.ProtoReflect.Descriptor instead.
func (*DomainMapping) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_domain_mapping_proto_rawDescGZIP(), []int{0}
}

func (x *DomainMapping) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DomainMapping) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DomainMapping) GetSslSettings() *SslSettings {
	if x != nil {
		return x.SslSettings
	}
	return nil
}

func (x *DomainMapping) GetResourceRecords() []*ResourceRecord {
	if x != nil {
		return x.ResourceRecords
	}
	return nil
}

// SSL configuration for a `DomainMapping` resource.
type SslSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the `AuthorizedCertificate` resource configuring SSL for the
	// application. Clearing this field will remove SSL support.
	//
	// By default, a managed certificate is automatically created for every
	// domain mapping. To omit SSL support or to configure SSL manually, specify
	// `SslManagementType.MANUAL` on a `CREATE` or `UPDATE` request. You must
	// be authorized to administer the `AuthorizedCertificate` resource to
	// manually map it to a `DomainMapping` resource.
	// Example: `12345`.
	CertificateId string `protobuf:"bytes,1,opt,name=certificate_id,json=certificateId,proto3" json:"certificate_id,omitempty"`
	// SSL management type for this domain. If `AUTOMATIC`, a managed certificate
	// is automatically provisioned. If `MANUAL`, `certificate_id` must be
	// manually specified in order to configure SSL for this domain.
	SslManagementType SslSettings_SslManagementType `protobuf:"varint,3,opt,name=ssl_management_type,json=sslManagementType,proto3,enum=google.appengine.v1.SslSettings_SslManagementType" json:"ssl_management_type,omitempty"`
	// ID of the managed `AuthorizedCertificate` resource currently being
	// provisioned, if applicable. Until the new managed certificate has been
	// successfully provisioned, the previous SSL state will be preserved. Once
	// the provisioning process completes, the `certificate_id` field will reflect
	// the new managed certificate and this field will be left empty. To remove
	// SSL support while there is still a pending managed certificate, clear the
	// `certificate_id` field with an `UpdateDomainMappingRequest`.
	//
	// @OutputOnly
	PendingManagedCertificateId string `protobuf:"bytes,4,opt,name=pending_managed_certificate_id,json=pendingManagedCertificateId,proto3" json:"pending_managed_certificate_id,omitempty"`
}

func (x *SslSettings) Reset() {
	*x = SslSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_domain_mapping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SslSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SslSettings) ProtoMessage() {}

func (x *SslSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_domain_mapping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SslSettings.ProtoReflect.Descriptor instead.
func (*SslSettings) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_domain_mapping_proto_rawDescGZIP(), []int{1}
}

func (x *SslSettings) GetCertificateId() string {
	if x != nil {
		return x.CertificateId
	}
	return ""
}

func (x *SslSettings) GetSslManagementType() SslSettings_SslManagementType {
	if x != nil {
		return x.SslManagementType
	}
	return SslSettings_SSL_MANAGEMENT_TYPE_UNSPECIFIED
}

func (x *SslSettings) GetPendingManagedCertificateId() string {
	if x != nil {
		return x.PendingManagedCertificateId
	}
	return ""
}

// A DNS resource record.
type ResourceRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Relative name of the object affected by this record. Only applicable for
	// `CNAME` records. Example: 'www'.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Data for this record. Values vary by record type, as defined in RFC 1035
	// (section 5) and RFC 1034 (section 3.6.1).
	Rrdata string `protobuf:"bytes,2,opt,name=rrdata,proto3" json:"rrdata,omitempty"`
	// Resource record type. Example: `AAAA`.
	Type ResourceRecord_RecordType `protobuf:"varint,3,opt,name=type,proto3,enum=google.appengine.v1.ResourceRecord_RecordType" json:"type,omitempty"`
}

func (x *ResourceRecord) Reset() {
	*x = ResourceRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_domain_mapping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceRecord) ProtoMessage() {}

func (x *ResourceRecord) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_domain_mapping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceRecord.ProtoReflect.Descriptor instead.
func (*ResourceRecord) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_domain_mapping_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceRecord) GetRrdata() string {
	if x != nil {
		return x.Rrdata
	}
	return ""
}

func (x *ResourceRecord) GetType() ResourceRecord_RecordType {
	if x != nil {
		return x.Type
	}
	return ResourceRecord_RECORD_TYPE_UNSPECIFIED
}

var File_google_appengine_v1_domain_mapping_proto protoreflect.FileDescriptor

var file_google_appengine_v1_domain_mapping_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01,
	0x0a, 0x0d, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0c, 0x73, 0x73, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x73, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0b, 0x73, 0x73, 0x6c,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4e, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xb2, 0x02, 0x0a, 0x0b, 0x53, 0x73, 0x6c,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x62, 0x0a, 0x13, 0x73, 0x73, 0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x73, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53,
	0x73, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x11, 0x73, 0x73, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x1e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x11, 0x53, 0x73, 0x6c, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a,
	0x1f, 0x53, 0x53, 0x4c, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x55, 0x54, 0x4f, 0x4d, 0x41, 0x54, 0x49, 0x43, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x10, 0x02, 0x22, 0xc7, 0x01,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x72, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x72, 0x64, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x45, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x41,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x41, 0x41, 0x41, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x43, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x03, 0x42, 0xc4, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x12, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0xaa, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_appengine_v1_domain_mapping_proto_rawDescOnce sync.Once
	file_google_appengine_v1_domain_mapping_proto_rawDescData = file_google_appengine_v1_domain_mapping_proto_rawDesc
)

func file_google_appengine_v1_domain_mapping_proto_rawDescGZIP() []byte {
	file_google_appengine_v1_domain_mapping_proto_rawDescOnce.Do(func() {
		file_google_appengine_v1_domain_mapping_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_appengine_v1_domain_mapping_proto_rawDescData)
	})
	return file_google_appengine_v1_domain_mapping_proto_rawDescData
}

var file_google_appengine_v1_domain_mapping_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_appengine_v1_domain_mapping_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_appengine_v1_domain_mapping_proto_goTypes = []interface{}{
	(SslSettings_SslManagementType)(0), // 0: google.appengine.v1.SslSettings.SslManagementType
	(ResourceRecord_RecordType)(0),     // 1: google.appengine.v1.ResourceRecord.RecordType
	(*DomainMapping)(nil),              // 2: google.appengine.v1.DomainMapping
	(*SslSettings)(nil),                // 3: google.appengine.v1.SslSettings
	(*ResourceRecord)(nil),             // 4: google.appengine.v1.ResourceRecord
}
var file_google_appengine_v1_domain_mapping_proto_depIdxs = []int32{
	3, // 0: google.appengine.v1.DomainMapping.ssl_settings:type_name -> google.appengine.v1.SslSettings
	4, // 1: google.appengine.v1.DomainMapping.resource_records:type_name -> google.appengine.v1.ResourceRecord
	0, // 2: google.appengine.v1.SslSettings.ssl_management_type:type_name -> google.appengine.v1.SslSettings.SslManagementType
	1, // 3: google.appengine.v1.ResourceRecord.type:type_name -> google.appengine.v1.ResourceRecord.RecordType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_appengine_v1_domain_mapping_proto_init() }
func file_google_appengine_v1_domain_mapping_proto_init() {
	if File_google_appengine_v1_domain_mapping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_appengine_v1_domain_mapping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainMapping); i {
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
		file_google_appengine_v1_domain_mapping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SslSettings); i {
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
		file_google_appengine_v1_domain_mapping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceRecord); i {
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
			RawDescriptor: file_google_appengine_v1_domain_mapping_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_appengine_v1_domain_mapping_proto_goTypes,
		DependencyIndexes: file_google_appengine_v1_domain_mapping_proto_depIdxs,
		EnumInfos:         file_google_appengine_v1_domain_mapping_proto_enumTypes,
		MessageInfos:      file_google_appengine_v1_domain_mapping_proto_msgTypes,
	}.Build()
	File_google_appengine_v1_domain_mapping_proto = out.File
	file_google_appengine_v1_domain_mapping_proto_rawDesc = nil
	file_google_appengine_v1_domain_mapping_proto_goTypes = nil
	file_google_appengine_v1_domain_mapping_proto_depIdxs = nil
}
