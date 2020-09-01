package main

import (
	"fmt"

	"github.com/corrots/data-structures/heap/leetcode/offer40/v2/priorityqueue"
)

func main() {
	nums1 := []int{3, 2, 1, 10, 9}
	fmt.Println(getLeastNumbers(nums1, 3))
}

// Top K, 求最大的k个元素
func getLeastNumbers(arr []int, k int) []int {
	var res []int
	if len(arr) == 0 {
		return res
	}
	pq := priorityqueue.NewPriorityQueue()
	// 将前K个元素加入优先队列
	for i := 0; i < k; i++ {
		pq.Enqueue(arr[i])
	}
	// 遍历arr[k:n)，若当前元素>优先队列中的最小值，则替换之
	for i := k; i < len(arr); i++ {
		if arr[i] > pq.GetFront().(int) {
			pq.Dequeue()
			pq.Enqueue(arr[i])
		}
	}
	for i := 0; i < k; i++ {
		res = append(res, pq.Dequeue().(int))
	}
	return res
}
