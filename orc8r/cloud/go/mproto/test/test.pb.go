// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// TestDataBasic contains basic field types for testing deterministic encoding.
type TestDataBasic struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestDataBasic) Reset()         { *m = TestDataBasic{} }
func (m *TestDataBasic) String() string { return proto.CompactTextString(m) }
func (*TestDataBasic) ProtoMessage()    {}
func (*TestDataBasic) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *TestDataBasic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestDataBasic.Unmarshal(m, b)
}
func (m *TestDataBasic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestDataBasic.Marshal(b, m, deterministic)
}
func (m *TestDataBasic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestDataBasic.Merge(m, src)
}
func (m *TestDataBasic) XXX_Size() int {
	return xxx_messageInfo_TestDataBasic.Size(m)
}
func (m *TestDataBasic) XXX_DiscardUnknown() {
	xxx_messageInfo_TestDataBasic.DiscardUnknown(m)
}

var xxx_messageInfo_TestDataBasic proto.InternalMessageInfo

func (m *TestDataBasic) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TestDataBasic) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// TestDataCompound contains compound field types for testing deterministic encoding.
type TestDataCompound struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SingleData           *TestDataBasic            `protobuf:"bytes,2,opt,name=single_data,json=singleData,proto3" json:"single_data,omitempty"`
	DataMap              map[string]*TestDataBasic `protobuf:"bytes,3,rep,name=data_map,json=dataMap,proto3" json:"data_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DataSlice            []*TestDataBasic          `protobuf:"bytes,4,rep,name=data_slice,json=dataSlice,proto3" json:"data_slice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TestDataCompound) Reset()         { *m = TestDataCompound{} }
func (m *TestDataCompound) String() string { return proto.CompactTextString(m) }
func (*TestDataCompound) ProtoMessage()    {}
func (*TestDataCompound) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *TestDataCompound) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestDataCompound.Unmarshal(m, b)
}
func (m *TestDataCompound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestDataCompound.Marshal(b, m, deterministic)
}
func (m *TestDataCompound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestDataCompound.Merge(m, src)
}
func (m *TestDataCompound) XXX_Size() int {
	return xxx_messageInfo_TestDataCompound.Size(m)
}
func (m *TestDataCompound) XXX_DiscardUnknown() {
	xxx_messageInfo_TestDataCompound.DiscardUnknown(m)
}

var xxx_messageInfo_TestDataCompound proto.InternalMessageInfo

func (m *TestDataCompound) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TestDataCompound) GetSingleData() *TestDataBasic {
	if m != nil {
		return m.SingleData
	}
	return nil
}

func (m *TestDataCompound) GetDataMap() map[string]*TestDataBasic {
	if m != nil {
		return m.DataMap
	}
	return nil
}

func (m *TestDataCompound) GetDataSlice() []*TestDataBasic {
	if m != nil {
		return m.DataSlice
	}
	return nil
}

func init() {
	proto.RegisterType((*TestDataBasic)(nil), "magma.orc8r.mproto_test.TestDataBasic")
	proto.RegisterType((*TestDataCompound)(nil), "magma.orc8r.mproto_test.TestDataCompound")
	proto.RegisterMapType((map[string]*TestDataBasic)(nil), "magma.orc8r.mproto_test.TestDataCompound.DataMapEntry")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcf, 0x4d, 0x4c, 0xcf, 0x4d, 0xd4, 0xcb, 0x2f,
	0x4a, 0xb6, 0x28, 0xd2, 0xcb, 0x05, 0x8b, 0xc5, 0x83, 0xa4, 0x95, 0xcc, 0xb9, 0x78, 0x43, 0x52,
	0x8b, 0x4b, 0x5c, 0x12, 0x4b, 0x12, 0x9d, 0x12, 0x8b, 0x33, 0x93, 0x85, 0x04, 0xb8, 0x98, 0xb3,
	0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2,
	0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08, 0x47, 0xe9, 0x25, 0x13,
	0x97, 0x00, 0x4c, 0xa7, 0x73, 0x7e, 0x6e, 0x41, 0x7e, 0x69, 0x5e, 0x8a, 0x10, 0x1f, 0x17, 0x53,
	0x66, 0x0a, 0x54, 0x2f, 0x53, 0x66, 0x8a, 0x90, 0x3b, 0x17, 0x77, 0x71, 0x66, 0x5e, 0x7a, 0x4e,
	0x6a, 0x7c, 0x4a, 0x62, 0x49, 0x22, 0xd8, 0x00, 0x6e, 0x23, 0x35, 0x3d, 0x1c, 0x8e, 0xd1, 0x43,
	0x71, 0x49, 0x10, 0x17, 0x44, 0x2b, 0x48, 0x40, 0x28, 0x90, 0x8b, 0x03, 0x64, 0x42, 0x7c, 0x6e,
	0x62, 0x81, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x19, 0x41, 0x53, 0x60, 0xae, 0xd2, 0x03,
	0x71, 0x7c, 0x13, 0x0b, 0x5c, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0xd8, 0x53, 0x20, 0x3c, 0x21, 0x57,
	0x2e, 0x2e, 0xb0, 0x91, 0xc5, 0x39, 0x99, 0xc9, 0xa9, 0x12, 0x2c, 0x60, 0x43, 0x89, 0x75, 0x1a,
	0x27, 0x48, 0x67, 0x30, 0x48, 0xa3, 0x54, 0x12, 0x17, 0x0f, 0xb2, 0xf9, 0x58, 0xc2, 0xcf, 0x06,
	0x39, 0xfc, 0x88, 0xb7, 0x03, 0xa2, 0xc9, 0x8a, 0xc9, 0x82, 0xd1, 0x89, 0x2d, 0x8a, 0x05, 0xa4,
	0x20, 0x89, 0x0d, 0xac, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x95, 0xf0, 0x9d, 0xda,
	0x01, 0x00, 0x00,
}
