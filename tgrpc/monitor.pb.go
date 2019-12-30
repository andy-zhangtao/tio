// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor.proto

package tio_control_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MonitorScalaRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MonitorScalaRequest) Reset()         { *m = MonitorScalaRequest{} }
func (m *MonitorScalaRequest) String() string { return proto.CompactTextString(m) }
func (*MonitorScalaRequest) ProtoMessage()    {}
func (*MonitorScalaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44174b7b2a306b71, []int{0}
}

func (m *MonitorScalaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorScalaRequest.Unmarshal(m, b)
}
func (m *MonitorScalaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorScalaRequest.Marshal(b, m, deterministic)
}
func (m *MonitorScalaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorScalaRequest.Merge(m, src)
}
func (m *MonitorScalaRequest) XXX_Size() int {
	return xxx_messageInfo_MonitorScalaRequest.Size(m)
}
func (m *MonitorScalaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorScalaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorScalaRequest proto.InternalMessageInfo

func (m *MonitorScalaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MonitorScalaRequest) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*MonitorScalaRequest)(nil), "MonitorScalaRequest")
}

func init() { proto.RegisterFile("monitor.proto", fileDescriptor_44174b7b2a306b71) }

var fileDescriptor_44174b7b2a306b71 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xcd, 0xcf, 0xcb,
	0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x4e, 0xce, 0xcf, 0x2b, 0x2e,
	0x81, 0x70, 0x94, 0xac, 0xb9, 0x84, 0x7d, 0x21, 0xb2, 0xc1, 0xc9, 0x89, 0x39, 0x89, 0x41, 0xa9,
	0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x00, 0x17, 0x73, 0x5e, 0x69, 0xae, 0x04, 0x93, 0x02,
	0xa3, 0x06, 0x6b, 0x10, 0x88, 0x69, 0x64, 0xc3, 0xc5, 0x07, 0xd3, 0x9c, 0x5a, 0x54, 0x96, 0x99,
	0x9c, 0x2a, 0xa4, 0xc5, 0xc5, 0x0a, 0x36, 0x47, 0x48, 0x44, 0x0f, 0x8b, 0xb1, 0x52, 0x9c, 0x7a,
	0x21, 0x99, 0xf9, 0x41, 0xa9, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x4e, 0x02, 0x51, 0x7c, 0x25, 0x99,
	0xf9, 0x7a, 0xc9, 0xf9, 0x79, 0x25, 0x45, 0xf9, 0x39, 0x7a, 0x65, 0x86, 0x49, 0x6c, 0x60, 0x37,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xcd, 0xfe, 0xa8, 0xb1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MonitorServiceClient is the client API for MonitorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorServiceClient interface {
	Scala(ctx context.Context, in *MonitorScalaRequest, opts ...grpc.CallOption) (*TioReply, error)
}

type monitorServiceClient struct {
	cc *grpc.ClientConn
}

func NewMonitorServiceClient(cc *grpc.ClientConn) MonitorServiceClient {
	return &monitorServiceClient{cc}
}

func (c *monitorServiceClient) Scala(ctx context.Context, in *MonitorScalaRequest, opts ...grpc.CallOption) (*TioReply, error) {
	out := new(TioReply)
	err := c.cc.Invoke(ctx, "/MonitorService/Scala", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorServiceServer is the server API for MonitorService service.
type MonitorServiceServer interface {
	Scala(context.Context, *MonitorScalaRequest) (*TioReply, error)
}

// UnimplementedMonitorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMonitorServiceServer struct {
}

func (*UnimplementedMonitorServiceServer) Scala(ctx context.Context, req *MonitorScalaRequest) (*TioReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scala not implemented")
}

func RegisterMonitorServiceServer(s *grpc.Server, srv MonitorServiceServer) {
	s.RegisterService(&_MonitorService_serviceDesc, srv)
}

func _MonitorService_Scala_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonitorScalaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).Scala(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MonitorService/Scala",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).Scala(ctx, req.(*MonitorScalaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MonitorService",
	HandlerType: (*MonitorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Scala",
			Handler:    _MonitorService_Scala_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor.proto",
}
