// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room.proto

package mimo_srv_room

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
	Mid                  string   `protobuf:"bytes,1,opt,name=mid,proto3" json:"mid,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{0}
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

func (m *Request) GetMid() string {
	if m != nil {
		return m.Mid
	}
	return ""
}

func (m *Request) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type Response struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{1}
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

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RespMember struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Member               []string `protobuf:"bytes,2,rep,name=member,proto3" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespMember) Reset()         { *m = RespMember{} }
func (m *RespMember) String() string { return proto.CompactTextString(m) }
func (*RespMember) ProtoMessage()    {}
func (*RespMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5fd27dd97284ef4, []int{2}
}

func (m *RespMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespMember.Unmarshal(m, b)
}
func (m *RespMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespMember.Marshal(b, m, deterministic)
}
func (m *RespMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespMember.Merge(m, src)
}
func (m *RespMember) XXX_Size() int {
	return xxx_messageInfo_RespMember.Size(m)
}
func (m *RespMember) XXX_DiscardUnknown() {
	xxx_messageInfo_RespMember.DiscardUnknown(m)
}

var xxx_messageInfo_RespMember proto.InternalMessageInfo

func (m *RespMember) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *RespMember) GetMember() []string {
	if m != nil {
		return m.Member
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "mimo.srv.room.Request")
	proto.RegisterType((*Response)(nil), "mimo.srv.room.Response")
	proto.RegisterType((*RespMember)(nil), "mimo.srv.room.RespMember")
}

func init() {
	proto.RegisterFile("room.proto", fileDescriptor_c5fd27dd97284ef4)
}

var fileDescriptor_c5fd27dd97284ef4 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xca, 0xcf, 0xcf,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcd, 0xcd, 0xcc, 0xcd, 0xd7, 0x2b, 0x2e, 0x2a,
	0xd3, 0x03, 0x09, 0x2a, 0xe9, 0x72, 0xb1, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09,
	0x70, 0x31, 0xe7, 0x66, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x20, 0x91,
	0xd2, 0xcc, 0x14, 0x09, 0x26, 0x88, 0x48, 0x69, 0x66, 0x8a, 0x92, 0x12, 0x17, 0x47, 0x50, 0x6a,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18, 0x17, 0x5b, 0x71, 0x49, 0x62, 0x49, 0x69, 0x31,
	0x54, 0x0b, 0x94, 0xa7, 0x64, 0xc3, 0xc5, 0x05, 0x52, 0xe3, 0x9b, 0x9a, 0x9b, 0x94, 0x5a, 0x84,
	0x4b, 0x15, 0x48, 0x3c, 0x17, 0xac, 0x42, 0x82, 0x49, 0x81, 0x19, 0x24, 0x0e, 0xe1, 0x19, 0xed,
	0x61, 0xe4, 0x62, 0x09, 0xca, 0xcf, 0xcf, 0x15, 0xb2, 0xe4, 0x62, 0xf1, 0xca, 0xcf, 0xcc, 0x13,
	0x12, 0xd3, 0x43, 0x71, 0xb1, 0x1e, 0xd4, 0xb9, 0x52, 0xe2, 0x18, 0xe2, 0x10, 0x77, 0x29, 0x31,
	0x08, 0x59, 0x71, 0xb1, 0xfa, 0xa4, 0x26, 0x96, 0xa5, 0x92, 0xa3, 0xd7, 0x96, 0x8b, 0x0d, 0xe6,
	0x72, 0x1c, 0x9a, 0x25, 0xb1, 0x68, 0x86, 0x68, 0x51, 0x62, 0x48, 0x62, 0x03, 0x87, 0xb2, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x51, 0x5d, 0xe1, 0x9e, 0x73, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RoomClient is the client API for Room service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoomClient interface {
	Join(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Leave(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Member(ctx context.Context, in *Request, opts ...grpc.CallOption) (*RespMember, error)
}

type roomClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomClient(cc grpc.ClientConnInterface) RoomClient {
	return &roomClient{cc}
}

func (c *roomClient) Join(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mimo.srv.room.Room/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) Leave(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mimo.srv.room.Room/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomClient) Member(ctx context.Context, in *Request, opts ...grpc.CallOption) (*RespMember, error) {
	out := new(RespMember)
	err := c.cc.Invoke(ctx, "/mimo.srv.room.Room/Member", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServer is the server API for Room service.
type RoomServer interface {
	Join(context.Context, *Request) (*Response, error)
	Leave(context.Context, *Request) (*Response, error)
	Member(context.Context, *Request) (*RespMember, error)
}

// UnimplementedRoomServer can be embedded to have forward compatible implementations.
type UnimplementedRoomServer struct {
}

func (*UnimplementedRoomServer) Join(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedRoomServer) Leave(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (*UnimplementedRoomServer) Member(ctx context.Context, req *Request) (*RespMember, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Member not implemented")
}

func RegisterRoomServer(s *grpc.Server, srv RoomServer) {
	s.RegisterService(&_Room_serviceDesc, srv)
}

func _Room_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mimo.srv.room.Room/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).Join(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mimo.srv.room.Room/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).Leave(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Room_Member_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServer).Member(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mimo.srv.room.Room/Member",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServer).Member(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Room_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mimo.srv.room.Room",
	HandlerType: (*RoomServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Room_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Room_Leave_Handler,
		},
		{
			MethodName: "Member",
			Handler:    _Room_Member_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "room.proto",
}
