// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/http_uri.proto

package envoy_api_v2_core

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type HttpUri struct {
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Types that are valid to be assigned to HttpUpstreamType:
	//	*HttpUri_Cluster
	HttpUpstreamType     isHttpUri_HttpUpstreamType `protobuf_oneof:"http_upstream_type"`
	Timeout              *duration.Duration         `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *HttpUri) Reset()         { *m = HttpUri{} }
func (m *HttpUri) String() string { return proto.CompactTextString(m) }
func (*HttpUri) ProtoMessage()    {}
func (*HttpUri) Descriptor() ([]byte, []int) {
	return fileDescriptor_1660b946db74c078, []int{0}
}

func (m *HttpUri) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpUri.Unmarshal(m, b)
}
func (m *HttpUri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpUri.Marshal(b, m, deterministic)
}
func (m *HttpUri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpUri.Merge(m, src)
}
func (m *HttpUri) XXX_Size() int {
	return xxx_messageInfo_HttpUri.Size(m)
}
func (m *HttpUri) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpUri.DiscardUnknown(m)
}

var xxx_messageInfo_HttpUri proto.InternalMessageInfo

func (m *HttpUri) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type isHttpUri_HttpUpstreamType interface {
	isHttpUri_HttpUpstreamType()
}

type HttpUri_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*HttpUri_Cluster) isHttpUri_HttpUpstreamType() {}

func (m *HttpUri) GetHttpUpstreamType() isHttpUri_HttpUpstreamType {
	if m != nil {
		return m.HttpUpstreamType
	}
	return nil
}

func (m *HttpUri) GetCluster() string {
	if x, ok := m.GetHttpUpstreamType().(*HttpUri_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *HttpUri) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpUri) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpUri_Cluster)(nil),
	}
}

func init() {
	proto.RegisterType((*HttpUri)(nil), "envoy.api.v2.core.HttpUri")
}

func init() { proto.RegisterFile("envoy/api/v2/core/http_uri.proto", fileDescriptor_1660b946db74c078) }

var fileDescriptor_1660b946db74c078 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0xad, 0x1b, 0x20, 0x60, 0x18, 0xc0, 0x0b, 0x49, 0x07, 0x88, 0x60, 0xe9, 0x64, 0x8b, 0xf0,
	0x07, 0x16, 0x48, 0x1d, 0xab, 0x48, 0xcc, 0x95, 0xdb, 0x9a, 0x62, 0x29, 0xed, 0x59, 0xd7, 0x73,
	0x44, 0x7e, 0x09, 0x7e, 0x10, 0x75, 0x42, 0x89, 0x9b, 0x01, 0x31, 0xdd, 0xe9, 0xde, 0x7b, 0x7a,
	0xef, 0x1d, 0x2f, 0xec, 0xae, 0x81, 0x56, 0x19, 0xef, 0x54, 0x53, 0xaa, 0x15, 0xa0, 0x55, 0x1f,
	0x44, 0x7e, 0x11, 0xd0, 0x49, 0x8f, 0x40, 0x20, 0x6e, 0x7a, 0x86, 0x34, 0xde, 0xc9, 0xa6, 0x94,
	0x1d, 0x63, 0x72, 0xb7, 0x01, 0xd8, 0xd4, 0x56, 0xf5, 0x84, 0x65, 0x78, 0x57, 0xeb, 0x80, 0x86,
	0x1c, 0xec, 0xa2, 0x64, 0x72, 0xdb, 0x98, 0xda, 0xad, 0x0d, 0x59, 0x35, 0x2c, 0x11, 0x78, 0xf8,
	0x66, 0x3c, 0x9d, 0x11, 0xf9, 0x37, 0x74, 0x22, 0xe7, 0x49, 0x40, 0x97, 0xb1, 0x82, 0x4d, 0x2f,
	0x74, 0x7a, 0xd0, 0x27, 0x38, 0x2e, 0x58, 0xd5, 0xdd, 0xc4, 0x23, 0x4f, 0x57, 0x75, 0xd8, 0x93,
	0xc5, 0x6c, 0xfc, 0x07, 0x9e, 0x8d, 0xaa, 0x01, 0x11, 0xaf, 0x3c, 0x25, 0xb7, 0xb5, 0x10, 0x28,
	0x4b, 0x0a, 0x36, 0xbd, 0x2c, 0x73, 0x19, 0x63, 0xc9, 0x21, 0x96, 0x7c, 0x39, 0xc6, 0xd2, 0xd7,
	0x07, 0x7d, 0xfa, 0xc5, 0xc6, 0xe5, 0x28, 0xce, 0x73, 0x56, 0x0d, 0x5a, 0x9d, 0x73, 0x11, 0x0b,
	0xfb, 0x3d, 0xa1, 0x35, 0xdb, 0x05, 0xb5, 0xde, 0x8a, 0xe4, 0x47, 0x33, 0xfd, 0xc4, 0xef, 0x1d,
	0xc8, 0xbe, 0xbe, 0x47, 0xf8, 0x6c, 0xe5, 0xbf, 0x4f, 0xe8, 0xab, 0x63, 0x9b, 0x79, 0x67, 0x39,
	0x67, 0xcb, 0xb3, 0xde, 0xfb, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x03, 0xc2, 0xe6, 0x57,
	0x01, 0x00, 0x00,
}
