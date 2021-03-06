package main

import (
	"fmt"

	v2 "github.com/corrots/data-structures/sort/merge/v2"
	v3 "github.com/corrots/data-structures/sort/merge/v3"

	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	count := 1000000
	nums1 := helper.GenerateRandArray(count, 0, count*10)
	nums2 := make([]int, count)
	copy(nums2, nums1)
	t1 := helper.TestSort(nums1, v2.MergeSort)
	fmt.Printf("MergeSort2 spend %d ms\n", t1)
	t2 := helper.TestSort(nums2, v3.MergeSort)
	fmt.Printf("MergeSort3 spend %d ms\n", t2)
}
