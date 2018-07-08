package main

import (
	"strings"
)

func WordCount(str string) map[string]int {
	strArr := strings.Fields(str)
	wordCount := make(map[string]int)
	for _, str := range strArr {
		wordCount[str]++
	}
	return wordCount
}
