package test

import (
	"github.com/stevezaluk/go-collections/list"
	"testing"
)

/*
TestListPrint - Test the Print() function of the linked list
*/
func TestListPrint(t *testing.T) {
	head := &list.Node[string]{Data: "headNode"}
	head.Next = &list.Node[string]{Data: "secondNode"}

	head.Print()
}
