package main

import (
	"fmt"

	v3 "github.com/corrots/data-structures/sort/merge/v3"

	v4 "github.com/corrots/data-structures/sort/merge/v4"

	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	count := 1000000
	nums1 := helper.GenerateRandArray(count, 0, count*10)
	nums2 := make([]int, count)
	copy(nums2, nums1)
	t1 := helper.TestSort(nums1, v4.MergeSort)
	fmt.Printf("MergeSort4 spend %d ms\n", t1)
	t2 := helper.TestSort(nums2, v3.MergeSort)
	fmt.Printf("MergeSort3 spend %d ms\n", t2)
}
