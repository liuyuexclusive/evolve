// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_basic_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AuthRequest struct {
	//id（账号）
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//key (token)
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AuthRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type AuthResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidateRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateRequest) Reset()         { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()    {}
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *ValidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateRequest.Unmarshal(m, b)
}
func (m *ValidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateRequest.Marshal(b, m, deterministic)
}
func (m *ValidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateRequest.Merge(m, src)
}
func (m *ValidateRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateRequest.Size(m)
}
func (m *ValidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateRequest proto.InternalMessageInfo

func (m *ValidateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidateResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateResponse) Reset()         { *m = ValidateResponse{} }
func (m *ValidateResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateResponse) ProtoMessage()    {}
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *ValidateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateResponse.Unmarshal(m, b)
}
func (m *ValidateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateResponse.Marshal(b, m, deterministic)
}
func (m *ValidateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateResponse.Merge(m, src)
}
func (m *ValidateResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateResponse.Size(m)
}
func (m *ValidateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateResponse proto.InternalMessageInfo

func (m *ValidateResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Access               []string `protobuf:"bytes,2,rep,name=access,proto3" json:"access,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetResponse) GetAccess() []string {
	if m != nil {
		return m.Access
	}
	return nil
}

func (m *GetResponse) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type TestRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{6}
}

func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TestResponse struct {
	Res                  string   `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{7}
}

func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (m *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(m, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "go.micro.srv.basic.user.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "go.micro.srv.basic.user.AuthResponse")
	proto.RegisterType((*ValidateRequest)(nil), "go.micro.srv.basic.user.ValidateRequest")
	proto.RegisterType((*ValidateResponse)(nil), "go.micro.srv.basic.user.ValidateResponse")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.basic.user.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "go.micro.srv.basic.user.GetResponse")
	proto.RegisterType((*TestRequest)(nil), "go.micro.srv.basic.user.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "go.micro.srv.basic.user.TestResponse")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0xc6, 0x5b, 0xc0, 0xc6, 0x4e, 0x1b, 0x6d, 0x36, 0xfe, 0x21, 0x3d, 0xe1, 0x5a, 0x15, 0x2f,
	0xdb, 0x44, 0x9f, 0xc0, 0x53, 0xcf, 0x12, 0xf1, 0xbe, 0x85, 0x89, 0x92, 0x5a, 0xb6, 0xee, 0x2e,
	0x4d, 0x7c, 0x66, 0x5f, 0xc2, 0xb0, 0x2c, 0x2d, 0x31, 0x05, 0xbd, 0x90, 0x99, 0xe1, 0x37, 0xdf,
	0x24, 0xdf, 0xb7, 0x70, 0xbe, 0x91, 0x42, 0x8b, 0x79, 0xa1, 0x50, 0x9a, 0x0f, 0x33, 0x3d, 0xb9,
	0x7c, 0x13, 0x6c, 0x9d, 0x25, 0x52, 0x30, 0x25, 0xb7, 0x6c, 0xc9, 0x55, 0x96, 0xb0, 0xf2, 0x37,
	0x9d, 0xc3, 0xe8, 0xa9, 0xd0, 0xef, 0x11, 0x7e, 0x16, 0xa8, 0x34, 0x39, 0x01, 0x27, 0x4b, 0xfd,
	0x7e, 0xd0, 0x0f, 0x87, 0x91, 0x93, 0xa5, 0x64, 0x02, 0xee, 0x0a, 0xbf, 0x7c, 0xc7, 0x0c, 0xca,
	0x92, 0xce, 0x60, 0x5c, 0x2d, 0xa8, 0x8d, 0xc8, 0x15, 0x92, 0x33, 0x38, 0xd2, 0x62, 0x85, 0xb9,
	0x5d, 0xaa, 0x1a, 0x7a, 0x07, 0xa7, 0xaf, 0xfc, 0x23, 0x4b, 0xb9, 0xc6, 0x5a, 0xfa, 0x30, 0x78,
	0x0b, 0x93, 0x3d, 0x68, 0x25, 0x09, 0x78, 0x39, 0x5f, 0xa3, 0x05, 0x4d, 0x4d, 0x03, 0x80, 0x05,
	0xea, 0x5a, 0xeb, 0x10, 0xf1, 0x0c, 0x23, 0x43, 0xb4, 0x8b, 0x90, 0x0b, 0x18, 0xf0, 0x24, 0x41,
	0xa5, 0x7c, 0x27, 0x70, 0xc3, 0x61, 0x64, 0x3b, 0x33, 0xdf, 0x72, 0xcd, 0xa5, 0xef, 0x1a, 0xda,
	0x76, 0xf4, 0x0a, 0x46, 0x2f, 0xa8, 0x3a, 0xaf, 0x06, 0x30, 0xae, 0x10, 0x7b, 0x76, 0x02, 0xae,
	0x44, 0x65, 0x91, 0xb2, 0x7c, 0xf8, 0x76, 0xc0, 0x8b, 0x15, 0x4a, 0x12, 0x83, 0x57, 0x3a, 0x47,
	0x66, 0xac, 0x25, 0x0c, 0xd6, 0x48, 0x62, 0x7a, 0xf3, 0x07, 0x55, 0xdd, 0xa3, 0x3d, 0xc2, 0xe1,
	0xb8, 0x76, 0x90, 0x84, 0xad, 0x4b, 0xbf, 0xd2, 0x98, 0xde, 0xff, 0x83, 0xdc, 0x9d, 0x88, 0xc0,
	0x5d, 0xa0, 0x26, 0xd7, 0xad, 0x3b, 0xfb, 0x68, 0xa6, 0xb3, 0x6e, 0x68, 0xa7, 0x19, 0x83, 0x57,
	0x1a, 0xd7, 0xe1, 0x46, 0xc3, 0xfa, 0x0e, 0x37, 0x9a, 0xee, 0xd3, 0xde, 0x72, 0x60, 0xde, 0xfb,
	0xe3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xb7, 0xc6, 0xb6, 0x08, 0x03, 0x00, 0x00,
}
