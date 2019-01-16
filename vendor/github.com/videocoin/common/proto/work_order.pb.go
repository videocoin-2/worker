// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: work_order.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkOrderStatus int32

const (
	WorkOrderStatusNone        WorkOrderStatus = 0
	WorkOrderStatusPending     WorkOrderStatus = 1
	WorkOrderStatusWorkStarted WorkOrderStatus = 2
	WorkOrderStatusTranscoding WorkOrderStatus = 3
	WorkOrderStatusCanceled    WorkOrderStatus = 4
	WorkOrderStatusFailed      WorkOrderStatus = 5
	WorkOrderStatusCompleted   WorkOrderStatus = 6
	WorkOrderStatusReady       WorkOrderStatus = 7
)

var WorkOrderStatus_name = map[int32]string{
	0: "none",
	1: "pending",
	2: "work_started",
	3: "transcoding",
	4: "canceld",
	5: "failed",
	6: "completed",
	7: "ready",
}
var WorkOrderStatus_value = map[string]int32{
	"none":         0,
	"pending":      1,
	"work_started": 2,
	"transcoding":  3,
	"canceld":      4,
	"failed":       5,
	"completed":    6,
	"ready":        7,
}

func (x WorkOrderStatus) String() string {
	return proto.EnumName(WorkOrderStatus_name, int32(x))
}
func (WorkOrderStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_work_order_a9acaa9863af700f, []int{0}
}

type WorkOrder struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TranscoderId         string   `protobuf:"bytes,2,opt,name=transcoder_id,json=transcoderId,proto3" json:"transcoder_id,omitempty"`
	StreamHash           string   `protobuf:"bytes,3,opt,name=stream_hash,json=streamHash,proto3" json:"stream_hash,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ApplicationId        string   `protobuf:"bytes,5,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	InputUrl             string   `protobuf:"bytes,6,opt,name=input_url,json=inputUrl,proto3" json:"input_url,omitempty"`
	OutputUrl            string   `protobuf:"bytes,7,opt,name=output_url,json=outputUrl,proto3" json:"output_url,omitempty"`
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Chunks               []byte   `protobuf:"bytes,9,opt,name=chunks,proto3" json:"chunks,omitempty" gorm:"chunks;type:json;DEFAULT:NULL"sql:"type:json;DEFAULT:null"`
	CreatedAt            string   `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" gorm:"created_at;type:DATETIME;DEFAULT:current_timestamp" sql:"type:DATETIME;DEFAULT:current_timestamp"`
	UpdatedAt            string   `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" gorm:"updated_at;type:DATETIME;DEFAULT:current_timestamp" sql:"type:DATETIME;DEFAULT:current_timestamp"`
	Worker               []byte   `protobuf:"bytes,12,opt,name=worker,proto3" json:"worker,omitempty" gorm:"worker;type:binary(32);DEFAULT:NULL"sql:"type:binary(32);DEFAULT:null"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkOrder) Reset()         { *m = WorkOrder{} }
func (m *WorkOrder) String() string { return proto.CompactTextString(m) }
func (*WorkOrder) ProtoMessage()    {}
func (*WorkOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_work_order_a9acaa9863af700f, []int{0}
}
func (m *WorkOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *WorkOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkOrder.Merge(dst, src)
}
func (m *WorkOrder) XXX_Size() int {
	return m.Size()
}
func (m *WorkOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkOrder.DiscardUnknown(m)
}

var xxx_messageInfo_WorkOrder proto.InternalMessageInfo

func (m *WorkOrder) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WorkOrder) GetTranscoderId() string {
	if m != nil {
		return m.TranscoderId
	}
	return ""
}

func (m *WorkOrder) GetStreamHash() string {
	if m != nil {
		return m.StreamHash
	}
	return ""
}

func (m *WorkOrder) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *WorkOrder) GetApplicationId() string {
	if m != nil {
		return m.ApplicationId
	}
	return ""
}

func (m *WorkOrder) GetInputUrl() string {
	if m != nil {
		return m.InputUrl
	}
	return ""
}

func (m *WorkOrder) GetOutputUrl() string {
	if m != nil {
		return m.OutputUrl
	}
	return ""
}

func (m *WorkOrder) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *WorkOrder) GetChunks() []byte {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func (m *WorkOrder) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *WorkOrder) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *WorkOrder) GetWorker() []byte {
	if m != nil {
		return m.Worker
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkOrder)(nil), "proto.WorkOrder")
	proto.RegisterEnum("proto.WorkOrderStatus", WorkOrderStatus_name, WorkOrderStatus_value)
}
func (m *WorkOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkOrder) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(m.Id))
	}
	if len(m.TranscoderId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(len(m.TranscoderId)))
		i += copy(dAtA[i:], m.TranscoderId)
	}
	if len(m.StreamHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(len(m.StreamHash)))
		i += copy(dAtA[i:], m.StreamHash)
	}
	if len(m.UserId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if len(m.ApplicationId) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(len(m.ApplicationId)))
		i += copy(dAtA[i:], m.ApplicationId)
	}
	if len(m.InputUrl) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(len(m.InputUrl)))
		i += copy(dAtA[i:], m.InputUrl)
	}
	if len(m.OutputUrl) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(len(m.OutputUrl)))
		i += copy(dAtA[i:], m.OutputUrl)
	}
	if len(m.Status) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(len(m.Status)))
		i += copy(dAtA[i:], m.Status)
	}
	if len(m.Chunks) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(len(m.Chunks)))
		i += copy(dAtA[i:], m.Chunks)
	}
	if len(m.CreatedAt) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(len(m.CreatedAt)))
		i += copy(dAtA[i:], m.CreatedAt)
	}
	if len(m.UpdatedAt) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(len(m.UpdatedAt)))
		i += copy(dAtA[i:], m.UpdatedAt)
	}
	if len(m.Worker) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintWorkOrder(dAtA, i, uint64(len(m.Worker)))
		i += copy(dAtA[i:], m.Worker)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintWorkOrder(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *WorkOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovWorkOrder(uint64(m.Id))
	}
	l = len(m.TranscoderId)
	if l > 0 {
		n += 1 + l + sovWorkOrder(uint64(l))
	}
	l = len(m.StreamHash)
	if l > 0 {
		n += 1 + l + sovWorkOrder(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovWorkOrder(uint64(l))
	}
	l = len(m.ApplicationId)
	if l > 0 {
		n += 1 + l + sovWorkOrder(uint64(l))
	}
	l = len(m.InputUrl)
	if l > 0 {
		n += 1 + l + sovWorkOrder(uint64(l))
	}
	l = len(m.OutputUrl)
	if l > 0 {
		n += 1 + l + sovWorkOrder(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovWorkOrder(uint64(l))
	}
	l = len(m.Chunks)
	if l > 0 {
		n += 1 + l + sovWorkOrder(uint64(l))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovWorkOrder(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 1 + l + sovWorkOrder(uint64(l))
	}
	l = len(m.Worker)
	if l > 0 {
		n += 1 + l + sovWorkOrder(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWorkOrder(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWorkOrder(x uint64) (n int) {
	return sovWorkOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorkOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkOrder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WorkOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TranscoderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TranscoderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreamHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StreamHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InputUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutputUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWorkOrder
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks[:0], dAtA[iNdEx:postIndex]...)
			if m.Chunks == nil {
				m.Chunks = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Worker", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWorkOrder
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Worker = append(m.Worker[:0], dAtA[iNdEx:postIndex]...)
			if m.Worker == nil {
				m.Worker = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkOrder
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
func skipWorkOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkOrder
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
					return 0, ErrIntOverflowWorkOrder
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkOrder
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthWorkOrder
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWorkOrder
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWorkOrder(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWorkOrder = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkOrder   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("work_order.proto", fileDescriptor_work_order_a9acaa9863af700f) }

var fileDescriptor_work_order_a9acaa9863af700f = []byte{
	// 690 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0x86, 0xe3, 0xb4, 0x71, 0x9a, 0x69, 0xda, 0x1b, 0xcd, 0xbd, 0xb7, 0x71, 0x5d, 0x70, 0x4c,
	0x2a, 0x44, 0x00, 0xd1, 0x20, 0xba, 0x4b, 0x37, 0xa4, 0xb4, 0x15, 0x91, 0x42, 0xa9, 0xd2, 0x54,
	0x48, 0x08, 0x29, 0x9a, 0x78, 0xa6, 0x8e, 0x89, 0xe3, 0x31, 0xe3, 0xb1, 0x50, 0xde, 0x00, 0xf9,
	0x09, 0xd8, 0x64, 0x05, 0x4f, 0xc1, 0x8a, 0x65, 0x97, 0xec, 0xd8, 0x55, 0xa8, 0x7d, 0x83, 0xae,
	0x59, 0x20, 0xcf, 0x38, 0x09, 0x32, 0x45, 0x62, 0xc3, 0xca, 0x3e, 0xe7, 0xfb, 0x7d, 0xce, 0x9f,
	0x33, 0x73, 0x02, 0x4a, 0x6f, 0x29, 0x1b, 0xf6, 0x28, 0xc3, 0x84, 0x6d, 0xf9, 0x8c, 0x72, 0x0a,
	0x73, 0xe2, 0xa1, 0x57, 0x6c, 0x4a, 0x6d, 0x97, 0xd4, 0x45, 0xd4, 0x0f, 0x4f, 0xeb, 0xdc, 0x19,
	0x91, 0x80, 0xa3, 0x91, 0x2f, 0x75, 0xba, 0x91, 0x16, 0xe0, 0x90, 0x21, 0xee, 0x50, 0x2f, 0xe1,
	0xeb, 0x69, 0x8e, 0xbc, 0x71, 0x82, 0x1e, 0xd8, 0x0e, 0x1f, 0x84, 0xfd, 0x2d, 0x8b, 0x8e, 0xea,
	0x36, 0xb5, 0xe9, 0x5c, 0x13, 0x47, 0x22, 0x10, 0x6f, 0x89, 0x1c, 0x78, 0xa1, 0xeb, 0xca, 0xf7,
	0xea, 0xd7, 0x1c, 0x28, 0xbc, 0xa0, 0x6c, 0xf8, 0x3c, 0x76, 0x0c, 0x57, 0x41, 0xd6, 0xc1, 0x9a,
	0x62, 0x2a, 0xb5, 0x95, 0x4e, 0xd6, 0xc1, 0x70, 0x13, 0xac, 0x70, 0x86, 0xbc, 0xc0, 0xa2, 0x98,
	0xb0, 0x9e, 0x83, 0xb5, 0xac, 0xa9, 0xd4, 0x0a, 0x9d, 0xe2, 0x3c, 0xd9, 0xc2, 0xb0, 0x02, 0x96,
	0x03, 0xce, 0x08, 0x1a, 0xf5, 0x06, 0x28, 0x18, 0x68, 0x0b, 0x42, 0x02, 0x64, 0xea, 0x29, 0x0a,
	0x06, 0xb0, 0x0c, 0xf2, 0x61, 0x20, 0xbf, 0x5f, 0x14, 0x50, 0x8d, 0xc3, 0x16, 0x86, 0xb7, 0xc1,
	0x2a, 0xf2, 0x7d, 0xd7, 0xb1, 0xc4, 0xef, 0x8c, 0x79, 0x4e, 0xf0, 0x95, 0x9f, 0xb2, 0x2d, 0x0c,
	0x37, 0x40, 0xc1, 0xf1, 0xfc, 0x90, 0xf7, 0x42, 0xe6, 0x6a, 0xaa, 0x50, 0x2c, 0x89, 0xc4, 0x09,
	0x73, 0xe1, 0x4d, 0x00, 0x68, 0xc8, 0xa7, 0x34, 0x2f, 0x68, 0x41, 0x66, 0x62, 0xbc, 0x06, 0xd4,
	0x80, 0x23, 0x1e, 0x06, 0xda, 0x92, 0x6c, 0x2d, 0x23, 0xf8, 0x0a, 0xa8, 0xd6, 0x20, 0xf4, 0x86,
	0x81, 0x56, 0x30, 0x95, 0x5a, 0x71, 0x77, 0xef, 0xea, 0xbc, 0xf2, 0xd8, 0xa6, 0x6c, 0xd4, 0xa8,
	0xca, 0xfc, 0x0e, 0x1f, 0xfb, 0xa4, 0xf1, 0x3a, 0xa0, 0xde, 0xce, 0xde, 0xfe, 0x41, 0xf3, 0xa4,
	0xdd, 0x6d, 0x1c, 0x9e, 0xb4, 0xdb, 0xd5, 0xe0, 0x8d, 0xdb, 0xa8, 0xfe, 0xca, 0xe2, 0x91, 0x56,
	0x3b, 0x49, 0x4d, 0x18, 0x29, 0x00, 0x58, 0x8c, 0x20, 0x4e, 0x70, 0x0f, 0x71, 0x0d, 0xc4, 0xad,
	0x77, 0x87, 0x57, 0xe7, 0x15, 0x3b, 0x69, 0x31, 0x63, 0xb2, 0xcd, 0x5e, 0xb3, 0xbb, 0xdf, 0x6d,
	0x3d, 0xdb, 0x9f, 0x95, 0xb3, 0x42, 0xc6, 0x88, 0xc7, 0x7b, 0xb3, 0x4b, 0x52, 0x35, 0xe7, 0x8d,
	0xff, 0x40, 0xdd, 0x29, 0x24, 0x2d, 0x9a, 0x5c, 0x98, 0x09, 0x7d, 0x3c, 0x35, 0xb3, 0x9c, 0x36,
	0x33, 0x67, 0x7f, 0xcd, 0x4c, 0xd2, 0xa2, 0xc9, 0xe1, 0x00, 0xa8, 0xf1, 0x86, 0x10, 0xa6, 0x15,
	0xc5, 0xdc, 0x8f, 0xae, 0xce, 0x2b, 0x6d, 0xe9, 0x43, 0xe6, 0xa5, 0x87, 0xbe, 0xe3, 0x21, 0x36,
	0xae, 0x6d, 0x3f, 0xba, 0xfb, 0xbb, 0xe9, 0x5f, 0xa3, 0x48, 0xce, 0x40, 0xd6, 0xb9, 0xf7, 0x3d,
	0x0b, 0xfe, 0x99, 0xdd, 0xec, 0x63, 0x79, 0xea, 0xb7, 0xc0, 0xa2, 0x47, 0x3d, 0x52, 0xca, 0xe8,
	0xe5, 0x68, 0x62, 0xfe, 0x9b, 0xc2, 0x87, 0xd4, 0x23, 0xf0, 0x0e, 0xc8, 0xfb, 0xc4, 0xc3, 0x8e,
	0x67, 0x97, 0x14, 0x5d, 0x8f, 0x26, 0xe6, 0x5a, 0x4a, 0x75, 0x24, 0x29, 0x7c, 0x08, 0x8a, 0x62,
	0xd7, 0x03, 0x8e, 0x18, 0x27, 0xb8, 0x94, 0xd5, 0x8d, 0x68, 0x62, 0xea, 0x29, 0x75, 0x1c, 0x1e,
	0x4b, 0x05, 0xac, 0x83, 0xe5, 0xe9, 0xe2, 0xc4, 0xe5, 0x17, 0xae, 0xfd, 0xa0, 0x3b, 0x57, 0xc0,
	0x1a, 0xc8, 0x5b, 0xc8, 0xb3, 0x88, 0x8b, 0x4b, 0x8b, 0xfa, 0x46, 0x34, 0x31, 0xcb, 0x29, 0xf1,
	0x13, 0x41, 0x49, 0xbc, 0x49, 0xea, 0x29, 0x72, 0x5c, 0x82, 0x4b, 0x39, 0x7d, 0x3d, 0x9a, 0x98,
	0xff, 0xa7, 0x84, 0x07, 0x02, 0xc2, 0xfb, 0xa0, 0x60, 0xd1, 0x91, 0xef, 0x92, 0xd8, 0xb0, 0xaa,
	0xdf, 0x88, 0x26, 0xa6, 0x96, 0x2e, 0x39, 0xe5, 0x70, 0x13, 0xe4, 0x18, 0x41, 0x78, 0x5c, 0xca,
	0xeb, 0x5a, 0x34, 0x31, 0xff, 0x4b, 0x09, 0x3b, 0x31, 0xd3, 0xcb, 0xef, 0x3e, 0x18, 0x99, 0x4f,
	0x1f, 0x8d, 0xf4, 0xa8, 0x77, 0x2b, 0x67, 0x17, 0x86, 0xf2, 0xe5, 0xc2, 0x50, 0xbe, 0x5d, 0x18,
	0xca, 0xfb, 0x4b, 0x23, 0xf3, 0xf9, 0xd2, 0x50, 0xce, 0x2e, 0x8d, 0xcc, 0x4b, 0xf9, 0x87, 0xd8,
	0x57, 0xc5, 0x63, 0xfb, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x47, 0x9b, 0x63, 0x32, 0x05,
	0x00, 0x00,
}
