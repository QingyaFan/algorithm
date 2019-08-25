package main

import (
	"algorithm/sort"
	"fmt"
)

var arr = []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}

func main() {
	fmt.Println(sort.QuickSort2(arr))
}
