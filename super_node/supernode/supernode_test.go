package supernode

import (
  "testing"
)

var (
  s *supernode
  err error
)

func TestNewSupernode(t *testing.T) {
  s = NewSupernode()
  if s == nil {
    t.Fatalf("Supernode unable to be created.")
  }
  if s.nodes == nil {
    t.Errorf("Nodes slice not initialized in supernode")
  }
}
