package main

import "fmt"

func main() {
	//nums := []int{1, 2, 3, 4, 5}
	nums := []int{1, 2}
	l1 := CreateLinkedList(nums, len(nums))
	PrintLinkedList(l1)
	//list := removeNthFromEnd(l1, 2)
	list := removeNthFromEnd(l1, 2)
	PrintLinkedList(list)
}

//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	var length int
	for cur != nil {
		length++
		cur = cur.Next
	}
	//给定的 n 保证是有效的。
	if length == n {
		return head.Next
	}
	dummy := &ListNode{Next: head}
	prev := dummy.Next
	for i := 1; i < length-n; i++ {
		prev = prev.Next
	}
	delNode := prev.Next
	prev.Next = delNode.Next
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
