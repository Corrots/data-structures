package v3

func MergeSort(nums []int) {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	mergeSort(nums, 0, len(nums)-1, tmp)
}

func mergeSort(nums []int, l, r int, tmp []int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergeSort(nums, l, mid, tmp)
	mergeSort(nums, mid+1, r, tmp)
	if nums[mid] > nums[mid+1] {
		merge(nums, l, mid, r, tmp)
	}
}

func merge(nums []int, l, mid, r int, tmp []int) {
	copy(tmp[l:r+1], nums[l:r+1])
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			nums[k] = tmp[j]
			j++
		} else if j > r {
			nums[k] = tmp[i]
			i++
		} else if tmp[i] < tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
		}
	}
}
