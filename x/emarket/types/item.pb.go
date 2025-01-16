// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emarket/emarket/item.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Item struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProductType string `protobuf:"bytes,3,opt,name=productType,proto3" json:"productType,omitempty"`
	Amount      int32  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Price       int32  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Discounted  bool   `protobuf:"varint,6,opt,name=discounted,proto3" json:"discounted,omitempty"`
	Creator     string `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_561579aa25b97567, []int{0}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Item.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return m.Size()
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetProductType() string {
	if m != nil {
		return m.ProductType
	}
	return ""
}

func (m *Item) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Item) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetDiscounted() bool {
	if m != nil {
		return m.Discounted
	}
	return false
}

func (m *Item) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Item)(nil), "emarket.emarket.Item")
}

func init() { proto.RegisterFile("emarket/emarket/item.proto", fileDescriptor_561579aa25b97567) }

var fileDescriptor_561579aa25b97567 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x3b, 0xb5, 0xed, 0xea, 0x08, 0x0a, 0x83, 0x68, 0xf0, 0x10, 0x82, 0xa7, 0x9e, 0x56,
	0xc4, 0x37, 0xf0, 0xe6, 0xb5, 0x78, 0xf2, 0x56, 0x93, 0x39, 0x04, 0xc9, 0x26, 0x64, 0xb3, 0xe0,
	0xbe, 0x85, 0x0f, 0xe3, 0x43, 0x78, 0xdc, 0xa3, 0x47, 0x69, 0x5f, 0x44, 0x8c, 0xad, 0xf4, 0xf4,
	0xe7, 0xfb, 0x32, 0x30, 0xfc, 0x83, 0xd7, 0xec, 0xfa, 0xf8, 0xca, 0xe9, 0x76, 0x4e, 0x9b, 0xd8,
	0xad, 0x43, 0xf4, 0xc9, 0xd3, 0xf9, 0xe4, 0xd6, 0x53, 0xde, 0x7c, 0x00, 0x56, 0x8f, 0x89, 0x1d,
	0x9d, 0x61, 0x69, 0x8d, 0x00, 0x05, 0x6d, 0xd5, 0x95, 0xd6, 0x10, 0x61, 0xb5, 0xe9, 0x1d, 0x8b,
	0x52, 0x41, 0x7b, 0xd2, 0xe5, 0x37, 0x29, 0x3c, 0x0d, 0xd1, 0x9b, 0x9d, 0x4e, 0x4f, 0xfb, 0xc0,
	0xe2, 0x28, 0x7f, 0x2d, 0x15, 0x5d, 0x62, 0xd3, 0x3b, 0xbf, 0xdb, 0x24, 0x51, 0x29, 0x68, 0xeb,
	0x6e, 0x22, 0xba, 0xc0, 0x3a, 0x44, 0xab, 0x59, 0xd4, 0x59, 0xff, 0x01, 0x49, 0x44, 0x63, 0xb7,
	0xfa, 0x77, 0x82, 0x8d, 0x68, 0x14, 0xb4, 0xc7, 0xdd, 0xc2, 0x90, 0xc0, 0x95, 0x8e, 0xdc, 0x27,
	0x1f, 0xc5, 0x2a, 0xef, 0x9a, 0xf1, 0xe1, 0xee, 0x73, 0x90, 0x70, 0x18, 0x24, 0x7c, 0x0f, 0x12,
	0xde, 0x47, 0x59, 0x1c, 0x46, 0x59, 0x7c, 0x8d, 0xb2, 0x78, 0xbe, 0x9a, 0x5b, 0xbf, 0xfd, 0xf7,
	0x4f, 0xfb, 0xc0, 0xdb, 0x97, 0x26, 0x5f, 0xe0, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x27, 0x16,
	0xdb, 0x80, 0x1f, 0x01, 0x00, 0x00,
}

func (m *Item) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Item) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Item) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Discounted {
		i--
		if m.Discounted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Price != 0 {
		i = encodeVarintItem(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x28
	}
	if m.Amount != 0 {
		i = encodeVarintItem(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ProductType) > 0 {
		i -= len(m.ProductType)
		copy(dAtA[i:], m.ProductType)
		i = encodeVarintItem(dAtA, i, uint64(len(m.ProductType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintItem(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintItem(dAtA []byte, offset int, v uint64) int {
	offset -= sovItem(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Item) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovItem(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.ProductType)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovItem(uint64(m.Amount))
	}
	if m.Price != 0 {
		n += 1 + sovItem(uint64(m.Price))
	}
	if m.Discounted {
		n += 2
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	return n
}

func sovItem(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozItem(x uint64) (n int) {
	return sovItem(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Item) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: Item: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Item: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProductType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Discounted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
			m.Discounted = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthItem
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
func skipItem(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowItem
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
					return 0, ErrIntOverflowItem
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
					return 0, ErrIntOverflowItem
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
				return 0, ErrInvalidLengthItem
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupItem
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthItem
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthItem        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowItem          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupItem = fmt.Errorf("proto: unexpected end of group")
)