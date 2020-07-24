package array

import (
	"fmt"
	"strconv"
	"strings"
)

type Array struct {
	data []int
	size int
}

func New(cap int) *Array {
	return &Array{
		data: make([]int, cap),
		size: 0,
	}
}

func Init() *Array {
	return New(10)
}

func (arr *Array) Len() int {
	return arr.size
}

func (arr *Array) Cap() int {
	return len(arr.data)
}

func (arr *Array) IsEmpty() bool {
	return arr.size == 0
}

func (arr *Array) AddLast(e int) {
	arr.Add(arr.size, e)
}

func (arr *Array) AddFirst(e int) {
	arr.Add(0, e)
}

// 向数组中指定位置i添加元素e
func (arr *Array) Add(i, e int) {
	if arr.size == len(arr.data) {
		panic("Add failed, Array is full")
	}
	if i < 0 || i > arr.size {
		panic("Add failed, required i >= 0 and <= size ")
	}
	for j := arr.size - 1; j >= i; j-- {
		arr.data[j+1] = arr.data[j]
	}
	arr.data[i] = e
	arr.size++
}

func (arr *Array) Get(i int) int {
	if i < 0 || i > arr.size {
		panic("Add failed, required i >= 0 and <= size ")
	}
	return arr.data[i]
}

func (arr *Array) Set(i, e int) {
	if i < 0 || i > arr.size {
		panic("Add failed, required i >= 0 and <= size ")
	}
	arr.data[i] = e
}

func (arr *Array) String() string {
	var res strings.Builder
	res.WriteString(fmt.Sprintf("Array: size: %d, capacity: %d\n", arr.size, len(arr.data)))
	res.WriteString("[")
	for i := 0; i < arr.size; i++ {
		res.WriteString(strconv.Itoa(arr.data[i]))
		if i != arr.size-1 {
			res.WriteString(", ")
		}
	}
	res.WriteString("]")
	return res.String()
}
