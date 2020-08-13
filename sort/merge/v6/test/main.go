package main

import (
	"fmt"

	"github.com/corrots/data-structures/sort/helper"
	v6 "github.com/corrots/data-structures/sort/merge/v6"
)

func main() {
	nums := []int{2, 3, 6, 8, 1, 4, 5, 7}
	t := helper.TestSort(nums, v6.MergeSort)
	fmt.Printf("MergeSort v6 spend %d ms\n", t)
	count := 1000000
	nums1 := helper.GenerateRandArray(count, 0, count*10)
	t1 := helper.TestSort(nums1, v6.MergeSort)
	fmt.Printf("MergeSort v6 spend %d ms\n", t1)
}
