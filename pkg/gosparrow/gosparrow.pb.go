// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gosparrow.proto

/*
Package gosparrow is a generated protocol buffer package.

It is generated from these files:
	gosparrow.proto

It has these top-level messages:
	GetNameReq
	GetNameResp
*/
package gosparrow

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// GetNameReq represents the request to retrieve the app name.
type GetNameReq struct {
	Prefix string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
}

func (m *GetNameReq) Reset()                    { *m = GetNameReq{} }
func (m *GetNameReq) String() string            { return proto.CompactTextString(m) }
func (*GetNameReq) ProtoMessage()               {}
func (*GetNameReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetNameReq) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// GetNameResp represents the response to retrieve the app name
type GetNameResp struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetNameResp) Reset()                    { *m = GetNameResp{} }
func (m *GetNameResp) String() string            { return proto.CompactTextString(m) }
func (*GetNameResp) ProtoMessage()               {}
func (*GetNameResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetNameResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetNameReq)(nil), "gosparrow.GetNameReq")
	proto.RegisterType((*GetNameResp)(nil), "gosparrow.GetNameResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Gosparrow service

type GosparrowClient interface {
	// GetName used to get the app name.
	GetName(ctx context.Context, in *GetNameReq, opts ...grpc.CallOption) (*GetNameResp, error)
}

type gosparrowClient struct {
	cc *grpc.ClientConn
}

func NewGosparrowClient(cc *grpc.ClientConn) GosparrowClient {
	return &gosparrowClient{cc}
}

func (c *gosparrowClient) GetName(ctx context.Context, in *GetNameReq, opts ...grpc.CallOption) (*GetNameResp, error) {
	out := new(GetNameResp)
	err := grpc.Invoke(ctx, "/gosparrow.Gosparrow/GetName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gosparrow service

type GosparrowServer interface {
	// GetName used to get the app name.
	GetName(context.Context, *GetNameReq) (*GetNameResp, error)
}

func RegisterGosparrowServer(s *grpc.Server, srv GosparrowServer) {
	s.RegisterService(&_Gosparrow_serviceDesc, srv)
}

func _Gosparrow_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GosparrowServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gosparrow.Gosparrow/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GosparrowServer).GetName(ctx, req.(*GetNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gosparrow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gosparrow.Gosparrow",
	HandlerType: (*GosparrowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetName",
			Handler:    _Gosparrow_GetName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gosparrow.proto",
}

func init() { proto.RegisterFile("gosparrow.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcf, 0x2f, 0x2e,
	0x48, 0x2c, 0x2a, 0xca, 0x2f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0xa9, 0x70, 0x71, 0xb9, 0xa7, 0x96, 0xf8, 0x25, 0xe6, 0xa6, 0x06, 0xa5, 0x16, 0x0a, 0x89, 0x71,
	0xb1, 0x15, 0x14, 0xa5, 0xa6, 0x65, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79,
	0x4a, 0x8a, 0x5c, 0xdc, 0x70, 0x55, 0xc5, 0x05, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9,
	0x50, 0x45, 0x60, 0xb6, 0x91, 0x3b, 0x17, 0xa7, 0x3b, 0xcc, 0x54, 0x21, 0x2b, 0x2e, 0x76, 0xa8,
	0x7a, 0x21, 0x51, 0x3d, 0x84, 0xed, 0x08, 0x9b, 0xa4, 0xc4, 0xb0, 0x09, 0x17, 0x17, 0x28, 0x31,
	0x38, 0xa9, 0x72, 0x09, 0x66, 0xe6, 0xeb, 0xa5, 0x17, 0x15, 0x24, 0x23, 0x94, 0x38, 0xf1, 0xc1,
	0xcd, 0x0e, 0x00, 0xf9, 0x20, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x15, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0x0c, 0x6e, 0xd9, 0xdd, 0x00, 0x00, 0x00,
}