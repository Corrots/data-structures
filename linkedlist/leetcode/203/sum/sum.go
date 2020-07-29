package sum

func Sum(nums []int) int {
	return sum(nums, 0)
}

func sum(nums []int, l int) int {
	if l == len(nums) {
		return 0
	}
	return nums[l] + sum(nums, l+1)
}
