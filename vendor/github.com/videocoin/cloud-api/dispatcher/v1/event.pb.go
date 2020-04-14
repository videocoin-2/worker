// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dispatcher/v1/event.proto

package v1

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EventType int32

const (
	EventTypeUnknown            EventType = 0
	EventTypeCreate             EventType = 1
	EventTypeUpdate             EventType = 2
	EventTypeDelete             EventType = 3
	EventTypeUpdateStatus       EventType = 4
	EventTypeTaskCompleted      EventType = 5
	EventTypeSegementTranscoded EventType = 6
)

var EventType_name = map[int32]string{
	0: "EVENT_TYPE_UNKNOWN",
	1: "EVENT_TYPE_CREATE",
	2: "EVENT_TYPE_UPDATE",
	3: "EVENT_TYPE_DELETE",
	4: "EVENT_TYPE_UPDATE_STATUS",
	5: "EVENT_TYPE_TASK_COMPLETED",
	6: "EVENT_TYPE_SEGEMENT_TRANSCODED",
}

var EventType_value = map[string]int32{
	"EVENT_TYPE_UNKNOWN":             0,
	"EVENT_TYPE_CREATE":              1,
	"EVENT_TYPE_UPDATE":              2,
	"EVENT_TYPE_DELETE":              3,
	"EVENT_TYPE_UPDATE_STATUS":       4,
	"EVENT_TYPE_TASK_COMPLETED":      5,
	"EVENT_TYPE_SEGEMENT_TRANSCODED": 6,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dadd60d0273236e0, []int{0}
}

type Event struct {
	Type                  EventType  `protobuf:"varint,1,opt,name=type,proto3,enum=cloud.api.dispatcher.v1.EventType" json:"type,omitempty"`
	TaskID                string     `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Status                TaskStatus `protobuf:"varint,3,opt,name=status,proto3,enum=cloud.api.dispatcher.v1.TaskStatus" json:"status,omitempty"`
	StreamID              string     `protobuf:"bytes,4,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	StreamName            string     `protobuf:"bytes,5,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	StreamContractAddress string     `protobuf:"bytes,6,opt,name=stream_contract_address,json=streamContractAddress,proto3" json:"stream_contract_address,omitempty"`
	StreamIsLive          bool       `protobuf:"varint,7,opt,name=stream_is_live,json=streamIsLive,proto3" json:"stream_is_live,omitempty"`
	ProfileID             string     `protobuf:"bytes,8,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	ProfileName           string     `protobuf:"bytes,9,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	ProfileCost           float64    `protobuf:"fixed64,10,opt,name=profile_cost,json=profileCost,proto3" json:"profile_cost,omitempty"`
	ClientID              string     `protobuf:"bytes,11,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientUserID          string     `protobuf:"bytes,12,opt,name=client_user_id,json=clientUserId,proto3" json:"client_user_id,omitempty"`
	UserID                string     `protobuf:"bytes,13,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChunkNum              uint64     `protobuf:"varint,14,opt,name=chunk_num,json=chunkNum,proto3" json:"chunk_num,omitempty"`
	Duration              float64    `protobuf:"fixed64,15,opt,name=duration,proto3" json:"duration,omitempty"`
	CostPerSec            float64    `protobuf:"fixed64,16,opt,name=cost_per_sec,json=costPerSec,proto3" json:"cost_per_sec,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}   `json:"-"`
	XXX_unrecognized      []byte     `json:"-"`
	XXX_sizecache         int32      `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_dadd60d0273236e0, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventTypeUnknown
}

func (m *Event) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *Event) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatusCreated
}

func (m *Event) GetStreamID() string {
	if m != nil {
		return m.StreamID
	}
	return ""
}

func (m *Event) GetStreamName() string {
	if m != nil {
		return m.StreamName
	}
	return ""
}

func (m *Event) GetStreamContractAddress() string {
	if m != nil {
		return m.StreamContractAddress
	}
	return ""
}

func (m *Event) GetStreamIsLive() bool {
	if m != nil {
		return m.StreamIsLive
	}
	return false
}

func (m *Event) GetProfileID() string {
	if m != nil {
		return m.ProfileID
	}
	return ""
}

func (m *Event) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

func (m *Event) GetProfileCost() float64 {
	if m != nil {
		return m.ProfileCost
	}
	return 0
}

func (m *Event) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *Event) GetClientUserID() string {
	if m != nil {
		return m.ClientUserID
	}
	return ""
}

func (m *Event) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Event) GetChunkNum() uint64 {
	if m != nil {
		return m.ChunkNum
	}
	return 0
}

func (m *Event) GetDuration() float64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Event) GetCostPerSec() float64 {
	if m != nil {
		return m.CostPerSec
	}
	return 0
}

func (*Event) XXX_MessageName() string {
	return "cloud.api.dispatcher.v1.Event"
}
func init() {
	proto.RegisterEnum("cloud.api.dispatcher.v1.EventType", EventType_name, EventType_value)
	golang_proto.RegisterEnum("cloud.api.dispatcher.v1.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Event)(nil), "cloud.api.dispatcher.v1.Event")
	golang_proto.RegisterType((*Event)(nil), "cloud.api.dispatcher.v1.Event")
}

func init() { proto.RegisterFile("dispatcher/v1/event.proto", fileDescriptor_dadd60d0273236e0) }
func init() { golang_proto.RegisterFile("dispatcher/v1/event.proto", fileDescriptor_dadd60d0273236e0) }

var fileDescriptor_dadd60d0273236e0 = []byte{
	// 745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xcf, 0x6e, 0xeb, 0x44,
	0x14, 0xc6, 0xaf, 0xdb, 0x24, 0x37, 0x99, 0xe6, 0xa6, 0xee, 0x40, 0xa9, 0xeb, 0x22, 0xc7, 0xb4,
	0x2c, 0x42, 0x55, 0x12, 0x15, 0xa4, 0x22, 0xc4, 0x2a, 0xb5, 0x2d, 0x64, 0xb5, 0x4d, 0x23, 0xdb,
	0x01, 0xc1, 0xc6, 0x72, 0xed, 0x69, 0x6a, 0x25, 0xf6, 0x58, 0x9e, 0x71, 0x50, 0xdf, 0x00, 0xf2,
	0x0e, 0xd9, 0x00, 0x4f, 0xc1, 0x8a, 0x65, 0x97, 0x3c, 0x41, 0x84, 0xd2, 0x17, 0x41, 0x33, 0x76,
	0x93, 0x90, 0xaa, 0x3b, 0x9f, 0xf3, 0xfd, 0xbe, 0x73, 0x3e, 0x8f, 0xff, 0x80, 0xc3, 0x20, 0x24,
	0x89, 0x47, 0xfd, 0x07, 0x94, 0x76, 0x26, 0xe7, 0x1d, 0x34, 0x41, 0x31, 0x6d, 0x27, 0x29, 0xa6,
	0x18, 0x1e, 0xf8, 0x63, 0x9c, 0x05, 0x6d, 0x2f, 0x09, 0xdb, 0x2b, 0xa8, 0x3d, 0x39, 0x97, 0x3f,
	0x1d, 0x62, 0x3c, 0x1c, 0xa3, 0x8e, 0x97, 0x84, 0x1d, 0x2f, 0x8e, 0x31, 0xf5, 0x68, 0x88, 0x63,
	0x92, 0xdb, 0xe4, 0x2f, 0x87, 0x21, 0x7d, 0xc8, 0xee, 0xda, 0x3e, 0x8e, 0x3a, 0x43, 0x3c, 0xc4,
	0x1d, 0xde, 0xbe, 0xcb, 0xee, 0x79, 0xc5, 0x0b, 0x7e, 0x55, 0xe0, 0xd2, 0xff, 0x03, 0x50, 0x8f,
	0x8c, 0x72, 0xe5, 0xf8, 0xf7, 0x32, 0x28, 0x1b, 0x2c, 0x0f, 0xbc, 0x00, 0x25, 0xfa, 0x98, 0x20,
	0x49, 0x50, 0x85, 0x56, 0xe3, 0xab, 0xe3, 0xf6, 0x1b, 0xc1, 0xda, 0x9c, 0x76, 0x1e, 0x13, 0x64,
	0x71, 0x1e, 0x9e, 0x80, 0xf7, 0x6c, 0x9e, 0x1b, 0x06, 0xd2, 0x96, 0x2a, 0xb4, 0x6a, 0x97, 0x60,
	0x31, 0x6f, 0x56, 0x1c, 0x8f, 0x8c, 0x4c, 0xdd, 0xaa, 0x30, 0xc9, 0x0c, 0xe0, 0x77, 0xa0, 0x42,
	0xa8, 0x47, 0x33, 0x22, 0x6d, 0xf3, 0xf1, 0x27, 0x6f, 0x8e, 0x67, 0x46, 0x9b, 0xa3, 0x56, 0x61,
	0x81, 0x5f, 0x80, 0x1a, 0xa1, 0x29, 0xf2, 0x22, 0xb6, 0xa3, 0xc4, 0x77, 0xd4, 0x17, 0xf3, 0x66,
	0xd5, 0xe6, 0x4d, 0x53, 0xb7, 0xaa, 0xb9, 0x6c, 0x06, 0xb0, 0x09, 0x76, 0x0a, 0x34, 0xf6, 0x22,
	0x24, 0x95, 0x19, 0x6c, 0x81, 0xbc, 0xd5, 0xf3, 0x22, 0x04, 0x2f, 0xc0, 0x41, 0x01, 0xf8, 0x38,
	0xa6, 0xa9, 0xe7, 0x53, 0xd7, 0x0b, 0x82, 0x14, 0x11, 0x22, 0x55, 0x38, 0xbc, 0x9f, 0xcb, 0x5a,
	0xa1, 0x76, 0x73, 0x11, 0x7e, 0x0e, 0x1a, 0x2f, 0x19, 0x88, 0x3b, 0x0e, 0x27, 0x48, 0x7a, 0xaf,
	0x0a, 0xad, 0xaa, 0x55, 0x2f, 0x56, 0x93, 0xeb, 0x70, 0x82, 0xe0, 0x19, 0x00, 0x49, 0x8a, 0xef,
	0xc3, 0x31, 0x62, 0x51, 0xab, 0x3c, 0xea, 0x87, 0xc5, 0xbc, 0x59, 0xeb, 0xe7, 0x5d, 0x53, 0xb7,
	0x6a, 0x05, 0x60, 0x06, 0xf0, 0x33, 0x50, 0x7f, 0xa1, 0x79, 0xda, 0x1a, 0x0f, 0xb0, 0x53, 0xf4,
	0x78, 0xdc, 0x35, 0xc4, 0xc7, 0x84, 0x4a, 0x40, 0x15, 0x5a, 0xc2, 0x12, 0xd1, 0x30, 0xa1, 0xec,
	0x74, 0xfc, 0x71, 0x88, 0x62, 0xca, 0x56, 0xee, 0xac, 0x4e, 0x47, 0xe3, 0x4d, 0x76, 0x3a, 0xb9,
	0x6c, 0x06, 0xf0, 0x02, 0x34, 0x0a, 0x34, 0x23, 0x28, 0x65, 0x7c, 0x9d, 0xf3, 0xe2, 0x62, 0xde,
	0xac, 0xe7, 0xfc, 0x80, 0xa0, 0xd4, 0xd4, 0xad, 0xba, 0xbf, 0xaa, 0x02, 0xf6, 0x88, 0x5f, 0x0c,
	0x1f, 0x56, 0x8f, 0xb8, 0x40, 0x2b, 0x59, 0x0e, 0x1d, 0x81, 0x9a, 0xff, 0x90, 0xc5, 0x23, 0x37,
	0xce, 0x22, 0xa9, 0xa1, 0x0a, 0xad, 0x92, 0x55, 0xe5, 0x8d, 0x5e, 0x16, 0x41, 0x19, 0x54, 0x83,
	0x2c, 0xe5, 0xaf, 0xb0, 0xb4, 0xcb, 0xef, 0x61, 0x59, 0x43, 0x15, 0xd4, 0xd9, 0xbd, 0xb9, 0x09,
	0x4a, 0x5d, 0x82, 0x7c, 0x49, 0xe4, 0x3a, 0x60, 0xbd, 0x3e, 0x4a, 0x6d, 0xe4, 0x9f, 0xfe, 0xb6,
	0x0d, 0x6a, 0xcb, 0xd7, 0x0e, 0x9e, 0x01, 0x68, 0xfc, 0x60, 0xf4, 0x1c, 0xd7, 0xf9, 0xa9, 0x6f,
	0xb8, 0x83, 0xde, 0x55, 0xef, 0xf6, 0xc7, 0x9e, 0xf8, 0x4e, 0xfe, 0x78, 0x3a, 0x53, 0xc5, 0x25,
	0x36, 0x88, 0x47, 0x31, 0xfe, 0x25, 0x86, 0xa7, 0x60, 0x6f, 0x8d, 0xd6, 0x2c, 0xa3, 0xeb, 0x18,
	0xa2, 0x20, 0x7f, 0x34, 0x9d, 0xa9, 0xbb, 0x4b, 0x58, 0x4b, 0x91, 0x47, 0xd1, 0x06, 0x3b, 0xe8,
	0xeb, 0x8c, 0xdd, 0xda, 0x60, 0x07, 0x49, 0xf0, 0x9a, 0xd5, 0x8d, 0x6b, 0xc3, 0x31, 0xc4, 0xed,
	0x0d, 0x56, 0x47, 0x63, 0x44, 0x11, 0xfc, 0x06, 0x48, 0xaf, 0xe6, 0xba, 0xb6, 0xd3, 0x75, 0x06,
	0xb6, 0x58, 0x92, 0x0f, 0xa7, 0x33, 0x75, 0x7f, 0x63, 0x7c, 0xfe, 0x05, 0xc0, 0x6f, 0xc1, 0xe1,
	0x9a, 0xd1, 0xe9, 0xda, 0x57, 0xae, 0x76, 0x7b, 0xd3, 0x67, 0xbb, 0x74, 0xb1, 0x2c, 0xcb, 0xd3,
	0x99, 0xfa, 0xc9, 0xd2, 0xc9, 0xbe, 0x1c, 0x0d, 0x47, 0x09, 0x5b, 0x19, 0x40, 0x0d, 0x28, 0x6b,
	0x56, 0xdb, 0xf8, 0xde, 0xb8, 0xe1, 0x95, 0xd5, 0xed, 0xd9, 0xda, 0xad, 0x6e, 0xe8, 0x62, 0x45,
	0x6e, 0x4e, 0x67, 0xea, 0xd1, 0xd2, 0x6f, 0xa3, 0x21, 0x8a, 0xd8, 0x75, 0xea, 0xc5, 0xc4, 0xc7,
	0x01, 0x0a, 0xe4, 0xbd, 0x5f, 0xff, 0x50, 0xde, 0xfd, 0xf5, 0xa7, 0xb2, 0x3a, 0xfd, 0x4b, 0xe9,
	0x69, 0xa1, 0x08, 0xff, 0x2c, 0x14, 0xe1, 0xdf, 0x85, 0x22, 0xfc, 0xfd, 0xac, 0x08, 0x4f, 0xcf,
	0x8a, 0xf0, 0xf3, 0xd6, 0xe4, 0xfc, 0xae, 0xc2, 0xff, 0x28, 0x5f, 0xff, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x1f, 0x28, 0x0d, 0xed, 0xee, 0x04, 0x00, 0x00,
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CostPerSec != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.CostPerSec))))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x81
	}
	if m.Duration != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Duration))))
		i--
		dAtA[i] = 0x79
	}
	if m.ChunkNum != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.ChunkNum))
		i--
		dAtA[i] = 0x70
	}
	if len(m.UserID) > 0 {
		i -= len(m.UserID)
		copy(dAtA[i:], m.UserID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.UserID)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.ClientUserID) > 0 {
		i -= len(m.ClientUserID)
		copy(dAtA[i:], m.ClientUserID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ClientUserID)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0x5a
	}
	if m.ProfileCost != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.ProfileCost))))
		i--
		dAtA[i] = 0x51
	}
	if len(m.ProfileName) > 0 {
		i -= len(m.ProfileName)
		copy(dAtA[i:], m.ProfileName)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ProfileName)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ProfileID) > 0 {
		i -= len(m.ProfileID)
		copy(dAtA[i:], m.ProfileID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ProfileID)))
		i--
		dAtA[i] = 0x42
	}
	if m.StreamIsLive {
		i--
		if m.StreamIsLive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.StreamContractAddress) > 0 {
		i -= len(m.StreamContractAddress)
		copy(dAtA[i:], m.StreamContractAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.StreamContractAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.StreamName) > 0 {
		i -= len(m.StreamName)
		copy(dAtA[i:], m.StreamName)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.StreamName)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.StreamID) > 0 {
		i -= len(m.StreamID)
		copy(dAtA[i:], m.StreamID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.StreamID)))
		i--
		dAtA[i] = 0x22
	}
	if m.Status != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TaskID) > 0 {
		i -= len(m.TaskID)
		copy(dAtA[i:], m.TaskID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.TaskID)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovEvent(uint64(m.Type))
	}
	l = len(m.TaskID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovEvent(uint64(m.Status))
	}
	l = len(m.StreamID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.StreamName)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.StreamContractAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.StreamIsLive {
		n += 2
	}
	l = len(m.ProfileID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.ProfileName)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.ProfileCost != 0 {
		n += 9
	}
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.ClientUserID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.ChunkNum != 0 {
		n += 1 + sovEvent(uint64(m.ChunkNum))
	}
	if m.Duration != 0 {
		n += 9
	}
	if m.CostPerSec != 0 {
		n += 10
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= EventType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaskID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StreamID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StreamName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StreamContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamIsLive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StreamIsLive = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfileName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfileCost", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.ProfileCost = float64(math.Float64frombits(v))
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientUserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientUserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkNum", wireType)
			}
			m.ChunkNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChunkNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Duration = float64(math.Float64frombits(v))
		case 16:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field CostPerSec", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.CostPerSec = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)