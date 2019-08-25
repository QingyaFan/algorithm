package sort

// QuickSort sort an array using quicksort algorithm
// return an sorted array
func QuickSort(arr []int) []int {

	length := len(arr)
	var leftSet, rightSet []int

	if length <= 1 {
		return arr
	}

	pivot := arr[0]
	for i := 1; i < length; i++ {
		if arr[i] <= pivot {
			leftSet = append(leftSet, arr[i])
		} else {
			rightSet = append(rightSet, arr[i])
		}
	}

	return append(append(QuickSort(leftSet), pivot), QuickSort(rightSet)...)

}

// QuickSort2 merge sort algorithm using less memory
func QuickSort2(arr []int) []int {

	length := len(arr)

	if length <= 1 {
		return arr
	}

	pivot := arr[0]
	leftI, rightI := 0, length-1

	for leftI != rightI {
		for arr[rightI] >= pivot && leftI < rightI {
			rightI = rightI - 1
		}
		for arr[leftI] < pivot && leftI < rightI {
			leftI = leftI + 1
		}

		arr[leftI], arr[rightI] = arr[rightI], arr[leftI]
	}
	arr[leftI], arr[0] = arr[0], arr[leftI]

	return append(QuickSort2(arr[0:leftI+1]), QuickSort2(arr[leftI+1:])...)

}
