package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	reversed := make([]int, len(arr))

	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		reversed[i] = arr[j]
	}

	fmt.Println("Original Array:", arr)
	fmt.Println("Reversed Array:", reversed)
}
