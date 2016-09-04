// package supernode holds all the logic for the supernode service in our system.
// The supernode is responsible for maintaining a list of nodes in the system.
// It handles joining and connecting to the system. It also gives a random node
// to a client in order to read and write from the system.

package supernode

import (
	"errors"
	"sync"
  "math/rand"
	"crypto/sha1"
	"strings"
	"bytes"
	"encoding/binary"

  "golang.org/x/net/context"

	pb "github.com/nathanpotter/go-chord/super_node/protos"
  "github.com/golang/protobuf/proto"
)

const (
	// m represents the size of the hash space, 2^m
	m = 6
	// hashSpace represents the size of the hash space for the nodes in the system
	hashSpace = 2 << (m-1)
)

var (
	NoNodesError = errors.New("There are no nodes in the system currently")
	BusyError    = errors.New("Busy connecting a node to the system")
  WrongNodeError = errors.New("Incorrect node calling PostJoin")
	NilNodeError = errors.New("Invalid formatting, Nil node values")
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
	// supernode is busy
	if s.busyWith != nil {
		return nil, BusyError
	}

	// return random node
	return s.getRandomNode()
}

func (s *supernode) Join(ctx context.Context, node *pb.Node) (*pb.Nodes, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	// TODO: make sure < 2^m nodes

	if s.busyWith != nil {
		return nil, BusyError
	}
	node, err = buildId(node)
	if err != nil {
		return nil, err
	}
	// TODO: validate uniqueness

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

func buildId(node *pb.Node) (*pb.Node, error) {
	if node.Ip == "" || node.Port == "" {
		return nil, NilNodeError
	}
	if !strings.HasPrefix(node.Port, ":") {
		node.Port = ":" + node.Port
	}
	byteArr := sha1.Sum([]byte(node.Ip + node.Port))
	result, err := putInHashSpace(byteArr[:])
	if err != nil {
		return nil, err
	}

	node.Id = result
	return node, nil
}

func putInHashSpace(b []byte) (uint64, error) {
	bReader := bytes.NewReader(b)
	result, err := binary.ReadUvarint(bReader)
	if err != nil {
		return 0, err
	}
	return result % hashSpace, nil
}
