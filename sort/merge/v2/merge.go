package v2

import "github.com/corrots/data-structures/sort/helper"

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) {
	if r-l <= 15 {
		helper.InsertionSort(nums, l, r)
		return
	}
	mid := (l + r) / 2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	if nums[mid] > nums[mid+1] {
		merge(nums, l, mid, r)
	}
}

func merge(nums []int, l, mid, r int) {
	temp := make([]int, (r-l)+1)
	copy(temp, nums[l:r+1])
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			nums[k] = temp[j-l]
			j++
		} else if j > r {
			nums[k] = temp[i-l]
			i++
		} else if temp[i-l] < temp[j-l] {
			nums[k] = temp[i-l]
			i++
		} else {
			nums[k] = temp[j-l]
			j++
		}
	}
}
