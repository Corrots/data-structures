package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	l1 := CreateLinkedList(nums, len(nums))
	PrintLinkedList(l1)
	list := removeNthFromEnd(l1, 2)
	PrintLinkedList(list)
}

//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	p, q := dummy, dummy
	for i := 0; i <= n; i++ {
		q = q.Next
	}
	//fmt.Printf("p: %d, q: %d\n", p.Val, q.Val)
	for q != nil {
		p = p.Next
		q = q.Next
	}
	delNode := p.Next
	p.Next = delNode.Next
	delNode = nil
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
