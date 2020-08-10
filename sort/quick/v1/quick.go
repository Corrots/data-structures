package v1

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func partition(nums []int, l, r int) int {
	v := nums[l]
	p := l
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			nums[p+1], nums[i] = nums[i], nums[p+1]
			p++
		}
	}
	nums[l], nums[p] = nums[p], nums[l]
	return p
}
