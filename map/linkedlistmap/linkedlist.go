package linkedlistmap

import (
	"fmt"
	"log"
	"strings"
)

type ListNode struct {
	key  *int
	val  interface{}
	next *ListNode
}

type LinkedList struct {
	dummyHead *ListNode
	size      int
}

func newNode(k *int, v interface{}, next *ListNode) *ListNode {
	return &ListNode{key: k, val: v, next: next}
}

func (l *LinkedList) getNode(k *int) *ListNode {
	cur := l.dummyHead
	for cur != nil {
		if cur.key == k {
			return cur
		}
		cur = cur.next
	}
	return nil
}

func NewLinkedList() *LinkedList {
	return &LinkedList{dummyHead: newNode(nil, nil, nil), size: 0}
}

func (l *LinkedList) Len() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 在链表的index位置添加新的元素e
func (l *LinkedList) Add(k int, v interface{}) {
	node := l.getNode(&k)
	if node == nil {
		l.dummyHead.next = newNode(&k, v, l.dummyHead.next)
		l.size++
		return
	}
	node.val = v
}

// 获取链表中index位置的元素
func (l *LinkedList) Get(index int) interface{} {
	node := l.getNode(&index)
	if node == nil {
		return nil
	}
	return node.val
}

//
//func (l *LinkedList) GetFirst() interface{} {
//	return l.Get(0)
//}
//
//func (l *LinkedList) GetLast() interface{} {
//	return l.Get(l.size - 1)
//}

func (l *LinkedList) Set(index int, e interface{}) {
	node := l.getNode(&index)
	if node == nil {
		log.Fatalf("%d doesn't exist\n", index)
	}
	node.val = e
}

func (l *LinkedList) Contains(k int) bool {
	return l.getNode(&k) != nil
}

func (l *LinkedList) Remove(index int) interface{} {
	prev := l.dummyHead
	// 先找到要删除节点的prev
	for prev.next != nil {
		if prev.next.key == &index {
			break
		}
		prev = prev.next
	}
	if prev.next != nil {
		delNode := prev.next
		prev.next = delNode.next
		delNode.next = nil
		l.size--
		return delNode.val
	}
	return nil
}

func (l *LinkedList) String() string {
	var res strings.Builder
	cur := l.dummyHead.next
	for cur != nil {
		res.WriteString(fmt.Sprintf("%v: %v->", cur.key, cur.val))
		cur = cur.next
	}
	res.WriteString("nil")
	return res.String()
}
