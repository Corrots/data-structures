package main

import (
	"fmt"

	linkedlist "github.com/corrots/data-structures/linkedlist/leetcode/707"
)

func main() {
	ll := linkedlist.Constructor()
	fmt.Println(ll)
	ll.AddAtHead(1)
	fmt.Println(ll)
	ll.AddAtTail(3)
	fmt.Println(ll)
	ll.AddAtIndex(1, 2)
	fmt.Println(ll)
	fmt.Println("get: ", ll.Get(1))
	ll.DeleteAtIndex(1)
	fmt.Println("get: ", ll.Get(1))
}
