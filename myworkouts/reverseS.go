package main

import "fmt"

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	original := "Hello, Golang!"
	reversed := reverseString(original)

	fmt.Printf("Original: %s\nReversed: %s\n", original, reversed)
}
