// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emitter/v1/event.proto

package v1

import (
	fmt "fmt"
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
	EventTypeUnknown        EventType = 0
	EventTypeAccountCreated EventType = 1
)

var EventType_name = map[int32]string{
	0: "EVENT_TYPE_UNKNOWN",
	1: "EVENT_TYPE_ACCOUNT_CREATED",
}

var EventType_value = map[string]int32{
	"EVENT_TYPE_UNKNOWN":         0,
	"EVENT_TYPE_ACCOUNT_CREATED": 1,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_71ff87d0e153db0a, []int{0}
}

type Event struct {
	Type                 EventType `protobuf:"varint,1,opt,name=type,proto3,enum=cloud.api.emitter.v1.EventType" json:"type,omitempty"`
	UserID               string    `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Address              string    `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_71ff87d0e153db0a, []int{0}
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

func (m *Event) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Event) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (*Event) XXX_MessageName() string {
	return "cloud.api.emitter.v1.Event"
}
func init() {
	proto.RegisterEnum("cloud.api.emitter.v1.EventType", EventType_name, EventType_value)
	golang_proto.RegisterEnum("cloud.api.emitter.v1.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Event)(nil), "cloud.api.emitter.v1.Event")
	golang_proto.RegisterType((*Event)(nil), "cloud.api.emitter.v1.Event")
}

func init() { proto.RegisterFile("emitter/v1/event.proto", fileDescriptor_71ff87d0e153db0a) }
func init() { golang_proto.RegisterFile("emitter/v1/event.proto", fileDescriptor_71ff87d0e153db0a) }

var fileDescriptor_71ff87d0e153db0a = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xcd, 0xcd, 0x2c,
	0x29, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x49, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x4b, 0x2c, 0xc8, 0xd4, 0x83, 0xaa,
	0xd0, 0x2b, 0x33, 0x94, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b, 0x4e, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b,
	0x62, 0x88, 0x52, 0x35, 0x17, 0xab, 0x2b, 0xc8, 0x4c, 0x21, 0x63, 0x2e, 0x96, 0x92, 0xca, 0x82,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x79, 0x3d, 0x6c, 0x86, 0xeb, 0x81, 0x95, 0x86,
	0x54, 0x16, 0xa4, 0x06, 0x81, 0x15, 0x0b, 0x29, 0x73, 0xb1, 0x97, 0x16, 0xa7, 0x16, 0xc5, 0x67,
	0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x3a, 0x71, 0x3d, 0xba, 0x27, 0xcf, 0x16, 0x5a, 0x9c,
	0x5a, 0xe4, 0xe9, 0x12, 0xc4, 0x06, 0x92, 0xf2, 0x4c, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x49,
	0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x06, 0x29, 0x0a, 0x82, 0x71, 0xb5, 0x3a, 0x19, 0xb9, 0x38,
	0xe1, 0x46, 0x0a, 0xe9, 0x70, 0x09, 0xb9, 0x86, 0xb9, 0xfa, 0x85, 0xc4, 0x87, 0x44, 0x06, 0xb8,
	0xc6, 0x87, 0xfa, 0x79, 0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30, 0x48, 0x89, 0x74, 0xcd, 0x55, 0x10,
	0x80, 0x2b, 0x0b, 0xcd, 0xcb, 0xce, 0xcb, 0x2f, 0xcf, 0x13, 0xb2, 0xe6, 0x92, 0x42, 0x52, 0xed,
	0xe8, 0xec, 0xec, 0x1f, 0xea, 0x17, 0x12, 0xef, 0x1c, 0xe4, 0xea, 0x18, 0xe2, 0xea, 0x22, 0xc0,
	0x28, 0x25, 0xdd, 0x35, 0x57, 0x41, 0x1c, 0xae, 0xcb, 0x31, 0x39, 0x39, 0xbf, 0x34, 0xaf, 0xc4,
	0xb9, 0x28, 0x35, 0xb1, 0x24, 0x35, 0x45, 0x4a, 0xb0, 0x63, 0xb1, 0x1c, 0xc3, 0xae, 0x25, 0x72,
	0x08, 0xdb, 0x9d, 0x24, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x03, 0x8f, 0xe5, 0x18, 0x4f, 0x3c, 0x96, 0x63, 0x8c, 0x62, 0x2a, 0x33, 0x4c, 0x62, 0x03,
	0x87, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x85, 0x23, 0x5f, 0x88, 0x01, 0x00, 0x00,
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
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.UserID) > 0 {
		i -= len(m.UserID)
		copy(dAtA[i:], m.UserID)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.UserID)))
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
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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