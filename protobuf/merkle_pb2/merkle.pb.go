// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/merkle_pb2/merkle.proto

package merkle_pb2

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

// An Entry in the change log for a given state root.
type ChangeLogEntry struct {
	// A root hash of a merkle trie this tree was based off.
	Parent []byte `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The hashes that were added for this root. These may be deleted during
	// pruning, if this root is being abandoned.
	Additions [][]byte `protobuf:"bytes,2,rep,name=additions,proto3" json:"additions,omitempty"`
	// The list of successors.
	Successors           []*ChangeLogEntry_Successor `protobuf:"bytes,3,rep,name=successors" json:"successors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ChangeLogEntry) Reset()         { *m = ChangeLogEntry{} }
func (m *ChangeLogEntry) String() string { return proto.CompactTextString(m) }
func (*ChangeLogEntry) ProtoMessage()    {}
func (*ChangeLogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_merkle_65730cd798e98cd8, []int{0}
}
func (m *ChangeLogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeLogEntry.Unmarshal(m, b)
}
func (m *ChangeLogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeLogEntry.Marshal(b, m, deterministic)
}
func (dst *ChangeLogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeLogEntry.Merge(dst, src)
}
func (m *ChangeLogEntry) XXX_Size() int {
	return xxx_messageInfo_ChangeLogEntry.Size(m)
}
func (m *ChangeLogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeLogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeLogEntry proto.InternalMessageInfo

func (m *ChangeLogEntry) GetParent() []byte {
	if m != nil {
		return m.Parent
	}
	return nil
}

func (m *ChangeLogEntry) GetAdditions() [][]byte {
	if m != nil {
		return m.Additions
	}
	return nil
}

func (m *ChangeLogEntry) GetSuccessors() []*ChangeLogEntry_Successor {
	if m != nil {
		return m.Successors
	}
	return nil
}

// A state root that succeed this root.
type ChangeLogEntry_Successor struct {
	// A root hash of a merkle trie based of off this root.
	Successor []byte `protobuf:"bytes,1,opt,name=successor,proto3" json:"successor,omitempty"`
	// The keys (i.e. hashes) that were replaced (i.e. deleted) by this
	// successor.  These may be deleted during pruning.
	Deletions            [][]byte `protobuf:"bytes,2,rep,name=deletions,proto3" json:"deletions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeLogEntry_Successor) Reset()         { *m = ChangeLogEntry_Successor{} }
func (m *ChangeLogEntry_Successor) String() string { return proto.CompactTextString(m) }
func (*ChangeLogEntry_Successor) ProtoMessage()    {}
func (*ChangeLogEntry_Successor) Descriptor() ([]byte, []int) {
	return fileDescriptor_merkle_65730cd798e98cd8, []int{0, 0}
}
func (m *ChangeLogEntry_Successor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeLogEntry_Successor.Unmarshal(m, b)
}
func (m *ChangeLogEntry_Successor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeLogEntry_Successor.Marshal(b, m, deterministic)
}
func (dst *ChangeLogEntry_Successor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeLogEntry_Successor.Merge(dst, src)
}
func (m *ChangeLogEntry_Successor) XXX_Size() int {
	return xxx_messageInfo_ChangeLogEntry_Successor.Size(m)
}
func (m *ChangeLogEntry_Successor) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeLogEntry_Successor.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeLogEntry_Successor proto.InternalMessageInfo

func (m *ChangeLogEntry_Successor) GetSuccessor() []byte {
	if m != nil {
		return m.Successor
	}
	return nil
}

func (m *ChangeLogEntry_Successor) GetDeletions() [][]byte {
	if m != nil {
		return m.Deletions
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeLogEntry)(nil), "ChangeLogEntry")
	proto.RegisterType((*ChangeLogEntry_Successor)(nil), "ChangeLogEntry.Successor")
}

func init() {
	proto.RegisterFile("github.com/daludaluking/sawtooth-sdk/protobuf/merkle_pb2/merkle.proto", fileDescriptor_merkle_65730cd798e98cd8)
}

var fileDescriptor_merkle_65730cd798e98cd8 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2d, 0x4e, 0x2c, 0x2f,
	0xc9, 0xcf, 0x2f, 0xc9, 0x88, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a,
	0x4d, 0xd3, 0xcf, 0x4d, 0x2d, 0xca, 0xce, 0x49, 0x8d, 0x2f, 0x48, 0x32, 0x82, 0x32, 0xf5, 0xc0,
	0x52, 0x4a, 0xa7, 0x18, 0xb9, 0xf8, 0x9c, 0x33, 0x12, 0xf3, 0xd2, 0x53, 0x7d, 0xf2, 0xd3, 0x5d,
	0xf3, 0x4a, 0x8a, 0x2a, 0x85, 0xc4, 0xb8, 0xd8, 0x0a, 0x12, 0x8b, 0x52, 0xf3, 0x4a, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x78, 0x82, 0xa0, 0x3c, 0x21, 0x19, 0x2e, 0xce, 0xc4, 0x94, 0x94, 0xcc, 0x92,
	0xcc, 0xfc, 0xbc, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x9e, 0x20, 0x84, 0x80, 0x90, 0x25, 0x17,
	0x57, 0x71, 0x69, 0x72, 0x72, 0x6a, 0x71, 0x71, 0x7e, 0x51, 0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06,
	0xb7, 0x91, 0xa4, 0x1e, 0xaa, 0xd1, 0x7a, 0xc1, 0x30, 0x15, 0x41, 0x48, 0x8a, 0xa5, 0xdc, 0xb9,
	0x38, 0xe1, 0x12, 0x20, 0x5b, 0xe0, 0x52, 0x50, 0x07, 0x20, 0x04, 0x40, 0xb2, 0x29, 0xa9, 0x39,
	0xa9, 0x28, 0x6e, 0x80, 0x0b, 0x38, 0xa9, 0x72, 0x89, 0xc2, 0x7c, 0xaf, 0x57, 0x9c, 0x92, 0xad,
	0x07, 0xf3, 0x7d, 0x00, 0x63, 0x14, 0x17, 0x22, 0x00, 0x92, 0xd8, 0xc0, 0xe2, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbd, 0xd3, 0x1e, 0x7d, 0x2b, 0x01, 0x00, 0x00,
}
