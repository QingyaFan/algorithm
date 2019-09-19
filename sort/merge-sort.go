package sort

// mergeSort 归并排序
func mergeSort(arr []int, left int, mid int, right int) {

}

func sort(arr []int, left int, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2
	sort(arr[0:mid], 0, mid)
	sort(arr[(mid+1):], mid+1, right)

	mergeSort(arr, left, mid, right)
}

// MergeSort merge sort
func MergeSort(arr []int) []int {
	sort()
}
