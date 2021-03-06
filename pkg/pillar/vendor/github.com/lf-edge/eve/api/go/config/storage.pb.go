// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package config

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

type DsType int32

const (
	DsType_DsUnknown DsType = 0
	DsType_DsHttp    DsType = 1
	DsType_DsHttps   DsType = 2
	DsType_DsS3      DsType = 3
	DsType_DsSFTP    DsType = 4
)

var DsType_name = map[int32]string{
	0: "DsUnknown",
	1: "DsHttp",
	2: "DsHttps",
	3: "DsS3",
	4: "DsSFTP",
}

var DsType_value = map[string]int32{
	"DsUnknown": 0,
	"DsHttp":    1,
	"DsHttps":   2,
	"DsS3":      3,
	"DsSFTP":    4,
}

func (x DsType) String() string {
	return proto.EnumName(DsType_name, int32(x))
}

func (DsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

type Format int32

const (
	Format_FmtUnknown Format = 0
	Format_RAW        Format = 1
	Format_QCOW       Format = 2
	Format_QCOW2      Format = 3
	Format_VHD        Format = 4
	Format_VMDK       Format = 5
	Format_OVA        Format = 6
	Format_VHDX       Format = 7
)

var Format_name = map[int32]string{
	0: "FmtUnknown",
	1: "RAW",
	2: "QCOW",
	3: "QCOW2",
	4: "VHD",
	5: "VMDK",
	6: "OVA",
	7: "VHDX",
}

var Format_value = map[string]int32{
	"FmtUnknown": 0,
	"RAW":        1,
	"QCOW":       2,
	"QCOW2":      3,
	"VHD":        4,
	"VMDK":       5,
	"OVA":        6,
	"VHDX":       7,
}

func (x Format) String() string {
	return proto.EnumName(Format_name, int32(x))
}

func (Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}

type Target int32

const (
	Target_TgtUnknown Target = 0
	Target_Disk       Target = 1
	Target_Kernel     Target = 2
	Target_Initrd     Target = 3
	Target_RamDisk    Target = 4
)

var Target_name = map[int32]string{
	0: "TgtUnknown",
	1: "Disk",
	2: "Kernel",
	3: "Initrd",
	4: "RamDisk",
}

var Target_value = map[string]int32{
	"TgtUnknown": 0,
	"Disk":       1,
	"Kernel":     2,
	"Initrd":     3,
	"RamDisk":    4,
}

func (x Target) String() string {
	return proto.EnumName(Target_name, int32(x))
}

func (Target) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}

type DriveType int32

const (
	DriveType_Unclassified DriveType = 0
	DriveType_CDROM        DriveType = 1
	DriveType_HDD          DriveType = 2
	DriveType_NET          DriveType = 3
	// this type is allocate the empty disk of maxsizebytes specified
	DriveType_HDD_EMPTY DriveType = 4
)

var DriveType_name = map[int32]string{
	0: "Unclassified",
	1: "CDROM",
	2: "HDD",
	3: "NET",
	4: "HDD_EMPTY",
}

var DriveType_value = map[string]int32{
	"Unclassified": 0,
	"CDROM":        1,
	"HDD":          2,
	"NET":          3,
	"HDD_EMPTY":    4,
}

func (x DriveType) String() string {
	return proto.EnumName(DriveType_name, int32(x))
}

func (DriveType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{3}
}

type SignatureInfo struct {
	Intercertsurl        string   `protobuf:"bytes,1,opt,name=intercertsurl,proto3" json:"intercertsurl,omitempty"`
	Signercerturl        string   `protobuf:"bytes,2,opt,name=signercerturl,proto3" json:"signercerturl,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignatureInfo) Reset()         { *m = SignatureInfo{} }
func (m *SignatureInfo) String() string { return proto.CompactTextString(m) }
func (*SignatureInfo) ProtoMessage()    {}
func (*SignatureInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

func (m *SignatureInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignatureInfo.Unmarshal(m, b)
}
func (m *SignatureInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignatureInfo.Marshal(b, m, deterministic)
}
func (m *SignatureInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureInfo.Merge(m, src)
}
func (m *SignatureInfo) XXX_Size() int {
	return xxx_messageInfo_SignatureInfo.Size(m)
}
func (m *SignatureInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureInfo proto.InternalMessageInfo

func (m *SignatureInfo) GetIntercertsurl() string {
	if m != nil {
		return m.Intercertsurl
	}
	return ""
}

func (m *SignatureInfo) GetSignercerturl() string {
	if m != nil {
		return m.Signercerturl
	}
	return ""
}

func (m *SignatureInfo) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type DatastoreConfig struct {
	Id       string `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
	DType    DsType `protobuf:"varint,1,opt,name=dType,proto3,enum=DsType" json:"dType,omitempty"`
	Fqdn     string `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	ApiKey   string `protobuf:"bytes,3,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// depending on datastore types, it could be bucket or path
	Dpath string `protobuf:"bytes,5,opt,name=dpath,proto3" json:"dpath,omitempty"`
	// Applies for some datastore types
	Region               string   `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatastoreConfig) Reset()         { *m = DatastoreConfig{} }
func (m *DatastoreConfig) String() string { return proto.CompactTextString(m) }
func (*DatastoreConfig) ProtoMessage()    {}
func (*DatastoreConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}

func (m *DatastoreConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatastoreConfig.Unmarshal(m, b)
}
func (m *DatastoreConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatastoreConfig.Marshal(b, m, deterministic)
}
func (m *DatastoreConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatastoreConfig.Merge(m, src)
}
func (m *DatastoreConfig) XXX_Size() int {
	return xxx_messageInfo_DatastoreConfig.Size(m)
}
func (m *DatastoreConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DatastoreConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DatastoreConfig proto.InternalMessageInfo

func (m *DatastoreConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DatastoreConfig) GetDType() DsType {
	if m != nil {
		return m.DType
	}
	return DsType_DsUnknown
}

func (m *DatastoreConfig) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *DatastoreConfig) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *DatastoreConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *DatastoreConfig) GetDpath() string {
	if m != nil {
		return m.Dpath
	}
	return ""
}

func (m *DatastoreConfig) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type Image struct {
	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion,proto3" json:"uuidandversion,omitempty"`
	// it could be relative path/name as well
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Sha256  string `protobuf:"bytes,3,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Iformat Format `protobuf:"varint,4,opt,name=iformat,proto3,enum=Format" json:"iformat,omitempty"`
	// if its signed image
	Siginfo              *SignatureInfo `protobuf:"bytes,5,opt,name=siginfo,proto3" json:"siginfo,omitempty"`
	DsId                 string         `protobuf:"bytes,6,opt,name=dsId,proto3" json:"dsId,omitempty"`
	SizeBytes            int64          `protobuf:"varint,8,opt,name=sizeBytes,proto3" json:"sizeBytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetUuidandversion() *UUIDandVersion {
	if m != nil {
		return m.Uuidandversion
	}
	return nil
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

func (m *Image) GetIformat() Format {
	if m != nil {
		return m.Iformat
	}
	return Format_FmtUnknown
}

func (m *Image) GetSiginfo() *SignatureInfo {
	if m != nil {
		return m.Siginfo
	}
	return nil
}

func (m *Image) GetDsId() string {
	if m != nil {
		return m.DsId
	}
	return ""
}

func (m *Image) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

type Drive struct {
	Image    *Image    `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Readonly bool      `protobuf:"varint,5,opt,name=readonly,proto3" json:"readonly,omitempty"`
	Preserve bool      `protobuf:"varint,6,opt,name=preserve,proto3" json:"preserve,omitempty"`
	Drvtype  DriveType `protobuf:"varint,8,opt,name=drvtype,proto3,enum=DriveType" json:"drvtype,omitempty"`
	Target   Target    `protobuf:"varint,9,opt,name=target,proto3,enum=Target" json:"target,omitempty"`
	// Initial image need to be resized to this size.
	// A value of 0 will indicate that no resizing is required
	Maxsizebytes         int64    `protobuf:"varint,10,opt,name=maxsizebytes,proto3" json:"maxsizebytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Drive) Reset()         { *m = Drive{} }
func (m *Drive) String() string { return proto.CompactTextString(m) }
func (*Drive) ProtoMessage()    {}
func (*Drive) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{3}
}

func (m *Drive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Drive.Unmarshal(m, b)
}
func (m *Drive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Drive.Marshal(b, m, deterministic)
}
func (m *Drive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Drive.Merge(m, src)
}
func (m *Drive) XXX_Size() int {
	return xxx_messageInfo_Drive.Size(m)
}
func (m *Drive) XXX_DiscardUnknown() {
	xxx_messageInfo_Drive.DiscardUnknown(m)
}

var xxx_messageInfo_Drive proto.InternalMessageInfo

func (m *Drive) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Drive) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

func (m *Drive) GetPreserve() bool {
	if m != nil {
		return m.Preserve
	}
	return false
}

func (m *Drive) GetDrvtype() DriveType {
	if m != nil {
		return m.Drvtype
	}
	return DriveType_Unclassified
}

func (m *Drive) GetTarget() Target {
	if m != nil {
		return m.Target
	}
	return Target_TgtUnknown
}

func (m *Drive) GetMaxsizebytes() int64 {
	if m != nil {
		return m.Maxsizebytes
	}
	return 0
}

func init() {
	proto.RegisterEnum("DsType", DsType_name, DsType_value)
	proto.RegisterEnum("Format", Format_name, Format_value)
	proto.RegisterEnum("Target", Target_name, Target_value)
	proto.RegisterEnum("DriveType", DriveType_name, DriveType_value)
	proto.RegisterType((*SignatureInfo)(nil), "SignatureInfo")
	proto.RegisterType((*DatastoreConfig)(nil), "DatastoreConfig")
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*Drive)(nil), "Drive")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 716 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x54, 0xe1, 0x6e, 0xdb, 0x36,
	0x10, 0xae, 0x6c, 0x4b, 0xb6, 0x2f, 0x89, 0x43, 0x10, 0xc3, 0x20, 0x14, 0x1d, 0x9a, 0x19, 0xf9,
	0x11, 0x04, 0x98, 0x0c, 0xb8, 0xd8, 0xf6, 0xbb, 0x8d, 0x9a, 0xda, 0x08, 0xb2, 0x74, 0x8a, 0xed,
	0x6e, 0xc3, 0x80, 0x81, 0x11, 0x69, 0x85, 0xa8, 0x45, 0x7a, 0x24, 0xed, 0xce, 0x7d, 0xb0, 0x3d,
	0xc3, 0xde, 0x65, 0x2f, 0x31, 0x1c, 0x69, 0x3b, 0x73, 0xff, 0xdd, 0xf7, 0xdd, 0xe9, 0xee, 0xfb,
	0xee, 0x24, 0xc1, 0x89, 0x75, 0xda, 0xb0, 0x4a, 0x64, 0x4b, 0xa3, 0x9d, 0x7e, 0x7e, 0xca, 0xc5,
	0xba, 0xd4, 0x75, 0xad, 0x55, 0x20, 0xfa, 0x1b, 0x38, 0xb9, 0x97, 0x95, 0x62, 0x6e, 0x65, 0xc4,
	0x58, 0xcd, 0x35, 0x3d, 0x87, 0x13, 0xa9, 0x9c, 0x30, 0xa5, 0x30, 0xce, 0xae, 0xcc, 0x22, 0x8d,
	0xce, 0xa2, 0x8b, 0x6e, 0x71, 0x48, 0x62, 0x95, 0x95, 0x95, 0x0a, 0x0c, 0x56, 0x35, 0x42, 0xd5,
	0x01, 0x49, 0x5f, 0x40, 0xd7, 0xee, 0x9a, 0xa7, 0xcd, 0xb3, 0xe8, 0xe2, 0xb8, 0x78, 0x22, 0xfa,
	0x7f, 0x47, 0x70, 0x9a, 0x33, 0xc7, 0x50, 0xa1, 0xb8, 0xd2, 0x6a, 0x2e, 0x2b, 0xda, 0x83, 0x86,
	0xe4, 0x29, 0xf7, 0xcd, 0x1a, 0x92, 0xd3, 0x6f, 0x20, 0xe6, 0x93, 0xcd, 0x52, 0x78, 0x15, 0xbd,
	0x61, 0x3b, 0xcb, 0x2d, 0xc2, 0x22, 0xb0, 0x94, 0x42, 0x6b, 0xfe, 0x27, 0x57, 0xdb, 0xe9, 0x3e,
	0xa6, 0x5f, 0x43, 0xc2, 0x96, 0xf2, 0x46, 0x6c, 0xfc, 0xc4, 0x6e, 0xb1, 0x45, 0xf4, 0x39, 0x74,
	0x96, 0xcc, 0xda, 0x4f, 0xda, 0xf0, 0xb4, 0xe5, 0x33, 0x7b, 0x4c, 0xbf, 0x82, 0x98, 0x2f, 0x99,
	0x7b, 0x4c, 0x63, 0x9f, 0x08, 0x00, 0x3b, 0x19, 0x51, 0x49, 0xad, 0xd2, 0x24, 0x74, 0x0a, 0xa8,
	0xff, 0x6f, 0x04, 0xf1, 0xb8, 0x66, 0x95, 0xa0, 0x3f, 0x42, 0x6f, 0xb5, 0x92, 0x9c, 0x29, 0xbe,
	0x16, 0xc6, 0x62, 0x25, 0xea, 0x3c, 0x1a, 0x9e, 0x66, 0xd3, 0xe9, 0x38, 0x67, 0x8a, 0xcf, 0x02,
	0x5d, 0x7c, 0x51, 0x86, 0xc2, 0x15, 0xab, 0xc5, 0x4e, 0x38, 0xc6, 0x38, 0xce, 0x3e, 0xb2, 0xe1,
	0xf7, 0x3f, 0xec, 0x84, 0x07, 0x44, 0xbf, 0x85, 0xb6, 0x9c, 0x6b, 0x53, 0x33, 0xe7, 0x75, 0xe3,
	0x16, 0xae, 0x3d, 0x2c, 0x76, 0x3c, 0xbd, 0x80, 0xb6, 0x95, 0x95, 0x54, 0x73, 0xed, 0x1d, 0x1c,
	0x0d, 0x7b, 0xd9, 0xc1, 0x55, 0x8b, 0x5d, 0x1a, 0x07, 0x73, 0x3b, 0xe6, 0x5b, 0x47, 0x3e, 0x0e,
	0x67, 0xfa, 0x2c, 0xde, 0x6c, 0x9c, 0xb0, 0x69, 0xe7, 0x2c, 0xba, 0x68, 0x16, 0x4f, 0x44, 0xff,
	0x9f, 0x08, 0xe2, 0xdc, 0xc8, 0xb5, 0xa0, 0x2f, 0x20, 0x96, 0x68, 0x7b, 0x6b, 0x32, 0xc9, 0xfc,
	0x12, 0x8a, 0x40, 0xe2, 0x7e, 0x8d, 0x60, 0x5c, 0xab, 0xc5, 0xc6, 0x8b, 0xe8, 0x14, 0x7b, 0xec,
	0x77, 0x6f, 0x84, 0x15, 0x66, 0x2d, 0xfc, 0xe4, 0x4e, 0xb1, 0xc7, 0xf4, 0x1c, 0xda, 0xdc, 0xac,
	0x1d, 0x1e, 0xb9, 0xe3, 0xed, 0x41, 0xe6, 0xc7, 0xf9, 0x3b, 0xef, 0x52, 0xf4, 0x25, 0x24, 0x8e,
	0x99, 0x4a, 0xb8, 0xb4, 0xbb, 0xdd, 0xc1, 0xc4, 0xc3, 0x62, 0x4b, 0xd3, 0x3e, 0x1c, 0xd7, 0xec,
	0x2f, 0x94, 0xfd, 0xe0, 0x7d, 0x80, 0xf7, 0x71, 0xc0, 0x5d, 0x5e, 0x43, 0x12, 0xde, 0x1f, 0x7a,
	0x02, 0xdd, 0xdc, 0x4e, 0xd5, 0x47, 0xa5, 0x3f, 0x29, 0xf2, 0x8c, 0x02, 0x26, 0x46, 0xce, 0x2d,
	0x49, 0x44, 0x8f, 0xa0, 0x1d, 0x62, 0x4b, 0x1a, 0xb4, 0x03, 0xad, 0xdc, 0xde, 0xbf, 0x22, 0xcd,
	0x50, 0x72, 0x7f, 0x3d, 0x79, 0x4f, 0x5a, 0x97, 0xbf, 0x43, 0x12, 0x2e, 0x40, 0x7b, 0x00, 0xd7,
	0xb5, 0x7b, 0x6a, 0xd4, 0x86, 0x66, 0xf1, 0xfa, 0x03, 0x89, 0xf0, 0xc1, 0x9f, 0xaf, 0xee, 0x3e,
	0x90, 0x06, 0xed, 0x42, 0x8c, 0xd1, 0x90, 0x34, 0x31, 0x3b, 0x1b, 0xe5, 0xa4, 0x85, 0xd9, 0xd9,
	0x6d, 0x7e, 0x43, 0x62, 0xa4, 0xee, 0x66, 0xaf, 0x49, 0xe2, 0xa9, 0x51, 0xfe, 0x0b, 0x69, 0x5f,
	0xbe, 0x83, 0x24, 0x78, 0xc3, 0xee, 0x93, 0xea, 0x7f, 0xdd, 0x51, 0x8d, 0xb4, 0x1f, 0x49, 0x84,
	0x6a, 0x6e, 0x84, 0x51, 0x62, 0x41, 0x1a, 0x18, 0x8f, 0x95, 0x74, 0x86, 0x93, 0x26, 0x8a, 0x2f,
	0x58, 0xed, 0x8b, 0x5a, 0x97, 0x63, 0xe8, 0xee, 0x37, 0x49, 0x09, 0x1c, 0x4f, 0x55, 0xb9, 0x60,
	0xd6, 0xca, 0xb9, 0x14, 0x9c, 0x3c, 0x43, 0x61, 0x57, 0x79, 0x71, 0x77, 0x4b, 0x22, 0x54, 0x31,
	0xca, 0x73, 0xd2, 0xc0, 0xe0, 0xa7, 0xb7, 0x13, 0xd2, 0xc4, 0x05, 0x8d, 0xf2, 0xfc, 0x8f, 0xb7,
	0xb7, 0xef, 0x27, 0xbf, 0x92, 0xd6, 0x9b, 0x77, 0xf0, 0xb2, 0xd4, 0x75, 0xf6, 0x59, 0x70, 0xc1,
	0x59, 0x56, 0x2e, 0xf4, 0x8a, 0x67, 0x2b, 0xbc, 0x9f, 0x2c, 0xb7, 0xbf, 0x96, 0xdf, 0xce, 0x2b,
	0xe9, 0x1e, 0x57, 0x0f, 0x59, 0xa9, 0xeb, 0xc1, 0x62, 0xfe, 0x9d, 0xe0, 0x95, 0x18, 0x88, 0xb5,
	0x18, 0xb0, 0xa5, 0x1c, 0x54, 0x7a, 0x50, 0xfa, 0x0f, 0xfc, 0x21, 0xf1, 0xc5, 0xaf, 0xfe, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x4e, 0xe2, 0x52, 0xac, 0x98, 0x04, 0x00, 0x00,
}
