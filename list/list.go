package list

import "fmt"

/*
LinkedList - A representation of a linear Doubly Linked List
*/
type LinkedList[T interface{}] struct {
	Head *Node[T]
	Tail *Node[T]

	Length int
}

/*
Node - A representation of a single node in the linked list
*/
type Node[T interface{}] struct {
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
func (list *LinkedList[T]) Print() {
	nodeCopy := list.Head

	for nodeCopy != nil {
		fmt.Println(nodeCopy.Data)

		nodeCopy = nodeCopy.Next
	}
}

/*
Append - Add a new node to the end of the linked list. If the head node is nil then it will simply add the head node to the
linked list with the data passed in the parameter
*/
func (list *LinkedList[T]) Append(data T) {

	// If the list head is nil then create a new node and set it as the head and tail
	if list.Head == nil {
		headNode := &Node[T]{
			Index: 0, // head node will always have an index of 0
			Next:  nil,
			Prev:  nil,
			Data:  data,
		}
		list.Head = headNode
		list.Tail = headNode
		list.Length = 1

		return
	}

	newNode := &Node[T]{
		Index: list.Tail.Index + 1,
		Next:  nil,
		Prev:  list.Tail,
		Data:  data,
	}

	// this was failing previously as I wasn't setting the tail of our linked list to the new node.
	// if we are appending here, then our new node will always be the new tail
	list.Tail.Next = newNode
	list.Tail = list.Tail.Next

	list.Length += 1
}

/*
GetData - Fetch the data stored in the node at the requested index
*/
func (list *LinkedList[T]) GetData(index int) interface{} {
	nodeCopy := list.Head

	for nodeCopy.Next != nil {
		if nodeCopy.Index == index {
			return nodeCopy.Data
		}
		nodeCopy = nodeCopy.Next
	}

	return nil
}
