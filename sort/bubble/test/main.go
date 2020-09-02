package main

import (
	"fmt"

	v1 "github.com/corrots/data-structures/sort/bubble/v1"
	v2 "github.com/corrots/data-structures/sort/bubble/v2"
	v3 "github.com/corrots/data-structures/sort/bubble/v3"
	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	count := 20000
	nums1 := helper.GenerateRandArray(count, 0, count*10)
	nums2 := make([]int, count)
	copy(nums2, nums1)
	t1 := helper.TestSort(nums1, v1.BubbleSort)
	fmt.Printf("bubble sort v1 spent: %d ms\n", t1)
	t2 := helper.TestSort(nums2, v2.BubbleSort)
	fmt.Printf("bubble sort v2 spent: %d ms\n", t2)

	fmt.Println("优化1")
	//nums3 := helper.GenerateNearlyOrderedArray(count, 0, count*10)
	nums3 := helper.GenerateOrderedArray(count)
	nums4 := make([]int, count)
	copy(nums4, nums3)
	t3 := helper.TestSort(nums3, v1.BubbleSort)
	fmt.Printf("bubble sort v1 spent: %d ms\n", t3)
	t4 := helper.TestSort(nums4, v2.BubbleSort)
	fmt.Printf("bubble sort v2 spent: %d ms\n", t4)

	fmt.Println("优化2")
	nums5 := helper.GenerateNearlyOrderedArray(count, 0, count*10)
	nums6 := make([]int, count)
	copy(nums6, nums5)
	t5 := helper.TestSort(nums5, v2.BubbleSort)
	fmt.Printf("bubble sort v2 spent: %d ms\n", t5)
	t6 := helper.TestSort(nums6, v3.BubbleSort)
	fmt.Printf("bubble sort v3 spent: %d ms\n", t6)
}
