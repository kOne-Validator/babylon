// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/zoneconcierge/zoneconcierge.proto

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

// IndexedHeader is the metadata of a CZ header
type IndexedHeader struct {
	// chain_id is the unique ID of the chain
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// hash is the hash of this header
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// height is the height of this header on CZ ledger
	// (hash, height) jointly provides the position of the header on CZ ledger
	Height uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	// babylon_block_height is the height of the Babylon block that includes this header
	BabylonBlockHeight uint64 `protobuf:"varint,4,opt,name=babylon_block_height,json=babylonBlockHeight,proto3" json:"babylon_block_height,omitempty"`
	// epoch is the epoch number of this header on Babylon ledger
	BabylonEpoch uint64 `protobuf:"varint,5,opt,name=babylon_epoch,json=babylonEpoch,proto3" json:"babylon_epoch,omitempty"`
	// babylon_tx_hash is the hash of the tx that includes this header
	// (babylon_block_height, babylon_tx_hash) jointly provides the position of the header on Babylon ledger
	BabylonTxHash []byte `protobuf:"bytes,6,opt,name=babylon_tx_hash,json=babylonTxHash,proto3" json:"babylon_tx_hash,omitempty"`
}

func (m *IndexedHeader) Reset()         { *m = IndexedHeader{} }
func (m *IndexedHeader) String() string { return proto.CompactTextString(m) }
func (*IndexedHeader) ProtoMessage()    {}
func (*IndexedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76d28ce8dde4532, []int{0}
}
func (m *IndexedHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexedHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexedHeader.Merge(m, src)
}
func (m *IndexedHeader) XXX_Size() int {
	return m.Size()
}
func (m *IndexedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_IndexedHeader proto.InternalMessageInfo

func (m *IndexedHeader) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *IndexedHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *IndexedHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *IndexedHeader) GetBabylonBlockHeight() uint64 {
	if m != nil {
		return m.BabylonBlockHeight
	}
	return 0
}

func (m *IndexedHeader) GetBabylonEpoch() uint64 {
	if m != nil {
		return m.BabylonEpoch
	}
	return 0
}

func (m *IndexedHeader) GetBabylonTxHash() []byte {
	if m != nil {
		return m.BabylonTxHash
	}
	return nil
}

// Forks is a list of non-canonical `IndexedHeader`s at the same height.
// For example, assuming the following blockchain
// ```
// A <- B <- C <- D <- E
//            \ -- D1
//            \ -- D2
// ```
// Then the fork will be {[D1, D2]} where each item is in struct `IndexedBlock`.
//
// Note that each `IndexedHeader` in the fork should have a valid quorum certificate.
// Such forks exist since Babylon considers CZs might have dishonest majority.
// Also note that the IBC-Go implementation will only consider the first header in a fork valid, since
// the subsequent headers cannot be verified without knowing the validator set in the previous header.
type Forks struct {
	// blocks is the list of non-canonical indexed headers at the same height
	Headers []*IndexedHeader `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (m *Forks) Reset()         { *m = Forks{} }
func (m *Forks) String() string { return proto.CompactTextString(m) }
func (*Forks) ProtoMessage()    {}
func (*Forks) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76d28ce8dde4532, []int{1}
}
func (m *Forks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Forks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Forks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Forks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Forks.Merge(m, src)
}
func (m *Forks) XXX_Size() int {
	return m.Size()
}
func (m *Forks) XXX_DiscardUnknown() {
	xxx_messageInfo_Forks.DiscardUnknown(m)
}

var xxx_messageInfo_Forks proto.InternalMessageInfo

func (m *Forks) GetHeaders() []*IndexedHeader {
	if m != nil {
		return m.Headers
	}
	return nil
}

// ChainInfo is the information of a CZ
type ChainInfo struct {
	// chain_id is the ID of the chain
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// latest_header is the latest header in the canonical chain of CZ
	LatestHeader *IndexedHeader `protobuf:"bytes,2,opt,name=latest_header,json=latestHeader,proto3" json:"latest_header,omitempty"`
	// latest_forks is the latest forks, formed as a series of IndexedHeader (from low to high)
	LatestForks *Forks `protobuf:"bytes,3,opt,name=latest_forks,json=latestForks,proto3" json:"latest_forks,omitempty"`
}

func (m *ChainInfo) Reset()         { *m = ChainInfo{} }
func (m *ChainInfo) String() string { return proto.CompactTextString(m) }
func (*ChainInfo) ProtoMessage()    {}
func (*ChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76d28ce8dde4532, []int{2}
}
func (m *ChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainInfo.Merge(m, src)
}
func (m *ChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *ChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChainInfo proto.InternalMessageInfo

func (m *ChainInfo) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ChainInfo) GetLatestHeader() *IndexedHeader {
	if m != nil {
		return m.LatestHeader
	}
	return nil
}

func (m *ChainInfo) GetLatestForks() *Forks {
	if m != nil {
		return m.LatestForks
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexedHeader)(nil), "babylon.zoneconcierge.v1.IndexedHeader")
	proto.RegisterType((*Forks)(nil), "babylon.zoneconcierge.v1.Forks")
	proto.RegisterType((*ChainInfo)(nil), "babylon.zoneconcierge.v1.ChainInfo")
}

func init() {
	proto.RegisterFile("babylon/zoneconcierge/zoneconcierge.proto", fileDescriptor_c76d28ce8dde4532)
}

var fileDescriptor_c76d28ce8dde4532 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xee, 0xda, 0x3f, 0x3b, 0x6d, 0x11, 0x16, 0x91, 0x78, 0x89, 0xa1, 0x82, 0xc6, 0x4b, 0xaa,
	0x15, 0x1f, 0xc0, 0x8a, 0xd2, 0x8a, 0x20, 0x04, 0x4f, 0x5e, 0x42, 0x7e, 0xb6, 0xdd, 0xd0, 0x9a,
	0x2d, 0xc9, 0x2a, 0xa9, 0x4f, 0xe1, 0xe3, 0xf8, 0x08, 0x1e, 0x7b, 0x11, 0x3c, 0x4a, 0xfb, 0x22,
	0xd2, 0xc9, 0xe6, 0x10, 0xa1, 0x82, 0xb7, 0x7c, 0x3b, 0xdf, 0x37, 0x33, 0xdf, 0x97, 0x81, 0x13,
	0xcf, 0xf5, 0xe6, 0x53, 0x11, 0x75, 0x5f, 0x45, 0xc4, 0x7c, 0x11, 0xf9, 0x21, 0x8b, 0xc7, 0xac,
	0x88, 0xac, 0x59, 0x2c, 0xa4, 0xa0, 0x9a, 0xa2, 0x5a, 0xc5, 0xe2, 0xcb, 0x59, 0xe7, 0x93, 0x40,
	0x7b, 0x18, 0x05, 0x2c, 0x65, 0xc1, 0x80, 0xb9, 0x01, 0x8b, 0xe9, 0x3e, 0x6c, 0xfb, 0xdc, 0x0d,
	0x23, 0x27, 0x0c, 0x34, 0x62, 0x10, 0xb3, 0x61, 0xd7, 0x11, 0x0f, 0x03, 0x4a, 0xa1, 0xc2, 0xdd,
	0x84, 0x6b, 0x5b, 0x06, 0x31, 0x5b, 0x36, 0x7e, 0xd3, 0x3d, 0xa8, 0x71, 0x16, 0x8e, 0xb9, 0xd4,
	0xca, 0x06, 0x31, 0x2b, 0xb6, 0x42, 0xf4, 0x14, 0x76, 0xd5, 0x50, 0xc7, 0x9b, 0x0a, 0x7f, 0xe2,
	0x28, 0x56, 0x05, 0x59, 0x54, 0xd5, 0xfa, 0xeb, 0xd2, 0x20, 0x53, 0x1c, 0x42, 0x3b, 0x57, 0xb0,
	0x99, 0xf0, 0xb9, 0x56, 0x45, 0x6a, 0x4b, 0x3d, 0x5e, 0xaf, 0xdf, 0xe8, 0x11, 0xec, 0xe4, 0x24,
	0x99, 0x3a, 0xb8, 0x4d, 0x0d, 0xb7, 0xc9, 0xb5, 0x0f, 0xe9, 0xc0, 0x4d, 0x78, 0xe7, 0x16, 0xaa,
	0x37, 0x22, 0x9e, 0x24, 0xf4, 0x12, 0xea, 0x1c, 0x8d, 0x25, 0x5a, 0xd9, 0x28, 0x9b, 0xcd, 0xde,
	0xb1, 0xb5, 0x29, 0x0c, 0xab, 0x10, 0x84, 0x9d, 0xeb, 0x3a, 0xef, 0x04, 0x1a, 0x57, 0x18, 0x41,
	0x34, 0x12, 0x7f, 0xe5, 0x73, 0x07, 0xed, 0xa9, 0x2b, 0x59, 0x22, 0x9d, 0x4c, 0x8a, 0x41, 0xfd,
	0x63, 0x62, 0x2b, 0x53, 0xab, 0x1f, 0xd1, 0x07, 0x85, 0x9d, 0xd1, 0xda, 0x09, 0xe6, 0xdb, 0xec,
	0x1d, 0x6c, 0x6e, 0x86, 0x86, 0xed, 0x66, 0x26, 0x42, 0xd0, 0xbf, 0xff, 0x58, 0xea, 0x64, 0xb1,
	0xd4, 0xc9, 0xf7, 0x52, 0x27, 0x6f, 0x2b, 0xbd, 0xb4, 0x58, 0xe9, 0xa5, 0xaf, 0x95, 0x5e, 0x7a,
	0xbc, 0x18, 0x87, 0x92, 0x3f, 0x7b, 0x96, 0x2f, 0x9e, 0xba, 0xaa, 0x23, 0xda, 0xc8, 0x41, 0x37,
	0xfd, 0x75, 0x57, 0x72, 0x3e, 0x63, 0x89, 0x57, 0xc3, 0x83, 0x3a, 0xff, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x48, 0x5f, 0x35, 0x4e, 0x7d, 0x02, 0x00, 0x00,
}

func (m *IndexedHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexedHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexedHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BabylonTxHash) > 0 {
		i -= len(m.BabylonTxHash)
		copy(dAtA[i:], m.BabylonTxHash)
		i = encodeVarintZoneconcierge(dAtA, i, uint64(len(m.BabylonTxHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.BabylonEpoch != 0 {
		i = encodeVarintZoneconcierge(dAtA, i, uint64(m.BabylonEpoch))
		i--
		dAtA[i] = 0x28
	}
	if m.BabylonBlockHeight != 0 {
		i = encodeVarintZoneconcierge(dAtA, i, uint64(m.BabylonBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.Height != 0 {
		i = encodeVarintZoneconcierge(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintZoneconcierge(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintZoneconcierge(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Forks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Forks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Forks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Headers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintZoneconcierge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	return len(dAtA) - i, nil
}

func (m *ChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestForks != nil {
		{
			size, err := m.LatestForks.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintZoneconcierge(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.LatestHeader != nil {
		{
			size, err := m.LatestHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintZoneconcierge(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintZoneconcierge(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintZoneconcierge(dAtA []byte, offset int, v uint64) int {
	offset -= sovZoneconcierge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IndexedHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovZoneconcierge(uint64(m.Height))
	}
	if m.BabylonBlockHeight != 0 {
		n += 1 + sovZoneconcierge(uint64(m.BabylonBlockHeight))
	}
	if m.BabylonEpoch != 0 {
		n += 1 + sovZoneconcierge(uint64(m.BabylonEpoch))
	}
	l = len(m.BabylonTxHash)
	if l > 0 {
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	return n
}

func (m *Forks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovZoneconcierge(uint64(l))
		}
	}
	return n
}

func (m *ChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	if m.LatestHeader != nil {
		l = m.LatestHeader.Size()
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	if m.LatestForks != nil {
		l = m.LatestForks.Size()
		n += 1 + l + sovZoneconcierge(uint64(l))
	}
	return n
}

func sovZoneconcierge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozZoneconcierge(x uint64) (n int) {
	return sovZoneconcierge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexedHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZoneconcierge
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
			return fmt.Errorf("proto: IndexedHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexedHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
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
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
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
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BabylonBlockHeight", wireType)
			}
			m.BabylonBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BabylonBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BabylonEpoch", wireType)
			}
			m.BabylonEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BabylonEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BabylonTxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
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
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BabylonTxHash = append(m.BabylonTxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BabylonTxHash == nil {
				m.BabylonTxHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZoneconcierge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZoneconcierge
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
func (m *Forks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZoneconcierge
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
			return fmt.Errorf("proto: Forks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Forks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
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
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &IndexedHeader{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZoneconcierge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZoneconcierge
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
func (m *ChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZoneconcierge
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
			return fmt.Errorf("proto: ChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
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
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
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
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LatestHeader == nil {
				m.LatestHeader = &IndexedHeader{}
			}
			if err := m.LatestHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestForks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZoneconcierge
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
				return ErrInvalidLengthZoneconcierge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthZoneconcierge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LatestForks == nil {
				m.LatestForks = &Forks{}
			}
			if err := m.LatestForks.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZoneconcierge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZoneconcierge
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
func skipZoneconcierge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowZoneconcierge
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
					return 0, ErrIntOverflowZoneconcierge
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
					return 0, ErrIntOverflowZoneconcierge
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
				return 0, ErrInvalidLengthZoneconcierge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupZoneconcierge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthZoneconcierge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthZoneconcierge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowZoneconcierge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupZoneconcierge = fmt.Errorf("proto: unexpected end of group")
)
