// package supernode holds all the logic for the supernode service in our system.
// The supernode is responsible for maintaining a list of nodes in the system.
// It handles joining and connecting to the system. It also gives a random node
// to a client in order to read and write from the system.

package supernode

import (
	"sync"

	pb "github.com/nathanpotter/go-chord/super_node/protos"
	"golang.org/x/net/context"
)

// type supernode represents the supernode in our system and holds all required state.
// Concurrent access to the data through the use of a mutex.
type supernode struct {
	nodes *pb.Nodes
	mtx   *sync.Mutex
}

func NewSupernode() *supernode {
	return &supernode{
		nodes: &pb.Nodes{Nodes: make([]*pb.Node, 0, 10)},
		mtx:   &sync.Mutex{},
	}
}

func (s *supernode) Join(ctx context.Context, node *pb.Node) (*pb.Nodes, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

  s.nodes.Nodes = append(s.nodes.Nodes, node)
	return s.nodes, nil
}
