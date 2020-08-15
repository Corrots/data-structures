package main

import "fmt"

func main() {
	nums := []int{1, 1, 3, 3, 5, 5}
	//fmt.Printf("%d ", Lower(nums, 2))
	//for i := 0; i <= 6; i++ {
	//	fmt.Printf("%d ", Lower(nums, i))
	//}
	fmt.Println(LowerFloor(nums, 1))
	fmt.Println(LowerFloor(nums, 3))
	fmt.Println(LowerFloor(nums, 5))
}

// <target的最大值的索引
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

//
func LowerFloor(nums []int, target int) int {
	i := Lower(nums, target)
	if i+1 < len(nums) && nums[i+1] == target {
		return i + 1
	}
	return i
}

// nums := []int{1, 1, 3, 3, 5, 5}
// <=target的最大索引
//func UpperFloor(nums []int, target int) int {
//
//}
