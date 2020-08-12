package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums1 := []int{3, 2, 1, 5, 6, 4}
	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums1, 2))
	fmt.Println(findKthLargest(nums2, 4))
}

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// 递归实现
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, l, r, k int) int {
	p := partition(nums, l, r)
	if k == p {
		return nums[p]
	} else if k > p {
		return quickSort(nums, p+1, r, k)
	}
	return quickSort(nums, l, p-1, k)
}

func randomPivot(nums []int, l, r int) {
	randomIndex := rand.Int()%(r-l+1) + l
	nums[randomIndex], nums[l] = nums[l], nums[randomIndex]
}

func partition(nums []int, l, r int) int {
	randomPivot(nums, l, r)
	v := nums[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			nums[i], nums[j+1] = nums[j+1], nums[i]
			j++
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
