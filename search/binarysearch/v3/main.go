package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 9}
	fmt.Println(BinarySearch(nums, 10))
}

// 二分查找法的递归实现
func BinarySearch(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, l, r, target int) int {
	if l > r {
		return -1
	}
	mid := (r-l)/2 + l
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binarySearch(nums, l, mid-1, target)
	}
	return binarySearch(nums, mid+1, r, target)
}
