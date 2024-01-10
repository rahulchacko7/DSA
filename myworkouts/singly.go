package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {
	list := LinkedList{}
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

// AppendEnd appends a node at the end of the linked list
func (list *LinkedList) AppendEnd(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

// InsertAtBeginning inserts a node at the beginning of the linked list
func (list *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data, next: nil}
	newNode.next = list.head
	list.head = newNode
}

// InsertAtPosition inserts a node at a specified position in the linked list
func (list *LinkedList) InsertAtPosition(data int, position int) {
	if position <= 0 {
		list.InsertAtBeginning(data)
		return
	}

	newNode := &Node{data: data}
	temp := list.head
	for i := 1; i < position-1 && temp != nil; i++ {
		temp = temp.next
	}

	if temp == nil {
		return
	}

	newNode.next = temp.next
	temp.next = newNode
}

// DisplayList displays the linked list
func (list *LinkedList) DisplayList() {
	temp := list.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
