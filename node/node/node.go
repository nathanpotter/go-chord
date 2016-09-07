// package node represents a node in the system. It maintains a DHT of other
// nodes in the system, and a client can read and write to the system through
// any node. Reads and Writes of the system are handled recursively.

package node

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"

	pb "github.com/nathanpotter/go-chord/protos/common"
	npb "github.com/nathanpotter/go-chord/protos/node"
	spb "github.com/nathanpotter/go-chord/protos/supernode"
	"github.com/nathanpotter/go-chord/super_node/supernode"
	"github.com/spf13/afero"
	"google.golang.org/grpc"
)

var (
	NilNameError       = errors.New("Cannot hash nil name")
	NilNodesError      = errors.New("Invalid argument, nodes cannot be nil")
	NodeNotFoundError  = errors.New("Unable to find local node in nodes list")
	NilFileError       = errors.New("Invalid argument, file cannot be nil")
	NotAllWrittenError = errors.New("Not all file contents written.")
)

const (
	m         = 6
	hashSpace = 2 << (m - 1) // 2^6 nodes allowed in the system
)

type node struct {
	this    *pb.Node
	fingers []*pb.Node
	mtx     *sync.Mutex
	fs      afero.Fs
}

func NewNode(ip, port string) *node {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return &node{
		this:    &pb.Node{Ip: ip, Port: port},
		fingers: make([]*pb.Node, m, m),
		mtx:     &sync.Mutex{},
		fs:      afero.NewMemMapFs(),
	}
}

func (n *node) findMyNode(nodes *pb.Nodes) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	for _, node := range nodes.Nodes {
		if node.Ip == n.this.Ip && node.Port == n.this.Port {
			n.this = node
			return nil
		}
	}
	return NodeNotFoundError
}

func (n *node) UpdateDHT(ctx context.Context, nodes *pb.Nodes) (*pb.Empty, error) {
	//log.Println("Received UpdateDHT request", nodes)
	err := n.updateDHT(nodes)
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}

func (n *node) updateDHT(nodes *pb.Nodes) error {
	if nodes == nil || nodes.Nodes == nil {
		return NilNodesError
	}

	n.mtx.Lock()
	defer n.mtx.Unlock()

	for i := 0; i < m; i++ {
		var val uint64
		if i == 0 {
			val = ((n.this.Id + 1) % hashSpace)
		} else {
			val = ((n.this.Id + (2 << uint64((i - 1)))) % hashSpace)
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
func findSuccessor(id uint64, nodes *pb.Nodes) (*pb.Node, error) {
	var n *pb.Node
	var min *pb.Node

	if nodes == nil || nodes.Nodes == nil {
		return &pb.Node{}, NilNodesError
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

	return &pb.Empty{}, nil
	// check if file is mine
	// write file if it is
	// else forward to closest preceding node
}

func (n *node) write(file *pb.File) error {
	if file == nil || file.Name == "" {
		return NilFileError
	}
	f, err := n.fs.Create(file.Name)
	if err != nil {
		return err
	}
	defer f.Close()

	num, err := f.Write(file.Contents)
	if err != nil {
		return err
	}
	if num != len(file.Contents) {
		return NotAllWrittenError
	}

	return nil
}

func (n *node) Read(ctx context.Context, file *pb.File) (*pb.File, error) {
	// hash file
	// is my file
	// if yes, read and return result
	// else, forward to closest preceding
	return &pb.File{}, nil
}

func (n *node) read(file *pb.File) (*pb.File, error) {
	if file == nil || file.Name == "" {
		return &pb.File{}, NilFileError
	}

	f, err := n.fs.Open(file.Name)
	if err != nil {
		return &pb.File{}, err
	}

	info, err := f.Stat()
	if err != nil {
		return &pb.File{}, err
	}
	size := info.Size()

	b := make([]byte, size)
	_, err = f.Read(b)
	if err != nil {
		return &pb.File{}, err
	}
	file.Contents = b

	return file, nil
}

func (n *node) Join(ip string, port string) error {
	// connect to supernode
	conn, err := grpc.Dial((ip + port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	s := spb.NewSupernodeClient(conn)

	nodes, err := s.Join(context.Background(), n.this)

	// if busy error, wait
	for err == supernode.BusyError {
		time.Sleep(1 * time.Second)
		nodes, err = s.Join(context.Background(), n.this)
	}

	err = n.findMyNode(nodes)
	if err != nil {
		return err
	}

	// when nodes are received, updateDHT
	err = n.updateDHT(nodes)
	if err != nil {
		log.Println(err)
		return err
	}

	// then call UpdateDHT on other nodes in system
	// TODO: parallelize this
	for _, node := range nodes.Nodes {
		if node.Id == n.this.Id {
			continue
		}
		conn, err := grpc.Dial((node.Ip + node.Port), grpc.WithInsecure())
		if err != nil {
			log.Println(err)
			return err
		}
		defer conn.Close()
		n := npb.NewNodeClient(conn)

		_, err = n.UpdateDHT(context.Background(), nodes)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	// when finished, call PostJoin to supernode
	_, err = s.PostJoin(context.Background(), n.this)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
	// then return, then main method will serve from this node
}

func hashFilename(word string) (uint64, error) {
	if word == "" {
		return 0, NilNameError
	}
	byteArr := sha1.Sum([]byte(word))
	result, err := putInHashSpace(byteArr[:])
	if err != nil {
		return 0, err
	}
	return result, nil
}

func putInHashSpace(b []byte) (uint64, error) {
	bReader := bytes.NewReader(b)
	result, err := binary.ReadUvarint(bReader)
	if err != nil {
		return 0, err
	}
	return result % hashSpace, nil
}
