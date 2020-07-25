package main

import (
	"fmt"

	"github.com/corrots/data-structures/queue/queue"
)

func main() {
	q := queue.InitLoopQueue()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
		fmt.Println(q)
		if i%3 == 2 {
			q.Dequeue()
			fmt.Println(q)
		}
	}
}