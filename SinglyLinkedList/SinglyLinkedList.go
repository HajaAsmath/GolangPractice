package main

import "fmt"

var tail *SSL

//SSL struct
type SSL struct {
	next  *SSL
	value int
}

type SingleLinkedList struct {
	head *SSL
	tail *SSL
}

func CreateSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

//Insert Node
func (list *SingleLinkedList) Insert(value int) {
	newNode := CreateNode()
	if list.head == nil && list.tail == nil {
		newNode.value = value
		list.head = newNode
	} else if list.head == list.tail {
		newNode.value = value
		list.head.next = newNode
	} else if list.tail != nil {
		newNode.value = value
		list.tail.next = newNode
	}
	list.tail = newNode
}

//CreateNew Node
func CreateNode() *SSL {
	return new(SSL)
}

//Traver SSL
func (list *SingleLinkedList) String() string {
	s := ""
	for n := list.head; n.next != nil; n = n.next {
		s += fmt.Sprintf("%d ", n.value)
	}
	s += fmt.Sprintf("%d", list.tail.value)
	return s
}

func main() {
	list := CreateSingleLinkedList()
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)
	fmt.Println(list)
}
