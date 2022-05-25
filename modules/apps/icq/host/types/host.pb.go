// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/icq/host/v1/host.proto

package types

import (
	fmt "fmt"
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

// Params defines the set of on-chain interchain query parameters.
// The following parameters may be used to disable the host submodule.
type Params struct {
	// host_enabled enables or disables the host submodule.
	HostEnabled bool `protobuf:"varint,1,opt,name=host_enabled,json=hostEnabled,proto3" json:"host_enabled,omitempty" yaml:"host_enabled"`
	// allow_height determines whether querying a specific height is allowed.
	AllowHeight bool `protobuf:"varint,2,opt,name=allow_height,json=allowHeight,proto3" json:"allow_height,omitempty" yaml:"allow_height"`
	// allow_proof determines whether requesting proof of the query is allowed.
	AllowProof bool `protobuf:"varint,3,opt,name=allow_proof,json=allowProof,proto3" json:"allow_proof,omitempty" yaml:"allow_height"`
	// allow_queries defines a list of query paths allowed to be queried on a host chain.
	AllowQueries []string `protobuf:"bytes,4,rep,name=allow_queries,json=allowQueries,proto3" json:"allow_queries,omitempty" yaml:"allow_queries"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_82640ca3c7c03171, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetHostEnabled() bool {
	if m != nil {
		return m.HostEnabled
	}
	return false
}

func (m *Params) GetAllowHeight() bool {
	if m != nil {
		return m.AllowHeight
	}
	return false
}

func (m *Params) GetAllowProof() bool {
	if m != nil {
		return m.AllowProof
	}
	return false
}

func (m *Params) GetAllowQueries() []string {
	if m != nil {
		return m.AllowQueries
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "ibc.applications.icq.host.v1.Params")
}

func init() {
	proto.RegisterFile("ibc/applications/icq/host/v1/host.proto", fileDescriptor_82640ca3c7c03171)
}

var fileDescriptor_82640ca3c7c03171 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x4b, 0x33, 0x31,
	0x1c, 0xc6, 0x7b, 0x6f, 0x5f, 0x8a, 0xa6, 0xba, 0x9c, 0x05, 0x0f, 0x91, 0xb4, 0xdc, 0x62, 0x17,
	0x2f, 0x94, 0x0e, 0x4a, 0xc1, 0xa5, 0x20, 0x38, 0xb6, 0x1d, 0x5d, 0x4a, 0x92, 0xc6, 0xbb, 0x40,
	0xae, 0xff, 0xeb, 0x25, 0xad, 0xf4, 0x5b, 0xf8, 0xb1, 0x1c, 0x3b, 0x3a, 0x15, 0xe9, 0x7d, 0x83,
	0x6e, 0x6e, 0x92, 0xc4, 0xe1, 0x70, 0x70, 0x4a, 0x7e, 0x3c, 0xcf, 0x0f, 0x42, 0x1e, 0x74, 0x23,
	0x19, 0x27, 0xb4, 0x28, 0x94, 0xe4, 0xd4, 0x48, 0x58, 0x6a, 0x22, 0xf9, 0x8a, 0x64, 0xa0, 0x0d,
	0xd9, 0x0c, 0xdc, 0x99, 0x14, 0x25, 0x18, 0x08, 0xaf, 0x25, 0xe3, 0x49, 0xbd, 0x98, 0x48, 0xbe,
	0x4a, 0x5c, 0x61, 0x33, 0xb8, 0xea, 0xa4, 0x90, 0x82, 0x2b, 0x12, 0x7b, 0xf3, 0x4e, 0xfc, 0x15,
	0xa0, 0xd6, 0x84, 0x96, 0x34, 0xd7, 0xe1, 0x08, 0x9d, 0xd9, 0xee, 0x5c, 0x2c, 0x29, 0x53, 0x62,
	0x11, 0x05, 0xbd, 0xa0, 0x7f, 0x32, 0xbe, 0x3c, 0xee, 0xbb, 0x17, 0x5b, 0x9a, 0xab, 0x51, 0x5c,
	0x4f, 0xe3, 0x59, 0xdb, 0xe2, 0xa3, 0x27, 0xeb, 0x52, 0xa5, 0xe0, 0x75, 0x9e, 0x09, 0x99, 0x66,
	0x26, 0xfa, 0xf7, 0xdb, 0xad, 0xa7, 0xf1, 0xac, 0xed, 0xf0, 0xc9, 0x51, 0x78, 0x8f, 0x3c, 0xce,
	0x8b, 0x12, 0xe0, 0x25, 0x6a, 0xfe, 0xad, 0x22, 0x87, 0x13, 0x5b, 0x0d, 0x1f, 0xd0, 0xb9, 0x0f,
	0x57, 0x6b, 0x51, 0x4a, 0xa1, 0xa3, 0xff, 0xbd, 0x66, 0xff, 0x74, 0x1c, 0x1d, 0xf7, 0xdd, 0x4e,
	0xdd, 0xfd, 0x89, 0xe3, 0x99, 0x7f, 0xe4, 0xd4, 0xe3, 0x78, 0xfa, 0x7e, 0xc0, 0xc1, 0xee, 0x80,
	0x83, 0xcf, 0x03, 0x0e, 0xde, 0x2a, 0xdc, 0xd8, 0x55, 0xb8, 0xf1, 0x51, 0xe1, 0xc6, 0xf3, 0x5d,
	0x2a, 0x4d, 0xb6, 0x66, 0x09, 0x87, 0x9c, 0x70, 0xd0, 0x39, 0x68, 0x22, 0x19, 0xbf, 0x4d, 0x81,
	0x6c, 0x86, 0x24, 0x87, 0xc5, 0x5a, 0x09, 0x6d, 0x27, 0xa9, 0x4d, 0x61, 0xb6, 0x85, 0xd0, 0xac,
	0xe5, 0x7e, 0x75, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x03, 0x73, 0x5d, 0x7e, 0xb4, 0x01, 0x00,
	0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowQueries) > 0 {
		for iNdEx := len(m.AllowQueries) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowQueries[iNdEx])
			copy(dAtA[i:], m.AllowQueries[iNdEx])
			i = encodeVarintHost(dAtA, i, uint64(len(m.AllowQueries[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.AllowProof {
		i--
		if m.AllowProof {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.AllowHeight {
		i--
		if m.AllowHeight {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.HostEnabled {
		i--
		if m.HostEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHost(dAtA []byte, offset int, v uint64) int {
	offset -= sovHost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HostEnabled {
		n += 2
	}
	if m.AllowHeight {
		n += 2
	}
	if m.AllowProof {
		n += 2
	}
	if len(m.AllowQueries) > 0 {
		for _, s := range m.AllowQueries {
			l = len(s)
			n += 1 + l + sovHost(uint64(l))
		}
	}
	return n
}

func sovHost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHost(x uint64) (n int) {
	return sovHost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
			m.HostEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowHeight", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
			m.AllowHeight = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowProof", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
			m.AllowProof = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowQueries", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowQueries = append(m.AllowQueries, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHost
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
func skipHost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
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
				return 0, ErrInvalidLengthHost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHost = fmt.Errorf("proto: unexpected end of group")
)
