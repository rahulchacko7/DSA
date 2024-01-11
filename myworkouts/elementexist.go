package main

import "fmt"

func doesElementExist(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	array := []int{5, 10, 15, 20, 25}
	target := 15
	exists := doesElementExist(array, target)

	if exists {
		fmt.Printf("Element %d exists in the array\n", target)
	} else {
		fmt.Printf("Element %d does not exist in the array\n", target)
	}
}
