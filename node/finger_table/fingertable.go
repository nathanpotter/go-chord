// package fingertable represents the fingertable that allows the nodes in the
// system to keep track of eachother and allows for O(log(n)) lookup time

package fingertable

type fingertable struct {

}

func NewFingertable() *fingertable {
  return &fingertable{}
}
