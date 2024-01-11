package main

import "fmt"

func power(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	return base * power(base, exponent-1)
}

func main() {
	result := power(2, 3)
	fmt.Println("2^3:", result)
}
