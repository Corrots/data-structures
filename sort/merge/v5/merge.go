package v5

import "github.com/corrots/data-structures/sort/helper"

// 归并排序的自底向上实现
// 小规模数据使用insertion sort
func MergeSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i += 16 {
		helper.InsertionSort(nums, i, min(n-1, i+15))
	}

	for sz := 16; sz < n; sz += sz {
		// 合并区间[i:i+sz-1]和[i+sz:i+sz+sz-1]
		// min(n-1, i+sz+sz-1)
		for i := 0; i+sz < n; i += sz + sz {
			if nums[i+sz-1] > nums[i+sz] {
				merge(nums, i, i+sz-1, min(n-1, i+sz+sz-1))
			}
		}
	}
}

func merge(nums []int, l, mid, r int) {
	tmp := make([]int, r-l+1)
	copy(tmp, nums[l:r+1])
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			nums[k] = tmp[j-l]
			j++
		} else if j > r {
			nums[k] = tmp[i-l]
			i++
		} else if tmp[i-l] < tmp[j-l] {
			nums[k] = tmp[i-l]
			i++
		} else {
			nums[k] = tmp[j-l]
			j++
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
