// Code generated by protoc-gen-go.
// source: logic.proto
// DO NOT EDIT!

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LoginReq struct {
	UID        int64  `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	AccessAddr string `protobuf:"bytes,3,opt,name=accessAddr" json:"accessAddr,omitempty"`
}

func (m *LoginReq) Reset()                    { *m = LoginReq{} }
func (m *LoginReq) String() string            { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()               {}
func (*LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *LoginReq) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *LoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginReq) GetAccessAddr() string {
	if m != nil {
		return m.AccessAddr
	}
	return ""
}

type LoginRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *LoginRes) Reset()                    { *m = LoginRes{} }
func (m *LoginRes) String() string            { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()               {}
func (*LoginRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *LoginRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *LoginRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

type PingReq struct {
	UID int64 `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
}

func (m *PingReq) Reset()                    { *m = PingReq{} }
func (m *PingReq) String() string            { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()               {}
func (*PingReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *PingReq) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type PingRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *PingRes) Reset()                    { *m = PingRes{} }
func (m *PingRes) String() string            { return proto.CompactTextString(m) }
func (*PingRes) ProtoMessage()               {}
func (*PingRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *PingRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *PingRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

type SendP2PMsgReq struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,3,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
}

func (m *SendP2PMsgReq) Reset()                    { *m = SendP2PMsgReq{} }
func (m *SendP2PMsgReq) String() string            { return proto.CompactTextString(m) }
func (*SendP2PMsgReq) ProtoMessage()               {}
func (*SendP2PMsgReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *SendP2PMsgReq) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *SendP2PMsgReq) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *SendP2PMsgReq) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *SendP2PMsgReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type SendP2PMsgRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *SendP2PMsgRes) Reset()                    { *m = SendP2PMsgRes{} }
func (m *SendP2PMsgRes) String() string            { return proto.CompactTextString(m) }
func (*SendP2PMsgRes) ProtoMessage()               {}
func (*SendP2PMsgRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *SendP2PMsgRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *SendP2PMsgRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

// p2p msg accept ack
type AcceptP2PMsgAckReq struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,3,opt,name=msgID" json:"msgID,omitempty"`
}

func (m *AcceptP2PMsgAckReq) Reset()                    { *m = AcceptP2PMsgAckReq{} }
func (m *AcceptP2PMsgAckReq) String() string            { return proto.CompactTextString(m) }
func (*AcceptP2PMsgAckReq) ProtoMessage()               {}
func (*AcceptP2PMsgAckReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *AcceptP2PMsgAckReq) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *AcceptP2PMsgAckReq) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *AcceptP2PMsgAckReq) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

type AcceptP2PMsgAckRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *AcceptP2PMsgAckRes) Reset()                    { *m = AcceptP2PMsgAckRes{} }
func (m *AcceptP2PMsgAckRes) String() string            { return proto.CompactTextString(m) }
func (*AcceptP2PMsgAckRes) ProtoMessage()               {}
func (*AcceptP2PMsgAckRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *AcceptP2PMsgAckRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *AcceptP2PMsgAckRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

// group msg
type SendGroupMsgReq struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	GroupID   int64  `protobuf:"varint,2,opt,name=groupID" json:"groupID,omitempty"`
	MsgID     string `protobuf:"bytes,3,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=msg" json:"msg,omitempty"`
}

func (m *SendGroupMsgReq) Reset()                    { *m = SendGroupMsgReq{} }
func (m *SendGroupMsgReq) String() string            { return proto.CompactTextString(m) }
func (*SendGroupMsgReq) ProtoMessage()               {}
func (*SendGroupMsgReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *SendGroupMsgReq) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *SendGroupMsgReq) GetGroupID() int64 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *SendGroupMsgReq) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *SendGroupMsgReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type SendGroupMsgRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *SendGroupMsgRes) Reset()                    { *m = SendGroupMsgRes{} }
func (m *SendGroupMsgRes) String() string            { return proto.CompactTextString(m) }
func (*SendGroupMsgRes) ProtoMessage()               {}
func (*SendGroupMsgRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *SendGroupMsgRes) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *SendGroupMsgRes) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "rpc.LoginReq")
	proto.RegisterType((*LoginRes)(nil), "rpc.LoginRes")
	proto.RegisterType((*PingReq)(nil), "rpc.PingReq")
	proto.RegisterType((*PingRes)(nil), "rpc.PingRes")
	proto.RegisterType((*SendP2PMsgReq)(nil), "rpc.SendP2PMsgReq")
	proto.RegisterType((*SendP2PMsgRes)(nil), "rpc.SendP2PMsgRes")
	proto.RegisterType((*AcceptP2PMsgAckReq)(nil), "rpc.AcceptP2PMsgAckReq")
	proto.RegisterType((*AcceptP2PMsgAckRes)(nil), "rpc.AcceptP2PMsgAckRes")
	proto.RegisterType((*SendGroupMsgReq)(nil), "rpc.SendGroupMsgReq")
	proto.RegisterType((*SendGroupMsgRes)(nil), "rpc.SendGroupMsgRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LogicRPC service

type LogicRPCClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error)
	SendP2PMsg(ctx context.Context, in *SendP2PMsgReq, opts ...grpc.CallOption) (*SendP2PMsgRes, error)
	AcceptP2PMsgAck(ctx context.Context, in *AcceptP2PMsgAckReq, opts ...grpc.CallOption) (*AcceptP2PMsgAckRes, error)
	SendGroupMsg(ctx context.Context, in *SendGroupMsgReq, opts ...grpc.CallOption) (*SendGroupMsgRes, error)
}

type logicRPCClient struct {
	cc *grpc.ClientConn
}

func NewLogicRPCClient(cc *grpc.ClientConn) LogicRPCClient {
	return &logicRPCClient{cc}
}

func (c *logicRPCClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := grpc.Invoke(ctx, "/rpc.LogicRPC/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicRPCClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error) {
	out := new(PingRes)
	err := grpc.Invoke(ctx, "/rpc.LogicRPC/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicRPCClient) SendP2PMsg(ctx context.Context, in *SendP2PMsgReq, opts ...grpc.CallOption) (*SendP2PMsgRes, error) {
	out := new(SendP2PMsgRes)
	err := grpc.Invoke(ctx, "/rpc.LogicRPC/SendP2PMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicRPCClient) AcceptP2PMsgAck(ctx context.Context, in *AcceptP2PMsgAckReq, opts ...grpc.CallOption) (*AcceptP2PMsgAckRes, error) {
	out := new(AcceptP2PMsgAckRes)
	err := grpc.Invoke(ctx, "/rpc.LogicRPC/AcceptP2PMsgAck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicRPCClient) SendGroupMsg(ctx context.Context, in *SendGroupMsgReq, opts ...grpc.CallOption) (*SendGroupMsgRes, error) {
	out := new(SendGroupMsgRes)
	err := grpc.Invoke(ctx, "/rpc.LogicRPC/SendGroupMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LogicRPC service

type LogicRPCServer interface {
	Login(context.Context, *LoginReq) (*LoginRes, error)
	Ping(context.Context, *PingReq) (*PingRes, error)
	SendP2PMsg(context.Context, *SendP2PMsgReq) (*SendP2PMsgRes, error)
	AcceptP2PMsgAck(context.Context, *AcceptP2PMsgAckReq) (*AcceptP2PMsgAckRes, error)
	SendGroupMsg(context.Context, *SendGroupMsgReq) (*SendGroupMsgRes, error)
}

func RegisterLogicRPCServer(s *grpc.Server, srv LogicRPCServer) {
	s.RegisterService(&_LogicRPC_serviceDesc, srv)
}

func _LogicRPC_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicRPCServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicRPC/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicRPCServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicRPC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicRPCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicRPC/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicRPCServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicRPC_SendP2PMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendP2PMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicRPCServer).SendP2PMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicRPC/SendP2PMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicRPCServer).SendP2PMsg(ctx, req.(*SendP2PMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicRPC_AcceptP2PMsgAck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptP2PMsgAckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicRPCServer).AcceptP2PMsgAck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicRPC/AcceptP2PMsgAck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicRPCServer).AcceptP2PMsgAck(ctx, req.(*AcceptP2PMsgAckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicRPC_SendGroupMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendGroupMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicRPCServer).SendGroupMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LogicRPC/SendGroupMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicRPCServer).SendGroupMsg(ctx, req.(*SendGroupMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogicRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.LogicRPC",
	HandlerType: (*LogicRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LogicRPC_Login_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _LogicRPC_Ping_Handler,
		},
		{
			MethodName: "SendP2PMsg",
			Handler:    _LogicRPC_SendP2PMsg_Handler,
		},
		{
			MethodName: "AcceptP2PMsgAck",
			Handler:    _LogicRPC_AcceptP2PMsgAck_Handler,
		},
		{
			MethodName: "SendGroupMsg",
			Handler:    _LogicRPC_SendGroupMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic.proto",
}

func init() { proto.RegisterFile("logic.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0xdb, 0x6a, 0xc2, 0x40,
	0x10, 0x55, 0xe3, 0xa5, 0x4e, 0x15, 0xcb, 0x20, 0x6d, 0x48, 0x4b, 0x91, 0xa5, 0x50, 0x9f, 0x7c,
	0xb0, 0xd0, 0x97, 0xfa, 0x12, 0xb4, 0x2d, 0x42, 0x0b, 0x21, 0xd2, 0x0f, 0xb0, 0x9b, 0x6d, 0x10,
	0x6b, 0x36, 0xee, 0xc6, 0x5f, 0xe9, 0xf7, 0x96, 0xdd, 0x5c, 0x8c, 0x97, 0x82, 0x04, 0xfa, 0xb6,
	0x67, 0xce, 0x78, 0xe6, 0xec, 0xec, 0x31, 0x70, 0xfe, 0xcd, 0xfd, 0x05, 0x1d, 0x84, 0x82, 0x47,
	0x1c, 0x0d, 0x11, 0x52, 0xe2, 0xc2, 0xd9, 0x1b, 0xf7, 0x17, 0x81, 0xcb, 0xd6, 0x78, 0x01, 0xc6,
	0xc7, 0x74, 0x62, 0x96, 0x7b, 0xe5, 0xbe, 0xe1, 0xaa, 0x23, 0x76, 0xa1, 0x16, 0xf1, 0x25, 0x0b,
	0xcc, 0x4a, 0xaf, 0xdc, 0x6f, 0xba, 0x31, 0xc0, 0x5b, 0x80, 0x39, 0xa5, 0x4c, 0x4a, 0xdb, 0xf3,
	0x84, 0x69, 0x68, 0x2a, 0x57, 0x21, 0xa3, 0x4c, 0x53, 0xa2, 0x09, 0x0d, 0x26, 0xc4, 0x98, 0x7b,
	0x4c, 0xeb, 0xb6, 0xdd, 0x14, 0xe2, 0x25, 0xd4, 0x99, 0x10, 0xb3, 0x48, 0x24, 0xe2, 0x09, 0x22,
	0xd7, 0xd0, 0x70, 0x16, 0x81, 0x7f, 0xd4, 0x10, 0x79, 0x4a, 0xc9, 0x22, 0xca, 0x12, 0xda, 0x33,
	0x16, 0x78, 0xce, 0xd0, 0x79, 0x97, 0x5a, 0xff, 0x06, 0x9a, 0x92, 0x6f, 0x04, 0x65, 0xdb, 0x29,
	0xdb, 0x82, 0x62, 0xa3, 0xb9, 0xf0, 0x59, 0xa4, 0xd8, 0x4a, 0xcc, 0x66, 0x05, 0xb5, 0x9a, 0x95,
	0xf4, 0xa7, 0x93, 0xe4, 0xfe, 0x31, 0x50, 0x8e, 0x57, 0xd2, 0x37, 0xab, 0xba, 0xa6, 0x8e, 0xc4,
	0xde, 0x1d, 0x5a, 0xc4, 0xf7, 0x17, 0xa0, 0x4d, 0x29, 0x0b, 0xa3, 0x58, 0xc4, 0xa6, 0xcb, 0x7f,
	0x31, 0x4f, 0x5e, 0x8e, 0xcc, 0x29, 0xe2, 0x77, 0x0d, 0x1d, 0x75, 0xe5, 0x57, 0xc1, 0x37, 0xe1,
	0x49, 0x9b, 0x36, 0xa1, 0xe1, 0xab, 0xe6, 0xcc, 0x6a, 0x0a, 0x4f, 0xde, 0xf2, 0x78, 0x7f, 0x64,
	0x01, 0xdf, 0xc3, 0x9f, 0x4a, 0x1c, 0x5c, 0xea, 0x3a, 0x63, 0xbc, 0x87, 0x9a, 0x0e, 0x31, 0xb6,
	0x07, 0x22, 0xa4, 0x83, 0xf4, 0x4f, 0x62, 0xed, 0x40, 0x49, 0x4a, 0x78, 0x07, 0x55, 0x15, 0x49,
	0x6c, 0x69, 0x22, 0x89, 0xae, 0x95, 0x47, 0xaa, 0xeb, 0x11, 0x60, 0x1b, 0x03, 0x44, 0xcd, 0xee,
	0x84, 0xd1, 0x3a, 0xac, 0xa9, 0xdf, 0x3d, 0x43, 0x67, 0xef, 0x4d, 0xf0, 0x4a, 0x37, 0x1e, 0x26,
	0xc2, 0xfa, 0x83, 0x50, 0x32, 0x23, 0x68, 0xe5, 0xf7, 0x83, 0xdd, 0x6c, 0x58, 0xee, 0x95, 0xac,
	0x63, 0x55, 0x49, 0x4a, 0x9f, 0x75, 0xfd, 0xc1, 0x78, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x83,
	0xe5, 0x67, 0xc1, 0x3f, 0x04, 0x00, 0x00,
}
