package main

import (
	"fmt"

	"github.com/corrots/data-structures/queue/linkedlistqueue"
)

func main() {
	llq := linkedlistqueue.NewLinkedListQueue()
	for i := 0; i < 10; i++ {
		llq.Enqueue(i)
		fmt.Println(llq)
		if i%3 == 2 {
			llq.Dequeue()
			fmt.Println(llq)
		}
	}
}
