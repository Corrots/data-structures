package linkedlist

import "log"

type Node struct {
	e    interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func newNode(e interface{}, next *Node) *Node {
	return &Node{
		e:    e,
		next: next,
	}
}

func New() *LinkedList {
	return &LinkedList{
		head: nil,
		size: 0,
	}
}

func (l *LinkedList) Len() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 在链表头添加新的元素e
func (l *LinkedList) AddFront(e interface{}) {
	l.head = newNode(e, l.head)
	l.size++
}

// 在链表的index位置添加新的元素e
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		log.Fatal("Add failed, illegal index")
	}
	if index == 0 {
		l.AddFront(e)
	} else {
		prev := l.head
		for i := 0; i < index-1; i++ {
			// 找到要插入位置的前一个节点
			prev = prev.next
		}
		//node := newNode(e, prev.next)
		prev.next = newNode(e, prev.next)
		l.size++
	}
}

// 在链表末尾添加新元素e
func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}
