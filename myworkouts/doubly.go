package main

import (
	"fmt"
)

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func main() {
	list := DoublyLinkedList{}
	list.AppendEnd(1)
	list.AppendEnd(2)
	list.AppendEnd(3)
	list.AppendEnd(4)
	list.AppendEnd(5)
	list.AppendEnd(6)

	list.InsertAtBeginning(0)
	list.InsertAtPosition(3, 2) // Inserting 3 at position 2

	list.DisplayList()
}

// AppendEnd appends a node at the end of the doubly linked list
func (list *DoublyLinkedList) AppendEnd(data int) {
	newNode := &Node{data: data, next: nil, prev: list.tail}
	if list.tail != nil {
		list.tail.next = newNode
	}
	list.tail = newNode

	if list.head == nil {
		list.head = newNode
	}
}

// InsertAtBeginning inserts a node at the beginning of the doubly linked list
func (list *DoublyLinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data, next: list.head}

	if list.head != nil {
		list.head.prev = newNode
	}

	list.head = newNode

	if list.tail == nil {
		list.tail = newNode
	}
}

// InsertAtPosition inserts a node at a specified position in the doubly linked list
func (list *DoublyLinkedList) InsertAtPosition(data int, position int) {
	if position <= 0 {
		list.InsertAtBeginning(data)
		return
	}

	newNode := &Node{data: data}
	temp := list.head
	for i := 1; i < position && temp != nil; i++ {
		temp = temp.next
	}

	if temp == nil {
		return
	}

	newNode.prev = temp
	newNode.next = temp.next
	if temp.next != nil {
		temp.next.prev = newNode
	}
	temp.next = newNode

	if newNode.next == nil {
		list.tail = newNode
	}
}

// DisplayList displays the doubly linked list
func (list *DoublyLinkedList) DisplayList() {
	temp := list.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
