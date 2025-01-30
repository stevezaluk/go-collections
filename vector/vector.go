package vector

/*
IVector - An interface describing all the functions that the Vector implements.
It accepts an interface as its generic type, allowing you to store any
data type needed
*/
type IVector[T interface{}] interface {
	Append(T)
	Get(int) interface{}
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
