package main

import (
	"fmt"

	"github.com/corrots/data-structures/queue/deque"
)

func main() {
	dq := deque.Init()
	// 入队
	for i := 0; i < 16; i++ {
		if i%2 == 0 {
			dq.AddLast(i)
		} else {
			dq.AddFront(i)
		}
	}
	fmt.Println(dq)
	// 出队
	for i := 0; !dq.IsEmpty(); i++ {
		if i%2 == 0 {
			dq.RemoveFront()
		} else {
			dq.RemoveLast()
		}
		fmt.Println(dq)
	}
}
