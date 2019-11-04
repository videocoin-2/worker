// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dispatcher/v1/task.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TaskStatus int32

const (
	TaskStatusCreated   TaskStatus = 0
	TaskStatusPending   TaskStatus = 1
	TaskStatusAssigned  TaskStatus = 2
	TaskStatusEncoding  TaskStatus = 3
	TaskStatusCompleted TaskStatus = 4
	TaskStatusFailed    TaskStatus = 5
	TaskStatusCanceled  TaskStatus = 6
)

var TaskStatus_name = map[int32]string{
	0: "CREATED",
	1: "PENDING",
	2: "ASSIGNED",
	3: "ENCODING",
	4: "COMPLETED",
	5: "FAILED",
	6: "CANCELED",
}

var TaskStatus_value = map[string]int32{
	"CREATED":   0,
	"PENDING":   1,
	"ASSIGNED":  2,
	"ENCODING":  3,
	"COMPLETED": 4,
	"FAILED":    5,
	"CANCELED":  6,
}

func (x TaskStatus) String() string {
	return proto.EnumName(TaskStatus_name, int32(x))
}

func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9fa62cfa120fc1c4, []int{0}
}

type TaskInput struct {
	URI                  string   `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskInput) Reset()         { *m = TaskInput{} }
func (m *TaskInput) String() string { return proto.CompactTextString(m) }
func (*TaskInput) ProtoMessage()    {}
func (*TaskInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa62cfa120fc1c4, []int{0}
}
func (m *TaskInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInput.Merge(m, src)
}
func (m *TaskInput) XXX_Size() int {
	return m.Size()
}
func (m *TaskInput) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskInput.DiscardUnknown(m)
}

var xxx_messageInfo_TaskInput proto.InternalMessageInfo

func (m *TaskInput) GetURI() string {
	if m != nil {
		return m.URI
	}
	return ""
}

func (*TaskInput) XXX_MessageName() string {
	return "cloud.api.dispatcher.v1.TaskInput"
}

type TaskOutput struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskOutput) Reset()         { *m = TaskOutput{} }
func (m *TaskOutput) String() string { return proto.CompactTextString(m) }
func (*TaskOutput) ProtoMessage()    {}
func (*TaskOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa62cfa120fc1c4, []int{1}
}
func (m *TaskOutput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskOutput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskOutput.Merge(m, src)
}
func (m *TaskOutput) XXX_Size() int {
	return m.Size()
}
func (m *TaskOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskOutput.DiscardUnknown(m)
}

var xxx_messageInfo_TaskOutput proto.InternalMessageInfo

func (m *TaskOutput) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (*TaskOutput) XXX_MessageName() string {
	return "cloud.api.dispatcher.v1.TaskOutput"
}

type Task struct {
	ID                    string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerID               int32       `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	CreatedAt             *time.Time  `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at,omitempty" db:"created_at"`
	Status                TaskStatus  `protobuf:"varint,4,opt,name=status,proto3,enum=cloud.api.dispatcher.v1.TaskStatus" json:"status,omitempty"`
	ProfileID             string      `protobuf:"bytes,5,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	Input                 *TaskInput  `protobuf:"bytes,6,opt,name=input,proto3" json:"input,omitempty"`
	Output                *TaskOutput `protobuf:"bytes,7,opt,name=output,proto3" json:"output,omitempty"`
	Cmdline               string      `protobuf:"bytes,8,opt,name=cmdline,proto3" json:"cmdline,omitempty"`
	ClientID              string      `protobuf:"bytes,9,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	StreamContractID      uint64      `protobuf:"varint,10,opt,name=stream_contract_id,json=streamContractId,proto3" json:"stream_contract_id,omitempty"`
	StreamContractAddress string      `protobuf:"bytes,11,opt,name=stream_contract_address,json=streamContractAddress,proto3" json:"stream_contract_address,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}    `json:"-"`
	XXX_unrecognized      []byte      `json:"-"`
	XXX_sizecache         int32       `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa62cfa120fc1c4, []int{2}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Task.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return m.Size()
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Task) GetOwnerID() int32 {
	if m != nil {
		return m.OwnerID
	}
	return 0
}

func (m *Task) GetCreatedAt() *time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Task) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatusCreated
}

func (m *Task) GetProfileID() string {
	if m != nil {
		return m.ProfileID
	}
	return ""
}

func (m *Task) GetInput() *TaskInput {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *Task) GetOutput() *TaskOutput {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *Task) GetCmdline() string {
	if m != nil {
		return m.Cmdline
	}
	return ""
}

func (m *Task) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *Task) GetStreamContractID() uint64 {
	if m != nil {
		return m.StreamContractID
	}
	return 0
}

func (m *Task) GetStreamContractAddress() string {
	if m != nil {
		return m.StreamContractAddress
	}
	return ""
}

func (*Task) XXX_MessageName() string {
	return "cloud.api.dispatcher.v1.Task"
}
func init() {
	proto.RegisterEnum("cloud.api.dispatcher.v1.TaskStatus", TaskStatus_name, TaskStatus_value)
	golang_proto.RegisterEnum("cloud.api.dispatcher.v1.TaskStatus", TaskStatus_name, TaskStatus_value)
	proto.RegisterType((*TaskInput)(nil), "cloud.api.dispatcher.v1.TaskInput")
	golang_proto.RegisterType((*TaskInput)(nil), "cloud.api.dispatcher.v1.TaskInput")
	proto.RegisterType((*TaskOutput)(nil), "cloud.api.dispatcher.v1.TaskOutput")
	golang_proto.RegisterType((*TaskOutput)(nil), "cloud.api.dispatcher.v1.TaskOutput")
	proto.RegisterType((*Task)(nil), "cloud.api.dispatcher.v1.Task")
	golang_proto.RegisterType((*Task)(nil), "cloud.api.dispatcher.v1.Task")
}

func init() { proto.RegisterFile("dispatcher/v1/task.proto", fileDescriptor_9fa62cfa120fc1c4) }
func init() { golang_proto.RegisterFile("dispatcher/v1/task.proto", fileDescriptor_9fa62cfa120fc1c4) }

var fileDescriptor_9fa62cfa120fc1c4 = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x6e, 0xd3, 0x4a,
	0x14, 0xc6, 0xeb, 0xfc, 0xf7, 0xf4, 0xfe, 0xc9, 0x9d, 0xdb, 0x36, 0xbe, 0xd6, 0x55, 0x6c, 0x05,
	0x54, 0x05, 0x04, 0x8e, 0x5a, 0x24, 0x84, 0x60, 0x95, 0xd8, 0x6e, 0x65, 0xa9, 0x24, 0x91, 0x53,
	0x36, 0x6c, 0xaa, 0x89, 0x67, 0x9a, 0x8e, 0xea, 0x78, 0x2c, 0x7b, 0x52, 0x5e, 0x01, 0x65, 0xc5,
	0x0b, 0x64, 0x45, 0x9f, 0x02, 0x09, 0x89, 0x65, 0x97, 0x3c, 0x41, 0x40, 0xee, 0x1b, 0xf0, 0x04,
	0xc8, 0x63, 0x97, 0xb4, 0x91, 0xe8, 0xce, 0xf3, 0x9d, 0xdf, 0x77, 0xe6, 0x3b, 0x67, 0x12, 0xa0,
	0x60, 0x1a, 0x87, 0x88, 0x7b, 0x67, 0x24, 0xea, 0x5c, 0xec, 0x75, 0x38, 0x8a, 0xcf, 0x8d, 0x30,
	0x62, 0x9c, 0xc1, 0x86, 0xe7, 0xb3, 0x19, 0x36, 0x50, 0x48, 0x8d, 0x15, 0x63, 0x5c, 0xec, 0xa9,
	0xff, 0x4f, 0x18, 0x9b, 0xf8, 0xa4, 0x83, 0x42, 0xda, 0x41, 0x41, 0xc0, 0x38, 0xe2, 0x94, 0x05,
	0x71, 0x66, 0x53, 0x9f, 0x4e, 0x28, 0x3f, 0x9b, 0x8d, 0x0d, 0x8f, 0x4d, 0x3b, 0x13, 0x36, 0x61,
	0x1d, 0x21, 0x8f, 0x67, 0xa7, 0xe2, 0x24, 0x0e, 0xe2, 0x2b, 0xc7, 0xb5, 0xbc, 0xd9, 0x2f, 0x8a,
	0xd3, 0x29, 0x89, 0x39, 0x9a, 0x86, 0x19, 0xd0, 0xda, 0x05, 0xf2, 0x31, 0x8a, 0xcf, 0x9d, 0x20,
	0x9c, 0x71, 0xf8, 0x1f, 0x28, 0xce, 0x22, 0xaa, 0x48, 0xba, 0xd4, 0x96, 0x7b, 0xd5, 0x64, 0xa9,
	0x15, 0xdf, 0xb8, 0x8e, 0x9b, 0x6a, 0x2d, 0x1d, 0x80, 0x94, 0x1b, 0xcc, 0x78, 0x0a, 0x42, 0x50,
	0x0a, 0x11, 0x3f, 0xcb, 0x48, 0x57, 0x7c, 0xb7, 0x3e, 0x97, 0x40, 0x29, 0x45, 0xe0, 0x0e, 0x28,
	0x50, 0x9c, 0x37, 0xa9, 0x24, 0x4b, 0xad, 0xe0, 0x58, 0x6e, 0x81, 0x62, 0xb8, 0x0b, 0x6a, 0xec,
	0x5d, 0x40, 0xa2, 0x13, 0x8a, 0x95, 0x82, 0x2e, 0xb5, 0xcb, 0xbd, 0xcd, 0x64, 0xa9, 0x55, 0x07,
	0xa9, 0xe6, 0x58, 0x6e, 0x55, 0x14, 0x1d, 0x0c, 0x5d, 0x00, 0xbc, 0x88, 0x20, 0x4e, 0xf0, 0x09,
	0xe2, 0x4a, 0x51, 0x97, 0xda, 0x9b, 0xfb, 0xaa, 0x91, 0x0d, 0x62, 0xdc, 0x0c, 0x62, 0x1c, 0xdf,
	0x0c, 0xd2, 0x6b, 0xfc, 0x58, 0x6a, 0x7f, 0xe3, 0xf1, 0xcb, 0xd6, 0xca, 0xd5, 0xfa, 0xf0, 0x4d,
	0x93, 0x5c, 0x39, 0x17, 0xba, 0x1c, 0xbe, 0x02, 0x95, 0x98, 0x23, 0x3e, 0x8b, 0x95, 0x92, 0x2e,
	0xb5, 0xff, 0xda, 0x7f, 0x60, 0xfc, 0x66, 0xfd, 0x46, 0x3a, 0xc2, 0x48, 0xa0, 0x6e, 0x6e, 0x81,
	0x4f, 0x00, 0x08, 0x23, 0x76, 0x4a, 0x7d, 0x92, 0x46, 0x2f, 0x8b, 0xc1, 0xfe, 0x4c, 0x96, 0x9a,
	0x3c, 0xcc, 0x54, 0xc7, 0x72, 0xe5, 0x1c, 0x70, 0x30, 0x7c, 0x01, 0xca, 0x34, 0xdd, 0xa6, 0x52,
	0x11, 0xc9, 0x5b, 0xf7, 0xde, 0x24, 0xf6, 0xee, 0x66, 0x86, 0x34, 0x24, 0x13, 0xfb, 0x55, 0xaa,
	0xc2, 0x7a, 0x7f, 0xc8, 0xec, 0x29, 0xdc, 0xdc, 0x02, 0x15, 0x50, 0xf5, 0xa6, 0xd8, 0xa7, 0x01,
	0x51, 0x6a, 0xe2, 0x55, 0x6e, 0x8e, 0xf0, 0x11, 0x90, 0x3d, 0x9f, 0x92, 0x80, 0xa7, 0xe9, 0x65,
	0x91, 0xfe, 0x8f, 0x64, 0xa9, 0xd5, 0x4c, 0x21, 0x3a, 0x96, 0x5b, 0xcb, 0xca, 0x0e, 0x86, 0x3d,
	0x00, 0x63, 0x1e, 0x11, 0x34, 0x3d, 0xf1, 0x58, 0xc0, 0x23, 0xe4, 0x09, 0x0f, 0xd0, 0xa5, 0x76,
	0xa9, 0xb7, 0x95, 0x2c, 0xb5, 0xfa, 0x48, 0x54, 0xcd, 0xbc, 0xe8, 0x58, 0x6e, 0x3d, 0xbe, 0xab,
	0x60, 0xf8, 0x1c, 0x34, 0xd6, 0x7b, 0x20, 0x8c, 0x23, 0x12, 0xc7, 0xca, 0xa6, 0x08, 0xb6, 0x7d,
	0xd7, 0xd2, 0xcd, 0x8a, 0x8f, 0x2f, 0x0b, 0xd9, 0x4f, 0x2c, 0x5b, 0x3e, 0x6c, 0x81, 0xaa, 0xe9,
	0xda, 0xdd, 0x63, 0xdb, 0xaa, 0x6f, 0xa8, 0xdb, 0xf3, 0x85, 0xfe, 0xcf, 0xaa, 0x68, 0x66, 0xef,
	0x9a, 0x32, 0x43, 0xbb, 0x6f, 0x39, 0xfd, 0xc3, 0xba, 0xb4, 0xce, 0x0c, 0x49, 0x80, 0x69, 0x30,
	0x81, 0x0f, 0x41, 0xad, 0x3b, 0x1a, 0x39, 0x87, 0x7d, 0xdb, 0xaa, 0x17, 0xd4, 0x9d, 0xf9, 0x42,
	0x87, 0x2b, 0xa8, 0x1b, 0xc7, 0x74, 0x12, 0x10, 0x9c, 0x52, 0x76, 0xdf, 0x1c, 0x88, 0x56, 0xc5,
	0x75, 0xca, 0x0e, 0x3c, 0x26, 0x7a, 0xed, 0x02, 0xd9, 0x1c, 0xbc, 0x1e, 0x1e, 0xd9, 0x69, 0xaa,
	0x92, 0xda, 0x98, 0x2f, 0xf4, 0x7f, 0x6f, 0xa5, 0x62, 0xd3, 0xd0, 0x27, 0x69, 0x2e, 0x1d, 0x54,
	0x0e, 0xba, 0xce, 0x91, 0x6d, 0xd5, 0xcb, 0xea, 0xd6, 0x7c, 0xa1, 0xd7, 0x57, 0xd0, 0x01, 0xa2,
	0x7e, 0x76, 0x9f, 0xd9, 0xed, 0x9b, 0x76, 0xca, 0x54, 0xd6, 0xef, 0x33, 0x51, 0xe0, 0x11, 0x9f,
	0x60, 0x15, 0xbe, 0xff, 0xd8, 0xdc, 0xf8, 0x74, 0xd9, 0xbc, 0xb5, 0x97, 0x9e, 0x72, 0x95, 0x34,
	0xa5, 0xaf, 0x49, 0x53, 0xfa, 0x9e, 0x34, 0xa5, 0x2f, 0xd7, 0x4d, 0xe9, 0xea, 0xba, 0x29, 0xbd,
	0x2d, 0x5c, 0xec, 0x8d, 0x2b, 0xe2, 0xbf, 0xf1, 0xec, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28,
	0x6d, 0xa4, 0x38, 0x74, 0x04, 0x00, 0x00,
}

func (m *TaskInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskInput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskInput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.URI) > 0 {
		i -= len(m.URI)
		copy(dAtA[i:], m.URI)
		i = encodeVarintTask(dAtA, i, uint64(len(m.URI)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TaskOutput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskOutput) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskOutput) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Task) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Task) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Task) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.StreamContractAddress) > 0 {
		i -= len(m.StreamContractAddress)
		copy(dAtA[i:], m.StreamContractAddress)
		i = encodeVarintTask(dAtA, i, uint64(len(m.StreamContractAddress)))
		i--
		dAtA[i] = 0x5a
	}
	if m.StreamContractID != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.StreamContractID))
		i--
		dAtA[i] = 0x50
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintTask(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Cmdline) > 0 {
		i -= len(m.Cmdline)
		copy(dAtA[i:], m.Cmdline)
		i = encodeVarintTask(dAtA, i, uint64(len(m.Cmdline)))
		i--
		dAtA[i] = 0x42
	}
	if m.Output != nil {
		{
			size, err := m.Output.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTask(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Input != nil {
		{
			size, err := m.Input.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTask(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.ProfileID) > 0 {
		i -= len(m.ProfileID)
		copy(dAtA[i:], m.ProfileID)
		i = encodeVarintTask(dAtA, i, uint64(len(m.ProfileID)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Status != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if m.CreatedAt != nil {
		n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintTask(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x1a
	}
	if m.OwnerID != 0 {
		i = encodeVarintTask(dAtA, i, uint64(m.OwnerID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTask(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTask(dAtA []byte, offset int, v uint64) int {
	offset -= sovTask(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TaskInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.URI)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TaskOutput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Task) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.OwnerID != 0 {
		n += 1 + sovTask(uint64(m.OwnerID))
	}
	if m.CreatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.CreatedAt)
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovTask(uint64(m.Status))
	}
	l = len(m.ProfileID)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Input != nil {
		l = m.Input.Size()
		n += 1 + l + sovTask(uint64(l))
	}
	if m.Output != nil {
		l = m.Output.Size()
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.Cmdline)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.StreamContractID != 0 {
		n += 1 + sovTask(uint64(m.StreamContractID))
	}
	l = len(m.StreamContractAddress)
	if l > 0 {
		n += 1 + l + sovTask(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTask(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTask(x uint64) (n int) {
	return sovTask(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TaskInput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaskInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TaskOutput) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaskOutput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskOutput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Task) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTask
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Task: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Task: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerID", wireType)
			}
			m.OwnerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OwnerID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatedAt == nil {
				m.CreatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= TaskStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Input", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Input == nil {
				m.Input = &TaskInput{}
			}
			if err := m.Input.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Output", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Output == nil {
				m.Output = &TaskOutput{}
			}
			if err := m.Output.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cmdline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cmdline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamContractID", wireType)
			}
			m.StreamContractID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreamContractID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTask
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTask
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTask
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StreamContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTask(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTask
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTask(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTask
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTask
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTask
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTask
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTask
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTask        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTask          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTask = fmt.Errorf("proto: unexpected end of group")
)