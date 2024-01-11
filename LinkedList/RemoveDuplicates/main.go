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
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func (list *LinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Printf("%d ", temp.data) // Added %d to print the integer data
		temp = temp.next
	}
	fmt.Println()
}

func (list *LinkedList) RemoveDuplicates() {
	if list.head == nil {
		return
	}
	temp := list.head
	for temp != nil {
		runner := temp
		for runner.next != nil {
			if runner.next.data == temp.data {
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}
		temp = temp.next
	}
}

func main() {
	mylist := LinkedList{} // Corrected struct instantiation
	mylist.Append(2)
	mylist.Append(3)
	mylist.Append(2)
	mylist.Append(9)
	mylist.Append(3)
	fmt.Println("List Before Removing Duplicates:")
	mylist.Print()

	mylist.RemoveDuplicates()

	fmt.Println("List After Removing Duplicates:")
	mylist.Print()
}
