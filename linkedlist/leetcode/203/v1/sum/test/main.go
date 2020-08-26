package main

import (
	"fmt"

	"github.com/corrots/data-structures/linkedlist/leetcode/203/sum"
)

func main() {
	var nums []int
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
	}
	fmt.Println(sum.Sum(nums))

}
