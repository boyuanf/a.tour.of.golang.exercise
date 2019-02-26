package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordmap := make(map[string]int)
	for _, word := range words {
		wordmap[word] += 1
	}
	return wordmap
}

func main() {
	wc.Test(WordCount)
}