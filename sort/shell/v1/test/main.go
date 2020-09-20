package main

import (
	"fmt"
	"math"

	v4 "github.com/corrots/data-structures/sort/merge/v4"

	"github.com/corrots/data-structures/sort/helper"

	v1 "github.com/corrots/data-structures/sort/shell/v1"
)

func main() {
	nums := []int{4, 3, 2, 7, 5, 1, 6, 8}
	v1.ShellSort(nums)
	fmt.Println(nums)
	count := 1000000
	nums1 := helper.GenerateRandArray(count, 0, math.MaxInt64)
	nums2 := make([]int, count)
	copy(nums2, nums1)
	//v1.ShellSort(nums1)
	t1 := helper.TestSort(nums1, v1.ShellSort)
	fmt.Printf("ShellSort: %d ms\n", t1)
	t2 := helper.TestSort(nums2, v4.MergeSort)
	fmt.Printf("MergeSort: %d ms\n", t2)
}
