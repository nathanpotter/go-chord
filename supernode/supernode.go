package supernode

import (
    "errors"
    "sync"
)

type (
    SuperNode struct {
        nodes NodeList
        busy bool
        mu *sync.Mutex        
    }
    
    Node struct {
        ip string
        port int
    }
    
    NodeList []Node
)

var (
    BusyError = errors.New("Busy waiting for Node to finish joining")
    ZeroNodesError = errors.New("There aren't any nodes in the system")
)

func NewSuperNode() *SuperNode {
    return &SuperNode{
        nodes: &NodeList{},
        busy: false,
        mu: &sync.Mutex{},
    }
}

func (s *SuperNode) GetNode(_ *struct{}, node *Node) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if len(s.nodes) == 0 {
        return ZeroNodesError
    }
    
    *node = NodeList[0]
    return nil
}

func (s *SuperNode) Join(node *Node, nodes *NodeList) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if s.busy {
        return BusyError
    }
    s.busy = true
    nodes = s.nodes
    Append(s.nodes, node)
    return nil
}

func (s *SuperNode) PostJoin(node *Node, tru bool) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.busy = false
    return nil
}