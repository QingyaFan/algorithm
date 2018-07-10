package main

func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else {
		return arr[0] + Sum(arr[1:])
	}
}

func FindMax(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	} else if arr[0] > FindMax(arr[1:]) {
		return arr[0]
	} else {
		return FindMax(arr[1:])
	}
}
