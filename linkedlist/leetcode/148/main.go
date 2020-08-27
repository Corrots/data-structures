package main

import "fmt"

func main() {
	nums := []int{4, 2, 1, 3}
	l1 := CreateLinkedList(nums, len(nums))
	PrintLinkedList(l1)
	list := sortList(l1)
	PrintLinkedList(list)
}

//https://leetcode-cn.com/problems/sort-list/solution/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	var length int
	p := head
	for p != nil {
		length += 1
		p = p.Next
	}

	for sz := 1; sz < length; sz += sz {
		cur := dummy.Next
		tail := dummy
		for cur != nil {
			left := cur
			right := cut(left, sz)
			//fmt.Printf("left: %d, right: %d\n", left.Val, right.Val)
			cur = cut(right, sz)

			tail.Next = merge(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return dummy.Next
}

//将链表切掉前 n 个节点，并返回后半部分的链表头
func cut(head *ListNode, n int) *ListNode {
	var newRoot *ListNode
	cur := head
	for i := 0; i < n && cur != nil; i++ {
		if i == n-1 {
			newRoot = cur.Next
			cur.Next = nil
		} else {
			cur = cur.Next
		}
	}
	return newRoot
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			cur.Next = l1
			l1 = l1.Next
		} else if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func CreateLinkedList(nums []int, n int) *ListNode {
	if n == 0 {
		return nil
	}
	head := NewListNode(nums[0])
	cur := head
	for i := 1; i < n; i++ {
		cur.Next = NewListNode(nums[i])
		cur = cur.Next
	}
	return head
}

func PrintLinkedList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("nil\n")
}
