package linkedlistqueue

import (
	"fmt"
	"log"
	"strings"
)

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

// 入队
func (l *LinkedListQueue) Enqueue(e interface{}) {
	// tail为nil，则表示head也为nil
	if l.tail == nil {
		l.tail = newNode(e, nil)
		l.head = l.tail
	} else {
		l.tail.next = newNode(e, nil)
		// @TODO
		l.tail = l.tail.next
	}
	l.size++
}

// 出队
func (l *LinkedListQueue) Dequeue() interface{} {
	if l.IsEmpty() {
		log.Fatal("queue is empty")
	}
	retNode := l.head
	l.head = l.head.next
	retNode.next = nil
	// 链表出队后变为空
	if l.head == nil {
		// 出队前的tail指向被出队的元素
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
	var res strings.Builder
	res.WriteString("Queue: front ")
	cur := l.head
	for cur != nil {
		res.WriteString(fmt.Sprintf("%v->", cur.e))
		cur = cur.next
	}
	res.WriteString("nil tail")
	return res.String()
}
