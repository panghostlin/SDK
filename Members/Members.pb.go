// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Members.proto

package members

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
//*	RPCS
//************************************************************************
type CreateMemberRequest struct {
	Email                string          `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string          `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PublicKey            string          `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	PrivateKey           *CryptedPrivate `protobuf:"bytes,4,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateMemberRequest) Reset()         { *m = CreateMemberRequest{} }
func (m *CreateMemberRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMemberRequest) ProtoMessage()    {}
func (*CreateMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d0e7a977b094cf8, []int{0}
}

func (m *CreateMemberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMemberRequest.Unmarshal(m, b)
}
func (m *CreateMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMemberRequest.Marshal(b, m, deterministic)
}
func (m *CreateMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMemberRequest.Merge(m, src)
}
func (m *CreateMemberRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMemberRequest.Size(m)
}
func (m *CreateMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMemberRequest proto.InternalMessageInfo

func (m *CreateMemberRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateMemberRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateMemberRequest) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *CreateMemberRequest) GetPrivateKey() *CryptedPrivate {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

type CreateMemberResponse struct {
	MemberID             string   `protobuf:"bytes,1,opt,name=memberID,proto3" json:"memberID,omitempty"`
	AccessToken          *Cookie  `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	Keys                 *Keys    `protobuf:"bytes,3,opt,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMemberResponse) Reset()         { *m = CreateMemberResponse{} }
func (m *CreateMemberResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMemberResponse) ProtoMessage()    {}
func (*CreateMemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d0e7a977b094cf8, []int{1}
}

func (m *CreateMemberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMemberResponse.Unmarshal(m, b)
}
func (m *CreateMemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMemberResponse.Marshal(b, m, deterministic)
}
func (m *CreateMemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMemberResponse.Merge(m, src)
}
func (m *CreateMemberResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMemberResponse.Size(m)
}
func (m *CreateMemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMemberResponse proto.InternalMessageInfo

func (m *CreateMemberResponse) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

func (m *CreateMemberResponse) GetAccessToken() *Cookie {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

func (m *CreateMemberResponse) GetKeys() *Keys {
	if m != nil {
		return m.Keys
	}
	return nil
}

type LoginMemberRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginMemberRequest) Reset()         { *m = LoginMemberRequest{} }
func (m *LoginMemberRequest) String() string { return proto.CompactTextString(m) }
func (*LoginMemberRequest) ProtoMessage()    {}
func (*LoginMemberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d0e7a977b094cf8, []int{2}
}

func (m *LoginMemberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginMemberRequest.Unmarshal(m, b)
}
func (m *LoginMemberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginMemberRequest.Marshal(b, m, deterministic)
}
func (m *LoginMemberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginMemberRequest.Merge(m, src)
}
func (m *LoginMemberRequest) XXX_Size() int {
	return xxx_messageInfo_LoginMemberRequest.Size(m)
}
func (m *LoginMemberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginMemberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginMemberRequest proto.InternalMessageInfo

func (m *LoginMemberRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginMemberRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginMemberResponse struct {
	MemberID             string   `protobuf:"bytes,1,opt,name=memberID,proto3" json:"memberID,omitempty"`
	AccessToken          *Cookie  `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	Keys                 *Keys    `protobuf:"bytes,3,opt,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginMemberResponse) Reset()         { *m = LoginMemberResponse{} }
func (m *LoginMemberResponse) String() string { return proto.CompactTextString(m) }
func (*LoginMemberResponse) ProtoMessage()    {}
func (*LoginMemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d0e7a977b094cf8, []int{3}
}

func (m *LoginMemberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginMemberResponse.Unmarshal(m, b)
}
func (m *LoginMemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginMemberResponse.Marshal(b, m, deterministic)
}
func (m *LoginMemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginMemberResponse.Merge(m, src)
}
func (m *LoginMemberResponse) XXX_Size() int {
	return xxx_messageInfo_LoginMemberResponse.Size(m)
}
func (m *LoginMemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginMemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginMemberResponse proto.InternalMessageInfo

func (m *LoginMemberResponse) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

func (m *LoginMemberResponse) GetAccessToken() *Cookie {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

func (m *LoginMemberResponse) GetKeys() *Keys {
	if m != nil {
		return m.Keys
	}
	return nil
}

type CheckAccessTokenRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAccessTokenRequest) Reset()         { *m = CheckAccessTokenRequest{} }
func (m *CheckAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*CheckAccessTokenRequest) ProtoMessage()    {}
func (*CheckAccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d0e7a977b094cf8, []int{4}
}

func (m *CheckAccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAccessTokenRequest.Unmarshal(m, b)
}
func (m *CheckAccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAccessTokenRequest.Marshal(b, m, deterministic)
}
func (m *CheckAccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAccessTokenRequest.Merge(m, src)
}
func (m *CheckAccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_CheckAccessTokenRequest.Size(m)
}
func (m *CheckAccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAccessTokenRequest proto.InternalMessageInfo

func (m *CheckAccessTokenRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type CheckAccessTokenResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	MemberID             string   `protobuf:"bytes,2,opt,name=memberID,proto3" json:"memberID,omitempty"`
	AccessToken          *Cookie  `protobuf:"bytes,3,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAccessTokenResponse) Reset()         { *m = CheckAccessTokenResponse{} }
func (m *CheckAccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*CheckAccessTokenResponse) ProtoMessage()    {}
func (*CheckAccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d0e7a977b094cf8, []int{5}
}

func (m *CheckAccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAccessTokenResponse.Unmarshal(m, b)
}
func (m *CheckAccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAccessTokenResponse.Marshal(b, m, deterministic)
}
func (m *CheckAccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAccessTokenResponse.Merge(m, src)
}
func (m *CheckAccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_CheckAccessTokenResponse.Size(m)
}
func (m *CheckAccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAccessTokenResponse proto.InternalMessageInfo

func (m *CheckAccessTokenResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CheckAccessTokenResponse) GetMemberID() string {
	if m != nil {
		return m.MemberID
	}
	return ""
}

func (m *CheckAccessTokenResponse) GetAccessToken() *Cookie {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

//*************************************************************************
//*	HELPERS
//************************************************************************
type Cookie struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Expiration           int64    `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cookie) Reset()         { *m = Cookie{} }
func (m *Cookie) String() string { return proto.CompactTextString(m) }
func (*Cookie) ProtoMessage()    {}
func (*Cookie) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d0e7a977b094cf8, []int{6}
}

func (m *Cookie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cookie.Unmarshal(m, b)
}
func (m *Cookie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cookie.Marshal(b, m, deterministic)
}
func (m *Cookie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cookie.Merge(m, src)
}
func (m *Cookie) XXX_Size() int {
	return xxx_messageInfo_Cookie.Size(m)
}
func (m *Cookie) XXX_DiscardUnknown() {
	xxx_messageInfo_Cookie.DiscardUnknown(m)
}

var xxx_messageInfo_Cookie proto.InternalMessageInfo

func (m *Cookie) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Cookie) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type Member struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PublicKey            string   `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	PrivateKey           string   `protobuf:"bytes,4,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d0e7a977b094cf8, []int{7}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Member) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Member) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *Member) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

type CryptedPrivate struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Salt                 string   `protobuf:"bytes,2,opt,name=Salt,proto3" json:"Salt,omitempty"`
	IV                   string   `protobuf:"bytes,3,opt,name=IV,proto3" json:"IV,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptedPrivate) Reset()         { *m = CryptedPrivate{} }
func (m *CryptedPrivate) String() string { return proto.CompactTextString(m) }
func (*CryptedPrivate) ProtoMessage()    {}
func (*CryptedPrivate) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d0e7a977b094cf8, []int{8}
}

func (m *CryptedPrivate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptedPrivate.Unmarshal(m, b)
}
func (m *CryptedPrivate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptedPrivate.Marshal(b, m, deterministic)
}
func (m *CryptedPrivate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptedPrivate.Merge(m, src)
}
func (m *CryptedPrivate) XXX_Size() int {
	return xxx_messageInfo_CryptedPrivate.Size(m)
}
func (m *CryptedPrivate) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptedPrivate.DiscardUnknown(m)
}

var xxx_messageInfo_CryptedPrivate proto.InternalMessageInfo

func (m *CryptedPrivate) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CryptedPrivate) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *CryptedPrivate) GetIV() string {
	if m != nil {
		return m.IV
	}
	return ""
}

type Keys struct {
	PrivateKey           string   `protobuf:"bytes,1,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	PrivateSalt          string   `protobuf:"bytes,2,opt,name=privateSalt,proto3" json:"privateSalt,omitempty"`
	PrivateIV            string   `protobuf:"bytes,3,opt,name=privateIV,proto3" json:"privateIV,omitempty"`
	PublicKey            string   `protobuf:"bytes,4,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Keys) Reset()         { *m = Keys{} }
func (m *Keys) String() string { return proto.CompactTextString(m) }
func (*Keys) ProtoMessage()    {}
func (*Keys) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d0e7a977b094cf8, []int{9}
}

func (m *Keys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keys.Unmarshal(m, b)
}
func (m *Keys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keys.Marshal(b, m, deterministic)
}
func (m *Keys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keys.Merge(m, src)
}
func (m *Keys) XXX_Size() int {
	return xxx_messageInfo_Keys.Size(m)
}
func (m *Keys) XXX_DiscardUnknown() {
	xxx_messageInfo_Keys.DiscardUnknown(m)
}

var xxx_messageInfo_Keys proto.InternalMessageInfo

func (m *Keys) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *Keys) GetPrivateSalt() string {
	if m != nil {
		return m.PrivateSalt
	}
	return ""
}

func (m *Keys) GetPrivateIV() string {
	if m != nil {
		return m.PrivateIV
	}
	return ""
}

func (m *Keys) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateMemberRequest)(nil), "CreateMemberRequest")
	proto.RegisterType((*CreateMemberResponse)(nil), "CreateMemberResponse")
	proto.RegisterType((*LoginMemberRequest)(nil), "LoginMemberRequest")
	proto.RegisterType((*LoginMemberResponse)(nil), "LoginMemberResponse")
	proto.RegisterType((*CheckAccessTokenRequest)(nil), "CheckAccessTokenRequest")
	proto.RegisterType((*CheckAccessTokenResponse)(nil), "CheckAccessTokenResponse")
	proto.RegisterType((*Cookie)(nil), "Cookie")
	proto.RegisterType((*Member)(nil), "Member")
	proto.RegisterType((*CryptedPrivate)(nil), "CryptedPrivate")
	proto.RegisterType((*Keys)(nil), "Keys")
}

func init() { proto.RegisterFile("Members.proto", fileDescriptor_8d0e7a977b094cf8) }

var fileDescriptor_8d0e7a977b094cf8 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x1d, 0x37, 0x6d, 0xc6, 0x10, 0xaa, 0x49, 0x10, 0x6e, 0x84, 0xaa, 0x68, 0x4f, 0xe5,
	0xb2, 0x48, 0xe1, 0x06, 0x02, 0x09, 0xa5, 0xaa, 0x14, 0x15, 0x24, 0xe4, 0xa2, 0xde, 0x1d, 0x77,
	0x04, 0x56, 0x9c, 0xac, 0xd9, 0x75, 0xd2, 0x46, 0x5c, 0xf9, 0x06, 0x3e, 0x8d, 0xef, 0x41, 0xde,
	0x75, 0x92, 0xb5, 0x93, 0xf4, 0x82, 0xc4, 0xcd, 0xf3, 0x66, 0xbd, 0x6f, 0xde, 0xcc, 0x9b, 0x85,
	0xa7, 0x9f, 0x69, 0x36, 0x21, 0xa9, 0x78, 0x26, 0x45, 0x2e, 0xd8, 0x6f, 0x07, 0xba, 0x23, 0x49,
	0x51, 0x4e, 0x06, 0x0f, 0xe9, 0xc7, 0x82, 0x54, 0x8e, 0x3d, 0x38, 0xa2, 0x59, 0x94, 0xa4, 0x81,
	0x33, 0x70, 0x2e, 0xda, 0xa1, 0x09, 0xb0, 0x0f, 0x27, 0x59, 0xa4, 0xd4, 0xbd, 0x90, 0x77, 0x81,
	0xab, 0x13, 0x9b, 0x18, 0x5f, 0x42, 0x3b, 0x5b, 0x4c, 0xd2, 0x24, 0xbe, 0xa6, 0x55, 0xd0, 0xd4,
	0xc9, 0x2d, 0x80, 0xaf, 0x01, 0x32, 0x99, 0x2c, 0xa3, 0x9c, 0x8a, 0xb4, 0x37, 0x70, 0x2e, 0xfc,
	0xe1, 0x33, 0x3e, 0x92, 0xab, 0x2c, 0xa7, 0xbb, 0x2f, 0x26, 0x13, 0x5a, 0x47, 0xd8, 0x03, 0xf4,
	0xaa, 0x75, 0xa9, 0x4c, 0xcc, 0x15, 0x15, 0x25, 0xcc, 0x34, 0x32, 0xbe, 0x2c, 0x6b, 0xdb, 0xc4,
	0xf8, 0x0a, 0xfc, 0x28, 0x8e, 0x49, 0xa9, 0xaf, 0x62, 0x4a, 0x73, 0x5d, 0xa1, 0x3f, 0x3c, 0xe6,
	0x23, 0x21, 0xa6, 0x09, 0x85, 0x76, 0x0e, 0xcf, 0xc0, 0x9b, 0xd2, 0x4a, 0xe9, 0x42, 0xfd, 0xe1,
	0x11, 0xbf, 0xa6, 0x95, 0x0a, 0x35, 0xc4, 0xae, 0x00, 0x3f, 0x89, 0x6f, 0xc9, 0xfc, 0x1f, 0x1b,
	0xc2, 0xee, 0xa1, 0x5b, 0xb9, 0xe7, 0xbf, 0x09, 0x78, 0x07, 0x2f, 0x46, 0xdf, 0x29, 0x9e, 0x7e,
	0xdc, 0x1e, 0x5f, 0xab, 0x18, 0x54, 0x09, 0x0c, 0xbf, 0x0d, 0xb1, 0x9f, 0x10, 0xec, 0xfe, 0x5c,
	0x96, 0x1e, 0xc0, 0xb1, 0x5a, 0x68, 0x5c, 0xff, 0x79, 0x12, 0xae, 0xc3, 0x8a, 0x28, 0xf7, 0x71,
	0x51, 0xcd, 0xc3, 0xa2, 0xd8, 0x07, 0x68, 0x19, 0xb8, 0x68, 0xf7, 0x32, 0x4a, 0x17, 0xb4, 0x6e,
	0xb7, 0x0e, 0xf0, 0x1c, 0x80, 0x1e, 0xb2, 0x44, 0x46, 0x79, 0x22, 0x4c, 0x7b, 0x9a, 0xa1, 0x85,
	0xb0, 0x14, 0x5a, 0xa6, 0xdb, 0xd8, 0x01, 0x77, 0xd3, 0x5f, 0x77, 0x7c, 0xb9, 0x1d, 0x9f, 0x6b,
	0x8f, 0xef, 0x71, 0xcf, 0x9e, 0xef, 0x78, 0xb6, 0x5d, 0xb1, 0xe8, 0x15, 0x74, 0xaa, 0x06, 0xc6,
	0x53, 0x68, 0x16, 0x47, 0x0d, 0x6d, 0xf1, 0x89, 0x08, 0xde, 0x4d, 0x94, 0xe6, 0x25, 0xad, 0xfe,
	0xd6, 0xb5, 0xdd, 0x96, 0x74, 0xee, 0xf8, 0x96, 0xfd, 0x72, 0xc0, 0x2b, 0xc6, 0x57, 0x23, 0x74,
	0xea, 0x84, 0xc5, 0xf4, 0xca, 0xc8, 0xba, 0xd3, 0x86, 0xb4, 0x20, 0x13, 0x6e, 0x18, 0xb6, 0x40,
	0x55, 0xae, 0x57, 0x93, 0x3b, 0xfc, 0xe3, 0x40, 0xa7, 0x7c, 0x1c, 0x6e, 0x48, 0x2e, 0x93, 0x98,
	0xf0, 0x3d, 0x3c, 0xb1, 0x97, 0x10, 0x7b, 0x7c, 0xcf, 0x5b, 0xd1, 0x7f, 0xce, 0xf7, 0x6d, 0x2a,
	0x6b, 0xe0, 0x5b, 0xf0, 0xad, 0x0d, 0xc0, 0x2e, 0xdf, 0xdd, 0xab, 0x7e, 0x8f, 0xef, 0x59, 0x12,
	0xd6, 0xc0, 0x31, 0x9c, 0xd6, 0x7d, 0x88, 0x01, 0x3f, 0xe0, 0xeb, 0xfe, 0x19, 0x3f, 0x64, 0x5a,
	0xd6, 0x98, 0xb4, 0xf4, 0x53, 0xf7, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x4f, 0x12,
	0x98, 0xfb, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MembersServiceClient is the client API for MembersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MembersServiceClient interface {
	CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*CreateMemberResponse, error)
	LoginMember(ctx context.Context, in *LoginMemberRequest, opts ...grpc.CallOption) (*LoginMemberResponse, error)
	CheckAccessToken(ctx context.Context, in *CheckAccessTokenRequest, opts ...grpc.CallOption) (*CheckAccessTokenResponse, error)
}

type membersServiceClient struct {
	cc *grpc.ClientConn
}

func NewMembersServiceClient(cc *grpc.ClientConn) MembersServiceClient {
	return &membersServiceClient{cc}
}

func (c *membersServiceClient) CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*CreateMemberResponse, error) {
	out := new(CreateMemberResponse)
	err := c.cc.Invoke(ctx, "/MembersService/CreateMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membersServiceClient) LoginMember(ctx context.Context, in *LoginMemberRequest, opts ...grpc.CallOption) (*LoginMemberResponse, error) {
	out := new(LoginMemberResponse)
	err := c.cc.Invoke(ctx, "/MembersService/LoginMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membersServiceClient) CheckAccessToken(ctx context.Context, in *CheckAccessTokenRequest, opts ...grpc.CallOption) (*CheckAccessTokenResponse, error) {
	out := new(CheckAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/MembersService/CheckAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MembersServiceServer is the server API for MembersService service.
type MembersServiceServer interface {
	CreateMember(context.Context, *CreateMemberRequest) (*CreateMemberResponse, error)
	LoginMember(context.Context, *LoginMemberRequest) (*LoginMemberResponse, error)
	CheckAccessToken(context.Context, *CheckAccessTokenRequest) (*CheckAccessTokenResponse, error)
}

// UnimplementedMembersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMembersServiceServer struct {
}

func (*UnimplementedMembersServiceServer) CreateMember(ctx context.Context, req *CreateMemberRequest) (*CreateMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMember not implemented")
}
func (*UnimplementedMembersServiceServer) LoginMember(ctx context.Context, req *LoginMemberRequest) (*LoginMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginMember not implemented")
}
func (*UnimplementedMembersServiceServer) CheckAccessToken(ctx context.Context, req *CheckAccessTokenRequest) (*CheckAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAccessToken not implemented")
}

func RegisterMembersServiceServer(s *grpc.Server, srv MembersServiceServer) {
	s.RegisterService(&_MembersService_serviceDesc, srv)
}

func _MembersService_CreateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembersServiceServer).CreateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MembersService/CreateMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembersServiceServer).CreateMember(ctx, req.(*CreateMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MembersService_LoginMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembersServiceServer).LoginMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MembersService/LoginMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembersServiceServer).LoginMember(ctx, req.(*LoginMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MembersService_CheckAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembersServiceServer).CheckAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MembersService/CheckAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembersServiceServer).CheckAccessToken(ctx, req.(*CheckAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MembersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MembersService",
	HandlerType: (*MembersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMember",
			Handler:    _MembersService_CreateMember_Handler,
		},
		{
			MethodName: "LoginMember",
			Handler:    _MembersService_LoginMember_Handler,
		},
		{
			MethodName: "CheckAccessToken",
			Handler:    _MembersService_CheckAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Members.proto",
}
