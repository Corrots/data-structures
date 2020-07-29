package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/corrots/data-structures/queue/linkedlistqueue"

	"github.com/corrots/data-structures/queue/basic"

	"github.com/corrots/data-structures/queue/queue"
)

func main() {
	opCount := 1000000
	//aq := queue.NewArrayQueue()
	//fmt.Printf("array queue spent: %d ms\n", TestQueue(aq, opCount))
	lq := queue.InitLoopQueue()
	fmt.Printf("queue1 queue spent: %d ms\n", TestQueue(lq, opCount))
	llq := linkedlistqueue.NewLinkedListQueue()
	fmt.Printf("linked list queue spent: %d ms\n", TestQueue(llq, opCount))
}

// 测试使用q运行opCount个enqueue和dequeue操作所需的时间，单位：s
func TestQueue(q basic.Queue, opCount int) int64 {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < opCount; i++ {
		q.Enqueue(rand.Int() % (opCount))
	}
	for i := 0; i < opCount; i++ {
		q.Dequeue()
	}
	return time.Since(start).Milliseconds()
	//fmt.Printf("spent: %d s")
}
