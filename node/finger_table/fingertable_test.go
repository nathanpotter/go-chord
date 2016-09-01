package fingertable

import (
  "testing"
)

var (
  tbl *fingertable
)

func TestNewFingertable(t *testing.T) {
  tbl = NewFingertable()
  if tbl == nil {
    t.Fatalf("NewFingertable should not return nil")
  }
}
