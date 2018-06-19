// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth-sdk/protobuf/state_context_pb2/state_context.proto

package state_context_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import events_pb2 "github.com/daludaluking/sawtooth-sdk/protobuf/events_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TpStateGetResponse_Status int32

const (
	TpStateGetResponse_STATUS_UNSET        TpStateGetResponse_Status = 0
	TpStateGetResponse_OK                  TpStateGetResponse_Status = 1
	TpStateGetResponse_AUTHORIZATION_ERROR TpStateGetResponse_Status = 2
)

var TpStateGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "AUTHORIZATION_ERROR",
}
var TpStateGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"AUTHORIZATION_ERROR": 2,
}

func (x TpStateGetResponse_Status) String() string {
	return proto.EnumName(TpStateGetResponse_Status_name, int32(x))
}
func (TpStateGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{2, 0}
}

type TpStateSetResponse_Status int32

const (
	TpStateSetResponse_STATUS_UNSET        TpStateSetResponse_Status = 0
	TpStateSetResponse_OK                  TpStateSetResponse_Status = 1
	TpStateSetResponse_AUTHORIZATION_ERROR TpStateSetResponse_Status = 2
)

var TpStateSetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "AUTHORIZATION_ERROR",
}
var TpStateSetResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"AUTHORIZATION_ERROR": 2,
}

func (x TpStateSetResponse_Status) String() string {
	return proto.EnumName(TpStateSetResponse_Status_name, int32(x))
}
func (TpStateSetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{4, 0}
}

type TpStateDeleteResponse_Status int32

const (
	TpStateDeleteResponse_STATUS_UNSET        TpStateDeleteResponse_Status = 0
	TpStateDeleteResponse_OK                  TpStateDeleteResponse_Status = 1
	TpStateDeleteResponse_AUTHORIZATION_ERROR TpStateDeleteResponse_Status = 2
)

var TpStateDeleteResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "AUTHORIZATION_ERROR",
}
var TpStateDeleteResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"AUTHORIZATION_ERROR": 2,
}

func (x TpStateDeleteResponse_Status) String() string {
	return proto.EnumName(TpStateDeleteResponse_Status_name, int32(x))
}
func (TpStateDeleteResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{6, 0}
}

type TpReceiptAddDataResponse_Status int32

const (
	TpReceiptAddDataResponse_STATUS_UNSET TpReceiptAddDataResponse_Status = 0
	TpReceiptAddDataResponse_OK           TpReceiptAddDataResponse_Status = 1
	TpReceiptAddDataResponse_ERROR        TpReceiptAddDataResponse_Status = 2
)

var TpReceiptAddDataResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "ERROR",
}
var TpReceiptAddDataResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"ERROR":        2,
}

func (x TpReceiptAddDataResponse_Status) String() string {
	return proto.EnumName(TpReceiptAddDataResponse_Status_name, int32(x))
}
func (TpReceiptAddDataResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{8, 0}
}

type TpEventAddResponse_Status int32

const (
	TpEventAddResponse_STATUS_UNSET TpEventAddResponse_Status = 0
	TpEventAddResponse_OK           TpEventAddResponse_Status = 1
	TpEventAddResponse_ERROR        TpEventAddResponse_Status = 2
)

var TpEventAddResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "ERROR",
}
var TpEventAddResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"ERROR":        2,
}

func (x TpEventAddResponse_Status) String() string {
	return proto.EnumName(TpEventAddResponse_Status_name, int32(x))
}
func (TpEventAddResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{10, 0}
}

// An entry in the State
type TpStateEntry struct {
	Address              string   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TpStateEntry) Reset()         { *m = TpStateEntry{} }
func (m *TpStateEntry) String() string { return proto.CompactTextString(m) }
func (*TpStateEntry) ProtoMessage()    {}
func (*TpStateEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{0}
}
func (m *TpStateEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpStateEntry.Unmarshal(m, b)
}
func (m *TpStateEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpStateEntry.Marshal(b, m, deterministic)
}
func (dst *TpStateEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpStateEntry.Merge(dst, src)
}
func (m *TpStateEntry) XXX_Size() int {
	return xxx_messageInfo_TpStateEntry.Size(m)
}
func (m *TpStateEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TpStateEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TpStateEntry proto.InternalMessageInfo

func (m *TpStateEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *TpStateEntry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// A request from a handler/tp for the values at a series of addresses
type TpStateGetRequest struct {
	// The context id that references a context in the contextmanager
	ContextId            string   `protobuf:"bytes,1,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	Addresses            []string `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TpStateGetRequest) Reset()         { *m = TpStateGetRequest{} }
func (m *TpStateGetRequest) String() string { return proto.CompactTextString(m) }
func (*TpStateGetRequest) ProtoMessage()    {}
func (*TpStateGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{1}
}
func (m *TpStateGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpStateGetRequest.Unmarshal(m, b)
}
func (m *TpStateGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpStateGetRequest.Marshal(b, m, deterministic)
}
func (dst *TpStateGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpStateGetRequest.Merge(dst, src)
}
func (m *TpStateGetRequest) XXX_Size() int {
	return xxx_messageInfo_TpStateGetRequest.Size(m)
}
func (m *TpStateGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TpStateGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TpStateGetRequest proto.InternalMessageInfo

func (m *TpStateGetRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *TpStateGetRequest) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// A response from the contextmanager/validator with a series of State entries
type TpStateGetResponse struct {
	Entries              []*TpStateEntry           `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	Status               TpStateGetResponse_Status `protobuf:"varint,2,opt,name=status,enum=TpStateGetResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TpStateGetResponse) Reset()         { *m = TpStateGetResponse{} }
func (m *TpStateGetResponse) String() string { return proto.CompactTextString(m) }
func (*TpStateGetResponse) ProtoMessage()    {}
func (*TpStateGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{2}
}
func (m *TpStateGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpStateGetResponse.Unmarshal(m, b)
}
func (m *TpStateGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpStateGetResponse.Marshal(b, m, deterministic)
}
func (dst *TpStateGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpStateGetResponse.Merge(dst, src)
}
func (m *TpStateGetResponse) XXX_Size() int {
	return xxx_messageInfo_TpStateGetResponse.Size(m)
}
func (m *TpStateGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TpStateGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TpStateGetResponse proto.InternalMessageInfo

func (m *TpStateGetResponse) GetEntries() []*TpStateEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *TpStateGetResponse) GetStatus() TpStateGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpStateGetResponse_STATUS_UNSET
}

// A request from the handler/tp to put entries in the state of a context
type TpStateSetRequest struct {
	ContextId            string          `protobuf:"bytes,1,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	Entries              []*TpStateEntry `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TpStateSetRequest) Reset()         { *m = TpStateSetRequest{} }
func (m *TpStateSetRequest) String() string { return proto.CompactTextString(m) }
func (*TpStateSetRequest) ProtoMessage()    {}
func (*TpStateSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{3}
}
func (m *TpStateSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpStateSetRequest.Unmarshal(m, b)
}
func (m *TpStateSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpStateSetRequest.Marshal(b, m, deterministic)
}
func (dst *TpStateSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpStateSetRequest.Merge(dst, src)
}
func (m *TpStateSetRequest) XXX_Size() int {
	return xxx_messageInfo_TpStateSetRequest.Size(m)
}
func (m *TpStateSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TpStateSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TpStateSetRequest proto.InternalMessageInfo

func (m *TpStateSetRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *TpStateSetRequest) GetEntries() []*TpStateEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// A response from the contextmanager/validator with the addresses that were set
type TpStateSetResponse struct {
	Addresses            []string                  `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	Status               TpStateSetResponse_Status `protobuf:"varint,2,opt,name=status,enum=TpStateSetResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TpStateSetResponse) Reset()         { *m = TpStateSetResponse{} }
func (m *TpStateSetResponse) String() string { return proto.CompactTextString(m) }
func (*TpStateSetResponse) ProtoMessage()    {}
func (*TpStateSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{4}
}
func (m *TpStateSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpStateSetResponse.Unmarshal(m, b)
}
func (m *TpStateSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpStateSetResponse.Marshal(b, m, deterministic)
}
func (dst *TpStateSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpStateSetResponse.Merge(dst, src)
}
func (m *TpStateSetResponse) XXX_Size() int {
	return xxx_messageInfo_TpStateSetResponse.Size(m)
}
func (m *TpStateSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TpStateSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TpStateSetResponse proto.InternalMessageInfo

func (m *TpStateSetResponse) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *TpStateSetResponse) GetStatus() TpStateSetResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpStateSetResponse_STATUS_UNSET
}

// A request from the handler/tp to delete state entries at an collection of addresses
type TpStateDeleteRequest struct {
	ContextId            string   `protobuf:"bytes,1,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	Addresses            []string `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TpStateDeleteRequest) Reset()         { *m = TpStateDeleteRequest{} }
func (m *TpStateDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*TpStateDeleteRequest) ProtoMessage()    {}
func (*TpStateDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{5}
}
func (m *TpStateDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpStateDeleteRequest.Unmarshal(m, b)
}
func (m *TpStateDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpStateDeleteRequest.Marshal(b, m, deterministic)
}
func (dst *TpStateDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpStateDeleteRequest.Merge(dst, src)
}
func (m *TpStateDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_TpStateDeleteRequest.Size(m)
}
func (m *TpStateDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TpStateDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TpStateDeleteRequest proto.InternalMessageInfo

func (m *TpStateDeleteRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *TpStateDeleteRequest) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

// A response form the contextmanager/validator with the addresses that were deleted
type TpStateDeleteResponse struct {
	Addresses            []string                     `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	Status               TpStateDeleteResponse_Status `protobuf:"varint,2,opt,name=status,enum=TpStateDeleteResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *TpStateDeleteResponse) Reset()         { *m = TpStateDeleteResponse{} }
func (m *TpStateDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*TpStateDeleteResponse) ProtoMessage()    {}
func (*TpStateDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{6}
}
func (m *TpStateDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpStateDeleteResponse.Unmarshal(m, b)
}
func (m *TpStateDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpStateDeleteResponse.Marshal(b, m, deterministic)
}
func (dst *TpStateDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpStateDeleteResponse.Merge(dst, src)
}
func (m *TpStateDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_TpStateDeleteResponse.Size(m)
}
func (m *TpStateDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TpStateDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TpStateDeleteResponse proto.InternalMessageInfo

func (m *TpStateDeleteResponse) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *TpStateDeleteResponse) GetStatus() TpStateDeleteResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpStateDeleteResponse_STATUS_UNSET
}

// The request from the transaction processor to the validator append data
// to a transaction receipt
type TpReceiptAddDataRequest struct {
	// The context id that references a context in the context manager
	ContextId            string   `protobuf:"bytes,1,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TpReceiptAddDataRequest) Reset()         { *m = TpReceiptAddDataRequest{} }
func (m *TpReceiptAddDataRequest) String() string { return proto.CompactTextString(m) }
func (*TpReceiptAddDataRequest) ProtoMessage()    {}
func (*TpReceiptAddDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{7}
}
func (m *TpReceiptAddDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpReceiptAddDataRequest.Unmarshal(m, b)
}
func (m *TpReceiptAddDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpReceiptAddDataRequest.Marshal(b, m, deterministic)
}
func (dst *TpReceiptAddDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpReceiptAddDataRequest.Merge(dst, src)
}
func (m *TpReceiptAddDataRequest) XXX_Size() int {
	return xxx_messageInfo_TpReceiptAddDataRequest.Size(m)
}
func (m *TpReceiptAddDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TpReceiptAddDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TpReceiptAddDataRequest proto.InternalMessageInfo

func (m *TpReceiptAddDataRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *TpReceiptAddDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The response from the validator to the transaction processor to verify that
// data has been appended to a transaction receipt
type TpReceiptAddDataResponse struct {
	Status               TpReceiptAddDataResponse_Status `protobuf:"varint,2,opt,name=status,enum=TpReceiptAddDataResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *TpReceiptAddDataResponse) Reset()         { *m = TpReceiptAddDataResponse{} }
func (m *TpReceiptAddDataResponse) String() string { return proto.CompactTextString(m) }
func (*TpReceiptAddDataResponse) ProtoMessage()    {}
func (*TpReceiptAddDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{8}
}
func (m *TpReceiptAddDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpReceiptAddDataResponse.Unmarshal(m, b)
}
func (m *TpReceiptAddDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpReceiptAddDataResponse.Marshal(b, m, deterministic)
}
func (dst *TpReceiptAddDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpReceiptAddDataResponse.Merge(dst, src)
}
func (m *TpReceiptAddDataResponse) XXX_Size() int {
	return xxx_messageInfo_TpReceiptAddDataResponse.Size(m)
}
func (m *TpReceiptAddDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TpReceiptAddDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TpReceiptAddDataResponse proto.InternalMessageInfo

func (m *TpReceiptAddDataResponse) GetStatus() TpReceiptAddDataResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpReceiptAddDataResponse_STATUS_UNSET
}

type TpEventAddRequest struct {
	ContextId            string            `protobuf:"bytes,1,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	Event                *events_pb2.Event `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TpEventAddRequest) Reset()         { *m = TpEventAddRequest{} }
func (m *TpEventAddRequest) String() string { return proto.CompactTextString(m) }
func (*TpEventAddRequest) ProtoMessage()    {}
func (*TpEventAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{9}
}
func (m *TpEventAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpEventAddRequest.Unmarshal(m, b)
}
func (m *TpEventAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpEventAddRequest.Marshal(b, m, deterministic)
}
func (dst *TpEventAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpEventAddRequest.Merge(dst, src)
}
func (m *TpEventAddRequest) XXX_Size() int {
	return xxx_messageInfo_TpEventAddRequest.Size(m)
}
func (m *TpEventAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TpEventAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TpEventAddRequest proto.InternalMessageInfo

func (m *TpEventAddRequest) GetContextId() string {
	if m != nil {
		return m.ContextId
	}
	return ""
}

func (m *TpEventAddRequest) GetEvent() *events_pb2.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type TpEventAddResponse struct {
	Status               TpEventAddResponse_Status `protobuf:"varint,2,opt,name=status,enum=TpEventAddResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TpEventAddResponse) Reset()         { *m = TpEventAddResponse{} }
func (m *TpEventAddResponse) String() string { return proto.CompactTextString(m) }
func (*TpEventAddResponse) ProtoMessage()    {}
func (*TpEventAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_context_e9993d0f0f547372, []int{10}
}
func (m *TpEventAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TpEventAddResponse.Unmarshal(m, b)
}
func (m *TpEventAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TpEventAddResponse.Marshal(b, m, deterministic)
}
func (dst *TpEventAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TpEventAddResponse.Merge(dst, src)
}
func (m *TpEventAddResponse) XXX_Size() int {
	return xxx_messageInfo_TpEventAddResponse.Size(m)
}
func (m *TpEventAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TpEventAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TpEventAddResponse proto.InternalMessageInfo

func (m *TpEventAddResponse) GetStatus() TpEventAddResponse_Status {
	if m != nil {
		return m.Status
	}
	return TpEventAddResponse_STATUS_UNSET
}

func init() {
	proto.RegisterType((*TpStateEntry)(nil), "TpStateEntry")
	proto.RegisterType((*TpStateGetRequest)(nil), "TpStateGetRequest")
	proto.RegisterType((*TpStateGetResponse)(nil), "TpStateGetResponse")
	proto.RegisterType((*TpStateSetRequest)(nil), "TpStateSetRequest")
	proto.RegisterType((*TpStateSetResponse)(nil), "TpStateSetResponse")
	proto.RegisterType((*TpStateDeleteRequest)(nil), "TpStateDeleteRequest")
	proto.RegisterType((*TpStateDeleteResponse)(nil), "TpStateDeleteResponse")
	proto.RegisterType((*TpReceiptAddDataRequest)(nil), "TpReceiptAddDataRequest")
	proto.RegisterType((*TpReceiptAddDataResponse)(nil), "TpReceiptAddDataResponse")
	proto.RegisterType((*TpEventAddRequest)(nil), "TpEventAddRequest")
	proto.RegisterType((*TpEventAddResponse)(nil), "TpEventAddResponse")
	proto.RegisterEnum("TpStateGetResponse_Status", TpStateGetResponse_Status_name, TpStateGetResponse_Status_value)
	proto.RegisterEnum("TpStateSetResponse_Status", TpStateSetResponse_Status_name, TpStateSetResponse_Status_value)
	proto.RegisterEnum("TpStateDeleteResponse_Status", TpStateDeleteResponse_Status_name, TpStateDeleteResponse_Status_value)
	proto.RegisterEnum("TpReceiptAddDataResponse_Status", TpReceiptAddDataResponse_Status_name, TpReceiptAddDataResponse_Status_value)
	proto.RegisterEnum("TpEventAddResponse_Status", TpEventAddResponse_Status_name, TpEventAddResponse_Status_value)
}

func init() {
	proto.RegisterFile("github.com/daludaluking/sawtooth-sdk/protobuf/state_context_pb2/state_context.proto", fileDescriptor_state_context_e9993d0f0f547372)
}

var fileDescriptor_state_context_e9993d0f0f547372 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x65, 0x1d, 0xea, 0x2a, 0xd3, 0x80, 0xd2, 0x85, 0xaa, 0x56, 0x45, 0xa5, 0xc8, 0x17, 0x72,
	0xa0, 0x5b, 0xc9, 0x08, 0x09, 0xa9, 0x5c, 0x8c, 0x1a, 0x41, 0x04, 0x6a, 0xaa, 0x5d, 0xe7, 0x52,
	0x0e, 0x96, 0xd3, 0x1d, 0x44, 0x54, 0x64, 0x9b, 0xec, 0x86, 0xc2, 0x99, 0x4f, 0xe1, 0xc4, 0x85,
	0x6f, 0x44, 0x5e, 0xaf, 0x5b, 0xd7, 0x49, 0x91, 0x25, 0xca, 0x6d, 0x33, 0xf3, 0xe6, 0x65, 0xde,
	0xf3, 0xcc, 0xc0, 0x91, 0x4a, 0x2e, 0x75, 0x96, 0xe9, 0x4f, 0xb1, 0x92, 0x17, 0x87, 0xf9, 0x22,
	0xd3, 0xd9, 0x6c, 0xf9, 0xf1, 0x50, 0xe9, 0x44, 0x63, 0x7c, 0x9e, 0xa5, 0x1a, 0xbf, 0xe9, 0x38,
	0x9f, 0x05, 0x37, 0x23, 0xcc, 0x00, 0xf7, 0x0e, 0xd6, 0x17, 0xe3, 0x57, 0x4c, 0xb5, 0x32, 0x55,
	0xe5, 0xb3, 0x84, 0xfb, 0xaf, 0xa0, 0x17, 0xe5, 0xa2, 0xe0, 0x19, 0xa5, 0x7a, 0xf1, 0x9d, 0x7a,
	0xb0, 0x99, 0x48, 0xb9, 0x40, 0xa5, 0x3c, 0x32, 0x20, 0xc3, 0x2e, 0xaf, 0x7e, 0x52, 0x0a, 0xf7,
	0x65, 0xa2, 0x13, 0xcf, 0x19, 0x90, 0x61, 0x8f, 0x9b, 0xb7, 0x7f, 0x0a, 0xdb, 0xb6, 0xfa, 0x0d,
	0x6a, 0x8e, 0x5f, 0x96, 0xa8, 0x34, 0xdd, 0x07, 0xa8, 0x9a, 0x9c, 0x4b, 0xcb, 0xd2, 0xb5, 0x91,
	0xb1, 0xa4, 0x4f, 0xa0, 0x6b, 0x29, 0x51, 0x79, 0xce, 0xa0, 0x53, 0x64, 0xaf, 0x02, 0xfe, 0x6f,
	0x02, 0xb4, 0x4e, 0xa9, 0xf2, 0x2c, 0x55, 0x48, 0x9f, 0xc2, 0x26, 0xa6, 0x7a, 0x31, 0xc7, 0xa2,
	0xad, 0xce, 0x70, 0x2b, 0x78, 0xc0, 0xea, 0x6d, 0xf3, 0x2a, 0x4b, 0x03, 0x70, 0x0b, 0x57, 0x96,
	0xca, 0xf4, 0xf9, 0x30, 0xd8, 0x63, 0xab, 0x6c, 0x4c, 0x18, 0x04, 0xb7, 0x48, 0xff, 0x08, 0xdc,
	0x32, 0x42, 0xfb, 0xd0, 0x13, 0x51, 0x18, 0x4d, 0x45, 0x3c, 0x3d, 0x11, 0xa3, 0xa8, 0x7f, 0x8f,
	0xba, 0xe0, 0x4c, 0xde, 0xf5, 0x09, 0xdd, 0x85, 0x47, 0xe1, 0x34, 0x7a, 0x3b, 0xe1, 0xe3, 0xb3,
	0x30, 0x1a, 0x4f, 0x4e, 0xe2, 0x11, 0xe7, 0x13, 0xde, 0x77, 0xfc, 0x0f, 0x57, 0x16, 0x88, 0xd6,
	0x16, 0xd4, 0xd4, 0x38, 0x7f, 0x53, 0xe3, 0xff, 0xbc, 0x76, 0x43, 0xd4, 0xdc, 0xb8, 0x61, 0x21,
	0x69, 0x58, 0x78, 0xbb, 0x05, 0xe2, 0x3f, 0x59, 0x20, 0xe0, 0xb1, 0xfd, 0x87, 0x63, 0xfc, 0x8c,
	0x1a, 0xef, 0x64, 0x10, 0x7e, 0x11, 0xd8, 0x69, 0xb0, 0xb6, 0x52, 0xff, 0xa2, 0xa1, 0x7e, 0x9f,
	0xad, 0x65, 0xb9, 0x53, 0x03, 0xde, 0xc3, 0x6e, 0x94, 0x73, 0x3c, 0xc7, 0x79, 0xae, 0x43, 0x29,
	0x8f, 0x13, 0x9d, 0xb4, 0xf4, 0xa0, 0x5a, 0xaa, 0x4e, 0x6d, 0xa9, 0x7e, 0x10, 0xf0, 0x56, 0xe9,
	0xac, 0xf8, 0x97, 0x0d, 0x79, 0x03, 0x76, 0x1b, 0xb4, 0xa9, 0xf0, 0xa0, 0x85, 0xc2, 0x2e, 0x6c,
	0x54, 0x9a, 0xcc, 0x6a, 0x8f, 0x8a, 0x53, 0x11, 0x4a, 0xd9, 0xfa, 0x8b, 0x6e, 0x98, 0xe3, 0x62,
	0x7a, 0xdb, 0x0a, 0x5c, 0x66, 0xea, 0x79, 0x19, 0xf4, 0x2f, 0x8b, 0x59, 0xbe, 0x66, 0xb4, 0x82,
	0xd6, 0x4d, 0x6b, 0x13, 0xf4, 0x6f, 0x52, 0x5e, 0x3f, 0x83, 0x9d, 0xea, 0x28, 0x32, 0x25, 0x2f,
	0x58, 0x75, 0x14, 0x4f, 0xc9, 0xd9, 0xf6, 0xca, 0x51, 0x9d, 0xb9, 0x26, 0xfd, 0xfc, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd9, 0xdd, 0x1e, 0xd6, 0x86, 0x05, 0x00, 0x00,
}
