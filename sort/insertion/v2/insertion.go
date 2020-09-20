package v2

func InsertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		e := nums[i]
		j := 0
		for j = i; j > 0 && nums[j-1] > e; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = e
	}
}
