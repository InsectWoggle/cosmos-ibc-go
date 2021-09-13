// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/port/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryAppVersionRequest is the request type for the Query/AppVersion RPC method
type QueryAppVersionRequest struct {
	// port unique identifier
	PortId string `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	// connection unique identifier
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// whether the channel is ordered or unordered
	Ordering types.Order `protobuf:"varint,3,opt,name=ordering,proto3,enum=ibc.core.channel.v1.Order" json:"ordering,omitempty"`
	// counterparty channel end
	Counterparty *types.Counterparty `protobuf:"bytes,4,opt,name=counterparty,proto3" json:"counterparty,omitempty"`
	// proposed version
	ProposedVersion string `protobuf:"bytes,5,opt,name=proposed_version,json=proposedVersion,proto3" json:"proposed_version,omitempty"`
}

func (m *QueryAppVersionRequest) Reset()         { *m = QueryAppVersionRequest{} }
func (m *QueryAppVersionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAppVersionRequest) ProtoMessage()    {}
func (*QueryAppVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a256596009a8334, []int{0}
}
func (m *QueryAppVersionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAppVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAppVersionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAppVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAppVersionRequest.Merge(m, src)
}
func (m *QueryAppVersionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAppVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAppVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAppVersionRequest proto.InternalMessageInfo

func (m *QueryAppVersionRequest) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *QueryAppVersionRequest) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *QueryAppVersionRequest) GetOrdering() types.Order {
	if m != nil {
		return m.Ordering
	}
	return types.NONE
}

func (m *QueryAppVersionRequest) GetCounterparty() *types.Counterparty {
	if m != nil {
		return m.Counterparty
	}
	return nil
}

func (m *QueryAppVersionRequest) GetProposedVersion() string {
	if m != nil {
		return m.ProposedVersion
	}
	return ""
}

// QueryAppVersionResponse is the response type for the Query/AppVersion RPC method.
type QueryAppVersionResponse struct {
	// port id associated with the request identifiers
	PortId string `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	// supported app version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *QueryAppVersionResponse) Reset()         { *m = QueryAppVersionResponse{} }
func (m *QueryAppVersionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAppVersionResponse) ProtoMessage()    {}
func (*QueryAppVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a256596009a8334, []int{1}
}
func (m *QueryAppVersionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAppVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAppVersionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAppVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAppVersionResponse.Merge(m, src)
}
func (m *QueryAppVersionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAppVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAppVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAppVersionResponse proto.InternalMessageInfo

func (m *QueryAppVersionResponse) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *QueryAppVersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryAppVersionRequest)(nil), "ibc.core.port.v1.QueryAppVersionRequest")
	proto.RegisterType((*QueryAppVersionResponse)(nil), "ibc.core.port.v1.QueryAppVersionResponse")
}

func init() { proto.RegisterFile("ibc/core/port/v1/query.proto", fileDescriptor_9a256596009a8334) }

var fileDescriptor_9a256596009a8334 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3f, 0x8f, 0xda, 0x30,
	0x14, 0x4f, 0x68, 0x81, 0xd6, 0xa5, 0x2d, 0xf2, 0x50, 0x22, 0x54, 0x45, 0x40, 0x97, 0x30, 0xe0,
	0x14, 0x50, 0xbb, 0xb7, 0x55, 0x07, 0xa4, 0x56, 0x55, 0x33, 0x74, 0xe8, 0x82, 0x88, 0x63, 0x81,
	0x25, 0xf0, 0x33, 0xb6, 0x13, 0x89, 0xed, 0x3e, 0xc2, 0x7d, 0xac, 0x1b, 0x19, 0x6f, 0x3c, 0xc1,
	0x17, 0x39, 0x39, 0x21, 0x70, 0x7f, 0x38, 0xe9, 0x36, 0xbf, 0xf7, 0xfb, 0xa3, 0xdf, 0x7b, 0x7e,
	0xe8, 0x23, 0x8f, 0x69, 0x48, 0x41, 0xb1, 0x50, 0x82, 0x32, 0x61, 0x36, 0x0c, 0xd7, 0x29, 0x53,
	0x1b, 0x22, 0x15, 0x18, 0xc0, 0x4d, 0x1e, 0x53, 0x62, 0x51, 0x62, 0x51, 0x92, 0x0d, 0xdb, 0xdd,
	0x23, 0x9f, 0x2e, 0x66, 0x42, 0xb0, 0xa5, 0x95, 0x1c, 0x9e, 0x85, 0xa8, 0x77, 0x51, 0x41, 0x1f,
	0xfe, 0x5a, 0x93, 0x6f, 0x52, 0xfe, 0x63, 0x4a, 0x73, 0x10, 0x11, 0x5b, 0xa7, 0x4c, 0x1b, 0xdc,
	0x42, 0x75, 0x6b, 0x34, 0xe5, 0x89, 0xe7, 0x76, 0xdc, 0xe0, 0x75, 0x54, 0xb3, 0xe5, 0x24, 0xc1,
	0x9f, 0xd0, 0x5b, 0x0a, 0x42, 0x30, 0x6a, 0x38, 0x08, 0x0b, 0x57, 0x72, 0xb8, 0x71, 0x6a, 0x4e,
	0x12, 0xfc, 0x15, 0xbd, 0x02, 0x95, 0x30, 0xc5, 0xc5, 0xdc, 0x7b, 0xd1, 0x71, 0x83, 0x77, 0xa3,
	0x36, 0x39, 0x06, 0x2c, 0x33, 0x64, 0x43, 0xf2, 0xc7, 0x92, 0xa2, 0x23, 0x17, 0xff, 0x44, 0x0d,
	0x0a, 0xa9, 0x30, 0x4c, 0xc9, 0x99, 0x32, 0x1b, 0xef, 0x65, 0xc7, 0x0d, 0xde, 0x8c, 0xba, 0x67,
	0xb5, 0x3f, 0xee, 0x10, 0xa3, 0x7b, 0x32, 0xdc, 0x47, 0x4d, 0xa9, 0x40, 0x82, 0x66, 0xc9, 0x34,
	0x2b, 0xe6, 0xf2, 0xaa, 0x79, 0xcc, 0xf7, 0x65, 0xff, 0x30, 0x6e, 0xef, 0x17, 0x6a, 0x3d, 0xda,
	0x80, 0x96, 0x20, 0x34, 0x7b, 0x7a, 0x05, 0x1e, 0xaa, 0x97, 0xae, 0xc5, 0xf0, 0x65, 0x39, 0x5a,
	0xa2, 0x6a, 0xee, 0x86, 0x29, 0x42, 0x27, 0x47, 0x1c, 0x90, 0x87, 0xbf, 0x43, 0xce, 0xaf, 0xbd,
	0xdd, 0x7f, 0x06, 0xb3, 0x88, 0xd7, 0x73, 0xbe, 0xff, 0xbe, 0xda, 0xf9, 0xee, 0x76, 0xe7, 0xbb,
	0x37, 0x3b, 0xdf, 0xbd, 0xdc, 0xfb, 0xce, 0x76, 0xef, 0x3b, 0xd7, 0x7b, 0xdf, 0xf9, 0x3f, 0x9e,
	0x73, 0xb3, 0x48, 0x63, 0x42, 0x61, 0x15, 0x52, 0xd0, 0x2b, 0xd0, 0x21, 0x8f, 0xe9, 0x60, 0x0e,
	0xe1, 0x0a, 0x92, 0x74, 0xc9, 0x74, 0x71, 0x18, 0x9f, 0xbf, 0x0c, 0xf2, 0x5b, 0x32, 0x1b, 0xc9,
	0x74, 0x5c, 0xcb, 0x8f, 0x62, 0x7c, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x19, 0xf0, 0xfb, 0x67, 0x69,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// AppVersion queries an IBC Port and determines the appropriate application version to be used
	AppVersion(ctx context.Context, in *QueryAppVersionRequest, opts ...grpc.CallOption) (*QueryAppVersionResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) AppVersion(ctx context.Context, in *QueryAppVersionRequest, opts ...grpc.CallOption) (*QueryAppVersionResponse, error) {
	out := new(QueryAppVersionResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.port.v1.Query/AppVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// AppVersion queries an IBC Port and determines the appropriate application version to be used
	AppVersion(context.Context, *QueryAppVersionRequest) (*QueryAppVersionResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) AppVersion(ctx context.Context, req *QueryAppVersionRequest) (*QueryAppVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppVersion not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_AppVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAppVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AppVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.port.v1.Query/AppVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AppVersion(ctx, req.(*QueryAppVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.core.port.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppVersion",
			Handler:    _Query_AppVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/core/port/v1/query.proto",
}

func (m *QueryAppVersionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAppVersionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAppVersionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProposedVersion) > 0 {
		i -= len(m.ProposedVersion)
		copy(dAtA[i:], m.ProposedVersion)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ProposedVersion)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Counterparty != nil {
		{
			size, err := m.Counterparty.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Ordering != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Ordering))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAppVersionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAppVersionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAppVersionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAppVersionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Ordering != 0 {
		n += 1 + sovQuery(uint64(m.Ordering))
	}
	if m.Counterparty != nil {
		l = m.Counterparty.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ProposedVersion)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAppVersionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAppVersionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAppVersionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAppVersionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ordering", wireType)
			}
			m.Ordering = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ordering |= types.Order(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counterparty", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Counterparty == nil {
				m.Counterparty = &types.Counterparty{}
			}
			if err := m.Counterparty.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposedVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposedVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAppVersionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAppVersionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAppVersionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
