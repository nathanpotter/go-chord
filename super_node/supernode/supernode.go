package supernode

import (
  pb "github.com/nathanpotter/go-chord/super_node/protos"
)

type supernode struct {
  nodes []*pb.Node
}

func NewSupernode() *supernode {
  return &supernode{nodes: []*pb.Node{}}
}
