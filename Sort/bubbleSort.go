package main

import "fmt"

func main() {
	Array := []int{5, 2, 8, 4, 6}
	BubbleSort(Array)
	fmt.Println("Sorted array is ", Array)
}

func BubbleSort(Array []int) {
	length := len(Array)

	for j := 0; j < length-1; j++ {
		for i := 0; i < length-j-1; i++ {
			if Array[i] > Array[i+1] {
				temp := Array[i]
				Array[i] = Array[i+1]
				Array[i+1] = temp
			}
		}
	}
}
