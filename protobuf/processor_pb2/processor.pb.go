// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/processor_pb2/processor.proto

package processor_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import transaction_pb2 "github.com/daludaluking/sawtooth-sdk/protobuf/transaction_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TpRegisterResponse_Status int32

const (
	TpRegisterResponse_STATUS_UNSET TpRegisterResponse_Status = 0
	TpRegisterResponse_OK           TpRegisterResponse_Status = 1
	TpRegisterResponse_ERROR        TpRegisterResponse_Status = 2
)

var TpRegisterResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "ERROR",
}
var TpRegisterResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"ERROR":        2,
}

func (x TpRegisterResponse_Status) String() string {
	return proto.EnumName(TpRegisterResponse_Status_name, int32(x))
}
func (TpRegisterResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_processor_f8a91421cdbbfbbf, []int{1, 0}
}

type TpUnregisterResponse_Status int32

const (
	TpUnregisterResponse_STATUS_UNSET TpUnregisterResponse_Status = 0
	TpUnregisterResponse_OK           TpUnregisterResponse_Status = 1
	TpUnregisterResponse_ERROR        TpUnregisterResponse_Status = 2
)

var TpUnregisterResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "ERROR",
}
var TpUnregisterResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"ERROR":        2,
}

func (x TpUnregisterResponse_Status) String() string {
	return proto.EnumName(TpUnregisterResponse_Status_name, int32(x))
}
func (TpUnregisterResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_processor_f8a91421cdbbfbbf, []int{3, 0}
}

type TpProcessResponse_Status int32

const (
	TpProcessResponse_STATUS_UNSET        TpProcessResponse_Status = 0
	TpProcessResponse_OK                  TpProcessResponse_Status = 1
	TpProcessResponse_INVALID_TRANSACTION TpProcessResponse_Status = 2
	TpProcessResponse_INTERNAL_ERROR      TpProcessResponse_Status = 3
)

var TpProcessResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INVALID_TRANSACTION",
	3: "INTERNAL_ERROR",
}
var TpProcessResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"INVALID_TRANSACTION": 2,
	"INTERNAL_ERROR":      3,
}

func (x TpProcessResponse_Status) String() string {
	return proto.EnumName(TpProcessResponse_Status_name, int32(x))
}
func (TpProcessResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_processor_f8a91421cdbbfbbf, []int{5, 0}
}

// The registration request from the transaction processor to the
// validator/executor
type TpRegisterRequest struct {
	// A settled upon name for the capabilities of the transaction processor.
	// For example: intkey, xo
	Family string `protobuf:"bytes,1,opt,name=family" json:"family,omitempty"`
	// The version supported.  For example:
	//      1.0  for version 1.0
	//      2.1  for version 2.1
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	// The namespaces this transaction processor expects to interact with
	// when processing transactions matching this specification; will be
	// enforced by the state API on the validator.
	Namespaces []string `protobuf:"bytes,4,rep,name=namespaces" json:"namespaces,omitempty"`
	// The maximum number of transactions that this transaction processor can
	// handle at once.
	MaxOccupancy         uint32   `protobuf:"varint,5,opt,name=max_occupancy,json=maxOccupancy" json:"max_occupancy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TpRegisterRequest) Reset()         { *m = TpRegisterRequest{} }
func (m *TpRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*TpRegisterRequest) ProtoMessage()    {}
func (*TpRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_processor_f8a91421cdbbfbbf, []int{0}
}
func (m *TpRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpRegisterRequest.Unmarshal(m, b)
}
func (m *TpRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpRegisterRequest.Marshal(b, m, deterministic)
}
func (dst *TpRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpRegisterRequest.Merge(dst, src)
}
func (m *TpRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_TpRegisterRequest.Size(m)
}
func (m *TpRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TpRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TpRegisterRequest proto.InternalMessageInfo

func (m *TpRegisterRequest) GetFamily() string {
	if m != nil {
		return m.Family
	}
	return ""
}

func (m *TpRegisterRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *TpRegisterRequest) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *TpRegisterRequest) GetMaxOccupancy() uint32 {
	if m != nil {
		return m.MaxOccupancy
	}
	return 0
}

// A response sent from the validator to the transaction processor
// acknowledging the registration
type TpRegisterResponse struct {
	Status               TpRegisterResponse_Status `protobuf:"varint,1,opt,name=status,enum=TpRegisterResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TpRegisterResponse) Reset()         { *m = TpRegisterResponse{} }
func (m *TpRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*TpRegisterResponse) ProtoMessage()    {}
func (*TpRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_processor_f8a91421cdbbfbbf, []int{1}
}
func (m *TpRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpRegisterResponse.Unmarshal(m, b)
}
func (m *TpRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpRegisterResponse.Marshal(b, m, deterministic)
}
func (dst *TpRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpRegisterResponse.Merge(dst, src)
}
func (m *TpRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_TpRegisterResponse.Size(m)
}
func (m *TpRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TpRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TpRegisterResponse proto.InternalMessageInfo

func (m *TpRegisterResponse) GetStatus() TpRegisterResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpRegisterResponse_STATUS_UNSET
}

// The unregistration request from the transaction processor to the
// validator/executor. The correct handlers are determined from the
// zeromq identity of the tp, on the validator side.
type TpUnregisterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TpUnregisterRequest) Reset()         { *m = TpUnregisterRequest{} }
func (m *TpUnregisterRequest) String() string { return proto.CompactTextString(m) }
func (*TpUnregisterRequest) ProtoMessage()    {}
func (*TpUnregisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_processor_f8a91421cdbbfbbf, []int{2}
}
func (m *TpUnregisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpUnregisterRequest.Unmarshal(m, b)
}
func (m *TpUnregisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpUnregisterRequest.Marshal(b, m, deterministic)
}
func (dst *TpUnregisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpUnregisterRequest.Merge(dst, src)
}
func (m *TpUnregisterRequest) XXX_Size() int {
	return xxx_messageInfo_TpUnregisterRequest.Size(m)
}
func (m *TpUnregisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TpUnregisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TpUnregisterRequest proto.InternalMessageInfo

// A response sent from the validator to the transaction processor
// acknowledging the unregistration
type TpUnregisterResponse struct {
	Status               TpUnregisterResponse_Status `protobuf:"varint,1,opt,name=status,enum=TpUnregisterResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TpUnregisterResponse) Reset()         { *m = TpUnregisterResponse{} }
func (m *TpUnregisterResponse) String() string { return proto.CompactTextString(m) }
func (*TpUnregisterResponse) ProtoMessage()    {}
func (*TpUnregisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_processor_f8a91421cdbbfbbf, []int{3}
}
func (m *TpUnregisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpUnregisterResponse.Unmarshal(m, b)
}
func (m *TpUnregisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpUnregisterResponse.Marshal(b, m, deterministic)
}
func (dst *TpUnregisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpUnregisterResponse.Merge(dst, src)
}
func (m *TpUnregisterResponse) XXX_Size() int {
	return xxx_messageInfo_TpUnregisterResponse.Size(m)
}
func (m *TpUnregisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TpUnregisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TpUnregisterResponse proto.InternalMessageInfo

func (m *TpUnregisterResponse) GetStatus() TpUnregisterResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpUnregisterResponse_STATUS_UNSET
}

// The request from the validator/executor of the transaction processor
// to verify a transaction.
type TpProcessRequest struct {
	Header               *transaction_pb2.TransactionHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Payload              []byte                             `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature            string                             `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
	ContextId            string                             `protobuf:"bytes,4,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *TpProcessRequest) Reset()         { *m = TpProcessRequest{} }
func (m *TpProcessRequest) String() string { return proto.CompactTextString(m) }
func (*TpProcessRequest) ProtoMessage()    {}
func (*TpProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_processor_f8a91421cdbbfbbf, []int{4}
}
func (m *TpProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpProcessRequest.Unmarshal(m, b)
}
func (m *TpProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpProcessRequest.Marshal(b, m, deterministic)
}
func (dst *TpProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpProcessRequest.Merge(dst, src)
}
func (m *TpProcessRequest) XXX_Size() int {
	return xxx_messageInfo_TpProcessRequest.Size(m)
}
func (m *TpProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TpProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TpProcessRequest proto.InternalMessageInfo

func (m *TpProcessRequest) GetHeader() *transaction_pb2.TransactionHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TpProcessRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TpProcessRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *TpProcessRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

// The response from the transaction processor to the validator/executor
// used to respond about the validity of a transaction
type TpProcessResponse struct {
	Status TpProcessResponse_Status `protobuf:"varint,1,opt,name=status,enum=TpProcessResponse_Status" json:"status,omitempty"`
	// A message to include on responses in the cases where
	// status is either INVALID_TRANSACTION or INTERNAL_ERROR
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// Information that may be included with the response.
	// This information is an opaque, application-specific encoded block of
	// data that will be propagated back to the transaction submitter.
	ExtendedData         []byte   `protobuf:"bytes,3,opt,name=extended_data,json=extendedData,proto3" json:"extended_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TpProcessResponse) Reset()         { *m = TpProcessResponse{} }
func (m *TpProcessResponse) String() string { return proto.CompactTextString(m) }
func (*TpProcessResponse) ProtoMessage()    {}
func (*TpProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_processor_f8a91421cdbbfbbf, []int{5}
}
func (m *TpProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpProcessResponse.Unmarshal(m, b)
}
func (m *TpProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpProcessResponse.Marshal(b, m, deterministic)
}
func (dst *TpProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpProcessResponse.Merge(dst, src)
}
func (m *TpProcessResponse) XXX_Size() int {
	return xxx_messageInfo_TpProcessResponse.Size(m)
}
func (m *TpProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TpProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TpProcessResponse proto.InternalMessageInfo

func (m *TpProcessResponse) GetStatus() TpProcessResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpProcessResponse_STATUS_UNSET
}

func (m *TpProcessResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TpProcessResponse) GetExtendedData() []byte {
	if m != nil {
		return m.ExtendedData
	}
	return nil
}

func init() {
	proto.RegisterType((*TpRegisterRequest)(nil), "TpRegisterRequest")
	proto.RegisterType((*TpRegisterResponse)(nil), "TpRegisterResponse")
	proto.RegisterType((*TpUnregisterRequest)(nil), "TpUnregisterRequest")
	proto.RegisterType((*TpUnregisterResponse)(nil), "TpUnregisterResponse")
	proto.RegisterType((*TpProcessRequest)(nil), "TpProcessRequest")
	proto.RegisterType((*TpProcessResponse)(nil), "TpProcessResponse")
	proto.RegisterEnum("TpRegisterResponse_Status", TpRegisterResponse_Status_name, TpRegisterResponse_Status_value)
	proto.RegisterEnum("TpUnregisterResponse_Status", TpUnregisterResponse_Status_name, TpUnregisterResponse_Status_value)
	proto.RegisterEnum("TpProcessResponse_Status", TpProcessResponse_Status_name, TpProcessResponse_Status_value)
}

func init() {
	proto.RegisterFile("github.com/daludaluking/sawtooth-sdk/protobuf/processor_pb2/processor.proto", fileDescriptor_processor_f8a91421cdbbfbbf)
}

var fileDescriptor_processor_f8a91421cdbbfbbf = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x14, 0xfc, 0x39, 0x69, 0xfd, 0x53, 0x1e, 0x49, 0x65, 0xb6, 0x14, 0x4c, 0x55, 0x50, 0x64, 0x2e,
	0x11, 0x12, 0x46, 0xa4, 0x48, 0x9c, 0x0d, 0x8d, 0x84, 0x45, 0x65, 0x57, 0xeb, 0x0d, 0x07, 0x2e,
	0xd6, 0xc6, 0xde, 0xa6, 0x56, 0xeb, 0x5d, 0xe3, 0x5d, 0xd3, 0x44, 0x7c, 0x03, 0xee, 0x7c, 0x35,
	0x3e, 0x0f, 0xca, 0xda, 0x26, 0x7f, 0xda, 0x03, 0x12, 0x37, 0xcf, 0xbc, 0x59, 0x79, 0x34, 0xf3,
	0x1e, 0x9c, 0x4a, 0x7a, 0xab, 0x84, 0x50, 0x57, 0xb1, 0x4c, 0xaf, 0x5f, 0x17, 0xa5, 0x50, 0x62,
	0x56, 0x5d, 0xae, 0x3e, 0x12, 0x26, 0xa5, 0x28, 0xe3, 0x62, 0x36, 0x5e, 0x23, 0x57, 0x0b, 0x8e,
	0xdf, 0xdd, 0xff, 0x48, 0x95, 0x94, 0x4b, 0x9a, 0xa8, 0x4c, 0x70, 0xfd, 0x6c, 0x03, 0xd7, 0x0f,
	0x9d, 0x1f, 0x06, 0x3c, 0x24, 0x05, 0x66, 0xf3, 0x4c, 0x2a, 0x56, 0x62, 0xf6, 0xb5, 0x62, 0x52,
	0xa1, 0xc7, 0x60, 0x5e, 0xd2, 0x3c, 0xbb, 0x59, 0xda, 0xc6, 0xd0, 0x18, 0xf5, 0x70, 0x83, 0x90,
	0x0d, 0xff, 0x7f, 0x63, 0xa5, 0xcc, 0x04, 0xb7, 0x3b, 0x7a, 0xd0, 0x42, 0xf4, 0x1c, 0x80, 0xd3,
	0x9c, 0xc9, 0x82, 0x26, 0x4c, 0xda, 0x7b, 0xc3, 0xee, 0xa8, 0x87, 0x37, 0x18, 0xf4, 0x02, 0x06,
	0x39, 0x5d, 0xc4, 0x22, 0x49, 0xaa, 0x82, 0xf2, 0x64, 0x69, 0xef, 0x0f, 0x8d, 0xd1, 0x00, 0xf7,
	0x73, 0xba, 0x08, 0x5b, 0xce, 0xb9, 0x05, 0xb4, 0xe9, 0x45, 0x16, 0x82, 0x4b, 0x86, 0xc6, 0x60,
	0x4a, 0x45, 0x55, 0x25, 0xb5, 0x99, 0x83, 0xf1, 0xb1, 0x7b, 0x57, 0xe4, 0x46, 0x5a, 0x81, 0x1b,
	0xa5, 0xf3, 0x0a, 0xcc, 0x9a, 0x41, 0x16, 0xf4, 0x23, 0xe2, 0x91, 0x69, 0x14, 0x4f, 0x83, 0x68,
	0x42, 0xac, 0xff, 0x90, 0x09, 0x9d, 0xf0, 0x93, 0x65, 0xa0, 0x1e, 0xec, 0x4f, 0x30, 0x0e, 0xb1,
	0xd5, 0x71, 0x8e, 0xe0, 0x90, 0x14, 0x53, 0x5e, 0x6e, 0xc7, 0xe0, 0x7c, 0x87, 0x47, 0xdb, 0x74,
	0xe3, 0xe8, 0xed, 0x8e, 0xa3, 0x13, 0xf7, 0x3e, 0xd9, 0x3f, 0x7a, 0xfa, 0x69, 0x80, 0x45, 0x8a,
	0x8b, 0xba, 0xe8, 0xb6, 0x98, 0x97, 0x60, 0x5e, 0x31, 0x9a, 0xb2, 0x52, 0xff, 0xf9, 0xc1, 0x18,
	0xb9, 0x64, 0x5d, 0xe9, 0x47, 0x3d, 0xc1, 0x8d, 0x62, 0x55, 0x56, 0x41, 0x97, 0x37, 0x82, 0xa6,
	0xba, 0xac, 0x3e, 0x6e, 0x21, 0x3a, 0x81, 0x9e, 0xcc, 0xe6, 0x9c, 0xaa, 0xaa, 0x64, 0x76, 0x57,
	0x17, 0xb9, 0x26, 0xd0, 0x33, 0x80, 0x44, 0x70, 0xc5, 0x16, 0x2a, 0xce, 0x52, 0x7b, 0xaf, 0x1e,
	0x37, 0x8c, 0x9f, 0x3a, 0xbf, 0xf4, 0xc6, 0xfc, 0xf1, 0xd5, 0x44, 0xf2, 0x66, 0x27, 0x92, 0xa7,
	0xee, 0x1d, 0xcd, 0x4e, 0x1e, 0x2b, 0x7f, 0x39, 0x93, 0x92, 0xce, 0x59, 0xbb, 0x4c, 0x0d, 0x5c,
	0x2d, 0x0b, 0x5b, 0x28, 0xc6, 0x53, 0x96, 0xc6, 0x29, 0x55, 0x54, 0x7b, 0xec, 0xe3, 0x7e, 0x4b,
	0x9e, 0x51, 0x45, 0x9d, 0xf0, 0x2f, 0xe2, 0x7c, 0x02, 0x87, 0x7e, 0xf0, 0xd9, 0x3b, 0xf7, 0xcf,
	0x62, 0x82, 0xbd, 0x20, 0xf2, 0x3e, 0x10, 0x3f, 0x0c, 0xac, 0x0e, 0x42, 0x70, 0xe0, 0x07, 0x64,
	0x82, 0x03, 0xef, 0x3c, 0xae, 0x03, 0xef, 0xbe, 0x1f, 0xc1, 0x51, 0x7b, 0x45, 0xae, 0x4c, 0xaf,
	0xdd, 0xf6, 0x8a, 0x2e, 0x8c, 0x2f, 0x83, 0xad, 0xeb, 0x9b, 0x99, 0x7a, 0x74, 0xfa, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0xd5, 0xb7, 0x20, 0x0b, 0xab, 0x03, 0x00, 0x00,
}
