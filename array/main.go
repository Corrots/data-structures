package main

import (
	"fmt"

	"github.com/corrots/data-structures/array/array"
)

func main() {
	arr := array.New(20)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)
	arr.Add(1, 100)
	fmt.Println(arr)
	arr.AddFirst(-1)
	fmt.Println(arr)
	arr.Delete(0)
	fmt.Println(arr)
	arr.Delete(1)
	fmt.Println(arr)
}
