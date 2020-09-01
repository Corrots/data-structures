package main

import (
	"fmt"

	v3 "github.com/corrots/data-structures/sort/merge/v3"

	"github.com/corrots/data-structures/sort/heapsort"
	"github.com/corrots/data-structures/sort/helper"
	v4 "github.com/corrots/data-structures/sort/quick/v4"
)

func main() {
	count := 1000000
	nums1 := helper.GenerateRandArray(count, 0, count*100)
	nums2 := make([]int, count)
	copy(nums2, nums1)
	nums3 := make([]int, count)
	copy(nums3, nums1)
	t1 := helper.TestSort(nums1, heapsort.HeapSort)
	fmt.Printf("HeapSort spend %d ms\n", t1)
	t2 := helper.TestSort(nums2, v4.QuickSort)
	fmt.Printf("QuickSort 2ways spend %d ms\n", t2)
	t3 := helper.TestSort(nums1, v3.MergeSort)
	fmt.Printf("MergeSort spend %d ms\n", t3)
}
