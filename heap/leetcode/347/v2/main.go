package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//nums := []int{1, 1, 1, 2, 2, 3}
	//nums := []int{-1, -1}
	nums := []int{5, 3, 1, 1, 1, 3, 73, 1}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

//https://leetcode-cn.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	// 统计所有元素的频率
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	//fmt.Println(freq)
	minHeap := &MinHeap{}
	for key, v := range freq {
		if minHeap.Len() == k {
			if v > minHeap.Peek().(*Feq).count {
				heap.Pop(minHeap)
				heap.Push(minHeap, &Feq{
					val:   key,
					count: v,
				})
			}
		} else {
			heap.Push(minHeap, &Feq{
				val:   key,
				count: v,
			})
		}
	}
	var res []int
	for minHeap.Len() > 0 {
		res = append(res, heap.Pop(minHeap).(*Feq).val)
	}
	return res
}

type Feq struct {
	val   int
	count int
}

type MinHeap []*Feq

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(e interface{}) {
	*h = append(*h, e.(*Feq))
}

func (h *MinHeap) Pop() interface{} {
	n := len(*h) - 1
	res := (*h)[n]
	*h = (*h)[:n]
	return res
}

func (h *MinHeap) Peek() interface{} {
	return (*h)[0]
}
