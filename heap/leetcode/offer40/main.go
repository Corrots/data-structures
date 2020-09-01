package main

import (
	"fmt"

	"github.com/corrots/data-structures/heap"
)

func main() {
	nums1 := []int{3, 2, 1}
	fmt.Println(getLeastNumbers(nums1, 2))
}

//https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func getLeastNumbers(arr []int, k int) []int {
	var res []int
	if len(arr) == 0 {
		return res
	}
	pq := heap.NewPriorityQueue()
	// 将前k个元素加入优先队列
	for i := 0; i < k; i++ {
		pq.Enqueue(arr[i])
	}
	// 遍历arr[k+1,n)，比较每个元素与队首元素的大小
	// 若arr[i]<队首元素，则交换之
	for i := k; i < len(arr); i++ {
		if !pq.IsEmpty() && arr[i] < pq.GetFront().(int) {
			pq.Dequeue()
			pq.Enqueue(arr[i])
		}
	}
	for i := 0; i < k; i++ {
		res = append(res, pq.Dequeue().(int))
	}
	return res
}
