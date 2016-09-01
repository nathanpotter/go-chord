// package node represents a node in the system. It maintains a DHT of other
// nodes in the system, and a client can read and write to the system through
// any node. Reads and Writes of the system are handled recursively.

package node

import (
  pb "github.com/nathanpotter/go-chord/node/protos"
)

type node struct {
  this *pb.Nodes_Node
}

func NewNode() *node {
  return &node{
    this: &pb.Nodes_Node{},
  }
}
