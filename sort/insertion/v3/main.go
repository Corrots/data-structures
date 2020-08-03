package main

import (
	"fmt"

	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	nums := []int{3, 2, 5, 1, 4, 0}
	InsertionSort(nums)
	fmt.Println(nums)

	opCount := 100000
	t2 := helper.TestSort(opCount, InsertionSort)
	fmt.Printf("InsertionSort spend %d ms", t2)
}

// 从后向前实现插入排序
// 小优化版
func InsertionSort(nums []int) {
	n := len(nums)
	// 循环不变量: nums[0,i)未排序, nums[i...n)已排序
	for i := n - 2; i >= 0; i-- {
		e := nums[i]
		j := 0
		for j = i; j < n-1 && nums[j+1] < e; j++ {
			nums[j] = nums[j+1]
		}
		nums[j] = e
	}
}
