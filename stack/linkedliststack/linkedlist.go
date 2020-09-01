package linkedliststack

import (
	"fmt"
	"log"
	"strings"
)

type Node struct {
	e    interface{}
	next *Node
}

type LinkedList struct {
	dummyHead *Node
	size      int
}

func newNode(e interface{}, next *Node) *Node {
	return &Node{e: e, next: next}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{dummyHead: newNode(nil, nil), size: 0}
}

func (l *LinkedList) Len() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 在链表的index位置添加新的元素e
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		log.Fatal("Add failed, illegal index")
	}
	// 从链表的虚拟头节点开始遍历
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		// 找到要插入位置的前一个节点
		prev = prev.next
	}
	//node := newNode(e, prev.next)
	prev.next = newNode(e, prev.next)
	l.size++
}

// 在链表头添加新的元素e
func (l *LinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

// 在链表末尾添加新元素e
func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

// 获取链表中index位置的元素
func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.size {
		log.Fatal("Get failed, illegal index")
	}
	// 从链表的第一个真实元素开始遍历
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

func (l *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= l.size {
		log.Fatal("Set failed, illegal index")
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

func (l *LinkedList) Contains(e interface{}) bool {
	cur := l.dummyHead.next
	for cur != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= l.size {
		log.Fatal("ExtractMax failed, illegal index")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	delNode := prev.next
	ret := delNode.e
	prev.next = delNode.next
	delNode = nil
	l.size--
	return ret
}

func (l *LinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}

func (l *LinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) String() string {
	var res strings.Builder
	cur := l.dummyHead.next
	for cur != nil {
		res.WriteString(fmt.Sprintf("%v->", cur.e))
		cur = cur.next
	}
	res.WriteString("nil")
	return res.String()
}
