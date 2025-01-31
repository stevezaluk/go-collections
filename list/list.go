package list

import "fmt"

/*
ILinkedList - An interface describing all the functions that a linked list
implements. Implementation is implicit in go, so no need to explicitly declare
this
*/
type ILinkedList[T any] interface {
	Append(T)
	Prepend(T)
	GetData(int) interface{}
	Get(int) *Node[T]
	All() []T
	Insert(T, int)
	Remove(int)
}

/*
LinkedList - A representation of a linear Doubly Linked List
*/
type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]

	Length int
}

/*
Node - A representation of a single node in the linked list
*/
type Node[T any] struct {
	Next *Node[T]
	Prev *Node[T]

	Data T
}

/*
NewLinkedList - Constructor for the linked list structure. This function
returns a Pointer to an empty LinkedList. This does not initialize the linked
list with a node, however any calls to Insert, Append, or Prepend will initialize
this for you
*/
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{Length: 0}
}

/*
Print - Iterate through each element of the linked list and print the data that it is storing

Time Complexity: o(N) where N is the length of the list
Reason: Each node must be iterated over in order to print the data stored in the node.
*/
func (list *LinkedList[T]) Print() {
	curr := list.Head

	for curr != nil {
		fmt.Println(curr.Data)

		curr = curr.Next
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
GetHead - Return a pointer to the node placed at the head of the linked list
*/
func (list *LinkedList[T]) GetHead() *Node[T] {
	return list.Head
}

/*
RemoveHead - Removes the head of the linked list, and sets the head to the node at the 1st index
*/
func (list *LinkedList[T]) RemoveHead() {
	newHead := list.Head.Next
	newHead.Prev = nil
	list.Head = newHead
	list.Length -= 1
}

/*
GetTail - Return a pointer to the node placed at the tail of the linked list
*/
func (list *LinkedList[T]) GetTail() *Node[T] {
	return list.Tail
}

/*
RemoveTail - Removes the tail of the linked list, and sets the tail to the node at the index before the tail
*/
func (list *LinkedList[T]) RemoveTail() {
	newTail := list.Tail.Prev
	newTail.Next = nil
	list.Tail = newTail
	list.Length -= 1
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
Get - Fetch a pointer to the node at the requested index. Returns nil if the node could not be found.
If 0 is passed to the index, then the head of the linked list is returned. If -1 is passed to the index,
then the tail of the linked list is returned.
*/
func (list *LinkedList[T]) Get(index int) *Node[T] {
	if index > (list.Length - 1) {
		return nil
	}

	if index == 0 {
		return list.GetHead()
	}

	if index == -1 {
		return list.GetTail()
	}

	curr := list.Head

	i := 0
	for curr != nil {
		if i == index {
			return curr
		}

		curr = curr.Next
		i += 1
	}

	return nil
}

/*
GetData - Fetch the data stored in the node at the requested index
*/
func (list *LinkedList[T]) GetData(index int) interface{} {
	node := list.Get(index)
	if node != nil {
		return node.Data
	}

	return nil
}

/*
All - Iterate through each node of the linked list, store them in order of their index and return them in a slice
*/
func (list *LinkedList[T]) All() []T {
	var ret []T

	curr := list.Head

	for curr != nil {
		ret = append(ret, curr.Data)

		curr = curr.Next
	}

	return ret
}

/*
Insert - Insert a new node at the requested index with the specified data. Calls to 0 are effectively
a Prepend operation and calls to -1 are an Append operation
*/
func (list *LinkedList[T]) Insert(data T, index int) {
	if index > (list.Length - 1) {
		return
	}

	if index == 0 {
		list.Prepend(data)
		return
	}

	if index == -1 {
		list.Append(data)
		return
	}

	curr := list.Head

	i := 0
	for curr != nil {
		if i == index {
			newNode := &Node[T]{
				Data: data,
			}
			newNode.Prev = curr.Prev
			curr.Prev.Next = newNode
			curr.Prev = newNode

			newNode.Next = curr
			list.Length += 1
			break
		}

		curr = curr.Next
		i += 1
	}
}

/*
Remove - Remove a node from the linked list. If the node could not be found at the respective index then
the linked list will not be modified
*/
func (list *LinkedList[T]) Remove(index int) {
	if index > (list.Length - 1) {
		return
	}

	if index == 0 {
		list.RemoveHead()
		return
	}

	if index == -1 {
		list.RemoveTail()
		return
	}

	curr := list.Head

	i := 0
	for curr != nil {
		if i == index {
			curr.Next.Prev = curr.Prev
			curr.Prev.Next = curr.Next

			list.Length -= 1
			break
		}

		curr = curr.Next
		i += 1
	}
}
