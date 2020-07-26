package queue2

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	GetFront() interface{}
	Len() int
	IsEmpty() bool
}
