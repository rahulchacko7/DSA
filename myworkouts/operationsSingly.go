package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) addNode(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) deleteNode(value int) {
	if list.head == nil {
		return
	}

	if list.head.data == value {
		list.head = list.head.next
		return
	}

	current := list.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next == nil {
		return
	}

	current.next = current.next.next
}

func (list *LinkedList) insertAfter(prevData, newData int) {
	current := list.head
	for current != nil && current.data != prevData {
		current = current.next
	}

	if current == nil {
		return
	}

	newNode := &Node{data: newData, next: current.next}
	current.next = newNode
}

func (list *LinkedList) printList() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (list *LinkedList) removeDuplicates() {
	current := list.head

	for current != nil && current.next != nil {
		if current.data == current.next.data {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}
}

func main() {
	linkedList := LinkedList{}

	linkedList.addNode(1)
	linkedList.addNode(2)
	linkedList.addNode(2)
	linkedList.addNode(3)
	linkedList.addNode(4)
}
