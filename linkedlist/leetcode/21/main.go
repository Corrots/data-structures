package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 4}
	nums2 := []int{1, 3, 4}
	l1 := CreateLinkedList(nums1, len(nums1))
	l2 := CreateLinkedList(nums2, len(nums2))
	head := mergeTwoLists(l1, l2)
	PrintLinkedList(head)
}

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummyHead := &ListNode{}
	p := l1
	q := l2
	cur := dummyHead
	for p != nil || q != nil {
		if p == nil {
			cur.Next = q
			q = q.Next
		} else if q == nil {
			cur.Next = p
			p = p.Next
		} else if p.Val <= q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}
	return dummyHead.Next
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
