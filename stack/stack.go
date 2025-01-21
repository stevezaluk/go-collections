package stack

type Stack[T comparable] struct {
	Data   []T
	Length int
}
