package tree

/*
BinaryTree - Represents an ordered binary tree
*/
type BinaryTree struct {
	Head *BinaryNode
}

/*
BinaryNode - Represents one node in the binary tree
*/
type BinaryNode struct {
	Data  int
	Left  *BinaryNode
	Right *BinaryNode
}

/*
New - Initialize a new binary tree and set the head node to the data passed in the parameter
*/
func New(data int) *BinaryTree {
	return &BinaryTree{Head: &BinaryNode{Data: data}}
}
