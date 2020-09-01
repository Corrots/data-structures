package heapsort

import "github.com/corrots/data-structures/heap"

// 堆排序
func HeapSort(nums []int) {
	maxHeap := heap.NewMaxHeap()
	for _, v := range nums {
		maxHeap.Add(v)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] = maxHeap.ExtractMax().(int)
	}
}
