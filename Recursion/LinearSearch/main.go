package main

import "fmt"

func main() {

	array := []int{1, 3, 7, 9, 4, 6, 2}

	target := 6

	result := LinearSearch(array, target)

	if result != -1 {
		fmt.Printf("Target %d is present at position %d\n", target, result)
		return
	}
	{
		fmt.Printf("target %d is not present\n", target)
		return
	}

}

func LinearSearch(array []int, target int) int {
	for i := range array {
		if array[i] == target {
			return i
		}
	}
	return -1
}
