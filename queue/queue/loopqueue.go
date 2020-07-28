package queue

import (
	"fmt"
	"log"
	"strings"

	"github.com/corrots/data-structures/queue/basic"
)

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func NewLoopQueue(cap int) basic.Queue {
	return &LoopQueue{
		data: make([]interface{}, cap+1),
	}
}

func InitLoopQueue() basic.Queue {
	return NewLoopQueue(10)
}

func (q *LoopQueue) resize(newCap int) {
	newData := make([]interface{}, newCap+1)
	for i := 0; i < q.size; i++ {
		newData[i] = q.data[(q.front+i)%len(q.data)]
	}
	q.data = newData
	q.front = 0
	q.tail = q.size
}

// 入队
func (q *LoopQueue) Enqueue(e interface{}) {
	if (q.tail+1)%len(q.data) == q.front {
		q.resize(q.Cap() * 2)
	}
	q.data[q.tail] = e
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
}

// 出队
func (q *LoopQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		log.Fatal("Dequeue failed, queue is empty")
	}
	res := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % len(q.data)
	q.size--
	if q.size <= q.Cap()/4 && q.Cap()/2 != 0 {
		q.resize(q.Cap() / 2)
	}
	return res
}

func (q *LoopQueue) GetFront() interface{} {
	if q.IsEmpty() {
		log.Fatal("queue is empty")
	}
	return q.data[q.front]
}

func (q *LoopQueue) Len() int {
	return q.size
}

func (q *LoopQueue) IsEmpty() bool {
	return q.front == q.tail
}

func (q *LoopQueue) Cap() int {
	return len(q.data) - 1
}

func (q *LoopQueue) String() string {
	var res strings.Builder
	res.WriteString(fmt.Sprintf("Queue: size: %d, capacity: %d\n", q.size, q.Cap()))
	res.WriteString("front [")
	for i := q.front; i != q.tail; i = (i + 1) % len(q.data) {
		res.WriteString(fmt.Sprintf("%v", q.data[i]))
		if (i+1)%len(q.data) != q.tail {
			res.WriteString(", ")
		}
	}
	res.WriteString("] tail")
	return res.String()
}
