package main

import "fmt"

func main() {

	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{4, 5, 6, 7, 8}

	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2)
	fmt.Println("Array 3:", arr3)

	fmt.Println("Element at index 2 in Array 2:", arr2[2])

	arr1[0] = 10
	fmt.Println("Updated Array 1:", arr1)

	fmt.Println("Length of Array 3:", len(arr3))

	fmt.Print("Array 3 elements:")
	for _, num := range arr3 {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
