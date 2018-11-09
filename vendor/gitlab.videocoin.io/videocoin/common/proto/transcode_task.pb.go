// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transcode_task.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TranscodeTask struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               TranscodeStatus      `protobuf:"varint,2,opt,name=status,proto3,enum=proto.TranscodeStatus" json:"status,omitempty"`
	CreatedAt            *time.Time           `protobuf:"bytes,3,opt,name=created_at,json=createdAt,stdtime" json:"created_at,omitempty"`
	UpdatedAt            *time.Time           `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at,omitempty"`
	Kind                 string               `protobuf:"bytes,20,opt,name=kind,proto3" json:"kind,omitempty"`
	Metadata             *Metadata            `protobuf:"bytes,21,opt,name=metadata" json:"metadata,omitempty"`
	Input                *TranscodeTaskInput  `protobuf:"bytes,40,opt,name=input" json:"input,omitempty"`
	Output               *TranscodeTaskOutput `protobuf:"bytes,41,opt,name=output" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TranscodeTask) Reset()         { *m = TranscodeTask{} }
func (m *TranscodeTask) String() string { return proto.CompactTextString(m) }
func (*TranscodeTask) ProtoMessage()    {}
func (*TranscodeTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcode_task_70730f148e8e8d7c, []int{0}
}
func (m *TranscodeTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscodeTask.Unmarshal(m, b)
}
func (m *TranscodeTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscodeTask.Marshal(b, m, deterministic)
}
func (dst *TranscodeTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscodeTask.Merge(dst, src)
}
func (m *TranscodeTask) XXX_Size() int {
	return xxx_messageInfo_TranscodeTask.Size(m)
}
func (m *TranscodeTask) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscodeTask.DiscardUnknown(m)
}

var xxx_messageInfo_TranscodeTask proto.InternalMessageInfo

func (m *TranscodeTask) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TranscodeTask) GetStatus() TranscodeStatus {
	if m != nil {
		return m.Status
	}
	return TranscodeStatus_NONE
}

func (m *TranscodeTask) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TranscodeTask) GetUpdatedAt() *time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *TranscodeTask) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *TranscodeTask) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TranscodeTask) GetInput() *TranscodeTaskInput {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *TranscodeTask) GetOutput() *TranscodeTaskOutput {
	if m != nil {
		return m.Output
	}
	return nil
}

type TranscodeTaskInput struct {
	Video                *VideoInput `protobuf:"bytes,1,opt,name=video" json:"video,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TranscodeTaskInput) Reset()         { *m = TranscodeTaskInput{} }
func (m *TranscodeTaskInput) String() string { return proto.CompactTextString(m) }
func (*TranscodeTaskInput) ProtoMessage()    {}
func (*TranscodeTaskInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcode_task_70730f148e8e8d7c, []int{1}
}
func (m *TranscodeTaskInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscodeTaskInput.Unmarshal(m, b)
}
func (m *TranscodeTaskInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscodeTaskInput.Marshal(b, m, deterministic)
}
func (dst *TranscodeTaskInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscodeTaskInput.Merge(dst, src)
}
func (m *TranscodeTaskInput) XXX_Size() int {
	return xxx_messageInfo_TranscodeTaskInput.Size(m)
}
func (m *TranscodeTaskInput) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscodeTaskInput.DiscardUnknown(m)
}

var xxx_messageInfo_TranscodeTaskInput proto.InternalMessageInfo

func (m *TranscodeTaskInput) GetVideo() *VideoInput {
	if m != nil {
		return m.Video
	}
	return nil
}

type TranscodeTaskOutput struct {
	Video                []*VideoOutput `protobuf:"bytes,1,rep,name=video" json:"video,omitempty"`
	Cmdline              string         `protobuf:"bytes,2,opt,name=cmdline,proto3" json:"cmdline,omitempty"`
	Auido                *AudioOutput   `protobuf:"bytes,3,opt,name=auido" json:"auido,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TranscodeTaskOutput) Reset()         { *m = TranscodeTaskOutput{} }
func (m *TranscodeTaskOutput) String() string { return proto.CompactTextString(m) }
func (*TranscodeTaskOutput) ProtoMessage()    {}
func (*TranscodeTaskOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcode_task_70730f148e8e8d7c, []int{2}
}
func (m *TranscodeTaskOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscodeTaskOutput.Unmarshal(m, b)
}
func (m *TranscodeTaskOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscodeTaskOutput.Marshal(b, m, deterministic)
}
func (dst *TranscodeTaskOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscodeTaskOutput.Merge(dst, src)
}
func (m *TranscodeTaskOutput) XXX_Size() int {
	return xxx_messageInfo_TranscodeTaskOutput.Size(m)
}
func (m *TranscodeTaskOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscodeTaskOutput.DiscardUnknown(m)
}

var xxx_messageInfo_TranscodeTaskOutput proto.InternalMessageInfo

func (m *TranscodeTaskOutput) GetVideo() []*VideoOutput {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *TranscodeTaskOutput) GetCmdline() string {
	if m != nil {
		return m.Cmdline
	}
	return ""
}

func (m *TranscodeTaskOutput) GetAuido() *AudioOutput {
	if m != nil {
		return m.Auido
	}
	return nil
}

type SimpleTranscodeTask struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InputUrl             string     `protobuf:"bytes,2,opt,name=input_url,json=inputUrl,proto3" json:"input_url,omitempty"`
	OutputUrl            string     `protobuf:"bytes,3,opt,name=output_url,json=outputUrl,proto3" json:"output_url,omitempty"`
	UserId               int32      `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt            *time.Time `protobuf:"bytes,5,opt,name=created_at,json=createdAt,stdtime" json:"created_at,omitempty"`
	UpdatedAt            *time.Time `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SimpleTranscodeTask) Reset()         { *m = SimpleTranscodeTask{} }
func (m *SimpleTranscodeTask) String() string { return proto.CompactTextString(m) }
func (*SimpleTranscodeTask) ProtoMessage()    {}
func (*SimpleTranscodeTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_transcode_task_70730f148e8e8d7c, []int{3}
}
func (m *SimpleTranscodeTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleTranscodeTask.Unmarshal(m, b)
}
func (m *SimpleTranscodeTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleTranscodeTask.Marshal(b, m, deterministic)
}
func (dst *SimpleTranscodeTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleTranscodeTask.Merge(dst, src)
}
func (m *SimpleTranscodeTask) XXX_Size() int {
	return xxx_messageInfo_SimpleTranscodeTask.Size(m)
}
func (m *SimpleTranscodeTask) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleTranscodeTask.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleTranscodeTask proto.InternalMessageInfo

func (m *SimpleTranscodeTask) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SimpleTranscodeTask) GetInputUrl() string {
	if m != nil {
		return m.InputUrl
	}
	return ""
}

func (m *SimpleTranscodeTask) GetOutputUrl() string {
	if m != nil {
		return m.OutputUrl
	}
	return ""
}

func (m *SimpleTranscodeTask) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SimpleTranscodeTask) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *SimpleTranscodeTask) GetUpdatedAt() *time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*TranscodeTask)(nil), "proto.TranscodeTask")
	proto.RegisterType((*TranscodeTaskInput)(nil), "proto.TranscodeTaskInput")
	proto.RegisterType((*TranscodeTaskOutput)(nil), "proto.TranscodeTaskOutput")
	proto.RegisterType((*SimpleTranscodeTask)(nil), "proto.SimpleTranscodeTask")
}

func init() {
	proto.RegisterFile("transcode_task.proto", fileDescriptor_transcode_task_70730f148e8e8d7c)
}

var fileDescriptor_transcode_task_70730f148e8e8d7c = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x6d, 0xda, 0x26, 0xbb, 0x99, 0x8a, 0x22, 0xbc, 0x0b, 0x1b, 0x8a, 0xe8, 0x56, 0xbd, 0x10,
	0x84, 0x48, 0xa5, 0x70, 0x46, 0xa8, 0x7b, 0xdb, 0x03, 0x42, 0xca, 0x16, 0x0e, 0x5c, 0x2a, 0xb7,
	0x36, 0xc1, 0x6a, 0x52, 0x47, 0x89, 0xcd, 0x0f, 0xf0, 0x03, 0x1c, 0xf9, 0x24, 0x8e, 0xfc, 0x01,
	0xa8, 0x7c, 0x04, 0x17, 0x0e, 0xc8, 0x13, 0xbb, 0x8b, 0x96, 0x15, 0x02, 0x71, 0x8a, 0xdf, 0xcc,
	0x9b, 0xe7, 0x19, 0xcf, 0x0b, 0x1c, 0xab, 0x9a, 0x6e, 0x9b, 0xb5, 0x64, 0x7c, 0xa9, 0x68, 0xb3,
	0x49, 0xaa, 0x5a, 0x2a, 0x49, 0x7c, 0xfc, 0x8c, 0x4e, 0x73, 0x29, 0xf3, 0x82, 0xcf, 0x10, 0xad,
	0xf4, 0x9b, 0x99, 0x12, 0x25, 0x6f, 0x14, 0x2d, 0xab, 0x96, 0x37, 0x7a, 0x9c, 0x0b, 0xf5, 0x56,
	0xaf, 0x92, 0xb5, 0x2c, 0x67, 0xb9, 0xcc, 0xe5, 0x25, 0xd3, 0x20, 0x04, 0x78, 0xb2, 0xf4, 0xc1,
	0x3b, 0xc1, 0xf8, 0x1e, 0x50, 0xcd, 0x84, 0x03, 0xc3, 0x92, 0x2b, 0xca, 0xa8, 0xa2, 0x2d, 0x9e,
	0x7e, 0xef, 0xc2, 0x8d, 0x85, 0xeb, 0x6c, 0x41, 0x9b, 0x0d, 0x19, 0x42, 0x57, 0xb0, 0xc8, 0x9b,
	0x78, 0x71, 0x98, 0x75, 0x05, 0x23, 0x09, 0x04, 0x8d, 0xa2, 0x4a, 0x37, 0x51, 0x77, 0xe2, 0xc5,
	0xc3, 0xf4, 0x4e, 0x5b, 0x99, 0xec, 0xab, 0x2e, 0x30, 0x9b, 0x59, 0x16, 0x79, 0x06, 0xb0, 0xae,
	0x39, 0x55, 0x9c, 0x2d, 0xa9, 0x8a, 0x7a, 0x13, 0x2f, 0x1e, 0xa4, 0xa3, 0xa4, 0x1d, 0x30, 0x71,
	0x6d, 0x27, 0x0b, 0x37, 0xe0, 0x59, 0xff, 0xc3, 0x97, 0x53, 0x2f, 0x0b, 0x6d, 0xcd, 0x5c, 0x19,
	0x01, 0x5d, 0x31, 0x27, 0xd0, 0xff, 0x5b, 0x01, 0x5b, 0x33, 0x57, 0x84, 0x40, 0x7f, 0x23, 0xb6,
	0x2c, 0x3a, 0xc6, 0x19, 0xf0, 0x4c, 0x1e, 0xc1, 0xa1, 0x9b, 0x3c, 0xba, 0x8d, 0x92, 0x37, 0xed,
	0x1c, 0xcf, 0x6d, 0x38, 0xdb, 0x13, 0xc8, 0x0c, 0x7c, 0xb1, 0xad, 0xb4, 0x8a, 0x62, 0x64, 0xde,
	0xbd, 0x3a, 0xb1, 0x79, 0xa7, 0x73, 0x43, 0xc8, 0x5a, 0x1e, 0x49, 0x21, 0x90, 0x5a, 0x99, 0x8a,
	0x87, 0xb6, 0xdd, 0x6b, 0x2a, 0x5e, 0x20, 0x23, 0xb3, 0xcc, 0xe9, 0x53, 0x20, 0xbf, 0x0b, 0x92,
	0x07, 0xe0, 0xe3, 0xee, 0x70, 0x01, 0x83, 0xf4, 0x96, 0x15, 0x7a, 0x65, 0x62, 0xf6, 0x4a, 0xcc,
	0x4f, 0xdf, 0x7b, 0x70, 0x74, 0x8d, 0x3c, 0x89, 0x2f, 0x05, 0x7a, 0xf1, 0x20, 0x25, 0xbf, 0x0a,
	0xd8, 0x0e, 0x5a, 0x02, 0x89, 0xe0, 0x60, 0x5d, 0xb2, 0x42, 0x6c, 0x39, 0x6e, 0x36, 0xcc, 0x1c,
	0x34, 0x1a, 0x54, 0x0b, 0x26, 0xed, 0xf6, 0x9c, 0xc6, 0xdc, 0xf8, 0xc8, 0x69, 0x20, 0x61, 0xfa,
	0xc3, 0x83, 0xa3, 0x0b, 0x51, 0x56, 0x05, 0xff, 0xb3, 0x89, 0xee, 0x41, 0x88, 0x2f, 0xb5, 0xd4,
	0x75, 0x61, 0x6f, 0x3b, 0xc4, 0xc0, 0xcb, 0xba, 0x20, 0xf7, 0x01, 0xda, 0x37, 0xc1, 0x6c, 0x0f,
	0xb3, 0x61, 0x1b, 0x31, 0xe9, 0x13, 0x38, 0xd0, 0x0d, 0xaf, 0x97, 0x82, 0xa1, 0x19, 0xfc, 0x2c,
	0x30, 0xf0, 0x9c, 0x5d, 0x71, 0x9a, 0xff, 0xbf, 0x4e, 0x0b, 0xfe, 0xd9, 0x69, 0x67, 0x27, 0x9f,
	0x76, 0xe3, 0xce, 0xe7, 0xdd, 0xb8, 0xf3, 0x75, 0x37, 0xee, 0x7c, 0xfc, 0x36, 0xee, 0xbc, 0x6e,
	0x7f, 0xe8, 0x55, 0x80, 0x9f, 0x27, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x74, 0x03, 0x6f, 0xb6,
	0xf6, 0x03, 0x00, 0x00,
}
