package main

import "fmt"

func main() {
	arr := []int{7, 2, 9, 1, 5}
	max := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	fmt.Println("Maximum element in array:", max)
}
