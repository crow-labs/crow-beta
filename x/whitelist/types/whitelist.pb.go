// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: whitelist/whitelist.proto

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

type Producer struct {
	Id      uint64 `protobuf:"varint,1,opt,name=producerId,proto3" json:"producerId,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Status  string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *Producer) Reset()         { *m = Producer{} }
func (m *Producer) String() string { return proto.CompactTextString(m) }
func (*Producer) ProtoMessage()    {}
func (*Producer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2235d805cc8f0fe2, []int{0}
}
func (m *Producer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Producer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Producer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Producer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Producer.Merge(m, src)
}
func (m *Producer) XXX_Size() int {
	return m.Size()
}
func (m *Producer) XXX_DiscardUnknown() {
	xxx_messageInfo_Producer.DiscardUnknown(m)
}

var xxx_messageInfo_Producer proto.InternalMessageInfo

func (m *Producer) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Producer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Producer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Producer) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type User struct {
	Id      uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Status  string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_2235d805cc8f0fe2, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Whitelist struct {
	Id        uint64     `protobuf:"varint,1,opt,name=whitelistId,proto3" json:"whitelistId,omitempty"`
	Governor  string     `protobuf:"bytes,2,opt,name=governor,proto3" json:"governor,omitempty"`
	Users     []User     `protobuf:"bytes,3,rep,name=users,proto3" json:"users"`
	Producers []Producer `protobuf:"bytes,4,rep,name=producers,proto3" json:"producers"`
}

func (m *Whitelist) Reset()         { *m = Whitelist{} }
func (m *Whitelist) String() string { return proto.CompactTextString(m) }
func (*Whitelist) ProtoMessage()    {}
func (*Whitelist) Descriptor() ([]byte, []int) {
	return fileDescriptor_2235d805cc8f0fe2, []int{2}
}
func (m *Whitelist) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Whitelist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Whitelist.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Whitelist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Whitelist.Merge(m, src)
}
func (m *Whitelist) XXX_Size() int {
	return m.Size()
}
func (m *Whitelist) XXX_DiscardUnknown() {
	xxx_messageInfo_Whitelist.DiscardUnknown(m)
}

var xxx_messageInfo_Whitelist proto.InternalMessageInfo

func (m *Whitelist) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Whitelist) GetGovernor() string {
	if m != nil {
		return m.Governor
	}
	return ""
}

func (m *Whitelist) GetUsers() []User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Whitelist) GetProducers() []Producer {
	if m != nil {
		return m.Producers
	}
	return nil
}

func init() {
	proto.RegisterType((*Producer)(nil), "crow.whitelist.Producer")
	proto.RegisterType((*User)(nil), "crow.whitelist.User")
	proto.RegisterType((*Whitelist)(nil), "crow.whitelist.Whitelist")
}

func init() { proto.RegisterFile("whitelist/whitelist.proto", fileDescriptor_2235d805cc8f0fe2) }

var fileDescriptor_2235d805cc8f0fe2 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xcf, 0xc8, 0x2c,
	0x49, 0xcd, 0xc9, 0x2c, 0x2e, 0xd1, 0x87, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xf8,
	0x92, 0x8b, 0xf2, 0xcb, 0xf5, 0xe0, 0xa2, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x29, 0x7d,
	0x10, 0x0b, 0xa2, 0x4a, 0xa9, 0x82, 0x8b, 0x23, 0xa0, 0x28, 0x3f, 0xa5, 0x34, 0x39, 0xb5, 0x48,
	0x48, 0x8d, 0x8b, 0xab, 0x00, 0xca, 0xf6, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x71, 0x62,
	0x7b, 0x74, 0x4f, 0x9e, 0xc9, 0x33, 0x25, 0x08, 0x49, 0x46, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31,
	0x37, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x92, 0xe0, 0x62, 0x4f, 0x4c,
	0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x06, 0x0b, 0xc3, 0xb8, 0x42, 0x62, 0x5c, 0x6c, 0xc5,
	0x25, 0x89, 0x25, 0xa5, 0xc5, 0x12, 0x2c, 0x60, 0x09, 0x28, 0x4f, 0x29, 0x87, 0x8b, 0x25, 0xb4,
	0x38, 0xb5, 0x48, 0x48, 0x8e, 0x8b, 0xad, 0xb4, 0x18, 0x8b, 0x8d, 0x50, 0x51, 0x2a, 0xd9, 0x76,
	0x90, 0x91, 0x8b, 0x33, 0x1c, 0x16, 0x16, 0x42, 0x1a, 0x5c, 0xdc, 0xf0, 0x80, 0xc1, 0xb0, 0x18,
	0x59, 0x4a, 0x48, 0x8a, 0x8b, 0x23, 0x3d, 0xbf, 0x2c, 0xb5, 0x28, 0x2f, 0xbf, 0x08, 0xea, 0x02,
	0x38, 0x5f, 0xc8, 0x80, 0x8b, 0x15, 0xe4, 0x46, 0x90, 0x1b, 0x98, 0x35, 0xb8, 0x8d, 0x44, 0xf4,
	0x50, 0x43, 0x5c, 0x0f, 0xe4, 0x3d, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0x20, 0x0a, 0x85,
	0x6c, 0xb8, 0x38, 0x61, 0xe1, 0x08, 0x72, 0x20, 0x48, 0x97, 0x04, 0xba, 0x2e, 0x58, 0x74, 0x40,
	0x75, 0x22, 0x34, 0x38, 0x19, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94,
	0x18, 0xc8, 0x0c, 0xfd, 0x0a, 0x44, 0x1a, 0xd0, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03,
	0x47, 0xb2, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x91, 0x9b, 0x0e, 0x27, 0x02, 0x00, 0x00,
}

func (m *Producer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Producer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Producer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintWhitelist(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintWhitelist(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWhitelist(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintWhitelist(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *User) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintWhitelist(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintWhitelist(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWhitelist(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintWhitelist(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Whitelist) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Whitelist) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Whitelist) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Producers) > 0 {
		for iNdEx := len(m.Producers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Producers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWhitelist(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Users) > 0 {
		for iNdEx := len(m.Users) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Users[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWhitelist(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Governor) > 0 {
		i -= len(m.Governor)
		copy(dAtA[i:], m.Governor)
		i = encodeVarintWhitelist(dAtA, i, uint64(len(m.Governor)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintWhitelist(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhitelist(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhitelist(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Producer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovWhitelist(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWhitelist(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovWhitelist(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovWhitelist(uint64(l))
	}
	return n
}

func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovWhitelist(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWhitelist(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovWhitelist(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovWhitelist(uint64(l))
	}
	return n
}

func (m *Whitelist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovWhitelist(uint64(m.Id))
	}
	l = len(m.Governor)
	if l > 0 {
		n += 1 + l + sovWhitelist(uint64(l))
	}
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovWhitelist(uint64(l))
		}
	}
	if len(m.Producers) > 0 {
		for _, e := range m.Producers {
			l = e.Size()
			n += 1 + l + sovWhitelist(uint64(l))
		}
	}
	return n
}

func sovWhitelist(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhitelist(x uint64) (n int) {
	return sovWhitelist(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Producer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhitelist
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
			return fmt.Errorf("proto: Producer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Producer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
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
				return ErrInvalidLengthWhitelist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
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
				return ErrInvalidLengthWhitelist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
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
				return ErrInvalidLengthWhitelist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhitelist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhitelist
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
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhitelist
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
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
				return ErrInvalidLengthWhitelist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
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
				return ErrInvalidLengthWhitelist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
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
				return ErrInvalidLengthWhitelist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhitelist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhitelist
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
func (m *Whitelist) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhitelist
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
			return fmt.Errorf("proto: Whitelist: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Whitelist: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Governor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
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
				return ErrInvalidLengthWhitelist
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Governor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
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
				return ErrInvalidLengthWhitelist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, User{})
			if err := m.Users[len(m.Users)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Producers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhitelist
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
				return ErrInvalidLengthWhitelist
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhitelist
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Producers = append(m.Producers, Producer{})
			if err := m.Producers[len(m.Producers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhitelist(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhitelist
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
func skipWhitelist(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhitelist
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
					return 0, ErrIntOverflowWhitelist
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
					return 0, ErrIntOverflowWhitelist
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
				return 0, ErrInvalidLengthWhitelist
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhitelist
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhitelist
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhitelist        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhitelist          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhitelist = fmt.Errorf("proto: unexpected end of group")
)
