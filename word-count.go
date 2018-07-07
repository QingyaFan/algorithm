package main

import (
	"fmt"
	"strings"
)

func wordCount(str string) map[string]int {
	strArr := strings.Fields(str)
	wordCount := make(map[string]int)
	for _, str := range strArr {
		wordCount[str]++
	}
	return wordCount
}

func main() {
	fmt.Println(wordCount("hello mr mr dj dj dj  "))
}
