package main

import (
	"fmt"

	"github.com/corrots/data-structures/array/array"
)

func main() {
	arr := array.Init()
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)
	arr.Add(1, "abc")
	fmt.Println(arr)
	arr.AddFirst(-1)
	fmt.Println(arr)
	fmt.Println(arr.Remove(0))
	fmt.Println(arr.Remove(1))
	fmt.Println(arr)
}
