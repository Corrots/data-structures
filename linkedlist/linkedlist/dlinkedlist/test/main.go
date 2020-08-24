package main

import (
	"fmt"

	"github.com/corrots/data-structures/linkedlist/linkedlist/dlinkedlist"
)

func main() {
	ll := dlinkedlist.NewLinkedList()
	for i := 0; i < 5; i++ {
		ll.AddAtHead(i)
		fmt.Println(ll)
	}

	for i := 5; i < 10; i++ {
		ll.AddAtTail(i)
		fmt.Println(ll)
	}
	ll.Add(2, 666)
	fmt.Println(ll)
	ll.Add(10, 999)
	fmt.Println(ll)
	// get
	fmt.Println(ll.Get(3))
	fmt.Println(ll.Get(10))
	// get
	ll.Set(2, 996)
	fmt.Println(ll)
	ll.Set(10, 669)
	fmt.Println(ll)
	// del
	fmt.Println("del: ", ll.Remove(10))
	fmt.Println(ll)
	fmt.Println("del: ", ll.Remove(2))
	fmt.Println(ll)
	//fmt.Println("del: ", ll.Remove(2))

	//ll.RemoveFirst()
	//ll.RemoveLast()
	//fmt.Println(ll)
}
