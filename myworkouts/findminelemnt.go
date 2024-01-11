package main

import "fmt"

func findMin(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	min := arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	array := []int{30, 10, 50, 20, 40}
	min := findMin(array)

	fmt.Printf("Minimum element in the array: %d\n", min)
}
