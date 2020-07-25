package queue

import "log"

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func NewLoopQueue(cap int) Queue {
	return &LoopQueue{
		data:  make([]interface{}, cap+1),
		front: 0,
		tail:  0,
		size:  0,
	}
}

func InitLoopQueue() Queue {
	return NewLoopQueue(10)
}

func (q *LoopQueue) resize(newCap int) {
	newData := make([]interface{}, newCap+1)
	for i := 0; i < q.size; i++ {
		// 新数组中索引为i的元素 = 原数组中 front+i位置的元素
		newData[i] = q.data[(q.front+i)%len(q.data)]
	}
	q.data = newData
	q.front = 0
	q.tail = q.size
}

func (q *LoopQueue) Enqueue(e interface{}) {
	// 判断队列是否已满
	if (q.tail+1)%len(q.data) == q.front {
		q.resize(q.Cap() * 2)
	}
	q.data[q.tail] = e
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
}

func (q *LoopQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		log.Fatal("Dequeue failed, queue is empty")
	}
	ret := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % len(q.data)
	q.size--
	// resize
	if q.size == q.Cap()/4 && q.Cap()/2 != 0 {
		q.resize(q.Cap() / 2)
	}
	return ret
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
