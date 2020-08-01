package main

import "fmt"

func main() {
	nums := []int{3, 2, 5, 1, 4, 0}
	SelectionSort(nums)
	fmt.Println(nums)
}

// 选择排序的另一种实现，从后向前
func SelectionSort(nums []int) {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		maxIndex := i
		for j := i; j >= 0; j-- {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[i], nums[maxIndex] = nums[maxIndex], nums[i]
	}
}
