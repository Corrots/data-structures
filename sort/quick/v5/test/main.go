package main

import (
	"fmt"

	v5 "github.com/corrots/data-structures/sort/quick/v5"

	v4 "github.com/corrots/data-structures/sort/quick/v4"

	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	count := 1000000
	nums3 := helper.GenerateRandArray(count, 0, 1)
	nums4 := make([]int, count)
	copy(nums4, nums3)
	t3 := helper.TestSort(nums3, v4.QuickSort)
	fmt.Printf("QuickSort v4 spend %d ms\n", t3)
	t4 := helper.TestSort(nums4, v5.QuickSort)
	fmt.Printf("QuickSort v5 spend %d ms\n", t4)
}
