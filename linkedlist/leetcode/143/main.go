package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	l1 := CreateLinkedList(nums, len(nums))
	PrintLinkedList(l1)
	reorderList(l1)
}

//https://leetcode-cn.com/problems/reorder-list/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var list []*ListNode
	cur := head
	for cur != nil {
		list = append(list, cur)
		cur = cur.Next
	}
	p, q := 0, len(list)-1
	for p < q {
		list[p].Next = list[q]
		list[q].Next = list[p+1]
		p++
		q--
	}
	list[p].Next = nil
	//PrintLinkedList(head)
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
