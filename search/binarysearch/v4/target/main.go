package main

import "fmt"

func main() {
	nums := []int{1, 1, 3, 3, 5, 5}
	fmt.Println(BinarySearch(nums, 1))
	fmt.Println(BinarySearch(nums, 3))
	fmt.Println(BinarySearch(nums, 5))
	fmt.Println(BinarySearch(nums, 4))
	fmt.Println(BinarySearch(nums, 10))
}

// 用>=target的最小值的思路，实现二分查找法
func BinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l < len(nums) && nums[l] == target {
		return l
	}
	return -1
}
