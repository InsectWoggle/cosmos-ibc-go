// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/lightclients/localhost/v1/localhost.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// ClientState defines a loopback (localhost) client. It requires (read-only)
// access to keys outside the client prefix.
type ClientState struct {
	// self chain ID
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	// self latest block height
	Height types.Height `protobuf:"bytes,2,opt,name=height,proto3" json:"height"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_acd9f5b22d41bf6d, []int{0}
}
func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return m.Size()
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientState)(nil), "ibc.lightclients.localhost.v1.ClientState")
}

func init() {
	proto.RegisterFile("ibc/lightclients/localhost/v1/localhost.proto", fileDescriptor_acd9f5b22d41bf6d)
}

var fileDescriptor_acd9f5b22d41bf6d = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcd, 0x4c, 0x4a, 0xd6,
	0xcf, 0xc9, 0x4c, 0xcf, 0x28, 0x49, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x29, 0xd6, 0xcf, 0xc9, 0x4f,
	0x4e, 0xcc, 0xc9, 0xc8, 0x2f, 0x2e, 0xd1, 0x2f, 0x33, 0x44, 0x70, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b,
	0xf2, 0x85, 0x64, 0x33, 0x93, 0x92, 0xf5, 0x90, 0x95, 0xeb, 0x21, 0x54, 0x94, 0x19, 0x4a, 0x89,
	0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x55, 0xea, 0x83, 0x58, 0x10, 0x4d, 0x52, 0xf2, 0x20, 0x3b, 0x92,
	0xf3, 0x8b, 0x52, 0xf5, 0x21, 0x9a, 0x40, 0x06, 0x43, 0x58, 0x10, 0x05, 0x4a, 0xb5, 0x5c, 0xdc,
	0xce, 0x60, 0x7e, 0x70, 0x49, 0x62, 0x49, 0xaa, 0x90, 0x1e, 0x17, 0x47, 0x72, 0x46, 0x62, 0x66,
	0x5e, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x93, 0xf0, 0xa7, 0x7b, 0xf2, 0xfc,
	0x95, 0x89, 0xb9, 0x39, 0x56, 0x4a, 0x30, 0x19, 0xa5, 0x20, 0x76, 0x30, 0xd3, 0x33, 0x45, 0xc8,
	0x82, 0x8b, 0x2d, 0x23, 0x15, 0xe4, 0x26, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x29, 0x3d,
	0x90, 0x2b, 0x41, 0x16, 0xea, 0x41, 0xad, 0x29, 0x33, 0xd4, 0xf3, 0x00, 0xab, 0x70, 0x62, 0x39,
	0x71, 0x4f, 0x9e, 0x21, 0x08, 0xaa, 0xde, 0x8a, 0xa5, 0x63, 0x81, 0x3c, 0x83, 0x53, 0xf4, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x39, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0xeb, 0x67, 0x26, 0x25, 0xeb,
	0xa6, 0xe7, 0xeb, 0xe7, 0xe6, 0xa7, 0x94, 0xe6, 0xa4, 0x16, 0x43, 0x82, 0x4e, 0x17, 0x16, 0x76,
	0x06, 0x96, 0xba, 0x88, 0xe0, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x7b, 0xd1, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xee, 0xdf, 0xcc, 0xab, 0x69, 0x01, 0x00, 0x00,
}

func (m *ClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Height.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLocalhost(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintLocalhost(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLocalhost(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocalhost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovLocalhost(uint64(l))
	}
	l = m.Height.Size()
	n += 1 + l + sovLocalhost(uint64(l))
	return n
}

func sovLocalhost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocalhost(x uint64) (n int) {
	return sovLocalhost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalhost
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
			return fmt.Errorf("proto: ClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalhost
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
				return ErrInvalidLengthLocalhost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalhost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalhost
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
				return ErrInvalidLengthLocalhost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLocalhost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Height.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalhost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLocalhost
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
func skipLocalhost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocalhost
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
					return 0, ErrIntOverflowLocalhost
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
					return 0, ErrIntOverflowLocalhost
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
				return 0, ErrInvalidLengthLocalhost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLocalhost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLocalhost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLocalhost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocalhost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLocalhost = fmt.Errorf("proto: unexpected end of group")
)
