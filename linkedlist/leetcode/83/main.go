package main

import "fmt"

func main() {
	nums := []int{1, 1, 1}
	head := CreateLinkedList(nums, len(nums))
	PrintLinkedList(head)
	mod1 := deleteDuplicates(head)
	PrintLinkedList(mod1)
	//nums1 := []int{1, 1, 2, 3, 3}

}

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			delNode := cur.Next
			cur.Next = delNode.Next
			delNode.Next = nil
		} else {
			cur = cur.Next
		}
	}
	return head
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
