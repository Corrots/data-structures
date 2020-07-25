package queue

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	GetFront() interface{}
	Len() int
	IsEmpty() bool
}
