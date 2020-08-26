package main

import "fmt"

func main() {
	nums1 := []int{2, 4, 3}
	nums2 := []int{5, 6, 4}
	l1 := CreateLinkedList(nums1, len(nums1))
	l2 := CreateLinkedList(nums2, len(nums2))
	l3 := addTwoNumbers(l1, l2)
	PrintLinkedList(l3)
}

// https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p, q, cur := l1, l2, dummyHead
	var carry int
	for p != nil || q != nil {
		var x, y int
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum := carry + x + y
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
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
