package node

import (
	"bytes"
	"testing"

	"github.com/golang/protobuf/proto"
	pb "github.com/nathanpotter/go-chord/protos/common"
)

var (
	n     *node
	nodes = &pb.Nodes{
		Nodes: []*pb.Node{
			&pb.Node{Ip: "192.76.91.5", Port: ":9000", Id: 3},
			&pb.Node{Ip: "127.0.0.1", Port: ":5050", Id: 28},
			&pb.Node{Ip: "176.80.20.18", Port: ":9090", Id: 41},
			&pb.Node{Ip: "176.58.21.37", Port: ":8000", Id: 19},
			&pb.Node{Ip: "192.54.38.76", Port: ":6000", Id: 56},
		},
	}
)

func TestNewNode(t *testing.T) {
	n = NewNode("127.0.0.1", "5050")
	if n == nil {
		t.Fatalf("NewNode should not return nil")
	}
	if n.this == nil {
		t.Fatalf("n.this should not be nil after returning from NewNode")
	}
	if n.fingers == nil {
		t.Fatalf("n.fingers should not be nil after returning from NewNode")
	}
	if n.fs == nil {
		t.Fatalf("n.fs should not be nil after returning from NewNode")
	}
}

func TestFindMyNode(t *testing.T) {

	ns := nodes.GetNodes()
	myNode := ns[1]
	err := n.findMyNode(nodes)
	if err != nil {
		t.Errorf("Should not receive error for findMyNode when valid argument")
	}
	if !proto.Equal(n.this, myNode) {
		t.Errorf("Should replace n.this with correct node which has Id", n.this, myNode)
	}
}

func TestNilNodesUpdateDHT(t *testing.T) {
	// UpdateDHT
	_, err := n.UpdateDHT(nil, nil)
	if err == nil {
		t.Errorf("Should receive NilNodesError when calling UpdateDHT with nil nodes")
	}
	_, err = n.UpdateDHT(nil, &pb.Nodes{})
	if err == nil {
		t.Errorf("Should receive NilNodesError when calling UpdateDHT with nil nodes.Nodes")
	}

	// updateDHT
	err = n.updateDHT(nil)
	if err == nil {
		t.Errorf("Should receive NilNodesError when calling updateDHT with nil nodes")
	}
	err = n.updateDHT(&pb.Nodes{})
	if err == nil {
		t.Errorf("Should receive NilNodesError when calling updateDHT with nil nodes.Nodes")
	}
}

func TestUpdateDHT(t *testing.T) {
	n.UpdateDHT(nil, nodes)
	if n.fingers[0] == nil {
		t.Errorf("node.fingers shouldn't be nil after updateDHT")
	}
	if !proto.Equal(n.fingers[0], nodes.Nodes[2]) {
		t.Errorf("n.fingers[0] should be equal to nodes.Nodes[2]")
	}
	if !proto.Equal(n.fingers[1], nodes.Nodes[2]) {
		t.Errorf("n.fingers[1] should be equal to nodes.Nodes[2]")
	}
	if !proto.Equal(n.fingers[5], nodes.Nodes[0]) {
		t.Errorf("n.fingers[5] should be equal to nodes.Nodes[0]")
	}
}

func TestFindSuccessor(t *testing.T) {
	// nil nodes
	node, err := findSuccessor(0, nil)
	if err == nil {
		t.Errorf("Should receive error from nil nodes argument")
	}
	if !proto.Equal(node, &pb.Node{}) {
		t.Errorf("Should receive empty node when nodes argument is nil")
	}

	node, err = findSuccessor(0, nodes)
	if err != nil {
		t.Errorf("Shouldn't receive error from valid findSuccessor call")
	}
	if node == nil {
		t.Errorf("Shouldn't receive nil node when nodes argument is valid")
	}
	if !proto.Equal(node, nodes.Nodes[0]) {
		t.Errorf("n should equal nodes.Nodes[0] when id = 0")
	}

	node, err = findSuccessor(3, nodes)
	if err != nil {
		t.Errorf("Shouldn't receive error from valid findSuccessor call")
	}
	if !proto.Equal(node, nodes.Nodes[3]) {
		t.Errorf("n should equal nodes.Nodes[3] when id = 3")
	}

	node, err = findSuccessor(25, nodes)
	if err != nil {
		t.Errorf("Shouldn't receive error from valid findSuccessor call")
	}
	if !proto.Equal(node, nodes.Nodes[1]) {
		t.Errorf("n should equal nodes.Nodes[1] when id = 25")
	}

	node, err = findSuccessor(60, nodes)
	if err != nil {
		t.Errorf("Shouldn't receive error from valid findSuccessor call")
	}
	if !proto.Equal(node, nodes.Nodes[0]) {
		t.Errorf("n should equal nodes.Nodes[0] when id = 60")
	}
}

func TestWrite(t *testing.T) {
	file := &pb.File{Name: "Test", Contents: []byte("This is a test file")}

	err := n.write(nil)
	if err == nil {
		t.Errorf("Should receive error when using nil as argument")
	}

	err = n.write(&pb.File{})
	if err == nil {
		t.Errorf("Should receive error when file argument doesn't have name")
	}

	err = n.write(file)
	if err != nil {
		t.Errorf("Should not receive error when writing valid file")
	}

	f, err := n.fs.Open(file.Name)
	if err != nil {
		t.Errorf("Should not receive error when opening a file that was written to the system")
	}
	if f == nil {
		t.Errorf("File should not be nil when opening a file that was written to the system")
	}

	info, err := f.Stat()
	if err != nil {
		t.Errorf("Should not receive error on call to stat")
	}
	size := info.Size()

	b := make([]byte, size)
	_, err = f.Read(b)
	if !bytes.Equal(b, file.Contents) {
		t.Errorf("File contents are not equal to what was written", b)
	}
}

func TestRead(t *testing.T) {
	file := &pb.File{Name: "Test2", Contents: []byte("This is a second test file")}
	nameOnly := &pb.File{Name: "Test2"}

	err := n.write(file)
	if err != nil {
		t.Errorf("Should not receive error from writing valid file")
	}

	_, err = n.read(nil)
	if err == nil {
		t.Errorf("Should receive error from nil file argument")
	}

	_, err = n.read(&pb.File{})
	if err == nil {
		t.Errorf("Should receive error from file with no name")
	}

	_, err = n.read(&pb.File{Name: "DoesNotExist"})
	if err == nil {
		t.Errorf("Should receive error from trying to open and read a non-existent file")
	}

	f, err := n.read(nameOnly)
	if err != nil {
		t.Errorf("Should not receive error from reading valid file")
	}
	if !proto.Equal(f, file) {
		t.Errorf("File read from n.read should be equal to file written")
	}
}

func TestHashFilename(t *testing.T) {
	_, err := hashFilename("")
	if err == nil {
		t.Errorf("Should receive error when trying to hash nil filename")
	}
	n, err := hashFilename("filename.txt")
	if err != nil {
		t.Errorf("Should not receive error from hashing valid filename")
	}
	if n != 3 {
		t.Errorf("Hashed filename should be equal to 3 when hashSpace = 64")
	}
}
