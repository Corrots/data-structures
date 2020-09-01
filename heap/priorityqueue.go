package heap

// 基于最大堆实现的优先队列
type PriorityQueue struct {
	maxHeap *MaxHeap
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{maxHeap: NewMaxHeap()}
}

func (q *PriorityQueue) Enqueue(e interface{}) {
	q.maxHeap.Add(e.(int))
}

func (q *PriorityQueue) Dequeue() interface{} {
	return q.maxHeap.ExtractMax()
}

func (q *PriorityQueue) GetFront() interface{} {
	return q.maxHeap.FindMax()
}

func (q *PriorityQueue) Len() int {
	return q.maxHeap.Len()
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.maxHeap.IsEmpty()
}
