package linkedlistqueue

import "log"

type Node struct {
	e    interface{}
	next *Node
}

func newNode(e interface{}, next *Node) *Node {
	return &Node{e: e, next: next}
}

type LinkedListQueue struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (l *LinkedListQueue) Enqueue(e interface{}) {
	if l.tail == nil {
		l.tail = newNode(e, nil)
		l.head = l.tail
	} else {
		l.tail.next = newNode(e, nil)
		l.tail = l.tail.next
	}
	l.size++
}

func (l *LinkedListQueue) Dequeue() interface{} {
	if l.IsEmpty() {
		log.Fatal("queue is empty")
	}
	retNode := l.head
	l.head = l.head.next
	retNode.next = nil
	if l.head == nil {
		l.tail = nil
	}
	l.size--
	return retNode.e
}

func (l *LinkedListQueue) GetFront() interface{} {
	if l.IsEmpty() {
		log.Fatal("queue is empty")
	}
	return l.head.e
}

func (l *LinkedListQueue) Len() int {
	return l.size
}

func (l *LinkedListQueue) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedListQueue) String() string {

}
