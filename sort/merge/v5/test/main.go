package main

import (
	"fmt"

	v5 "github.com/corrots/data-structures/sort/merge/v5"

	v4 "github.com/corrots/data-structures/sort/merge/v4"

	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	count := 5000000
	nums1 := helper.GenerateRandArray(count, 0, count*10)
	nums2 := make([]int, count)
	copy(nums2, nums1)
	t1 := helper.TestSort(nums1, v4.MergeSort)
	fmt.Printf("MergeSort4 spend %d ms\n", t1)
	t2 := helper.TestSort(nums2, v5.MergeSort)
	fmt.Printf("MergeSort5 spend %d ms\n", t2)
}
