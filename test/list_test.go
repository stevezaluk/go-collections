package test

import (
	"fmt"
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

func TestListGetData(t *testing.T) {
	_list := &list.LinkedList[string]{}
	_list.Append("headNode")
	_list.Append("secondNode")
	_list.Append("thirdNode")
	_list.Append("fourthNode")
	_list.Append("fifthNode")

	data := _list.GetData(3)
	if data != "fourthNode" {
		log.Fatal("Data does not equal the expected value", "currentData", data, "expectedData", "fourthNode")
	}
}

func TestListGet(t *testing.T) {
	_list := &list.LinkedList[string]{}
	_list.Append("headNode")
	_list.Append("secondNode")
	_list.Append("thirdNode")
	_list.Append("fourthNode")
	_list.Append("fifthNode")

	node := _list.Get(3)
	if node.Data != "fourthNode" {
		log.Fatal("Data does not equal the expected value", "currentData", node.Data, "expectedData", "fourthNode")
	}
}
func TestListAll(t *testing.T) {
	_list := &list.LinkedList[string]{}
	_list.Append("headNode")
	_list.Append("secondNode")
	_list.Append("thirdNode")
	_list.Append("fourthNode")
	_list.Append("fifthNode")

	_list.Print()
	all := _list.All()

	fmt.Println(all)
	if len(all) != 5 {
		log.Fatal("Length of linked list content does not equal the expected value", "currentLength", len(all), "expectedLength", 5)
	}

}

func TestListIndex(t *testing.T) {
	_list := &list.LinkedList[string]{}
	_list.Append("headNode")
	_list.Append("secondNode")
	_list.Append("thirdNode")
	_list.Append("fourthNode")
	_list.Append("fifthNode")

	_list.Insert("insertedNode", 3)

	_list.Print()
}
