package v4

// 归并排序的自底向上实现
func MergeSort(nums []int) {
	//mergeSort(nums, 0, len(nums)-1)
	n := len(nums)
	// 遍历合并的区间长度
	for sz := 1; sz < n; sz += sz {
		// 遍历合并的两个区间的起始位置i
		// 合并 [i, i+sz-1]和[i+sz, min(n-1, i+sz+sz-1)]
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
