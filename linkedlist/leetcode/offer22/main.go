package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	l1 := CreateLinkedList(nums, len(nums))
	PrintLinkedList(l1)
	head := getKthFromEnd(l1, 2)
	PrintLinkedList(head)
}

//https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	p, q := dummy, dummy
	for i := 0; i < k+1; i++ {
		q = q.Next
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}
	return p.Next
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
