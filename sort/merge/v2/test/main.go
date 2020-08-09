package main

import (
	"fmt"

	"github.com/corrots/data-structures/sort/helper"
	v2 "github.com/corrots/data-structures/sort/merge/v2"
)

func main() {
	//nums := []int{3, 2, 5, 1, 4, 0}
	//MergeSort(nums)
	//fmt.Println(nums)
	opCount := 50
	t1 := helper.TestSort(opCount, v2.MergeSort)
	fmt.Printf("MergeSort spend %d ms", t1)
}
