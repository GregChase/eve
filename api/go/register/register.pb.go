// Code generated by protoc-gen-go. DO NOT EDIT.
// source: register.proto

package register

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

// XXX is this used? Not in client.go; deprecate
type ZRegisterResult int32

const (
	ZRegisterResult_ZRegNone        ZRegisterResult = 0
	ZRegisterResult_ZRegSuccess     ZRegisterResult = 1
	ZRegisterResult_ZRegNotActive   ZRegisterResult = 2
	ZRegisterResult_ZRegAlreadyDone ZRegisterResult = 3
	ZRegisterResult_ZRegDeviceNA    ZRegisterResult = 4
	ZRegisterResult_ZRegFailed      ZRegisterResult = 5
)

var ZRegisterResult_name = map[int32]string{
	0: "ZRegNone",
	1: "ZRegSuccess",
	2: "ZRegNotActive",
	3: "ZRegAlreadyDone",
	4: "ZRegDeviceNA",
	5: "ZRegFailed",
}

var ZRegisterResult_value = map[string]int32{
	"ZRegNone":        0,
	"ZRegSuccess":     1,
	"ZRegNotActive":   2,
	"ZRegAlreadyDone": 3,
	"ZRegDeviceNA":    4,
	"ZRegFailed":      5,
}

func (x ZRegisterResult) String() string {
	return proto.EnumName(ZRegisterResult_name, int32(x))
}

func (ZRegisterResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}

// XXX is this used? Not in client.go; deprecate
type ZRegisterResp struct {
	Result               ZRegisterResult `protobuf:"varint,2,opt,name=result,proto3,enum=ZRegisterResult" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ZRegisterResp) Reset()         { *m = ZRegisterResp{} }
func (m *ZRegisterResp) String() string { return proto.CompactTextString(m) }
func (*ZRegisterResp) ProtoMessage()    {}
func (*ZRegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}

func (m *ZRegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZRegisterResp.Unmarshal(m, b)
}
func (m *ZRegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZRegisterResp.Marshal(b, m, deterministic)
}
func (m *ZRegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZRegisterResp.Merge(m, src)
}
func (m *ZRegisterResp) XXX_Size() int {
	return xxx_messageInfo_ZRegisterResp.Size(m)
}
func (m *ZRegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ZRegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_ZRegisterResp proto.InternalMessageInfo

func (m *ZRegisterResp) GetResult() ZRegisterResult {
	if m != nil {
		return m.Result
	}
	return ZRegisterResult_ZRegNone
}

type ZRegisterMsg struct {
	OnBoardKey           string   `protobuf:"bytes,1,opt,name=onBoardKey,proto3" json:"onBoardKey,omitempty"`
	PemCert              []byte   `protobuf:"bytes,2,opt,name=pemCert,proto3" json:"pemCert,omitempty"`
	Serial               string   `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZRegisterMsg) Reset()         { *m = ZRegisterMsg{} }
func (m *ZRegisterMsg) String() string { return proto.CompactTextString(m) }
func (*ZRegisterMsg) ProtoMessage()    {}
func (*ZRegisterMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{1}
}

func (m *ZRegisterMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZRegisterMsg.Unmarshal(m, b)
}
func (m *ZRegisterMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZRegisterMsg.Marshal(b, m, deterministic)
}
func (m *ZRegisterMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZRegisterMsg.Merge(m, src)
}
func (m *ZRegisterMsg) XXX_Size() int {
	return xxx_messageInfo_ZRegisterMsg.Size(m)
}
func (m *ZRegisterMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ZRegisterMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ZRegisterMsg proto.InternalMessageInfo

func (m *ZRegisterMsg) GetOnBoardKey() string {
	if m != nil {
		return m.OnBoardKey
	}
	return ""
}

func (m *ZRegisterMsg) GetPemCert() []byte {
	if m != nil {
		return m.PemCert
	}
	return nil
}

func (m *ZRegisterMsg) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func init() {
	proto.RegisterEnum("ZRegisterResult", ZRegisterResult_name, ZRegisterResult_value)
	proto.RegisterType((*ZRegisterResp)(nil), "ZRegisterResp")
	proto.RegisterType((*ZRegisterMsg)(nil), "ZRegisterMsg")
}

func init() { proto.RegisterFile("register.proto", fileDescriptor_1303fe8288f4efb6) }

var fileDescriptor_1303fe8288f4efb6 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0xed, 0xa6, 0x55, 0x3f, 0x6b, 0x17, 0x3f, 0x41, 0x7a, 0x92, 0xb1, 0x83, 0x14, 0xc1,
	0x16, 0xf4, 0xe4, 0xb1, 0x73, 0x78, 0x11, 0x7b, 0xa8, 0xb7, 0x9d, 0xec, 0xd2, 0xcf, 0x18, 0xc8,
	0x96, 0x92, 0xa4, 0x85, 0xf9, 0xd7, 0x4b, 0xd7, 0x4e, 0x86, 0xc7, 0xf7, 0x7b, 0xbc, 0x97, 0x7c,
	0x0f, 0x42, 0x43, 0x42, 0x5a, 0x47, 0x26, 0xa9, 0x8d, 0x76, 0x7a, 0xf6, 0x0c, 0x97, 0xcb, 0x62,
	0x40, 0x05, 0xd9, 0x1a, 0x63, 0xf0, 0x0d, 0xd9, 0x46, 0xb9, 0x68, 0x34, 0xf5, 0xe2, 0xf0, 0x91,
	0x25, 0x87, 0x7e, 0xa3, 0x5c, 0x31, 0xf8, 0xb3, 0x4f, 0x08, 0xfe, 0xac, 0x77, 0x2b, 0xf0, 0x16,
	0x40, 0x6f, 0xe6, 0xba, 0x34, 0xd5, 0x1b, 0x6d, 0x23, 0x6f, 0xea, 0xc5, 0xe7, 0xc5, 0x01, 0xc1,
	0x08, 0x4e, 0x6b, 0x5a, 0xbf, 0x90, 0xe9, 0xab, 0x83, 0x62, 0x2f, 0xf1, 0x06, 0x7c, 0x4b, 0x46,
	0x96, 0x2a, 0x1a, 0xef, 0x52, 0x83, 0xba, 0xff, 0x81, 0xc9, 0xbf, 0xc7, 0x31, 0x80, 0xb3, 0x0e,
	0xe5, 0x7a, 0x43, 0xec, 0x08, 0x27, 0x70, 0xd1, 0xa9, 0x8f, 0x86, 0x73, 0xb2, 0x96, 0x79, 0x78,
	0xd5, 0x9f, 0x93, 0x6b, 0x97, 0x71, 0x27, 0x5b, 0x62, 0x23, 0xbc, 0xee, 0x4b, 0x32, 0x65, 0xa8,
	0xac, 0xb6, 0x8b, 0x2e, 0x38, 0x46, 0xd6, 0xff, 0x7d, 0x41, 0xad, 0xe4, 0x94, 0x67, 0xec, 0x18,
	0x43, 0x80, 0x8e, 0xbc, 0x96, 0x52, 0x51, 0xc5, 0x4e, 0xe6, 0xf1, 0xf2, 0x4e, 0x48, 0xf7, 0xdd,
	0xac, 0x12, 0xae, 0xd7, 0xa9, 0xfa, 0x7a, 0xa0, 0x4a, 0x50, 0x4a, 0x2d, 0xa5, 0x65, 0x2d, 0x53,
	0xa1, 0xd3, 0xfd, 0x90, 0x2b, 0x7f, 0xb7, 0xe4, 0xd3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5,
	0x59, 0x04, 0xb0, 0x5b, 0x01, 0x00, 0x00,
}
