package main

import (
	"fmt"

	v1 "github.com/corrots/data-structures/sort/merge/v1"
	v2 "github.com/corrots/data-structures/sort/merge/v2"

	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	count := 1000000
	nums1 := helper.GenerateRandArray(count, 0, count*10)
	nums2 := make([]int, count)
	copy(nums2, nums1)
	t1 := helper.TestSort(nums1, v1.MergeSort)
	fmt.Printf("MergeSort1 spend %d ms\n", t1)
	t2 := helper.TestSort(nums2, v2.MergeSort)
	fmt.Printf("MergeSort2 spend %d ms\n", t2)
}
