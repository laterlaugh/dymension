// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/sponsorship/sponsorship.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params is a module parameters.
type Params struct {
	// MinAllocationWeight is a minimum portion of the user's voting power that
	// one can allocate to a single gauge. The value is measured in percentages
	// and must fall between 1 and 100 * 10^18, inclusive. The base unit is
	// 10^-18%, so
	//  * 1 --> 10^-18%
	//  * 10^18 --> 1%
	//  * 100 * 10^18 --> 100%.
	//
	// Example! If this parameter is 20 * 10^18, which implies 20%, then the min
	// allocation is 20%, and consequently, one can vote on a max of 5 gauges:
	// 100 / 20 = 5.
	MinAllocationWeight cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=min_allocation_weight,json=minAllocationWeight,proto3,customtype=cosmossdk.io/math.Int" json:"min_allocation_weight"`
	// MinVotingPower is a minimum voting power a user must have in order to be
	// able to vote. Denominated in aDYM.
	MinVotingPower cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=min_voting_power,json=minVotingPower,proto3,customtype=cosmossdk.io/math.Int" json:"min_voting_power"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b03084b45de066, []int{0}
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

// Distribution holds the distribution plan among gauges. Distribution with the
// Merge operation forms an Abelian group:
// https://en.wikipedia.org/wiki/Abelian_group. Which helps to safely operate
// with it. That is, Distribution:
//  1. Is commutative:           a + b = b + a
//  2. Is associative:           a + (b + c) = (a + b) + c
//  3. Has the identity element: e + a = a + e = a
//  4. Has inverse elements:     i + a = a + i = e
//
// where
// a, b, c, i, e : Distribution type,
// + : Merge operation (Merge method)
// i : inverse of a (Negate method),
// e : identity element (zero, NewDistribution method).
//
// Example 1:
//
//	a     : [100, 1] [50, 2] [0, 3] power 100
//	b     : [160, 1] [40, 2] [5, 3] power 200
//	a + b : [260, 1] [90, 2] [5, 3] power 300
//
// Example 2:
//
//	a     : [100, 1] [50, 2] [0, 3] power 100
//	b     : [160, 1]                power 200
//	a + b : [260, 1] [50, 2] [0, 3] power 300
//
// Example 3:
//
//	a     : [100, 1] [50, 2] [0, 3]          power 100
//	b     :                         [160, 4] power 200
//	a + b : [100, 1] [50, 2] [0, 3] [160, 4] power 300
//
// Example 4:
//
//	a     : [210, 1] [180, 2] [210, 3] power 600
//	-b    : [-40, 1] [-10, 2]          power -50
//	a - b : [170, 1] [180, 2] [210, 3] power 550
//
// Example 5:
//
//	a         : [210, 1] [180, 2] [210, 3] power 600
//	e         :                            power 0
//	a + e = a : [210, 1] [180, 2] [210, 3] power 600
//
// Example 6:
//
//	a         : [ 210, 1] [ 180, 2] [ 210, 3] power  600
//	i = -a    : [-210, 1] [-180, 2] [-210, 3] power -600
//	a + i = e :                               power 0
//
// CONTRACT: Gauges are sorted by the gauge ID.
// CONTRACT: Gauges hold gauges only with non-zero power.
type Distribution struct {
	// VotingPower is the total voting power that the distribution holds.
	VotingPower cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=voting_power,json=votingPower,proto3,customtype=cosmossdk.io/math.Int" json:"voting_power"`
	// Gauges is a breakdown of the voting power for different gauges.
	Gauges []Gauge `protobuf:"bytes,2,rep,name=gauges,proto3" json:"gauges"`
}

func (m *Distribution) Reset()         { *m = Distribution{} }
func (m *Distribution) String() string { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()    {}
func (*Distribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b03084b45de066, []int{1}
}
func (m *Distribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Distribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Distribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Distribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Distribution.Merge(m, src)
}
func (m *Distribution) XXX_Size() int {
	return m.Size()
}
func (m *Distribution) XXX_DiscardUnknown() {
	xxx_messageInfo_Distribution.DiscardUnknown(m)
}

var xxx_messageInfo_Distribution proto.InternalMessageInfo

func (m *Distribution) GetGauges() []Gauge {
	if m != nil {
		return m.Gauges
	}
	return nil
}

// Gauge represents a single gauge with its absolute power.
type Gauge struct {
	// GaugeID is the ID of the gauge.
	GaugeId uint64 `protobuf:"varint,1,opt,name=gauge_id,json=gaugeId,proto3" json:"gauge_id,omitempty"`
	// Power is a total voting power distributed to this gauge.
	Power cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=power,proto3,customtype=cosmossdk.io/math.Int" json:"power"`
}

func (m *Gauge) Reset()         { *m = Gauge{} }
func (m *Gauge) String() string { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()    {}
func (*Gauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b03084b45de066, []int{2}
}
func (m *Gauge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Gauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Gauge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Gauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gauge.Merge(m, src)
}
func (m *Gauge) XXX_Size() int {
	return m.Size()
}
func (m *Gauge) XXX_DiscardUnknown() {
	xxx_messageInfo_Gauge.DiscardUnknown(m)
}

var xxx_messageInfo_Gauge proto.InternalMessageInfo

func (m *Gauge) GetGaugeId() uint64 {
	if m != nil {
		return m.GaugeId
	}
	return 0
}

// Vote represents the user's vote.
type Vote struct {
	// Voting power is a total voting power of the vote.
	VotingPower cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=voting_power,json=votingPower,proto3,customtype=cosmossdk.io/math.Int" json:"voting_power"`
	// Weights is a breakdown of the vote for different gauges.
	Weights []GaugeWeight `protobuf:"bytes,2,rep,name=weights,proto3" json:"weights"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b03084b45de066, []int{3}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetWeights() []GaugeWeight {
	if m != nil {
		return m.Weights
	}
	return nil
}

// GaugeWeight is a weight distributed to the specified gauge.
type GaugeWeight struct {
	// GaugeID is the ID of the gauge.
	GaugeId uint64 `protobuf:"varint,1,opt,name=gauge_id,json=gaugeId,proto3" json:"gauge_id,omitempty"`
	// Weight is a portion of the voting power that is allocated for the given
	// gauge. The value is measured in percentages and must fall between 1 and 100
	// * 10^18, inclusive. The base unit is 10^-18%, so
	//   - 1 --> 10^-18%
	//   - 10^18 --> 1%
	//   - 100 * 10^18 --> 100%.
	Weight cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=weight,proto3,customtype=cosmossdk.io/math.Int" json:"weight"`
}

func (m *GaugeWeight) Reset()         { *m = GaugeWeight{} }
func (m *GaugeWeight) String() string { return proto.CompactTextString(m) }
func (*GaugeWeight) ProtoMessage()    {}
func (*GaugeWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2b03084b45de066, []int{4}
}
func (m *GaugeWeight) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GaugeWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GaugeWeight.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GaugeWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GaugeWeight.Merge(m, src)
}
func (m *GaugeWeight) XXX_Size() int {
	return m.Size()
}
func (m *GaugeWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_GaugeWeight.DiscardUnknown(m)
}

var xxx_messageInfo_GaugeWeight proto.InternalMessageInfo

func (m *GaugeWeight) GetGaugeId() uint64 {
	if m != nil {
		return m.GaugeId
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "dymensionxyz.dymension.sponsorship.Params")
	proto.RegisterType((*Distribution)(nil), "dymensionxyz.dymension.sponsorship.Distribution")
	proto.RegisterType((*Gauge)(nil), "dymensionxyz.dymension.sponsorship.Gauge")
	proto.RegisterType((*Vote)(nil), "dymensionxyz.dymension.sponsorship.Vote")
	proto.RegisterType((*GaugeWeight)(nil), "dymensionxyz.dymension.sponsorship.GaugeWeight")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/sponsorship/sponsorship.proto", fileDescriptor_f2b03084b45de066)
}

var fileDescriptor_f2b03084b45de066 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0xce, 0xd2, 0x40,
	0x14, 0x6d, 0x3f, 0xa1, 0xe8, 0x40, 0x8c, 0xa9, 0x92, 0x00, 0x8b, 0x42, 0xba, 0xc2, 0x18, 0xa7,
	0x89, 0xb8, 0x70, 0x0b, 0x9a, 0x10, 0x36, 0x4a, 0x9a, 0x88, 0x89, 0x9b, 0xa6, 0xd0, 0xa6, 0x9d,
	0xc8, 0xcc, 0x6d, 0x3a, 0xc3, 0x9f, 0x4f, 0xe1, 0x5b, 0xf8, 0x02, 0x3e, 0x80, 0x4b, 0x96, 0xc4,
	0x95, 0x71, 0x41, 0x0c, 0xbc, 0x88, 0xe9, 0x4c, 0xc1, 0xb2, 0xf0, 0x8f, 0xb8, 0x9b, 0x33, 0x77,
	0xce, 0x39, 0x37, 0x67, 0xee, 0x45, 0x4f, 0x83, 0x0d, 0x0d, 0x19, 0x27, 0xc0, 0xd6, 0x9b, 0xf7,
	0xce, 0x19, 0x38, 0x3c, 0x01, 0xc6, 0x21, 0xe5, 0x31, 0x49, 0x8a, 0x67, 0x9c, 0xa4, 0x20, 0xc0,
	0xb4, 0x8b, 0x2c, 0x7c, 0x06, 0xb8, 0xf0, 0xb2, 0xf5, 0x20, 0x82, 0x08, 0xe4, 0x73, 0x27, 0x3b,
	0x29, 0x66, 0xab, 0x39, 0x03, 0x4e, 0x81, 0x7b, 0xaa, 0xa0, 0x80, 0x2a, 0xd9, 0x9f, 0x75, 0x64,
	0x8c, 0xfd, 0xd4, 0xa7, 0xdc, 0xf4, 0x50, 0x9d, 0x12, 0xe6, 0xf9, 0xf3, 0x39, 0xcc, 0x7c, 0x41,
	0x80, 0x79, 0xab, 0x90, 0x44, 0xb1, 0x68, 0xe8, 0x1d, 0xbd, 0x7b, 0x67, 0xf0, 0x68, 0xbb, 0x6f,
	0x6b, 0xdf, 0xf6, 0xed, 0xba, 0xe2, 0xf3, 0xe0, 0x1d, 0x26, 0xe0, 0x50, 0x5f, 0xc4, 0x78, 0xc4,
	0xc4, 0x97, 0x4f, 0x8f, 0x51, 0x2e, 0x3c, 0x62, 0xc2, 0xbd, 0x4f, 0x09, 0xeb, 0x9f, 0x85, 0xde,
	0x48, 0x1d, 0xf3, 0x35, 0xba, 0x97, 0x19, 0x2c, 0x41, 0x10, 0x16, 0x79, 0x09, 0xac, 0xc2, 0xb4,
	0x71, 0xf3, 0xef, 0xda, 0x77, 0x29, 0x61, 0x13, 0xa9, 0x31, 0xce, 0x24, 0xec, 0x8f, 0x3a, 0xaa,
	0xbd, 0x20, 0x5c, 0xa4, 0x64, 0xba, 0xc8, 0xdc, 0xcc, 0x97, 0xa8, 0x76, 0xe1, 0x71, 0x45, 0xff,
	0xd5, 0xe5, 0x4f, 0x03, 0x73, 0x88, 0x8c, 0xc8, 0x5f, 0x44, 0x21, 0x6f, 0xdc, 0x74, 0x6e, 0x75,
	0xab, 0x4f, 0x1e, 0xe2, 0x3f, 0xff, 0x04, 0x1e, 0x66, 0x8c, 0x41, 0x29, 0x33, 0x75, 0x73, 0xba,
	0x1d, 0xa2, 0xb2, 0xbc, 0x36, 0x9b, 0xe8, 0xb6, 0xbc, 0xf2, 0x48, 0x20, 0xbb, 0x2b, 0xb9, 0x15,
	0x89, 0x47, 0x81, 0xd9, 0x47, 0xe5, 0xab, 0x93, 0x51, 0xcc, 0x2c, 0x90, 0xd2, 0x04, 0x44, 0xf8,
	0xdf, 0x83, 0x78, 0x85, 0x2a, 0x6a, 0x24, 0x4e, 0x49, 0x38, 0x7f, 0x9d, 0x84, 0x1a, 0x81, 0x3c,
	0x8f, 0x93, 0x8a, 0x4d, 0x51, 0xb5, 0x50, 0xfd, 0x5d, 0x2c, 0xcf, 0x91, 0x91, 0x4f, 0xe3, 0x15,
	0xb9, 0xe4, 0xd4, 0x81, 0xbb, 0x3d, 0x58, 0xfa, 0xee, 0x60, 0xe9, 0xdf, 0x0f, 0x96, 0xfe, 0xe1,
	0x68, 0x69, 0xbb, 0xa3, 0xa5, 0x7d, 0x3d, 0x5a, 0xda, 0xdb, 0x67, 0x11, 0x11, 0xf1, 0x62, 0x8a,
	0x67, 0x40, 0x9d, 0x5f, 0x2c, 0xe7, 0xb2, 0xe7, 0xac, 0x2f, 0x36, 0x54, 0x6c, 0x92, 0x90, 0x4f,
	0x0d, 0xb9, 0x47, 0xbd, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x68, 0x67, 0x68, 0xd4, 0x03,
	0x00, 0x00,
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
	{
		size := m.MinVotingPower.Size()
		i -= size
		if _, err := m.MinVotingPower.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSponsorship(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinAllocationWeight.Size()
		i -= size
		if _, err := m.MinAllocationWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSponsorship(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Distribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Distribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Distribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Gauges) > 0 {
		for iNdEx := len(m.Gauges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Gauges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSponsorship(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.VotingPower.Size()
		i -= size
		if _, err := m.VotingPower.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSponsorship(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Gauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gauge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Gauge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Power.Size()
		i -= size
		if _, err := m.Power.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSponsorship(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.GaugeId != 0 {
		i = encodeVarintSponsorship(dAtA, i, uint64(m.GaugeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Weights) > 0 {
		for iNdEx := len(m.Weights) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Weights[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSponsorship(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.VotingPower.Size()
		i -= size
		if _, err := m.VotingPower.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSponsorship(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GaugeWeight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GaugeWeight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GaugeWeight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSponsorship(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.GaugeId != 0 {
		i = encodeVarintSponsorship(dAtA, i, uint64(m.GaugeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSponsorship(dAtA []byte, offset int, v uint64) int {
	offset -= sovSponsorship(v)
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
	l = m.MinAllocationWeight.Size()
	n += 1 + l + sovSponsorship(uint64(l))
	l = m.MinVotingPower.Size()
	n += 1 + l + sovSponsorship(uint64(l))
	return n
}

func (m *Distribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.VotingPower.Size()
	n += 1 + l + sovSponsorship(uint64(l))
	if len(m.Gauges) > 0 {
		for _, e := range m.Gauges {
			l = e.Size()
			n += 1 + l + sovSponsorship(uint64(l))
		}
	}
	return n
}

func (m *Gauge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GaugeId != 0 {
		n += 1 + sovSponsorship(uint64(m.GaugeId))
	}
	l = m.Power.Size()
	n += 1 + l + sovSponsorship(uint64(l))
	return n
}

func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.VotingPower.Size()
	n += 1 + l + sovSponsorship(uint64(l))
	if len(m.Weights) > 0 {
		for _, e := range m.Weights {
			l = e.Size()
			n += 1 + l + sovSponsorship(uint64(l))
		}
	}
	return n
}

func (m *GaugeWeight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GaugeId != 0 {
		n += 1 + sovSponsorship(uint64(m.GaugeId))
	}
	l = m.Weight.Size()
	n += 1 + l + sovSponsorship(uint64(l))
	return n
}

func sovSponsorship(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSponsorship(x uint64) (n int) {
	return sovSponsorship(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSponsorship
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAllocationWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSponsorship
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
				return ErrInvalidLengthSponsorship
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSponsorship
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinAllocationWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinVotingPower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSponsorship
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
				return ErrInvalidLengthSponsorship
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSponsorship
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinVotingPower.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSponsorship(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSponsorship
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
func (m *Distribution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSponsorship
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
			return fmt.Errorf("proto: Distribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Distribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSponsorship
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
				return ErrInvalidLengthSponsorship
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSponsorship
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VotingPower.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gauges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSponsorship
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
				return ErrInvalidLengthSponsorship
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSponsorship
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gauges = append(m.Gauges, Gauge{})
			if err := m.Gauges[len(m.Gauges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSponsorship(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSponsorship
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
func (m *Gauge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSponsorship
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
			return fmt.Errorf("proto: Gauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeId", wireType)
			}
			m.GaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSponsorship
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSponsorship
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
				return ErrInvalidLengthSponsorship
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSponsorship
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Power.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSponsorship(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSponsorship
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
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSponsorship
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSponsorship
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
				return ErrInvalidLengthSponsorship
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSponsorship
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VotingPower.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weights", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSponsorship
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
				return ErrInvalidLengthSponsorship
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSponsorship
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Weights = append(m.Weights, GaugeWeight{})
			if err := m.Weights[len(m.Weights)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSponsorship(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSponsorship
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
func (m *GaugeWeight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSponsorship
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
			return fmt.Errorf("proto: GaugeWeight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GaugeWeight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeId", wireType)
			}
			m.GaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSponsorship
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSponsorship
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
				return ErrInvalidLengthSponsorship
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSponsorship
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSponsorship(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSponsorship
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
func skipSponsorship(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSponsorship
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
					return 0, ErrIntOverflowSponsorship
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
					return 0, ErrIntOverflowSponsorship
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
				return 0, ErrInvalidLengthSponsorship
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSponsorship
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSponsorship
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSponsorship        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSponsorship          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSponsorship = fmt.Errorf("proto: unexpected end of group")
)
