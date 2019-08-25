package main

import (
	"algorithm/sort"
	"testing"
)

// BenchmarkQuickSort benchmark quick sort algorithm
func BenchmarkQuickSort(b *testing.B) {
	testArr := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	for i := 0; i < b.N; i++ {
		sort.QuickSort2(testArr)
	}
}
