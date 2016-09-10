// Code generated by protoc-gen-go.
// source: supernode/supernode.proto
// DO NOT EDIT!

/*
Package supernode is a generated protocol buffer package.

It is generated from these files:
	supernode/supernode.proto

It has these top-level messages:
*/
package supernode

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/nathanpotter/go-chord/protos/common"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Supernode service

type SupernodeClient interface {
	Join(ctx context.Context, in *common.Node, opts ...grpc.CallOption) (*common.Nodes, error)
	PostJoin(ctx context.Context, in *common.Node, opts ...grpc.CallOption) (*common.Empty, error)
	GetNode(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*common.Node, error)
}

type supernodeClient struct {
	cc *grpc.ClientConn
}

func NewSupernodeClient(cc *grpc.ClientConn) SupernodeClient {
	return &supernodeClient{cc}
}

func (c *supernodeClient) Join(ctx context.Context, in *common.Node, opts ...grpc.CallOption) (*common.Nodes, error) {
	out := new(common.Nodes)
	err := grpc.Invoke(ctx, "/supernode.Supernode/Join", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supernodeClient) PostJoin(ctx context.Context, in *common.Node, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/supernode.Supernode/PostJoin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supernodeClient) GetNode(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*common.Node, error) {
	out := new(common.Node)
	err := grpc.Invoke(ctx, "/supernode.Supernode/GetNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Supernode service

type SupernodeServer interface {
	Join(context.Context, *common.Node) (*common.Nodes, error)
	PostJoin(context.Context, *common.Node) (*common.Empty, error)
	GetNode(context.Context, *common.Empty) (*common.Node, error)
}

func RegisterSupernodeServer(s *grpc.Server, srv SupernodeServer) {
	s.RegisterService(&_Supernode_serviceDesc, srv)
}

func _Supernode_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupernodeServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supernode.Supernode/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupernodeServer).Join(ctx, req.(*common.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supernode_PostJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupernodeServer).PostJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supernode.Supernode/PostJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupernodeServer).PostJoin(ctx, req.(*common.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supernode_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupernodeServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supernode.Supernode/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupernodeServer).GetNode(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Supernode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "supernode.Supernode",
	HandlerType: (*SupernodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Supernode_Join_Handler,
		},
		{
			MethodName: "PostJoin",
			Handler:    _Supernode_PostJoin_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _Supernode_GetNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("supernode/supernode.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x2e, 0x2d, 0x48,
	0x2d, 0xca, 0xcb, 0x4f, 0x49, 0xd5, 0x87, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38,
	0xe1, 0x02, 0x52, 0xc2, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0xfa, 0x10, 0x0a, 0x22, 0x6f, 0xd4,
	0xce, 0xc8, 0xc5, 0x19, 0x0c, 0x53, 0x22, 0xa4, 0xca, 0xc5, 0xe2, 0x95, 0x9f, 0x99, 0x27, 0xc4,
	0xa3, 0x07, 0x55, 0xe4, 0x07, 0xd2, 0xc8, 0x8b, 0xcc, 0x2b, 0x56, 0x62, 0x10, 0xd2, 0xe4, 0xe2,
	0x08, 0xc8, 0x2f, 0x2e, 0xc1, 0xa7, 0xd4, 0x35, 0xb7, 0xa0, 0xa4, 0x52, 0x89, 0x41, 0x48, 0x83,
	0x8b, 0xdd, 0x3d, 0xb5, 0x04, 0x24, 0x27, 0x84, 0x2a, 0x27, 0x85, 0xa2, 0x51, 0x89, 0x21, 0x89,
	0x0d, 0xec, 0x20, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x43, 0x77, 0x0c, 0xcd, 0x00,
	0x00, 0x00,
}
