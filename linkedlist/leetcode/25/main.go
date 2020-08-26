package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	l1 := CreateLinkedList(nums, len(nums))
	PrintLinkedList(l1)
	head := reverseKGroup(l1, 2)
	PrintLinkedList(head)
}

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	prev := hair
	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next
		head, tail = reverse(head, tail)
		prev.Next = head
		tail.Next = next
		prev = tail
		head = tail.Next
	}
	return hair.Next
}

// 翻转部分链表
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	cur := head
	for prev != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return tail, head
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
