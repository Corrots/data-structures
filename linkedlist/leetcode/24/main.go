package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	l1 := CreateLinkedList(nums, len(nums))
	PrintLinkedList(l1)
	head := swapPairs(l1)
	PrintLinkedList(head)
}

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	p := dummyHead
	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := node1.Next
		next := node2.Next
		node2.Next = node1
		node1.Next = next
		p.Next = node2
		p = node1
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
