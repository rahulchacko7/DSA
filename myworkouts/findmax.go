package main

import "fmt"

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	max := arr[len(arr)-1]
	return max
}

func main() {
	array := []int{30, 10, 50, 20, 40}
	max := findMax(array)

	fmt.Printf("Maximum element in the array: %d\n", max)
}
