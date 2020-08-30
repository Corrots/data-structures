package main

import "fmt"

func main() {
	//nums := []int{1, 0, 1}
	nums := []int{1, 2, 3, 4}
	l1 := CreateLinkedList(nums, len(nums))
	fmt.Println(reversePrint(l1))
}

//https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/submissions/
func reversePrint(head *ListNode) []int {
	var res []int
	if head == nil {
		return res
	}
	rev := reverseList(head)
	for rev != nil {
		res = append(res, rev.Val)
		rev = rev.Next
	}
	return res
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
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
