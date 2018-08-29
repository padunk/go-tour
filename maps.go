package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	l := len(words)
	result := make(map[string]int, l)

	for i := 0; i < l; i++ {
		if result[words[i]] == 0 {
			result[words[i]] = 1
		} else {
			result[words[i]] += 1
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
