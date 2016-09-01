package node

import (
  "testing"
)

var (
  n *node
)

func TestNewNode(t *testing.T) {
  n = NewNode()
  if n == nil {
    t.Fatalf("NewNode should not return nil")
  }
  if n.this == nil {
    t.Fatalf("n.this should not be nil after returning from NewNode")
  }
}
