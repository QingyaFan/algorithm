package sort

func QuickSort(arr []int) []int {
	arrLength := len(arr)
	var lessArr, largerArr []int
	if arrLength <= 1 {
		return arr
	} else {
		pivot := arr[0]
		for i := 1; i < arrLength; i++ {
			if arr[i] <= pivot {
				lessArr = append(lessArr, arr[i])
			} else {
				largerArr = append(largerArr, arr[i])
			}
		}
		return append(append(QuickSort(lessArr), pivot), QuickSort(largerArr)...)
	}
}
