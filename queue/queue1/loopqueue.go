package queue1

import (
	"fmt"
	"log"
	"strings"

	"github.com/corrots/data-structures/queue/basic"
)

// 不浪费一个空间的循环队列
type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func NewLoopQueue(cap int) basic.Queue {
	return &LoopQueue{
		data: make([]interface{}, cap),
	}
}

func InitLoopQueue() basic.Queue {
	return NewLoopQueue(10)
}

func (q *LoopQueue) resize(newCap int) {
	newData := make([]interface{}, newCap)
	for i := 0; i < q.size; i++ {
		newData[i] = q.data[(q.front+i)%len(q.data)]
	}
	q.data = newData
	q.front = 0
	q.tail = q.size
}

// 入队
func (q *LoopQueue) Enqueue(e interface{}) {
	if q.size == q.Cap() {
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
	return q.size == 0
}

func (q *LoopQueue) Cap() int {
	return len(q.data)
}

func (q *LoopQueue) String() string {
	var res strings.Builder
	res.WriteString(fmt.Sprintf("Queue: size: %d, capacity: %d\n", q.size, q.Cap()))
	res.WriteString("front [")
	for i := 0; i < q.size; i++ {
		res.WriteString(fmt.Sprintf("%v", q.data[(q.front+i)%len(q.data)]))
		if i != q.size-1 {
			res.WriteString(", ")
		}
	}
	res.WriteString("] tail")
	return res.String()
}
