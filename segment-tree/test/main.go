package main

import (
	"fmt"

	segmenttree "github.com/corrots/data-structures/segment-tree"
)

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	segTree := segmenttree.NewSegmentTree(nums, sumFunc)
	fmt.Println(segTree)
}

func sumFunc(v1, v2 int) int {
	return v1 + v2
}
