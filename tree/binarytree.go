package tree

type BinaryRoot struct {
	Left, Right *BinaryRoot
	Value       interface{}
}
