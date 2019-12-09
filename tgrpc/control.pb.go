// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control.proto

package tgrpc

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

type JobStatus int32

const (
	JobStatus_BuildSucc   JobStatus = 0
	JobStatus_BuildFailed JobStatus = 1
	JobStatus_BuildIng    JobStatus = 2
)

var JobStatus_name = map[int32]string{
	0: "BuildSucc",
	1: "BuildFailed",
	2: "BuildIng",
}

var JobStatus_value = map[string]int32{
	"BuildSucc":   0,
	"BuildFailed": 1,
	"BuildIng":    2,
}

func (x JobStatus) String() string {
	return proto.EnumName(JobStatus_name, int32(x))
}

func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{0}
}

type BuildStatus struct {
	User                 string    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image                string    `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Api                  string    `protobuf:"bytes,4,opt,name=api,proto3" json:"api,omitempty"`
	Rate                 int32     `protobuf:"varint,5,opt,name=rate,proto3" json:"rate,omitempty"`
	Status               JobStatus `protobuf:"varint,6,opt,name=status,proto3,enum=JobStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BuildStatus) Reset()         { *m = BuildStatus{} }
func (m *BuildStatus) String() string { return proto.CompactTextString(m) }
func (*BuildStatus) ProtoMessage()    {}
func (*BuildStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{0}
}

func (m *BuildStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildStatus.Unmarshal(m, b)
}
func (m *BuildStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildStatus.Marshal(b, m, deterministic)
}
func (m *BuildStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildStatus.Merge(m, src)
}
func (m *BuildStatus) XXX_Size() int {
	return xxx_messageInfo_BuildStatus.Size(m)
}
func (m *BuildStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildStatus.DiscardUnknown(m)
}

var xxx_messageInfo_BuildStatus proto.InternalMessageInfo

func (m *BuildStatus) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *BuildStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BuildStatus) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *BuildStatus) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *BuildStatus) GetRate() int32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *BuildStatus) GetStatus() JobStatus {
	if m != nil {
		return m.Status
	}
	return JobStatus_BuildSucc
}

type BuildReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildReply) Reset()         { *m = BuildReply{} }
func (m *BuildReply) String() string { return proto.CompactTextString(m) }
func (*BuildReply) ProtoMessage()    {}
func (*BuildReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{1}
}

func (m *BuildReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildReply.Unmarshal(m, b)
}
func (m *BuildReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildReply.Marshal(b, m, deterministic)
}
func (m *BuildReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildReply.Merge(m, src)
}
func (m *BuildReply) XXX_Size() int {
	return xxx_messageInfo_BuildReply.Size(m)
}
func (m *BuildReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildReply.DiscardUnknown(m)
}

var xxx_messageInfo_BuildReply proto.InternalMessageInfo

func (m *BuildReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BuildReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("JobStatus", JobStatus_name, JobStatus_value)
	proto.RegisterType((*BuildStatus)(nil), "BuildStatus")
	proto.RegisterType((*BuildReply)(nil), "BuildReply")
}

func init() { proto.RegisterFile("control.proto", fileDescriptor_0c5120591600887d) }

var fileDescriptor_0c5120591600887d = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4f, 0x83, 0x40,
	0x10, 0x85, 0xbb, 0x6d, 0x69, 0x64, 0xda, 0x22, 0x6e, 0x3c, 0x6c, 0x3c, 0x11, 0x4e, 0xc4, 0x44,
	0xa2, 0x78, 0xf2, 0x5a, 0x13, 0x8d, 0x1e, 0x69, 0xbc, 0x78, 0xdb, 0xc2, 0x84, 0x6c, 0x02, 0x2c,
	0xc2, 0xd2, 0xc4, 0x1f, 0xe2, 0xff, 0x35, 0x3b, 0x8b, 0xda, 0xa3, 0xb7, 0x6f, 0x1e, 0xec, 0xbc,
	0x37, 0x33, 0xb0, 0x2d, 0x74, 0x6b, 0x7a, 0x5d, 0xa7, 0x5d, 0xaf, 0x8d, 0x8e, 0xbf, 0x18, 0xac,
	0x77, 0xa3, 0xaa, 0xcb, 0xbd, 0x91, 0x66, 0x1c, 0x38, 0x87, 0xe5, 0x38, 0x60, 0x2f, 0x58, 0xc4,
	0x12, 0x3f, 0x27, 0xb6, 0x5a, 0x2b, 0x1b, 0x14, 0x73, 0xa7, 0x59, 0xe6, 0x97, 0xe0, 0xa9, 0x46,
	0x56, 0x28, 0x16, 0x24, 0xba, 0x82, 0x87, 0xb0, 0x90, 0x9d, 0x12, 0x4b, 0xd2, 0x2c, 0xda, 0xb7,
	0xbd, 0x34, 0x28, 0xbc, 0x88, 0x25, 0x5e, 0x4e, 0xcc, 0x63, 0x58, 0x0d, 0xe4, 0x26, 0x56, 0x11,
	0x4b, 0x82, 0x0c, 0xd2, 0x57, 0x7d, 0x70, 0xfe, 0xf9, 0xf4, 0x25, 0xce, 0x00, 0x28, 0x56, 0x8e,
	0x5d, 0xfd, 0x69, 0xbb, 0x14, 0xba, 0x44, 0x4a, 0xe5, 0xe5, 0xc4, 0xd6, 0xab, 0x19, 0xaa, 0x29,
	0x94, 0xc5, 0xeb, 0x07, 0xf0, 0x7f, 0x1b, 0xf1, 0x2d, 0xf8, 0x6e, 0xae, 0xb1, 0x28, 0xc2, 0x19,
	0x3f, 0x9f, 0xc6, 0x7c, 0x92, 0xaa, 0xc6, 0x32, 0x64, 0x7c, 0x03, 0x67, 0x24, 0xbc, 0xb4, 0x55,
	0x38, 0xcf, 0x3e, 0x20, 0x78, 0x74, 0x7b, 0xd9, 0x63, 0x7f, 0x54, 0x05, 0xf2, 0x5b, 0xb8, 0x78,
	0xeb, 0x4a, 0x69, 0xf0, 0x74, 0x3b, 0x9b, 0xf4, 0xa4, 0xba, 0x5a, 0xa7, 0x7f, 0x11, 0xe3, 0x19,
	0xbf, 0x81, 0xe0, 0x19, 0xcd, 0x7f, 0x7f, 0xdf, 0x85, 0xef, 0x81, 0x51, 0x3a, 0xfd, 0x39, 0xc7,
	0xf1, 0xee, 0xb0, 0xa2, 0x93, 0xdc, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x74, 0x27, 0xc6, 0x12,
	0xa3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControlServiceClient is the client API for ControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControlServiceClient interface {
	UpdateBuildStatus(ctx context.Context, in *BuildStatus, opts ...grpc.CallOption) (*BuildReply, error)
	GetBuildStatus(ctx context.Context, in *BuildStatus, opts ...grpc.CallOption) (*BuildReply, error)
}

type controlServiceClient struct {
	cc *grpc.ClientConn
}

func NewControlServiceClient(cc *grpc.ClientConn) ControlServiceClient {
	return &controlServiceClient{cc}
}

func (c *controlServiceClient) UpdateBuildStatus(ctx context.Context, in *BuildStatus, opts ...grpc.CallOption) (*BuildReply, error) {
	out := new(BuildReply)
	err := c.cc.Invoke(ctx, "/ControlService/UpdateBuildStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) GetBuildStatus(ctx context.Context, in *BuildStatus, opts ...grpc.CallOption) (*BuildReply, error) {
	out := new(BuildReply)
	err := c.cc.Invoke(ctx, "/ControlService/GetBuildStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServiceServer is the server API for ControlService service.
type ControlServiceServer interface {
	UpdateBuildStatus(context.Context, *BuildStatus) (*BuildReply, error)
	GetBuildStatus(context.Context, *BuildStatus) (*BuildReply, error)
}

// UnimplementedControlServiceServer can be embedded to have forward compatible implementations.
type UnimplementedControlServiceServer struct {
}

func (*UnimplementedControlServiceServer) UpdateBuildStatus(ctx context.Context, req *BuildStatus) (*BuildReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBuildStatus not implemented")
}
func (*UnimplementedControlServiceServer) GetBuildStatus(ctx context.Context, req *BuildStatus) (*BuildReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuildStatus not implemented")
}

func RegisterControlServiceServer(s *grpc.Server, srv ControlServiceServer) {
	s.RegisterService(&_ControlService_serviceDesc, srv)
}

func _ControlService_UpdateBuildStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).UpdateBuildStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ControlService/UpdateBuildStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).UpdateBuildStatus(ctx, req.(*BuildStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_GetBuildStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).GetBuildStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ControlService/GetBuildStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).GetBuildStatus(ctx, req.(*BuildStatus))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ControlService",
	HandlerType: (*ControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateBuildStatus",
			Handler:    _ControlService_UpdateBuildStatus_Handler,
		},
		{
			MethodName: "GetBuildStatus",
			Handler:    _ControlService_GetBuildStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}
