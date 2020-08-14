package main

import "fmt"

func main() {
	nums := []int{1, 1, 3, 3, 5, 5}
	//fmt.Printf("%d ", Lower(nums, 2))
	for i := 0; i <= 6; i++ {
		fmt.Printf("%d ", Lower(nums, i))
	}
}

// <target的最大值索引
func Lower(nums []int, target int) int {
	l, r := -1, len(nums)-1
	for l < r {
		//fmt.Printf("l: %d, r: %d\n", l, r)
		mid := (r-l+1)/2 + l
		if nums[mid] < target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
