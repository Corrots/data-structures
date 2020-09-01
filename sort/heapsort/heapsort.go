package heapsort

import "github.com/corrots/data-structures/heap"

// 堆排序
func HeapSort(nums []int) {
	maxHeap := heap.NewMaxHeap()
	for _, v := range nums {
		maxHeap.Add(v)
	}
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i] = maxHeap.ExtractMax()
	}
}

// 原地堆排序
func InPlaceHeapSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	// heapify构建最大堆
	lastNoneLeaf := (n - 2) / 2
	for i := lastNoneLeaf; i >= 0; i-- {
		siftDown(nums, i, n)
	}
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		siftDown(nums, 0, i)
	}
}

// 对以数组data[0,n)形成的最大堆中，对索引为k的元素做siftDown
func siftDown(data []int, k, n int) {
	// 循环中止条件:k的左孩子已越界
	for 2*k+1 < n {
		j := 2*k + 1
		if j+1 < n && data[j] < data[j+1] {
			j++
		}
		// 索引j记录左右孩子中的最大值
		if data[k] > data[j] {
			break
		}
		data[k], data[j] = data[j], data[k]
		k = j
	}
}
