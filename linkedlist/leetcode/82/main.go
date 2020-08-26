package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 3, 4, 4, 5}
	head := CreateLinkedList(nums1, len(nums1))
	PrintLinkedList(head)
	l1 := deleteDuplicates(head)
	PrintLinkedList(l1)
}

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	a := dummyHead
	b := head
	for b != nil && b.Next != nil {
		if a.Next.Val == b.Next.Val {
			for b != nil && b.Next != nil && a.Next.Val == b.Next.Val {
				b = b.Next
			}
			a.Next = b.Next
			b = b.Next
		} else {
			a = a.Next
			b = b.Next
		}
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
