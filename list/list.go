package list

/*
Node - A representation of a linear Doubly Linked List
*/
type Node[T interface{}] struct {
	Index int
	Next  *Node[T]
	Prev  *Node[T]

	Data T
}
