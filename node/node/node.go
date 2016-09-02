// package node represents a node in the system. It maintains a DHT of other
// nodes in the system, and a client can read and write to the system through
// any node. Reads and Writes of the system are handled recursively.

package node

import (
  "strings"
  pb "github.com/nathanpotter/go-chord/node/protos"
)

type node struct {
  this *pb.Nodes_Node
}

func NewNode(ip, port string) *node {
  if !strings.HasPrefix(port, ":") {
    port = ":" + port
  }
  return &node{
    this: &pb.Nodes_Node{Ip: ip, Port: port},
  }
}

func (n *node) findMyNode(nodes *pb.Nodes) *pb.Nodes_Node {
  for _, node := range nodes.Nodes {
    if node.Ip == n.this.Ip && node.Port == n.this.Port {
      return node
    }
  }
  return nil
}

func (n *node) UpdateDHT(ctx context.Context, nodes *pb.Nodes) (*pb.Empty, error) {
  
}
