// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: syncer/v1/syncer_service.proto

package v1

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	rpc "github.com/videocoin/cloud-api/rpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SyncRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	ContentType          string   `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Duration             float64  `protobuf:"fixed64,4,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRequest) Reset()         { *m = SyncRequest{} }
func (m *SyncRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()    {}
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e555e157b8570a45, []int{0}
}
func (m *SyncRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRequest.Merge(m, src)
}
func (m *SyncRequest) XXX_Size() int {
	return m.Size()
}
func (m *SyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRequest proto.InternalMessageInfo

func (m *SyncRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SyncRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *SyncRequest) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *SyncRequest) GetDuration() float64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (*SyncRequest) XXX_MessageName() string {
	return "cloud.api.syncer.v1.SyncRequest"
}
func init() {
	proto.RegisterType((*SyncRequest)(nil), "cloud.api.syncer.v1.SyncRequest")
	golang_proto.RegisterType((*SyncRequest)(nil), "cloud.api.syncer.v1.SyncRequest")
}

func init() { proto.RegisterFile("syncer/v1/syncer_service.proto", fileDescriptor_e555e157b8570a45) }
func init() {
	golang_proto.RegisterFile("syncer/v1/syncer_service.proto", fileDescriptor_e555e157b8570a45)
}

var fileDescriptor_e555e157b8570a45 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x7d, 0xd3, 0x57, 0x4a, 0xdf, 0xb4, 0x0f, 0x64, 0x04, 0x09, 0xa9, 0x84, 0xd8, 0x55, 0x37,
	0x9d, 0xa1, 0xfa, 0x07, 0x05, 0xc1, 0x9d, 0x90, 0xb8, 0x72, 0x53, 0xa6, 0x93, 0x31, 0x09, 0xb4,
	0x33, 0x63, 0x72, 0x13, 0x88, 0x4b, 0x7f, 0xc1, 0x5f, 0xf0, 0x43, 0x5c, 0x76, 0x29, 0xf8, 0x03,
	0xd2, 0xfa, 0x21, 0x92, 0x49, 0xb4, 0x5d, 0xe8, 0xee, 0xdc, 0x7b, 0xce, 0x9d, 0x7b, 0xcf, 0x19,
	0xec, 0xe5, 0x95, 0x12, 0x32, 0x63, 0xe5, 0x8c, 0x35, 0x68, 0x91, 0xcb, 0xac, 0x4c, 0x85, 0xa4,
	0x26, 0xd3, 0xa0, 0xc9, 0xb1, 0x58, 0xe9, 0x22, 0xa2, 0xdc, 0xa4, 0xb4, 0xe1, 0x69, 0x39, 0x73,
	0x47, 0xb1, 0xd6, 0xf1, 0x4a, 0x32, 0x2b, 0x59, 0x16, 0x77, 0x4c, 0xae, 0x0d, 0x54, 0xcd, 0x84,
	0x7b, 0xda, 0x92, 0xdc, 0xa4, 0x8c, 0x2b, 0xa5, 0x81, 0x43, 0xaa, 0x55, 0xde, 0xb2, 0xd3, 0x38,
	0x85, 0xa4, 0x58, 0x52, 0xa1, 0xd7, 0x2c, 0xd6, 0xb1, 0xde, 0xbf, 0x51, 0x57, 0xb6, 0xb0, 0xa8,
	0x95, 0xb3, 0x03, 0x79, 0x99, 0x46, 0x52, 0x0b, 0x9d, 0x2a, 0x66, 0x6f, 0x9a, 0xd6, 0x0b, 0x32,
	0x23, 0x58, 0x22, 0xf9, 0x0a, 0x92, 0x66, 0x60, 0x0c, 0x78, 0x10, 0x56, 0x4a, 0x04, 0xf2, 0xbe,
	0x90, 0x39, 0x10, 0x82, 0xbb, 0x11, 0x07, 0xee, 0x20, 0x1f, 0x4d, 0x86, 0x81, 0xc5, 0x75, 0xcf,
	0x70, 0x48, 0x9c, 0x8e, 0x8f, 0x26, 0xff, 0x02, 0x8b, 0xc9, 0x19, 0x1e, 0x0a, 0xad, 0x40, 0x2a,
	0x58, 0x40, 0x65, 0xa4, 0xf3, 0xd7, 0x72, 0x83, 0xb6, 0x77, 0x53, 0x19, 0x49, 0x5c, 0xdc, 0x8f,
	0x8a, 0xcc, 0x9a, 0x71, 0xba, 0x3e, 0x9a, 0xa0, 0xe0, 0xbb, 0x3e, 0x7f, 0x46, 0xf8, 0x7f, 0x68,
	0xe3, 0x09, 0x9b, 0xf4, 0xc8, 0x35, 0xee, 0x5d, 0xd9, 0xbb, 0xc8, 0x09, 0x6d, 0x02, 0xa1, 0x5f,
	0x4e, 0xe9, 0x65, 0x9d, 0x96, 0x3b, 0xa2, 0xfb, 0x68, 0x33, 0x23, 0x68, 0x23, 0x0f, 0x81, 0x43,
	0x91, 0x8f, 0x8f, 0x1e, 0xdf, 0x3e, 0x9e, 0x3a, 0x98, 0xf4, 0x5b, 0x77, 0x0f, 0x64, 0x8e, 0xbb,
	0xf5, 0x06, 0xe2, 0xd3, 0x1f, 0x7e, 0x84, 0x1e, 0x78, 0x76, 0x7f, 0x59, 0x38, 0xfe, 0x33, 0x77,
	0x36, 0x5b, 0x0f, 0xbd, 0x6e, 0x3d, 0xf4, 0xbe, 0xf5, 0xd0, 0xcb, 0xce, 0x43, 0x9b, 0x9d, 0x87,
	0x6e, 0x3b, 0xe5, 0x6c, 0xd9, 0xb3, 0xda, 0x8b, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0xfe,
	0x82, 0x49, 0x0f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SyncerServiceClient is the client API for SyncerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncerServiceClient interface {
	Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error)
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type syncerServiceClient struct {
	cc *grpc.ClientConn
}

func NewSyncerServiceClient(cc *grpc.ClientConn) SyncerServiceClient {
	return &syncerServiceClient{cc}
}

func (c *syncerServiceClient) Health(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*rpc.HealthStatus, error) {
	out := new(rpc.HealthStatus)
	err := c.cc.Invoke(ctx, "/cloud.api.syncer.v1.SyncerService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncerServiceClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/cloud.api.syncer.v1.SyncerService/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncerServiceServer is the server API for SyncerService service.
type SyncerServiceServer interface {
	Health(context.Context, *types.Empty) (*rpc.HealthStatus, error)
	Sync(context.Context, *SyncRequest) (*types.Empty, error)
}

// UnimplementedSyncerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSyncerServiceServer struct {
}

func (*UnimplementedSyncerServiceServer) Health(ctx context.Context, req *types.Empty) (*rpc.HealthStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedSyncerServiceServer) Sync(ctx context.Context, req *SyncRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}

func RegisterSyncerServiceServer(s *grpc.Server, srv SyncerServiceServer) {
	s.RegisterService(&_SyncerService_serviceDesc, srv)
}

func _SyncerService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.syncer.v1.SyncerService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).Health(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncerService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncerServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.api.syncer.v1.SyncerService/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncerServiceServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SyncerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.api.syncer.v1.SyncerService",
	HandlerType: (*SyncerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _SyncerService_Health_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _SyncerService_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syncer/v1/syncer_service.proto",
}

func (m *SyncRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSyncerService(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if len(m.Path) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSyncerService(dAtA, i, uint64(len(m.Path)))
		i += copy(dAtA[i:], m.Path)
	}
	if len(m.ContentType) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSyncerService(dAtA, i, uint64(len(m.ContentType)))
		i += copy(dAtA[i:], m.ContentType)
	}
	if m.Duration != 0 {
		dAtA[i] = 0x21
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Duration))))
		i += 8
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSyncerService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SyncRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovSyncerService(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovSyncerService(uint64(l))
	}
	l = len(m.ContentType)
	if l > 0 {
		n += 1 + l + sovSyncerService(uint64(l))
	}
	if m.Duration != 0 {
		n += 9
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSyncerService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSyncerService(x uint64) (n int) {
	return sovSyncerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SyncRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSyncerService
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
			return fmt.Errorf("proto: SyncRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSyncerService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSyncerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncerService
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
				return ErrInvalidLengthSyncerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSyncerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncerService
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
				return ErrInvalidLengthSyncerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSyncerService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
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
		default:
			iNdEx = preIndex
			skippy, err := skipSyncerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSyncerService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSyncerService
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
func skipSyncerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSyncerService
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
					return 0, ErrIntOverflowSyncerService
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
					return 0, ErrIntOverflowSyncerService
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
				return 0, ErrInvalidLengthSyncerService
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSyncerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSyncerService
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
				next, err := skipSyncerService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSyncerService
				}
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
	ErrInvalidLengthSyncerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSyncerService   = fmt.Errorf("proto: integer overflow")
)