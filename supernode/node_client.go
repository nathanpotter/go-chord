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
    node := &Node{ip: ip, port: port}
    var nodes *NodeList
    err := c.connection.Call("SuperNode.Join", node, &nodes)
    return nodes, err
}

func (c *NodeClient) PostJoin(ip string, port int) (bool, error) {
    var added bool
    node := &Node{ip: ip, port: port}
    err := c.connection.Call("SuperNode.PostJoin", node, &added)
    return added, err
}