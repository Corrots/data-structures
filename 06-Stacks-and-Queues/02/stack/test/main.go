package main

import (
	"fmt"
	"github.com/corrots/data-structures/06-Stacks-and-Queues/02/stack"
)

func main() {
	s := stack.New(10)
	s.Push("abc")
	fmt.Println(s.Peek())
}
