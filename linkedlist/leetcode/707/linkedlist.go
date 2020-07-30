package linkedlist

import (
	"fmt"
	"strings"
)

type listNode struct {
	e    int
	next *listNode
}

func newNode(e int, next *listNode) *listNode {
	return &listNode{e: e, next: next}
}

type MyLinkedList struct {
	dummyHead *listNode
	size      int
}

/** Initialize your data structure here. */
func Constructor() *MyLinkedList {
	return &MyLinkedList{dummyHead: &listNode{}}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	prev := this.dummyHead
	for prev.next != nil {
		prev = prev.next
	}
	prev.next = newNode(val, prev.next)
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	prev := this.dummyHead
	// 找到待插入节点的前一个节点
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = newNode(val, prev.next)
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	prev := this.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	//delNode := prev.next
	prev.next = prev.next.next
	//delNode.next = nil
	this.size--
}

func (this *MyLinkedList) String() string {
	var res strings.Builder
	cur := this.dummyHead.next
	for cur != nil {
		res.WriteString(fmt.Sprintf("%v->", cur.e))
		cur = cur.next
	}
	res.WriteString("nil")
	return res.String()
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
