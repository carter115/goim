// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello/hello.proto

package mimo_srv_hello

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

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_49e10c42a6c052d6, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_49e10c42a6c052d6, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "mimo.srv.hello.Request")
	proto.RegisterType((*Response)(nil), "mimo.srv.hello.Response")
}

func init() {
	proto.RegisterFile("hello/hello.proto", fileDescriptor_49e10c42a6c052d6)
}

var fileDescriptor_49e10c42a6c052d6 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x7c, 0xb9, 0x99, 0xb9, 0xf9,
	0x7a, 0xc5, 0x45, 0x65, 0x7a, 0x60, 0x51, 0x25, 0x59, 0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x30, 0x5b, 0x49, 0x86, 0x8b, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48,
	0x80, 0x8b, 0x39, 0xb7, 0x38, 0x1d, 0x2a, 0x0d, 0x62, 0x1a, 0xb9, 0x71, 0x71, 0x04, 0x27, 0x56,
	0x7a, 0x80, 0x0c, 0x12, 0xb2, 0xe2, 0x62, 0x0e, 0x4e, 0xac, 0x14, 0x12, 0xd7, 0x43, 0xb5, 0x40,
	0x0f, 0x6a, 0xba, 0x94, 0x04, 0xa6, 0x04, 0xc4, 0x5c, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xdb, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x13, 0x76, 0x22, 0xb0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SayHelloClient is the client API for SayHello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SayHelloClient interface {
	Say(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type sayHelloClient struct {
	cc grpc.ClientConnInterface
}

func NewSayHelloClient(cc grpc.ClientConnInterface) SayHelloClient {
	return &sayHelloClient{cc}
}

func (c *sayHelloClient) Say(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mimo.srv.hello.SayHello/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SayHelloServer is the server API for SayHello service.
type SayHelloServer interface {
	Say(context.Context, *Request) (*Response, error)
}

// UnimplementedSayHelloServer can be embedded to have forward compatible implementations.
type UnimplementedSayHelloServer struct {
}

func (*UnimplementedSayHelloServer) Say(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}

func RegisterSayHelloServer(s *grpc.Server, srv SayHelloServer) {
	s.RegisterService(&_SayHello_serviceDesc, srv)
}

func _SayHello_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayHelloServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mimo.srv.hello.SayHello/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayHelloServer).Say(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _SayHello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mimo.srv.hello.SayHello",
	HandlerType: (*SayHelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _SayHello_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello/hello.proto",
}
