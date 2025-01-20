package list

import "fmt"

/*
Node - A representation of a linear Doubly Linked List
*/
type Node[T any] struct {
	Index int
	Next  *Node[T]
	Prev  *Node[T]

	Data T
}

/*
Print - Iterate through each element of the linked list and print the data that it is storing

Time Complexity: o(N) where N is the length of the list
Reason: Each node must be iterated over in order to print the data stored in the node.
*/
func (list *Node[T]) Print() {
	nodeCopy := list

	for nodeCopy != nil {
		fmt.Println(nodeCopy.Data)

		nodeCopy = nodeCopy.Next
	}
}
