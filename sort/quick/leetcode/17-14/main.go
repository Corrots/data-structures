package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums1 := []int{1, 3, 5, 7, 2, 4, 6, 8}
	fmt.Println(smallestK(nums1, 4))
}

// https://leetcode-cn.com/problems/smallest-k-lcci/
func smallestK(arr []int, k int) []int {
	if k >= len(arr)-1 {
		return arr
	}
	rand.Seed(time.Now().UnixNano())
	return quickSelect(arr, 0, len(arr)-1, k)
}

func quickSelect(nums []int, l, r, k int) (res []int) {
	if l > r {
		return
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
	nums[randIndex], nums[l] = nums[l], nums[randIndex]
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
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
