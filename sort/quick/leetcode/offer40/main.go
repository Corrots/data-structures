package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums1 := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(getLeastNumbers(nums1, 4))
	nums2 := []int{0, 0, 2, 3, 2, 1, 1, 2, 0, 4}
	fmt.Println(getLeastNumbers(nums2, 10))
}

// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func getLeastNumbers(arr []int, k int) []int {
	rand.Seed(time.Now().UnixNano())
	if k >= len(arr) {
		k = len(arr)
	}
	return quickSelect(arr, 0, len(arr)-1, k)
}

func quickSelect(nums []int, l, r, k int) []int {
	var res []int
	if l > r {
		return res
	}
	p := partition(nums, l, r)
	if p == k {
		return append(res, nums[:p]...)
	} else if p < k {
		return quickSelect(nums, p+1, r, k)
	}
	return quickSelect(nums, l, p-1, k)
}

func randPivot(nums []int, l, r int) {
	randIndex := rand.Int()%(r-l+1) + l
	nums[l], nums[randIndex] = nums[randIndex], nums[l]
}

func partition(nums []int, l, r int) int {
	randPivot(nums, l, r)
	v := nums[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			j++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[j], nums[l] = nums[l], nums[j]
	return j
}
