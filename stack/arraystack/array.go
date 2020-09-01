package arraystack

import (
	"fmt"
	"log"
	"strings"
)

type Array struct {
	data []interface{}
	size int
}

func NewArray(cap int) *Array {
	return &Array{
		data: make([]interface{}, cap),
		size: 0,
	}
}

func Init() *Array {
	return NewArray(10)
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

func (arr *Array) AddLast(e interface{}) {
	arr.Add(arr.size, e)
}

func (arr *Array) AddFirst(e interface{}) {
	arr.Add(0, e)
}

// 向数组中指定位置i添加元素e
func (arr *Array) Add(i int, e interface{}) {
	if i < 0 || i > arr.size {
		log.Fatal("Add failed, required i >= 0 and <= size")
	}
	if arr.size == len(arr.data) {
		arr.resize(2 * len(arr.data))
	}
	for j := arr.size - 1; j >= i; j-- {
		arr.data[j+1] = arr.data[j]
	}
	arr.data[i] = e
	arr.size++
}

func (arr *Array) Get(i int) interface{} {
	if i < 0 || i > arr.size {
		log.Fatal("Add failed, required i >= 0 and <= size")
	}
	return arr.data[i]
}

func (arr *Array) GetLast() interface{} {
	return arr.Get(arr.size - 1)
}

func (arr *Array) GetFirst() interface{} {
	return arr.Get(0)
}

func (arr *Array) Set(i int, e interface{}) {
	if i < 0 || i > arr.size {
		log.Fatal("Add failed, required i >= 0 and <= size")
	}
	arr.data[i] = e
}

func (arr *Array) Contains(e interface{}) bool {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return true
		}
	}
	return false
}

func (arr *Array) Find(e interface{}) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return i
		}
	}
	return -1
}

func (arr *Array) Remove(i int) interface{} {
	if i < 0 || i > arr.size {
		log.Fatal("ExtractMax failed, required i >= 0 and <= size")
	}
	ret := arr.data[i]
	for j := i + 1; j < arr.size; j++ {
		arr.data[j-1] = arr.data[j]
	}
	arr.size--
	arr.data[arr.size] = nil // loitering objects != memory leak
	if arr.size == len(arr.data)/4 && len(arr.data)/2 != 0 {
		arr.resize(len(arr.data) / 2)
	}
	return ret
}

func (arr *Array) RemoveFirst() interface{} {
	return arr.Remove(0)
}

func (arr *Array) RemoveLast() interface{} {
	return arr.Remove(arr.size - 1)
}

func (arr *Array) RemoveElement(e interface{}) {
	i := arr.Find(e)
	if i != -1 {
		arr.Remove(i)
	}
}

func (arr *Array) String() string {
	var res strings.Builder
	res.WriteString(fmt.Sprintf("Array: size: %d, capacity: %d\n", arr.size, len(arr.data)))
	res.WriteString("[")
	for i := 0; i < arr.size; i++ {
		res.WriteString(fmt.Sprintf("%v", arr.data[i]))
		if i != arr.size-1 {
			res.WriteString(", ")
		}
	}
	res.WriteString("]")
	return res.String()
}

func (arr *Array) resize(n int) {
	newData := make([]interface{}, n)
	copy(newData, arr.data)
	//for i := 0; i < arr.size; i++ {
	//	newData[i] = arr.data[i]
	//}
	arr.data = newData
	//fmt.Println("arr.data: ", arr.data)
}
