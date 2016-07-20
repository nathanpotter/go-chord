package supernode

import (
    "log"
    "testing"
    "time"
)

var (
    c *NodeClient
    err error
    dsn = "localhost:9876"
    s *SuperNode
)

func startSuperNode() {
    
}



func init() {
    // start server
    c, err = NewNodeClient(dsn, time.Millisecond*500)
    if err != nil {
        log.Fatal(err)
    }
}

func TestJoin(t *testing.T) {

}

func TestPostJoin(t *testing.T) {

}

