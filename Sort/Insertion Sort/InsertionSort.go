package main

import "fmt"

func main() {
	arr := []int{5, 9, 2, 7, 1, 3, 8}
	n := len(arr)
	fmt.Println("Array before soring", arr)
	result := InsertionSort(arr, n)
	fmt.Println("Array before soring", result)
}

func InsertionSort(arr []int, n int) []int {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] >= key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
