package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := CreateLinkedList(nums, len(nums))
	PrintLinkedList(head)
	o := oddEvenList(head)
	PrintLinkedList(o)
}

// https://leetcode-cn.com/problems/odd-even-linked-list/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	oddHead, evenHead := &ListNode{}, &ListNode{}
	odd := oddHead
	even := evenHead
	isOdd := true
	for head != nil {
		if isOdd {
			odd.Next = head
			odd = odd.Next
			isOdd = false
		} else {
			even.Next = head
			even = even.Next
			isOdd = true
		}
		head = head.Next
	}
	even.Next = nil
	odd.Next = evenHead.Next
	return oddHead.Next
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
