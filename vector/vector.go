package vector

/*
IVector - An interface describing all the functions that the Vector implements.
It accepts an interface as its generic type, allowing you to store any
data type needed
*/
type IVector[T any] interface {
	Append(T)
	Get(int) any
	Insert(T, int)
	Remove(int)
	Length() int
}

/*
Vector - A representation of a vector (or a dynamic array). Its underlying array
is allocated with 0 elements
*/
type Vector[T interface{}] struct {
	data   [0]T
	length int
}

/*
NewVector - Serves as a constructor for the vector. Calling this function
will return an empty vector with its length assigned to 0
*/
func NewVector[T any]() *Vector[T] {
	return &Vector[T]{length: 0}
}
