package main

import "fmt"

func countOccurrences(input string, char rune) int {
	count := 0

	for _, c := range input {
		if c == char {
			count++
		}
	}

	return count
}

func main() {
	sentence := "Golang is awesome!"
	charToCount := 'o'
	occurrences := countOccurrences(sentence, charToCount)

	fmt.Printf("Occurrences of '%c' in '%s': %d\n", charToCount, sentence, occurrences)
}
