package hash

/*
KeyPair - A structure representing a key pair in a Hash Map
*/
type KeyPair[T any] struct {
	Key   string
	Value T
}

/*
NewKeyPair - A constructor for the KeyPair structure. Returns a pointer
to the KeyPair initialized with the values that are provided as arguments
*/
func NewKeyPair[T any](key string, value T) *KeyPair[T] {
	return &KeyPair[T]{Key: key, Value: value}
}

/*
HashMap - Represents the HashMap as a whole and encapsulates it's all of its
logic.
*/
type HashMap[T any] struct {
	data     []*KeyPair[T]
	length   uint64
	capacity uint64
}

/*
NewHashMap - A constructor for the HashMap structure. Returns a pointer
to the HashMap structure with
*/
func NewHashMap[T string, K any]() *HashMap[T] {
	return &HashMap[T]{data: make([]*KeyPair[T], 0), length: 0, capacity: 0}
}
