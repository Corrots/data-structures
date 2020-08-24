package main

import (
	"fmt"

	"github.com/corrots/data-structures/linkedlist/linkedlist/dlinkedlist"
)

func main() {
	ll := dlinkedlist.NewLinkedList()
	for i := 0; i < 5; i++ {
		ll.AddHead(i)
		fmt.Println(ll)
	}

	for i := 5; i < 10; i++ {
		ll.AddTail(i)
		fmt.Println(ll)
	}
	ll.Add(2, 666)
	fmt.Println(ll)
	ll.Add(10, 999)
	fmt.Println(ll)
	//
	//fmt.Println("del: ", ll.Remove(2))
	//fmt.Println(ll)
	//ll.RemoveFirst()
	//fmt.Println(ll)
	//ll.RemoveLast()
	//fmt.Println(ll)
}
