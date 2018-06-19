// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth-sdk/protobuf/client_status_pb2/client_status.proto

package client_status

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

// The status of the response message, not the validator's status
type ClientStatusGetResponse_Status int32

const (
	ClientStatusGetResponse_STATUS_UNSET ClientStatusGetResponse_Status = 0
	ClientStatusGetResponse_OK           ClientStatusGetResponse_Status = 1
	ClientStatusGetResponse_ERROR        ClientStatusGetResponse_Status = 2
)

var ClientStatusGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "ERROR",
}
var ClientStatusGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET": 0,
	"OK":           1,
	"ERROR":        2,
}

func (x ClientStatusGetResponse_Status) String() string {
	return proto.EnumName(ClientStatusGetResponse_Status_name, int32(x))
}
func (ClientStatusGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_status_9ab2e4c7ff7bcd72, []int{1, 0}
}

// A request to get miscellaneous information about the validator
type ClientStatusGetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStatusGetRequest) Reset()         { *m = ClientStatusGetRequest{} }
func (m *ClientStatusGetRequest) String() string { return proto.CompactTextString(m) }
func (*ClientStatusGetRequest) ProtoMessage()    {}
func (*ClientStatusGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_status_9ab2e4c7ff7bcd72, []int{0}
}
func (m *ClientStatusGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStatusGetRequest.Unmarshal(m, b)
}
func (m *ClientStatusGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStatusGetRequest.Marshal(b, m, deterministic)
}
func (dst *ClientStatusGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStatusGetRequest.Merge(dst, src)
}
func (m *ClientStatusGetRequest) XXX_Size() int {
	return xxx_messageInfo_ClientStatusGetRequest.Size(m)
}
func (m *ClientStatusGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStatusGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStatusGetRequest proto.InternalMessageInfo

type ClientStatusGetResponse struct {
	Status ClientStatusGetResponse_Status  `protobuf:"varint,1,opt,name=status,enum=ClientStatusGetResponse_Status" json:"status,omitempty"`
	Peers  []*ClientStatusGetResponse_Peer `protobuf:"bytes,2,rep,name=peers" json:"peers,omitempty"`
	// The validator's public network endpoint
	Endpoint             string   `protobuf:"bytes,3,opt,name=endpoint" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStatusGetResponse) Reset()         { *m = ClientStatusGetResponse{} }
func (m *ClientStatusGetResponse) String() string { return proto.CompactTextString(m) }
func (*ClientStatusGetResponse) ProtoMessage()    {}
func (*ClientStatusGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_status_9ab2e4c7ff7bcd72, []int{1}
}
func (m *ClientStatusGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStatusGetResponse.Unmarshal(m, b)
}
func (m *ClientStatusGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStatusGetResponse.Marshal(b, m, deterministic)
}
func (dst *ClientStatusGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStatusGetResponse.Merge(dst, src)
}
func (m *ClientStatusGetResponse) XXX_Size() int {
	return xxx_messageInfo_ClientStatusGetResponse.Size(m)
}
func (m *ClientStatusGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStatusGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStatusGetResponse proto.InternalMessageInfo

func (m *ClientStatusGetResponse) GetStatus() ClientStatusGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientStatusGetResponse_STATUS_UNSET
}

func (m *ClientStatusGetResponse) GetPeers() []*ClientStatusGetResponse_Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *ClientStatusGetResponse) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

// Information about the validator's peers
type ClientStatusGetResponse_Peer struct {
	// The peer's public network endpoint
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientStatusGetResponse_Peer) Reset()         { *m = ClientStatusGetResponse_Peer{} }
func (m *ClientStatusGetResponse_Peer) String() string { return proto.CompactTextString(m) }
func (*ClientStatusGetResponse_Peer) ProtoMessage()    {}
func (*ClientStatusGetResponse_Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_status_9ab2e4c7ff7bcd72, []int{1, 0}
}
func (m *ClientStatusGetResponse_Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientStatusGetResponse_Peer.Unmarshal(m, b)
}
func (m *ClientStatusGetResponse_Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientStatusGetResponse_Peer.Marshal(b, m, deterministic)
}
func (dst *ClientStatusGetResponse_Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientStatusGetResponse_Peer.Merge(dst, src)
}
func (m *ClientStatusGetResponse_Peer) XXX_Size() int {
	return xxx_messageInfo_ClientStatusGetResponse_Peer.Size(m)
}
func (m *ClientStatusGetResponse_Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientStatusGetResponse_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_ClientStatusGetResponse_Peer proto.InternalMessageInfo

func (m *ClientStatusGetResponse_Peer) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientStatusGetRequest)(nil), "ClientStatusGetRequest")
	proto.RegisterType((*ClientStatusGetResponse)(nil), "ClientStatusGetResponse")
	proto.RegisterType((*ClientStatusGetResponse_Peer)(nil), "ClientStatusGetResponse.Peer")
	proto.RegisterEnum("ClientStatusGetResponse_Status", ClientStatusGetResponse_Status_name, ClientStatusGetResponse_Status_value)
}

func init() {
	proto.RegisterFile("github.com/daludaluking/sawtooth-sdk/protobuf/client_status_pb2/client_status.proto", fileDescriptor_client_status_9ab2e4c7ff7bcd72)
}

var fileDescriptor_client_status_9ab2e4c7ff7bcd72 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0xd4, 0x06, 0x3b, 0xfe, 0x20, 0x2c, 0xa8, 0xa1, 0x20, 0x86, 0x9c, 0x72, 0x71,
	0x0b, 0xe9, 0xc1, 0x83, 0x27, 0x95, 0xe2, 0x41, 0xb0, 0x65, 0x93, 0x5e, 0xbc, 0x84, 0xc6, 0x8c,
	0x58, 0x2a, 0xd9, 0x35, 0x33, 0xc1, 0xff, 0xdc, 0xb3, 0x98, 0x58, 0x61, 0x91, 0x1e, 0xe7, 0x7b,
	0xdf, 0xcc, 0xc0, 0x83, 0x1b, 0x5a, 0x7d, 0xb2, 0x31, 0xfc, 0x56, 0x50, 0xb5, 0x99, 0xd8, 0xc6,
	0xb0, 0x29, 0xdb, 0xd7, 0xc9, 0xcb, 0xfb, 0x1a, 0x6b, 0x2e, 0x88, 0x57, 0xdc, 0x52, 0x61, 0xcb,
	0xd4, 0x25, 0xaa, 0x13, 0xe3, 0x10, 0xce, 0xee, 0x3b, 0x9c, 0x75, 0xf4, 0x01, 0x59, 0xe3, 0x47,
	0x8b, 0xc4, 0xf1, 0x97, 0x80, 0xf3, 0x7f, 0x11, 0x59, 0x53, 0x13, 0xca, 0x6b, 0xf0, 0xfb, 0x2b,
	0xa1, 0x88, 0x44, 0x72, 0x92, 0x5e, 0xaa, 0x1d, 0xa6, 0xea, 0x89, 0xfe, 0xd5, 0xe5, 0x14, 0x86,
	0x16, 0xb1, 0xa1, 0xd0, 0x8b, 0x06, 0xc9, 0x61, 0x7a, 0xb1, 0x73, 0x6f, 0x81, 0xd8, 0xe8, 0xde,
	0x95, 0x63, 0x38, 0xc0, 0xba, 0xb2, 0x66, 0x5d, 0x73, 0x38, 0x88, 0x44, 0x32, 0xd2, 0x7f, 0xf3,
	0x38, 0x86, 0xfd, 0x1f, 0xd5, 0x71, 0x84, 0xeb, 0xc4, 0x57, 0xe0, 0xf7, 0x0f, 0x64, 0x00, 0x47,
	0x59, 0x7e, 0x9b, 0x2f, 0xb3, 0x62, 0xf9, 0x94, 0xcd, 0xf2, 0x60, 0x4f, 0xfa, 0xe0, 0xcd, 0x1f,
	0x03, 0x21, 0x47, 0x30, 0x9c, 0x69, 0x3d, 0xd7, 0x81, 0x77, 0x97, 0xc0, 0xe9, 0xb6, 0x51, 0x45,
	0xd5, 0x46, 0x6d, 0x1b, 0x5d, 0x88, 0xe7, 0x63, 0xa7, 0xc2, 0xd2, 0xef, 0xa2, 0xe9, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x22, 0x79, 0xa9, 0x15, 0x82, 0x01, 0x00, 0x00,
}
