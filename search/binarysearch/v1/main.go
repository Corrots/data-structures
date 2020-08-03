package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 9}
	fmt.Println(BinarySearch(nums, 0))
}

func BinarySearch(nums []int, target int) int {
	// 循环不变量：nums[l:r]范围内寻找target
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
