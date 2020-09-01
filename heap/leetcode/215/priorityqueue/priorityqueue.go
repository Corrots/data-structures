package priorityqueue

import "github.com/corrots/data-structures/heap"

type PriorityQueue struct {
	minHeap *heap.MinHeap
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{minHeap: heap.NewMinHeap()}
}

func (q *PriorityQueue) Enqueue(e interface{}) {
	q.minHeap.Add(e.(int))
}

func (q *PriorityQueue) Dequeue() interface{} {
	return q.minHeap.ExtractMin()
}

func (q *PriorityQueue) GetFront() interface{} {
	return q.minHeap.FindMin()
}

func (q *PriorityQueue) Len() int {
	return q.minHeap.Len()
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.minHeap.IsEmpty()
}
