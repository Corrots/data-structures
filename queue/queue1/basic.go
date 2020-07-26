package queue1

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	GetFront() interface{}
	Len() int
	IsEmpty() bool
}
