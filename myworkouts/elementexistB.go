package main

import "fmt"

func doesElementExistBinary(arr []int, target int) bool {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	array := []int{5, 10, 15, 20, 25}
	target := 15
	exists := doesElementExistBinary(array, target)

	if exists {
		fmt.Printf("Element %d exists in the array\n", target)
	} else {
		fmt.Printf("Element %d does not exist in the array\n", target)
	}
}
