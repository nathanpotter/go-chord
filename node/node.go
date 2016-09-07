package main

import (
	"flag"
	"log"
	"net"
	"strconv"

	"github.com/nathanpotter/go-chord/node/node"
	"google.golang.org/grpc"
	npb "github.com/nathanpotter/go-chord/protos/node"
)

var (
	port      int
	superIp   string
	superPort int
)

func main() {
	flag.IntVar(&port, "port", 10001, "Specify port to use")
	flag.StringVar(&superIp, "superIp", "localhost", "Specify supernode's Ip address")
	flag.IntVar(&superPort, "superPort", 10000, "Specify supernode's Port")
	flag.Parse()

	// convert to string and add : to front
	p := strconv.Itoa(port)
	p = ":" + p

	sp := strconv.Itoa(superPort)
	sp = ":" + sp

	lis, err := net.Listen("tcp", p)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ip, err := getIpAddr()
	if err != nil || ip == "" {
		log.Fatalf("Unable to get local ip address")
	}
	n := node.NewNode(ip, p)

	err = n.Join(superIp, sp)
	if err != nil {
		log.Fatalf("Unable to join chord network")
	}

	s := grpc.NewServer()
	npb.RegisterNodeServer(s, n)
	log.Println("Listening...")
	s.Serve(lis)

}

func getIpAddr() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("Unable to get network interfaces", err)
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}
