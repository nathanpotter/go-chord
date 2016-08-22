package main

import (
  "log"
  "net"
  "errors"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/nathanpotter/go-chord/supernode"
)

const (
  port = ":50051"

  NilRequestNodeError = errors.New("Cannot compute request with nil NodeInfo")
  InvalidPostJoinRequest = errors.New("Can't invoke PostJoin, supernode isn't currently handling you")
)

type supernode struct {
  nodes []*pb.NodeInfo
  busy bool
  busyWith *pb.NodeInfo
}

func NewSupernode() *supernode {
  return &supernode{}
}

func (s *supernode) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
  if s.busy {
    return &pb.JoinResponse{Nodes: nil, Busy: true}, nil
  }
  if req.GetNode() == nil {
    return nil, NilRequestNodeError
  }
  s.busy = true
  s.busyWith = req.GetNode()

  return &pb.JoinResponse{Nodes: s.nodes}, nil
}

func (s *supernode) PostJoin(ctx context.Context, req *pb.PostJoinRequest) (*pb.Empty, error) {
  if req.GetNode() == nil {
    return nil, NilRequestNodeError
  }
  if !s.busyWith(req.GetNode) {
    return nil, InvalidPostJoinRequest
  }

  s.nodes = append(s.nodes, req.GetNode())
  s.busy = false
  s.busyWith = nil

  return &pb.Empty{}, nil

}

func (s *supernode) GetNode(ctx context.Context, req *pb.Empty) (*pb.NodeResponse, error) {
  
}

func (s *supernode) busyWith(node *pb.NodeInfo) bool {
  if s.busyWith.Ip != node.Ip || s.busyWith.Port != node.Port {
    return false
  }
  return true
}
