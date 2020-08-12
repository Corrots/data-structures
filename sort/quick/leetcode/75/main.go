package main

import "fmt"

func main() {
	nums1 := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums1)
	fmt.Println(nums1)
}

// https://leetcode-cn.com/problems/sort-colors/
func sortColors(nums []int) {
	lt, gt := -1, len(nums)
	i := 0
	for i < gt {
		if nums[i] == 1 {
			i++
		} else if nums[i] == 0 {
			lt++
			nums[i], nums[lt] = nums[lt], nums[i]
			i++
		} else {
			gt--
			nums[i], nums[gt] = nums[gt], nums[i]
		}
	}
}
