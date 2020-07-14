package array

import (
	"errors"
	"fmt"
)

type Array struct {
	data []int
	size int
}

func New(cap int) *Array {
	if cap == 0 {
		cap = 10
	}
	return &Array{
		data: make([]int, cap),
		size: 0,
	}
}

func (arr *Array) String() string {
	fmt.Print(arr.data)
	res := fmt.Sprintf("Array: size: %d, capacity: %d\n", arr.size, len(arr.data))
	res += "["
	for i := 0; i < arr.size; i++ {
		res += string(arr.data[i])
		if i != arr.size-1 {
			res += " "
		}
	}
	res += "]"
	return res
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

func (arr *Array) AddLast(e int) error {
	return arr.Add(arr.size, e)
}

func (arr *Array) AddFirst(e int) error {
	return arr.Add(0, e)
}

func (arr *Array) Add(index, e int) error {
	if arr.size == len(arr.data) {
		return errors.New("Add failed, Array is full ")
	}
	if index < 0 || index > arr.size {
		return errors.New("Add failed, required index >= 0 and <= size ")
	}
	for i := arr.size - 1; i >= index; i-- {
		arr.data[i+1] = arr.data[i]
	}
	arr.data[index] = e
	arr.size++
	return nil
}
