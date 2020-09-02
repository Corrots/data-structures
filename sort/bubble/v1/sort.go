package v1

// 冒泡排序-基础版
func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i+1 < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
