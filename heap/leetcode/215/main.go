package main

import (
	"fmt"

	"github.com/corrots/data-structures/heap/leetcode/215/priorityqueue"
)

func main() {
	nums1 := []int{3, 2, 1, 10, 9}
	fmt.Println(findKthLargest(nums1, 2))
}

// 数组中的第K个最大元素
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
func findKthLargest(arr []int, k int) int {
	if len(arr) == 0 {
		return -1
	}
	pq := priorityqueue.NewPriorityQueue()
	// 将前K个元素加入优先队列
	for i := 0; i < k; i++ {
		pq.Enqueue(arr[i])
	}
	// 遍历arr[k:n)，若当前元素>优先队列中的最小值，则替换之
	for i := k; i < len(arr); i++ {
		if !pq.IsEmpty() && arr[i] > pq.GetFront().(int) {
			pq.Dequeue()
			pq.Enqueue(arr[i])
		}
	}
	return pq.Dequeue().(int)
}
