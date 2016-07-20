package supernode

import (
    "net"
    "net/rpc"
    "time"
)

type (
    NodeClient struct {
        connection *rpc.Client
    }
)

func NewNodeClient(dsn string, timeout time.Duration) (*NodeClient, error) {
    connection, err := net.DialTimeout("tcp", dsn, timeout)
    if err != nil {
        return nil, err
    }
    return &NodeClient{connection: rpc.NewClient(connection)}, nil
}

func (c *NodeClient) Join(ip string, port int) (*NodeList, error) {
    var node *Node
    err := c.connection.Call("SuperNode.GetNode", nil, &node)
    return node, err
}