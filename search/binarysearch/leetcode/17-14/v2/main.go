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
	var res []int
	// nums[l:r)
	l, r := 0, len(arr)
	for l < r {
		p := partition(arr, l, r)
		if p == k {
			return append(res, arr[:p]...)
		} else if p < k {
			l = p + 1
		} else {
			r = p
		}
	}
	return nil
}

func randPivot(nums []int, l, r int) {
	randIndex := rand.Int()%(r-l) + l
	nums[randIndex], nums[l] = nums[l], nums[randIndex]
}

// nums[l, r)
func partition(nums []int, l, r int) int {
	randPivot(nums, l, r)
	v := nums[l]
	j := l
	for i := l + 1; i < r; i++ {
		if nums[i] < v {
			j++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
