package main

import "fmt"

func main() {
	nums := []int{4, 2, 1, 3}
	l1 := CreateLinkedList(nums, len(nums))
	PrintLinkedList(l1)
	list := insertionSortList(l1)
	PrintLinkedList(list)
}

// https://leetcode-cn.com/problems/insertion-sort-list/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	tail, sort := head, head.Next
	for sort != nil {
		if sort.Val < tail.Val {
			pos := dummy
			for pos.Next.Val < sort.Val {
				pos = pos.Next
			}
			tail.Next = sort.Next
			sort.Next = pos.Next
			pos.Next = sort
			sort = tail.Next
		} else {
			tail = tail.Next
			sort = sort.Next
		}
	}
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
