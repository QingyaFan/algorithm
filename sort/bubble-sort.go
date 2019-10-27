package sort

// bubbleSort 冒泡排序
func bubbleSort(nums []int) []int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		for j := 0; j < (numsLen - i); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
