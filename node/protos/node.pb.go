// Code generated by protoc-gen-go.
// source: node.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	node.proto

It has these top-level messages:
	File
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

type File struct {
	Name     string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	Contents []byte `protobuf:"bytes,2,opt,name=Contents,json=contents,proto3" json:"Contents,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Nodes struct {
	Nodes []*Nodes_Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *Nodes) Reset()                    { *m = Nodes{} }
func (m *Nodes) String() string            { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()               {}
func (*Nodes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Nodes) GetNodes() []*Nodes_Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Nodes_Node struct {
	Ip   string `protobuf:"bytes,1,opt,name=Ip,json=ip" json:"Ip,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=Port,json=port" json:"Port,omitempty"`
	Id   []byte `protobuf:"bytes,3,opt,name=Id,json=id,proto3" json:"Id,omitempty"`
}

func (m *Nodes_Node) Reset()                    { *m = Nodes_Node{} }
func (m *Nodes_Node) String() string            { return proto.CompactTextString(m) }
func (*Nodes_Node) ProtoMessage()               {}
func (*Nodes_Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*File)(nil), "protos.File")
	proto.RegisterType((*Nodes)(nil), "protos.Nodes")
	proto.RegisterType((*Nodes_Node)(nil), "protos.Nodes.Node")
	proto.RegisterType((*Empty)(nil), "protos.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Node service

type NodeClient interface {
	Write(ctx context.Context, in *File, opts ...grpc.CallOption) (*Empty, error)
	Read(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error)
	UpdateDHT(ctx context.Context, in *Nodes, opts ...grpc.CallOption) (*Empty, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Write(ctx context.Context, in *File, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/protos.node/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Read(ctx context.Context, in *File, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := grpc.Invoke(ctx, "/protos.node/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) UpdateDHT(ctx context.Context, in *Nodes, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/protos.node/UpdateDHT", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	Write(context.Context, *File) (*Empty, error)
	Read(context.Context, *File) (*File, error)
	UpdateDHT(context.Context, *Nodes) (*Empty, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.node/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Write(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.node/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Read(ctx, req.(*File))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_UpdateDHT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nodes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).UpdateDHT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.node/UpdateDHT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).UpdateDHT(ctx, req.(*Nodes))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _Node_Write_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Node_Read_Handler,
		},
		{
			MethodName: "UpdateDHT",
			Handler:    _Node_UpdateDHT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0x3f, 0x4f, 0x04, 0x21,
	0x10, 0xc5, 0x5d, 0x0e, 0xf4, 0x6e, 0xbc, 0xb3, 0x98, 0x6a, 0xb3, 0x95, 0x21, 0xc6, 0x6c, 0x62,
	0x72, 0xc5, 0x99, 0x58, 0xd8, 0xfa, 0x27, 0xda, 0x5c, 0x0c, 0xd1, 0x58, 0xaf, 0x42, 0x41, 0xe2,
	0x02, 0x59, 0x68, 0x2c, 0xfd, 0xe6, 0x0e, 0xec, 0x6d, 0xb1, 0xda, 0xc0, 0x7b, 0xf0, 0x78, 0xbf,
	0x0c, 0x00, 0xce, 0x6b, 0xb3, 0x0d, 0x83, 0x4f, 0x1e, 0x8f, 0xcb, 0x16, 0xe5, 0x0d, 0xf0, 0x47,
	0xfb, 0x65, 0x10, 0x81, 0xef, 0xbb, 0xde, 0xd4, 0xd5, 0x79, 0xd5, 0xae, 0x14, 0x77, 0xa4, 0xb1,
	0x81, 0xe5, 0x9d, 0x77, 0xc9, 0xb8, 0x14, 0x6b, 0x46, 0xe7, 0x6b, 0xb5, 0xfc, 0x3c, 0x78, 0xd9,
	0x83, 0xd8, 0x53, 0x5b, 0xc4, 0x16, 0x44, 0xae, 0x8d, 0xf4, 0x72, 0xd1, 0x9e, 0xee, 0x70, 0xec,
	0x8f, 0xdb, 0x72, 0x5b, 0x56, 0x35, 0x06, 0x9a, 0x5b, 0x42, 0x90, 0xc0, 0x33, 0x60, 0xcf, 0xe1,
	0x00, 0x62, 0x36, 0x64, 0xf4, 0x8b, 0x1f, 0x52, 0x41, 0x10, 0x3a, 0x90, 0x2e, 0x19, 0x5d, 0x2f,
	0x0a, 0x94, 0x59, 0x2d, 0x4f, 0x40, 0x3c, 0xf4, 0x21, 0x7d, 0xef, 0x7e, 0x2a, 0xe0, 0xb9, 0x0e,
	0x2f, 0x41, 0xbc, 0x0f, 0x36, 0x19, 0x5c, 0x4f, 0xc4, 0x3c, 0x47, 0xb3, 0x99, 0x5c, 0x89, 0xcb,
	0x23, 0xbc, 0x00, 0xae, 0x4c, 0xa7, 0xff, 0xc4, 0x66, 0x8e, 0x52, 0x57, 0xb0, 0x7a, 0x0b, 0xba,
	0x4b, 0xe6, 0xfe, 0xe9, 0x15, 0x37, 0xb3, 0x19, 0xfe, 0x55, 0x7e, 0x8c, 0x7f, 0x77, 0xfd, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xa9, 0x2d, 0xc6, 0x1f, 0x50, 0x01, 0x00, 0x00,
}
