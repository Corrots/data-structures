package main

import "fmt"

func main() {
	nums := []int{3, 2, 5, 1, 4, 0}
	InsertionSort(nums)
	fmt.Println(nums)
}

// 从后向前实现插入排序
func InsertionSort(nums []int) {
	n := len(nums)
	// 循环不变量: nums[0,i)未排序, nums[i...n)已排序
	for i := n - 2; i >= 0; i-- {
		for j := i; j < n-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			} else {
				break
			}
		}
	}
}
