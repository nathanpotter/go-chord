// Code generated by protoc-gen-go.
// source: supernode.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	supernode.proto

It has these top-level messages:
	Node
	Nodes
	Empty
*/
package protos

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

type Node struct {
	Ip   string `protobuf:"bytes,1,opt,name=Ip,json=ip" json:"Ip,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=Port,json=port" json:"Port,omitempty"`
	Id   []byte `protobuf:"bytes,3,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Nodes struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *Nodes) Reset()                    { *m = Nodes{} }
func (m *Nodes) String() string            { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()               {}
func (*Nodes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Nodes) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Node)(nil), "protos.Node")
	proto.RegisterType((*Nodes)(nil), "protos.Nodes")
	proto.RegisterType((*Empty)(nil), "protos.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Supernode service

type SupernodeClient interface {
	Join(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Nodes, error)
	PostJoin(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Empty, error)
	GetNode(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Node, error)
}

type supernodeClient struct {
	cc *grpc.ClientConn
}

func NewSupernodeClient(cc *grpc.ClientConn) SupernodeClient {
	return &supernodeClient{cc}
}

func (c *supernodeClient) Join(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Nodes, error) {
	out := new(Nodes)
	err := grpc.Invoke(ctx, "/protos.Supernode/Join", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supernodeClient) PostJoin(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/protos.Supernode/PostJoin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supernodeClient) GetNode(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := grpc.Invoke(ctx, "/protos.Supernode/GetNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Supernode service

type SupernodeServer interface {
	Join(context.Context, *Node) (*Nodes, error)
	PostJoin(context.Context, *Node) (*Empty, error)
	GetNode(context.Context, *Empty) (*Node, error)
}

func RegisterSupernodeServer(s *grpc.Server, srv SupernodeServer) {
	s.RegisterService(&_Supernode_serviceDesc, srv)
}

func _Supernode_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupernodeServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Supernode/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupernodeServer).Join(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supernode_PostJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupernodeServer).PostJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Supernode/PostJoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupernodeServer).PostJoin(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supernode_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupernodeServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Supernode/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupernodeServer).GetNode(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Supernode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Supernode",
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

func init() { proto.RegisterFile("supernode.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2e, 0x2d, 0x48,
	0x2d, 0xca, 0xcb, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5,
	0x4a, 0x56, 0x5c, 0x2c, 0x7e, 0x40, 0x51, 0x21, 0x3e, 0x2e, 0x26, 0xcf, 0x02, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x02, 0x21, 0x21, 0x2e, 0x96, 0x80, 0xfc, 0xa2, 0x12, 0x09,
	0x26, 0xb0, 0x08, 0x4b, 0x01, 0x90, 0x0d, 0x56, 0x93, 0x22, 0xc1, 0x0c, 0x14, 0xe1, 0x01, 0xaa,
	0x49, 0x51, 0xd2, 0xe6, 0x62, 0x05, 0xe9, 0x2d, 0x16, 0x52, 0xe2, 0x62, 0x05, 0x19, 0x5d, 0x0c,
	0xd4, 0xcf, 0xac, 0xc1, 0x6d, 0xc4, 0x03, 0xb1, 0xa3, 0x58, 0x0f, 0x24, 0x1b, 0x04, 0x91, 0x52,
	0x62, 0xe7, 0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x34, 0x6a, 0x67, 0xe4, 0xe2, 0x0c, 0x86, 0xb9,
	0x46, 0x48, 0x95, 0x8b, 0xc5, 0x2b, 0x3f, 0x33, 0x4f, 0x08, 0x45, 0x8f, 0x14, 0x2f, 0x32, 0xaf,
	0x58, 0x89, 0x41, 0x48, 0x93, 0x8b, 0x23, 0x20, 0xbf, 0xb8, 0x04, 0x9f, 0x52, 0xb0, 0xe9, 0x40,
	0xa5, 0x1a, 0x5c, 0xec, 0xee, 0xa9, 0x25, 0x60, 0x4f, 0xa1, 0xca, 0x49, 0xa1, 0x68, 0x54, 0x62,
	0x48, 0x82, 0x84, 0x81, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xd9, 0xd1, 0xe5, 0x1d, 0x01,
	0x00, 0x00,
}