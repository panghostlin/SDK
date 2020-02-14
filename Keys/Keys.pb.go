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

//*************************************************************************
//*	RPCS - Members
//************************************************************************
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

type GetPrivateKeyRequest struct {
	MemberID             string   `protobuf:"bytes,1,opt,name=memberID,proto3" json:"memberID,omitempty"`
	HashKey              string   `protobuf:"bytes,2,opt,name=hashKey,proto3" json:"hashKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPrivateKeyRequest) Reset()         { *m = GetPrivateKeyRequest{} }
func (m *GetPrivateKeyRequest) String() string { return proto.CompactTextString(m) }
func (*GetPrivateKeyRequest) ProtoMessage()    {}
func (*GetPrivateKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{6}
}

func (m *GetPrivateKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPrivateKeyRequest.Unmarshal(m, b)
}
func (m *GetPrivateKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPrivateKeyRequest.Marshal(b, m, deterministic)
}
func (m *GetPrivateKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPrivateKeyRequest.Merge(m, src)
}
func (m *GetPrivateKeyRequest) XXX_Size() int {
	return xxx_messageInfo_GetPrivateKeyRequest.Size(m)
}
func (m *GetPrivateKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPrivateKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPrivateKeyRequest proto.InternalMessageInfo

func (m *GetPrivateKeyRequest) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

func (m *GetPrivateKeyRequest) GetHashKey() string {
	if m != nil {
		return m.HashKey
	}
	return ""
}

type GetPrivateKeyResponse struct {
	PrivateKey           string   `protobuf:"bytes,1,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPrivateKeyResponse) Reset()         { *m = GetPrivateKeyResponse{} }
func (m *GetPrivateKeyResponse) String() string { return proto.CompactTextString(m) }
func (*GetPrivateKeyResponse) ProtoMessage()    {}
func (*GetPrivateKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{7}
}

func (m *GetPrivateKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPrivateKeyResponse.Unmarshal(m, b)
}
func (m *GetPrivateKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPrivateKeyResponse.Marshal(b, m, deterministic)
}
func (m *GetPrivateKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPrivateKeyResponse.Merge(m, src)
}
func (m *GetPrivateKeyResponse) XXX_Size() int {
	return xxx_messageInfo_GetPrivateKeyResponse.Size(m)
}
func (m *GetPrivateKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPrivateKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPrivateKeyResponse proto.InternalMessageInfo

func (m *GetPrivateKeyResponse) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

type GetKeyRequest struct {
	MemberID             string   `protobuf:"bytes,1,opt,name=memberID,proto3" json:"memberID,omitempty"`
	HashKey              string   `protobuf:"bytes,2,opt,name=hashKey,proto3" json:"hashKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeyRequest) Reset()         { *m = GetKeyRequest{} }
func (m *GetKeyRequest) String() string { return proto.CompactTextString(m) }
func (*GetKeyRequest) ProtoMessage()    {}
func (*GetKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{8}
}

func (m *GetKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeyRequest.Unmarshal(m, b)
}
func (m *GetKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeyRequest.Marshal(b, m, deterministic)
}
func (m *GetKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeyRequest.Merge(m, src)
}
func (m *GetKeyRequest) XXX_Size() int {
	return xxx_messageInfo_GetKeyRequest.Size(m)
}
func (m *GetKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeyRequest proto.InternalMessageInfo

func (m *GetKeyRequest) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

func (m *GetKeyRequest) GetHashKey() string {
	if m != nil {
		return m.HashKey
	}
	return ""
}

type GetKeyResponse struct {
	PublicKey            string   `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	PrivateKey           string   `protobuf:"bytes,2,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKeyResponse) Reset()         { *m = GetKeyResponse{} }
func (m *GetKeyResponse) String() string { return proto.CompactTextString(m) }
func (*GetKeyResponse) ProtoMessage()    {}
func (*GetKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dcf5ddb16cc394, []int{9}
}

func (m *GetKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKeyResponse.Unmarshal(m, b)
}
func (m *GetKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKeyResponse.Marshal(b, m, deterministic)
}
func (m *GetKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKeyResponse.Merge(m, src)
}
func (m *GetKeyResponse) XXX_Size() int {
	return xxx_messageInfo_GetKeyResponse.Size(m)
}
func (m *GetKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetKeyResponse proto.InternalMessageInfo

func (m *GetKeyResponse) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *GetKeyResponse) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

//*************************************************************************
//*	RPCS - Pictures
//************************************************************************
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
	return fileDescriptor_d1dcf5ddb16cc394, []int{10}
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
	return fileDescriptor_d1dcf5ddb16cc394, []int{11}
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
	return fileDescriptor_d1dcf5ddb16cc394, []int{12}
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
	return fileDescriptor_d1dcf5ddb16cc394, []int{13}
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
	proto.RegisterType((*GetPrivateKeyRequest)(nil), "GetPrivateKeyRequest")
	proto.RegisterType((*GetPrivateKeyResponse)(nil), "GetPrivateKeyResponse")
	proto.RegisterType((*GetKeyRequest)(nil), "GetKeyRequest")
	proto.RegisterType((*GetKeyResponse)(nil), "GetKeyResponse")
	proto.RegisterType((*EncryptPictureRequest)(nil), "EncryptPictureRequest")
	proto.RegisterType((*EncryptPictureResponse)(nil), "EncryptPictureResponse")
	proto.RegisterType((*DecryptPictureRequest)(nil), "DecryptPictureRequest")
	proto.RegisterType((*DecryptPictureResponse)(nil), "DecryptPictureResponse")
}

func init() { proto.RegisterFile("Keys.proto", fileDescriptor_d1dcf5ddb16cc394) }

var fileDescriptor_d1dcf5ddb16cc394 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x61, 0x6b, 0xd4, 0x40,
	0x10, 0x86, 0x93, 0x9c, 0x56, 0x3b, 0xb6, 0xa7, 0x4e, 0x93, 0xf4, 0x58, 0x44, 0x64, 0x3f, 0x15,
	0x84, 0xc5, 0xaa, 0xd0, 0x4f, 0xc2, 0x41, 0x5b, 0x8e, 0x12, 0x29, 0xc7, 0xf9, 0x0b, 0xee, 0xe2,
	0x40, 0xc2, 0xd9, 0x4b, 0xdc, 0x24, 0x95, 0xfc, 0x43, 0x7f, 0x96, 0xa4, 0xbb, 0xc9, 0x65, 0x73,
	0x8b, 0x16, 0xce, 0x6f, 0x99, 0xd9, 0xec, 0x33, 0x6f, 0x76, 0xdf, 0x37, 0x00, 0x11, 0xd5, 0x85,
	0xc8, 0x65, 0x56, 0x66, 0x3c, 0x82, 0xd7, 0x97, 0x92, 0x96, 0x25, 0x35, 0xbd, 0x05, 0xfd, 0xac,
	0xa8, 0x28, 0x91, 0xc1, 0xf3, 0x7c, 0x59, 0x14, 0xbf, 0x32, 0xf9, 0x7d, 0xe2, 0xbe, 0x73, 0xcf,
	0x0e, 0x17, 0x5d, 0xdd, 0xac, 0xdd, 0xd1, 0xdd, 0x8a, 0xe4, 0xcd, 0xd5, 0xc4, 0x53, 0x6b, 0x6d,
	0xcd, 0x05, 0x60, 0x1f, 0x56, 0xe4, 0xd9, 0xa6, 0x20, 0x9c, 0xc0, 0xb3, 0x64, 0x59, 0x24, 0x11,
	0xd5, 0x1a, 0xd6, 0x96, 0xfc, 0x16, 0xfc, 0xcb, 0x84, 0xe2, 0xf5, 0x5c, 0xc3, 0xf7, 0x9d, 0x7f,
	0x0e, 0xc1, 0x80, 0xf7, 0x4f, 0x09, 0xe7, 0x70, 0x32, 0xa3, 0x72, 0x5e, 0xad, 0x7e, 0xa4, 0x71,
	0x44, 0x75, 0x4f, 0x41, 0x37, 0xc5, 0x1d, 0x4c, 0xf9, 0x0c, 0xbe, 0xb9, 0x45, 0x0f, 0x79, 0x03,
	0x87, 0x79, 0xdb, 0xd4, 0x9b, 0xb6, 0x0d, 0xfe, 0x55, 0xed, 0x92, 0xe9, 0xbd, 0x3a, 0x9f, 0x47,
	0x4c, 0xea, 0xcb, 0xf6, 0x4c, 0xd9, 0x17, 0x10, 0x0c, 0x68, 0x5a, 0xc4, 0x5b, 0x80, 0xbc, 0xeb,
	0x6a, 0x60, 0xaf, 0xc3, 0xaf, 0xe1, 0x78, 0x46, 0xe5, 0xde, 0xf3, 0x6f, 0x61, 0xdc, 0x62, 0x1e,
	0xf3, 0xf5, 0x03, 0x59, 0xde, 0x8e, 0xac, 0x1b, 0x08, 0xae, 0x37, 0xb1, 0xac, 0xf3, 0x72, 0x9e,
	0xc6, 0x65, 0x25, 0xa9, 0x95, 0xe7, 0xc3, 0xd3, 0x38, 0xa9, 0x36, 0xeb, 0x07, 0xe4, 0xd1, 0x42,
	0x15, 0x7f, 0x35, 0xc1, 0x14, 0xc2, 0x21, 0x4a, 0x4b, 0xb4, 0xb3, 0x5e, 0xc1, 0x68, 0xdd, 0x69,
	0x6a, 0x1e, 0x79, 0x05, 0xc1, 0x15, 0xfd, 0x17, 0x31, 0x2d, 0x7c, 0xd4, 0xc1, 0xfb, 0x67, 0xfa,
	0xc4, 0x3c, 0xd3, 0x29, 0x84, 0xc3, 0xb1, 0x5a, 0xb8, 0xa6, 0xb8, 0x5b, 0x4a, 0xa7, 0xc4, 0xeb,
	0x29, 0xf9, 0xf8, 0x7b, 0x04, 0x2f, 0x9a, 0xe8, 0x7d, 0x23, 0x79, 0x9f, 0xc6, 0x84, 0x17, 0x00,
	0xdb, 0x3c, 0x22, 0x8a, 0x9d, 0xa4, 0xb3, 0x13, 0xb1, 0x1b, 0x58, 0xee, 0xe0, 0x14, 0x8e, 0x8d,
	0x20, 0x61, 0x20, 0x6c, 0x41, 0x65, 0xa1, 0xb0, 0xe6, 0x8d, 0x3b, 0xf8, 0x05, 0x8e, 0xfa, 0x21,
	0x41, 0x5f, 0x58, 0x62, 0xc6, 0x02, 0x61, 0x4b, 0x92, 0x12, 0x60, 0xf8, 0x1b, 0xd5, 0x9b, 0xc3,
	0xf4, 0xb0, 0x50, 0x58, 0x63, 0xc0, 0x1d, 0x7c, 0x0f, 0x07, 0xca, 0xa1, 0x38, 0x16, 0x86, 0xe3,
	0xd9, 0x4b, 0x61, 0x5a, 0x97, 0x3b, 0x38, 0x83, 0xb1, 0xe9, 0x19, 0x0c, 0x85, 0xd5, 0x8f, 0xec,
	0x54, 0xd8, 0xcd, 0xc5, 0x9d, 0x33, 0xf7, 0x83, 0xdb, 0x80, 0xcc, 0x3b, 0xc4, 0x50, 0x58, 0xbd,
	0xc4, 0x4e, 0x85, 0xfd, 0xb2, 0x15, 0x68, 0x75, 0xf0, 0xf0, 0x7b, 0xfe, 0xf4, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xc1, 0xa1, 0x48, 0x2a, 0xac, 0x05, 0x00, 0x00,
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
	GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequest, opts ...grpc.CallOption) (*GetPrivateKeyResponse, error)
	GetKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*GetKeyResponse, error)
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

func (c *keysServiceClient) GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequest, opts ...grpc.CallOption) (*GetPrivateKeyResponse, error) {
	out := new(GetPrivateKeyResponse)
	err := c.cc.Invoke(ctx, "/KeysService/GetPrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keysServiceClient) GetKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*GetKeyResponse, error) {
	out := new(GetKeyResponse)
	err := c.cc.Invoke(ctx, "/KeysService/GetKey", in, out, opts...)
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
	GetPrivateKey(context.Context, *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error)
	GetKey(context.Context, *GetKeyRequest) (*GetKeyResponse, error)
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
func (*UnimplementedKeysServiceServer) GetPrivateKey(ctx context.Context, req *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateKey not implemented")
}
func (*UnimplementedKeysServiceServer) GetKey(ctx context.Context, req *GetKeyRequest) (*GetKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
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

func _KeysService_GetPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServiceServer).GetPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeysService/GetPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServiceServer).GetPrivateKey(ctx, req.(*GetPrivateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeysService_GetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeysServiceServer).GetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeysService/GetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeysServiceServer).GetKey(ctx, req.(*GetKeyRequest))
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
		{
			MethodName: "GetPrivateKey",
			Handler:    _KeysService_GetPrivateKey_Handler,
		},
		{
			MethodName: "GetKey",
			Handler:    _KeysService_GetKey_Handler,
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
