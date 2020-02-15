// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/api/proto/banner.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Status int32

const (
	Status_default Status = 0
	Status_success Status = 1
	Status_error   Status = 2
)

var Status_name = map[int32]string{
	0: "default",
	1: "success",
	2: "error",
}

var Status_value = map[string]int32{
	"default": 0,
	"success": 1,
	"error":   2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Slot struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Slot) Reset()         { *m = Slot{} }
func (m *Slot) String() string { return proto.CompactTextString(m) }
func (*Slot) ProtoMessage()    {}
func (*Slot) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{1}
}

func (m *Slot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Slot.Unmarshal(m, b)
}
func (m *Slot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Slot.Marshal(b, m, deterministic)
}
func (m *Slot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slot.Merge(m, src)
}
func (m *Slot) XXX_Size() int {
	return xxx_messageInfo_Slot.Size(m)
}
func (m *Slot) XXX_DiscardUnknown() {
	xxx_messageInfo_Slot.DiscardUnknown(m)
}

var xxx_messageInfo_Slot proto.InternalMessageInfo

func (m *Slot) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Audience struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Audience) Reset()         { *m = Audience{} }
func (m *Audience) String() string { return proto.CompactTextString(m) }
func (*Audience) ProtoMessage()    {}
func (*Audience) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{2}
}

func (m *Audience) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Audience.Unmarshal(m, b)
}
func (m *Audience) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Audience.Marshal(b, m, deterministic)
}
func (m *Audience) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Audience.Merge(m, src)
}
func (m *Audience) XXX_Size() int {
	return xxx_messageInfo_Audience.Size(m)
}
func (m *Audience) XXX_DiscardUnknown() {
	xxx_messageInfo_Audience.DiscardUnknown(m)
}

var xxx_messageInfo_Audience proto.InternalMessageInfo

func (m *Audience) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Banner struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Banner) Reset()         { *m = Banner{} }
func (m *Banner) String() string { return proto.CompactTextString(m) }
func (*Banner) ProtoMessage()    {}
func (*Banner) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{3}
}

func (m *Banner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Banner.Unmarshal(m, b)
}
func (m *Banner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Banner.Marshal(b, m, deterministic)
}
func (m *Banner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Banner.Merge(m, src)
}
func (m *Banner) XXX_Size() int {
	return xxx_messageInfo_Banner.Size(m)
}
func (m *Banner) XXX_DiscardUnknown() {
	xxx_messageInfo_Banner.DiscardUnknown(m)
}

var xxx_messageInfo_Banner proto.InternalMessageInfo

func (m *Banner) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Date struct {
	Date                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{4}
}

func (m *Date) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Date.Unmarshal(m, b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Date.Marshal(b, m, deterministic)
}
func (m *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(m, src)
}
func (m *Date) XXX_Size() int {
	return xxx_messageInfo_Date.Size(m)
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type Response struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{5}
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

func (m *Response) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_default
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ResponseBannerMessage struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResponseBannerMessage) Reset()         { *m = ResponseBannerMessage{} }
func (m *ResponseBannerMessage) String() string { return proto.CompactTextString(m) }
func (*ResponseBannerMessage) ProtoMessage()    {}
func (*ResponseBannerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{6}
}

func (m *ResponseBannerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseBannerMessage.Unmarshal(m, b)
}
func (m *ResponseBannerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseBannerMessage.Marshal(b, m, deterministic)
}
func (m *ResponseBannerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseBannerMessage.Merge(m, src)
}
func (m *ResponseBannerMessage) XXX_Size() int {
	return xxx_messageInfo_ResponseBannerMessage.Size(m)
}
func (m *ResponseBannerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseBannerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseBannerMessage proto.InternalMessageInfo

func (m *ResponseBannerMessage) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type GetBannerRequestMessage struct {
	Audience             *Audience `protobuf:"bytes,1,opt,name=audience,proto3" json:"audience,omitempty"`
	Slot                 *Slot     `protobuf:"bytes,2,opt,name=slot,proto3" json:"slot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetBannerRequestMessage) Reset()         { *m = GetBannerRequestMessage{} }
func (m *GetBannerRequestMessage) String() string { return proto.CompactTextString(m) }
func (*GetBannerRequestMessage) ProtoMessage()    {}
func (*GetBannerRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{7}
}

func (m *GetBannerRequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBannerRequestMessage.Unmarshal(m, b)
}
func (m *GetBannerRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBannerRequestMessage.Marshal(b, m, deterministic)
}
func (m *GetBannerRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBannerRequestMessage.Merge(m, src)
}
func (m *GetBannerRequestMessage) XXX_Size() int {
	return xxx_messageInfo_GetBannerRequestMessage.Size(m)
}
func (m *GetBannerRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBannerRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetBannerRequestMessage proto.InternalMessageInfo

func (m *GetBannerRequestMessage) GetAudience() *Audience {
	if m != nil {
		return m.Audience
	}
	return nil
}

func (m *GetBannerRequestMessage) GetSlot() *Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

type GetBannerResponseMessage struct {
	Banner               *Banner   `protobuf:"bytes,1,opt,name=banner,proto3" json:"banner,omitempty"`
	Response             *Response `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetBannerResponseMessage) Reset()         { *m = GetBannerResponseMessage{} }
func (m *GetBannerResponseMessage) String() string { return proto.CompactTextString(m) }
func (*GetBannerResponseMessage) ProtoMessage()    {}
func (*GetBannerResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{8}
}

func (m *GetBannerResponseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBannerResponseMessage.Unmarshal(m, b)
}
func (m *GetBannerResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBannerResponseMessage.Marshal(b, m, deterministic)
}
func (m *GetBannerResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBannerResponseMessage.Merge(m, src)
}
func (m *GetBannerResponseMessage) XXX_Size() int {
	return xxx_messageInfo_GetBannerResponseMessage.Size(m)
}
func (m *GetBannerResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBannerResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetBannerResponseMessage proto.InternalMessageInfo

func (m *GetBannerResponseMessage) GetBanner() *Banner {
	if m != nil {
		return m.Banner
	}
	return nil
}

func (m *GetBannerResponseMessage) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type AddBannerToSlotRequestMessage struct {
	Audience             *Audience `protobuf:"bytes,1,opt,name=audience,proto3" json:"audience,omitempty"`
	Slot                 *Slot     `protobuf:"bytes,2,opt,name=slot,proto3" json:"slot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddBannerToSlotRequestMessage) Reset()         { *m = AddBannerToSlotRequestMessage{} }
func (m *AddBannerToSlotRequestMessage) String() string { return proto.CompactTextString(m) }
func (*AddBannerToSlotRequestMessage) ProtoMessage()    {}
func (*AddBannerToSlotRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{9}
}

func (m *AddBannerToSlotRequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBannerToSlotRequestMessage.Unmarshal(m, b)
}
func (m *AddBannerToSlotRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBannerToSlotRequestMessage.Marshal(b, m, deterministic)
}
func (m *AddBannerToSlotRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBannerToSlotRequestMessage.Merge(m, src)
}
func (m *AddBannerToSlotRequestMessage) XXX_Size() int {
	return xxx_messageInfo_AddBannerToSlotRequestMessage.Size(m)
}
func (m *AddBannerToSlotRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBannerToSlotRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddBannerToSlotRequestMessage proto.InternalMessageInfo

func (m *AddBannerToSlotRequestMessage) GetAudience() *Audience {
	if m != nil {
		return m.Audience
	}
	return nil
}

func (m *AddBannerToSlotRequestMessage) GetSlot() *Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

type DeleteBannerFromSlotRequestMessage struct {
	Audience             *Audience `protobuf:"bytes,1,opt,name=audience,proto3" json:"audience,omitempty"`
	Slot                 *Slot     `protobuf:"bytes,2,opt,name=slot,proto3" json:"slot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteBannerFromSlotRequestMessage) Reset()         { *m = DeleteBannerFromSlotRequestMessage{} }
func (m *DeleteBannerFromSlotRequestMessage) String() string { return proto.CompactTextString(m) }
func (*DeleteBannerFromSlotRequestMessage) ProtoMessage()    {}
func (*DeleteBannerFromSlotRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{10}
}

func (m *DeleteBannerFromSlotRequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBannerFromSlotRequestMessage.Unmarshal(m, b)
}
func (m *DeleteBannerFromSlotRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBannerFromSlotRequestMessage.Marshal(b, m, deterministic)
}
func (m *DeleteBannerFromSlotRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBannerFromSlotRequestMessage.Merge(m, src)
}
func (m *DeleteBannerFromSlotRequestMessage) XXX_Size() int {
	return xxx_messageInfo_DeleteBannerFromSlotRequestMessage.Size(m)
}
func (m *DeleteBannerFromSlotRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBannerFromSlotRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBannerFromSlotRequestMessage proto.InternalMessageInfo

func (m *DeleteBannerFromSlotRequestMessage) GetAudience() *Audience {
	if m != nil {
		return m.Audience
	}
	return nil
}

func (m *DeleteBannerFromSlotRequestMessage) GetSlot() *Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

type AddClickRequestMessage struct {
	Banner               *Banner   `protobuf:"bytes,1,opt,name=banner,proto3" json:"banner,omitempty"`
	Slot                 *Slot     `protobuf:"bytes,2,opt,name=slot,proto3" json:"slot,omitempty"`
	Audience             *Audience `protobuf:"bytes,3,opt,name=audience,proto3" json:"audience,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddClickRequestMessage) Reset()         { *m = AddClickRequestMessage{} }
func (m *AddClickRequestMessage) String() string { return proto.CompactTextString(m) }
func (*AddClickRequestMessage) ProtoMessage()    {}
func (*AddClickRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f76317f5ef7e5ef, []int{11}
}

func (m *AddClickRequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddClickRequestMessage.Unmarshal(m, b)
}
func (m *AddClickRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddClickRequestMessage.Marshal(b, m, deterministic)
}
func (m *AddClickRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddClickRequestMessage.Merge(m, src)
}
func (m *AddClickRequestMessage) XXX_Size() int {
	return xxx_messageInfo_AddClickRequestMessage.Size(m)
}
func (m *AddClickRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddClickRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddClickRequestMessage proto.InternalMessageInfo

func (m *AddClickRequestMessage) GetBanner() *Banner {
	if m != nil {
		return m.Banner
	}
	return nil
}

func (m *AddClickRequestMessage) GetSlot() *Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

func (m *AddClickRequestMessage) GetAudience() *Audience {
	if m != nil {
		return m.Audience
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.Status", Status_name, Status_value)
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*Slot)(nil), "proto.Slot")
	proto.RegisterType((*Audience)(nil), "proto.Audience")
	proto.RegisterType((*Banner)(nil), "proto.Banner")
	proto.RegisterType((*Date)(nil), "proto.Date")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*ResponseBannerMessage)(nil), "proto.ResponseBannerMessage")
	proto.RegisterType((*GetBannerRequestMessage)(nil), "proto.GetBannerRequestMessage")
	proto.RegisterType((*GetBannerResponseMessage)(nil), "proto.GetBannerResponseMessage")
	proto.RegisterType((*AddBannerToSlotRequestMessage)(nil), "proto.AddBannerToSlotRequestMessage")
	proto.RegisterType((*DeleteBannerFromSlotRequestMessage)(nil), "proto.DeleteBannerFromSlotRequestMessage")
	proto.RegisterType((*AddClickRequestMessage)(nil), "proto.AddClickRequestMessage")
}

func init() { proto.RegisterFile("server/api/proto/banner.proto", fileDescriptor_9f76317f5ef7e5ef) }

var fileDescriptor_9f76317f5ef7e5ef = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4b, 0x6f, 0xd3, 0x4c,
	0x14, 0xfd, 0x92, 0xe6, 0xd5, 0x9b, 0xaf, 0x25, 0x1a, 0x41, 0x30, 0x16, 0x21, 0xc8, 0xa2, 0x12,
	0x0f, 0xe1, 0x48, 0x41, 0x62, 0x9f, 0x36, 0x3c, 0x24, 0xc4, 0xc6, 0xa9, 0x04, 0x62, 0xc5, 0xc4,
	0x73, 0x9b, 0x9a, 0xda, 0x1e, 0x33, 0x33, 0xae, 0xc4, 0x8f, 0xe0, 0xef, 0xf1, 0x7b, 0x50, 0xe6,
	0x11, 0xea, 0xa4, 0xd4, 0x6c, 0xba, 0x1a, 0xcd, 0x9c, 0x3b, 0xe7, 0x9c, 0x7b, 0xe7, 0xd8, 0x30,
	0x92, 0x28, 0x2e, 0x51, 0x4c, 0x68, 0x91, 0x4c, 0x0a, 0xc1, 0x15, 0x9f, 0x2c, 0x69, 0x9e, 0xa3,
	0x08, 0xf5, 0x86, 0xb4, 0xf5, 0xe2, 0x8f, 0x57, 0x9c, 0xaf, 0x52, 0x34, 0x15, 0xcb, 0xf2, 0x6c,
	0xa2, 0x92, 0x0c, 0xa5, 0xa2, 0x59, 0x61, 0xea, 0x82, 0x2e, 0xb4, 0xdf, 0x64, 0x85, 0xfa, 0x11,
	0x0c, 0xa1, 0xb5, 0x48, 0xb9, 0x22, 0x87, 0xd0, 0x4c, 0x98, 0xd7, 0x78, 0xdc, 0x78, 0xba, 0x1f,
	0x35, 0x13, 0x16, 0xf8, 0xd0, 0x9b, 0x95, 0x2c, 0xc1, 0x3c, 0xc6, 0x1d, 0xcc, 0x83, 0xce, 0xb1,
	0x16, 0xdd, 0x41, 0x5e, 0x43, 0x6b, 0x4e, 0x15, 0x92, 0x10, 0x5a, 0x8c, 0x2a, 0xd4, 0x48, 0x7f,
	0xea, 0x87, 0xc6, 0x4e, 0xe8, 0xec, 0x84, 0xa7, 0xce, 0x4e, 0xa4, 0xeb, 0x82, 0x0f, 0xd0, 0x8b,
	0x50, 0x16, 0x3c, 0x97, 0x48, 0x8e, 0xa0, 0x23, 0x15, 0x55, 0xa5, 0xd4, 0xb7, 0x0f, 0xa7, 0x07,
	0xe6, 0x5a, 0xb8, 0xd0, 0x87, 0x91, 0x05, 0x89, 0x07, 0xdd, 0x0c, 0xa5, 0xa4, 0x2b, 0xf4, 0x9a,
	0x5a, 0xdf, 0x6d, 0x83, 0x39, 0xdc, 0x73, 0x64, 0xc6, 0xe6, 0x47, 0x03, 0x90, 0x17, 0xd0, 0x13,
	0x16, 0xb0, 0xce, 0xee, 0x58, 0x6e, 0x57, 0x1f, 0x6d, 0x0a, 0x82, 0x15, 0xdc, 0x7f, 0x87, 0xca,
	0x10, 0x44, 0xf8, 0xbd, 0x44, 0xa9, 0xae, 0xf0, 0x50, 0x3b, 0x9b, 0x2d, 0x1e, 0x37, 0xb2, 0x68,
	0x53, 0x40, 0xc6, 0xd0, 0x92, 0x29, 0x57, 0xda, 0x64, 0x7f, 0xda, 0x77, 0xcd, 0xa4, 0x5c, 0x45,
	0x1a, 0x08, 0x72, 0xf0, 0xae, 0x08, 0x19, 0x75, 0xa7, 0x74, 0x04, 0x1d, 0xf3, 0xbc, 0x56, 0xc7,
	0xcd, 0xc2, 0x56, 0x5b, 0xb0, 0xd2, 0x58, 0xb3, 0xae, 0xb1, 0x0c, 0x46, 0x33, 0xc6, 0x0c, 0xc3,
	0x29, 0xd7, 0x46, 0x6e, 0xb3, 0x3d, 0x01, 0xc1, 0x1c, 0x53, 0x54, 0xf6, 0x2d, 0xde, 0x0a, 0x9e,
	0xdd, 0xba, 0xe6, 0xcf, 0x06, 0x0c, 0x67, 0x8c, 0x9d, 0xa4, 0x49, 0x7c, 0xb1, 0x25, 0xf4, 0x8f,
	0x13, 0xad, 0x93, 0xa8, 0x18, 0xde, 0xab, 0x31, 0xfc, 0xfc, 0x25, 0x74, 0x4c, 0x7a, 0x49, 0x1f,
	0xba, 0x0c, 0xcf, 0x68, 0x99, 0xaa, 0xc1, 0x7f, 0xeb, 0x8d, 0x2c, 0xe3, 0x18, 0xa5, 0x1c, 0x34,
	0xc8, 0x3e, 0xb4, 0x51, 0x08, 0x2e, 0x06, 0xcd, 0xe9, 0xaf, 0x3d, 0x38, 0x30, 0x7e, 0x16, 0x28,
	0x2e, 0x93, 0x18, 0xc9, 0x1c, 0x86, 0x12, 0x73, 0xf6, 0x1e, 0x69, 0xaa, 0xce, 0x4f, 0xce, 0x31,
	0xbe, 0x70, 0xfd, 0xfc, 0x6f, 0x55, 0xf5, 0xd7, 0xec, 0x3f, 0xdc, 0x7a, 0xf6, 0x6a, 0xfe, 0x3f,
	0xc1, 0xdd, 0x35, 0xcb, 0x26, 0x6d, 0xee, 0xfc, 0x91, 0xbd, 0xf5, 0x97, 0xbc, 0xfb, 0xe3, 0x5d,
	0xbc, 0x1a, 0xd3, 0xaf, 0xe0, 0xaf, 0x89, 0xb7, 0x62, 0xe5, 0xd0, 0x27, 0x6e, 0x30, 0x37, 0xa5,
	0xae, 0xc6, 0xfa, 0x37, 0x18, 0xaf, 0x15, 0xae, 0x4b, 0x92, 0x2b, 0x79, 0x66, 0x09, 0xea, 0xd3,
	0x56, 0xa3, 0xf5, 0x19, 0x1e, 0xd8, 0x6e, 0x74, 0x80, 0xaa, 0xe0, 0xe8, 0x4f, 0x33, 0xd7, 0xc4,
	0xeb, 0x66, 0xe6, 0xe3, 0xee, 0x17, 0xf3, 0x7f, 0x5e, 0x76, 0xf4, 0xf2, 0xea, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x99, 0x38, 0x70, 0x2b, 0xce, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BannerServiceClient is the client API for BannerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BannerServiceClient interface {
	SendHealthCheckMessage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseBannerMessage, error)
	SendGetBannerMessage(ctx context.Context, in *GetBannerRequestMessage, opts ...grpc.CallOption) (*GetBannerResponseMessage, error)
	SendAddBannerToSlotMessage(ctx context.Context, in *AddBannerToSlotRequestMessage, opts ...grpc.CallOption) (*ResponseBannerMessage, error)
	SendDeleteBannerFromSlotMessage(ctx context.Context, in *DeleteBannerFromSlotRequestMessage, opts ...grpc.CallOption) (*ResponseBannerMessage, error)
	SendAddClickBannerMessage(ctx context.Context, in *AddClickRequestMessage, opts ...grpc.CallOption) (*ResponseBannerMessage, error)
}

type bannerServiceClient struct {
	cc *grpc.ClientConn
}

func NewBannerServiceClient(cc *grpc.ClientConn) BannerServiceClient {
	return &bannerServiceClient{cc}
}

func (c *bannerServiceClient) SendHealthCheckMessage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResponseBannerMessage, error) {
	out := new(ResponseBannerMessage)
	err := c.cc.Invoke(ctx, "/proto.BannerService/sendHealthCheckMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) SendGetBannerMessage(ctx context.Context, in *GetBannerRequestMessage, opts ...grpc.CallOption) (*GetBannerResponseMessage, error) {
	out := new(GetBannerResponseMessage)
	err := c.cc.Invoke(ctx, "/proto.BannerService/sendGetBannerMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) SendAddBannerToSlotMessage(ctx context.Context, in *AddBannerToSlotRequestMessage, opts ...grpc.CallOption) (*ResponseBannerMessage, error) {
	out := new(ResponseBannerMessage)
	err := c.cc.Invoke(ctx, "/proto.BannerService/sendAddBannerToSlotMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) SendDeleteBannerFromSlotMessage(ctx context.Context, in *DeleteBannerFromSlotRequestMessage, opts ...grpc.CallOption) (*ResponseBannerMessage, error) {
	out := new(ResponseBannerMessage)
	err := c.cc.Invoke(ctx, "/proto.BannerService/sendDeleteBannerFromSlotMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) SendAddClickBannerMessage(ctx context.Context, in *AddClickRequestMessage, opts ...grpc.CallOption) (*ResponseBannerMessage, error) {
	out := new(ResponseBannerMessage)
	err := c.cc.Invoke(ctx, "/proto.BannerService/sendAddClickBannerMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BannerServiceServer is the server API for BannerService service.
type BannerServiceServer interface {
	SendHealthCheckMessage(context.Context, *Empty) (*ResponseBannerMessage, error)
	SendGetBannerMessage(context.Context, *GetBannerRequestMessage) (*GetBannerResponseMessage, error)
	SendAddBannerToSlotMessage(context.Context, *AddBannerToSlotRequestMessage) (*ResponseBannerMessage, error)
	SendDeleteBannerFromSlotMessage(context.Context, *DeleteBannerFromSlotRequestMessage) (*ResponseBannerMessage, error)
	SendAddClickBannerMessage(context.Context, *AddClickRequestMessage) (*ResponseBannerMessage, error)
}

// UnimplementedBannerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBannerServiceServer struct {
}

func (*UnimplementedBannerServiceServer) SendHealthCheckMessage(ctx context.Context, req *Empty) (*ResponseBannerMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHealthCheckMessage not implemented")
}
func (*UnimplementedBannerServiceServer) SendGetBannerMessage(ctx context.Context, req *GetBannerRequestMessage) (*GetBannerResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGetBannerMessage not implemented")
}
func (*UnimplementedBannerServiceServer) SendAddBannerToSlotMessage(ctx context.Context, req *AddBannerToSlotRequestMessage) (*ResponseBannerMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAddBannerToSlotMessage not implemented")
}
func (*UnimplementedBannerServiceServer) SendDeleteBannerFromSlotMessage(ctx context.Context, req *DeleteBannerFromSlotRequestMessage) (*ResponseBannerMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDeleteBannerFromSlotMessage not implemented")
}
func (*UnimplementedBannerServiceServer) SendAddClickBannerMessage(ctx context.Context, req *AddClickRequestMessage) (*ResponseBannerMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAddClickBannerMessage not implemented")
}

func RegisterBannerServiceServer(s *grpc.Server, srv BannerServiceServer) {
	s.RegisterService(&_BannerService_serviceDesc, srv)
}

func _BannerService_SendHealthCheckMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).SendHealthCheckMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BannerService/SendHealthCheckMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).SendHealthCheckMessage(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_SendGetBannerMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBannerRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).SendGetBannerMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BannerService/SendGetBannerMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).SendGetBannerMessage(ctx, req.(*GetBannerRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_SendAddBannerToSlotMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBannerToSlotRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).SendAddBannerToSlotMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BannerService/SendAddBannerToSlotMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).SendAddBannerToSlotMessage(ctx, req.(*AddBannerToSlotRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_SendDeleteBannerFromSlotMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBannerFromSlotRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).SendDeleteBannerFromSlotMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BannerService/SendDeleteBannerFromSlotMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).SendDeleteBannerFromSlotMessage(ctx, req.(*DeleteBannerFromSlotRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_SendAddClickBannerMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClickRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).SendAddClickBannerMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BannerService/SendAddClickBannerMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).SendAddClickBannerMessage(ctx, req.(*AddClickRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _BannerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BannerService",
	HandlerType: (*BannerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendHealthCheckMessage",
			Handler:    _BannerService_SendHealthCheckMessage_Handler,
		},
		{
			MethodName: "sendGetBannerMessage",
			Handler:    _BannerService_SendGetBannerMessage_Handler,
		},
		{
			MethodName: "sendAddBannerToSlotMessage",
			Handler:    _BannerService_SendAddBannerToSlotMessage_Handler,
		},
		{
			MethodName: "sendDeleteBannerFromSlotMessage",
			Handler:    _BannerService_SendDeleteBannerFromSlotMessage_Handler,
		},
		{
			MethodName: "sendAddClickBannerMessage",
			Handler:    _BannerService_SendAddClickBannerMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/api/proto/banner.proto",
}
