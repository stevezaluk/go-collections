package stack

/*
Stack - Represents a LIFO stack data structure. We use an array internally to store the data as this provides us
with a O(1) time when randomly accessing elements
*/
type Stack[T comparable] struct {
	Data   []T
	Length int
}

/*
Push - Push a new element to the top of the stack
*/
func (stack *Stack[T]) Push(data T) {
	stack.Data = append(stack.Data, data)

	stack.Length += 1
}
