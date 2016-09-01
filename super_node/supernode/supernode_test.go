package supernode

import (
	"testing"

  "github.com/golang/protobuf/proto"
	pb "github.com/nathanpotter/go-chord/super_node/protos"
)

var (
	s    *supernode
	node *pb.Node
	err  error
)

func TestNewSupernode(t *testing.T) {
	s = NewSupernode()
	if s == nil {
		t.Fatalf("Supernode unable to be created.")
	}
	if s.nodes == nil {
		t.Fatalf("Nodes slice not initialized in supernode")
	}
	if s.nodes.Nodes == nil {
		t.Fatalf("Nodes inside pb.Nodes not initialized in supernode")
	}
}

func TestColdGet(t *testing.T) {
	n, err := s.GetNode(nil, nil)
	if err == nil {
		t.Errorf("Should receive NoNodesError from GetNode when there are no nodes are in the system")
	}
	if n != nil {
		t.Errorf("Should return nil node when no nodes are in the system")
	}
}

func TestJoin(t *testing.T) {
	node = &pb.Node{Ip: "localhost", Port: ":50001"}

	nodes, err := s.Join(nil, node)
	if err != nil {
		t.Errorf("Error joining supernode:", err)
	}
	if len(nodes.Nodes) == 0 {
		t.Errorf("Node not added to supernodes's node list")
	}
}

func TestMultiJoin(t *testing.T) {
	otherNode := &pb.Node{Ip: "localhost", Port: ":10000"}

	nodes, err := s.Join(nil, otherNode)
	if err == nil {
		t.Errorf("Should receive busy error when trying to join system")
	}
  if nodes != nil {
    t.Errorf("Nodes should be nil when supernode is busy")
  }
}

func TestNotSamePostJoin(t *testing.T) {
  wrongNode := &pb.Node{Ip: "badIpAddress", Port: ":5050"}

  _, err = s.PostJoin(nil, wrongNode)
  if err == nil {
    t.Errorf("Should receive WrongNodeError when calling PostJoin with incorrect node")
  }
}

func TestGoodPostJoin(t *testing.T) {
  _, err = s.PostJoin(nil, node)
  if err != nil {
    t.Errorf("Should not have error when calling PostJoin with correct node")
  }
}

func TestWarmGetNode(t *testing.T) {
  n, err := s.GetNode(nil, nil)
  if err != nil {
    t.Errorf("Should not have error when calling GetNode and there is a node in the system")
  }
  // only 1 node in system, n should be equal to node
  if !proto.Equal(n, node) {
    t.Errorf("Should receive valid node from GetNode")
  }
}
