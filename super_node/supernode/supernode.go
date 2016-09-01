// package supernode holds all the logic for the supernode service in our system.
// The supernode is responsible for maintaining a list of nodes in the system.
// It handles joining and connecting to the system. It also gives a random node
// to a client in order to read and write from the system.

package supernode

import (
	"errors"
	"sync"
  "math/rand"

  "golang.org/x/net/context"

	pb "github.com/nathanpotter/go-chord/super_node/protos"
  "github.com/golang/protobuf/proto"
)

var (
	NoNodesError = errors.New("There are no nodes in the system currently")
	BusyError    = errors.New("Busy connecting a node to the system")
  WrongNodeError = errors.New("Incorrect node calling PostJoin")
)

// type supernode represents the supernode in our system and holds all required state.
// Concurrent access to the data through the use of a mutex.
type supernode struct {
	nodes    *pb.Nodes
	mtx      *sync.Mutex
	busyWith *pb.Node
}

func NewSupernode() *supernode {
	return &supernode{
		nodes: &pb.Nodes{Nodes: make([]*pb.Node, 0, 10)},
		mtx:   &sync.Mutex{},
	}
}

// Returns a random node from the supernode nodes list for a client to call Read/Write on
func (s *supernode) GetNode(ctx context.Context, empty *pb.Empty) (*pb.Node, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	// no nodes in the system
	if len(s.nodes.Nodes) == 0 {
		return nil, NoNodesError
	}

	// return random node
	return s.getRandomNode()
}

func (s *supernode) Join(ctx context.Context, node *pb.Node) (*pb.Nodes, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.busyWith != nil {
		return nil, BusyError
	}

	s.nodes.Nodes = append(s.nodes.Nodes, node)
	s.busyWith = node

	return s.nodes, nil
}

func (s *supernode) PostJoin(ctx context.Context, node *pb.Node) (*pb.Empty, error) {
  s.mtx.Lock()
  defer s.mtx.Unlock()

  if !proto.Equal(node, s.busyWith) {
    return nil, WrongNodeError
  }

  s.busyWith = nil
  return nil, nil
}

func (s *supernode) getRandomNode() (*pb.Node, error) {
  if len(s.nodes.Nodes) == 0 {
    return nil, NoNodesError
  }
  randNode := s.nodes.Nodes[rand.Intn(len(s.nodes.Nodes))]
  return randNode, nil
}
