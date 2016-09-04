// package node represents a node in the system. It maintains a DHT of other
// nodes in the system, and a client can read and write to the system through
// any node. Reads and Writes of the system are handled recursively.

package node

import (
	"errors"
	"strings"

	"golang.org/x/net/context"

	pb "github.com/nathanpotter/go-chord/node/protos"
)

var (
	NilNodesError = errors.New("Invalid argument, nodes cannot be nil")
)

const (
	m         = 6
	hashSpace = 2 << (m - 1) // 2^6 nodes allowed in the system
)

type node struct {
	this    *pb.Nodes_Node
	fingers []*pb.Nodes_Node
  mtx      *sync.Mutex
}

func NewNode(ip, port string) *node {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return &node{
		this:    &pb.Nodes_Node{Ip: ip, Port: port},
		fingers: make([]*pb.Nodes_Node, m, m),
	}
}

func (n *node) findMyNode(nodes *pb.Nodes) *pb.Nodes_Node {
  n.mtx.Lock()
  defer n.mtx.Unlock()

	for _, node := range nodes.Nodes {
		if node.Ip == n.this.Ip && node.Port == n.this.Port {
			return node
		}
	}
	return nil
}

func (n *node) UpdateDHT(ctx context.Context, nodes *pb.Nodes) (*pb.Empty, error) {
	err := n.updateDHT(nodes)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (n *node) updateDHT(nodes *pb.Nodes) error {
	if nodes == nil || nodes.Nodes == nil {
		return NilNodesError
	}

  n.mtx.Lock()
  defer n.mtx.Unlock()

  for i := 0; i<m; i++ {
    var val uint64
    if i == 0 {
      val = ((n.this.Id + 1) % hashSpace)
    } else {
      val = ((n.this.Id + (2 << uint64((i-1)))) % hashSpace)
    }
    node, err := findSuccessor(val, nodes)
    if err != nil {
      return err
    }
    n.fingers[i] = node
  }
	return nil
}

// Finds successor to id in the nodes set. Rolls over to min node if no node is
// between id and upper limit of hashSpace.
func findSuccessor(id uint64, nodes *pb.Nodes) (*pb.Nodes_Node, error) {
  var n *pb.Nodes_Node
  var min *pb.Nodes_Node

	if nodes == nil || nodes.Nodes == nil {
		return nil, NilNodesError
	}

  for _, node := range nodes.Nodes {
    if node.Id > id && n == nil {
      n = node
    } else if node.Id > id && node.Id < n.Id {
      n = node
    } else if node.Id < id && min == nil {
      min = node
    } else if node.Id < id && node.Id < min.Id {
      min = node
    } else {
      continue
    }
  }

  if n == nil {
    return min, nil
  }
	return n, nil
}

func (n *node) Write(ctx context.Context, file *pb.File) (*pb.Empty, error) {
  n.mtx.Lock()
  defer n.mtx.Unlock()
  return nil, nil
}

func (n *node) Read(ctx context.Context, file *pb.File) (*pb.File, error) {
  n.mtx.Lock()
  defer n.mtx.Unlock()
  return nil, nil
}
