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
	cur := dummyHead
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			cur.Next = l1
			l1 = l1.Next
		} else if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
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
