// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: share/ipldv2/pb/ipldv2pb.proto

package ipldv2pb

import (
	fmt "fmt"
	pb "github.com/celestiaorg/nmt/pb"
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

type AxisType int32

const (
	AxisType_Row AxisType = 0
	AxisType_Col AxisType = 1
)

var AxisType_name = map[int32]string{
	0: "Row",
	1: "Col",
}

var AxisType_value = map[string]int32{
	"Row": 0,
	"Col": 1,
}

func (x AxisType) String() string {
	return proto.EnumName(AxisType_name, int32(x))
}

func (AxisType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb41c3a4f982a271, []int{0}
}

type SampleType int32

const (
	SampleType_Data   SampleType = 0
	SampleType_Parity SampleType = 1
)

var SampleType_name = map[int32]string{
	0: "Data",
	1: "Parity",
}

var SampleType_value = map[string]int32{
	"Data":   0,
	"Parity": 1,
}

func (x SampleType) String() string {
	return proto.EnumName(SampleType_name, int32(x))
}

func (SampleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb41c3a4f982a271, []int{1}
}

type AxisID struct {
	Type   AxisType `protobuf:"varint,1,opt,name=type,proto3,enum=AxisType" json:"type,omitempty"`
	Height uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Index  uint32   `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Hash   []byte   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *AxisID) Reset()         { *m = AxisID{} }
func (m *AxisID) String() string { return proto.CompactTextString(m) }
func (*AxisID) ProtoMessage()    {}
func (*AxisID) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb41c3a4f982a271, []int{0}
}
func (m *AxisID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AxisID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AxisID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AxisID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AxisID.Merge(m, src)
}
func (m *AxisID) XXX_Size() int {
	return m.Size()
}
func (m *AxisID) XXX_DiscardUnknown() {
	xxx_messageInfo_AxisID.DiscardUnknown(m)
}

var xxx_messageInfo_AxisID proto.InternalMessageInfo

func (m *AxisID) GetType() AxisType {
	if m != nil {
		return m.Type
	}
	return AxisType_Row
}

func (m *AxisID) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *AxisID) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AxisID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Axis struct {
	Id       *AxisID  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AxisHalf [][]byte `protobuf:"bytes,2,rep,name=axis_half,json=axisHalf,proto3" json:"axis_half,omitempty"`
}

func (m *Axis) Reset()         { *m = Axis{} }
func (m *Axis) String() string { return proto.CompactTextString(m) }
func (*Axis) ProtoMessage()    {}
func (*Axis) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb41c3a4f982a271, []int{1}
}
func (m *Axis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Axis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Axis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Axis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Axis.Merge(m, src)
}
func (m *Axis) XXX_Size() int {
	return m.Size()
}
func (m *Axis) XXX_DiscardUnknown() {
	xxx_messageInfo_Axis.DiscardUnknown(m)
}

var xxx_messageInfo_Axis proto.InternalMessageInfo

func (m *Axis) GetId() *AxisID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Axis) GetAxisHalf() [][]byte {
	if m != nil {
		return m.AxisHalf
	}
	return nil
}

type SampleID struct {
	AxisId     *AxisID `protobuf:"bytes,1,opt,name=axis_id,json=axisId,proto3" json:"axis_id,omitempty"`
	ShareIndex uint32  `protobuf:"varint,2,opt,name=share_index,json=shareIndex,proto3" json:"share_index,omitempty"`
}

func (m *SampleID) Reset()         { *m = SampleID{} }
func (m *SampleID) String() string { return proto.CompactTextString(m) }
func (*SampleID) ProtoMessage()    {}
func (*SampleID) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb41c3a4f982a271, []int{2}
}
func (m *SampleID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SampleID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SampleID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SampleID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleID.Merge(m, src)
}
func (m *SampleID) XXX_Size() int {
	return m.Size()
}
func (m *SampleID) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleID.DiscardUnknown(m)
}

var xxx_messageInfo_SampleID proto.InternalMessageInfo

func (m *SampleID) GetAxisId() *AxisID {
	if m != nil {
		return m.AxisId
	}
	return nil
}

func (m *SampleID) GetShareIndex() uint32 {
	if m != nil {
		return m.ShareIndex
	}
	return 0
}

type Sample struct {
	Id    *SampleID  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type  SampleType `protobuf:"varint,2,opt,name=type,proto3,enum=SampleType" json:"type,omitempty"`
	Share []byte     `protobuf:"bytes,3,opt,name=share,proto3" json:"share,omitempty"`
	Proof *pb.Proof  `protobuf:"bytes,4,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb41c3a4f982a271, []int{3}
}
func (m *Sample) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(m, src)
}
func (m *Sample) XXX_Size() int {
	return m.Size()
}
func (m *Sample) XXX_DiscardUnknown() {
	xxx_messageInfo_Sample.DiscardUnknown(m)
}

var xxx_messageInfo_Sample proto.InternalMessageInfo

func (m *Sample) GetId() *SampleID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Sample) GetType() SampleType {
	if m != nil {
		return m.Type
	}
	return SampleType_Data
}

func (m *Sample) GetShare() []byte {
	if m != nil {
		return m.Share
	}
	return nil
}

func (m *Sample) GetProof() *pb.Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterEnum("AxisType", AxisType_name, AxisType_value)
	proto.RegisterEnum("SampleType", SampleType_name, SampleType_value)
	proto.RegisterType((*AxisID)(nil), "AxisID")
	proto.RegisterType((*Axis)(nil), "Axis")
	proto.RegisterType((*SampleID)(nil), "SampleID")
	proto.RegisterType((*Sample)(nil), "Sample")
}

func init() { proto.RegisterFile("share/ipldv2/pb/ipldv2pb.proto", fileDescriptor_cb41c3a4f982a271) }

var fileDescriptor_cb41c3a4f982a271 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x4f, 0x6b, 0xab, 0x40,
	0x1c, 0x74, 0x8d, 0x31, 0xe6, 0x67, 0x5e, 0x9e, 0x2c, 0xe1, 0x3d, 0xfb, 0xcf, 0x88, 0x50, 0x90,
	0x1c, 0x0c, 0xd8, 0x6b, 0x2f, 0x6d, 0x73, 0xa8, 0x87, 0x42, 0xd8, 0xf6, 0x1e, 0x56, 0x34, 0x75,
	0xc1, 0xd6, 0x45, 0xa5, 0x4d, 0x8e, 0xfd, 0x06, 0xfd, 0x58, 0x3d, 0xe6, 0xd8, 0x63, 0x49, 0xbe,
	0x48, 0x71, 0xd7, 0xd0, 0x40, 0x6f, 0x33, 0xb3, 0x32, 0x33, 0xbf, 0x11, 0x9c, 0x2a, 0xa3, 0x65,
	0x3a, 0x65, 0x3c, 0x4f, 0x5e, 0xc2, 0x29, 0x8f, 0x5b, 0xc4, 0xe3, 0x80, 0x97, 0x45, 0x5d, 0x1c,
	0x0f, 0x79, 0x3c, 0xe5, 0x65, 0x51, 0x2c, 0x25, 0xf7, 0x18, 0xe8, 0x57, 0x2b, 0x56, 0x45, 0x33,
	0x7c, 0x06, 0x5a, 0xbd, 0xe6, 0xa9, 0x8d, 0x5c, 0xe4, 0x0f, 0xc3, 0x7e, 0xd0, 0xc8, 0x0f, 0x6b,
	0x9e, 0x12, 0x21, 0xe3, 0x7f, 0xa0, 0x67, 0x29, 0x7b, 0xcc, 0x6a, 0x5b, 0x75, 0x91, 0xaf, 0x91,
	0x96, 0xe1, 0x11, 0x74, 0xd9, 0x73, 0x92, 0xae, 0xec, 0x8e, 0x8b, 0xfc, 0x3f, 0x44, 0x12, 0x8c,
	0x41, 0xcb, 0x68, 0x95, 0xd9, 0x9a, 0x8b, 0xfc, 0x01, 0x11, 0xd8, 0xbb, 0x04, 0xad, 0xf1, 0xc4,
	0xff, 0x41, 0x65, 0x89, 0x88, 0x31, 0xc3, 0x5e, 0x20, 0xd3, 0x89, 0xca, 0x12, 0x7c, 0x02, 0x7d,
	0xba, 0x62, 0xd5, 0x22, 0xa3, 0xf9, 0xd2, 0x56, 0xdd, 0x8e, 0x3f, 0x20, 0x46, 0x23, 0xdc, 0xd2,
	0x7c, 0xe9, 0xdd, 0x81, 0x71, 0x4f, 0x9f, 0x78, 0x9e, 0x46, 0x33, 0xec, 0x42, 0x4f, 0x7c, 0xf8,
	0xdb, 0x46, 0x6f, 0xf4, 0x28, 0xc1, 0x63, 0x30, 0xc5, 0x10, 0x0b, 0xd9, 0x4d, 0x15, 0xdd, 0x40,
	0x48, 0x51, 0xa3, 0x78, 0x6f, 0x08, 0x74, 0xe9, 0x87, 0x8f, 0x0e, 0xfa, 0xf4, 0x83, 0x7d, 0x88,
	0x68, 0x34, 0x6e, 0x37, 0x51, 0xc5, 0x26, 0x66, 0xfb, 0x78, 0xb0, 0xca, 0x08, 0xba, 0xc2, 0x54,
	0x5c, 0x3f, 0x20, 0x92, 0xe0, 0x73, 0xe8, 0x8a, 0x8d, 0xc5, 0xf9, 0x66, 0xf8, 0x37, 0x68, 0x17,
	0x8f, 0x83, 0x79, 0x03, 0x88, 0x7c, 0x9d, 0x9c, 0x82, 0xb1, 0x1f, 0x19, 0xf7, 0xa0, 0x43, 0x8a,
	0x57, 0x4b, 0x69, 0xc0, 0x4d, 0x91, 0x5b, 0x68, 0xe2, 0x01, 0xfc, 0xc4, 0x61, 0x03, 0xb4, 0x19,
	0xad, 0xa9, 0xa5, 0x60, 0x00, 0x7d, 0x4e, 0x4b, 0x56, 0xaf, 0x2d, 0x74, 0x6d, 0x7f, 0x6c, 0x1d,
	0xb4, 0xd9, 0x3a, 0xe8, 0x6b, 0xeb, 0xa0, 0xf7, 0x9d, 0xa3, 0x6c, 0x76, 0x8e, 0xf2, 0xb9, 0x73,
	0x94, 0x58, 0x17, 0xbf, 0xf7, 0xe2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xb7, 0xed, 0xa5, 0x10,
	0x02, 0x00, 0x00,
}

func (m *AxisID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AxisID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AxisID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintIpldv2Pb(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x22
	}
	if m.Index != 0 {
		i = encodeVarintIpldv2Pb(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if m.Height != 0 {
		i = encodeVarintIpldv2Pb(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintIpldv2Pb(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Axis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Axis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Axis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AxisHalf) > 0 {
		for iNdEx := len(m.AxisHalf) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AxisHalf[iNdEx])
			copy(dAtA[i:], m.AxisHalf[iNdEx])
			i = encodeVarintIpldv2Pb(dAtA, i, uint64(len(m.AxisHalf[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Id != nil {
		{
			size, err := m.Id.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIpldv2Pb(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SampleID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SampleID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SampleID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ShareIndex != 0 {
		i = encodeVarintIpldv2Pb(dAtA, i, uint64(m.ShareIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.AxisId != nil {
		{
			size, err := m.AxisId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIpldv2Pb(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Sample) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sample) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sample) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIpldv2Pb(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Share) > 0 {
		i -= len(m.Share)
		copy(dAtA[i:], m.Share)
		i = encodeVarintIpldv2Pb(dAtA, i, uint64(len(m.Share)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Type != 0 {
		i = encodeVarintIpldv2Pb(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != nil {
		{
			size, err := m.Id.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIpldv2Pb(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIpldv2Pb(dAtA []byte, offset int, v uint64) int {
	offset -= sovIpldv2Pb(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AxisID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovIpldv2Pb(uint64(m.Type))
	}
	if m.Height != 0 {
		n += 1 + sovIpldv2Pb(uint64(m.Height))
	}
	if m.Index != 0 {
		n += 1 + sovIpldv2Pb(uint64(m.Index))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovIpldv2Pb(uint64(l))
	}
	return n
}

func (m *Axis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovIpldv2Pb(uint64(l))
	}
	if len(m.AxisHalf) > 0 {
		for _, b := range m.AxisHalf {
			l = len(b)
			n += 1 + l + sovIpldv2Pb(uint64(l))
		}
	}
	return n
}

func (m *SampleID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AxisId != nil {
		l = m.AxisId.Size()
		n += 1 + l + sovIpldv2Pb(uint64(l))
	}
	if m.ShareIndex != 0 {
		n += 1 + sovIpldv2Pb(uint64(m.ShareIndex))
	}
	return n
}

func (m *Sample) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != nil {
		l = m.Id.Size()
		n += 1 + l + sovIpldv2Pb(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovIpldv2Pb(uint64(m.Type))
	}
	l = len(m.Share)
	if l > 0 {
		n += 1 + l + sovIpldv2Pb(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovIpldv2Pb(uint64(l))
	}
	return n
}

func sovIpldv2Pb(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIpldv2Pb(x uint64) (n int) {
	return sovIpldv2Pb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AxisID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIpldv2Pb
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
			return fmt.Errorf("proto: AxisID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AxisID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= AxisType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIpldv2Pb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIpldv2Pb
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
func (m *Axis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIpldv2Pb
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
			return fmt.Errorf("proto: Axis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Axis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Id == nil {
				m.Id = &AxisID{}
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AxisHalf", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AxisHalf = append(m.AxisHalf, make([]byte, postIndex-iNdEx))
			copy(m.AxisHalf[len(m.AxisHalf)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIpldv2Pb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIpldv2Pb
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
func (m *SampleID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIpldv2Pb
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
			return fmt.Errorf("proto: SampleID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SampleID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AxisId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AxisId == nil {
				m.AxisId = &AxisID{}
			}
			if err := m.AxisId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareIndex", wireType)
			}
			m.ShareIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ShareIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIpldv2Pb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIpldv2Pb
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
func (m *Sample) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIpldv2Pb
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
			return fmt.Errorf("proto: Sample: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sample: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Id == nil {
				m.Id = &SampleID{}
			}
			if err := m.Id.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= SampleType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Share", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Share = append(m.Share[:0], dAtA[iNdEx:postIndex]...)
			if m.Share == nil {
				m.Share = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpldv2Pb
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
				return ErrInvalidLengthIpldv2Pb
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIpldv2Pb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &pb.Proof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIpldv2Pb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIpldv2Pb
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
func skipIpldv2Pb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIpldv2Pb
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
					return 0, ErrIntOverflowIpldv2Pb
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
					return 0, ErrIntOverflowIpldv2Pb
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
				return 0, ErrInvalidLengthIpldv2Pb
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIpldv2Pb
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIpldv2Pb
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIpldv2Pb        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIpldv2Pb          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIpldv2Pb = fmt.Errorf("proto: unexpected end of group")
)