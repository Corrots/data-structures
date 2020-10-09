package main

import (
	"fmt"

	segmenttree "github.com/corrots/data-structures/segment-tree"
)

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	segTree := segmenttree.NewSegmentTree(nums, sumFunc)
	fmt.Println(segTree)
	fmt.Println(segTree.Query(0, 2))
	fmt.Println(segTree.Query(2, 5))
	fmt.Println(segTree.Query(0, 5))
}

func sumFunc(v1, v2 int) int {
	return v1 + v2
}
