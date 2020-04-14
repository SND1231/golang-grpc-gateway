// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/shared_set.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// SharedSets are used for sharing criterion exclusions across multiple
// campaigns.
type SharedSet struct {
	// The resource name of the shared set.
	// Shared set resource names have the form:
	//
	// `customers/{customer_id}/sharedSets/{shared_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of this shared set. Read only.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The type of this shared set: each shared set holds only a single kind
	// of resource. Required. Immutable.
	Type enums.SharedSetTypeEnum_SharedSetType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.SharedSetTypeEnum_SharedSetType" json:"type,omitempty"`
	// The name of this shared set. Required.
	// Shared Sets must have names that are unique among active shared sets of
	// the same type.
	// The length of this string should be between 1 and 255 UTF-8 bytes,
	// inclusive.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The status of this shared set. Read only.
	Status enums.SharedSetStatusEnum_SharedSetStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.SharedSetStatusEnum_SharedSetStatus" json:"status,omitempty"`
	// The number of shared criteria within this shared set. Read only.
	MemberCount *wrappers.Int64Value `protobuf:"bytes,6,opt,name=member_count,json=memberCount,proto3" json:"member_count,omitempty"`
	// The number of campaigns associated with this shared set. Read only.
	ReferenceCount       *wrappers.Int64Value `protobuf:"bytes,7,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SharedSet) Reset()         { *m = SharedSet{} }
func (m *SharedSet) String() string { return proto.CompactTextString(m) }
func (*SharedSet) ProtoMessage()    {}
func (*SharedSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a1581938e7cc60d, []int{0}
}

func (m *SharedSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSet.Unmarshal(m, b)
}
func (m *SharedSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSet.Marshal(b, m, deterministic)
}
func (m *SharedSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSet.Merge(m, src)
}
func (m *SharedSet) XXX_Size() int {
	return xxx_messageInfo_SharedSet.Size(m)
}
func (m *SharedSet) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSet.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSet proto.InternalMessageInfo

func (m *SharedSet) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SharedSet) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SharedSet) GetType() enums.SharedSetTypeEnum_SharedSetType {
	if m != nil {
		return m.Type
	}
	return enums.SharedSetTypeEnum_UNSPECIFIED
}

func (m *SharedSet) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SharedSet) GetStatus() enums.SharedSetStatusEnum_SharedSetStatus {
	if m != nil {
		return m.Status
	}
	return enums.SharedSetStatusEnum_UNSPECIFIED
}

func (m *SharedSet) GetMemberCount() *wrappers.Int64Value {
	if m != nil {
		return m.MemberCount
	}
	return nil
}

func (m *SharedSet) GetReferenceCount() *wrappers.Int64Value {
	if m != nil {
		return m.ReferenceCount
	}
	return nil
}

func init() {
	proto.RegisterType((*SharedSet)(nil), "google.ads.googleads.v1.resources.SharedSet")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/shared_set.proto", fileDescriptor_3a1581938e7cc60d)
}

var fileDescriptor_3a1581938e7cc60d = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0xd9, 0x4d, 0x8c, 0x74, 0x5a, 0x23, 0xec, 0x69, 0xa9, 0x45, 0x52, 0xa5, 0x10, 0x10,
	0x66, 0xdd, 0x54, 0x3d, 0x8c, 0x50, 0xd8, 0xa8, 0x14, 0x3d, 0x48, 0xd9, 0x48, 0x0e, 0x25, 0x10,
	0x26, 0x99, 0xd7, 0x75, 0x21, 0x3b, 0xb3, 0xcc, 0xcc, 0x56, 0xfa, 0x75, 0x3c, 0xfa, 0x3d, 0xbc,
	0xf8, 0x51, 0xfc, 0x0a, 0x5e, 0x64, 0x67, 0x76, 0x06, 0x8b, 0xb4, 0xc9, 0xed, 0xe5, 0xbd, 0xff,
	0xff, 0x97, 0xff, 0x7b, 0xcc, 0xa2, 0x49, 0x21, 0x44, 0xb1, 0x81, 0x84, 0x32, 0x95, 0xd8, 0xb2,
	0xad, 0xae, 0xd3, 0x44, 0x82, 0x12, 0x8d, 0x5c, 0x83, 0x4a, 0xd4, 0x57, 0x2a, 0x81, 0x2d, 0x15,
	0x68, 0x5c, 0x4b, 0xa1, 0x45, 0x74, 0x6c, 0x85, 0x98, 0x32, 0x85, 0xbd, 0x07, 0x5f, 0xa7, 0xd8,
	0x7b, 0x0e, 0x5f, 0xdf, 0x85, 0x05, 0xde, 0x54, 0xff, 0x22, 0x97, 0x4a, 0x53, 0xdd, 0x28, 0x4b,
	0x3e, 0x3c, 0xdd, 0xd9, 0xa6, 0x6f, 0x6a, 0xe8, 0x4c, 0x47, 0xce, 0x54, 0x97, 0x09, 0xe5, 0x5c,
	0x68, 0xaa, 0x4b, 0xc1, 0x1d, 0xf2, 0x69, 0x37, 0x35, 0xbf, 0x56, 0xcd, 0x55, 0xf2, 0x4d, 0xd2,
	0xba, 0x06, 0xd9, 0xcd, 0x9f, 0xfd, 0xec, 0xa1, 0xbd, 0x99, 0xe1, 0xce, 0x40, 0x47, 0xcf, 0xd1,
	0x23, 0xb7, 0xc4, 0x92, 0xd3, 0x0a, 0xe2, 0x60, 0x14, 0x8c, 0xf7, 0xf2, 0x03, 0xd7, 0xfc, 0x4c,
	0x2b, 0x88, 0x5e, 0xa0, 0xb0, 0x64, 0x71, 0x38, 0x0a, 0xc6, 0xfb, 0x93, 0x27, 0xdd, 0x05, 0xb0,
	0xe3, 0xe3, 0x8f, 0x5c, 0xbf, 0x79, 0x35, 0xa7, 0x9b, 0x06, 0xf2, 0xb0, 0x64, 0x51, 0x8e, 0xfa,
	0x6d, 0xd6, 0xb8, 0x37, 0x0a, 0xc6, 0xc3, 0xc9, 0x19, 0xbe, 0xeb, 0x76, 0x66, 0x43, 0xec, 0x93,
	0x7c, 0xb9, 0xa9, 0xe1, 0x03, 0x6f, 0xaa, 0xdb, 0x9d, 0xdc, 0xb0, 0xa2, 0x97, 0xa8, 0x6f, 0xc2,
	0xf5, 0x4d, 0x84, 0xa3, 0xff, 0x22, 0xcc, 0xb4, 0x2c, 0x79, 0x61, 0x33, 0x18, 0x65, 0x74, 0x89,
	0x06, 0xf6, 0xd0, 0xf1, 0x03, 0x93, 0x63, 0xba, 0x6b, 0x8e, 0x99, 0x71, 0xdd, 0x4e, 0x62, 0x7b,
	0x79, 0x47, 0x8c, 0xce, 0xd0, 0x41, 0x05, 0xd5, 0x0a, 0xe4, 0x72, 0x2d, 0x1a, 0xae, 0xe3, 0xc1,
	0xf6, 0xc3, 0xec, 0x5b, 0xc3, 0xbb, 0x56, 0x1f, 0xbd, 0x47, 0x8f, 0x25, 0x5c, 0x81, 0x04, 0xbe,
	0x86, 0x0e, 0xf1, 0x70, 0x3b, 0x62, 0xe8, 0x3d, 0x86, 0x32, 0xfd, 0x13, 0xa0, 0x93, 0xb5, 0xa8,
	0xf0, 0xd6, 0xb7, 0x39, 0x1d, 0xfa, 0x45, 0x2e, 0x5a, 0xee, 0x45, 0x70, 0xf9, 0xa9, 0x33, 0x15,
	0x62, 0x43, 0x79, 0x81, 0x85, 0x2c, 0x92, 0x02, 0xb8, 0xf9, 0x57, 0xf7, 0x0c, 0xeb, 0x52, 0xdd,
	0xf3, 0x8d, 0xbc, 0xf5, 0xd5, 0xf7, 0xb0, 0x77, 0x9e, 0x65, 0x3f, 0xc2, 0xe3, 0x73, 0x8b, 0xcc,
	0x98, 0xc2, 0xb6, 0x6c, 0xab, 0x79, 0x8a, 0x73, 0xa7, 0xfc, 0xe5, 0x34, 0x8b, 0x8c, 0xa9, 0x85,
	0xd7, 0x2c, 0xe6, 0xe9, 0xc2, 0x6b, 0x7e, 0x87, 0x27, 0x76, 0x40, 0x48, 0xc6, 0x14, 0x21, 0x5e,
	0x45, 0xc8, 0x3c, 0x25, 0xc4, 0xeb, 0x56, 0x03, 0x13, 0xf6, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x08, 0xe0, 0x65, 0x49, 0xcf, 0x03, 0x00, 0x00,
}
