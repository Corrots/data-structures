package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//nums := []int{1, 1, 1, 2, 2, 3}
	nums := []int{-1, -1}
	k := 1
	fmt.Println(topKFrequent(nums, k))
}

//https://leetcode-cn.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	// 统计每个元素出现的频率
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	minHeap := &FeqMinHeap{}
	// 将hashMap中的所有数据对放入堆中
	for k, v := range freq {
		*minHeap = append(*minHeap, Feq{
			val:   k,
			count: v,
		})
	}
	heap.Init(minHeap)
	res := make([]int, k)
	// 从堆中取出k个元素即为频率最高的k个元素
	for i := 0; i < k && minHeap.Len() > 0; i++ {
		res[i] = heap.Pop(minHeap).(Feq).val
	}
	return res
}

type Feq struct {
	val   int
	count int
}

// 基于最大堆实现的优先队列
type FeqMinHeap []Feq

func (pq FeqMinHeap) Len() int {
	return len(pq)
}

func (pq FeqMinHeap) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq FeqMinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *FeqMinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(Feq))
}

func (pq *FeqMinHeap) Pop() interface{} {
	n := len(*pq) - 1
	res := (*pq)[n]
	*pq = (*pq)[:n]
	return res
}
