package main

import (
	"fmt"

	v2 "github.com/corrots/data-structures/bst/v2"
)

func main() {
	bst := v2.NewBST()
	nums := []int{5, 3, 6, 8, 4, 2}
	for _, v := range nums {
		bst.Add(v)
	}
	//bst.PreOrder()
	//fmt.Println()
	//bst.PreOrderNR()
	//bst.InOrder()
	//fmt.Println()
	//bst.PostOrder()
	//fmt.Println(bst)
	// 层序遍历
	bst.LevelOrder()
	fmt.Println()
	//fmt.Println(bst.Minimum())
	//fmt.Println(bst.Maximum())
	//fmt.Println(bst.RemoveMin())
	//fmt.Println(bst.RemoveMin())
	//fmt.Println(bst.RemoveMax())
	//fmt.Println(bst.RemoveMax())
	bst.Remove(4)
	bst.LevelOrder()
	fmt.Println()
	bst.Remove(6)
	bst.LevelOrder()
}
