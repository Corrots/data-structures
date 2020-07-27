package main

import (
	"fmt"

	"github.com/corrots/data-structures/queue/leetcode/225/stack"
)

func main() {
	s := stack.Constructor()
	s.Push(1)
	s.Push(2)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Top())
	fmt.Println(s.Empty())
}
