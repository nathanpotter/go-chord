package supernode

import (
	"testing"
	"crypto/sha1"
	"bytes"

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
		t.Errorf("Should receive NoNodesError from GetNode when there are no nodes in the system")
	}
	if n != nil {
		t.Errorf("Should return nil node when no nodes are in the system")
	}
}

func TestColdGetRandomNode(t *testing.T) {
  n, err := s.getRandomNode()
  if err == nil {
    t.Errorf("Should receive NoNodesError from getRandomNode when there are no nodes in the system")
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

func TestGetRandomNode(t *testing.T) {
  n, err := s.getRandomNode()
  if err != nil {
    t.Errorf("Should not receive error when there is a node in the system")
  }
  // only 1 node in system, n should be equal to node
  if !proto.Equal(n, node) {
    t.Errorf("Should receive valid node from GetNode")
  }
}

func TestBuildId(t *testing.T) {
	n, err := buildId(&pb.Node{})
	if err == nil {
		t.Errorf("Should return error when trying to build Id from nil Node")
	}
	if n != nil {
		t.Errorf("Should return nil node when trying to build Id from nil node")
	}

	data := []byte("localhost:50001")
	byteArr := sha1.Sum(data)
	result := byteArr[:]

	n, err = buildId(&pb.Node{Ip:"localhost", Port:":50001"})
	if err != nil {
		t.Errorf("Should not receive error when building Id from valid node")
	}
	if !bytes.Equal(n.Id, result) {
		t.Errorf("Should have same Id between sha1.Sum and node")
	}
	// Node with incorrect Port format, add ':' to beginning and hash
	n, err = buildId(&pb.Node{Ip:"localhost", Port:"50001"})
	if err != nil {
		t.Errorf("Should not receive error when building Id from valid node")
	}
	if !bytes.Equal(n.Id, result) {
		t.Errorf("Should have same Id between sha1.Sum and node")
	}
}