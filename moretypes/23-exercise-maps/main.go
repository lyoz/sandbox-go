package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	freq := map[string]int{}
	for _, word := range strings.Fields(s) {
		freq[word]++
	}
	return freq
}

func main() {
	wc.Test(WordCount)
}
