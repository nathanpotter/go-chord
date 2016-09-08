package supernode

import (
	"testing"

	"github.com/golang/protobuf/proto"
	pb "github.com/nathanpotter/go-chord/protos/common"
)

var (
	s    *Supernode
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
	if len(s.nodes.Nodes) != 1 {
		t.Errorf("Node not added to supernodes's node list")
	}
	if len(nodes.Nodes) != 1 {
		t.Errorf("Node not added to supernodes's node list")
	}
}

func TestMultiJoin(t *testing.T) {
	otherNode := &pb.Node{Ip: "localhost", Port: ":10000"}

	nodes, err := s.Join(nil, otherNode)
	if err == nil {
		t.Errorf("Should receive busy error when trying to join system")
	}
	if !proto.Equal(nodes, &pb.Nodes{}) {
		t.Errorf("Nodes should be empty when supernode is busy")
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

	n, err = buildId(&pb.Node{Ip: "localhost", Port: ":50001"})
	if err != nil {
		t.Errorf("Should not receive error when building Id from valid node")
	}
	if n.Id != 18 {
		t.Errorf("Id should be 18 after going through buildId")
	}
	// Node with incorrect Port format, add ':' to beginning and hash
	n, err = buildId(&pb.Node{Ip: "localhost", Port: "50001"})
	if err != nil {
		t.Errorf("Should not receive error when building Id from valid node")
	}
	if n.Id != 18 {
		t.Errorf("Should have same Id between sha1.Sum and node")
	}
}

func TestUnique(t *testing.T) {
	n := &pb.Node{Id: 18}
	ok := s.unique(n)
	if ok {
		t.Errorf("Should receive false since Id: 18 is already in the system")
	}

	uniqueN := &pb.Node{Id: 20}
	ok = s.unique(uniqueN)
	if !ok {
		t.Errorf("s.unique should not return false when node.Id is unique in the system")
	}
}
