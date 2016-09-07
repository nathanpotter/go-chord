package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	pb "github.com/nathanpotter/go-chord/protos/supernode"
	"github.com/nathanpotter/go-chord/super_node/supernode"
	"google.golang.org/grpc"
)

var (
	port int
)

func main() {
	flag.IntVar(&port, "port", 10000, "Specify port to use")
	flag.Parse()

	// convert to string and add : to front
	p := strconv.Itoa(port)
	p = ":" + p

	lis, err := net.Listen("tcp", p)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSupernodeServer(s, supernode.NewSupernode())
	log.Println("Listening...")
	s.Serve(lis)

}
