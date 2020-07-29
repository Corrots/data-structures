package main

import (
	"fmt"

	solution "github.com/corrots/data-structures/queue/leetcode/203"
)

func main() {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	head := solution.NewListNode(nums)
	fmt.Println(head)
	solution.RemoveElements3(head, 6)
	fmt.Println(head)
}
