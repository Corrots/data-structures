package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/corrots/data-structures/queue/queue"
)

func main() {
	opCount := 100000
	aq := queue.NewArrayQueue()
	fmt.Printf("array queue spent: %d ms\n", TestQueue(aq, opCount))
	lq := queue.InitLoopQueue()
	fmt.Printf("loop queue spent: %d ms\n", TestQueue(lq, opCount))
}

// 测试使用q运行opCount个enqueue和dequeue操作所需的时间，单位：s
func TestQueue(q queue.Queue, opCount int) int64 {
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
