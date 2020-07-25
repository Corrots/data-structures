package main

import (
	"fmt"

	"github.com/corrots/data-structures/queue/queue"
)

func main() {
	queue := queue.NewArrayQueue()
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
