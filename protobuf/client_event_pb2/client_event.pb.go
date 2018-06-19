// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/daludaluking/sawtooth_sdk/protobuf/client_event_pb2/client_event.proto

package client_event_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import events_pb2 "github.com/daludaluking/sawtooth_sdk/protobuf/events_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientEventsSubscribeResponse_Status int32

const (
	ClientEventsSubscribeResponse_STATUS_UNSET   ClientEventsSubscribeResponse_Status = 0
	ClientEventsSubscribeResponse_OK             ClientEventsSubscribeResponse_Status = 1
	ClientEventsSubscribeResponse_INVALID_FILTER ClientEventsSubscribeResponse_Status = 2
	ClientEventsSubscribeResponse_UNKNOWN_BLOCK  ClientEventsSubscribeResponse_Status = 3
)

var ClientEventsSubscribeResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INVALID_FILTER",
	3: "UNKNOWN_BLOCK",
}
var ClientEventsSubscribeResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INVALID_FILTER": 2,
	"UNKNOWN_BLOCK":  3,
}

func (x ClientEventsSubscribeResponse_Status) String() string {
	return proto.EnumName(ClientEventsSubscribeResponse_Status_name, int32(x))
}
func (ClientEventsSubscribeResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_event_bd0be6b41971519b, []int{1, 0}
}

type ClientEventsUnsubscribeResponse_Status int32

const (
	ClientEventsUnsubscribeResponse_STATUS_UNSET   ClientEventsUnsubscribeResponse_Status = 0
	ClientEventsUnsubscribeResponse_OK             ClientEventsUnsubscribeResponse_Status = 1
	ClientEventsUnsubscribeResponse_INTERNAL_ERROR ClientEventsUnsubscribeResponse_Status = 2
)

var ClientEventsUnsubscribeResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
}
var ClientEventsUnsubscribeResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
}

func (x ClientEventsUnsubscribeResponse_Status) String() string {
	return proto.EnumName(ClientEventsUnsubscribeResponse_Status_name, int32(x))
}
func (ClientEventsUnsubscribeResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_event_bd0be6b41971519b, []int{3, 0}
}

type ClientEventsGetResponse_Status int32

const (
	ClientEventsGetResponse_STATUS_UNSET   ClientEventsGetResponse_Status = 0
	ClientEventsGetResponse_OK             ClientEventsGetResponse_Status = 1
	ClientEventsGetResponse_INTERNAL_ERROR ClientEventsGetResponse_Status = 2
	ClientEventsGetResponse_INVALID_FILTER ClientEventsGetResponse_Status = 3
	ClientEventsGetResponse_UNKNOWN_BLOCK  ClientEventsGetResponse_Status = 4
)

var ClientEventsGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "INVALID_FILTER",
	4: "UNKNOWN_BLOCK",
}
var ClientEventsGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET":   0,
	"OK":             1,
	"INTERNAL_ERROR": 2,
	"INVALID_FILTER": 3,
	"UNKNOWN_BLOCK":  4,
}

func (x ClientEventsGetResponse_Status) String() string {
	return proto.EnumName(ClientEventsGetResponse_Status_name, int32(x))
}
func (ClientEventsGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_client_event_bd0be6b41971519b, []int{5, 0}
}

type ClientEventsSubscribeRequest struct {
	Subscriptions []*events_pb2.EventSubscription `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
	// The block id (or ids, if trying to walk back a fork) the subscriber last
	// received events on. It can be set to empty if it has not yet received the
	// genesis block.
	LastKnownBlockIds    []string `protobuf:"bytes,2,rep,name=last_known_block_ids,json=lastKnownBlockIds" json:"last_known_block_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientEventsSubscribeRequest) Reset()         { *m = ClientEventsSubscribeRequest{} }
func (m *ClientEventsSubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*ClientEventsSubscribeRequest) ProtoMessage()    {}
func (*ClientEventsSubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_event_bd0be6b41971519b, []int{0}
}
func (m *ClientEventsSubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEventsSubscribeRequest.Unmarshal(m, b)
}
func (m *ClientEventsSubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEventsSubscribeRequest.Marshal(b, m, deterministic)
}
func (dst *ClientEventsSubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEventsSubscribeRequest.Merge(dst, src)
}
func (m *ClientEventsSubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_ClientEventsSubscribeRequest.Size(m)
}
func (m *ClientEventsSubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEventsSubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEventsSubscribeRequest proto.InternalMessageInfo

func (m *ClientEventsSubscribeRequest) GetSubscriptions() []*events_pb2.EventSubscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *ClientEventsSubscribeRequest) GetLastKnownBlockIds() []string {
	if m != nil {
		return m.LastKnownBlockIds
	}
	return nil
}

type ClientEventsSubscribeResponse struct {
	Status ClientEventsSubscribeResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientEventsSubscribeResponse_Status" json:"status,omitempty"`
	// Additional information about the response status
	ResponseMessage      string   `protobuf:"bytes,2,opt,name=response_message,json=responseMessage" json:"response_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientEventsSubscribeResponse) Reset()         { *m = ClientEventsSubscribeResponse{} }
func (m *ClientEventsSubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*ClientEventsSubscribeResponse) ProtoMessage()    {}
func (*ClientEventsSubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_event_bd0be6b41971519b, []int{1}
}
func (m *ClientEventsSubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEventsSubscribeResponse.Unmarshal(m, b)
}
func (m *ClientEventsSubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEventsSubscribeResponse.Marshal(b, m, deterministic)
}
func (dst *ClientEventsSubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEventsSubscribeResponse.Merge(dst, src)
}
func (m *ClientEventsSubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_ClientEventsSubscribeResponse.Size(m)
}
func (m *ClientEventsSubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEventsSubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEventsSubscribeResponse proto.InternalMessageInfo

func (m *ClientEventsSubscribeResponse) GetStatus() ClientEventsSubscribeResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientEventsSubscribeResponse_STATUS_UNSET
}

func (m *ClientEventsSubscribeResponse) GetResponseMessage() string {
	if m != nil {
		return m.ResponseMessage
	}
	return ""
}

type ClientEventsUnsubscribeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientEventsUnsubscribeRequest) Reset()         { *m = ClientEventsUnsubscribeRequest{} }
func (m *ClientEventsUnsubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*ClientEventsUnsubscribeRequest) ProtoMessage()    {}
func (*ClientEventsUnsubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_event_bd0be6b41971519b, []int{2}
}
func (m *ClientEventsUnsubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEventsUnsubscribeRequest.Unmarshal(m, b)
}
func (m *ClientEventsUnsubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEventsUnsubscribeRequest.Marshal(b, m, deterministic)
}
func (dst *ClientEventsUnsubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEventsUnsubscribeRequest.Merge(dst, src)
}
func (m *ClientEventsUnsubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_ClientEventsUnsubscribeRequest.Size(m)
}
func (m *ClientEventsUnsubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEventsUnsubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEventsUnsubscribeRequest proto.InternalMessageInfo

type ClientEventsUnsubscribeResponse struct {
	Status               ClientEventsUnsubscribeResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientEventsUnsubscribeResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *ClientEventsUnsubscribeResponse) Reset()         { *m = ClientEventsUnsubscribeResponse{} }
func (m *ClientEventsUnsubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*ClientEventsUnsubscribeResponse) ProtoMessage()    {}
func (*ClientEventsUnsubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_event_bd0be6b41971519b, []int{3}
}
func (m *ClientEventsUnsubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEventsUnsubscribeResponse.Unmarshal(m, b)
}
func (m *ClientEventsUnsubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEventsUnsubscribeResponse.Marshal(b, m, deterministic)
}
func (dst *ClientEventsUnsubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEventsUnsubscribeResponse.Merge(dst, src)
}
func (m *ClientEventsUnsubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_ClientEventsUnsubscribeResponse.Size(m)
}
func (m *ClientEventsUnsubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEventsUnsubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEventsUnsubscribeResponse proto.InternalMessageInfo

func (m *ClientEventsUnsubscribeResponse) GetStatus() ClientEventsUnsubscribeResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientEventsUnsubscribeResponse_STATUS_UNSET
}

type ClientEventsGetRequest struct {
	Subscriptions        []*events_pb2.EventSubscription `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
	BlockIds             []string                        `protobuf:"bytes,2,rep,name=block_ids,json=blockIds" json:"block_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ClientEventsGetRequest) Reset()         { *m = ClientEventsGetRequest{} }
func (m *ClientEventsGetRequest) String() string { return proto.CompactTextString(m) }
func (*ClientEventsGetRequest) ProtoMessage()    {}
func (*ClientEventsGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_event_bd0be6b41971519b, []int{4}
}
func (m *ClientEventsGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEventsGetRequest.Unmarshal(m, b)
}
func (m *ClientEventsGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEventsGetRequest.Marshal(b, m, deterministic)
}
func (dst *ClientEventsGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEventsGetRequest.Merge(dst, src)
}
func (m *ClientEventsGetRequest) XXX_Size() int {
	return xxx_messageInfo_ClientEventsGetRequest.Size(m)
}
func (m *ClientEventsGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEventsGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEventsGetRequest proto.InternalMessageInfo

func (m *ClientEventsGetRequest) GetSubscriptions() []*events_pb2.EventSubscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *ClientEventsGetRequest) GetBlockIds() []string {
	if m != nil {
		return m.BlockIds
	}
	return nil
}

type ClientEventsGetResponse struct {
	Status               ClientEventsGetResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientEventsGetResponse_Status" json:"status,omitempty"`
	Events               []*events_pb2.Event            `protobuf:"bytes,2,rep,name=events" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ClientEventsGetResponse) Reset()         { *m = ClientEventsGetResponse{} }
func (m *ClientEventsGetResponse) String() string { return proto.CompactTextString(m) }
func (*ClientEventsGetResponse) ProtoMessage()    {}
func (*ClientEventsGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_client_event_bd0be6b41971519b, []int{5}
}
func (m *ClientEventsGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientEventsGetResponse.Unmarshal(m, b)
}
func (m *ClientEventsGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientEventsGetResponse.Marshal(b, m, deterministic)
}
func (dst *ClientEventsGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientEventsGetResponse.Merge(dst, src)
}
func (m *ClientEventsGetResponse) XXX_Size() int {
	return xxx_messageInfo_ClientEventsGetResponse.Size(m)
}
func (m *ClientEventsGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientEventsGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientEventsGetResponse proto.InternalMessageInfo

func (m *ClientEventsGetResponse) GetStatus() ClientEventsGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientEventsGetResponse_STATUS_UNSET
}

func (m *ClientEventsGetResponse) GetEvents() []*events_pb2.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientEventsSubscribeRequest)(nil), "ClientEventsSubscribeRequest")
	proto.RegisterType((*ClientEventsSubscribeResponse)(nil), "ClientEventsSubscribeResponse")
	proto.RegisterType((*ClientEventsUnsubscribeRequest)(nil), "ClientEventsUnsubscribeRequest")
	proto.RegisterType((*ClientEventsUnsubscribeResponse)(nil), "ClientEventsUnsubscribeResponse")
	proto.RegisterType((*ClientEventsGetRequest)(nil), "ClientEventsGetRequest")
	proto.RegisterType((*ClientEventsGetResponse)(nil), "ClientEventsGetResponse")
	proto.RegisterEnum("ClientEventsSubscribeResponse_Status", ClientEventsSubscribeResponse_Status_name, ClientEventsSubscribeResponse_Status_value)
	proto.RegisterEnum("ClientEventsUnsubscribeResponse_Status", ClientEventsUnsubscribeResponse_Status_name, ClientEventsUnsubscribeResponse_Status_value)
	proto.RegisterEnum("ClientEventsGetResponse_Status", ClientEventsGetResponse_Status_name, ClientEventsGetResponse_Status_value)
}

func init() {
	proto.RegisterFile("github.com/daludaluking/sawtooth_sdk/protobuf/client_event_pb2/client_event.proto", fileDescriptor_client_event_bd0be6b41971519b)
}

var fileDescriptor_client_event_bd0be6b41971519b = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xeb, 0x8a, 0xd3, 0x40,
	0x14, 0x76, 0x52, 0x09, 0xf6, 0xac, 0xbb, 0x66, 0x07, 0x2f, 0xc5, 0xcb, 0x6e, 0x19, 0x10, 0x2b,
	0xe2, 0x14, 0x2a, 0x78, 0xf9, 0x21, 0xd2, 0xae, 0x51, 0x42, 0x6b, 0x2a, 0x93, 0x54, 0x41, 0x90,
	0xa1, 0x69, 0x47, 0x2d, 0xad, 0x99, 0xda, 0x33, 0x75, 0x9f, 0xc1, 0x57, 0xf0, 0xad, 0x04, 0x1f,
	0x48, 0x32, 0x49, 0xa0, 0x9b, 0xdd, 0xaa, 0xe8, 0x9f, 0x10, 0xbe, 0x4b, 0xce, 0x25, 0xdf, 0x81,
	0x27, 0x38, 0x3e, 0x36, 0x5a, 0x9b, 0x4f, 0x12, 0xa7, 0xf3, 0xf6, 0x72, 0xa5, 0x8d, 0x4e, 0xd6,
	0x1f, 0xda, 0x93, 0xc5, 0x4c, 0xa5, 0x46, 0xaa, 0xaf, 0xd9, 0x73, 0x99, 0x74, 0x4e, 0x00, 0xdc,
	0xca, 0xae, 0xdf, 0x3f, 0xdb, 0x6a, 0x25, 0x68, 0x4d, 0xf9, 0x6b, 0x2e, 0x67, 0xdf, 0x08, 0xdc,
	0x3c, 0xb2, 0x5f, 0xf1, 0x2d, 0x1c, 0xad, 0x13, 0x9c, 0xac, 0x66, 0x89, 0x12, 0xea, 0xcb, 0x5a,
	0xa1, 0xa1, 0x8f, 0x61, 0x17, 0x73, 0x6c, 0x69, 0x66, 0x3a, 0xc5, 0x06, 0x69, 0xd6, 0x5a, 0x3b,
	0x1d, 0xca, 0xad, 0x3e, 0xda, 0xa0, 0xc4, 0x49, 0x21, 0x6d, 0xc3, 0xe5, 0xc5, 0x18, 0x8d, 0x9c,
	0xa7, 0xfa, 0x38, 0x95, 0xc9, 0x42, 0x4f, 0xe6, 0x72, 0x36, 0xc5, 0x86, 0xd3, 0xac, 0xb5, 0xea,
	0x62, 0x3f, 0xe3, 0xfa, 0x19, 0xd5, 0xcb, 0x98, 0x60, 0x8a, 0xec, 0x27, 0x81, 0x5b, 0x5b, 0x7a,
	0xc1, 0xa5, 0x4e, 0x51, 0xd1, 0xa7, 0xe0, 0xa2, 0x19, 0x9b, 0x75, 0xd6, 0x05, 0x69, 0xed, 0x75,
	0x6e, 0xf3, 0xdf, 0xea, 0x79, 0x64, 0xc5, 0xa2, 0x30, 0xd1, 0xbb, 0xe0, 0xad, 0x0a, 0x4a, 0x7e,
	0x56, 0x88, 0xe3, 0x8f, 0xaa, 0xe1, 0x34, 0x49, 0xab, 0x2e, 0x2e, 0x95, 0xf8, 0xab, 0x1c, 0x66,
	0x01, 0xb8, 0xb9, 0x99, 0x7a, 0x70, 0x31, 0x8a, 0xbb, 0xf1, 0x28, 0x92, 0xa3, 0x30, 0xf2, 0x63,
	0xef, 0x1c, 0x75, 0xc1, 0x19, 0xf6, 0x3d, 0x42, 0x29, 0xec, 0x05, 0xe1, 0x9b, 0xee, 0x20, 0x78,
	0x2e, 0x5f, 0x04, 0x83, 0xd8, 0x17, 0x9e, 0x43, 0xf7, 0x61, 0x77, 0x14, 0xf6, 0xc3, 0xe1, 0xdb,
	0x50, 0xf6, 0x06, 0xc3, 0xa3, 0xbe, 0x57, 0x63, 0x4d, 0x38, 0xd8, 0xec, 0x72, 0x94, 0x62, 0x65,
	0xc7, 0xec, 0x3b, 0x81, 0xc3, 0xad, 0x92, 0x62, 0xf4, 0x67, 0x95, 0xd1, 0xef, 0xf0, 0x3f, 0x38,
	0x2a, 0xc3, 0xb3, 0x87, 0x7f, 0x3b, 0x51, 0xec, 0x8b, 0xb0, 0x3b, 0x90, 0xbe, 0x10, 0x43, 0xe1,
	0x39, 0x4c, 0xc3, 0xd5, 0xcd, 0x4a, 0x2f, 0x95, 0xf9, 0xff, 0x68, 0xdc, 0x80, 0x7a, 0x35, 0x0f,
	0x17, 0x92, 0x32, 0x06, 0x3f, 0x08, 0x5c, 0x3b, 0x55, 0xb1, 0xd8, 0xc2, 0xa3, 0xca, 0x16, 0x0e,
	0xf9, 0x16, 0x65, 0xf5, 0xd7, 0x1f, 0x80, 0x9b, 0xe7, 0xde, 0x96, 0xdb, 0xe9, 0xb8, 0x79, 0x93,
	0xa2, 0x40, 0xd9, 0xfb, 0x7f, 0xdb, 0xce, 0x19, 0x19, 0xa8, 0x9d, 0xce, 0xc0, 0xf9, 0xde, 0x3d,
	0xb8, 0x52, 0xde, 0x25, 0xc7, 0xe9, 0x9c, 0x97, 0x77, 0xf9, 0x9a, 0xbc, 0xf3, 0xaa, 0x57, 0x9d,
	0xb8, 0x96, 0x7d, 0xf0, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x01, 0x3c, 0x95, 0x77, 0x06, 0x04, 0x00,
	0x00,
}
