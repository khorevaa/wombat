// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.5.1
// source: foobar.proto

package server

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Baz int32

const (
	Baz_UNKNOWN Baz = 0
	Baz_FOO     Baz = 1
	Baz_BAR     Baz = 2
	Baz_BAZ     Baz = 3
)

// Enum value maps for Baz.
var (
	Baz_name = map[int32]string{
		0: "UNKNOWN",
		1: "FOO",
		2: "BAR",
		3: "BAZ",
	}
	Baz_value = map[string]int32{
		"UNKNOWN": 0,
		"FOO":     1,
		"BAR":     2,
		"BAZ":     3,
	}
)

func (x Baz) Enum() *Baz {
	p := new(Baz)
	*p = x
	return p
}

func (x Baz) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Baz) Descriptor() protoreflect.EnumDescriptor {
	return file_foobar_proto_enumTypes[0].Descriptor()
}

func (Baz) Type() protoreflect.EnumType {
	return &file_foobar_proto_enumTypes[0]
}

func (x Baz) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Baz.Descriptor instead.
func (Baz) EnumDescriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{0}
}

type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Names []*Bar_Nested `protobuf:"bytes,3,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bar.ProtoReflect.Descriptor instead.
func (*Bar) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{0}
}

func (x *Bar) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bar) GetNames() []*Bar_Nested {
	if x != nil {
		return x.Names
	}
	return nil
}

type FooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeDouble   float64 `protobuf:"fixed64,1,opt,name=type_double,json=typeDouble,proto3" json:"type_double,omitempty"`
	TypeFloat    float32 `protobuf:"fixed32,2,opt,name=type_float,json=typeFloat,proto3" json:"type_float,omitempty"`
	TypeInt64    int64   `protobuf:"varint,3,opt,name=type_int64,json=typeInt64,proto3" json:"type_int64,omitempty"`
	TypeUint64   uint64  `protobuf:"varint,4,opt,name=type_uint64,json=typeUint64,proto3" json:"type_uint64,omitempty"`
	TypeInt32    int32   `protobuf:"varint,5,opt,name=type_int32,json=typeInt32,proto3" json:"type_int32,omitempty"`
	TypeFixed64  uint64  `protobuf:"fixed64,6,opt,name=type_fixed64,json=typeFixed64,proto3" json:"type_fixed64,omitempty"`
	TypeFixed32  uint32  `protobuf:"fixed32,7,opt,name=type_fixed32,json=typeFixed32,proto3" json:"type_fixed32,omitempty"`
	TypeBool     bool    `protobuf:"varint,8,opt,name=type_bool,json=typeBool,proto3" json:"type_bool,omitempty"`
	TypeString   string  `protobuf:"bytes,9,opt,name=type_string,json=typeString,proto3" json:"type_string,omitempty"`
	TypeMessage  *Bar    `protobuf:"bytes,10,opt,name=type_message,json=typeMessage,proto3" json:"type_message,omitempty"`
	TypeBytes    []byte  `protobuf:"bytes,11,opt,name=type_bytes,json=typeBytes,proto3" json:"type_bytes,omitempty"`
	TypeUint32   uint32  `protobuf:"varint,12,opt,name=type_uint32,json=typeUint32,proto3" json:"type_uint32,omitempty"`
	TypeEnum     Baz     `protobuf:"varint,13,opt,name=type_enum,json=typeEnum,proto3,enum=wombat.v1.Baz" json:"type_enum,omitempty"`
	TypeSfixed32 int32   `protobuf:"fixed32,14,opt,name=type_sfixed32,json=typeSfixed32,proto3" json:"type_sfixed32,omitempty"`
	TypeSfixed64 int64   `protobuf:"fixed64,15,opt,name=type_sfixed64,json=typeSfixed64,proto3" json:"type_sfixed64,omitempty"`
	TypeSint32   int32   `protobuf:"zigzag32,16,opt,name=type_sint32,json=typeSint32,proto3" json:"type_sint32,omitempty"`
	TypeSint64   int64   `protobuf:"zigzag64,17,opt,name=type_sint64,json=typeSint64,proto3" json:"type_sint64,omitempty"`
}

func (x *FooRequest) Reset() {
	*x = FooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooRequest) ProtoMessage() {}

func (x *FooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooRequest.ProtoReflect.Descriptor instead.
func (*FooRequest) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{1}
}

func (x *FooRequest) GetTypeDouble() float64 {
	if x != nil {
		return x.TypeDouble
	}
	return 0
}

func (x *FooRequest) GetTypeFloat() float32 {
	if x != nil {
		return x.TypeFloat
	}
	return 0
}

func (x *FooRequest) GetTypeInt64() int64 {
	if x != nil {
		return x.TypeInt64
	}
	return 0
}

func (x *FooRequest) GetTypeUint64() uint64 {
	if x != nil {
		return x.TypeUint64
	}
	return 0
}

func (x *FooRequest) GetTypeInt32() int32 {
	if x != nil {
		return x.TypeInt32
	}
	return 0
}

func (x *FooRequest) GetTypeFixed64() uint64 {
	if x != nil {
		return x.TypeFixed64
	}
	return 0
}

func (x *FooRequest) GetTypeFixed32() uint32 {
	if x != nil {
		return x.TypeFixed32
	}
	return 0
}

func (x *FooRequest) GetTypeBool() bool {
	if x != nil {
		return x.TypeBool
	}
	return false
}

func (x *FooRequest) GetTypeString() string {
	if x != nil {
		return x.TypeString
	}
	return ""
}

func (x *FooRequest) GetTypeMessage() *Bar {
	if x != nil {
		return x.TypeMessage
	}
	return nil
}

func (x *FooRequest) GetTypeBytes() []byte {
	if x != nil {
		return x.TypeBytes
	}
	return nil
}

func (x *FooRequest) GetTypeUint32() uint32 {
	if x != nil {
		return x.TypeUint32
	}
	return 0
}

func (x *FooRequest) GetTypeEnum() Baz {
	if x != nil {
		return x.TypeEnum
	}
	return Baz_UNKNOWN
}

func (x *FooRequest) GetTypeSfixed32() int32 {
	if x != nil {
		return x.TypeSfixed32
	}
	return 0
}

func (x *FooRequest) GetTypeSfixed64() int64 {
	if x != nil {
		return x.TypeSfixed64
	}
	return 0
}

func (x *FooRequest) GetTypeSint32() int32 {
	if x != nil {
		return x.TypeSint32
	}
	return 0
}

func (x *FooRequest) GetTypeSint64() int64 {
	if x != nil {
		return x.TypeSint64
	}
	return 0
}

type FooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FooResponse) Reset() {
	*x = FooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooResponse) ProtoMessage() {}

func (x *FooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooResponse.ProtoReflect.Descriptor instead.
func (*FooResponse) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{2}
}

type BarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeRepeatedString  []string  `protobuf:"bytes,1,rep,name=type_repeated_string,json=typeRepeatedString,proto3" json:"type_repeated_string,omitempty"`
	TypeRepeatedFloat   []float32 `protobuf:"fixed32,2,rep,packed,name=type_repeated_float,json=typeRepeatedFloat,proto3" json:"type_repeated_float,omitempty"`
	TypeRepeatedBool    []bool    `protobuf:"varint,3,rep,packed,name=type_repeated_bool,json=typeRepeatedBool,proto3" json:"type_repeated_bool,omitempty"`
	TypeRepeatedEnum    []Baz     `protobuf:"varint,4,rep,packed,name=type_repeated_enum,json=typeRepeatedEnum,proto3,enum=wombat.v1.Baz" json:"type_repeated_enum,omitempty"`
	TypeRepeatedBytes   [][]byte  `protobuf:"bytes,5,rep,name=type_repeated_bytes,json=typeRepeatedBytes,proto3" json:"type_repeated_bytes,omitempty"`
	TypeRepeatedMessage []*Bar    `protobuf:"bytes,6,rep,name=type_repeated_message,json=typeRepeatedMessage,proto3" json:"type_repeated_message,omitempty"`
}

func (x *BarRequest) Reset() {
	*x = BarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarRequest) ProtoMessage() {}

func (x *BarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarRequest.ProtoReflect.Descriptor instead.
func (*BarRequest) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{3}
}

func (x *BarRequest) GetTypeRepeatedString() []string {
	if x != nil {
		return x.TypeRepeatedString
	}
	return nil
}

func (x *BarRequest) GetTypeRepeatedFloat() []float32 {
	if x != nil {
		return x.TypeRepeatedFloat
	}
	return nil
}

func (x *BarRequest) GetTypeRepeatedBool() []bool {
	if x != nil {
		return x.TypeRepeatedBool
	}
	return nil
}

func (x *BarRequest) GetTypeRepeatedEnum() []Baz {
	if x != nil {
		return x.TypeRepeatedEnum
	}
	return nil
}

func (x *BarRequest) GetTypeRepeatedBytes() [][]byte {
	if x != nil {
		return x.TypeRepeatedBytes
	}
	return nil
}

func (x *BarRequest) GetTypeRepeatedMessage() []*Bar {
	if x != nil {
		return x.TypeRepeatedMessage
	}
	return nil
}

type BarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BarResponse) Reset() {
	*x = BarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BarResponse) ProtoMessage() {}

func (x *BarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BarResponse.ProtoReflect.Descriptor instead.
func (*BarResponse) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{4}
}

type BazRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*BazRequest_Foo
	//	*BazRequest_Bar
	//	*BazRequest_Baz
	Request isBazRequest_Request `protobuf_oneof:"request"`
}

func (x *BazRequest) Reset() {
	*x = BazRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BazRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BazRequest) ProtoMessage() {}

func (x *BazRequest) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BazRequest.ProtoReflect.Descriptor instead.
func (*BazRequest) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{5}
}

func (m *BazRequest) GetRequest() isBazRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *BazRequest) GetFoo() *FooRequest {
	if x, ok := x.GetRequest().(*BazRequest_Foo); ok {
		return x.Foo
	}
	return nil
}

func (x *BazRequest) GetBar() *BarRequest {
	if x, ok := x.GetRequest().(*BazRequest_Bar); ok {
		return x.Bar
	}
	return nil
}

func (x *BazRequest) GetBaz() string {
	if x, ok := x.GetRequest().(*BazRequest_Baz); ok {
		return x.Baz
	}
	return ""
}

type isBazRequest_Request interface {
	isBazRequest_Request()
}

type BazRequest_Foo struct {
	Foo *FooRequest `protobuf:"bytes,1,opt,name=foo,proto3,oneof"`
}

type BazRequest_Bar struct {
	Bar *BarRequest `protobuf:"bytes,2,opt,name=bar,proto3,oneof"`
}

type BazRequest_Baz struct {
	Baz string `protobuf:"bytes,3,opt,name=baz,proto3,oneof"`
}

func (*BazRequest_Foo) isBazRequest_Request() {}

func (*BazRequest_Bar) isBazRequest_Request() {}

func (*BazRequest_Baz) isBazRequest_Request() {}

type BazResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BazResponse) Reset() {
	*x = BazResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BazResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BazResponse) ProtoMessage() {}

func (x *BazResponse) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BazResponse.ProtoReflect.Descriptor instead.
func (*BazResponse) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{6}
}

type Bar_Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsNested bool   `protobuf:"varint,2,opt,name=is_nested,json=isNested,proto3" json:"is_nested,omitempty"`
}

func (x *Bar_Nested) Reset() {
	*x = Bar_Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foobar_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar_Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar_Nested) ProtoMessage() {}

func (x *Bar_Nested) ProtoReflect() protoreflect.Message {
	mi := &file_foobar_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bar_Nested.ProtoReflect.Descriptor instead.
func (*Bar_Nested) Descriptor() ([]byte, []int) {
	return file_foobar_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Bar_Nested) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bar_Nested) GetIsNested() bool {
	if x != nil {
		return x.IsNested
	}
	return false
}

var File_foobar_proto protoreflect.FileDescriptor

var file_foobar_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x77, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x7d, 0x0a, 0x03, 0x42, 0x61, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2b, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x77, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x72, 0x2e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x39, 0x0a,
	0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x22, 0xdb, 0x04, 0x0a, 0x0a, 0x46, 0x6f, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x79,
	0x70, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x79,
	0x70, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x79, 0x70,
	0x65, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x79, 0x70,
	0x65, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x79, 0x70,
	0x65, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0b, 0x74, 0x79,
	0x70, 0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x07, 0x52,
	0x0b, 0x74, 0x79, 0x70, 0x65, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x74, 0x79, 0x70, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x0c, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x77, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x72,
	0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x74, 0x79, 0x70, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x2b, 0x0a,
	0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x77, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x7a,
	0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0f, 0x52, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12,
	0x23, 0x0a, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x10, 0x52, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x53, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x18, 0x10, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x53,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x11, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65,
	0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x22, 0x0d, 0x0a, 0x0b, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xce, 0x02, 0x0a, 0x0a, 0x42, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x12, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x02, 0x52, 0x11, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x08, 0x52, 0x10, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x3c, 0x0a, 0x12, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x77, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x7a,
	0x52, 0x10, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x11, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x42, 0x0a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x77, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61,
	0x72, 0x52, 0x13, 0x74, 0x79, 0x70, 0x65, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x42, 0x61, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x0a, 0x42, 0x61, 0x7a, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x77, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12,
	0x29, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77,
	0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x03, 0x62, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x03, 0x62, 0x61,
	0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x62, 0x61, 0x7a, 0x42, 0x09,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x42, 0x61, 0x7a,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x2d, 0x0a, 0x03, 0x42, 0x61, 0x7a, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x46, 0x4f, 0x4f, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x41, 0x52, 0x10, 0x02, 0x12, 0x07,
	0x0a, 0x03, 0x42, 0x41, 0x5a, 0x10, 0x03, 0x32, 0xb0, 0x01, 0x0a, 0x06, 0x46, 0x6f, 0x6f, 0x62,
	0x61, 0x72, 0x12, 0x36, 0x0a, 0x03, 0x42, 0x61, 0x7a, 0x12, 0x15, 0x2e, 0x77, 0x6f, 0x6d, 0x62,
	0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x77, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x7a,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x03, 0x42, 0x61,
	0x72, 0x12, 0x15, 0x2e, 0x77, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x77, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x15, 0x2e, 0x77, 0x6f, 0x6d, 0x62,
	0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x77, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foobar_proto_rawDescOnce sync.Once
	file_foobar_proto_rawDescData = file_foobar_proto_rawDesc
)

func file_foobar_proto_rawDescGZIP() []byte {
	file_foobar_proto_rawDescOnce.Do(func() {
		file_foobar_proto_rawDescData = protoimpl.X.CompressGZIP(file_foobar_proto_rawDescData)
	})
	return file_foobar_proto_rawDescData
}

var file_foobar_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_foobar_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_foobar_proto_goTypes = []interface{}{
	(Baz)(0),            // 0: wombat.v1.Baz
	(*Bar)(nil),         // 1: wombat.v1.Bar
	(*FooRequest)(nil),  // 2: wombat.v1.FooRequest
	(*FooResponse)(nil), // 3: wombat.v1.FooResponse
	(*BarRequest)(nil),  // 4: wombat.v1.BarRequest
	(*BarResponse)(nil), // 5: wombat.v1.BarResponse
	(*BazRequest)(nil),  // 6: wombat.v1.BazRequest
	(*BazResponse)(nil), // 7: wombat.v1.BazResponse
	(*Bar_Nested)(nil),  // 8: wombat.v1.Bar.Nested
}
var file_foobar_proto_depIdxs = []int32{
	8,  // 0: wombat.v1.Bar.names:type_name -> wombat.v1.Bar.Nested
	1,  // 1: wombat.v1.FooRequest.type_message:type_name -> wombat.v1.Bar
	0,  // 2: wombat.v1.FooRequest.type_enum:type_name -> wombat.v1.Baz
	0,  // 3: wombat.v1.BarRequest.type_repeated_enum:type_name -> wombat.v1.Baz
	1,  // 4: wombat.v1.BarRequest.type_repeated_message:type_name -> wombat.v1.Bar
	2,  // 5: wombat.v1.BazRequest.foo:type_name -> wombat.v1.FooRequest
	4,  // 6: wombat.v1.BazRequest.bar:type_name -> wombat.v1.BarRequest
	6,  // 7: wombat.v1.Foobar.Baz:input_type -> wombat.v1.BazRequest
	4,  // 8: wombat.v1.Foobar.Bar:input_type -> wombat.v1.BarRequest
	2,  // 9: wombat.v1.Foobar.Foo:input_type -> wombat.v1.FooRequest
	7,  // 10: wombat.v1.Foobar.Baz:output_type -> wombat.v1.BazResponse
	5,  // 11: wombat.v1.Foobar.Bar:output_type -> wombat.v1.BarResponse
	3,  // 12: wombat.v1.Foobar.Foo:output_type -> wombat.v1.FooResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_foobar_proto_init() }
func file_foobar_proto_init() {
	if File_foobar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foobar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_foobar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_foobar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_foobar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_foobar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BarResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_foobar_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BazRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_foobar_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BazResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_foobar_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar_Nested); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_foobar_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*BazRequest_Foo)(nil),
		(*BazRequest_Bar)(nil),
		(*BazRequest_Baz)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_foobar_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_foobar_proto_goTypes,
		DependencyIndexes: file_foobar_proto_depIdxs,
		EnumInfos:         file_foobar_proto_enumTypes,
		MessageInfos:      file_foobar_proto_msgTypes,
	}.Build()
	File_foobar_proto = out.File
	file_foobar_proto_rawDesc = nil
	file_foobar_proto_goTypes = nil
	file_foobar_proto_depIdxs = nil
}