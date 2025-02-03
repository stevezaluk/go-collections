package hash

import (
	"github.com/stevezaluk/fnv-hash/fnv"
	"github.com/stevezaluk/fnv-hash/prime"
)

/*
KeyPair - A structure representing a key pair in a Hash Map
*/
type KeyPair[T interface{}] struct {
	Key   string
	Value T
}

/*
NewKeyPair - A constructor for the KeyPair structure. Returns a pointer
to the KeyPair initialized with the values that are provided as arguments
*/
func NewKeyPair[T interface{}](key string, value T) *KeyPair[T] {
	return &KeyPair[T]{Key: key, Value: value}
}

/*
HashMap - Represents the HashMap as a whole and encapsulates it's all of its
logic.
*/
type HashMap[T interface{}] struct {
	data     []*KeyPair[T]
	length   int
	capacity int
}

/*
NewHashMap - A constructor for the HashMap structure. Returns a pointer
to the HashMap structure with
*/
func NewHashMap[T interface{}]() *HashMap[T] {
	return &HashMap[T]{data: make([]*KeyPair[T], 0), length: 0, capacity: 0}
}

/*
index - Return the index of a given key as an integer
*/
func (hashMap *HashMap[T]) index(key string) int {
	return int(fnv.NewFNV1AHash(prime.BitSize64, []byte(key)) % 6)
}

/*
reallocate - Extend the underlying array by the size passed in the parameter.
The result will be n + size, where n is the current length of the array
*/
func (hashMap *HashMap[T]) reallocate(size int) {
	var newItem T

	for i := 0; i < size; i++ {
		hashMap.data = append(hashMap.data, NewKeyPair("empty", newItem))
	}
}
