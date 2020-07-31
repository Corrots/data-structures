package main

import (
	"fmt"

	"github.com/corrots/data-structures/linkedlist/linkedlist/linkedlistr"
)

func main() {
	ll := linkedlistr.New()
	for i := 0; i < 5; i++ {
		ll.AddFirst(i)
		fmt.Println(ll)
	}
	ll.Add(2, 666)
	fmt.Println(ll)

	fmt.Println("del: ", ll.Remove(2))
	fmt.Println(ll)
	ll.RemoveFirst()
	fmt.Println(ll)
	ll.RemoveLast()
	fmt.Println(ll)
}
