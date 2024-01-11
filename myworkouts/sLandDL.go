package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node // For Doubly Linked List
}

type SinglyLinkedList struct {
	head *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func main() {
	// Construction of Singly Linked List
	singlyList := SinglyLinkedList{}
	singlyList.head = &Node{data: 1}
	singlyList.head.next = &Node{data: 2}
	singlyList.head.next.next = &Node{data: 3}

	// Construction of Doubly Linked List
	doublyList := DoublyLinkedList{}
	doublyList.head = &Node{data: 1}
	doublyList.tail = doublyList.head
	doublyList.tail.next = &Node{data: 2, prev: doublyList.head}
	doublyList.tail = doublyList.tail.next
	doublyList.tail.next = &Node{data: 3, prev: doublyList.tail}

	// Print the constructed lists
	fmt.Println("Singly Linked List:")
	printLinkedList(&singlyList)

	fmt.Println("\nDoubly Linked List:")
	printLinkedList(&doublyList)
}

func printLinkedList(list *SinglyLinkedList) {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
