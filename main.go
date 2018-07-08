package main

import (
	"fmt"
)

func main() {
	const wordStr = "golang nodejs nodejs docker docker kubernetes linux linux linux"
	res := WordCount(wordStr)
	fmt.Println(res)

	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
