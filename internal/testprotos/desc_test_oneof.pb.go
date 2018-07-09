// Code generated by protoc-gen-go. DO NOT EDIT.
// source: desc_test_oneof.proto

package testprotos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Value struct {
	// Types that are valid to be assigned to Value:
	//	*Value_BinaryValue
	//	*Value_StringValue
	//	*Value_BooleanValue
	//	*Value_IntValue
	//	*Value_Int64Value
	//	*Value_DoubleValue
	//	*Value_FloatValue
	Value                isValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_desc_test_oneof_2c7a585345c15437, []int{0}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (dst *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(dst, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Value interface {
	isValue_Value()
}

type Value_BinaryValue struct {
	BinaryValue []byte `protobuf:"bytes,1,opt,name=binary_value,json=binaryValue,proto3,oneof"`
}
type Value_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}
type Value_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,3,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}
type Value_IntValue struct {
	IntValue int32 `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3,oneof"`
}
type Value_Int64Value struct {
	Int64Value int64 `protobuf:"varint,5,opt,name=int64_value,json=int64Value,proto3,oneof"`
}
type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,6,opt,name=double_value,json=doubleValue,proto3,oneof"`
}
type Value_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,7,opt,name=float_value,json=floatValue,proto3,oneof"`
}

func (*Value_BinaryValue) isValue_Value()  {}
func (*Value_StringValue) isValue_Value()  {}
func (*Value_BooleanValue) isValue_Value() {}
func (*Value_IntValue) isValue_Value()     {}
func (*Value_Int64Value) isValue_Value()   {}
func (*Value_DoubleValue) isValue_Value()  {}
func (*Value_FloatValue) isValue_Value()   {}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Value) GetBinaryValue() []byte {
	if x, ok := m.GetValue().(*Value_BinaryValue); ok {
		return x.BinaryValue
	}
	return nil
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetValue().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBooleanValue() bool {
	if x, ok := m.GetValue().(*Value_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (m *Value) GetIntValue() int32 {
	if x, ok := m.GetValue().(*Value_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *Value) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*Value_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Value) GetFloatValue() float32 {
	if x, ok := m.GetValue().(*Value_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_BinaryValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BooleanValue)(nil),
		(*Value_IntValue)(nil),
		(*Value_Int64Value)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_FloatValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_BinaryValue:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BinaryValue)
	case *Value_StringValue:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Value_BooleanValue:
		t := uint64(0)
		if x.BooleanValue {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_IntValue:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntValue))
	case *Value_Int64Value:
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64Value))
	case *Value_DoubleValue:
		b.EncodeVarint(6<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *Value_FloatValue:
		b.EncodeVarint(7<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FloatValue)))
	case nil:
	default:
		return fmt.Errorf("Value.Value has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // value.binary_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &Value_BinaryValue{x}
		return true, err
	case 2: // value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Value_StringValue{x}
		return true, err
	case 3: // value.boolean_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_BooleanValue{x != 0}
		return true, err
	case 4: // value.int_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_IntValue{int32(x)}
		return true, err
	case 5: // value.int64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Value_Int64Value{int64(x)}
		return true, err
	case 6: // value.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &Value_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 7: // value.float_value
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &Value_FloatValue{math.Float32frombits(uint32(x))}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// value
	switch x := m.Value.(type) {
	case *Value_BinaryValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.BinaryValue)))
		n += len(x.BinaryValue)
	case *Value_StringValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Value_BooleanValue:
		n += 1 // tag and wire
		n += 1
	case *Value_IntValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.IntValue))
	case *Value_Int64Value:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Int64Value))
	case *Value_DoubleValue:
		n += 1 // tag and wire
		n += 8
	case *Value_FloatValue:
		n += 1 // tag and wire
		n += 4
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Value)(nil), "testprotos.Value")
}

func init() {
	proto.RegisterFile("desc_test_oneof.proto", fileDescriptor_desc_test_oneof_2c7a585345c15437)
}

var fileDescriptor_desc_test_oneof_2c7a585345c15437 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xd0, 0xbd, 0x4a, 0x04, 0x31,
	0x10, 0xc0, 0xf1, 0xcc, 0x9e, 0xb9, 0x8f, 0xb9, 0xb5, 0x39, 0x10, 0x6c, 0x84, 0xa8, 0x08, 0xa9,
	0x6c, 0x14, 0x1f, 0xc0, 0x2a, 0x75, 0x0a, 0xdb, 0x25, 0xf1, 0x72, 0x12, 0x08, 0x19, 0xd9, 0xe4,
	0x04, 0x9f, 0xc1, 0x97, 0x96, 0x7c, 0xdc, 0x76, 0xbb, 0xff, 0xf9, 0x31, 0x3b, 0x2c, 0xde, 0x1c,
	0x5d, 0xfa, 0x9c, 0xb2, 0x4b, 0x79, 0xa2, 0xe8, 0xe8, 0xf4, 0xfc, 0x3d, 0x53, 0xa6, 0x03, 0x96,
	0x52, 0x1f, 0xd3, 0xc3, 0xdf, 0x80, 0xfc, 0xc3, 0x84, 0xb3, 0x3b, 0x3c, 0xe2, 0x68, 0x7d, 0x34,
	0xf3, 0xef, 0xf4, 0x53, 0xde, 0x6f, 0x41, 0x80, 0x1c, 0x15, 0xd3, 0xfb, 0x56, 0x17, 0x94, 0xf2,
	0xec, 0xe3, 0x57, 0x47, 0x83, 0x00, 0xb9, 0x2b, 0xa8, 0xd5, 0x86, 0x9e, 0xf0, 0xda, 0x12, 0x05,
	0x67, 0x62, 0x57, 0x2b, 0x01, 0x72, 0xab, 0x98, 0x1e, 0x7b, 0x6e, 0xec, 0x0e, 0x77, 0x3e, 0xe6,
	0x4e, 0xae, 0x04, 0x48, 0xae, 0x98, 0xde, 0xfa, 0x98, 0xdb, 0xf8, 0x1e, 0xf7, 0x3e, 0xe6, 0xb7,
	0xd7, 0x0e, 0xb8, 0x00, 0xb9, 0x52, 0x4c, 0x63, 0x8d, 0xcb, 0x35, 0x47, 0x3a, 0xdb, 0xe0, 0xba,
	0x59, 0x0b, 0x90, 0x50, 0xae, 0x69, 0x75, 0xd9, 0x73, 0x0a, 0x64, 0x2e, 0x1f, 0xda, 0x08, 0x90,
	0x43, 0xd9, 0x53, 0x63, 0x25, 0xef, 0x1b, 0xe4, 0x75, 0x68, 0xd7, 0xf5, 0xaf, 0xbc, 0xfc, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x91, 0xfe, 0x92, 0x1b, 0x39, 0x01, 0x00, 0x00,
}
