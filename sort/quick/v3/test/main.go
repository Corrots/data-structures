package main

import (
	"fmt"

	v2 "github.com/corrots/data-structures/sort/quick/v2"

	v1 "github.com/corrots/data-structures/sort/quick/v1"

	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	count := 100000
	//nums1 := helper.GenerateRandArray(count, 0, count*10)
	//nums2 := make([]int, count)
	//copy(nums2, nums1)
	//t1 := helper.TestSort(nums1, v1.QuickSort)
	//fmt.Printf("QuickSort v1 spend %d ms\n", t1)
	//t2 := helper.TestSort(nums2, v2.QuickSort)
	//fmt.Printf("QuickSort v2 spend %d ms\n", t2)

	nums3 := helper.GenerateOrderedArray(count)
	nums4 := make([]int, count)
	copy(nums4, nums3)
	t3 := helper.TestSort(nums3, v1.QuickSort)
	fmt.Printf("QuickSort v1 spend %d ms\n", t3)
	t4 := helper.TestSort(nums4, v2.QuickSort)
	fmt.Printf("QuickSort v2 spend %d ms\n", t4)
}
