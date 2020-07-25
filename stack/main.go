package main

import (
	"fmt"

	"github.com/corrots/data-structures/stack/arraystack"
)

func main() {
	stack := arraystack.NewStack(10)
	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}
	fmt.Println("Pop: ", stack.Pop())
	fmt.Println("Peek: ", stack.Peek())
}
