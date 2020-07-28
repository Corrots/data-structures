package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/corrots/data-structures/stack"
	"github.com/corrots/data-structures/stack/arraystack"
)

func main() {
	opCount := 1000000
	as := arraystack.NewStack(10)
	t1 := TestStack(as, opCount)
	fmt.Printf("array stack spent: %d ms\n", t1)
	ls := arraystack.NewStack(10)
	t2 := TestStack(ls, opCount)
	fmt.Printf("linked list stack spent: %d ms\n", t2)
}

func TestStack(s stack.Stack, opCount int) int64 {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < opCount; i++ {
		s.Push(rand.Intn(opCount))
	}
	for i := 0; i < opCount; i++ {
		s.Pop()
	}
	return time.Since(start).Milliseconds()
}
