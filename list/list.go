package list

import "fmt"

/*
LinkedList - A representation of a linear Doubly Linked List
*/
type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]

	Length int
}

/*
Node - A representation of a single node in the linked list
*/
type Node[T comparable] struct {
	Next *Node[T]
	Prev *Node[T]

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
init - Initialize the linked list by creating a new node and setting it as both the head and tail. Calls to Append, Prepend, and
Insert (with an index of 0) do this for you automatically so it would be arbitrary to call this function after declaring your list,
hence why it is not exported
*/
func (list *LinkedList[T]) init(data T) {
	head := &Node[T]{
		Next: nil,
		Prev: nil,
		Data: data,
	}

	list.Head = head
	list.Tail = head

	list.Length = 1
}

/*
Append - Add a new node to the end of the linked list. If the head node is nil then it will simply add the head node to the
linked list with the data passed in the parameter
*/
func (list *LinkedList[T]) Append(data T) {

	// If the list head is nil then create a new node and set it as the head and tail
	if list.Head == nil {
		list.init(data)
		return
	}

	newNode := &Node[T]{
		Next: nil,
		Prev: list.Tail,
		Data: data,
	}

	// this was failing previously as I wasn't setting the tail of our linked list to the new node.
	// if we are appending here, then our new node will always be the new tail
	list.Tail.Next = newNode
	list.Tail = list.Tail.Next

	list.Length += 1
}

/*
Prepend - Prepend the list with a new node containing the data requested with the 'data' parameter. Effectively
replaces the head of the linked list
*/
func (list *LinkedList[T]) Prepend(data T) {
	if list.Head == nil {
		list.init(data)
		return
	}

	newNode := &Node[T]{
		Prev: nil,
		Data: data,
	}

	newNode.Next = list.Head
	list.Head = newNode

	list.Length += 1
}

/*
GetData - Fetch the data stored in the node at the requested index
*/
func (list *LinkedList[T]) GetData(index int) interface{} {
	nodeCopy := list.Head

	i := 0
	for nodeCopy != nil {
		if i == index {
			return nodeCopy.Data
		}

		nodeCopy = nodeCopy.Next
		i += 1
	}

	return nil
}

/*
Get - Fetch a pointer to the node at the requested index. Returns nil if the node could not be found
*/
func (list *LinkedList[T]) Get(index int) *Node[T] {
	nodeCopy := list.Head

	i := 0
	for nodeCopy != nil {
		if i == index {
			return nodeCopy
		}

		nodeCopy = nodeCopy.Next
		i += 1
	}

	return nil
}

/*
All - Iterate through each node of the linked list, store them in order of their index and return them in a slice
*/
func (list *LinkedList[T]) All() []T {
	var ret []T

	nodeCopy := list.Head

	for nodeCopy != nil {
		ret = append(ret, nodeCopy.Data)

		nodeCopy = nodeCopy.Next
	}

	return ret
}

/*
Insert - Insert a new node at the requested index with the specified data.

! This needs to be updated to support this operation at 0 index (track head here along with tail)
*/
func (list *LinkedList[T]) Insert(data T, index int) {
	nodeCopy := list.Head

	i := 0
	for nodeCopy != nil {
		if i == index {
			newNode := &Node[T]{
				Data: data,
			}
			newNode.Prev = nodeCopy.Prev
			nodeCopy.Prev.Next = newNode
			nodeCopy.Prev = newNode

			newNode.Next = nodeCopy
			list.Length += 1
			break
		}

		nodeCopy = nodeCopy.Next
		i += 1
	}
}

/*
Remove - Remove a node from the linked list. If the node could not be found at the respective index then
the linked list will not be modified
*/
func (list *LinkedList[T]) Remove(index int) {
	nodeCopy := list.Head

	i := 0
	for nodeCopy != nil {
		if i == index {
			nodeCopy.Next.Prev = nodeCopy.Prev
			nodeCopy.Prev.Next = nodeCopy.Next

			list.Length -= 1
			break
		}

		nodeCopy = nodeCopy.Next
		i += 1
	}
}

/*
Search - Iterate through each element of the list starting at the head until the data stored
in the node matches the data passed in the 'data' parameter. Returns both the index and a pointer
to the node
*/
func (list *LinkedList[T]) Search(data T) (int, *Node[T]) {
	nodeCopy := list.Head

	i := 0
	for nodeCopy != nil {
		if nodeCopy.Data == data {
			return i, nodeCopy
		}

		nodeCopy = nodeCopy.Next
		i += 1
	}

	return -1, nil
}
