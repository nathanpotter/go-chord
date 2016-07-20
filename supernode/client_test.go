package supernode

import (
    "log"
    "testing"
    "time"
)

var (
    c *Client
    err error
    dsn = "localhost:9876"
    s *SuperNode
    n *node.Node
)

func startSuperNode() {
    
}

func startNode() {
    
}

func init() {
    // start server
    c, err = NewClient(dsn, time.Millisecond*500)
    if err != nil {
        log.Fatal(err)
    }
}

func TestColdGetNode(t *testing.T) {
    node, err := c.GetNode()
    if err == nil {
        t.Error("Should have received error message, no nodes in the system")
    }
}

func TestWarmGetNode(t *testing.T) {
    // start node server
    node, err := c.GetNode()
    if node == nil {
        t.Errorf("Node shouldn't be nil", err)
    }
}

