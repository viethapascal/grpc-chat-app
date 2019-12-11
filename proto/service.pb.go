// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service.proto

package proto

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type Message struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{1}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Message) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type Connect struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Active               bool     `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{2}
}

func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (m *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(m, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Connect) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type Close struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Close) Reset()         { *m = Close{} }
func (m *Close) String() string { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()    {}
func (*Close) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33392ef2c1961ba, []int{3}
}

func (m *Close) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Close.Unmarshal(m, b)
}
func (m *Close) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Close.Marshal(b, m, deterministic)
}
func (m *Close) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Close.Merge(m, src)
}
func (m *Close) XXX_Size() int {
	return xxx_messageInfo_Close.Size(m)
}
func (m *Close) XXX_DiscardUnknown() {
	xxx_messageInfo_Close.DiscardUnknown(m)
}

var xxx_messageInfo_Close proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*Message)(nil), "proto.Message")
	proto.RegisterType((*Connect)(nil), "proto.Connect")
	proto.RegisterType((*Close)(nil), "proto.Close")
}

func init() { proto.RegisterFile("proto/service.proto", fileDescriptor_c33392ef2c1961ba) }

var fileDescriptor_c33392ef2c1961ba = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x69, 0xed, 0x56, 0xfb, 0x5a, 0x86, 0x44, 0x90, 0x30, 0x04, 0xb5, 0x27, 0x4f, 0xb3,
	0xcc, 0x93, 0xd7, 0xf5, 0xac, 0x87, 0x8a, 0x67, 0x79, 0xb6, 0x0f, 0x09, 0x2c, 0x4d, 0xc9, 0x8b,
	0x03, 0xff, 0x7b, 0x59, 0x9a, 0x56, 0x44, 0x76, 0x2a, 0xdf, 0xfb, 0xfa, 0x7b, 0xbf, 0x24, 0x70,
	0x39, 0x58, 0xe3, 0xcc, 0x03, 0x93, 0x3d, 0xa8, 0x96, 0x36, 0x3e, 0x89, 0x85, 0xff, 0x94, 0x4f,
	0x90, 0xbc, 0x31, 0x59, 0xb1, 0x82, 0x58, 0x75, 0x32, 0xba, 0x8d, 0xee, 0xb3, 0x26, 0x56, 0x9d,
	0xb8, 0x83, 0xa2, 0x53, 0x3c, 0xec, 0xf1, 0xfb, 0xbd, 0x47, 0x4d, 0x32, 0xf6, 0x4d, 0x1e, 0x66,
	0x2f, 0xa8, 0xa9, 0xb4, 0x90, 0x3e, 0x13, 0x33, 0x7e, 0xd2, 0x3f, 0xfa, 0x06, 0x92, 0x2f, 0x26,
	0xeb, 0xa9, 0x7c, 0x9b, 0x8f, 0xca, 0xcd, 0x51, 0xd4, 0xf8, 0x42, 0x48, 0x48, 0xf5, 0xc8, 0xca,
	0x33, 0x4f, 0x4d, 0x51, 0x5c, 0x43, 0xe6, 0x94, 0x26, 0x76, 0xa8, 0x07, 0x99, 0xf8, 0xee, 0x77,
	0x50, 0xee, 0x20, 0xad, 0x4d, 0xdf, 0x53, 0xeb, 0x66, 0x47, 0x74, 0xca, 0x71, 0x05, 0x4b, 0x6c,
	0x9d, 0x3a, 0x8c, 0x87, 0x3f, 0x6f, 0x42, 0x2a, 0x53, 0x58, 0xd4, 0x7b, 0xc3, 0xb4, 0x35, 0x90,
	0xed, 0xac, 0xc1, 0xae, 0x45, 0x76, 0xa2, 0x82, 0xa2, 0xb6, 0x84, 0x8e, 0x5e, 0x9d, 0x25, 0xd4,
	0x62, 0x15, 0x16, 0x06, 0xdd, 0x7a, 0xca, 0xe1, 0xca, 0x55, 0x24, 0x2a, 0xb8, 0x98, 0xf1, 0xf9,
	0x21, 0xfe, 0xfe, 0xb5, 0x2e, 0xa6, 0x2d, 0x47, 0xe1, 0xc7, 0xd2, 0x87, 0xc7, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x88, 0x46, 0x63, 0x1a, 0x91, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BroadcastClient is the client API for Broadcast service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BroadcastClient interface {
	CreateStream(ctx context.Context, in *Connect, opts ...grpc.CallOption) (Broadcast_CreateStreamClient, error)
	BroadcastMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Close, error)
}

type broadcastClient struct {
	cc *grpc.ClientConn
}

func NewBroadcastClient(cc *grpc.ClientConn) BroadcastClient {
	return &broadcastClient{cc}
}

func (c *broadcastClient) CreateStream(ctx context.Context, in *Connect, opts ...grpc.CallOption) (Broadcast_CreateStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Broadcast_serviceDesc.Streams[0], "/proto.Broadcast/CreateStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &broadcastCreateStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Broadcast_CreateStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type broadcastCreateStreamClient struct {
	grpc.ClientStream
}

func (x *broadcastCreateStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *broadcastClient) BroadcastMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Close, error) {
	out := new(Close)
	err := c.cc.Invoke(ctx, "/proto.Broadcast/BroadcastMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastServer is the server API for Broadcast service.
type BroadcastServer interface {
	CreateStream(*Connect, Broadcast_CreateStreamServer) error
	BroadcastMessage(context.Context, *Message) (*Close, error)
}

// UnimplementedBroadcastServer can be embedded to have forward compatible implementations.
type UnimplementedBroadcastServer struct {
}

func (*UnimplementedBroadcastServer) CreateStream(req *Connect, srv Broadcast_CreateStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateStream not implemented")
}
func (*UnimplementedBroadcastServer) BroadcastMessage(ctx context.Context, req *Message) (*Close, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastMessage not implemented")
}

func RegisterBroadcastServer(s *grpc.Server, srv BroadcastServer) {
	s.RegisterService(&_Broadcast_serviceDesc, srv)
}

func _Broadcast_CreateStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Connect)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BroadcastServer).CreateStream(m, &broadcastCreateStreamServer{stream})
}

type Broadcast_CreateStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type broadcastCreateStreamServer struct {
	grpc.ServerStream
}

func (x *broadcastCreateStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Broadcast_BroadcastMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).BroadcastMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Broadcast/BroadcastMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).BroadcastMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Broadcast_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Broadcast",
	HandlerType: (*BroadcastServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadcastMessage",
			Handler:    _Broadcast_BroadcastMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateStream",
			Handler:       _Broadcast_CreateStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}
