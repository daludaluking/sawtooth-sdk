// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/daludaluking/sawtooth_sdk/protobuf/client_batch_pb2/client_batch.proto

package client_batch_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import batch_pb2 "github.com/daludaluking/sawtooth_sdk/protobuf/batch_pb2"
import client_list_control_pb2 "github.com/daludaluking/sawtooth_sdk/protobuf/client_list_control_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientBatchListResponse_Status int32

const (
	ClientBatchListResponse_STATUS_UNSET   ClientBatchListResponse_Status = 0
	ClientBatchListResponse_OK             ClientBatchListResponse_Status = 1
	ClientBatchListResponse_INTERNAL_ERROR ClientBatchListResponse_Status = 2
	ClientBatchListResponse_NOT_READY      ClientBatchListResponse_Status = 3
	ClientBatchListResponse_NO_ROOT        ClientBatchListResponse_Status = 4
	ClientBatchListResponse_NO_RESOURCE    ClientBatchListResponse_Status = 5
	ClientBatchListResponse_INVALID_PAGING ClientBatchListResponse_Status = 6
	ClientBatchListResponse_INVALID_SORT   ClientBatchListResponse_Status = 7
	ClientBatchListResponse_INVALID_ID     ClientBatchListResponse_Status = 8
)

var ClientBatchListResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "NOT_READY",
	4: "NO_ROOT",
	5: "NO_RESOURCE",
	6: "INVALID_PAGING",
	7: "INVALID_SORT",
	8: "INVALID_ID",
}
var ClientBatchListResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NOT_READY":      3,
	"NO_ROOT":        4,
	"NO_RESOURCE":    5,
	"INVALID_PAGING": 6,
	"INVALID_SORT":   7,
	"INVALID_ID":     8,
}

func (x ClientBatchListResponse_Status) String() string {
	return proto.EnumName(ClientBatchListResponse_Status_name, int32(x))
}
func (ClientBatchListResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_batch_dfb7c8db88c2fcf8, []int{1, 0}
}

type ClientBatchGetResponse_Status int32

const (
	ClientBatchGetResponse_STATUS_UNSET   ClientBatchGetResponse_Status = 0
	ClientBatchGetResponse_OK             ClientBatchGetResponse_Status = 1
	ClientBatchGetResponse_INTERNAL_ERROR ClientBatchGetResponse_Status = 2
	ClientBatchGetResponse_NO_RESOURCE    ClientBatchGetResponse_Status = 5
	ClientBatchGetResponse_INVALID_ID     ClientBatchGetResponse_Status = 8
)

var ClientBatchGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	5: "NO_RESOURCE",
	8: "INVALID_ID",
}
var ClientBatchGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NO_RESOURCE":    5,
	"INVALID_ID":     8,
}

func (x ClientBatchGetResponse_Status) String() string {
	return proto.EnumName(ClientBatchGetResponse_Status_name, int32(x))
}
func (ClientBatchGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_batch_dfb7c8db88c2fcf8, []int{3, 0}
}

// A request to return a list of batches from the validator. May include the id
// of a particular block to be the `head` of the chain being requested. In that
// case the list will include the batches from that block, and all batches
// previous to that block on the chain. Filter with specific `batch_ids`.
type ClientBatchListRequest struct {
	HeadId               string                                        `protobuf:"bytes,1,opt,name=head_id,json=headId" json:"head_id,omitempty"`
	BatchIds             []string                                      `protobuf:"bytes,2,rep,name=batch_ids,json=batchIds" json:"batch_ids,omitempty"`
	Paging               *client_list_control_pb2.ClientPagingControls `protobuf:"bytes,3,opt,name=paging" json:"paging,omitempty"`
	Sorting              []*client_list_control_pb2.ClientSortControls `protobuf:"bytes,4,rep,name=sorting" json:"sorting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *ClientBatchListRequest) Reset()         { *m = ClientBatchListRequest{} }
func (m *ClientBatchListRequest) String() string { return proto.CompactTextString(m) }
func (*ClientBatchListRequest) ProtoMessage()    {}
func (*ClientBatchListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_batch_dfb7c8db88c2fcf8, []int{0}
}
func (m *ClientBatchListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientBatchListRequest.Unmarshal(m, b)
}
func (m *ClientBatchListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientBatchListRequest.Marshal(b, m, deterministic)
}
func (dst *ClientBatchListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientBatchListRequest.Merge(dst, src)
}
func (m *ClientBatchListRequest) XXX_Size() int {
	return xxx_messageInfo_ClientBatchListRequest.Size(m)
}
func (m *ClientBatchListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientBatchListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientBatchListRequest proto.InternalMessageInfo

func (m *ClientBatchListRequest) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientBatchListRequest) GetBatchIds() []string {
	if m != nil {
		return m.BatchIds
	}
	return nil
}

func (m *ClientBatchListRequest) GetPaging() *client_list_control_pb2.ClientPagingControls {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *ClientBatchListRequest) GetSorting() []*client_list_control_pb2.ClientSortControls {
	if m != nil {
		return m.Sorting
	}
	return nil
}

// A response that lists batches from newest to oldest.
//
// Statuses:
//   * OK - everything worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NOT_READY - the validator does not yet have a genesis block
//   * NO_ROOT - the head block specified was not found
//   * NO_RESOURCE - no batches were found with the parameters specified
//   * INVALID_PAGING - the paging controls were malformed or out of range
//   * INVALID_SORT - the sorting controls were malformed or invalid
type ClientBatchListResponse struct {
	Status               ClientBatchListResponse_Status                `protobuf:"varint,1,opt,name=status,enum=ClientBatchListResponse_Status" json:"status,omitempty"`
	Batches              []*batch_pb2.Batch                            `protobuf:"bytes,2,rep,name=batches" json:"batches,omitempty"`
	HeadId               string                                        `protobuf:"bytes,3,opt,name=head_id,json=headId" json:"head_id,omitempty"`
	Paging               *client_list_control_pb2.ClientPagingResponse `protobuf:"bytes,4,opt,name=paging" json:"paging,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *ClientBatchListResponse) Reset()         { *m = ClientBatchListResponse{} }
func (m *ClientBatchListResponse) String() string { return proto.CompactTextString(m) }
func (*ClientBatchListResponse) ProtoMessage()    {}
func (*ClientBatchListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_batch_dfb7c8db88c2fcf8, []int{1}
}
func (m *ClientBatchListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientBatchListResponse.Unmarshal(m, b)
}
func (m *ClientBatchListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientBatchListResponse.Marshal(b, m, deterministic)
}
func (dst *ClientBatchListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientBatchListResponse.Merge(dst, src)
}
func (m *ClientBatchListResponse) XXX_Size() int {
	return xxx_messageInfo_ClientBatchListResponse.Size(m)
}
func (m *ClientBatchListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientBatchListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientBatchListResponse proto.InternalMessageInfo

func (m *ClientBatchListResponse) GetStatus() ClientBatchListResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBatchListResponse_STATUS_UNSET
}

func (m *ClientBatchListResponse) GetBatches() []*batch_pb2.Batch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func (m *ClientBatchListResponse) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientBatchListResponse) GetPaging() *client_list_control_pb2.ClientPagingResponse {
	if m != nil {
		return m.Paging
	}
	return nil
}

// Fetches a specific batch by its id (header_signature) from the blockchain.
type ClientBatchGetRequest struct {
	BatchId              string   `protobuf:"bytes,1,opt,name=batch_id,json=batchId" json:"batch_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientBatchGetRequest) Reset()         { *m = ClientBatchGetRequest{} }
func (m *ClientBatchGetRequest) String() string { return proto.CompactTextString(m) }
func (*ClientBatchGetRequest) ProtoMessage()    {}
func (*ClientBatchGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_batch_dfb7c8db88c2fcf8, []int{2}
}
func (m *ClientBatchGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientBatchGetRequest.Unmarshal(m, b)
}
func (m *ClientBatchGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientBatchGetRequest.Marshal(b, m, deterministic)
}
func (dst *ClientBatchGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientBatchGetRequest.Merge(dst, src)
}
func (m *ClientBatchGetRequest) XXX_Size() int {
	return xxx_messageInfo_ClientBatchGetRequest.Size(m)
}
func (m *ClientBatchGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientBatchGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientBatchGetRequest proto.InternalMessageInfo

func (m *ClientBatchGetRequest) GetBatchId() string {
	if m != nil {
		return m.BatchId
	}
	return ""
}

// A response that returns the batch specified by a ClientBatchGetRequest.
//
// Statuses:
//   * OK - everything worked as expected, batch has been fetched
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NO_RESOURCE - no batch with the specified id exists
type ClientBatchGetResponse struct {
	Status               ClientBatchGetResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientBatchGetResponse_Status" json:"status,omitempty"`
	Batch                *batch_pb2.Batch              `protobuf:"bytes,2,opt,name=batch" json:"batch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ClientBatchGetResponse) Reset()         { *m = ClientBatchGetResponse{} }
func (m *ClientBatchGetResponse) String() string { return proto.CompactTextString(m) }
func (*ClientBatchGetResponse) ProtoMessage()    {}
func (*ClientBatchGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_batch_dfb7c8db88c2fcf8, []int{3}
}
func (m *ClientBatchGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientBatchGetResponse.Unmarshal(m, b)
}
func (m *ClientBatchGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientBatchGetResponse.Marshal(b, m, deterministic)
}
func (dst *ClientBatchGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientBatchGetResponse.Merge(dst, src)
}
func (m *ClientBatchGetResponse) XXX_Size() int {
	return xxx_messageInfo_ClientBatchGetResponse.Size(m)
}
func (m *ClientBatchGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientBatchGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientBatchGetResponse proto.InternalMessageInfo

func (m *ClientBatchGetResponse) GetStatus() ClientBatchGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBatchGetResponse_STATUS_UNSET
}

func (m *ClientBatchGetResponse) GetBatch() *batch_pb2.Batch {
	if m != nil {
		return m.Batch
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientBatchListRequest)(nil), "ClientBatchListRequest")
	proto.RegisterType((*ClientBatchListResponse)(nil), "ClientBatchListResponse")
	proto.RegisterType((*ClientBatchGetRequest)(nil), "ClientBatchGetRequest")
	proto.RegisterType((*ClientBatchGetResponse)(nil), "ClientBatchGetResponse")
	proto.RegisterEnum("ClientBatchListResponse_Status", ClientBatchListResponse_Status_name, ClientBatchListResponse_Status_value)
	proto.RegisterEnum("ClientBatchGetResponse_Status", ClientBatchGetResponse_Status_name, ClientBatchGetResponse_Status_value)
}

func init() {
	proto.RegisterFile("github.com/daludaluking/sawtooth_sdk/protobuf/client_batch_pb2/client_batch.proto", fileDescriptor_client_batch_dfb7c8db88c2fcf8)
}

var fileDescriptor_client_batch_dfb7c8db88c2fcf8 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x8e, 0x93, 0x40,
	0x14, 0x96, 0xb2, 0x0b, 0xed, 0x41, 0xeb, 0x64, 0x4c, 0xdd, 0xfa, 0x13, 0x25, 0x5c, 0x35, 0xd9,
	0x2c, 0x9b, 0x60, 0xa2, 0xf1, 0x92, 0x6d, 0x49, 0x43, 0x6c, 0xa0, 0x19, 0xa8, 0x46, 0x6f, 0x26,
	0xb4, 0xe0, 0x96, 0x6c, 0x53, 0x6a, 0x67, 0x1a, 0x9f, 0xc5, 0x77, 0xf0, 0x3d, 0x7c, 0x05, 0x1f,
	0xc7, 0x30, 0x30, 0xbb, 0xb4, 0x5b, 0xbd, 0xf0, 0x72, 0xbe, 0xf9, 0xce, 0xf9, 0xce, 0xf9, 0xbe,
	0x19, 0x78, 0xcf, 0x92, 0xef, 0xbc, 0x28, 0xf8, 0x92, 0xb2, 0xf4, 0xe6, 0x72, 0xb3, 0x2d, 0x78,
	0x31, 0xdf, 0x7d, 0xbd, 0x5c, 0xac, 0xf2, 0x6c, 0xcd, 0xe9, 0x3c, 0xe1, 0x8b, 0x25, 0xdd, 0xcc,
	0x9d, 0x3d, 0xc0, 0x16, 0xb4, 0xe7, 0xe7, 0xc7, 0x4b, 0xef, 0x6a, 0x9a, 0xe4, 0xf1, 0x3f, 0x75,
	0x56, 0x39, 0xe3, 0x74, 0x51, 0xac, 0xf9, 0xb6, 0x58, 0x35, 0xe5, 0x9a, 0x78, 0xd5, 0xc8, 0xfa,
	0xa9, 0xc0, 0xd3, 0xa1, 0xb8, 0xbd, 0x2a, 0xdb, 0x4f, 0x72, 0xc6, 0x49, 0xf6, 0x6d, 0x97, 0x31,
	0x8e, 0xcf, 0x40, 0x5f, 0x66, 0x49, 0x4a, 0xf3, 0xb4, 0xaf, 0x98, 0xca, 0xa0, 0x43, 0xb4, 0xf2,
	0xe8, 0xa7, 0xf8, 0x05, 0x74, 0xaa, 0xa9, 0xf2, 0x94, 0xf5, 0x5b, 0xa6, 0x3a, 0xe8, 0x90, 0xb6,
	0x00, 0xfc, 0x94, 0xe1, 0x0b, 0xd0, 0x36, 0xc9, 0x75, 0xbe, 0xbe, 0xee, 0xab, 0xa6, 0x32, 0x30,
	0x9c, 0x9e, 0x5d, 0xb5, 0x9f, 0x0a, 0x70, 0x58, 0x89, 0x33, 0x52, 0x93, 0xf0, 0x05, 0xe8, 0xac,
	0xd8, 0xf2, 0x92, 0x7f, 0x62, 0xaa, 0x03, 0xc3, 0x79, 0x52, 0xf3, 0xa3, 0x62, 0xcb, 0x6f, 0xd9,
	0x92, 0x63, 0xfd, 0x6e, 0xc1, 0xd9, 0xbd, 0x71, 0xd9, 0xa6, 0x58, 0xb3, 0x0c, 0xbf, 0x03, 0x8d,
	0xf1, 0x84, 0xef, 0x98, 0x18, 0xb7, 0xeb, 0xbc, 0xb6, 0xff, 0xc2, 0xb4, 0x23, 0x41, 0x23, 0x35,
	0x1d, 0x9b, 0xa0, 0x8b, 0xf1, 0xb3, 0x6a, 0x1b, 0xc3, 0xd1, 0x6c, 0x51, 0x43, 0x24, 0xdc, 0xb4,
	0x42, 0xdd, 0xb3, 0xe2, 0x6e, 0xdb, 0x93, 0x23, 0xdb, 0x4a, 0x41, 0xb9, 0xad, 0xf5, 0x43, 0x01,
	0xad, 0x12, 0xc7, 0x08, 0x1e, 0x46, 0xb1, 0x1b, 0xcf, 0x22, 0x3a, 0x0b, 0x22, 0x2f, 0x46, 0x0f,
	0xb0, 0x06, 0xad, 0xf0, 0x03, 0x52, 0x30, 0x86, 0xae, 0x1f, 0xc4, 0x1e, 0x09, 0xdc, 0x09, 0xf5,
	0x08, 0x09, 0x09, 0x6a, 0xe1, 0x47, 0xd0, 0x09, 0xc2, 0x98, 0x12, 0xcf, 0x1d, 0x7d, 0x46, 0x2a,
	0x36, 0x40, 0x0f, 0x42, 0x4a, 0xc2, 0x30, 0x46, 0x27, 0xf8, 0x31, 0x18, 0xe5, 0xc1, 0x8b, 0xc2,
	0x19, 0x19, 0x7a, 0xe8, 0xb4, 0x6a, 0xf0, 0xd1, 0x9d, 0xf8, 0x23, 0x3a, 0x75, 0xc7, 0x7e, 0x30,
	0x46, 0x5a, 0x29, 0x27, 0xb1, 0x28, 0x24, 0x31, 0xd2, 0x71, 0x17, 0x40, 0x22, 0xfe, 0x08, 0xb5,
	0x2d, 0x07, 0x7a, 0x0d, 0xbf, 0xc6, 0xd9, 0xed, 0x3b, 0x78, 0x06, 0x6d, 0x19, 0x77, 0xfd, 0x10,
	0xf4, 0x3a, 0x6d, 0xeb, 0xd7, 0xfe, 0xeb, 0x11, 0x45, 0x75, 0x1a, 0x6f, 0x0f, 0xd2, 0x78, 0x65,
	0x1f, 0x27, 0x1e, 0x86, 0xf1, 0x12, 0x4e, 0x45, 0xf7, 0x7e, 0x4b, 0x18, 0x2a, 0xa3, 0xa8, 0x40,
	0xeb, 0xd3, 0x7f, 0xfa, 0x77, 0xcf, 0xa3, 0x83, 0xed, 0xaf, 0xce, 0xa1, 0x27, 0xbf, 0x94, 0xcd,
	0xd2, 0x1b, 0x5b, 0x7e, 0xa9, 0xa9, 0xf2, 0x05, 0x1d, 0xfe, 0xde, 0xb9, 0x26, 0x6e, 0xdf, 0xfc,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x22, 0xee, 0x60, 0xee, 0x03, 0x00, 0x00,
}
