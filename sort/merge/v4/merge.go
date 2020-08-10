package v4

// 归并排序的自底向上实现
func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (r-l)/2 + l
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	if nums[mid] > nums[mid+1] {
		merge(nums, l, mid, r)
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
