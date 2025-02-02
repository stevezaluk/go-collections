package hash

/*
KeyPair - A structure representing a key pair in a Hash Map
*/
type KeyPair[T comparable, K any] struct {
	Key   T
	Value K
}

/*
NewKeyPair - A constructor for the KeyPair structure. Returns a pointer
to the KeyPair initialized with the values that are provided as arguments
*/
func NewKeyPair[T comparable, K any](key T, value K) *KeyPair[T, K] {
	return &KeyPair[T, K]{Key: key, Value: value}
}

/*
HashMap - Represents the HashMap as a whole and encapsulates it's all of its
logic.
*/
type HashMap[T comparable, K any] struct {
	data   []*KeyPair[T, K]
	length uint64
}

/*
NewHashMap - A constructor for the HashMap structure. Returns a pointer
to the HashMap structure with
*/
func NewHashMap[T comparable, K any]() *HashMap[T, K] {
	return &HashMap[T, K]{data: make([]*KeyPair[T, K], 0), length: 0}
}
