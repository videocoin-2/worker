// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: null.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

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

type NullString struct {
	Str                  string   `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NullString) Reset()         { *m = NullString{} }
func (m *NullString) String() string { return proto.CompactTextString(m) }
func (*NullString) ProtoMessage()    {}
func (*NullString) Descriptor() ([]byte, []int) {
	return fileDescriptor_null_a1714022790c5c6b, []int{0}
}
func (m *NullString) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NullString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NullString.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NullString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullString.Merge(dst, src)
}
func (m *NullString) XXX_Size() int {
	return m.Size()
}
func (m *NullString) XXX_DiscardUnknown() {
	xxx_messageInfo_NullString.DiscardUnknown(m)
}

var xxx_messageInfo_NullString proto.InternalMessageInfo

func (m *NullString) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *NullString) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type NullInt32 struct {
	Int32                int32    `protobuf:"varint,1,opt,name=int32,proto3" json:"int32,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NullInt32) Reset()         { *m = NullInt32{} }
func (m *NullInt32) String() string { return proto.CompactTextString(m) }
func (*NullInt32) ProtoMessage()    {}
func (*NullInt32) Descriptor() ([]byte, []int) {
	return fileDescriptor_null_a1714022790c5c6b, []int{1}
}
func (m *NullInt32) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NullInt32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NullInt32.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NullInt32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullInt32.Merge(dst, src)
}
func (m *NullInt32) XXX_Size() int {
	return m.Size()
}
func (m *NullInt32) XXX_DiscardUnknown() {
	xxx_messageInfo_NullInt32.DiscardUnknown(m)
}

var xxx_messageInfo_NullInt32 proto.InternalMessageInfo

func (m *NullInt32) GetInt32() int32 {
	if m != nil {
		return m.Int32
	}
	return 0
}

func (m *NullInt32) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type NullInt64 struct {
	Int64                int64    `protobuf:"varint,1,opt,name=int64,proto3" json:"int64,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NullInt64) Reset()         { *m = NullInt64{} }
func (m *NullInt64) String() string { return proto.CompactTextString(m) }
func (*NullInt64) ProtoMessage()    {}
func (*NullInt64) Descriptor() ([]byte, []int) {
	return fileDescriptor_null_a1714022790c5c6b, []int{2}
}
func (m *NullInt64) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NullInt64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NullInt64.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NullInt64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullInt64.Merge(dst, src)
}
func (m *NullInt64) XXX_Size() int {
	return m.Size()
}
func (m *NullInt64) XXX_DiscardUnknown() {
	xxx_messageInfo_NullInt64.DiscardUnknown(m)
}

var xxx_messageInfo_NullInt64 proto.InternalMessageInfo

func (m *NullInt64) GetInt64() int64 {
	if m != nil {
		return m.Int64
	}
	return 0
}

func (m *NullInt64) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type NullTime struct {
	Time                 *time.Time `protobuf:"bytes,1,opt,name=time,stdtime" json:"time,omitempty"`
	Valid                bool       `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NullTime) Reset()         { *m = NullTime{} }
func (m *NullTime) String() string { return proto.CompactTextString(m) }
func (*NullTime) ProtoMessage()    {}
func (*NullTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_null_a1714022790c5c6b, []int{3}
}
func (m *NullTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NullTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NullTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NullTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullTime.Merge(dst, src)
}
func (m *NullTime) XXX_Size() int {
	return m.Size()
}
func (m *NullTime) XXX_DiscardUnknown() {
	xxx_messageInfo_NullTime.DiscardUnknown(m)
}

var xxx_messageInfo_NullTime proto.InternalMessageInfo

func (m *NullTime) GetTime() *time.Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *NullTime) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*NullString)(nil), "proto.NullString")
	proto.RegisterType((*NullInt32)(nil), "proto.NullInt32")
	proto.RegisterType((*NullInt64)(nil), "proto.NullInt64")
	proto.RegisterType((*NullTime)(nil), "proto.NullTime")
}
func (m *NullString) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NullString) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Str) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNull(dAtA, i, uint64(len(m.Str)))
		i += copy(dAtA[i:], m.Str)
	}
	if m.Valid {
		dAtA[i] = 0x10
		i++
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *NullInt32) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NullInt32) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Int32 != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNull(dAtA, i, uint64(m.Int32))
	}
	if m.Valid {
		dAtA[i] = 0x10
		i++
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *NullInt64) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NullInt64) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Int64 != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNull(dAtA, i, uint64(m.Int64))
	}
	if m.Valid {
		dAtA[i] = 0x10
		i++
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *NullTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NullTime) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Time != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNull(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.Time)))
		n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.Time, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Valid {
		dAtA[i] = 0x10
		i++
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintNull(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NullString) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Str)
	if l > 0 {
		n += 1 + l + sovNull(uint64(l))
	}
	if m.Valid {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NullInt32) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Int32 != 0 {
		n += 1 + sovNull(uint64(m.Int32))
	}
	if m.Valid {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NullInt64) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Int64 != 0 {
		n += 1 + sovNull(uint64(m.Int64))
	}
	if m.Valid {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *NullTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Time != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.Time)
		n += 1 + l + sovNull(uint64(l))
	}
	if m.Valid {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNull(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNull(x uint64) (n int) {
	return sovNull(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NullString) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNull
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
			return fmt.Errorf("proto: NullString: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NullString: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Str", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNull
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
				return ErrInvalidLengthNull
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Str = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNull
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Valid = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNull(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNull
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
func (m *NullInt32) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNull
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
			return fmt.Errorf("proto: NullInt32: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NullInt32: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int32", wireType)
			}
			m.Int32 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNull
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Int32 |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNull
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Valid = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNull(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNull
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
func (m *NullInt64) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNull
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
			return fmt.Errorf("proto: NullInt64: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NullInt64: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64", wireType)
			}
			m.Int64 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNull
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Int64 |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNull
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Valid = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNull(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNull
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
func (m *NullTime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNull
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
			return fmt.Errorf("proto: NullTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NullTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNull
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNull
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Time == nil {
				m.Time = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNull
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Valid = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNull(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNull
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
func skipNull(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNull
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
					return 0, ErrIntOverflowNull
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
					return 0, ErrIntOverflowNull
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
				return 0, ErrInvalidLengthNull
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNull
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
				next, err := skipNull(dAtA[start:])
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
	ErrInvalidLengthNull = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNull   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("null.proto", fileDescriptor_null_a1714022790c5c6b) }

var fileDescriptor_null_a1714022790c5c6b = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x2b, 0xcd, 0xc9,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0xf2, 0xe9, 0xf9, 0xf9, 0xe9,
	0x39, 0xa9, 0xfa, 0x60, 0x5e, 0x52, 0x69, 0x9a, 0x7e, 0x49, 0x66, 0x6e, 0x6a, 0x71, 0x49, 0x62,
	0x6e, 0x01, 0x44, 0x9d, 0x94, 0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae,
	0x7e, 0x7a, 0x7e, 0x7a, 0x3e, 0x42, 0x25, 0x88, 0x07, 0xe6, 0x80, 0x59, 0x10, 0xe5, 0x4a, 0x26,
	0x5c, 0x5c, 0x7e, 0xa5, 0x39, 0x39, 0xc1, 0x25, 0x45, 0x99, 0x79, 0xe9, 0x42, 0x02, 0x5c, 0xcc,
	0xc5, 0x25, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b,
	0x59, 0x62, 0x4e, 0x66, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x84, 0xa3, 0x64, 0xce,
	0xc5, 0x09, 0xd2, 0xe5, 0x99, 0x57, 0x62, 0x6c, 0x04, 0x52, 0x92, 0x09, 0x62, 0x80, 0xb5, 0xb1,
	0x06, 0x41, 0x38, 0x04, 0x35, 0x9a, 0x99, 0x40, 0x35, 0x9a, 0x99, 0x80, 0x35, 0x32, 0x07, 0x41,
	0x38, 0x38, 0x34, 0x86, 0x71, 0x71, 0x80, 0x34, 0x86, 0x64, 0xe6, 0xa6, 0x0a, 0x99, 0x70, 0xb1,
	0x80, 0x7c, 0x0d, 0xd6, 0xc6, 0x6d, 0x24, 0xa5, 0x07, 0x09, 0x12, 0x3d, 0x98, 0x47, 0xf5, 0x42,
	0x60, 0x41, 0xe2, 0xc4, 0x32, 0xe1, 0xbe, 0x3c, 0x63, 0x10, 0x58, 0x35, 0x76, 0x73, 0x9d, 0x24,
	0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5,
	0x18, 0x3c, 0x18, 0xa3, 0x20, 0x41, 0x9d, 0xc4, 0x06, 0xa6, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xcf, 0xac, 0x70, 0xea, 0x86, 0x01, 0x00, 0x00,
}
