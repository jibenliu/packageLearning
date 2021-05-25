// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package pro

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

type PhoneType int32

const (
	PhoneType_HOME PhoneType = 0
	PhoneType_WORK PhoneType = 1
)

var PhoneType_name = map[int32]string{
	0: "HOME",
	1: "WORK",
}

var PhoneType_value = map[string]int32{
	"HOME": 0,
	"WORK": 1,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}

func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

//每个message名都要大写
type Phone struct {
	Type                 PhoneType `protobuf:"varint,1,opt,name=type,proto3,enum=pro.PhoneType" json:"type,omitempty"`
	Number               string    `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Phone) Reset()         { *m = Phone{} }
func (m *Phone) String() string { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()    {}
func (*Phone) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

func (m *Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phone.Unmarshal(m, b)
}
func (m *Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phone.Marshal(b, m, deterministic)
}
func (m *Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phone.Merge(m, src)
}
func (m *Phone) XXX_Size() int {
	return xxx_messageInfo_Phone.Size(m)
}
func (m *Phone) XXX_DiscardUnknown() {
	xxx_messageInfo_Phone.DiscardUnknown(m)
}

var xxx_messageInfo_Phone proto.InternalMessageInfo

func (m *Phone) GetType() PhoneType {
	if m != nil {
		return m.Type
	}
	return PhoneType_HOME
}

func (m *Phone) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type Person struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phones               []*Phone `protobuf:"bytes,3,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{1}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetPhones() []*Phone {
	if m != nil {
		return m.Phones
	}
	return nil
}

type ContactBook struct {
	Persons              []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ContactBook) Reset()         { *m = ContactBook{} }
func (m *ContactBook) String() string { return proto.CompactTextString(m) }
func (*ContactBook) ProtoMessage()    {}
func (*ContactBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{2}
}

func (m *ContactBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactBook.Unmarshal(m, b)
}
func (m *ContactBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactBook.Marshal(b, m, deterministic)
}
func (m *ContactBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactBook.Merge(m, src)
}
func (m *ContactBook) XXX_Size() int {
	return xxx_messageInfo_ContactBook.Size(m)
}
func (m *ContactBook) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactBook.DiscardUnknown(m)
}

var xxx_messageInfo_ContactBook proto.InternalMessageInfo

func (m *ContactBook) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

func init() {
	proto.RegisterEnum("pro.PhoneType", PhoneType_name, PhoneType_value)
	proto.RegisterType((*Phone)(nil), "pro.Phone")
	proto.RegisterType((*Person)(nil), "pro.Person")
	proto.RegisterType((*ContactBook)(nil), "pro.ContactBook")
}

func init() {
	proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150)
}

var fileDescriptor_db1b6b0986796150 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x4d, 0xdb, 0x8d, 0xee, 0x14, 0xca, 0x32, 0x07, 0xc9, 0xcd, 0x12, 0x10, 0x8a, 0x87,
	0x3d, 0xac, 0x3e, 0x81, 0x8b, 0x20, 0x88, 0x6c, 0x09, 0x82, 0xe7, 0xd6, 0x06, 0x2c, 0xd2, 0x4c,
	0x48, 0xe2, 0xa1, 0x6f, 0x2f, 0x1d, 0xab, 0x7b, 0xfb, 0x92, 0xdf, 0xf7, 0x27, 0x01, 0xe8, 0xbb,
	0x68, 0xf7, 0x3e, 0x50, 0x22, 0xcc, 0x7d, 0x20, 0x7d, 0x84, 0x4d, 0xfb, 0x49, 0xce, 0xa2, 0x86,
	0x22, 0xcd, 0xde, 0x2a, 0x51, 0x8b, 0xa6, 0x3a, 0x54, 0x8b, 0x67, 0xcf, 0xe4, 0x6d, 0xf6, 0xd6,
	0x30, 0xc3, 0x6b, 0x90, 0xee, 0x7b, 0xea, 0x6d, 0x50, 0x59, 0x2d, 0x9a, 0xad, 0x59, 0x4f, 0xba,
	0x05, 0xd9, 0xda, 0x10, 0xc9, 0x61, 0x05, 0xd9, 0x38, 0x70, 0xc7, 0xc6, 0x64, 0xe3, 0x80, 0x08,
	0x85, 0xeb, 0x26, 0xbb, 0xfa, 0x59, 0xa3, 0x06, 0xe9, 0x97, 0xe2, 0xa8, 0xf2, 0x3a, 0x6f, 0xca,
	0x03, 0x9c, 0xb7, 0xcc, 0x4a, 0xf4, 0x03, 0x94, 0x47, 0x72, 0xa9, 0xfb, 0x48, 0x8f, 0x44, 0x5f,
	0x78, 0x0b, 0x97, 0x9e, 0x07, 0xa2, 0x12, 0x9c, 0x29, 0x7f, 0x33, 0x7c, 0x67, 0xfe, 0xd8, 0xdd,
	0x0d, 0x6c, 0xff, 0x9f, 0x8c, 0x57, 0x50, 0x3c, 0x9f, 0x5e, 0x9f, 0x76, 0x17, 0x8b, 0x7a, 0x3f,
	0x99, 0x97, 0x9d, 0xe8, 0x25, 0xff, 0xfc, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x6c, 0xac,
	0x69, 0x07, 0x01, 0x00, 0x00,
}