package main

import "fmt"

func main() {
	nums := []int{1, 0, 1}
	l1 := CreateLinkedList(nums, len(nums))
	PrintLinkedList(l1)
	fmt.Println(isPalindrome(l1))
}

//https://leetcode-cn.com/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	right, n := halfTail(head)
	//PrintLinkedList(right)
	prev := reverseList(right)
	for i := 0; i < n/2; i++ {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}
	return true
}

func halfTail(head *ListNode) (*ListNode, int) {
	var n int
	p := head
	for n = 1; p.Next != nil; n++ {
		p = p.Next
	}
	for i := 0; i < n/2; i++ {
		head = head.Next
	}
	return head, n
}

func reverseList(head *ListNode) *ListNode {
	prev := &ListNode{Next: head}
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
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
