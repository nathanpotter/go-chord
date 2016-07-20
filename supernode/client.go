package supernode

import (
    "net"
    "net/rpc"
    "time"
)

type (
    Client struct {
        connection *rpc.Client
    }
)

func NewClient(dsn string, timeout time.Duration) (*Client, error) {
    connection, err := net.DialTimeout("tcp", dsn, timeout)
    if err != nil {
        return nil, err
    }
    return &Client{connection: rpc.NewClient(connection)}, nil
}

func (c *Client) GetNode() (*node.Client, error) {
    var node *Node
    err := c.connection.Call("SuperNode.GetNode", nil, &node)
    return node, err
}