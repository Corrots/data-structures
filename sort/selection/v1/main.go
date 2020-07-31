package main

import "fmt"

func main() {
	nums := []int{3, 2, 5, 1, 4, 0}
	SelectionSort(nums)
	fmt.Println(nums)
}

func SelectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
