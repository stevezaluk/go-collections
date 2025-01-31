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
