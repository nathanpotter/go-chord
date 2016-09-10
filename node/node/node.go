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
	predecessor *pb.Node
	this        *pb.Node
	fingers     []*pb.Node
	mtx         *sync.Mutex
	fs          afero.Fs
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

func (n *node) findMyPredecessor(nodes *pb.Nodes) error {

	if len(nodes.Nodes) == 1 {
		return nil
	}
	n.predecessor = nil

	for _, node := range nodes.Nodes {
		if node.Id == n.this.Id {
			continue
		}
		if n.predecessor == nil && node.Id < n.this.Id {
			n.predecessor = node
		} else if node.Id < n.this.Id && node.Id > n.predecessor.Id {
			n.predecessor = node
		}
	}
	if n.predecessor == nil {
		n.predecessor = findMax(nodes)
	}
	if n.predecessor == nil {
		return NodeNotFoundError
	}
	return nil
}

func findMax(nodes *pb.Nodes) *pb.Node {
	var max *pb.Node

	for _, node := range nodes.Nodes {
		if max == nil {
			max = node
		} else if node.Id > max.Id {
			max = node
		}
	}

	return max
}

func (n *node) UpdateDHT(ctx context.Context, nodes *pb.Nodes) (*pb.Empty, error) {
	log.Println("Received UpdateDHT request", nodes)
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

	err := n.findMyPredecessor(nodes)
	if err != nil {
		return err
	}

	err = n.updateFingers(nodes)
	if err != nil {
		return err
	}

	return nil
}

func (n *node) updateFingers(nodes *pb.Nodes) error {
	if nodes == nil || nodes.Nodes == nil {
		return NilNodesError
	}
	for i := 0; i < m; i++ {
		var val uint64
		if i == 0 {
			val = ((n.this.Id + 1) % hashSpace)
		} else {
			val = ((n.this.Id + (2 << uint64(i-1))) % hashSpace)
		}
		node, err := findSuccessor(val, nodes.Nodes)
		if err != nil {
			return err
		}
		n.fingers[i] = node
	}
	return nil
}

// Finds successor to id in the nodes set. Rolls over to min node if no node is
// between id and upper limit of hashSpace.
func findSuccessor(id uint64, nodes []*pb.Node) (*pb.Node, error) {
	var n *pb.Node
	var min *pb.Node

	if nodes == nil {
		return &pb.Node{}, NilNodesError
	}

	for _, node := range nodes {
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

func (n *node) closestPredecessor(id uint64) *pb.Node {
	for i := (len(n.fingers) - 1); i >= 0; i-- {
		if n.fingers[i].Id < id && n.fingers[i].Id > n.this.Id {
			return n.fingers[i]
		}
	}
	return n.fingers[0]
}

func (n *node) Write(ctx context.Context, file *pb.File) (*pb.Empty, error) {

	if file == nil || file.Name == "" {
		return &pb.Empty{}, NilFileError
	}
	// hash file
	id, err := hashFilename(file.Name)
	if err != nil {
		return &pb.Empty{}, err
	}

	// if file is mine, write and return result
	if n.myFile(id) {
		log.Printf("Writing file to local filesystem: %s\n", file.Name)
		return n.write(file)
	}

	// if not, forward to closest predecessor
	node := n.closestPredecessor(id)
	if err != nil {
		return &pb.Empty{}, err
	}

	conn, err := grpc.Dial((node.Ip + node.Port), grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return &pb.Empty{}, err
	}
	defer conn.Close()
	nodeClient := npb.NewNodeClient(conn)
	log.Printf("Writing: %s, To: %v\n", file.Name, node)
	return nodeClient.Write(context.Background(), file)
}

func (n *node) write(file *pb.File) (*pb.Empty, error) {
	if file == nil || file.Name == "" {
		return &pb.Empty{}, NilFileError
	}
	f, err := n.fs.Create(file.Name)
	if err != nil {
		return &pb.Empty{}, err
	}
	defer f.Close()

	num, err := f.Write(file.Contents)
	if err != nil {
		return &pb.Empty{}, err
	}
	if num != len(file.Contents) {
		return &pb.Empty{}, NotAllWrittenError
	}

	return &pb.Empty{}, nil
}

func (n *node) Read(ctx context.Context, file *pb.File) (*pb.File, error) {

	if file == nil || file.Name == "" {
		return &pb.File{}, NilFileError
	}
	// hash file
	id, err := hashFilename(file.Name)
	if err != nil {
		return &pb.File{}, err
	}

	// if file is mine, read and return result
	if n.myFile(id) {
		log.Printf("Reading file from local filesystem: %s\n", file.Name)
		return n.read(file)
	}

	// if not, forward to closest predecessor
	node := n.closestPredecessor(id)
	if err != nil {
		return &pb.File{}, err
	}

	conn, err := grpc.Dial((node.Ip + node.Port), grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return &pb.File{}, err
	}
	defer conn.Close()
	nodeClient := npb.NewNodeClient(conn)

	log.Printf("Reading: %s, From: %v\n", file.Name, node)
	return nodeClient.Read(context.Background(), file)
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

func (n *node) myFile(id uint64) bool {
	if n.this.Id < n.predecessor.Id {
		thisId := n.this.Id + n.predecessor.Id
		newId := id + n.predecessor.Id
		if id > n.predecessor.Id {
			if id < thisId {
				return true
			}
			return false
		}
		if id < n.this.Id {
			if newId > n.predecessor.Id {
				return true
			}
			return false
		}
	}
	if id <= n.this.Id && id > n.predecessor.Id {
		return true
	}
	return false
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
	if err != nil {
		return err
	}

	// find my own node
	err = n.findMyNode(nodes)
	if err != nil {
		return err
	}

	// update local DHT
	err = n.updateDHT(nodes)
	if err != nil {
		log.Println(err)
		return err
	}

	// then call UpdateDHT on other nodes in system
	var wg sync.WaitGroup

	for _, node := range nodes.Nodes {
		if node.Id == n.this.Id {
			continue
		}
		wg.Add(1)
		go sendUpdateDHT(node, nodes, &wg)
	}

	wg.Wait()
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

func sendUpdateDHT(node *pb.Node, nodes *pb.Nodes, wg *sync.WaitGroup) {

	defer wg.Done()

	conn, err := grpc.Dial((node.Ip + node.Port), grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	n := npb.NewNodeClient(conn)

	_, err = n.UpdateDHT(context.Background(), nodes)
	if err != nil {
		log.Println(err)
	}
}
