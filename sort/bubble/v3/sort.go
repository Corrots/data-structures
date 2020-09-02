package v3

// 冒泡排序-优化2
func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; {
		lastSwappedIndex := 0
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				lastSwappedIndex = j + 1
			}
		}
		i = n - lastSwappedIndex
	}
}
