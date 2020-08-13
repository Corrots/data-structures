package v6

// 修改循环不变量: nums[l:r)
func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums))
}

// 对nums[l:r)进行归并排序
func mergeSort(nums []int, l, r int) {
	if r-l <= 1 {
		return
	}
	mid := (r-l)/2 + l
	mergeSort(nums, l, mid) // nums[l:mid)
	mergeSort(nums, mid, r) // nums[mid:r)
	if nums[mid-1] > nums[mid] {
		merge(nums, l, mid, r)
	}
}

// 合并nums[l:mid)和nums[mid:r)中的元素
func merge(nums []int, l, mid, r int) {
	tmp := make([]int, r-l)
	copy(tmp, nums[l:r])
	i, j := l, mid
	for k := l; k < r; k++ {
		if i >= mid {
			nums[k] = tmp[j-l]
			j++
		} else if j >= r {
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
