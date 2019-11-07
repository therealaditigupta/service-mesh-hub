// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/v1/ingress.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

//
//MeshIngress represents a managed ingress (edge router) which can proxy connections
//for services in a Mesh managed by SuperGloo
type MeshIngress struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by supergloo during validation
	Status core.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
	// Types that are valid to be assigned to IngressType:
	//	*MeshIngress_Gloo
	IngressType          isMeshIngress_IngressType `protobuf_oneof:"ingress_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *MeshIngress) Reset()         { *m = MeshIngress{} }
func (m *MeshIngress) String() string { return proto.CompactTextString(m) }
func (*MeshIngress) ProtoMessage()    {}
func (*MeshIngress) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f8926cb2b4e4f9d, []int{0}
}
func (m *MeshIngress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshIngress.Unmarshal(m, b)
}
func (m *MeshIngress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshIngress.Marshal(b, m, deterministic)
}
func (m *MeshIngress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshIngress.Merge(m, src)
}
func (m *MeshIngress) XXX_Size() int {
	return xxx_messageInfo_MeshIngress.Size(m)
}
func (m *MeshIngress) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshIngress.DiscardUnknown(m)
}

var xxx_messageInfo_MeshIngress proto.InternalMessageInfo

type isMeshIngress_IngressType interface {
	isMeshIngress_IngressType()
	Equal(interface{}) bool
}

type MeshIngress_Gloo struct {
	Gloo *MeshIngress_GlooIngress `protobuf:"bytes,3,opt,name=gloo,proto3,oneof" json:"gloo,omitempty"`
}

func (*MeshIngress_Gloo) isMeshIngress_IngressType() {}

func (m *MeshIngress) GetIngressType() isMeshIngress_IngressType {
	if m != nil {
		return m.IngressType
	}
	return nil
}

func (m *MeshIngress) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *MeshIngress) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *MeshIngress) GetGloo() *MeshIngress_GlooIngress {
	if x, ok := m.GetIngressType().(*MeshIngress_Gloo); ok {
		return x.Gloo
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MeshIngress) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MeshIngress_Gloo)(nil),
	}
}

type MeshIngress_GlooIngress struct {
	// namespace which the ingress is located in.
	Namespace   string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// name of port which will be used for mesh traffic.
	Port                 string   `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshIngress_GlooIngress) Reset()         { *m = MeshIngress_GlooIngress{} }
func (m *MeshIngress_GlooIngress) String() string { return proto.CompactTextString(m) }
func (*MeshIngress_GlooIngress) ProtoMessage()    {}
func (*MeshIngress_GlooIngress) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f8926cb2b4e4f9d, []int{0, 0}
}
func (m *MeshIngress_GlooIngress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshIngress_GlooIngress.Unmarshal(m, b)
}
func (m *MeshIngress_GlooIngress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshIngress_GlooIngress.Marshal(b, m, deterministic)
}
func (m *MeshIngress_GlooIngress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshIngress_GlooIngress.Merge(m, src)
}
func (m *MeshIngress_GlooIngress) XXX_Size() int {
	return xxx_messageInfo_MeshIngress_GlooIngress.Size(m)
}
func (m *MeshIngress_GlooIngress) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshIngress_GlooIngress.DiscardUnknown(m)
}

var xxx_messageInfo_MeshIngress_GlooIngress proto.InternalMessageInfo

func (m *MeshIngress_GlooIngress) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *MeshIngress_GlooIngress) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *MeshIngress_GlooIngress) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func init() {
	proto.RegisterType((*MeshIngress)(nil), "zephyr.solo.io.MeshIngress")
	proto.RegisterType((*MeshIngress_GlooIngress)(nil), "zephyr.solo.io.MeshIngress.GlooIngress")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/v1/ingress.proto", fileDescriptor_4f8926cb2b4e4f9d)
}

var fileDescriptor_4f8926cb2b4e4f9d = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0xce, 0xd2, 0x40,
	0x10, 0x06, 0x6c, 0x88, 0xdd, 0x22, 0xc6, 0x0d, 0x31, 0x88, 0x46, 0x94, 0x8b, 0x26, 0x86, 0x6d,
	0xd0, 0x98, 0x18, 0x12, 0x2f, 0xbd, 0xa0, 0x07, 0x3c, 0xd4, 0x9b, 0x17, 0xb2, 0x94, 0xa1, 0x5d,
	0xa1, 0xcc, 0x66, 0x77, 0x21, 0xc1, 0x23, 0x4f, 0xe3, 0xa3, 0xf8, 0x14, 0xc4, 0xf8, 0x06, 0xf8,
	0x04, 0xa6, 0xdb, 0xad, 0x3f, 0x24, 0x7f, 0xf2, 0xff, 0xff, 0xa9, 0x33, 0xdf, 0xcc, 0xf7, 0x75,
	0xbe, 0x9d, 0x21, 0xef, 0x53, 0x61, 0xb2, 0xed, 0x9c, 0x25, 0x98, 0x87, 0x1a, 0xd7, 0x38, 0x14,
	0x18, 0xe6, 0xa0, 0xb3, 0xa1, 0x54, 0xf8, 0x1d, 0x12, 0xa3, 0x43, 0x2e, 0x45, 0xb8, 0x1b, 0x85,
	0x62, 0x93, 0x2a, 0xd0, 0x9a, 0x49, 0x85, 0x06, 0x69, 0xfb, 0x07, 0xc8, 0x6c, 0xaf, 0x58, 0x41,
	0x61, 0x02, 0x7b, 0x9d, 0x14, 0x53, 0xb4, 0xa5, 0xb0, 0x88, 0xca, 0xae, 0xde, 0xe8, 0x1a, 0x71,
	0xfb, 0x5d, 0x09, 0x53, 0xe9, 0xe6, 0x60, 0xf8, 0x82, 0x1b, 0xee, 0x28, 0xe1, 0x2d, 0x28, 0xda,
	0x70, 0xb3, 0xd5, 0x77, 0xf8, 0x47, 0x95, 0x97, 0x94, 0xc1, 0xef, 0x06, 0x09, 0xa6, 0xa0, 0xb3,
	0xcf, 0xa5, 0x25, 0x3a, 0x21, 0xcd, 0x52, 0xb2, 0x5b, 0x7f, 0x51, 0x7f, 0x1d, 0xbc, 0xed, 0xb0,
	0x04, 0x15, 0x54, 0xde, 0xd8, 0x57, 0x5b, 0x8b, 0x9e, 0xfc, 0x3a, 0xf6, 0x6b, 0x7f, 0x8f, 0xfd,
	0x47, 0x06, 0xb4, 0x59, 0x88, 0xe5, 0x72, 0x3c, 0x10, 0xe9, 0x06, 0x15, 0x0c, 0x62, 0x47, 0xa7,
	0x1f, 0xc8, 0xfd, 0xca, 0x4e, 0xb7, 0x61, 0xa5, 0x1e, 0x5f, 0x4a, 0x4d, 0x5d, 0x35, 0xf2, 0x0a,
	0xb1, 0xf8, 0x7f, 0x37, 0xfd, 0x48, 0xbc, 0x74, 0x8d, 0xd8, 0xbd, 0x67, 0x59, 0xaf, 0xd8, 0xe5,
	0xf3, 0xb2, 0xb3, 0x69, 0xd9, 0x64, 0x8d, 0xe8, 0xe2, 0x4f, 0xb5, 0xd8, 0xd2, 0x7a, 0x73, 0x12,
	0x9c, 0xc1, 0xf4, 0x19, 0xf1, 0x37, 0x3c, 0x07, 0x2d, 0x79, 0x02, 0xd6, 0x93, 0x1f, 0x5f, 0x01,
	0xf4, 0x25, 0x69, 0x69, 0x50, 0x3b, 0x91, 0xc0, 0xac, 0x00, 0xed, 0xa4, 0x7e, 0x1c, 0x38, 0xec,
	0x0b, 0xcf, 0x81, 0x52, 0xe2, 0x49, 0x54, 0xc6, 0x8e, 0xe3, 0xc7, 0x36, 0x1e, 0x3f, 0x3d, 0x9c,
	0x3c, 0x8f, 0x34, 0x72, 0x71, 0x38, 0x79, 0x0f, 0xe9, 0x83, 0xe2, 0x48, 0xdc, 0x4d, 0x80, 0x8e,
	0xda, 0xa4, 0xe5, 0x92, 0x99, 0xd9, 0x4b, 0x88, 0x46, 0x3f, 0xff, 0x3c, 0xaf, 0x7f, 0x7b, 0x73,
	0xe3, 0x71, 0xc9, 0x55, 0xea, 0x96, 0x34, 0x6f, 0xda, 0xe5, 0xbc, 0xfb, 0x17, 0x00, 0x00, 0xff,
	0xff, 0xcc, 0xbe, 0xcc, 0xcc, 0x92, 0x02, 0x00, 0x00,
}

func (this *MeshIngress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshIngress)
	if !ok {
		that2, ok := that.(MeshIngress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if that1.IngressType == nil {
		if this.IngressType != nil {
			return false
		}
	} else if this.IngressType == nil {
		return false
	} else if !this.IngressType.Equal(that1.IngressType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshIngress_Gloo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshIngress_Gloo)
	if !ok {
		that2, ok := that.(MeshIngress_Gloo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Gloo.Equal(that1.Gloo) {
		return false
	}
	return true
}
func (this *MeshIngress_GlooIngress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshIngress_GlooIngress)
	if !ok {
		that2, ok := that.(MeshIngress_GlooIngress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
