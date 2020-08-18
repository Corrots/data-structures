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
	bst.PreOrder()
	fmt.Println()
	bst.PreOrderNR()
	//bst.InOrder()
	//fmt.Println()
	//bst.PostOrder()
	//fmt.Println(bst)
}
