// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transcoder/v1/transcoder.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v11 "github.com/VideoCoin/cloud-api/profiles/v1"
import v1 "github.com/VideoCoin/cloud-api/workorder/v1"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TranscoderStatus int32

const (
	TranscoderStatusAvailable TranscoderStatus = 0
	TranscoderStatusOffline   TranscoderStatus = 1
	TranscoderStatusError     TranscoderStatus = 2
	TranscoderStatusBusy      TranscoderStatus = 3
)

var TranscoderStatus_name = map[int32]string{
	0: "available",
	1: "offline",
	2: "error",
	3: "busy",
}
var TranscoderStatus_value = map[string]int32{
	"available": 0,
	"offline":   1,
	"error":     2,
	"busy":      3,
}

func (x TranscoderStatus) String() string {
	return proto.EnumName(TranscoderStatus_name, int32(x))
}
func (TranscoderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_transcoder_81849387b72a5fc9, []int{0}
}

type Transcoder struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CpuCores             int32            `protobuf:"varint,2,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	CpuMhz               float64          `protobuf:"fixed64,3,opt,name=cpu_mhz,json=cpuMhz,proto3" json:"cpu_mhz,omitempty"`
	TotalMemory          uint64           `protobuf:"varint,4,opt,name=total_memory,json=totalMemory,proto3" json:"total_memory,omitempty"`
	Status               TranscoderStatus `protobuf:"varint,6,opt,name=status,proto3,enum=cloud.api.transcoder.v1.TranscoderStatus" json:"status,omitempty"`
	Worker               []byte           `protobuf:"bytes,7,opt,name=worker,proto3" json:"worker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Transcoder) Reset()         { *m = Transcoder{} }
func (m *Transcoder) String() string { return proto.CompactTextString(m) }
func (*Transcoder) ProtoMessage()    {}
func (*Transcoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcoder_81849387b72a5fc9, []int{0}
}
func (m *Transcoder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transcoder.Unmarshal(m, b)
}
func (m *Transcoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transcoder.Marshal(b, m, deterministic)
}
func (dst *Transcoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transcoder.Merge(dst, src)
}
func (m *Transcoder) XXX_Size() int {
	return xxx_messageInfo_Transcoder.Size(m)
}
func (m *Transcoder) XXX_DiscardUnknown() {
	xxx_messageInfo_Transcoder.DiscardUnknown(m)
}

var xxx_messageInfo_Transcoder proto.InternalMessageInfo

func (m *Transcoder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transcoder) GetCpuCores() int32 {
	if m != nil {
		return m.CpuCores
	}
	return 0
}

func (m *Transcoder) GetCpuMhz() float64 {
	if m != nil {
		return m.CpuMhz
	}
	return 0
}

func (m *Transcoder) GetTotalMemory() uint64 {
	if m != nil {
		return m.TotalMemory
	}
	return 0
}

func (m *Transcoder) GetStatus() TranscoderStatus {
	if m != nil {
		return m.Status
	}
	return TranscoderStatusAvailable
}

func (m *Transcoder) GetWorker() []byte {
	if m != nil {
		return m.Worker
	}
	return nil
}

type Assignment struct {
	Workorder            *v1.WorkOrder `protobuf:"bytes,1,opt,name=workorder,proto3" json:"workorder,omitempty"`
	Profile              *v11.Profile  `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Assignment) Reset()         { *m = Assignment{} }
func (m *Assignment) String() string { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()    {}
func (*Assignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcoder_81849387b72a5fc9, []int{1}
}
func (m *Assignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assignment.Unmarshal(m, b)
}
func (m *Assignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assignment.Marshal(b, m, deterministic)
}
func (dst *Assignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assignment.Merge(dst, src)
}
func (m *Assignment) XXX_Size() int {
	return xxx_messageInfo_Assignment.Size(m)
}
func (m *Assignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assignment.DiscardUnknown(m)
}

var xxx_messageInfo_Assignment proto.InternalMessageInfo

func (m *Assignment) GetWorkorder() *v1.WorkOrder {
	if m != nil {
		return m.Workorder
	}
	return nil
}

func (m *Assignment) GetProfile() *v11.Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*Transcoder)(nil), "cloud.api.transcoder.v1.Transcoder")
	proto.RegisterType((*Assignment)(nil), "cloud.api.transcoder.v1.Assignment")
	proto.RegisterEnum("cloud.api.transcoder.v1.TranscoderStatus", TranscoderStatus_name, TranscoderStatus_value)
}

func init() {
	proto.RegisterFile("transcoder/v1/transcoder.proto", fileDescriptor_transcoder_81849387b72a5fc9)
}

var fileDescriptor_transcoder_81849387b72a5fc9 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x3d, 0xce, 0x8d, 0x4c, 0xaa, 0xca, 0x1a, 0x15, 0xe2, 0xba, 0xc2, 0x72, 0x23, 0x16,
	0x06, 0x11, 0x5b, 0x09, 0x2b, 0xc4, 0x2a, 0x29, 0xac, 0x50, 0x55, 0x64, 0x10, 0x48, 0x6c, 0x2a,
	0xc7, 0x9e, 0x24, 0xa3, 0x3a, 0x3e, 0xd6, 0xf8, 0x02, 0xe9, 0x13, 0xa0, 0x3c, 0x01, 0x9b, 0x88,
	0x05, 0x3c, 0x45, 0x57, 0x2c, 0x59, 0xa1, 0x3e, 0x02, 0x0a, 0x2b, 0xde, 0x02, 0x79, 0xe2, 0x5c,
	0x48, 0x85, 0xba, 0x3b, 0xff, 0x9c, 0xff, 0x3b, 0x3e, 0x17, 0x63, 0x3d, 0xe1, 0x6e, 0x18, 0x7b,
	0xe0, 0x53, 0x6e, 0x67, 0x1d, 0x7b, 0xa3, 0xac, 0x88, 0x43, 0x02, 0xa4, 0xe9, 0x05, 0x90, 0xfa,
	0x96, 0x1b, 0x31, 0x6b, 0x2b, 0x97, 0x75, 0xb4, 0xf6, 0x88, 0x25, 0xe3, 0x74, 0x60, 0x79, 0x30,
	0xb1, 0x47, 0x30, 0x02, 0x5b, 0xf8, 0x07, 0xe9, 0x50, 0x28, 0x21, 0x44, 0xb4, 0xac, 0xa3, 0x3d,
	0xdb, 0xb2, 0xbf, 0x65, 0x3e, 0x85, 0x13, 0x60, 0xa1, 0x2d, 0x8a, 0xb7, 0xdd, 0x88, 0xd9, 0x1f,
	0x80, 0x5f, 0x00, 0x2f, 0xba, 0x58, 0x8b, 0x02, 0x7e, 0x7a, 0x0b, 0x1c, 0x71, 0x18, 0xb2, 0x80,
	0xc6, 0x39, 0xbb, 0x8a, 0x0b, 0xb4, 0xb7, 0x85, 0xd2, 0x30, 0x83, 0x69, 0xc4, 0xe1, 0xe3, 0x74,
	0xd9, 0xac, 0xd7, 0x1e, 0xd1, 0xb0, 0x9d, 0xb9, 0x01, 0xf3, 0xdd, 0x84, 0xda, 0x37, 0x82, 0x65,
	0x89, 0xd6, 0x35, 0xc2, 0xf8, 0xcd, 0x7a, 0x76, 0xb2, 0x8f, 0x65, 0xe6, 0xab, 0xc8, 0x40, 0x66,
	0xdd, 0x91, 0x99, 0x4f, 0x8e, 0x70, 0xdd, 0x8b, 0xd2, 0x73, 0x0f, 0x38, 0x8d, 0x55, 0xd9, 0x40,
	0x66, 0xc5, 0xb9, 0xe3, 0x45, 0xe9, 0x49, 0xae, 0x49, 0x13, 0xd7, 0xf2, 0xe4, 0x64, 0x7c, 0xa9,
	0x96, 0x0c, 0x64, 0x22, 0xa7, 0xea, 0x45, 0xe9, 0xe9, 0xf8, 0x92, 0x1c, 0xe3, 0xbd, 0x04, 0x12,
	0x37, 0x38, 0x9f, 0xd0, 0x09, 0xf0, 0xa9, 0x5a, 0x36, 0x90, 0x59, 0x76, 0x1a, 0xe2, 0xed, 0x54,
	0x3c, 0x91, 0x1e, 0xae, 0xc6, 0x89, 0x9b, 0xa4, 0xb1, 0x5a, 0x35, 0x90, 0xb9, 0xdf, 0x7d, 0x68,
	0xfd, 0xe7, 0x16, 0xd6, 0xa6, 0xbb, 0xd7, 0x02, 0x70, 0x0a, 0x90, 0xdc, 0xc3, 0xd5, 0x7c, 0x97,
	0x94, 0xab, 0x35, 0x03, 0x99, 0x7b, 0x4e, 0xa1, 0x5a, 0x5f, 0x10, 0xc6, 0xbd, 0x38, 0x66, 0xa3,
	0x70, 0x42, 0xc3, 0x84, 0xbc, 0xc4, 0xf5, 0xf5, 0xca, 0xc5, 0x64, 0x8d, 0xee, 0xf1, 0xd6, 0xc7,
	0x36, 0xe7, 0xc8, 0x3a, 0xd6, 0x3b, 0xe0, 0x17, 0x67, 0xb9, 0xe8, 0xe3, 0xab, 0x3f, 0xdf, 0x4b,
	0x95, 0x19, 0x92, 0x15, 0xe4, 0x6c, 0x78, 0xf2, 0x1c, 0xd7, 0x8a, 0x1b, 0x88, 0x6d, 0x34, 0xba,
	0xfa, 0x56, 0xa9, 0xf5, 0x75, 0xb2, 0x8e, 0xf5, 0x6a, 0x19, 0xff, 0x53, 0x67, 0x85, 0x3e, 0xfa,
	0x89, 0xb0, 0xb2, 0x3b, 0x16, 0x79, 0x8c, 0xeb, 0x6e, 0xe6, 0xb2, 0xc0, 0x1d, 0x04, 0x54, 0x91,
	0xb4, 0xfb, 0xb3, 0xb9, 0x71, 0xb8, 0x6b, 0xea, 0xad, 0x0c, 0xc4, 0xc4, 0x35, 0x18, 0x0e, 0x03,
	0x16, 0x52, 0x05, 0x69, 0x47, 0xb3, 0xb9, 0xd1, 0xdc, 0xf5, 0x9e, 0x2d, 0xd3, 0xe4, 0x01, 0xae,
	0x50, 0xce, 0x81, 0x2b, 0xb2, 0x76, 0x38, 0x9b, 0x1b, 0x77, 0x77, 0x7d, 0x2f, 0xf2, 0x24, 0x69,
	0xe1, 0xf2, 0x20, 0x8d, 0xa7, 0x4a, 0x49, 0x53, 0x67, 0x73, 0xe3, 0x60, 0xd7, 0xd4, 0x4f, 0xe3,
	0xa9, 0xa6, 0x7e, 0xfa, 0xaa, 0x4b, 0x57, 0xdf, 0xf4, 0x1b, 0xbd, 0xf7, 0x0f, 0x7e, 0x2c, 0x74,
	0xe9, 0x7a, 0xa1, 0x4b, 0xbf, 0x16, 0xba, 0xf4, 0xf9, 0xb7, 0x2e, 0xbd, 0x97, 0xb3, 0xce, 0xa0,
	0x2a, 0x7e, 0xb1, 0x27, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xb0, 0xa0, 0xef, 0x87, 0x03,
	0x00, 0x00,
}
