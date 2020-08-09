package main

import (
	"fmt"

	v2 "github.com/corrots/data-structures/sort/merge/v2"

	v3 "github.com/corrots/data-structures/sort/merge/v3"

	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	opCount := 5000000
	t1 := helper.TestSort(opCount, v3.MergeSort)
	fmt.Printf("MergeSort3 spend %d ms\n", t1)
	t2 := helper.TestSort(opCount, v2.MergeSort)
	fmt.Printf("MergeSort2 spend %d ms\n", t2)
}
