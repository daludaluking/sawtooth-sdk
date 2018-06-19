// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/daludaluking/sawtooth_sdk/protobuf/client_block_pb2/client_block.proto

package client_block_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import block_pb2 "github.com/daludaluking/sawtooth_sdk/protobuf/block_pb2"
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

type ClientBlockListResponse_Status int32

const (
	ClientBlockListResponse_STATUS_UNSET   ClientBlockListResponse_Status = 0
	ClientBlockListResponse_OK             ClientBlockListResponse_Status = 1
	ClientBlockListResponse_INTERNAL_ERROR ClientBlockListResponse_Status = 2
	ClientBlockListResponse_NOT_READY      ClientBlockListResponse_Status = 3
	ClientBlockListResponse_NO_ROOT        ClientBlockListResponse_Status = 4
	ClientBlockListResponse_NO_RESOURCE    ClientBlockListResponse_Status = 5
	ClientBlockListResponse_INVALID_PAGING ClientBlockListResponse_Status = 6
	ClientBlockListResponse_INVALID_SORT   ClientBlockListResponse_Status = 7
	ClientBlockListResponse_INVALID_ID     ClientBlockListResponse_Status = 8
)

var ClientBlockListResponse_Status_name = map[int32]string{
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
var ClientBlockListResponse_Status_value = map[string]int32{
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

func (x ClientBlockListResponse_Status) String() string {
	return proto.EnumName(ClientBlockListResponse_Status_name, int32(x))
}
func (ClientBlockListResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_block_b2f3f7e0db13ccb7, []int{1, 0}
}

type ClientBlockGetResponse_Status int32

const (
	ClientBlockGetResponse_STATUS_UNSET   ClientBlockGetResponse_Status = 0
	ClientBlockGetResponse_OK             ClientBlockGetResponse_Status = 1
	ClientBlockGetResponse_INTERNAL_ERROR ClientBlockGetResponse_Status = 2
	ClientBlockGetResponse_NO_RESOURCE    ClientBlockGetResponse_Status = 5
	ClientBlockGetResponse_INVALID_ID     ClientBlockGetResponse_Status = 8
)

var ClientBlockGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	5: "NO_RESOURCE",
	8: "INVALID_ID",
}
var ClientBlockGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"NO_RESOURCE":    5,
	"INVALID_ID":     8,
}

func (x ClientBlockGetResponse_Status) String() string {
	return proto.EnumName(ClientBlockGetResponse_Status_name, int32(x))
}
func (ClientBlockGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_block_b2f3f7e0db13ccb7, []int{6, 0}
}

// A request to return a list of blocks from the validator. May include the id
// of a particular block to be the `head` of the chain being requested. In that
// case the list will include that block (if found), and all blocks previous
// to it on the chain. Can be filtered using specific `block_ids`.
type ClientBlockListRequest struct {
	HeadId               string                                        `protobuf:"bytes,1,opt,name=head_id,json=headId" json:"head_id,omitempty"`
	BlockIds             []string                                      `protobuf:"bytes,2,rep,name=block_ids,json=blockIds" json:"block_ids,omitempty"`
	Paging               *client_list_control_pb2.ClientPagingControls `protobuf:"bytes,3,opt,name=paging" json:"paging,omitempty"`
	Sorting              []*client_list_control_pb2.ClientSortControls `protobuf:"bytes,4,rep,name=sorting" json:"sorting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *ClientBlockListRequest) Reset()         { *m = ClientBlockListRequest{} }
func (m *ClientBlockListRequest) String() string { return proto.CompactTextString(m) }
func (*ClientBlockListRequest) ProtoMessage()    {}
func (*ClientBlockListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_block_b2f3f7e0db13ccb7, []int{0}
}
func (m *ClientBlockListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientBlockListRequest.Unmarshal(m, b)
}
func (m *ClientBlockListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientBlockListRequest.Marshal(b, m, deterministic)
}
func (dst *ClientBlockListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientBlockListRequest.Merge(dst, src)
}
func (m *ClientBlockListRequest) XXX_Size() int {
	return xxx_messageInfo_ClientBlockListRequest.Size(m)
}
func (m *ClientBlockListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientBlockListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientBlockListRequest proto.InternalMessageInfo

func (m *ClientBlockListRequest) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientBlockListRequest) GetBlockIds() []string {
	if m != nil {
		return m.BlockIds
	}
	return nil
}

func (m *ClientBlockListRequest) GetPaging() *client_list_control_pb2.ClientPagingControls {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *ClientBlockListRequest) GetSorting() []*client_list_control_pb2.ClientSortControls {
	if m != nil {
		return m.Sorting
	}
	return nil
}

// A response that lists a chain of blocks with the newest at the beginning,
// and the oldest (genesis) block at the end.
//
// Statuses:
//   * OK - everything worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NOT_READY - the validator does not yet have a genesis block
//   * NO_ROOT - the head block specified was not found
//   * NO_RESOURCE - no blocks were found with the parameters specified
//   * INVALID_PAGING - the paging controls were malformed or out of range
//   * INVALID_SORT - the sorting controls were malformed or invalid
type ClientBlockListResponse struct {
	Status               ClientBlockListResponse_Status                `protobuf:"varint,1,opt,name=status,enum=ClientBlockListResponse_Status" json:"status,omitempty"`
	Blocks               []*block_pb2.Block                            `protobuf:"bytes,2,rep,name=blocks" json:"blocks,omitempty"`
	HeadId               string                                        `protobuf:"bytes,3,opt,name=head_id,json=headId" json:"head_id,omitempty"`
	Paging               *client_list_control_pb2.ClientPagingResponse `protobuf:"bytes,4,opt,name=paging" json:"paging,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *ClientBlockListResponse) Reset()         { *m = ClientBlockListResponse{} }
func (m *ClientBlockListResponse) String() string { return proto.CompactTextString(m) }
func (*ClientBlockListResponse) ProtoMessage()    {}
func (*ClientBlockListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_block_b2f3f7e0db13ccb7, []int{1}
}
func (m *ClientBlockListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientBlockListResponse.Unmarshal(m, b)
}
func (m *ClientBlockListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientBlockListResponse.Marshal(b, m, deterministic)
}
func (dst *ClientBlockListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientBlockListResponse.Merge(dst, src)
}
func (m *ClientBlockListResponse) XXX_Size() int {
	return xxx_messageInfo_ClientBlockListResponse.Size(m)
}
func (m *ClientBlockListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientBlockListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientBlockListResponse proto.InternalMessageInfo

func (m *ClientBlockListResponse) GetStatus() ClientBlockListResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBlockListResponse_STATUS_UNSET
}

func (m *ClientBlockListResponse) GetBlocks() []*block_pb2.Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *ClientBlockListResponse) GetHeadId() string {
	if m != nil {
		return m.HeadId
	}
	return ""
}

func (m *ClientBlockListResponse) GetPaging() *client_list_control_pb2.ClientPagingResponse {
	if m != nil {
		return m.Paging
	}
	return nil
}

// A request to return a specific block from the validator. The block must be
// specified by its unique id, in this case the block's header signature
type ClientBlockGetByIdRequest struct {
	BlockId              string   `protobuf:"bytes,1,opt,name=block_id,json=blockId" json:"block_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientBlockGetByIdRequest) Reset()         { *m = ClientBlockGetByIdRequest{} }
func (m *ClientBlockGetByIdRequest) String() string { return proto.CompactTextString(m) }
func (*ClientBlockGetByIdRequest) ProtoMessage()    {}
func (*ClientBlockGetByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_block_b2f3f7e0db13ccb7, []int{2}
}
func (m *ClientBlockGetByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientBlockGetByIdRequest.Unmarshal(m, b)
}
func (m *ClientBlockGetByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientBlockGetByIdRequest.Marshal(b, m, deterministic)
}
func (dst *ClientBlockGetByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientBlockGetByIdRequest.Merge(dst, src)
}
func (m *ClientBlockGetByIdRequest) XXX_Size() int {
	return xxx_messageInfo_ClientBlockGetByIdRequest.Size(m)
}
func (m *ClientBlockGetByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientBlockGetByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientBlockGetByIdRequest proto.InternalMessageInfo

func (m *ClientBlockGetByIdRequest) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

// A request to return a specific block from the validator. The block must be
// specified by its block number
type ClientBlockGetByNumRequest struct {
	BlockNum             uint64   `protobuf:"varint,1,opt,name=block_num,json=blockNum" json:"block_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientBlockGetByNumRequest) Reset()         { *m = ClientBlockGetByNumRequest{} }
func (m *ClientBlockGetByNumRequest) String() string { return proto.CompactTextString(m) }
func (*ClientBlockGetByNumRequest) ProtoMessage()    {}
func (*ClientBlockGetByNumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_block_b2f3f7e0db13ccb7, []int{3}
}
func (m *ClientBlockGetByNumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientBlockGetByNumRequest.Unmarshal(m, b)
}
func (m *ClientBlockGetByNumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientBlockGetByNumRequest.Marshal(b, m, deterministic)
}
func (dst *ClientBlockGetByNumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientBlockGetByNumRequest.Merge(dst, src)
}
func (m *ClientBlockGetByNumRequest) XXX_Size() int {
	return xxx_messageInfo_ClientBlockGetByNumRequest.Size(m)
}
func (m *ClientBlockGetByNumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientBlockGetByNumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientBlockGetByNumRequest proto.InternalMessageInfo

func (m *ClientBlockGetByNumRequest) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

// A request to return a specific block from the validator. The block
// containing the given transaction is returned. If no block on the current
// chain contains the transaction, NO_RESOURCE is returned.
type ClientBlockGetByTransactionIdRequest struct {
	TransactionId        string   `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientBlockGetByTransactionIdRequest) Reset()         { *m = ClientBlockGetByTransactionIdRequest{} }
func (m *ClientBlockGetByTransactionIdRequest) String() string { return proto.CompactTextString(m) }
func (*ClientBlockGetByTransactionIdRequest) ProtoMessage()    {}
func (*ClientBlockGetByTransactionIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_block_b2f3f7e0db13ccb7, []int{4}
}
func (m *ClientBlockGetByTransactionIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientBlockGetByTransactionIdRequest.Unmarshal(m, b)
}
func (m *ClientBlockGetByTransactionIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientBlockGetByTransactionIdRequest.Marshal(b, m, deterministic)
}
func (dst *ClientBlockGetByTransactionIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientBlockGetByTransactionIdRequest.Merge(dst, src)
}
func (m *ClientBlockGetByTransactionIdRequest) XXX_Size() int {
	return xxx_messageInfo_ClientBlockGetByTransactionIdRequest.Size(m)
}
func (m *ClientBlockGetByTransactionIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientBlockGetByTransactionIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientBlockGetByTransactionIdRequest proto.InternalMessageInfo

func (m *ClientBlockGetByTransactionIdRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

// A request to return a specific block from the validator. The block
// containing the given batch is returned. If no block on the current chain
// contains the batch, NO_RESOURCE is returned.
type ClientBlockGetByBatchIdRequest struct {
	BatchId              string   `protobuf:"bytes,1,opt,name=batch_id,json=batchId" json:"batch_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientBlockGetByBatchIdRequest) Reset()         { *m = ClientBlockGetByBatchIdRequest{} }
func (m *ClientBlockGetByBatchIdRequest) String() string { return proto.CompactTextString(m) }
func (*ClientBlockGetByBatchIdRequest) ProtoMessage()    {}
func (*ClientBlockGetByBatchIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_block_b2f3f7e0db13ccb7, []int{5}
}
func (m *ClientBlockGetByBatchIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientBlockGetByBatchIdRequest.Unmarshal(m, b)
}
func (m *ClientBlockGetByBatchIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientBlockGetByBatchIdRequest.Marshal(b, m, deterministic)
}
func (dst *ClientBlockGetByBatchIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientBlockGetByBatchIdRequest.Merge(dst, src)
}
func (m *ClientBlockGetByBatchIdRequest) XXX_Size() int {
	return xxx_messageInfo_ClientBlockGetByBatchIdRequest.Size(m)
}
func (m *ClientBlockGetByBatchIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientBlockGetByBatchIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientBlockGetByBatchIdRequest proto.InternalMessageInfo

func (m *ClientBlockGetByBatchIdRequest) GetBatchId() string {
	if m != nil {
		return m.BatchId
	}
	return ""
}

// A response that returns the block specified by a ClientBlockGetByIdRequest
// or  ClientBlockGetByNumRequest.
//
// Statuses:
//   * OK - everything worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NO_RESOURCE - no block with the specified id exists
type ClientBlockGetResponse struct {
	Status               ClientBlockGetResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientBlockGetResponse_Status" json:"status,omitempty"`
	Block                *block_pb2.Block              `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ClientBlockGetResponse) Reset()         { *m = ClientBlockGetResponse{} }
func (m *ClientBlockGetResponse) String() string { return proto.CompactTextString(m) }
func (*ClientBlockGetResponse) ProtoMessage()    {}
func (*ClientBlockGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_block_b2f3f7e0db13ccb7, []int{6}
}
func (m *ClientBlockGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientBlockGetResponse.Unmarshal(m, b)
}
func (m *ClientBlockGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientBlockGetResponse.Marshal(b, m, deterministic)
}
func (dst *ClientBlockGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientBlockGetResponse.Merge(dst, src)
}
func (m *ClientBlockGetResponse) XXX_Size() int {
	return xxx_messageInfo_ClientBlockGetResponse.Size(m)
}
func (m *ClientBlockGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientBlockGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientBlockGetResponse proto.InternalMessageInfo

func (m *ClientBlockGetResponse) GetStatus() ClientBlockGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientBlockGetResponse_STATUS_UNSET
}

func (m *ClientBlockGetResponse) GetBlock() *block_pb2.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientBlockListRequest)(nil), "ClientBlockListRequest")
	proto.RegisterType((*ClientBlockListResponse)(nil), "ClientBlockListResponse")
	proto.RegisterType((*ClientBlockGetByIdRequest)(nil), "ClientBlockGetByIdRequest")
	proto.RegisterType((*ClientBlockGetByNumRequest)(nil), "ClientBlockGetByNumRequest")
	proto.RegisterType((*ClientBlockGetByTransactionIdRequest)(nil), "ClientBlockGetByTransactionIdRequest")
	proto.RegisterType((*ClientBlockGetByBatchIdRequest)(nil), "ClientBlockGetByBatchIdRequest")
	proto.RegisterType((*ClientBlockGetResponse)(nil), "ClientBlockGetResponse")
	proto.RegisterEnum("ClientBlockListResponse_Status", ClientBlockListResponse_Status_name, ClientBlockListResponse_Status_value)
	proto.RegisterEnum("ClientBlockGetResponse_Status", ClientBlockGetResponse_Status_name, ClientBlockGetResponse_Status_value)
}

func init() {
	proto.RegisterFile("github.com/daludaluking/sawtooth_sdk/protobuf/client_block_pb2/client_block.proto", fileDescriptor_client_block_b2f3f7e0db13ccb7)
}

var fileDescriptor_client_block_b2f3f7e0db13ccb7 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x25, 0x6d, 0x97, 0xae, 0xb7, 0xac, 0x58, 0x46, 0x63, 0xdd, 0x40, 0xa3, 0x8a, 0x40, 0xaa,
	0x34, 0x2d, 0x93, 0x8a, 0x34, 0x34, 0xf1, 0xd4, 0xae, 0x51, 0x15, 0x51, 0x92, 0xca, 0x49, 0x41,
	0xf0, 0x12, 0xa5, 0x4d, 0xb6, 0x46, 0x6b, 0xe3, 0x52, 0x3b, 0x42, 0x7c, 0x0a, 0xff, 0xc0, 0x7f,
	0xf0, 0x07, 0x7c, 0x0f, 0x8a, 0x93, 0x2c, 0x69, 0x56, 0x78, 0xe0, 0x31, 0xc7, 0xe7, 0xde, 0x73,
	0x8f, 0xef, 0x89, 0xe1, 0x8a, 0xb9, 0xdf, 0x38, 0xa5, 0x7c, 0xe1, 0x30, 0xef, 0xee, 0x62, 0xbd,
	0xa1, 0x9c, 0xce, 0xa2, 0x9b, 0x8b, 0xf9, 0x32, 0xf0, 0x43, 0xee, 0xcc, 0x96, 0x74, 0x7e, 0xe7,
	0xac, 0x67, 0xbd, 0x2d, 0x40, 0x15, 0xb4, 0x93, 0xb3, 0xdd, 0xa5, 0x79, 0x4d, 0x91, 0x3c, 0xfa,
	0xa7, 0xce, 0x32, 0x60, 0xdc, 0x99, 0xd3, 0x90, 0x6f, 0xe8, 0xb2, 0x28, 0x57, 0xc4, 0x93, 0x46,
	0xca, 0x4f, 0x09, 0x9e, 0x5d, 0x8b, 0xd3, 0x41, 0xdc, 0x7e, 0x1c, 0x30, 0x4e, 0xfc, 0xaf, 0x91,
	0xcf, 0x38, 0x3e, 0x82, 0xfa, 0xc2, 0x77, 0x3d, 0x27, 0xf0, 0xda, 0x52, 0x47, 0xea, 0x36, 0x88,
	0x1c, 0x7f, 0xea, 0x1e, 0x7e, 0x0e, 0x8d, 0x64, 0xaa, 0xc0, 0x63, 0xed, 0x4a, 0xa7, 0xda, 0x6d,
	0x90, 0x7d, 0x01, 0xe8, 0x1e, 0xc3, 0xe7, 0x20, 0xaf, 0xdd, 0xdb, 0x20, 0xbc, 0x6d, 0x57, 0x3b,
	0x52, 0xb7, 0xd9, 0x3b, 0x54, 0x93, 0xf6, 0x13, 0x01, 0x5e, 0x27, 0xe2, 0x8c, 0xa4, 0x24, 0x7c,
	0x0e, 0x75, 0x46, 0x37, 0x3c, 0xe6, 0xd7, 0x3a, 0xd5, 0x6e, 0xb3, 0xf7, 0x34, 0xe5, 0x5b, 0x74,
	0xc3, 0xef, 0xd9, 0x19, 0x47, 0xf9, 0x5d, 0x81, 0xa3, 0x07, 0xe3, 0xb2, 0x35, 0x0d, 0x99, 0x8f,
	0xdf, 0x82, 0xcc, 0xb8, 0xcb, 0x23, 0x26, 0xc6, 0x6d, 0xf5, 0x5e, 0xaa, 0x7f, 0x61, 0xaa, 0x96,
	0xa0, 0x91, 0x94, 0x8e, 0x4f, 0x41, 0x16, 0xe3, 0x27, 0x66, 0x9a, 0x3d, 0x59, 0x15, 0x25, 0x24,
	0x45, 0x8b, 0x17, 0x51, 0xdd, 0xba, 0x88, 0xdc, 0x6b, 0x6d, 0x87, 0xd7, 0x4c, 0x2e, 0xf3, 0xaa,
	0xfc, 0x90, 0x40, 0x4e, 0xa4, 0x31, 0x82, 0xc7, 0x96, 0xdd, 0xb7, 0xa7, 0x96, 0x33, 0x35, 0x2c,
	0xcd, 0x46, 0x8f, 0xb0, 0x0c, 0x15, 0xf3, 0x3d, 0x92, 0x30, 0x86, 0x96, 0x6e, 0xd8, 0x1a, 0x31,
	0xfa, 0x63, 0x47, 0x23, 0xc4, 0x24, 0xa8, 0x82, 0x0f, 0xa0, 0x61, 0x98, 0xb6, 0x43, 0xb4, 0xfe,
	0xf0, 0x33, 0xaa, 0xe2, 0x26, 0xd4, 0x0d, 0xd3, 0x21, 0xa6, 0x69, 0xa3, 0x1a, 0x7e, 0x02, 0xcd,
	0xf8, 0x43, 0xb3, 0xcc, 0x29, 0xb9, 0xd6, 0xd0, 0x5e, 0xd2, 0xe0, 0x63, 0x7f, 0xac, 0x0f, 0x9d,
	0x49, 0x7f, 0xa4, 0x1b, 0x23, 0x24, 0xc7, 0x72, 0x19, 0x66, 0x99, 0xc4, 0x46, 0x75, 0xdc, 0x02,
	0xc8, 0x10, 0x7d, 0x88, 0xf6, 0x95, 0x4b, 0x38, 0x2e, 0xdc, 0xd6, 0xc8, 0xe7, 0x83, 0xef, 0xba,
	0x97, 0x25, 0xe1, 0x18, 0xf6, 0xb3, 0x85, 0xa7, 0x51, 0xa8, 0xa7, 0xfb, 0x56, 0xae, 0xe0, 0xa4,
	0x5c, 0x67, 0x44, 0xab, 0xac, 0xf0, 0x3e, 0x29, 0x61, 0xb4, 0x12, 0x95, 0xb5, 0x34, 0x29, 0x46,
	0xb4, 0x52, 0x3e, 0xc0, 0xab, 0x72, 0xa9, 0xbd, 0x71, 0x43, 0xe6, 0xce, 0x79, 0x40, 0xc3, 0x5c,
	0xfd, 0x35, 0xb4, 0x78, 0x8e, 0xe7, 0x33, 0x1c, 0xf0, 0x22, 0x5b, 0x79, 0x07, 0xa7, 0xe5, 0x76,
	0x03, 0x97, 0xcf, 0x17, 0xdb, 0x36, 0x62, 0xa4, 0x68, 0x23, 0x61, 0x28, 0xbf, 0xb6, 0x7f, 0x83,
	0x91, 0x9f, 0xc7, 0xea, 0xb2, 0x14, 0xab, 0x53, 0x75, 0x37, 0xb1, 0x9c, 0xaa, 0x17, 0xb0, 0x27,
	0xac, 0xb6, 0x2b, 0x22, 0x1b, 0x59, 0xa8, 0x12, 0x50, 0xf9, 0xf4, 0x9f, 0x51, 0x78, 0xb0, 0xee,
	0xd2, 0x22, 0x07, 0x67, 0x70, 0x98, 0xbd, 0x0d, 0x2a, 0xf3, 0xd2, 0xf7, 0x62, 0x16, 0xdd, 0x4c,
	0xa4, 0x2f, 0xa8, 0xfc, 0x0c, 0xcd, 0x64, 0x71, 0xfa, 0xe6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc2, 0x95, 0x07, 0x5d, 0xb7, 0x04, 0x00, 0x00,
}