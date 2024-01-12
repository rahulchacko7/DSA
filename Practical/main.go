package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	last := list.head
	for last.next != nil {
		last = last.next
	}
	last.next = newNode
}

func (list *LinkedList) Print() {
	current := list.head

	for current != nil {
		fmt.Println(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func deleteMiddle(list *LinkedList) {
	if list.head == nil || list.head.next == nil {
		return
	}

	slow := list.head
	fast := list.head

	var prev *Node

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		prev = slow
		slow = slow.next
	}

	prev.next = slow.next
}

func main() {
	linkedList := LinkedList{}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)

	fmt.Println("Orginal List is :")
	linkedList.Print()

	deleteMiddle(&linkedList)
	fmt.Println("List after delete is :")
	linkedList.Print()
}
