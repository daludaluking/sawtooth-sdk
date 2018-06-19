// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sawtooth_sdk/protobuf/identity_pb2/identity.proto

package identity

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

type Policy_EntryType int32

const (
	Policy_ENTRY_TYPE_UNSET Policy_EntryType = 0
	Policy_PERMIT_KEY       Policy_EntryType = 1
	Policy_DENY_KEY         Policy_EntryType = 2
)

var Policy_EntryType_name = map[int32]string{
	0: "ENTRY_TYPE_UNSET",
	1: "PERMIT_KEY",
	2: "DENY_KEY",
}
var Policy_EntryType_value = map[string]int32{
	"ENTRY_TYPE_UNSET": 0,
	"PERMIT_KEY":       1,
	"DENY_KEY":         2,
}

func (x Policy_EntryType) String() string {
	return proto.EnumName(Policy_EntryType_name, int32(x))
}
func (Policy_EntryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_identity_a83001c1ef7d245c, []int{0, 0}
}

type Policy struct {
	// name of the policy, this should be unique.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// list of Entries
	// The entries will be processed in order from first to last.
	Entries              []*Policy_Entry `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_a83001c1ef7d245c, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (dst *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(dst, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetEntries() []*Policy_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Policy_Entry struct {
	// Whether this is a Permit_KEY or Deny_KEY entry
	Type Policy_EntryType `protobuf:"varint,1,opt,name=type,enum=Policy_EntryType" json:"type,omitempty"`
	// This should be a list of public keys or * to refer to all participants.
	// If using *, it should be the only key in the list.
	Key                  string   `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy_Entry) Reset()         { *m = Policy_Entry{} }
func (m *Policy_Entry) String() string { return proto.CompactTextString(m) }
func (*Policy_Entry) ProtoMessage()    {}
func (*Policy_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_a83001c1ef7d245c, []int{0, 0}
}
func (m *Policy_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy_Entry.Unmarshal(m, b)
}
func (m *Policy_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy_Entry.Marshal(b, m, deterministic)
}
func (dst *Policy_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy_Entry.Merge(dst, src)
}
func (m *Policy_Entry) XXX_Size() int {
	return xxx_messageInfo_Policy_Entry.Size(m)
}
func (m *Policy_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Policy_Entry proto.InternalMessageInfo

func (m *Policy_Entry) GetType() Policy_EntryType {
	if m != nil {
		return m.Type
	}
	return Policy_ENTRY_TYPE_UNSET
}

func (m *Policy_Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// Policy will be stored in a Policy list to account for state collisions
type PolicyList struct {
	Policies             []*Policy `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PolicyList) Reset()         { *m = PolicyList{} }
func (m *PolicyList) String() string { return proto.CompactTextString(m) }
func (*PolicyList) ProtoMessage()    {}
func (*PolicyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_a83001c1ef7d245c, []int{1}
}
func (m *PolicyList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyList.Unmarshal(m, b)
}
func (m *PolicyList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyList.Marshal(b, m, deterministic)
}
func (dst *PolicyList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyList.Merge(dst, src)
}
func (m *PolicyList) XXX_Size() int {
	return xxx_messageInfo_PolicyList.Size(m)
}
func (m *PolicyList) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyList.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyList proto.InternalMessageInfo

func (m *PolicyList) GetPolicies() []*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type Role struct {
	// Role name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Name of corresponding policy
	PolicyName           string   `protobuf:"bytes,2,opt,name=policy_name,json=policyName" json:"policy_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_a83001c1ef7d245c, []int{2}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (dst *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(dst, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

// Roles will be stored in a RoleList to account for state collisions
type RoleList struct {
	Roles                []*Role  `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleList) Reset()         { *m = RoleList{} }
func (m *RoleList) String() string { return proto.CompactTextString(m) }
func (*RoleList) ProtoMessage()    {}
func (*RoleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_a83001c1ef7d245c, []int{3}
}
func (m *RoleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleList.Unmarshal(m, b)
}
func (m *RoleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleList.Marshal(b, m, deterministic)
}
func (dst *RoleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleList.Merge(dst, src)
}
func (m *RoleList) XXX_Size() int {
	return xxx_messageInfo_RoleList.Size(m)
}
func (m *RoleList) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleList.DiscardUnknown(m)
}

var xxx_messageInfo_RoleList proto.InternalMessageInfo

func (m *RoleList) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "Policy")
	proto.RegisterType((*Policy_Entry)(nil), "Policy.Entry")
	proto.RegisterType((*PolicyList)(nil), "PolicyList")
	proto.RegisterType((*Role)(nil), "Role")
	proto.RegisterType((*RoleList)(nil), "RoleList")
	proto.RegisterEnum("Policy_EntryType", Policy_EntryType_name, Policy_EntryType_value)
}

func init() {
	proto.RegisterFile("github.com/daludaluking/sawtooth-sdk/protobuf/identity_pb2/identity.proto", fileDescriptor_identity_a83001c1ef7d245c)
}

var fileDescriptor_identity_a83001c1ef7d245c = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xbf, 0xd9, 0xcf, 0xee, 0xdd, 0xd7, 0x51, 0x83, 0x87, 0x32, 0x41, 0x47, 0x45, 0xb6,
	0x53, 0xc7, 0xea, 0xd1, 0x83, 0x22, 0xe6, 0x20, 0x6a, 0x29, 0xb1, 0x1e, 0x7a, 0x0a, 0xad, 0xcb,
	0xb0, 0xac, 0x6b, 0x4a, 0x9b, 0x21, 0xf9, 0x0f, 0xfd, 0xb3, 0xa4, 0x89, 0x2b, 0x0c, 0xbc, 0xbd,
	0xef, 0xf3, 0x79, 0xda, 0xe7, 0x49, 0x02, 0xab, 0x3a, 0xf9, 0x92, 0x42, 0xc8, 0x4f, 0x56, 0xaf,
	0xb7, 0xcb, 0xb2, 0x12, 0x52, 0xa4, 0xfb, 0xcd, 0x32, 0x5b, 0xf3, 0x42, 0x66, 0x52, 0xb1, 0x32,
	0xf5, 0xdb, 0xc5, 0xd3, 0xd8, 0xfd, 0x46, 0x30, 0x08, 0x45, 0x9e, 0x7d, 0x28, 0x8c, 0xa1, 0x57,
	0x24, 0x3b, 0xee, 0xa0, 0x19, 0x5a, 0x8c, 0xa8, 0x9e, 0xf1, 0x1c, 0x86, 0xbc, 0x90, 0x55, 0xc6,
	0x6b, 0xa7, 0x33, 0xeb, 0x2e, 0xc6, 0xfe, 0x89, 0x67, 0xdc, 0x1e, 0x29, 0x64, 0xa5, 0xe8, 0x81,
	0x4e, 0xef, 0xa1, 0xaf, 0x15, 0x7c, 0x0d, 0x3d, 0xa9, 0x4a, 0xf3, 0x97, 0x89, 0x7f, 0x7a, 0x64,
	0x8f, 0x54, 0xc9, 0xa9, 0xc6, 0xd8, 0x86, 0xee, 0x96, 0x2b, 0xa7, 0xa3, 0xb3, 0x9a, 0xd1, 0xbd,
	0x83, 0x51, 0x6b, 0xc2, 0x67, 0x60, 0x93, 0x20, 0xa2, 0x31, 0x8b, 0xe2, 0x90, 0xb0, 0xf7, 0xe0,
	0x8d, 0x44, 0xf6, 0x3f, 0x3c, 0x01, 0x08, 0x09, 0x7d, 0x7d, 0x8a, 0xd8, 0x33, 0x89, 0x6d, 0x84,
	0xff, 0x83, 0xf5, 0x48, 0x82, 0x58, 0x6f, 0x1d, 0x77, 0x05, 0x60, 0xc2, 0x5e, 0xb2, 0x5a, 0xe2,
	0x2b, 0xb0, 0xca, 0x66, 0x6b, 0xaa, 0x23, 0x5d, 0x7d, 0xf8, 0xdb, 0x85, 0xb6, 0xc0, 0xbd, 0x85,
	0x1e, 0x15, 0x39, 0xff, 0xf3, 0xe8, 0x97, 0x30, 0xd6, 0x3e, 0xc5, 0x34, 0x32, 0x4d, 0xc1, 0x48,
	0x41, 0xb2, 0xe3, 0xee, 0x1c, 0xac, 0xe6, 0x63, 0x9d, 0x76, 0x0e, 0xfd, 0x4a, 0xe4, 0x6d, 0x54,
	0xdf, 0x6b, 0x08, 0x35, 0xda, 0xc3, 0x05, 0x4c, 0x0f, 0x0f, 0xe3, 0x1d, 0x5f, 0x7f, 0xba, 0xdf,
	0x84, 0x28, 0x1d, 0xe8, 0xf9, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x6b, 0xa9, 0x0d, 0xbf,
	0x01, 0x00, 0x00,
}
