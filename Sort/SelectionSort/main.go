package main

import "fmt"

func main() {
	arr := []int{2, 64, 44, 9, 23, 7, 5}
	n := len(arr)
	fmt.Println("Array before sort :", arr)
	result := SelectionSort(arr, n)
	fmt.Println("Array after sort", result)
}

func SelectionSort(arr []int, n int) []int {
	for i := 0; i < n-1; i++ {

		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
