package supernode

import (
	"sync"

	pb "github.com/nathanpotter/go-chord/super_node/protos"
	"golang.org/x/net/context"
)

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
  
	return s.nodes, nil
}
