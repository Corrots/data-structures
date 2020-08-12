package main

import (
	"fmt"

	v4 "github.com/corrots/data-structures/sort/quick/v4"

	v2 "github.com/corrots/data-structures/sort/quick/v2"

	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	count := 100000
	nums3 := helper.GenerateRandArray(count, 0, 1)
	nums4 := make([]int, count)
	copy(nums4, nums3)
	t3 := helper.TestSort(nums3, v2.QuickSort)
	fmt.Printf("QuickSort v2 spend %d ms\n", t3)
	t4 := helper.TestSort(nums4, v4.QuickSort)
	fmt.Printf("QuickSort v4 spend %d ms\n", t4)
}
