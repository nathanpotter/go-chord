package supernode

import (
	"testing"

	pb "github.com/nathanpotter/go-chord/super_node/protos"
)

var (
	s   *supernode
	err error
)

func TestNewSupernode(t *testing.T) {
	s = NewSupernode()
	if s == nil {
		t.Fatalf("Supernode unable to be created.")
	}
	if s.nodes == nil {
		t.Fatalf("Nodes slice not initialized in supernode")
	}
}

func TestJoin(t *testing.T) {
	node := &pb.Node{Ip: "localhost", Port: ":50001"}
  
	_, err := s.Join(nil, node)
	if err != nil {
		t.Errorf("Error joining supernode:", err)
	}
}
