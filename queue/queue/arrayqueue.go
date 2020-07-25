package queue

import (
	"fmt"
	"strings"
)

type ArrayQueue struct {
	array *Array
}

func NewArrayQueue() Queue {
	return &ArrayQueue{array: NewArray(10)}
}

func (q *ArrayQueue) Enqueue(e interface{}) {
	q.array.AddLast(e)
}

func (q *ArrayQueue) Dequeue() interface{} {
	return q.array.RemoveFirst()
}

func (q *ArrayQueue) GetFront() interface{} {
	return q.array.GetFirst()
}

func (q *ArrayQueue) String() string {
	var res strings.Builder
	res.WriteString(fmt.Sprintf("Queue: "))
	res.WriteString("front [")
	for i := 0; i < q.array.size; i++ {
		res.WriteString(fmt.Sprintf("%v", q.array.Get(i)))
		if i != q.array.size-1 {
			res.WriteString(", ")
		}
	}
	res.WriteString("] tail")
	return res.String()
}

func (q *ArrayQueue) Len() int {
	return q.array.Len()
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.array.IsEmpty()
}

func (q *ArrayQueue) Cap() int {
	return q.array.Cap()
}
