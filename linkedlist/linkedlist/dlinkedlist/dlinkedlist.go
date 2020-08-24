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
	// 判断从head还是tail进行遍历更优
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

// 向链表头添加新的元素e
func (l *LinkedList) AddAtHead(e interface{}) {
	pred, succ := l.head, l.head.next
	toAdd := newListNode(e, pred, succ)
	pred.next = toAdd
	succ.prev = toAdd
	l.size++
}

// 向链表尾添加新元素e
func (l *LinkedList) AddAtTail(e interface{}) {
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
	var cur *ListNode
	if index+1 < l.size-index {
		cur = l.head
		// 这里是直接获取index位置的节点,故需index+1
		for i := 0; i < index+1; i++ {
			cur = cur.next
		}
	} else {
		cur = l.tail
		for i := 0; i < l.size-index; i++ {
			cur = cur.prev
		}
	}
	return cur.val
}

func (l *LinkedList) Set(index int, e interface{}) {
	l.checkIndex(index)
	var cur *ListNode
	if index+1 < l.size-index {
		cur = l.head
		for i := 0; i < index+1; i++ {
			cur = cur.next
		}
	} else {
		cur = l.tail
		for i := 0; i < l.size-index; i++ {
			cur = cur.prev
		}
	}
	cur.val = e
}

func (l *LinkedList) Remove(index int) interface{} {
	l.checkIndex(index)
	var pred, succ, delNode *ListNode
	if index < l.size-index {
		pred = l.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		delNode = pred.next
		succ = pred.next.next
	} else {
		succ = l.tail
		for i := 0; i < l.size-index-1; i++ {
			succ = succ.prev
		}
		delNode = succ.prev
		pred = succ.prev.prev
	}
	pred.next = succ
	succ.prev = pred
	l.size--
	return delNode.val
}

// 删除头节点
//func (l *LinkedList) RemoveFirst() interface{} {
//	delNode := l.head
//	l.head = delNode.next
//	l.head.next.prev = l.head
//	l.size--
//	return delNode.val
//}

// 删除尾节点
//func (l *LinkedList) RemoveLast() interface{} {
//
//}

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
