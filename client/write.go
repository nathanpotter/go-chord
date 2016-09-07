package main

import (
  "flag"
  "log"
  "strconv"
  "time"

  "golang.org/x/net/context"

  pb "github.com/nathanpotter/go-chord/protos/common"
	npb "github.com/nathanpotter/go-chord/protos/node"
	spb "github.com/nathanpotter/go-chord/protos/supernode"
  "github.com/nathanpotter/go-chord/super_node/supernode"
  "google.golang.org/grpc"
  "github.com/golang/protobuf/proto"
)

var (
  superIp string
  superPort int
)

func main() {
  flag.StringVar(&superIp, "superIp", "localhost", "Specify supernode's Ip address")
	flag.IntVar(&superPort, "superPort", 10000, "Specify supernode's Port")
  flag.Parse()

  // convert to string and add : to front
	p := strconv.Itoa(superPort)
	p = ":" + p

  conn, err := grpc.Dial((superIp + p), grpc.WithInsecure())
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()

  s := spb.NewSupernodeClient(conn)

  node, err := s.GetNode(context.Background(), &pb.Empty{})
  for err == supernode.BusyError {
    time.Sleep(1 * time.Second)
    node, err = s.GetNode(context.Background(), &pb.Empty{})
  }
  if err != nil {
    log.Fatal(err)
  }

  log.Println(node)

  nConn, err := grpc.Dial((node.Ip + node.Port), grpc.WithInsecure())
  if err != nil {
    log.Fatal(err)
  }
  defer nConn.Close()

  n := npb.NewNodeClient(nConn)
  if err != nil {
    log.Fatal(err)
  }

  file1 := &pb.File{Name: "file1", Contents: []byte("Here are the contents of file1")}
  file2 := &pb.File{Name: "file2", Contents: []byte("Here is file 2's contents")}
  file3 := &pb.File{Name: "file3", Contents: []byte("Here is file 3's contents")}
  file4 := &pb.File{Name: "file4", Contents: []byte("Here is file 4's contents")}
  file5 := &pb.File{Name: "file5", Contents: []byte("Here is file 5's contents")}
  file6 := &pb.File{Name: "file6", Contents: []byte("Here is file 6's contents")}


  // test we can read file 1
  testFile1 := &pb.File{Name: "file1"}

  _, err = n.Write(context.Background(), file1)
  if err != nil {
    log.Println("Error writing file1 to the system:", err)
  }

  _, err = n.Write(context.Background(), file2)
  if err != nil {
    log.Println("Error writing file2 to the system:", err)
  }

  _, err = n.Write(context.Background(), file3)
  if err != nil {
    log.Println("Error writing file2 to the system:", err)
  }

  _, err = n.Write(context.Background(), file4)
  if err != nil {
    log.Println("Error writing file2 to the system:", err)
  }

  _, err = n.Write(context.Background(), file5)
  if err != nil {
    log.Println("Error writing file2 to the system:", err)
  }

  _, err = n.Write(context.Background(), file6)
  if err != nil {
    log.Println("Error writing file2 to the system:", err)
  }

  result, err := n.Read(context.Background(), testFile1)
  if err != nil {
    log.Println("Error reading file1 from the system:", err)
  }

  if !proto.Equal(result, file1) {
    log.Println("Error, result and file1 don't match: ", result, file1)
  }
  log.Printf("\nfile1: %v\nresult: %v\n", file1, result)

  result, err = n.Read(context.Background(), &pb.File{Name: "file6"})
  if err != nil {
    log.Println("Error reading file1 from the system:", err)
  }

  if !proto.Equal(result, file6) {
    log.Println("Error, result and file1 don't match: ", result, file6)
  }
  log.Printf("\nfile1: %v\nresult: %v\n", file6, result)
}
