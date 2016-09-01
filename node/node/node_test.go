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
}
