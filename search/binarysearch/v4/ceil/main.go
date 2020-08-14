package main

import "fmt"

func main() {
	nums := []int{1, 1, 3, 3, 5, 5}
	//for i := 0; i <= 6; i++ {
	//	fmt.Printf("%d ", Upper(nums, i))
	//}

	//for i := 0; i <= 6; i++ {
	//	fmt.Printf("%d ", UpperCeil(nums, i))
	//}
	//fmt.Println()

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

// >= target的索引值，若找到多个=target元素，返回索引最小值
func UpperCeil(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if r-1 >= 0 && nums[r-1] == target {
		return r - 1
	}
	return r
}

// nums := []int{1, 1, 3, 3, 5, 5}
// >= target的最小索引值
func LowerCeil(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r
}
