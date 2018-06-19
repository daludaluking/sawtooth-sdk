// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/network_pb2/network.proto

package network_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GossipMessage_ContentType int32

const (
	GossipMessage_CONTENT_TYPE_UNSET GossipMessage_ContentType = 0
	GossipMessage_BLOCK              GossipMessage_ContentType = 1
	GossipMessage_BATCH              GossipMessage_ContentType = 2
)

var GossipMessage_ContentType_name = map[int32]string{
	0: "CONTENT_TYPE_UNSET",
	1: "BLOCK",
	2: "BATCH",
}
var GossipMessage_ContentType_value = map[string]int32{
	"CONTENT_TYPE_UNSET": 0,
	"BLOCK":              1,
	"BATCH":              2,
}

func (x GossipMessage_ContentType) String() string {
	return proto.EnumName(GossipMessage_ContentType_name, int32(x))
}
func (GossipMessage_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{7, 0}
}

type NetworkAcknowledgement_Status int32

const (
	NetworkAcknowledgement_STATUS_UNSET NetworkAcknowledgement_Status = 0
	NetworkAcknowledgement_OK           NetworkAcknowledgement_Status = 1
	NetworkAcknowledgement_ERROR        NetworkAcknowledgement_Status = 2
)

var NetworkAcknowledgement_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "ERROR",
}
var NetworkAcknowledgement_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"ERROR":        2,
}

func (x NetworkAcknowledgement_Status) String() string {
	return proto.EnumName(NetworkAcknowledgement_Status_name, int32(x))
}
func (NetworkAcknowledgement_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{8, 0}
}

// The disconnect message from a client to the server
type DisconnectMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectMessage) Reset()         { *m = DisconnectMessage{} }
func (m *DisconnectMessage) String() string { return proto.CompactTextString(m) }
func (*DisconnectMessage) ProtoMessage()    {}
func (*DisconnectMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{0}
}
func (m *DisconnectMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectMessage.Unmarshal(m, b)
}
func (m *DisconnectMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectMessage.Marshal(b, m, deterministic)
}
func (dst *DisconnectMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectMessage.Merge(dst, src)
}
func (m *DisconnectMessage) XXX_Size() int {
	return xxx_messageInfo_DisconnectMessage.Size(m)
}
func (m *DisconnectMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectMessage proto.InternalMessageInfo

// The registration request from a peer to the validator
type PeerRegisterRequest struct {
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	// The current version of the network protocol that is being used by the
	// sender.  This version is an increasing value.
	ProtocolVersion      uint32   `protobuf:"varint,2,opt,name=protocol_version,json=protocolVersion" json:"protocol_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerRegisterRequest) Reset()         { *m = PeerRegisterRequest{} }
func (m *PeerRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*PeerRegisterRequest) ProtoMessage()    {}
func (*PeerRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{1}
}
func (m *PeerRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerRegisterRequest.Unmarshal(m, b)
}
func (m *PeerRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerRegisterRequest.Marshal(b, m, deterministic)
}
func (dst *PeerRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerRegisterRequest.Merge(dst, src)
}
func (m *PeerRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_PeerRegisterRequest.Size(m)
}
func (m *PeerRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeerRegisterRequest proto.InternalMessageInfo

func (m *PeerRegisterRequest) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *PeerRegisterRequest) GetProtocolVersion() uint32 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

// The unregistration request from a peer to the validator
type PeerUnregisterRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerUnregisterRequest) Reset()         { *m = PeerUnregisterRequest{} }
func (m *PeerUnregisterRequest) String() string { return proto.CompactTextString(m) }
func (*PeerUnregisterRequest) ProtoMessage()    {}
func (*PeerUnregisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{2}
}
func (m *PeerUnregisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerUnregisterRequest.Unmarshal(m, b)
}
func (m *PeerUnregisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerUnregisterRequest.Marshal(b, m, deterministic)
}
func (dst *PeerUnregisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerUnregisterRequest.Merge(dst, src)
}
func (m *PeerUnregisterRequest) XXX_Size() int {
	return xxx_messageInfo_PeerUnregisterRequest.Size(m)
}
func (m *PeerUnregisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerUnregisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeerUnregisterRequest proto.InternalMessageInfo

type GetPeersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPeersRequest) Reset()         { *m = GetPeersRequest{} }
func (m *GetPeersRequest) String() string { return proto.CompactTextString(m) }
func (*GetPeersRequest) ProtoMessage()    {}
func (*GetPeersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{3}
}
func (m *GetPeersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPeersRequest.Unmarshal(m, b)
}
func (m *GetPeersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPeersRequest.Marshal(b, m, deterministic)
}
func (dst *GetPeersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPeersRequest.Merge(dst, src)
}
func (m *GetPeersRequest) XXX_Size() int {
	return xxx_messageInfo_GetPeersRequest.Size(m)
}
func (m *GetPeersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPeersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPeersRequest proto.InternalMessageInfo

type GetPeersResponse struct {
	PeerEndpoints        []string `protobuf:"bytes,1,rep,name=peer_endpoints,json=peerEndpoints" json:"peer_endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPeersResponse) Reset()         { *m = GetPeersResponse{} }
func (m *GetPeersResponse) String() string { return proto.CompactTextString(m) }
func (*GetPeersResponse) ProtoMessage()    {}
func (*GetPeersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{4}
}
func (m *GetPeersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPeersResponse.Unmarshal(m, b)
}
func (m *GetPeersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPeersResponse.Marshal(b, m, deterministic)
}
func (dst *GetPeersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPeersResponse.Merge(dst, src)
}
func (m *GetPeersResponse) XXX_Size() int {
	return xxx_messageInfo_GetPeersResponse.Size(m)
}
func (m *GetPeersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPeersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPeersResponse proto.InternalMessageInfo

func (m *GetPeersResponse) GetPeerEndpoints() []string {
	if m != nil {
		return m.PeerEndpoints
	}
	return nil
}

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{5}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{6}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (dst *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(dst, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

type GossipMessage struct {
	Content              []byte                    `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	ContentType          GossipMessage_ContentType `protobuf:"varint,2,opt,name=content_type,json=contentType,enum=GossipMessage_ContentType" json:"content_type,omitempty"`
	TimeToLive           uint32                    `protobuf:"varint,3,opt,name=time_to_live,json=timeToLive" json:"time_to_live,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GossipMessage) Reset()         { *m = GossipMessage{} }
func (m *GossipMessage) String() string { return proto.CompactTextString(m) }
func (*GossipMessage) ProtoMessage()    {}
func (*GossipMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{7}
}
func (m *GossipMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GossipMessage.Unmarshal(m, b)
}
func (m *GossipMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GossipMessage.Marshal(b, m, deterministic)
}
func (dst *GossipMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipMessage.Merge(dst, src)
}
func (m *GossipMessage) XXX_Size() int {
	return xxx_messageInfo_GossipMessage.Size(m)
}
func (m *GossipMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GossipMessage proto.InternalMessageInfo

func (m *GossipMessage) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *GossipMessage) GetContentType() GossipMessage_ContentType {
	if m != nil {
		return m.ContentType
	}
	return GossipMessage_CONTENT_TYPE_UNSET
}

func (m *GossipMessage) GetTimeToLive() uint32 {
	if m != nil {
		return m.TimeToLive
	}
	return 0
}

// A response sent from the validator to the peer acknowledging message
// receipt
type NetworkAcknowledgement struct {
	Status               NetworkAcknowledgement_Status `protobuf:"varint,1,opt,name=status,enum=NetworkAcknowledgement_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *NetworkAcknowledgement) Reset()         { *m = NetworkAcknowledgement{} }
func (m *NetworkAcknowledgement) String() string { return proto.CompactTextString(m) }
func (*NetworkAcknowledgement) ProtoMessage()    {}
func (*NetworkAcknowledgement) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{8}
}
func (m *NetworkAcknowledgement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkAcknowledgement.Unmarshal(m, b)
}
func (m *NetworkAcknowledgement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkAcknowledgement.Marshal(b, m, deterministic)
}
func (dst *NetworkAcknowledgement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkAcknowledgement.Merge(dst, src)
}
func (m *NetworkAcknowledgement) XXX_Size() int {
	return xxx_messageInfo_NetworkAcknowledgement.Size(m)
}
func (m *NetworkAcknowledgement) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkAcknowledgement.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkAcknowledgement proto.InternalMessageInfo

func (m *NetworkAcknowledgement) GetStatus() NetworkAcknowledgement_Status {
	if m != nil {
		return m.Status
	}
	return NetworkAcknowledgement_STATUS_UNSET
}

type GossipBlockRequest struct {
	// The id of the block that is being requested
	BlockId string `protobuf:"bytes,1,opt,name=block_id,json=blockId" json:"block_id,omitempty"`
	// A random string that provides uniqueness for requests with
	// otherwise identical fields.
	Nonce                string   `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`
	TimeToLive           uint32   `protobuf:"varint,3,opt,name=time_to_live,json=timeToLive" json:"time_to_live,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GossipBlockRequest) Reset()         { *m = GossipBlockRequest{} }
func (m *GossipBlockRequest) String() string { return proto.CompactTextString(m) }
func (*GossipBlockRequest) ProtoMessage()    {}
func (*GossipBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{9}
}
func (m *GossipBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GossipBlockRequest.Unmarshal(m, b)
}
func (m *GossipBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GossipBlockRequest.Marshal(b, m, deterministic)
}
func (dst *GossipBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipBlockRequest.Merge(dst, src)
}
func (m *GossipBlockRequest) XXX_Size() int {
	return xxx_messageInfo_GossipBlockRequest.Size(m)
}
func (m *GossipBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GossipBlockRequest proto.InternalMessageInfo

func (m *GossipBlockRequest) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

func (m *GossipBlockRequest) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *GossipBlockRequest) GetTimeToLive() uint32 {
	if m != nil {
		return m.TimeToLive
	}
	return 0
}

type GossipBlockResponse struct {
	// The block
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GossipBlockResponse) Reset()         { *m = GossipBlockResponse{} }
func (m *GossipBlockResponse) String() string { return proto.CompactTextString(m) }
func (*GossipBlockResponse) ProtoMessage()    {}
func (*GossipBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{10}
}
func (m *GossipBlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GossipBlockResponse.Unmarshal(m, b)
}
func (m *GossipBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GossipBlockResponse.Marshal(b, m, deterministic)
}
func (dst *GossipBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipBlockResponse.Merge(dst, src)
}
func (m *GossipBlockResponse) XXX_Size() int {
	return xxx_messageInfo_GossipBlockResponse.Size(m)
}
func (m *GossipBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GossipBlockResponse proto.InternalMessageInfo

func (m *GossipBlockResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type GossipBatchResponse struct {
	// The batch
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GossipBatchResponse) Reset()         { *m = GossipBatchResponse{} }
func (m *GossipBatchResponse) String() string { return proto.CompactTextString(m) }
func (*GossipBatchResponse) ProtoMessage()    {}
func (*GossipBatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{11}
}
func (m *GossipBatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GossipBatchResponse.Unmarshal(m, b)
}
func (m *GossipBatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GossipBatchResponse.Marshal(b, m, deterministic)
}
func (dst *GossipBatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipBatchResponse.Merge(dst, src)
}
func (m *GossipBatchResponse) XXX_Size() int {
	return xxx_messageInfo_GossipBatchResponse.Size(m)
}
func (m *GossipBatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipBatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GossipBatchResponse proto.InternalMessageInfo

func (m *GossipBatchResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type GossipBatchByBatchIdRequest struct {
	// The id of the batch that is being requested
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// A random string that provides uniqueness for requests with
	// otherwise identical fields.
	Nonce                string   `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`
	TimeToLive           uint32   `protobuf:"varint,3,opt,name=time_to_live,json=timeToLive" json:"time_to_live,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GossipBatchByBatchIdRequest) Reset()         { *m = GossipBatchByBatchIdRequest{} }
func (m *GossipBatchByBatchIdRequest) String() string { return proto.CompactTextString(m) }
func (*GossipBatchByBatchIdRequest) ProtoMessage()    {}
func (*GossipBatchByBatchIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{12}
}
func (m *GossipBatchByBatchIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GossipBatchByBatchIdRequest.Unmarshal(m, b)
}
func (m *GossipBatchByBatchIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GossipBatchByBatchIdRequest.Marshal(b, m, deterministic)
}
func (dst *GossipBatchByBatchIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipBatchByBatchIdRequest.Merge(dst, src)
}
func (m *GossipBatchByBatchIdRequest) XXX_Size() int {
	return xxx_messageInfo_GossipBatchByBatchIdRequest.Size(m)
}
func (m *GossipBatchByBatchIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipBatchByBatchIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GossipBatchByBatchIdRequest proto.InternalMessageInfo

func (m *GossipBatchByBatchIdRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GossipBatchByBatchIdRequest) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *GossipBatchByBatchIdRequest) GetTimeToLive() uint32 {
	if m != nil {
		return m.TimeToLive
	}
	return 0
}

type GossipBatchByTransactionIdRequest struct {
	// The id's of the transaction that are in the batches requested
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	// A random string that provides uniqueness for requests with
	// otherwise identical fields.
	Nonce                string   `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`
	TimeToLive           uint32   `protobuf:"varint,3,opt,name=time_to_live,json=timeToLive" json:"time_to_live,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GossipBatchByTransactionIdRequest) Reset()         { *m = GossipBatchByTransactionIdRequest{} }
func (m *GossipBatchByTransactionIdRequest) String() string { return proto.CompactTextString(m) }
func (*GossipBatchByTransactionIdRequest) ProtoMessage()    {}
func (*GossipBatchByTransactionIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_network_a64939f5a4ea11c7, []int{13}
}
func (m *GossipBatchByTransactionIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GossipBatchByTransactionIdRequest.Unmarshal(m, b)
}
func (m *GossipBatchByTransactionIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GossipBatchByTransactionIdRequest.Marshal(b, m, deterministic)
}
func (dst *GossipBatchByTransactionIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipBatchByTransactionIdRequest.Merge(dst, src)
}
func (m *GossipBatchByTransactionIdRequest) XXX_Size() int {
	return xxx_messageInfo_GossipBatchByTransactionIdRequest.Size(m)
}
func (m *GossipBatchByTransactionIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipBatchByTransactionIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GossipBatchByTransactionIdRequest proto.InternalMessageInfo

func (m *GossipBatchByTransactionIdRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *GossipBatchByTransactionIdRequest) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *GossipBatchByTransactionIdRequest) GetTimeToLive() uint32 {
	if m != nil {
		return m.TimeToLive
	}
	return 0
}

func init() {
	proto.RegisterType((*DisconnectMessage)(nil), "DisconnectMessage")
	proto.RegisterType((*PeerRegisterRequest)(nil), "PeerRegisterRequest")
	proto.RegisterType((*PeerUnregisterRequest)(nil), "PeerUnregisterRequest")
	proto.RegisterType((*GetPeersRequest)(nil), "GetPeersRequest")
	proto.RegisterType((*GetPeersResponse)(nil), "GetPeersResponse")
	proto.RegisterType((*PingRequest)(nil), "PingRequest")
	proto.RegisterType((*PingResponse)(nil), "PingResponse")
	proto.RegisterType((*GossipMessage)(nil), "GossipMessage")
	proto.RegisterType((*NetworkAcknowledgement)(nil), "NetworkAcknowledgement")
	proto.RegisterType((*GossipBlockRequest)(nil), "GossipBlockRequest")
	proto.RegisterType((*GossipBlockResponse)(nil), "GossipBlockResponse")
	proto.RegisterType((*GossipBatchResponse)(nil), "GossipBatchResponse")
	proto.RegisterType((*GossipBatchByBatchIdRequest)(nil), "GossipBatchByBatchIdRequest")
	proto.RegisterType((*GossipBatchByTransactionIdRequest)(nil), "GossipBatchByTransactionIdRequest")
	proto.RegisterEnum("GossipMessage_ContentType", GossipMessage_ContentType_name, GossipMessage_ContentType_value)
	proto.RegisterEnum("NetworkAcknowledgement_Status", NetworkAcknowledgement_Status_name, NetworkAcknowledgement_Status_value)
}

func init() {
	proto.RegisterFile("github.com/daludaluking/sawtooth-sdk/protobuf/network_pb2/network.proto", fileDescriptor_network_a64939f5a4ea11c7)
}

var fileDescriptor_network_a64939f5a4ea11c7 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5b, 0x8b, 0xd3, 0x40,
	0x14, 0x36, 0x29, 0xdb, 0xdd, 0x9e, 0x5e, 0x36, 0x3b, 0x75, 0xd7, 0xba, 0x82, 0xd4, 0x80, 0x52,
	0x1f, 0x4c, 0xa1, 0x82, 0x20, 0xe2, 0x43, 0x5b, 0xcb, 0xba, 0xb8, 0xb6, 0x25, 0x4d, 0x05, 0x45,
	0x08, 0x69, 0x72, 0xb6, 0x3b, 0xb4, 0x9d, 0x89, 0x99, 0x69, 0x4b, 0x9f, 0xfc, 0x79, 0xfe, 0x2d,
	0xc9, 0x65, 0x7a, 0x01, 0x85, 0x65, 0x9f, 0x72, 0xbe, 0x6f, 0xce, 0xed, 0x3b, 0x73, 0x32, 0xd0,
	0x14, 0xde, 0x5a, 0x72, 0x2e, 0xef, 0x5c, 0x11, 0xcc, 0x9a, 0x61, 0xc4, 0x25, 0x9f, 0x2c, 0x6f,
	0x9b, 0x0c, 0xe5, 0x9a, 0x47, 0x33, 0x37, 0x9c, 0xb4, 0x94, 0x6d, 0x25, 0x87, 0x66, 0x15, 0xce,
	0x3e, 0x51, 0xe1, 0x73, 0xc6, 0xd0, 0x97, 0x5f, 0x51, 0x08, 0x6f, 0x8a, 0xe6, 0x4f, 0xa8, 0x0e,
	0x11, 0x23, 0x1b, 0xa7, 0x54, 0xc8, 0xf8, 0xfb, 0x6b, 0x89, 0x42, 0x92, 0x4b, 0x38, 0x41, 0x16,
	0x84, 0x9c, 0x32, 0x59, 0xd3, 0xea, 0x5a, 0xa3, 0x60, 0x6f, 0x31, 0x79, 0x0d, 0x46, 0x92, 0xd0,
	0xe7, 0x73, 0x77, 0x85, 0x91, 0xa0, 0x9c, 0xd5, 0xf4, 0xba, 0xd6, 0x28, 0xdb, 0xa7, 0x8a, 0xff,
	0x96, 0xd2, 0xe6, 0x13, 0x38, 0x8f, 0xb3, 0x8f, 0x59, 0x74, 0x98, 0xdf, 0x3c, 0x83, 0xd3, 0x2b,
	0x94, 0xf1, 0x99, 0x50, 0xd4, 0x7b, 0x30, 0x76, 0x94, 0x08, 0x39, 0x13, 0x48, 0x5e, 0x42, 0x25,
	0x44, 0x8c, 0x5c, 0x55, 0x5b, 0xd4, 0xb4, 0x7a, 0xae, 0x51, 0xb0, 0xcb, 0x31, 0xdb, 0x53, 0xa4,
	0x59, 0x86, 0xe2, 0x90, 0xb2, 0xa9, 0xca, 0x54, 0x81, 0x52, 0x0a, 0xd3, 0x2c, 0xe6, 0x1f, 0x0d,
	0xca, 0x57, 0x5c, 0x08, 0x1a, 0x66, 0xaa, 0x49, 0x0d, 0x8e, 0x7d, 0xce, 0x24, 0x66, 0xea, 0x4a,
	0xb6, 0x82, 0xe4, 0x23, 0x94, 0x32, 0xd3, 0x95, 0x9b, 0x10, 0x13, 0x61, 0x95, 0xd6, 0xa5, 0x75,
	0x10, 0x6f, 0x75, 0x53, 0x17, 0x67, 0x13, 0xa2, 0x5d, 0xf4, 0x77, 0x80, 0xd4, 0xa1, 0x24, 0xe9,
	0x02, 0x5d, 0xc9, 0xdd, 0x39, 0x5d, 0x61, 0x2d, 0x97, 0xcc, 0x05, 0x62, 0xce, 0xe1, 0x37, 0x74,
	0x85, 0xe6, 0x07, 0x28, 0xee, 0x45, 0x93, 0x0b, 0x20, 0xdd, 0x41, 0xdf, 0xe9, 0xf5, 0x1d, 0xd7,
	0xf9, 0x3e, 0xec, 0xb9, 0xe3, 0xfe, 0xa8, 0xe7, 0x18, 0x8f, 0x48, 0x01, 0x8e, 0x3a, 0x37, 0x83,
	0xee, 0x17, 0x43, 0x4b, 0xcc, 0xb6, 0xd3, 0xfd, 0x6c, 0xe8, 0xe6, 0x6f, 0xb8, 0xe8, 0xa7, 0x77,
	0xda, 0xf6, 0x67, 0x8c, 0xaf, 0xe7, 0x18, 0x4c, 0x71, 0x11, 0xf7, 0xfd, 0x0e, 0xf2, 0x42, 0x7a,
	0x72, 0x29, 0x12, 0x41, 0x95, 0xd6, 0x73, 0xeb, 0xdf, 0x8e, 0xd6, 0x28, 0xf1, 0xb2, 0x33, 0x6f,
	0xf3, 0x0d, 0xe4, 0x53, 0x86, 0x18, 0x50, 0x1a, 0x39, 0x6d, 0x67, 0x3c, 0xda, 0xf6, 0x90, 0x07,
	0x7d, 0x90, 0x35, 0xd0, 0xb3, 0xed, 0x81, 0x6d, 0xe8, 0xe6, 0x14, 0x48, 0x3a, 0x89, 0xce, 0x9c,
	0xfb, 0x33, 0xb5, 0x2d, 0x4f, 0xe1, 0x64, 0x12, 0x63, 0x97, 0x06, 0xd9, 0xb6, 0x1c, 0x27, 0xf8,
	0x3a, 0x20, 0x8f, 0xe1, 0x88, 0x71, 0xe6, 0xa7, 0x83, 0x2c, 0xd8, 0x29, 0xb8, 0xc7, 0x98, 0x9a,
	0x50, 0x3d, 0x28, 0x94, 0x2d, 0xc4, 0x7f, 0x2f, 0x6e, 0x2f, 0xc0, 0x93, 0xfe, 0xdd, 0x3d, 0x02,
	0x10, 0x9e, 0xed, 0x05, 0x74, 0x36, 0xc9, 0xe7, 0x3a, 0x50, 0x9a, 0x2a, 0xa0, 0x6f, 0xd5, 0xe8,
	0xf4, 0xe1, 0x42, 0x16, 0xf0, 0xe2, 0xa0, 0x8c, 0x13, 0x79, 0x4c, 0x78, 0xbe, 0xa4, 0x9c, 0xed,
	0x8a, 0x19, 0x90, 0xa3, 0x81, 0x5a, 0xee, 0xd8, 0x7c, 0x68, 0xb9, 0xce, 0x2b, 0x38, 0x57, 0xef,
	0x82, 0x25, 0x82, 0xec, 0xd7, 0x9f, 0x2c, 0x6f, 0x87, 0xda, 0x8f, 0xe2, 0xde, 0xd3, 0x30, 0xc9,
	0x27, 0x07, 0x6f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x78, 0x6a, 0x4c, 0x60, 0x46, 0x04, 0x00,
	0x00,
}
