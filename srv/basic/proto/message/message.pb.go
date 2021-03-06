// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/message/message.proto

package go_micro_srv_basic_message

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

type ChangeStatusRequest_Status int32

const (
	ChangeStatusRequest_Unread ChangeStatusRequest_Status = 0
	ChangeStatusRequest_Readed ChangeStatusRequest_Status = 10
	ChangeStatusRequest_Trash  ChangeStatusRequest_Status = 20
)

var ChangeStatusRequest_Status_name = map[int32]string{
	0:  "Unread",
	10: "Readed",
	20: "Trash",
}

var ChangeStatusRequest_Status_value = map[string]int32{
	"Unread": 0,
	"Readed": 10,
	"Trash":  20,
}

func (x ChangeStatusRequest_Status) String() string {
	return proto.EnumName(ChangeStatusRequest_Status_name, int32(x))
}

func (ChangeStatusRequest_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f274517418484e40, []int{2, 0}
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f274517418484e40, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

type SendRequest struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	ToList               []string `protobuf:"bytes,2,rep,name=to_list,json=toList,proto3" json:"to_list,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f274517418484e40, []int{1}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *SendRequest) GetToList() []string {
	if m != nil {
		return m.ToList
	}
	return nil
}

func (m *SendRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SendRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ChangeStatusRequest struct {
	Status               ChangeStatusRequest_Status `protobuf:"varint,1,opt,name=status,proto3,enum=go.micro.srv.basic.message.ChangeStatusRequest_Status" json:"status,omitempty"`
	To                   string                     `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Id                   int64                      `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ChangeStatusRequest) Reset()         { *m = ChangeStatusRequest{} }
func (m *ChangeStatusRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusRequest) ProtoMessage()    {}
func (*ChangeStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f274517418484e40, []int{2}
}

func (m *ChangeStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusRequest.Unmarshal(m, b)
}
func (m *ChangeStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusRequest.Marshal(b, m, deterministic)
}
func (m *ChangeStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusRequest.Merge(m, src)
}
func (m *ChangeStatusRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusRequest.Size(m)
}
func (m *ChangeStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusRequest proto.InternalMessageInfo

func (m *ChangeStatusRequest) GetStatus() ChangeStatusRequest_Status {
	if m != nil {
		return m.Status
	}
	return ChangeStatusRequest_Unread
}

func (m *ChangeStatusRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ChangeStatusRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type InitRequest struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitRequest) Reset()         { *m = InitRequest{} }
func (m *InitRequest) String() string { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()    {}
func (*InitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f274517418484e40, []int{3}
}

func (m *InitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitRequest.Unmarshal(m, b)
}
func (m *InitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitRequest.Marshal(b, m, deterministic)
}
func (m *InitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitRequest.Merge(m, src)
}
func (m *InitRequest) XXX_Size() int {
	return xxx_messageInfo_InitRequest.Size(m)
}
func (m *InitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitRequest proto.InternalMessageInfo

func (m *InitRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type InitResponse struct {
	To                   string                  `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Unread               []*InitResponse_Message `protobuf:"bytes,2,rep,name=Unread,proto3" json:"Unread,omitempty"`
	Readed               []*InitResponse_Message `protobuf:"bytes,3,rep,name=Readed,proto3" json:"Readed,omitempty"`
	Trash                []*InitResponse_Message `protobuf:"bytes,4,rep,name=Trash,proto3" json:"Trash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *InitResponse) Reset()         { *m = InitResponse{} }
func (m *InitResponse) String() string { return proto.CompactTextString(m) }
func (*InitResponse) ProtoMessage()    {}
func (*InitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f274517418484e40, []int{4}
}

func (m *InitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitResponse.Unmarshal(m, b)
}
func (m *InitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitResponse.Marshal(b, m, deterministic)
}
func (m *InitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitResponse.Merge(m, src)
}
func (m *InitResponse) XXX_Size() int {
	return xxx_messageInfo_InitResponse.Size(m)
}
func (m *InitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitResponse proto.InternalMessageInfo

func (m *InitResponse) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *InitResponse) GetUnread() []*InitResponse_Message {
	if m != nil {
		return m.Unread
	}
	return nil
}

func (m *InitResponse) GetReaded() []*InitResponse_Message {
	if m != nil {
		return m.Readed
	}
	return nil
}

func (m *InitResponse) GetTrash() []*InitResponse_Message {
	if m != nil {
		return m.Trash
	}
	return nil
}

type InitResponse_Message struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitResponse_Message) Reset()         { *m = InitResponse_Message{} }
func (m *InitResponse_Message) String() string { return proto.CompactTextString(m) }
func (*InitResponse_Message) ProtoMessage()    {}
func (*InitResponse_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f274517418484e40, []int{4, 0}
}

func (m *InitResponse_Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitResponse_Message.Unmarshal(m, b)
}
func (m *InitResponse_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitResponse_Message.Marshal(b, m, deterministic)
}
func (m *InitResponse_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitResponse_Message.Merge(m, src)
}
func (m *InitResponse_Message) XXX_Size() int {
	return xxx_messageInfo_InitResponse_Message.Size(m)
}
func (m *InitResponse_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_InitResponse_Message.DiscardUnknown(m)
}

var xxx_messageInfo_InitResponse_Message proto.InternalMessageInfo

func (m *InitResponse_Message) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InitResponse_Message) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *InitResponse_Message) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type GetRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f274517418484e40, []int{5}
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

func (m *GetRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f274517418484e40, []int{6}
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

func (m *GetResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *GetResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GetResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterEnum("go.micro.srv.basic.message.ChangeStatusRequest_Status", ChangeStatusRequest_Status_name, ChangeStatusRequest_Status_value)
	proto.RegisterType((*Response)(nil), "go.micro.srv.basic.message.Response")
	proto.RegisterType((*SendRequest)(nil), "go.micro.srv.basic.message.SendRequest")
	proto.RegisterType((*ChangeStatusRequest)(nil), "go.micro.srv.basic.message.ChangeStatusRequest")
	proto.RegisterType((*InitRequest)(nil), "go.micro.srv.basic.message.InitRequest")
	proto.RegisterType((*InitResponse)(nil), "go.micro.srv.basic.message.InitResponse")
	proto.RegisterType((*InitResponse_Message)(nil), "go.micro.srv.basic.message.InitResponse.Message")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.basic.message.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "go.micro.srv.basic.message.GetResponse")
}

func init() { proto.RegisterFile("proto/message/message.proto", fileDescriptor_f274517418484e40) }

var fileDescriptor_f274517418484e40 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xef, 0x6a, 0xd4, 0x40,
	0x10, 0xbf, 0xfc, 0x69, 0xce, 0x9b, 0x2b, 0xe5, 0x18, 0x0b, 0x2e, 0x51, 0xa1, 0x2c, 0x62, 0x0f,
	0x84, 0xad, 0x9c, 0xe0, 0x0b, 0x14, 0xac, 0x82, 0xfa, 0x21, 0x55, 0x14, 0xfc, 0x20, 0xdb, 0xcb,
	0x9a, 0xae, 0xf4, 0xb2, 0x35, 0x3b, 0xf5, 0x01, 0x7c, 0x24, 0x5f, 0xc0, 0x57, 0x93, 0xec, 0x66,
	0x7b, 0x41, 0xda, 0x78, 0xbd, 0x4f, 0xd9, 0x99, 0x9d, 0xf9, 0xcd, 0xcc, 0x6f, 0x7f, 0x13, 0x78,
	0x78, 0xd9, 0x18, 0x32, 0x47, 0x2b, 0x65, 0xad, 0xac, 0x54, 0xf8, 0x0a, 0xe7, 0xc5, 0xbc, 0x32,
	0x62, 0xa5, 0x97, 0x8d, 0x11, 0xb6, 0xf9, 0x29, 0xce, 0xa4, 0xd5, 0x4b, 0xd1, 0x45, 0x70, 0x80,
	0x7b, 0x85, 0xb2, 0x97, 0xa6, 0xb6, 0x8a, 0x7f, 0x87, 0xe9, 0xa9, 0xaa, 0xcb, 0x42, 0xfd, 0xb8,
	0x52, 0x96, 0x10, 0x21, 0xfd, 0xd6, 0x98, 0x15, 0x8b, 0x0e, 0xa2, 0xf9, 0xa4, 0x70, 0x67, 0x7c,
	0x00, 0x63, 0x32, 0x5f, 0x2f, 0xb4, 0x25, 0x16, 0x1f, 0x24, 0xf3, 0x49, 0x91, 0x91, 0x79, 0xab,
	0x2d, 0xe1, 0x3e, 0xec, 0x90, 0xa6, 0x0b, 0xc5, 0x12, 0x17, 0xed, 0x0d, 0x64, 0x30, 0x5e, 0x9a,
	0x9a, 0x54, 0x4d, 0x2c, 0x75, 0xfe, 0x60, 0xf2, 0xdf, 0x11, 0xdc, 0x3f, 0x3e, 0x97, 0x75, 0xa5,
	0x4e, 0x49, 0xd2, 0x95, 0x0d, 0x45, 0xdf, 0x43, 0x66, 0x9d, 0xc3, 0x95, 0xdd, 0x5b, 0xbc, 0x14,
	0xb7, 0x37, 0x2f, 0x6e, 0x00, 0x10, 0x9d, 0xd5, 0xa1, 0xe0, 0x1e, 0xc4, 0x64, 0x58, 0xec, 0x8a,
	0xc7, 0x64, 0x5a, 0x5b, 0x97, 0xae, 0xc9, 0xa4, 0x88, 0x75, 0xc9, 0x9f, 0x41, 0xe6, 0x33, 0x10,
	0x20, 0xfb, 0x58, 0x37, 0x4a, 0x96, 0xb3, 0x51, 0x7b, 0x2e, 0x94, 0x2c, 0x55, 0x39, 0x03, 0x9c,
	0xc0, 0xce, 0x87, 0x46, 0xda, 0xf3, 0xd9, 0x3e, 0x7f, 0x0c, 0xd3, 0x37, 0xb5, 0xa6, 0xd0, 0xab,
	0xc7, 0x8e, 0x02, 0x36, 0xff, 0x13, 0xc3, 0xae, 0xbf, 0xf7, 0x84, 0xfe, 0x1b, 0x80, 0xaf, 0x43,
	0x09, 0x47, 0xde, 0x74, 0xf1, 0x7c, 0x68, 0xb8, 0x3e, 0x92, 0x78, 0xe7, 0x9d, 0x45, 0x97, 0xdf,
	0x22, 0xf9, 0x06, 0x59, 0xb2, 0x2d, 0x92, 0xcf, 0xc7, 0x57, 0xdd, 0x78, 0x2c, 0xdd, 0x12, 0xc8,
	0xa7, 0xe7, 0xc7, 0x30, 0xee, 0x3c, 0x1d, 0xc7, 0x51, 0xe0, 0xf8, 0x5a, 0x48, 0x71, 0x4f, 0x48,
	0x37, 0xea, 0x85, 0x3f, 0x02, 0x38, 0x51, 0x7d, 0x7e, 0xfb, 0x38, 0x5c, 0xc2, 0xd4, 0xdd, 0xae,
	0xd9, 0xdd, 0xae, 0xcc, 0xed, 0xb2, 0x5c, 0xfc, 0x4a, 0xd6, 0x63, 0x7c, 0x82, 0xb4, 0x5d, 0x07,
	0x3c, 0x1c, 0xa2, 0xa4, 0xb7, 0x30, 0xf9, 0x93, 0xa1, 0xc0, 0xeb, 0x2d, 0x1b, 0x61, 0x05, 0xbb,
	0x7d, 0xe5, 0xe2, 0xd1, 0x1d, 0x35, 0xbe, 0x71, 0xa1, 0x2f, 0x90, 0xb6, 0x4f, 0x36, 0x3c, 0x41,
	0x4f, 0xd1, 0xf9, 0x7c, 0xd3, 0xd7, 0xe7, 0x23, 0xfc, 0x0c, 0xc9, 0x89, 0x22, 0x7c, 0x3a, 0x94,
	0xb2, 0x7e, 0xcc, 0xfc, 0xf0, 0xbf, 0x71, 0x01, 0xf9, 0x2c, 0x73, 0xbf, 0xad, 0x17, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x04, 0xa1, 0xbe, 0xaa, 0xd5, 0x04, 0x00, 0x00,
}
