package main

import (
	"fmt"
	"math"

	"github.com/corrots/data-structures/sort/helper"

	v1 "github.com/corrots/data-structures/sort/shell/v1"
	v2 "github.com/corrots/data-structures/sort/shell/v2"
)

func main() {
	nums := []int{4, 3, 2, 7, 5, 1, 6, 8}
	v1.ShellSort(nums)
	fmt.Println(nums)
	count := 5000000
	nums1 := helper.GenerateRandArray(count, 0, math.MaxInt64)
	nums2 := make([]int, count)
	copy(nums2, nums1)
	//v1.ShellSort(nums1)
	t1 := helper.TestSort(nums1, v2.ShellSort)
	fmt.Printf("ShellSort: %d ms\n", t1)
	t2 := helper.TestSort(nums2, v2.ShellSort2)
	fmt.Printf("ShellSort2: %d ms\n", t2)
}
