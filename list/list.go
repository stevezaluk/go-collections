package list

import "fmt"

/*
Node - A representation of a linear Doubly Linked List
*/
type Node[T interface{}] struct {
	Index int
	Next  *Node[T]
	Prev  *Node[T]

	Data T
}

/*
Print - Iterate through each element of the linked list and print the data that it is storing
*/
func (list *Node[T]) Print() {
	nodeCopy := list

	for nodeCopy != nil {
		fmt.Println(nodeCopy.Data)

		nodeCopy = nodeCopy.Next
	}
}
