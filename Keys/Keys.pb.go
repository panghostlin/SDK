// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Keys.proto

package keys

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type CreateKeysRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	MemberID             string   `protobuf:"bytes,2,opt,name=memberID,proto3" json:"memberID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateKeysRequest) Reset()         { *m = CreateKeysRequest{} }
func (m *CreateKeysRequest) String() string { return proto.CompactTextString(m) }
func (*CreateKeysRequest) ProtoMessage()    {}
func (*CreateKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{0}
}

func (m *CreateKeysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateKeysRequest.Unmarshal(m, b)
}
func (m *CreateKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateKeysRequest.Marshal(b, m, deterministic)
}
func (m *CreateKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateKeysRequest.Merge(m, src)
}
func (m *CreateKeysRequest) XXX_Size() int {
	return xxx_messageInfo_CreateKeysRequest.Size(m)
}
func (m *CreateKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateKeysRequest proto.InternalMessageInfo

func (m *CreateKeysRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateKeysRequest) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

type CreateKeysResponse struct {
	HashKey              string   `protobuf:"bytes,1,opt,name=hashKey,proto3" json:"hashKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateKeysResponse) Reset()         { *m = CreateKeysResponse{} }
func (m *CreateKeysResponse) String() string { return proto.CompactTextString(m) }
func (*CreateKeysResponse) ProtoMessage()    {}
func (*CreateKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{1}
}

func (m *CreateKeysResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateKeysResponse.Unmarshal(m, b)
}
func (m *CreateKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateKeysResponse.Marshal(b, m, deterministic)
}
func (m *CreateKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateKeysResponse.Merge(m, src)
}
func (m *CreateKeysResponse) XXX_Size() int {
	return xxx_messageInfo_CreateKeysResponse.Size(m)
}
func (m *CreateKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateKeysResponse proto.InternalMessageInfo

func (m *CreateKeysResponse) GetHashKey() string {
	if m != nil {
		return m.HashKey
	}
	return ""
}

type CheckPasswordRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	MemberID             string   `protobuf:"bytes,2,opt,name=memberID,proto3" json:"memberID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPasswordRequest) Reset()         { *m = CheckPasswordRequest{} }
func (m *CheckPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*CheckPasswordRequest) ProtoMessage()    {}
func (*CheckPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{2}
}

func (m *CheckPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPasswordRequest.Unmarshal(m, b)
}
func (m *CheckPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPasswordRequest.Marshal(b, m, deterministic)
}
func (m *CheckPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPasswordRequest.Merge(m, src)
}
func (m *CheckPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_CheckPasswordRequest.Size(m)
}
func (m *CheckPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPasswordRequest proto.InternalMessageInfo

func (m *CheckPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CheckPasswordRequest) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

type CheckPasswordResponse struct {
	HashKey              string   `protobuf:"bytes,1,opt,name=hashKey,proto3" json:"hashKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckPasswordResponse) Reset()         { *m = CheckPasswordResponse{} }
func (m *CheckPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*CheckPasswordResponse) ProtoMessage()    {}
func (*CheckPasswordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{3}
}

func (m *CheckPasswordResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckPasswordResponse.Unmarshal(m, b)
}
func (m *CheckPasswordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckPasswordResponse.Marshal(b, m, deterministic)
}
func (m *CheckPasswordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckPasswordResponse.Merge(m, src)
}
func (m *CheckPasswordResponse) XXX_Size() int {
	return xxx_messageInfo_CheckPasswordResponse.Size(m)
}
func (m *CheckPasswordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckPasswordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckPasswordResponse proto.InternalMessageInfo

func (m *CheckPasswordResponse) GetHashKey() string {
	if m != nil {
		return m.HashKey
	}
	return ""
}

type GetPublicKeyRequest struct {
	MemberID             string   `protobuf:"bytes,1,opt,name=memberID,proto3" json:"memberID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPublicKeyRequest) Reset()         { *m = GetPublicKeyRequest{} }
func (m *GetPublicKeyRequest) String() string { return proto.CompactTextString(m) }
func (*GetPublicKeyRequest) ProtoMessage()    {}
func (*GetPublicKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{4}
}

func (m *GetPublicKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicKeyRequest.Unmarshal(m, b)
}
func (m *GetPublicKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicKeyRequest.Marshal(b, m, deterministic)
}
func (m *GetPublicKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicKeyRequest.Merge(m, src)
}
func (m *GetPublicKeyRequest) XXX_Size() int {
	return xxx_messageInfo_GetPublicKeyRequest.Size(m)
}
func (m *GetPublicKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicKeyRequest proto.InternalMessageInfo

func (m *GetPublicKeyRequest) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

type GetPublicKeyResponse struct {
	PublicKey            string   `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPublicKeyResponse) Reset()         { *m = GetPublicKeyResponse{} }
func (m *GetPublicKeyResponse) String() string { return proto.CompactTextString(m) }
func (*GetPublicKeyResponse) ProtoMessage()    {}
func (*GetPublicKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{5}
}

func (m *GetPublicKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicKeyResponse.Unmarshal(m, b)
}
func (m *GetPublicKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicKeyResponse.Marshal(b, m, deterministic)
}
func (m *GetPublicKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicKeyResponse.Merge(m, src)
}
func (m *GetPublicKeyResponse) XXX_Size() int {
	return xxx_messageInfo_GetPublicKeyResponse.Size(m)
}
func (m *GetPublicKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicKeyResponse proto.InternalMessageInfo

func (m *GetPublicKeyResponse) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type EncryptPictureRequest struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	MemberID             string   `protobuf:"bytes,2,opt,name=memberID,proto3" json:"memberID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptPictureRequest) Reset()         { *m = EncryptPictureRequest{} }
func (m *EncryptPictureRequest) String() string { return proto.CompactTextString(m) }
func (*EncryptPictureRequest) ProtoMessage()    {}
func (*EncryptPictureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{6}
}

func (m *EncryptPictureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptPictureRequest.Unmarshal(m, b)
}
func (m *EncryptPictureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptPictureRequest.Marshal(b, m, deterministic)
}
func (m *EncryptPictureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptPictureRequest.Merge(m, src)
}
func (m *EncryptPictureRequest) XXX_Size() int {
	return xxx_messageInfo_EncryptPictureRequest.Size(m)
}
func (m *EncryptPictureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptPictureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptPictureRequest proto.InternalMessageInfo

func (m *EncryptPictureRequest) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *EncryptPictureRequest) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

type EncryptPictureResponse struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptPictureResponse) Reset()         { *m = EncryptPictureResponse{} }
func (m *EncryptPictureResponse) String() string { return proto.CompactTextString(m) }
func (*EncryptPictureResponse) ProtoMessage()    {}
func (*EncryptPictureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{7}
}

func (m *EncryptPictureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptPictureResponse.Unmarshal(m, b)
}
func (m *EncryptPictureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptPictureResponse.Marshal(b, m, deterministic)
}
func (m *EncryptPictureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptPictureResponse.Merge(m, src)
}
func (m *EncryptPictureResponse) XXX_Size() int {
	return xxx_messageInfo_EncryptPictureResponse.Size(m)
}
func (m *EncryptPictureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptPictureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptPictureResponse proto.InternalMessageInfo

func (m *EncryptPictureResponse) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *EncryptPictureResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DecryptPictureRequest struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	MemberID             string   `protobuf:"bytes,2,opt,name=memberID,proto3" json:"memberID,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	HashKey              string   `protobuf:"bytes,4,opt,name=hashKey,proto3" json:"hashKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptPictureRequest) Reset()         { *m = DecryptPictureRequest{} }
func (m *DecryptPictureRequest) String() string { return proto.CompactTextString(m) }
func (*DecryptPictureRequest) ProtoMessage()    {}
func (*DecryptPictureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{8}
}

func (m *DecryptPictureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptPictureRequest.Unmarshal(m, b)
}
func (m *DecryptPictureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptPictureRequest.Marshal(b, m, deterministic)
}
func (m *DecryptPictureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptPictureRequest.Merge(m, src)
}
func (m *DecryptPictureRequest) XXX_Size() int {
	return xxx_messageInfo_DecryptPictureRequest.Size(m)
}
func (m *DecryptPictureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptPictureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptPictureRequest proto.InternalMessageInfo

func (m *DecryptPictureRequest) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *DecryptPictureRequest) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

func (m *DecryptPictureRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DecryptPictureRequest) GetHashKey() string {
	if m != nil {
		return m.HashKey
	}
	return ""
}

type DecryptPictureResponse struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Chunk                []byte   `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptPictureResponse) Reset()         { *m = DecryptPictureResponse{} }
func (m *DecryptPictureResponse) String() string { return proto.CompactTextString(m) }
func (*DecryptPictureResponse) ProtoMessage()    {}
func (*DecryptPictureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{9}
}

func (m *DecryptPictureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptPictureResponse.Unmarshal(m, b)
}
func (m *DecryptPictureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptPictureResponse.Marshal(b, m, deterministic)
}
func (m *DecryptPictureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptPictureResponse.Merge(m, src)
}
func (m *DecryptPictureResponse) XXX_Size() int {
	return xxx_messageInfo_DecryptPictureResponse.Size(m)
}
func (m *DecryptPictureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptPictureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptPictureResponse proto.InternalMessageInfo

func (m *DecryptPictureResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DecryptPictureResponse) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateKeysRequest)(nil), "CreateKeysRequest")
	proto.RegisterType((*CreateKeysResponse)(nil), "CreateKeysResponse")
	proto.RegisterType((*CheckPasswordRequest)(nil), "CheckPasswordRequest")
	proto.RegisterType((*CheckPasswordResponse)(nil), "CheckPasswordResponse")
	proto.RegisterType((*GetPublicKeyRequest)(nil), "GetPublicKeyRequest")
	proto.RegisterType((*GetPublicKeyResponse)(nil), "GetPublicKeyResponse")
	proto.RegisterType((*EncryptPictureRequest)(nil), "EncryptPictureRequest")
	proto.RegisterType((*EncryptPictureResponse)(nil), "EncryptPictureResponse")
	proto.RegisterType((*DecryptPictureRequest)(nil), "DecryptPictureRequest")
	proto.RegisterType((*DecryptPictureResponse)(nil), "DecryptPictureResponse")
}

func init() { proto.RegisterFile("Keys.proto", fileDescriptor_d1dcf5ddb16cc394) }

var fileDescriptor_d1dcf5ddb16cc394 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x5b, 0xf0, 0x1f, 0x23, 0x1a, 0x1d, 0xda, 0x42, 0x1a, 0x0f, 0x66, 0x4f, 0x9c, 0x36,
	0xa2, 0x26, 0x9e, 0x4c, 0x48, 0xc0, 0x10, 0xd2, 0xc4, 0x10, 0x7c, 0x02, 0xa8, 0x93, 0x94, 0x20,
	0xb4, 0x76, 0x5b, 0x4d, 0x1f, 0xd5, 0xb7, 0x31, 0x85, 0x6d, 0x69, 0xcb, 0x46, 0x0f, 0x7a, 0x63,
	0x66, 0x33, 0xbf, 0xef, 0x63, 0xe6, 0x4b, 0x01, 0x1c, 0x4a, 0x04, 0x0f, 0x42, 0x3f, 0xf2, 0x99,
	0x03, 0x97, 0x83, 0x90, 0x66, 0x11, 0xa5, 0xbd, 0x29, 0xbd, 0xc7, 0x24, 0x22, 0xb4, 0xe1, 0x24,
	0x98, 0x09, 0xf1, 0xe9, 0x87, 0xaf, 0x1d, 0xfd, 0x5a, 0xef, 0x36, 0xa6, 0x79, 0x9d, 0xbe, 0xad,
	0x68, 0x35, 0xa7, 0x70, 0x3c, 0xec, 0xd4, 0xb6, 0x6f, 0x59, 0xcd, 0x38, 0x60, 0x11, 0x26, 0x02,
	0x7f, 0x2d, 0x08, 0x3b, 0x70, 0xec, 0xcd, 0x84, 0xe7, 0x50, 0x22, 0x61, 0x59, 0xc9, 0x9e, 0xc1,
	0x18, 0x78, 0xe4, 0x2e, 0x27, 0x12, 0xfe, 0x57, 0xfd, 0x1e, 0x98, 0x15, 0xde, 0xaf, 0x16, 0x7a,
	0xd0, 0x1a, 0x51, 0x34, 0x89, 0xe7, 0x6f, 0x0b, 0xd7, 0xa1, 0xa4, 0xe0, 0x20, 0x57, 0xd1, 0x2b,
	0x2a, 0xf7, 0x60, 0x94, 0x47, 0xa4, 0xc8, 0x15, 0x34, 0x82, 0xac, 0x29, 0x87, 0x76, 0x0d, 0x36,
	0x06, 0xf3, 0x69, 0xed, 0x86, 0x49, 0x10, 0x4d, 0x16, 0x6e, 0x14, 0x87, 0x94, 0x49, 0x19, 0x70,
	0xe8, 0x7a, 0xf1, 0x7a, 0xb9, 0x19, 0x69, 0x4e, 0xb7, 0xc5, 0x8f, 0x7f, 0xb3, 0x0f, 0x56, 0x15,
	0x25, 0x2d, 0xa8, 0x59, 0x17, 0x50, 0x5f, 0x52, 0x22, 0x31, 0xe9, 0x4f, 0x16, 0x83, 0x39, 0xa4,
	0x7f, 0x31, 0x93, 0xc1, 0xeb, 0x39, 0xbc, 0xb8, 0xec, 0x83, 0xf2, 0xb2, 0xfb, 0x60, 0x55, 0x65,
	0xa5, 0x71, 0x49, 0xd1, 0x77, 0x94, 0xdc, 0x49, 0xad, 0xe0, 0xe4, 0xf6, 0xab, 0x06, 0xa7, 0x69,
	0xb8, 0x5e, 0x28, 0xfc, 0x58, 0xb8, 0x84, 0x0f, 0x00, 0xbb, 0xc4, 0x21, 0xf2, 0xbd, 0x2c, 0xdb,
	0x2d, 0xbe, 0x1f, 0x49, 0xa6, 0x61, 0x1f, 0xce, 0x4a, 0x51, 0x41, 0x93, 0xab, 0xa2, 0x68, 0x5b,
	0x5c, 0x99, 0x28, 0xa6, 0xe1, 0x23, 0x34, 0x8b, 0x31, 0x40, 0x83, 0x2b, 0x82, 0x64, 0x9b, 0x5c,
	0x95, 0x15, 0xa6, 0xe1, 0x08, 0xce, 0xcb, 0x47, 0x44, 0x8b, 0x2b, 0x03, 0x62, 0xb7, 0xb9, 0xfa,
	0xda, 0x4c, 0xeb, 0xea, 0x37, 0x7a, 0x0a, 0x2a, 0x2f, 0x15, 0x2d, 0xae, 0x3c, 0xae, 0xdd, 0xe6,
	0xea, 0xed, 0x6f, 0x41, 0xf3, 0xa3, 0xcd, 0x17, 0xe1, 0xee, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x78,
	0x0e, 0x30, 0xaa, 0x1f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeysServiceClient is the client API for KeysService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeysServiceClient interface {
	CreateKeys(ctx context.Context, in *CreateKeysRequest, opts ...grpc.CallOption) (*CreateKeysResponse, error)
	CheckPassword(ctx context.Context, in *CheckPasswordRequest, opts ...grpc.CallOption) (*CheckPasswordResponse, error)
	GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error)
	EncryptPicture(ctx context.Context, opts ...grpc.CallOption) (KeysService_EncryptPictureClient, error)
	DecryptPicture(ctx context.Context, opts ...grpc.CallOption) (KeysService_DecryptPictureClient, error)
}

type keysServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeysServiceClient(cc *grpc.ClientConn) KeysServiceClient {
	return &keysServiceClient{cc}
}

func (c *keysServiceClient) CreateKeys(ctx context.Context, in *CreateKeysRequest, opts ...grpc.CallOption) (*CreateKeysResponse, error) {
	out := new(CreateKeysResponse)
	err := c.cc.Invoke(ctx, "/KeysService/CreateKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysServiceClient) CheckPassword(ctx context.Context, in *CheckPasswordRequest, opts ...grpc.CallOption) (*CheckPasswordResponse, error) {
	out := new(CheckPasswordResponse)
	err := c.cc.Invoke(ctx, "/KeysService/CheckPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysServiceClient) GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error) {
	out := new(GetPublicKeyResponse)
	err := c.cc.Invoke(ctx, "/KeysService/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysServiceClient) EncryptPicture(ctx context.Context, opts ...grpc.CallOption) (KeysService_EncryptPictureClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KeysService_serviceDesc.Streams[0], "/KeysService/EncryptPicture", opts...)
	if err != nil {
		return nil, err
	}
	x := &keysServiceEncryptPictureClient{stream}
	return x, nil
}

type KeysService_EncryptPictureClient interface {
	Send(*EncryptPictureRequest) error
	Recv() (*EncryptPictureResponse, error)
	grpc.ClientStream
}

type keysServiceEncryptPictureClient struct {
	grpc.ClientStream
}

func (x *keysServiceEncryptPictureClient) Send(m *EncryptPictureRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *keysServiceEncryptPictureClient) Recv() (*EncryptPictureResponse, error) {
	m := new(EncryptPictureResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *keysServiceClient) DecryptPicture(ctx context.Context, opts ...grpc.CallOption) (KeysService_DecryptPictureClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KeysService_serviceDesc.Streams[1], "/KeysService/DecryptPicture", opts...)
	if err != nil {
		return nil, err
	}
	x := &keysServiceDecryptPictureClient{stream}
	return x, nil
}

type KeysService_DecryptPictureClient interface {
	Send(*DecryptPictureRequest) error
	Recv() (*DecryptPictureResponse, error)
	grpc.ClientStream
}

type keysServiceDecryptPictureClient struct {
	grpc.ClientStream
}

func (x *keysServiceDecryptPictureClient) Send(m *DecryptPictureRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *keysServiceDecryptPictureClient) Recv() (*DecryptPictureResponse, error) {
	m := new(DecryptPictureResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KeysServiceServer is the server API for KeysService service.
type KeysServiceServer interface {
	CreateKeys(context.Context, *CreateKeysRequest) (*CreateKeysResponse, error)
	CheckPassword(context.Context, *CheckPasswordRequest) (*CheckPasswordResponse, error)
	GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error)
	EncryptPicture(KeysService_EncryptPictureServer) error
	DecryptPicture(KeysService_DecryptPictureServer) error
}

// UnimplementedKeysServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeysServiceServer struct {
}

func (*UnimplementedKeysServiceServer) CreateKeys(ctx context.Context, req *CreateKeysRequest) (*CreateKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKeys not implemented")
}
func (*UnimplementedKeysServiceServer) CheckPassword(ctx context.Context, req *CheckPasswordRequest) (*CheckPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassword not implemented")
}
func (*UnimplementedKeysServiceServer) GetPublicKey(ctx context.Context, req *GetPublicKeyRequest) (*GetPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (*UnimplementedKeysServiceServer) EncryptPicture(srv KeysService_EncryptPictureServer) error {
	return status.Errorf(codes.Unimplemented, "method EncryptPicture not implemented")
}
func (*UnimplementedKeysServiceServer) DecryptPicture(srv KeysService_DecryptPictureServer) error {
	return status.Errorf(codes.Unimplemented, "method DecryptPicture not implemented")
}

func RegisterKeysServiceServer(s *grpc.Server, srv KeysServiceServer) {
	s.RegisterService(&_KeysService_serviceDesc, srv)
}

func _KeysService_CreateKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServiceServer).CreateKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeysService/CreateKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServiceServer).CreateKeys(ctx, req.(*CreateKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeysService_CheckPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServiceServer).CheckPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeysService/CheckPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServiceServer).CheckPassword(ctx, req.(*CheckPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeysService_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServiceServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeysService/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServiceServer).GetPublicKey(ctx, req.(*GetPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeysService_EncryptPicture_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KeysServiceServer).EncryptPicture(&keysServiceEncryptPictureServer{stream})
}

type KeysService_EncryptPictureServer interface {
	Send(*EncryptPictureResponse) error
	Recv() (*EncryptPictureRequest, error)
	grpc.ServerStream
}

type keysServiceEncryptPictureServer struct {
	grpc.ServerStream
}

func (x *keysServiceEncryptPictureServer) Send(m *EncryptPictureResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *keysServiceEncryptPictureServer) Recv() (*EncryptPictureRequest, error) {
	m := new(EncryptPictureRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KeysService_DecryptPicture_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KeysServiceServer).DecryptPicture(&keysServiceDecryptPictureServer{stream})
}

type KeysService_DecryptPictureServer interface {
	Send(*DecryptPictureResponse) error
	Recv() (*DecryptPictureRequest, error)
	grpc.ServerStream
}

type keysServiceDecryptPictureServer struct {
	grpc.ServerStream
}

func (x *keysServiceDecryptPictureServer) Send(m *DecryptPictureResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *keysServiceDecryptPictureServer) Recv() (*DecryptPictureRequest, error) {
	m := new(DecryptPictureRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _KeysService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "KeysService",
	HandlerType: (*KeysServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKeys",
			Handler:    _KeysService_CreateKeys_Handler,
		},
		{
			MethodName: "CheckPassword",
			Handler:    _KeysService_CheckPassword_Handler,
		},
		{
			MethodName: "GetPublicKey",
			Handler:    _KeysService_GetPublicKey_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EncryptPicture",
			Handler:       _KeysService_EncryptPicture_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DecryptPicture",
			Handler:       _KeysService_DecryptPicture_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Keys.proto",
}
