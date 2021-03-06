// Code generated by protoc-gen-go. DO NOT EDIT.
// source: process.proto

package process

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

type ProcessRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessRequest) Reset()         { *m = ProcessRequest{} }
func (m *ProcessRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessRequest) ProtoMessage()    {}
func (*ProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0}
}

func (m *ProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessRequest.Unmarshal(m, b)
}
func (m *ProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessRequest.Marshal(b, m, deterministic)
}
func (m *ProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessRequest.Merge(m, src)
}
func (m *ProcessRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessRequest.Size(m)
}
func (m *ProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessRequest proto.InternalMessageInfo

func (m *ProcessRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProcessRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ProcessResponse struct {
	Result               uint64   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessResponse) Reset()         { *m = ProcessResponse{} }
func (m *ProcessResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessResponse) ProtoMessage()    {}
func (*ProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{1}
}

func (m *ProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessResponse.Unmarshal(m, b)
}
func (m *ProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessResponse.Marshal(b, m, deterministic)
}
func (m *ProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessResponse.Merge(m, src)
}
func (m *ProcessResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessResponse.Size(m)
}
func (m *ProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessResponse proto.InternalMessageInfo

func (m *ProcessResponse) GetResult() uint64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*ProcessRequest)(nil), "process.ProcessRequest")
	proto.RegisterType((*ProcessResponse)(nil), "process.ProcessResponse")
}

func init() { proto.RegisterFile("process.proto", fileDescriptor_54c4d0e8c0aaf5c3) }

var fileDescriptor_54c4d0e8c0aaf5c3 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x4f,
	0x4e, 0x2d, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xac, 0xb8,
	0xf8, 0x02, 0x20, 0xcc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x84, 0x8b, 0xb5,
	0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0xc2, 0x51, 0xd2, 0xe4,
	0xe2, 0x87, 0xeb, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe3, 0x62, 0x2b, 0x4a, 0x2d,
	0x2e, 0xcd, 0x29, 0x01, 0x6b, 0x67, 0x09, 0x82, 0xf2, 0x8c, 0x02, 0xe0, 0xd6, 0x04, 0xa7, 0x16,
	0x95, 0x65, 0x26, 0xa7, 0x0a, 0xd9, 0x71, 0xb1, 0x43, 0x45, 0x84, 0xc4, 0xf5, 0x60, 0x8e, 0x43,
	0x75, 0x8a, 0x94, 0x04, 0xa6, 0x04, 0xc4, 0x1e, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x47, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x21, 0xd0, 0x0e, 0x4e, 0xd9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProcessServiceClient is the client API for ProcessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessServiceClient interface {
	Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error)
}

type processServiceClient struct {
	cc *grpc.ClientConn
}

func NewProcessServiceClient(cc *grpc.ClientConn) ProcessServiceClient {
	return &processServiceClient{cc}
}

func (c *processServiceClient) Process(ctx context.Context, in *ProcessRequest, opts ...grpc.CallOption) (*ProcessResponse, error) {
	out := new(ProcessResponse)
	err := c.cc.Invoke(ctx, "/process.ProcessService/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessServiceServer is the server API for ProcessService service.
type ProcessServiceServer interface {
	Process(context.Context, *ProcessRequest) (*ProcessResponse, error)
}

// UnimplementedProcessServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProcessServiceServer struct {
}

func (*UnimplementedProcessServiceServer) Process(ctx context.Context, req *ProcessRequest) (*ProcessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}

func RegisterProcessServiceServer(s *grpc.Server, srv ProcessServiceServer) {
	s.RegisterService(&_ProcessService_serviceDesc, srv)
}

func _ProcessService_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServiceServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/process.ProcessService/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServiceServer).Process(ctx, req.(*ProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProcessService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "process.ProcessService",
	HandlerType: (*ProcessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _ProcessService_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "process.proto",
}
