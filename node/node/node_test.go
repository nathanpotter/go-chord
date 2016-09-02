package node

import (
  "testing"

  pb "github.com/nathanpotter/go-chord/node/protos"
  "github.com/golang/protobuf/proto"
)

var (
  n *node
  nodes *pb.Nodes
)

func TestNewNode(t *testing.T) {
  n = NewNode("127.0.0.1", "5050")
  if n == nil {
    t.Fatalf("NewNode should not return nil")
  }
  if n.this == nil {
    t.Fatalf("n.this should not be nil after returning from NewNode")
  }
}

func TestFindMyNode(t *testing.T) {
  nodes := &pb.Nodes{
    Nodes: []*pb.Nodes_Node{
      &pb.Nodes_Node{Ip: "192.76.91.5", Port:":9000", Id: 3},
      &pb.Nodes_Node{Ip: "127.0.0.1", Port:":5050", Id: 28},
      &pb.Nodes_Node{Ip: "176.80.20.18", Port:":9090", Id: 41},
      &pb.Nodes_Node{Ip: "176.58.21.37", Port:":8000", Id: 19},
      &pb.Nodes_Node{Ip: "192.54.38.76", Port:":6000", Id: 56},
    },
  }
  ns := nodes.GetNodes()
  myNode := ns[1]
  n.this = n.findMyNode(nodes)
  if !proto.Equal(n.this, myNode) {
    t.Errorf("Should replace n.this with correct node which has Id", n.this, myNode)
  }
}

func TestUpdateDHT(t *testing.T) {
  
}
