package linkedlistr

import (
	"fmt"
	"strings"
)

//
type Node struct {
	val  interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func newNode(e interface{}, next *Node) *Node {
	return &Node{e, next}
}

func New() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

// 在链表的index位置添加新的元素e
func (l *LinkedList) Add(index int, e interface{}) {
	// @TODO check if 0 < index < l.size
	l.head = l.add(l.head, index, e)
	l.size++
}

func (l *LinkedList) add(node *Node, index int, e interface{}) *Node {
	if index == 0 {
		return newNode(e, node)
	}
	node.next = l.add(node.next, index-1, e)
	return node
}

func (l *LinkedList) Get(index int) interface{} {
	// @TODO check if 0 < index < l.size
	return l.get(l.head, index)
}

func (l *LinkedList) get(node *Node, index int) interface{} {
	if index == 0 {
		return node.val
	}
	return l.get(node.next, index-1)
}

func (l *LinkedList) Set(index int, e interface{}) {
	// @TODO check if 0 < index < l.size
	l.set(l.head, index, e)
}

func (l *LinkedList) set(node *Node, index int, e interface{}) {
	if index == 0 {
		node.val = e
	}
	l.set(node.next, index-1, e)
}

func (l *LinkedList) Remove(index int) (e interface{}) {
	// @TODO check if 0 < index < l.size
	l.head, e = l.remove(l.head, index)
	l.size--
	return e
}

func (l *LinkedList) remove(node *Node, index int) (*Node, interface{}) {
	if index == 0 {
		return node.next, node.val
	}
	var e interface{}
	node.next, e = l.remove(node.next, index-1)
	return node, e
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
		res.WriteString(fmt.Sprintf("%v->", cur.val))
		cur = cur.next
	}
	res.WriteString("nil")
	return res.String()
}
