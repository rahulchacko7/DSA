package main

import "fmt"

func main() {

	arr := []int{6, 2, 5, 1, 8, 9, 4}
	BubbleSort(arr)
	fmt.Println("Sorted Array is...", arr)
}

func BubbleSort(arr []int) {

	limit := len(arr)

	for j := 0; j < limit-1; j++ {
		for i := 0; i < limit-j-1; i++ {
			if arr[i] < arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
			}
		}
	}

}
