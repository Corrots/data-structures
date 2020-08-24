package dlinkedlist

import (
	"fmt"
	"log"
	"strings"
)

type ListNode struct {
	val  interface{}
	prev *ListNode
	next *ListNode
}

func newListNode(value interface{}, prev, next *ListNode) *ListNode {
	return &ListNode{val: value, prev: prev, next: next}
}

// 双链表
type LinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

func NewLinkedList() *LinkedList {
	head := newListNode(nil, nil, nil)
	tail := newListNode(nil, nil, nil)
	head.next = tail
	tail.prev = head

	return &LinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

// 在链表的index位置添加新的元素e
func (l *LinkedList) Add(index int, e interface{}) {
	l.checkIndex(index)
	var pred, succ *ListNode
	if index+1 < l.size-index {
		pred = l.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		succ = pred.next
	} else {
		succ = l.tail
		for i := 0; i < l.size-index; i++ {
			succ = succ.prev
		}
		pred = succ.prev
	}
	toAdd := newListNode(e, pred, succ)
	pred.next = toAdd
	succ.prev = toAdd
	l.size++
}

// 在链表头添加新的元素e
func (l *LinkedList) AddHead(e interface{}) {
	pred, succ := l.head, l.head.next
	toAdd := newListNode(e, pred, succ)
	pred.next = toAdd
	succ.prev = toAdd
	l.size++
}

// 在链表末尾添加新元素e
func (l *LinkedList) AddTail(e interface{}) {
	succ, pred := l.tail, l.tail.prev
	toAdd := newListNode(e, pred, succ)
	pred.next = toAdd
	succ.prev = toAdd
	l.size++
}

// 获取链表中index位置的元素
func (l *LinkedList) Get(index int) interface{} {
	l.checkIndex(index)
	// 选择从head还是tail进行遍历
	if index+1 < l.size-index {
		cur := l.head.next
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		return cur.val
	}
	cur := l.tail.prev
	for i := 0; i < l.size-index; i++ {
		cur = cur.prev
	}
	return cur.val
}

func (l *LinkedList) Set(index int, e interface{}) {
	l.checkIndex(index)
	if index+1 < l.size-index {
		cur := l.head.next
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		cur.val = e
	}
	cur := l.tail.prev
	for i := 0; i < l.size-index; i++ {
		cur = cur.prev
	}
	cur.val = e
}

func (l *LinkedList) Remove(index int) interface{} {
	l.checkIndex(index)
	var res interface{}
	if index+1 < l.size-index {
		prev := l.head
		for i := 0; i < index; i++ {
			prev = prev.next
		}
		delNode := prev.next
		res = delNode.val
		prev.next = delNode.next
		delNode.next, delNode.prev = nil, nil
	} else {
		pred := l.tail
		for i := 0; i < l.size-index-1; i++ {
			pred = pred.prev
		}
		delNode := pred.prev
		res = delNode.val
		pred.prev = delNode.prev
		delNode.prev, delNode.next = nil, nil
	}
	l.size--
	return res
}

func (l *LinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}

func (l *LinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) checkIndex(index int) {
	if index < 0 || index > l.size {
		log.Fatalf("invalid index %d\n", index)
	}
}

func (l *LinkedList) Len() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) String() string {
	var res strings.Builder
	cur := l.head.next
	for cur != nil {
		res.WriteString(fmt.Sprintf("%v->", cur.val))
		cur = cur.next
	}
	res.WriteString("nil")
	return res.String()
}
