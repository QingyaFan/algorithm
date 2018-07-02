package main

import (
	"fmt"
	"math"
)

var arr [1024] int

func initArr() {
	for i := 0; i < len(arr); i++ {
		arr[i] = i*2
	}
}

func binarySearch(search int) int {
	low := 0
	high := len(arr) - 1
	for low < high {
		mid := int(math.Ceil(float64(low + high)/2))
		val := arr[mid]
		if val == search {
			return mid
		} else if val < search {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	initArr()
	result := binarySearch(100)
	fmt.Println(result)
}