package linkedlistr

import (
	"fmt"
	"log"
	"strings"
)

type listNode struct {
	e    interface{}
	next *listNode
}

// 链表的递归实现
type LinkedList struct {
	head *listNode
	size int
}

func newNode(e interface{}, next *listNode) *listNode {
	return &listNode{e: e, next: next}
}

func New() *LinkedList {
	return &LinkedList{}
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
	l.head = l.add(l.head, index, e)
	l.size++
}

// 在以node为头节点的index位置插入元素e
func (l *LinkedList) add(node *listNode, index int, e interface{}) *listNode {
	if index == 0 {
		return newNode(e, node)
	}
	node.next = l.add(node.next, index-1, e)
	return node
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
	return l.get(l.head, index)
}

func (l *LinkedList) get(node *listNode, index int) interface{} {
	if index == 0 {
		return node.e
	}
	return l.get(node.next, index-1)
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
	l.set(l.head, index, e)
}

func (l *LinkedList) set(node *listNode, index int, e interface{}) {
	if index == 0 {
		node.e = e
		return
	}
	l.set(node.next, index-1, e)
}

func (l *LinkedList) Contains(e interface{}) bool {
	return l.contains(l.head, e)
}

func (l *LinkedList) contains(node *listNode, e interface{}) bool {
	if node == nil {
		return false
	}
	if node.e == e {
		return true
	}
	return l.contains(node.next, e)
}

func (l *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= l.size {
		log.Fatal("Remove failed, illegal index")
	}
	N, E := l.remove(l.head, index)
	l.head = N
	l.size--
	return E
}

func (l *LinkedList) remove(node *listNode, index int) (*listNode, interface{}) {
	if index == 0 {
		return node.next, node.e
	}
	N, E := l.remove(node.next, index-1)
	node.next = N
	return node, E

}

func (l *LinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}

func (l *LinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) String() string {
	var res strings.Builder
	cur := l.head
	for cur != nil {
		res.WriteString(fmt.Sprintf("%v->", cur.e))
		cur = cur.next
	}
	res.WriteString("nil")
	return res.String()
}
