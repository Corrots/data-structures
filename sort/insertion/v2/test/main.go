package main

import (
	"fmt"

	v2 "github.com/corrots/data-structures/sort/insertion/v2"
)

func main() {
	nums := []int{3, 2, 5, 1, 4, 0}
	v2.InsertionSort(nums)
	fmt.Println(nums)
}
