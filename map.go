// Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

// You might find strings.Fields helpful.

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordsField := make(map[string]int)
	words := strings.Fields(s)
	for _, k := range words {
		wordsField[k]++
	}
	return wordsField
}

func main() {
	wc.Test(WordCount)
}
