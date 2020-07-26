package deque

type Deque interface {
	AddFront(interface{})
	AddLast(interface{})
	RemoveFront() interface{}
	RemoveLast() interface{}
	IsEmpty() bool
}
