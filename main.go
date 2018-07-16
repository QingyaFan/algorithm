package main

import (
	"fmt"

	"./sort"
)

var arr = []int{1, 23, 2, 30, 4, 5, 66, 7, 8, 9}

func main() {
	fmt.Println(sort.QuickSort(arr))
	// test := FindMax(arr)
	// fmt.Println(test)
}
