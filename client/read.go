package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/context"

	pb "github.com/nathanpotter/go-chord/protos/common"
	npb "github.com/nathanpotter/go-chord/protos/node"
	spb "github.com/nathanpotter/go-chord/protos/supernode"
	"github.com/nathanpotter/go-chord/super_node/supernode"
	"google.golang.org/grpc"
)

var (
	superIp   string
	superPort int
	baseDir   string
)

func main() {
	flag.StringVar(&superIp, "superIp", "localhost", "Specify supernode's Ip address")
	flag.IntVar(&superPort, "superPort", 10000, "Specify supernode's Port")
	flag.StringVar(&baseDir, "baseDir", "resources", "Base directory for files")
	flag.Parse()

	if len(os.Args) != 2 {
		fmt.Println("Need filename to read from system as argument")
		os.Exit(-1)
	}

	filename := os.Args[1]

	contents, err := read(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("Filename: %s\n", filename)
	fmt.Printf("Contents: %v\n", string(contents))

}

func read(filename string) ([]byte, error) {

	node, err := getNode()
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial((node.Ip + node.Port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	n := npb.NewNodeClient(conn)
	if err != nil {
		return nil, err
	}

	file, err := n.Read(context.Background(), &pb.File{Name: filename})
	if err != nil {
		return nil, err
	}
	return file.Contents, nil
}

func getNode() (*pb.Node, error) {
	p := strconv.Itoa(superPort)
	p = ":" + p

	conn, err := grpc.Dial((superIp + p), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	s := spb.NewSupernodeClient(conn)

	node, err := s.GetNode(context.Background(), &pb.Empty{})
	for err == supernode.BusyError {
		time.Sleep(1 * time.Second)
		node, err = s.GetNode(context.Background(), &pb.Empty{})
	}
	return node, err
}
