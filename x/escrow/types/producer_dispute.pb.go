// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: escrow/producer_dispute.proto

package types

import (
	fmt "fmt"
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

type ProducerDispute struct {
	EscrowId    uint64 `protobuf:"varint,1,opt,name=escrowId,proto3" json:"escrowId,omitempty"`
	ProducerId  uint64 `protobuf:"varint,2,opt,name=producerId,proto3" json:"producerId,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *ProducerDispute) Reset()         { *m = ProducerDispute{} }
func (m *ProducerDispute) String() string { return proto.CompactTextString(m) }
func (*ProducerDispute) ProtoMessage()    {}
func (*ProducerDispute) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7b32cef7bf01d92, []int{0}
}
func (m *ProducerDispute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProducerDispute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProducerDispute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProducerDispute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProducerDispute.Merge(m, src)
}
func (m *ProducerDispute) XXX_Size() int {
	return m.Size()
}
func (m *ProducerDispute) XXX_DiscardUnknown() {
	xxx_messageInfo_ProducerDispute.DiscardUnknown(m)
}

var xxx_messageInfo_ProducerDispute proto.InternalMessageInfo

func (m *ProducerDispute) GetEscrowId() uint64 {
	if m != nil {
		return m.EscrowId
	}
	return 0
}

func (m *ProducerDispute) GetProducerId() uint64 {
	if m != nil {
		return m.ProducerId
	}
	return 0
}

func (m *ProducerDispute) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ProducerDispute)(nil), "crow.escrow.ProducerDispute")
}

func init() { proto.RegisterFile("escrow/producer_dispute.proto", fileDescriptor_d7b32cef7bf01d92) }

var fileDescriptor_d7b32cef7bf01d92 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2d, 0x4e, 0x2e,
	0xca, 0x2f, 0xd7, 0x2f, 0x28, 0xca, 0x4f, 0x29, 0x4d, 0x4e, 0x2d, 0x8a, 0x4f, 0xc9, 0x2c, 0x2e,
	0x28, 0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x49, 0xea, 0x41, 0xd4,
	0x28, 0xe5, 0x73, 0xf1, 0x07, 0x40, 0x95, 0xb9, 0x40, 0x54, 0x09, 0x49, 0x71, 0x71, 0x40, 0x24,
	0x3d, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0xe0, 0x7c, 0x21, 0x39, 0x2e, 0x2e, 0x98,
	0xa9, 0x9e, 0x29, 0x12, 0x4c, 0x60, 0x59, 0x24, 0x11, 0x21, 0x05, 0x2e, 0xee, 0x14, 0x90, 0xe2,
	0xcc, 0x82, 0x92, 0xcc, 0xfc, 0x3c, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x64, 0x21, 0x27,
	0xdd, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x12, 0x06, 0x3b, 0xba, 0x42,
	0x1f, 0xea, 0xfa, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x9b, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xfa, 0x23, 0x91, 0xdf, 0xd4, 0x00, 0x00, 0x00,
}

func (m *ProducerDispute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProducerDispute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProducerDispute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProducerDispute(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ProducerId != 0 {
		i = encodeVarintProducerDispute(dAtA, i, uint64(m.ProducerId))
		i--
		dAtA[i] = 0x10
	}
	if m.EscrowId != 0 {
		i = encodeVarintProducerDispute(dAtA, i, uint64(m.EscrowId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProducerDispute(dAtA []byte, offset int, v uint64) int {
	offset -= sovProducerDispute(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProducerDispute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EscrowId != 0 {
		n += 1 + sovProducerDispute(uint64(m.EscrowId))
	}
	if m.ProducerId != 0 {
		n += 1 + sovProducerDispute(uint64(m.ProducerId))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProducerDispute(uint64(l))
	}
	return n
}

func sovProducerDispute(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProducerDispute(x uint64) (n int) {
	return sovProducerDispute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProducerDispute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProducerDispute
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
			return fmt.Errorf("proto: ProducerDispute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProducerDispute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowId", wireType)
			}
			m.EscrowId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerDispute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EscrowId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProducerId", wireType)
			}
			m.ProducerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerDispute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProducerId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProducerDispute
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
				return ErrInvalidLengthProducerDispute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProducerDispute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProducerDispute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProducerDispute
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
func skipProducerDispute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProducerDispute
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
					return 0, ErrIntOverflowProducerDispute
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
					return 0, ErrIntOverflowProducerDispute
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
				return 0, ErrInvalidLengthProducerDispute
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProducerDispute
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProducerDispute
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProducerDispute        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProducerDispute          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProducerDispute = fmt.Errorf("proto: unexpected end of group")
)
