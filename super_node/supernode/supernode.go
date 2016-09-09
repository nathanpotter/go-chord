// package Supernode holds all the logic for the Supernode service in our system.
// The Supernode is responsible for maintaining a list of nodes in the system.
// It handles joining and connecting to the system. It also gives a random node
// to a client in order to read and write from the system.

package supernode

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"log"
	"math/rand"
	"strings"
	"sync"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/proto"
	pb "github.com/nathanpotter/go-chord/protos/common"
)

const (
	// m represents the size of the hash space, 2^m
	m = 6
	// hashSpace represents the size of the hash space for the nodes in the system
	hashSpace = 2 << (m - 1)
)

var (
	NoNodesError    = errors.New("There are no nodes in the system currently")
	BusyError       = errors.New("Busy connecting a node to the system")
	WrongNodeError  = errors.New("Incorrect node calling PostJoin")
	NilNodeError    = errors.New("Invalid formatting, Nil node values")
	SystemFullError = errors.New("System is full, unable to join")
)

// type Supernode represents the Supernode in our system and holds all required state.
// Concurrent access to the data through the use of a mutex.
type Supernode struct {
	nodes    *pb.Nodes
	mtx      *sync.Mutex
	busyWith *pb.Node
}

func NewSupernode() *Supernode {
	return &Supernode{
		nodes: &pb.Nodes{Nodes: make([]*pb.Node, 0, 10)},
		mtx:   &sync.Mutex{},
	}
}

// Returns a random node from the Supernode nodes list for a client to call Read/Write on
func (s *Supernode) GetNode(ctx context.Context, empty *pb.Empty) (*pb.Node, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	// no nodes in the system
	if len(s.nodes.Nodes) == 0 {
		return nil, NoNodesError
	}
	// Supernode is busy
	if s.busyWith != nil {
		return nil, BusyError
	}

	// return random node
	return s.getRandomNode()
}

func (s *Supernode) Join(ctx context.Context, node *pb.Node) (*pb.Nodes, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if len(s.nodes.Nodes) == hashSpace {
		return &pb.Nodes{}, SystemFullError
	}

	log.Println("Join request", node)

	if s.busyWith != nil {
		return &pb.Nodes{}, BusyError
	}
	node, err := buildId(node)
	if err != nil {
		return &pb.Nodes{}, err
	}
	ok := s.unique(node)
	for !ok {
		log.Println("Node not unique")
		node.Id++
		ok = s.unique(node)
		if ok {
			log.Println("Node is now unique")
		}
	}

	s.nodes.Nodes = append(s.nodes.Nodes, node)
	s.busyWith = node

	return s.nodes, nil
}

func (s *Supernode) PostJoin(ctx context.Context, node *pb.Node) (*pb.Empty, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	log.Println("PostJoin request", node)

	if !proto.Equal(node, s.busyWith) {
		return &pb.Empty{}, WrongNodeError
	}

	s.busyWith = nil
	return &pb.Empty{}, nil
}

func (s *Supernode) getRandomNode() (*pb.Node, error) {
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

func (s *Supernode) unique(node *pb.Node) bool {
	for _, n := range s.nodes.Nodes {
		if n.Id == node.Id {
			return false
		}
	}
	return true
}
