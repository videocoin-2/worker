// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: verifier.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/mwitkow/go-proto-validators"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

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

type VerifyRequest struct {
	Bucket               string   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	JobId                uint32   `protobuf:"varint,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	SourceChunk          string   `protobuf:"bytes,3,opt,name=source_chunk,json=sourceChunk,proto3" json:"source_chunk,omitempty"`
	ResultChunk          string   `protobuf:"bytes,4,opt,name=result_chunk,json=resultChunk,proto3" json:"result_chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_verifier_b4efed012f6be4e9, []int{0}
}
func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(dst, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return m.Size()
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *VerifyRequest) GetJobId() uint32 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func (m *VerifyRequest) GetSourceChunk() string {
	if m != nil {
		return m.SourceChunk
	}
	return ""
}

func (m *VerifyRequest) GetResultChunk() string {
	if m != nil {
		return m.ResultChunk
	}
	return ""
}

func init() {
	proto.RegisterType((*VerifyRequest)(nil), "proto.VerifyRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VerifierServiceClient is the client API for VerifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VerifierServiceClient interface {
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type verifierServiceClient struct {
	cc *grpc.ClientConn
}

func NewVerifierServiceClient(cc *grpc.ClientConn) VerifierServiceClient {
	return &verifierServiceClient{cc}
}

func (c *verifierServiceClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.VerifierService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerifierServiceServer is the server API for VerifierService service.
type VerifierServiceServer interface {
	Verify(context.Context, *VerifyRequest) (*empty.Empty, error)
}

func RegisterVerifierServiceServer(s *grpc.Server, srv VerifierServiceServer) {
	s.RegisterService(&_VerifierService_serviceDesc, srv)
}

func _VerifierService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifierServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VerifierService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifierServiceServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VerifierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VerifierService",
	HandlerType: (*VerifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _VerifierService_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "verifier.proto",
}

func (m *VerifyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Bucket) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintVerifier(dAtA, i, uint64(len(m.Bucket)))
		i += copy(dAtA[i:], m.Bucket)
	}
	if m.JobId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintVerifier(dAtA, i, uint64(m.JobId))
	}
	if len(m.SourceChunk) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintVerifier(dAtA, i, uint64(len(m.SourceChunk)))
		i += copy(dAtA[i:], m.SourceChunk)
	}
	if len(m.ResultChunk) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintVerifier(dAtA, i, uint64(len(m.ResultChunk)))
		i += copy(dAtA[i:], m.ResultChunk)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintVerifier(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *VerifyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bucket)
	if l > 0 {
		n += 1 + l + sovVerifier(uint64(l))
	}
	if m.JobId != 0 {
		n += 1 + sovVerifier(uint64(m.JobId))
	}
	l = len(m.SourceChunk)
	if l > 0 {
		n += 1 + l + sovVerifier(uint64(l))
	}
	l = len(m.ResultChunk)
	if l > 0 {
		n += 1 + l + sovVerifier(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVerifier(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozVerifier(x uint64) (n int) {
	return sovVerifier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VerifyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifier
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
			return fmt.Errorf("proto: VerifyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bucket", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifier
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
				return ErrInvalidLengthVerifier
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bucket = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobId", wireType)
			}
			m.JobId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JobId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChunk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifier
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
				return ErrInvalidLengthVerifier
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChunk = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultChunk", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifier
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
				return ErrInvalidLengthVerifier
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultChunk = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerifier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerifier
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
func skipVerifier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVerifier
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
					return 0, ErrIntOverflowVerifier
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
					return 0, ErrIntOverflowVerifier
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
				return 0, ErrInvalidLengthVerifier
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowVerifier
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
				next, err := skipVerifier(dAtA[start:])
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
	ErrInvalidLengthVerifier = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVerifier   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("verifier.proto", fileDescriptor_verifier_b4efed012f6be4e9) }

var fileDescriptor_verifier_b4efed012f6be4e9 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcf, 0x4e, 0x02, 0x31,
	0x10, 0xc6, 0x29, 0xca, 0x26, 0x56, 0x41, 0xb3, 0x51, 0x42, 0xd0, 0xac, 0xc8, 0x89, 0x0b, 0xdb,
	0xa8, 0x89, 0x0f, 0xa0, 0xf1, 0xe0, 0xc5, 0x03, 0x26, 0x1c, 0xbc, 0x90, 0xed, 0x52, 0x4a, 0xf9,
	0xb3, 0x83, 0xdd, 0x76, 0x09, 0x57, 0xe3, 0x1b, 0x78, 0xf1, 0x71, 0x3c, 0x72, 0x34, 0xf1, 0x05,
	0x0c, 0xf8, 0x20, 0x66, 0xdb, 0xfa, 0xef, 0xd4, 0xf9, 0x66, 0x7e, 0xd3, 0x7c, 0xf9, 0x06, 0x57,
	0x32, 0x26, 0xc5, 0x40, 0x30, 0x19, 0xce, 0x24, 0x28, 0xf0, 0x4b, 0xe6, 0xa9, 0x1f, 0x72, 0x00,
	0x3e, 0x61, 0xc4, 0x28, 0xaa, 0x07, 0x84, 0x4d, 0x67, 0x6a, 0x61, 0x99, 0xfa, 0x91, 0x1b, 0x46,
	0x33, 0x41, 0xa2, 0x24, 0x01, 0x15, 0x29, 0x01, 0x49, 0xea, 0xa6, 0x6d, 0x2e, 0xd4, 0x50, 0xd3,
	0x30, 0x86, 0x29, 0xe1, 0xc0, 0xe1, 0xf7, 0x8f, 0x5c, 0x19, 0x61, 0x2a, 0x87, 0x5f, 0xfc, 0xc1,
	0xa7, 0x73, 0xa1, 0xc6, 0x30, 0x27, 0x1c, 0xda, 0x66, 0xd8, 0xce, 0xa2, 0x89, 0xe8, 0x47, 0x0a,
	0x64, 0x4a, 0x7e, 0x4a, 0xbb, 0xd7, 0x7c, 0x42, 0xb8, 0xdc, 0xcd, 0xbd, 0x2f, 0x3a, 0xec, 0x41,
	0xb3, 0x54, 0xf9, 0x55, 0xec, 0x51, 0x1d, 0x8f, 0x99, 0xaa, 0xa1, 0x06, 0x6a, 0x6d, 0x75, 0x9c,
	0xf2, 0x0f, 0xb0, 0x37, 0x02, 0xda, 0x13, 0xfd, 0x5a, 0xb1, 0x81, 0x5a, 0xe5, 0x4e, 0x69, 0x04,
	0xf4, 0xa6, 0xef, 0x9f, 0xe0, 0x9d, 0x14, 0xb4, 0x8c, 0x59, 0x2f, 0x1e, 0xea, 0x64, 0x5c, 0xdb,
	0x30, 0x4b, 0xdb, 0xb6, 0x77, 0x95, 0xb7, 0x72, 0x44, 0xb2, 0x54, 0x4f, 0x94, 0x43, 0x36, 0x2d,
	0x62, 0x7b, 0x06, 0x39, 0x8b, 0xf0, 0x6e, 0xd7, 0x25, 0x78, 0xc7, 0x64, 0x26, 0x62, 0xe6, 0xdf,
	0x62, 0xcf, 0x1a, 0xf3, 0xf7, 0xad, 0xd7, 0xf0, 0x9f, 0xcf, 0x7a, 0x35, 0xb4, 0xf9, 0x85, 0xdf,
	0xc1, 0x84, 0xd7, 0x79, 0xb8, 0xcd, 0xea, 0xe3, 0xfb, 0xe7, 0x73, 0x71, 0xaf, 0x59, 0x31, 0xc1,
	0x66, 0xa7, 0xc4, 0x9c, 0x66, 0x71, 0x79, 0xbc, 0x5c, 0x05, 0xe8, 0x6d, 0x15, 0xa0, 0x8f, 0x55,
	0x80, 0x5e, 0xd6, 0x41, 0xe1, 0x75, 0x1d, 0xa0, 0xe5, 0x3a, 0x28, 0xdc, 0xdb, 0x63, 0x51, 0xcf,
	0x3c, 0xe7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xba, 0x65, 0xbc, 0xcc, 0x01, 0x00, 0x00,
}
