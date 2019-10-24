package miscellaneous

// 找出数组中重复元素的首次出现位置，利用二分查找
// 数组中元素有序，且可能存在重复
func findRepeatElement(nums []int, target int) int {
	arrLen := len(nums)
	high, low := arrLen-1, 0
	for low < high {
		middle := (low + high) / 2

		if nums[middle] == target {
			// 关键，克服重复元素的关键
			// 如果找到了val,则继续对low~mid段数据进行二分查找
			// 此处能low = middle，否则将找到target最后出现的位置(当然还有第九行需要改成math.ceil(middle))
			high = middle
		}

		if nums[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}

	}

	if nums[high] == target {
		return high
	}

	return -1
}
