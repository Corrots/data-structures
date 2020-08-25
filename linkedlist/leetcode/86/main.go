package main

import "fmt"

func main() {
	nums := []int{1, 4, 3, 2, 5, 2}
	head := CreateLinkedList(nums, len(nums))
	PrintLinkedList(head)
	p := partition(head, 3)
	PrintLinkedList(p)
}

// https://leetcode-cn.com/problems/partition-list/
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	beforeHead, afterHead := &ListNode{}, &ListNode{}
	before := beforeHead
	after := afterHead
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
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
