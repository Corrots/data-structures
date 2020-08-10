package main

import (
	"fmt"

	v1 "github.com/corrots/data-structures/sort/quick/v1"

	"github.com/corrots/data-structures/sort/helper"
	v2 "github.com/corrots/data-structures/sort/merge/v2"
)

func main() {
	count := 3000000
	nums1 := helper.GenerateRandArray(count, 0, count*10)
	nums2 := make([]int, count)
	copy(nums2, nums1)
	t1 := helper.TestSort(nums1, v2.MergeSort)
	fmt.Printf("MergeSort v2 spend %d ms\n", t1)
	t2 := helper.TestSort(nums2, v1.QuickSort)
	fmt.Printf("QuickSort v1 spend %d ms\n", t2)
}
