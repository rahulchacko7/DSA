package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	array := []int{10, 20, 30, 40, 50, 60}
	target := 40
	result := linearSearch(array, target)

	if result != -1 {
		fmt.Printf("Element %d found at index %d\n", target, result)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
