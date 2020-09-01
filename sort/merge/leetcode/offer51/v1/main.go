package main

import "fmt"

func main() {
	nums := []int{7, 5, 6, 4}
	fmt.Println(reversePairs(nums))
}

// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func reversePairs(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				res++
			}
		}
	}
	return res
}
