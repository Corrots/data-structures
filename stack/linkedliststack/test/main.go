package main

import (
	"fmt"

	"github.com/corrots/data-structures/stack/linkedliststack"
)

func main() {
	s := linkedliststack.NewStack()
	for i := 0; i < 10; i++ {
		s.Push(i)
		fmt.Println(s)
	}
	for !s.IsEmpty() {
		s.Pop()
		fmt.Println(s)
	}
}
