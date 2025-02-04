package hash

import (
	"errors"
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

WARNING: THIS CANNOT CURRENTLY HANDLE HASH COLLISIONS
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
Capacity - Returns the size of the underlying array, including both real elements and placeholder elements
*/
func (hashMap *HashMap[T]) Capacity() int {
	return hashMap.capacity
}

/*
Length - Returns the number of real elements stored in the array
*/
func (hashMap *HashMap[T]) Length() int {
	return hashMap.length
}

/*
Index - Return the index of a given key as an integer
*/
func (hashMap *HashMap[T]) index(key string) int {
	return int(fnv.NewFNV1AHash(prime.BitSize64, []byte(key)) % 11)
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

	hashMap.capacity += size
}

/*
Set - Insert a new value into the hash map and
*/
func (hashMap *HashMap[T]) Set(key string, value T) {
	index := hashMap.index(key)

	// if our index is greater than our current greatest possible index,
	// then re-allocate space for it in the underlying array
	if index > (hashMap.capacity - 1) {
		hashMap.reallocate(index - (hashMap.capacity - 1))
	}

	// array has the space for our new element
	keyPair := hashMap.data[index]

	if keyPair.Key != "empty" && keyPair.Key != key {
		// hash collision has happened. Deal with it accordingly
	}

	hashMap.data[index] = NewKeyPair(key, value)
	hashMap.length += 1
}

/*
Get - Fetch an item stored in the hash map
*/
func (hashMap *HashMap[T]) Get(key string) (T, error) {
	var ret T

	index := hashMap.index(key)

	if index > (hashMap.capacity - 1) {
		return ret, errors.New("hashmap: Failed to find value with the specified key (index is greater then the greatest possible index)")
	}

	keyPair := hashMap.data[index]
	if keyPair.Key == "empty" {
		return ret, errors.New("hashmap: Failed to find value with the specified key (key was not found)")
	}

	return keyPair.Value, nil
}
