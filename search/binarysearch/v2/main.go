package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 9}
	fmt.Println(BinarySearch(nums, 0))
	fmt.Println(BinarySearch(nums, 3))
	fmt.Println(BinarySearch(nums, 7))
}

// 修改变量区间定义
func BinarySearch(nums []int, target int) int {
	// 循环不变量：nums[l:r)范围内寻找target
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r - 1) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return -1
}
