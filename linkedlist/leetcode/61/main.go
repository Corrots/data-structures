package main

import "fmt"

func main() {
	//nums := []int{1, 2, 3, 4, 5}
	nums := []int{0, 1, 2}
	l1 := CreateLinkedList(nums, len(nums))
	PrintLinkedList(l1)
	list := rotateRight(l1, 4)
	PrintLinkedList(list)
}

//https://leetcode-cn.com/problems/rotate-list/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oldTail := head
	var n int
	for n = 1; oldTail.Next != nil; n++ {
		oldTail = oldTail.Next
	}
	oldTail.Next = head
	tail := head
	for i := 1; i < n-k%n; i++ {
		tail = tail.Next
	}
	newHead := tail.Next
	tail.Next = nil
	return newHead
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
