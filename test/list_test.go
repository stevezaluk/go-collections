package test

import (
	"github.com/stevezaluk/go-collections/list"
	"log"
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

func TestListAppend(t *testing.T) {
	_list := &list.LinkedList[string]{}
	_list.Append("headNode")
	_list.Append("secondNode")
	_list.Append("thirdNode")
	_list.Append("fourthNode")
	_list.Append("fifthNode")

	if _list.Length != 5 {
		log.Fatal("Length of linked list is not expected", "currentLength", _list.Length, "expectedLength", 5)
	}
}
