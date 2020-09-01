package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/corrots/data-structures/sort/heapsort"

	"github.com/corrots/data-structures/heap"
)

func main() {
	//TestMaxHeap()
	//TestHeapSort()
	//TestHeapify()
	TestMinHeap()
}

func TestHeapSort() {
	opCount := 1000000
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, opCount)
	for i := 0; i < opCount; i++ {
		nums[i] = rand.Intn(opCount * 10)
	}
	start := time.Now()
	//heapsort.HeapSort(nums)
	heapsort.InPlaceHeapSort(nums)
	since := time.Since(start).Milliseconds()
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			log.Fatalf("Sort error")
		}
	}
	fmt.Printf("spent %d ms\n", since)
}

// 测试比较普通方式与Heapify方式来构建二叉堆的耗时
func TestHeapify() {
	opCount := 1000000
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, opCount)
	for i := 0; i < opCount; i++ {
		nums[i] = rand.Intn(math.MaxInt64)
	}
	nums2 := make([]int, opCount)
	copy(nums2, nums)
	start1 := time.Now()
	heap1 := heap.NewMaxHeap()
	for _, v := range nums {
		heap1.Add(v)
	}
	since := time.Since(start1).Milliseconds()
	if isMaxHeap(heap1, opCount) {
		fmt.Printf("max heap spent %d ms\n", since)
	}
	start2 := time.Now()
	heap2 := heap.HeapifyNew(nums2)
	since1 := time.Since(start2).Milliseconds()
	if isMaxHeap(heap2, opCount) {
		fmt.Printf("heapify spent %d ms\n", since1)
	}
}

func isMaxHeap(h *heap.MaxHeap, length int) bool {
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		nums[i] = h.ExtractMax()
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			return false
		}
	}
	return true
}

// 测试最小堆
func TestMinHeap() {
	count := 10
	minHeap := heap.NewMinHeap()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		minHeap.Add(rand.Intn(count * 10))
	}
	nums := make([]int, count)
	for i := 0; i < count; i++ {
		nums[i] = minHeap.ExtractMin()
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			log.Fatalf("Error ")
		}
	}
	fmt.Println(nums)
}

// 测试最大堆
func TestMaxHeap() {
	opCount := 1000000
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	maxHeap := heap.NewMaxHeap()
	for i := 0; i < opCount; i++ {
		maxHeap.Add(rand.Intn(math.MaxInt64))
	}

	nums := make([]int, opCount)
	for i := 0; i < opCount; i++ {
		nums[i] = maxHeap.ExtractMax()
	}
	SortedCheck(nums)
	since := time.Since(start).Milliseconds()
	fmt.Printf("spent %d ms\n", since)
}

func SortedCheck(nums []int) {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			log.Fatalf("Error ")
		}
	}
}
