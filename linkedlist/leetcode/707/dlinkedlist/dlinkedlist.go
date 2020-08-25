package dlinkedlist

type ListNode struct {
	val  int
	prev *ListNode
	next *ListNode
}

func newListNode(v int, prev, next *ListNode) *ListNode {
	return &ListNode{v, prev, next}
}

type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	// 初始化伪头，伪尾
	head, tail := &ListNode{}, &ListNode{}
	head.next = tail
	tail.prev = head
	return MyLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.size {
		return -1
	}
	var cur *ListNode
	if index <= this.size-index {
		cur = this.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
	} else {
		cur = this.tail
		for i := 0; i < this.size-index; i++ {
			cur = cur.prev
		}
	}
	return cur.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	// 前驱=伪头, 后继=伪头.next
	pred, succ := this.head, this.head.next
	toAdd := newListNode(val, pred, succ)
	pred.next = toAdd
	succ.prev = toAdd
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	succ, pred := this.tail, this.tail.prev
	toAdd := newListNode(val, pred, succ)
	pred.next = toAdd
	succ.prev = toAdd
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	var pred, succ *ListNode
	if index <= this.size-index {
		pred = this.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		succ = pred.next
	} else {
		succ = this.tail
		for i := 0; i < this.size-index; i++ {
			succ = succ.prev
		}
		pred = succ.prev
	}
	toAdd := newListNode(val, pred, succ)
	pred.next = toAdd
	succ.prev = toAdd
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.size {
		return
	}
	var pred, succ *ListNode
	if index <= this.size-index {
		pred = this.head
		for i := 0; i < index; i++ {
			pred = pred.next
		}
		succ = pred.next.next
	} else {
		succ = this.tail
		for i := 0; i < this.size-index; i++ {
			succ = succ.prev
		}
		pred = succ.prev.prev
	}
	pred.next = succ
	succ.prev = pred
	this.size--
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
