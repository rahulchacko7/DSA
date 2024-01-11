package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func arrayToLinkedList(arr []int) *LinkedList {
	list := LinkedList{}

	for _, num := range arr {
		list.addNode(num)
	}

	return &list
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

func main() {
	array := []int{1, 2, 3, 4, 5}
	linkedList := arrayToLinkedList(array)

	// Print the converted linked list
	current := linkedList.head
	fmt.Println("Linked List from Array:")
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
