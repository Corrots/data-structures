package main

import "fmt"

func main() {
	nums := []int{7, 5, 6, 4}
	fmt.Println(reversePairs(nums))
	nums1 := []int{1, 3, 2, 3, 1}
	fmt.Println(reversePairs(nums1))
}

// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
// 归并排序思想解法
func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) int {
	if l >= r {
		return 0
	}
	mid := (r-l)/2 + l
	var res int
	res += mergeSort(nums, l, mid)
	res += mergeSort(nums, mid+1, r)
	if nums[mid] > nums[mid+1] {
		res += merge(nums, l, mid, r)
	}
	return res
}

func merge(nums []int, l, mid, r int) int {
	tmp := make([]int, r-l+1)
	copy(tmp, nums[l:r+1])
	i, j := l, mid+1
	var res int
	for k := l; k <= r; k++ {
		if i > mid {
			nums[k] = tmp[j-l]
			j++
		} else if j > r || tmp[i-l] <= tmp[j-l] {
			nums[k] = tmp[i-l]
			i++
		} else {
			res += mid - i + 1
			nums[k] = tmp[j-l]
			j++
		}
	}
	return res
}
