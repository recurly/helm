// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_bid_modifier_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for
// [CampaignBidModifierService.GetCampaignBidModifier][google.ads.googleads.v1.services.CampaignBidModifierService.GetCampaignBidModifier].
type GetCampaignBidModifierRequest struct {
	// The resource name of the campaign bid modifier to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignBidModifierRequest) Reset()         { *m = GetCampaignBidModifierRequest{} }
func (m *GetCampaignBidModifierRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignBidModifierRequest) ProtoMessage()    {}
func (*GetCampaignBidModifierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d42c41001293b8, []int{0}
}

func (m *GetCampaignBidModifierRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignBidModifierRequest.Unmarshal(m, b)
}
func (m *GetCampaignBidModifierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignBidModifierRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignBidModifierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignBidModifierRequest.Merge(m, src)
}
func (m *GetCampaignBidModifierRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignBidModifierRequest.Size(m)
}
func (m *GetCampaignBidModifierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignBidModifierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignBidModifierRequest proto.InternalMessageInfo

func (m *GetCampaignBidModifierRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignBidModifierService.MutateCampaignBidModifier][].
type MutateCampaignBidModifiersRequest struct {
	// ID of the customer whose campaign bid modifiers are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual campaign bid modifiers.
	Operations []*CampaignBidModifierOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignBidModifiersRequest) Reset()         { *m = MutateCampaignBidModifiersRequest{} }
func (m *MutateCampaignBidModifiersRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignBidModifiersRequest) ProtoMessage()    {}
func (*MutateCampaignBidModifiersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d42c41001293b8, []int{1}
}

func (m *MutateCampaignBidModifiersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignBidModifiersRequest.Unmarshal(m, b)
}
func (m *MutateCampaignBidModifiersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignBidModifiersRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignBidModifiersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignBidModifiersRequest.Merge(m, src)
}
func (m *MutateCampaignBidModifiersRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignBidModifiersRequest.Size(m)
}
func (m *MutateCampaignBidModifiersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignBidModifiersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignBidModifiersRequest proto.InternalMessageInfo

func (m *MutateCampaignBidModifiersRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignBidModifiersRequest) GetOperations() []*CampaignBidModifierOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignBidModifiersRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignBidModifiersRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove, update) on a campaign bid modifier.
type CampaignBidModifierOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignBidModifierOperation_Create
	//	*CampaignBidModifierOperation_Update
	//	*CampaignBidModifierOperation_Remove
	Operation            isCampaignBidModifierOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *CampaignBidModifierOperation) Reset()         { *m = CampaignBidModifierOperation{} }
func (m *CampaignBidModifierOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignBidModifierOperation) ProtoMessage()    {}
func (*CampaignBidModifierOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d42c41001293b8, []int{2}
}

func (m *CampaignBidModifierOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignBidModifierOperation.Unmarshal(m, b)
}
func (m *CampaignBidModifierOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignBidModifierOperation.Marshal(b, m, deterministic)
}
func (m *CampaignBidModifierOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignBidModifierOperation.Merge(m, src)
}
func (m *CampaignBidModifierOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignBidModifierOperation.Size(m)
}
func (m *CampaignBidModifierOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignBidModifierOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignBidModifierOperation proto.InternalMessageInfo

func (m *CampaignBidModifierOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignBidModifierOperation_Operation interface {
	isCampaignBidModifierOperation_Operation()
}

type CampaignBidModifierOperation_Create struct {
	Create *resources.CampaignBidModifier `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignBidModifierOperation_Update struct {
	Update *resources.CampaignBidModifier `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignBidModifierOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignBidModifierOperation_Create) isCampaignBidModifierOperation_Operation() {}

func (*CampaignBidModifierOperation_Update) isCampaignBidModifierOperation_Operation() {}

func (*CampaignBidModifierOperation_Remove) isCampaignBidModifierOperation_Operation() {}

func (m *CampaignBidModifierOperation) GetOperation() isCampaignBidModifierOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignBidModifierOperation) GetCreate() *resources.CampaignBidModifier {
	if x, ok := m.GetOperation().(*CampaignBidModifierOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignBidModifierOperation) GetUpdate() *resources.CampaignBidModifier {
	if x, ok := m.GetOperation().(*CampaignBidModifierOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignBidModifierOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignBidModifierOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignBidModifierOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignBidModifierOperation_Create)(nil),
		(*CampaignBidModifierOperation_Update)(nil),
		(*CampaignBidModifierOperation_Remove)(nil),
	}
}

// Response message for campaign bid modifiers mutate.
type MutateCampaignBidModifiersResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignBidModifierResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *MutateCampaignBidModifiersResponse) Reset()         { *m = MutateCampaignBidModifiersResponse{} }
func (m *MutateCampaignBidModifiersResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignBidModifiersResponse) ProtoMessage()    {}
func (*MutateCampaignBidModifiersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d42c41001293b8, []int{3}
}

func (m *MutateCampaignBidModifiersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignBidModifiersResponse.Unmarshal(m, b)
}
func (m *MutateCampaignBidModifiersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignBidModifiersResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignBidModifiersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignBidModifiersResponse.Merge(m, src)
}
func (m *MutateCampaignBidModifiersResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignBidModifiersResponse.Size(m)
}
func (m *MutateCampaignBidModifiersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignBidModifiersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignBidModifiersResponse proto.InternalMessageInfo

func (m *MutateCampaignBidModifiersResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignBidModifiersResponse) GetResults() []*MutateCampaignBidModifierResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateCampaignBidModifierResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignBidModifierResult) Reset()         { *m = MutateCampaignBidModifierResult{} }
func (m *MutateCampaignBidModifierResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignBidModifierResult) ProtoMessage()    {}
func (*MutateCampaignBidModifierResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d42c41001293b8, []int{4}
}

func (m *MutateCampaignBidModifierResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignBidModifierResult.Unmarshal(m, b)
}
func (m *MutateCampaignBidModifierResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignBidModifierResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignBidModifierResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignBidModifierResult.Merge(m, src)
}
func (m *MutateCampaignBidModifierResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignBidModifierResult.Size(m)
}
func (m *MutateCampaignBidModifierResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignBidModifierResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignBidModifierResult proto.InternalMessageInfo

func (m *MutateCampaignBidModifierResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignBidModifierRequest)(nil), "google.ads.googleads.v1.services.GetCampaignBidModifierRequest")
	proto.RegisterType((*MutateCampaignBidModifiersRequest)(nil), "google.ads.googleads.v1.services.MutateCampaignBidModifiersRequest")
	proto.RegisterType((*CampaignBidModifierOperation)(nil), "google.ads.googleads.v1.services.CampaignBidModifierOperation")
	proto.RegisterType((*MutateCampaignBidModifiersResponse)(nil), "google.ads.googleads.v1.services.MutateCampaignBidModifiersResponse")
	proto.RegisterType((*MutateCampaignBidModifierResult)(nil), "google.ads.googleads.v1.services.MutateCampaignBidModifierResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_bid_modifier_service.proto", fileDescriptor_27d42c41001293b8)
}

var fileDescriptor_27d42c41001293b8 = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4f, 0x6e, 0xd3, 0x4c,
	0x14, 0xc0, 0x3f, 0x3b, 0x9f, 0x0a, 0x9d, 0x14, 0x90, 0x06, 0x01, 0x51, 0x54, 0x68, 0x30, 0x95,
	0xa8, 0xb2, 0xb0, 0x95, 0x20, 0x55, 0xe0, 0xaa, 0x45, 0x49, 0x43, 0x5a, 0x16, 0xa5, 0x95, 0x2b,
	0x75, 0x01, 0x11, 0xd6, 0xd4, 0x9e, 0x58, 0xa3, 0xda, 0x1e, 0x33, 0x33, 0x0e, 0xaa, 0xaa, 0x6e,
	0x58, 0x70, 0x01, 0x6e, 0xc0, 0x82, 0x05, 0x37, 0x81, 0x1d, 0xe2, 0x02, 0x2c, 0x58, 0x21, 0xb1,
	0xe1, 0x04, 0xc8, 0x1e, 0x4f, 0x68, 0xab, 0xb8, 0x41, 0x74, 0xf7, 0xfc, 0xe6, 0xf9, 0xf7, 0xfe,
	0xce, 0x1b, 0xd0, 0x0b, 0x28, 0x0d, 0x42, 0x6c, 0x21, 0x9f, 0x5b, 0x52, 0xcc, 0xa4, 0x51, 0xcb,
	0xe2, 0x98, 0x8d, 0x88, 0x87, 0xb9, 0xe5, 0xa1, 0x28, 0x41, 0x24, 0x88, 0xdd, 0x7d, 0xe2, 0xbb,
	0x11, 0xf5, 0xc9, 0x90, 0x60, 0xe6, 0x16, 0xc7, 0x66, 0xc2, 0xa8, 0xa0, 0xb0, 0x21, 0x7f, 0x35,
	0x91, 0xcf, 0xcd, 0x31, 0xc5, 0x1c, 0xb5, 0x4c, 0x45, 0xa9, 0xaf, 0x96, 0xf9, 0x61, 0x98, 0xd3,
	0x94, 0x95, 0x3a, 0x92, 0x0e, 0xea, 0xf3, 0xea, 0xf7, 0x84, 0x58, 0x28, 0x8e, 0xa9, 0x40, 0x82,
	0xd0, 0x98, 0x17, 0xa7, 0x85, 0x7b, 0x2b, 0xff, 0xda, 0x4f, 0x87, 0xd6, 0x90, 0xe0, 0xd0, 0x77,
	0x23, 0xc4, 0x0f, 0x0a, 0x8b, 0x3b, 0x67, 0x2d, 0x5e, 0x33, 0x94, 0x24, 0x98, 0x29, 0xc2, 0xad,
	0xe2, 0x9c, 0x25, 0x9e, 0xc5, 0x05, 0x12, 0x69, 0x71, 0x60, 0xf4, 0xc0, 0xed, 0x0d, 0x2c, 0xd6,
	0x8b, 0xd0, 0xba, 0xc4, 0xdf, 0x2a, 0x02, 0x73, 0xf0, 0xab, 0x14, 0x73, 0x01, 0xef, 0x81, 0x2b,
	0x2a, 0x05, 0x37, 0x46, 0x11, 0xae, 0x69, 0x0d, 0x6d, 0x69, 0xd6, 0x99, 0x53, 0xca, 0x67, 0x28,
	0xc2, 0xc6, 0x2f, 0x0d, 0xdc, 0xdd, 0x4a, 0x05, 0x12, 0x78, 0x02, 0x89, 0x2b, 0xd4, 0x02, 0xa8,
	0x7a, 0x29, 0x17, 0x34, 0xc2, 0xcc, 0x25, 0x7e, 0x01, 0x02, 0x4a, 0xf5, 0xd4, 0x87, 0x2f, 0x01,
	0xa0, 0x09, 0x66, 0x32, 0xf7, 0x9a, 0xde, 0xa8, 0x2c, 0x55, 0xdb, 0x6b, 0xe6, 0xb4, 0xda, 0x9b,
	0x13, 0x7c, 0x6e, 0x2b, 0x8c, 0x73, 0x82, 0x08, 0xef, 0x83, 0x6b, 0x09, 0x62, 0x82, 0xa0, 0xd0,
	0x1d, 0x22, 0x12, 0xa6, 0x0c, 0xd7, 0x2a, 0x0d, 0x6d, 0xe9, 0xb2, 0x73, 0xb5, 0x50, 0xf7, 0xa5,
	0x36, 0x4b, 0x7a, 0x84, 0x42, 0xe2, 0x23, 0x81, 0x5d, 0x1a, 0x87, 0x87, 0xb5, 0xff, 0x73, 0xb3,
	0x39, 0xa5, 0xdc, 0x8e, 0xc3, 0x43, 0xe3, 0x83, 0x0e, 0xe6, 0xcf, 0x73, 0x0d, 0x57, 0x40, 0x35,
	0x4d, 0x72, 0x46, 0xd6, 0xa9, 0x9c, 0x51, 0x6d, 0xd7, 0x55, 0x3e, 0xaa, 0x55, 0x66, 0x3f, 0x6b,
	0xe6, 0x16, 0xe2, 0x07, 0x0e, 0x90, 0xe6, 0x99, 0x0c, 0x77, 0xc0, 0x8c, 0xc7, 0x30, 0x12, 0xb2,
	0xe0, 0xd5, 0xf6, 0x72, 0x69, 0x1d, 0xc6, 0x13, 0x36, 0xa9, 0x10, 0x9b, 0xff, 0x39, 0x05, 0x27,
	0x23, 0x4a, 0x7e, 0x4d, 0xbf, 0x28, 0x51, 0x72, 0x60, 0x0d, 0xcc, 0x30, 0x1c, 0xd1, 0x91, 0x2c,
	0xe3, 0x6c, 0x76, 0x22, 0xbf, 0xbb, 0x55, 0x30, 0x3b, 0xae, 0xbb, 0xf1, 0x49, 0x03, 0xc6, 0x79,
	0xd3, 0xc1, 0x13, 0x1a, 0x73, 0x0c, 0xfb, 0xe0, 0xc6, 0x99, 0xee, 0xb8, 0x98, 0x31, 0xca, 0x72,
	0x78, 0xb5, 0x0d, 0x55, 0xb8, 0x2c, 0xf1, 0xcc, 0xdd, 0x7c, 0x86, 0x9d, 0xeb, 0xa7, 0xfb, 0xf6,
	0x24, 0x33, 0x87, 0x2f, 0xc0, 0x25, 0x86, 0x79, 0x1a, 0x0a, 0x35, 0x42, 0x9d, 0xe9, 0x23, 0x54,
	0x1a, 0x9e, 0x93, 0x93, 0x1c, 0x45, 0x34, 0xfa, 0x60, 0x61, 0x8a, 0xed, 0x5f, 0xdd, 0x98, 0xf6,
	0xb7, 0x0a, 0xa8, 0x4f, 0x40, 0xec, 0xca, 0x80, 0xe0, 0x17, 0x0d, 0xdc, 0x9c, 0x7c, 0x2f, 0xe1,
	0xe3, 0xe9, 0xd9, 0x9c, 0x7b, 0xa3, 0xeb, 0xff, 0xd8, 0x77, 0x63, 0xed, 0xcd, 0xd7, 0xef, 0xef,
	0xf4, 0x87, 0x70, 0x39, 0x5b, 0x6b, 0x47, 0xa7, 0x52, 0x5c, 0x55, 0x97, 0x98, 0x5b, 0xcd, 0xf1,
	0x9e, 0x3b, 0xd9, 0x64, 0xab, 0x79, 0x0c, 0x7f, 0x6a, 0xa0, 0x5e, 0x3e, 0x06, 0x70, 0xfd, 0x02,
	0x5d, 0x52, 0x2b, 0xa6, 0xde, 0xbb, 0x18, 0x44, 0x4e, 0xa2, 0xd1, 0xcb, 0x33, 0x5d, 0x33, 0x1e,
	0x65, 0x99, 0xfe, 0x49, 0xed, 0xe8, 0xc4, 0xf6, 0x5a, 0x6d, 0x1e, 0x4f, 0x4c, 0xd4, 0x8e, 0x72,
	0xbc, 0xad, 0x35, 0xbb, 0x6f, 0x75, 0xb0, 0xe8, 0xd1, 0x68, 0x6a, 0x44, 0xdd, 0x85, 0xf2, 0x41,
	0xd8, 0xc9, 0x96, 0xc4, 0x8e, 0xf6, 0x7c, 0xb3, 0x80, 0x04, 0x34, 0x44, 0x71, 0x60, 0x52, 0x16,
	0x58, 0x01, 0x8e, 0xf3, 0x15, 0xa2, 0x9e, 0x9b, 0x84, 0xf0, 0xf2, 0x57, 0x6e, 0x45, 0x09, 0xef,
	0xf5, 0xca, 0x46, 0xa7, 0xf3, 0x51, 0x6f, 0x6c, 0x48, 0x60, 0xc7, 0xe7, 0xa6, 0x14, 0x33, 0x69,
	0xaf, 0x65, 0x16, 0x8e, 0xf9, 0x67, 0x65, 0x32, 0xe8, 0xf8, 0x7c, 0x30, 0x36, 0x19, 0xec, 0xb5,
	0x06, 0xca, 0xe4, 0x87, 0xbe, 0x28, 0xf5, 0xb6, 0xdd, 0xf1, 0xb9, 0x6d, 0x8f, 0x8d, 0x6c, 0x7b,
	0xaf, 0x65, 0xdb, 0xca, 0x6c, 0x7f, 0x26, 0x8f, 0xf3, 0xc1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x6b, 0x80, 0xdc, 0x8d, 0x8c, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignBidModifierServiceClient is the client API for CampaignBidModifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignBidModifierServiceClient interface {
	// Returns the requested campaign bid modifier in full detail.
	GetCampaignBidModifier(ctx context.Context, in *GetCampaignBidModifierRequest, opts ...grpc.CallOption) (*resources.CampaignBidModifier, error)
	// Creates, updates, or removes campaign bid modifiers.
	// Operation statuses are returned.
	MutateCampaignBidModifiers(ctx context.Context, in *MutateCampaignBidModifiersRequest, opts ...grpc.CallOption) (*MutateCampaignBidModifiersResponse, error)
}

type campaignBidModifierServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignBidModifierServiceClient(cc *grpc.ClientConn) CampaignBidModifierServiceClient {
	return &campaignBidModifierServiceClient{cc}
}

func (c *campaignBidModifierServiceClient) GetCampaignBidModifier(ctx context.Context, in *GetCampaignBidModifierRequest, opts ...grpc.CallOption) (*resources.CampaignBidModifier, error) {
	out := new(resources.CampaignBidModifier)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignBidModifierService/GetCampaignBidModifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignBidModifierServiceClient) MutateCampaignBidModifiers(ctx context.Context, in *MutateCampaignBidModifiersRequest, opts ...grpc.CallOption) (*MutateCampaignBidModifiersResponse, error) {
	out := new(MutateCampaignBidModifiersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignBidModifierService/MutateCampaignBidModifiers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignBidModifierServiceServer is the server API for CampaignBidModifierService service.
type CampaignBidModifierServiceServer interface {
	// Returns the requested campaign bid modifier in full detail.
	GetCampaignBidModifier(context.Context, *GetCampaignBidModifierRequest) (*resources.CampaignBidModifier, error)
	// Creates, updates, or removes campaign bid modifiers.
	// Operation statuses are returned.
	MutateCampaignBidModifiers(context.Context, *MutateCampaignBidModifiersRequest) (*MutateCampaignBidModifiersResponse, error)
}

// UnimplementedCampaignBidModifierServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignBidModifierServiceServer struct {
}

func (*UnimplementedCampaignBidModifierServiceServer) GetCampaignBidModifier(ctx context.Context, req *GetCampaignBidModifierRequest) (*resources.CampaignBidModifier, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCampaignBidModifier not implemented")
}
func (*UnimplementedCampaignBidModifierServiceServer) MutateCampaignBidModifiers(ctx context.Context, req *MutateCampaignBidModifiersRequest) (*MutateCampaignBidModifiersResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignBidModifiers not implemented")
}

func RegisterCampaignBidModifierServiceServer(s *grpc.Server, srv CampaignBidModifierServiceServer) {
	s.RegisterService(&_CampaignBidModifierService_serviceDesc, srv)
}

func _CampaignBidModifierService_GetCampaignBidModifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignBidModifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignBidModifierServiceServer).GetCampaignBidModifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignBidModifierService/GetCampaignBidModifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignBidModifierServiceServer).GetCampaignBidModifier(ctx, req.(*GetCampaignBidModifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignBidModifierService_MutateCampaignBidModifiers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignBidModifiersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignBidModifierServiceServer).MutateCampaignBidModifiers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignBidModifierService/MutateCampaignBidModifiers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignBidModifierServiceServer).MutateCampaignBidModifiers(ctx, req.(*MutateCampaignBidModifiersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignBidModifierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignBidModifierService",
	HandlerType: (*CampaignBidModifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignBidModifier",
			Handler:    _CampaignBidModifierService_GetCampaignBidModifier_Handler,
		},
		{
			MethodName: "MutateCampaignBidModifiers",
			Handler:    _CampaignBidModifierService_MutateCampaignBidModifiers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_bid_modifier_service.proto",
}
