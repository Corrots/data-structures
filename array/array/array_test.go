package array

import "testing"

func TestArray_Add(t *testing.T) {
	arr := New(20)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	t.Log(arr)
	arr.Add(1, 100)
	t.Log(arr)
	arr.AddFirst(-1)
	t.Log(arr)
	t.Log(arr.Remove(0))
	t.Log(arr.Remove(1))
}
