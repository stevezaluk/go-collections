package queue

import "github.com/stevezaluk/go-collections/list"

/*
IQueue - An interface describing all the functions that a queue implements. Implementation is
implicit in go, so no need to explicitly declare this
*/
type IQueue[T comparable] interface {
	Enqueue(T)
	Dequeue() T
	Front() T
	IsEmpty() bool
	Size() int
}

/*
Queue - Represents a basic first-in-first-out (FIFO) queue
*/
type Queue[T comparable] struct {
	data *list.LinkedList[T]
}

/*
NewQueue - A constructor for the queue. Creates a pointer to a new linked list and then
returns a pointer to the queue
*/
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{data: list.NewLinkedList[T]()}
}

/*
Enqueue - Place data at the end of the queue
*/
func (queue *Queue[T]) Enqueue(data T) {
	queue.data.Append(data)
}

/*
Dequeue - Remove an object from the front of the queue and return it
*/
func (queue *Queue[T]) Dequeue() T {
	ret := queue.data.GetHead().Data
	queue.data.RemoveHead()

	return ret
}

/*
Front - Return the data stored at the front of the queue without removing it
*/
func (queue *Queue[T]) Front() T {
	return queue.data.GetHead().Data
}

/*
Size - Get the amount of items stored in the queue
*/
func (queue *Queue[T]) Size() int {
	return queue.data.Length
}

/*
IsEmpty - Returns a boolean value if the list is empty. If the length of the underlying
linked list is 0, then this returns true. Otherwise, it returns false
*/
func (queue *Queue[T]) IsEmpty() bool {
	if queue.Size() == 0 {
		return true
	}

	return false
}
