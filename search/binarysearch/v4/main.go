package main

import "fmt"

func main() {
	nums := []int{1, 1, 3, 3, 5, 5}
	for i := 0; i <= 6; i++ {
		fmt.Printf("%d ", Upper(nums, i))
	}

	fmt.Println()
	for i := 0; i <= 6; i++ {
		fmt.Printf("%d ", Ceil(nums, i))
	}
	fmt.Println()

	for i := 0; i <= 6; i++ {
		fmt.Printf("%d ", LowerCeil(nums, i))
	}
	fmt.Println()
}

// 查找大于target的最小值的索引
func Upper(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func Ceil(nums []int, target int) int {
	i := Upper(nums, target)
	if i-1 >= 0 && nums[i-1] == target {
		return i - 1
	}
	return i
}

// >= target的最小索引值
func LowerCeil(nums []int, target int) int {
	u := Upper(nums, target)
	res := u
	for i := u; i >= 0 && i < len(nums); i-- {
		if nums[i] == target {
			res = i
		}
	}
	return res
}
