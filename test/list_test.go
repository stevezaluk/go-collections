package test

import (
	"github.com/stevezaluk/go-collections/list"
	"testing"
)

/*
TestListPrint - Test the Print() function of the linked list
*/
func TestListPrint(t *testing.T) {
	_list := &list.LinkedList[string]{
		Head: &list.Node[string]{
			Data: "headNode",
		},
	}

	_list.Head.Next = &list.Node[string]{
		Data: "secondNode",
	}

	_list.Print()
}
