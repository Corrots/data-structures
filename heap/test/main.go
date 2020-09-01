package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/corrots/data-structures/heap"
	"github.com/corrots/data-structures/sort/heapsort"
)

func main() {
	//TestMaxHeap()
	TestHeapSort()
}

func TestHeapSort() {
	opCount := 1000000
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, opCount)
	for i := 0; i < opCount; i++ {
		nums[i] = rand.Intn(math.MaxInt64)
	}
	start := time.Now()
	heapsort.HeapSort(nums)
	since := time.Since(start).Milliseconds()
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			log.Fatalf("Sort error")
		}
	}
	fmt.Printf("spent %d ms\n", since)
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
		nums[i] = maxHeap.ExtractMax().(int)
	}
	since := time.Since(start).Milliseconds()
	SortedCheck(nums)
	fmt.Printf("spent %d ms\n", since)
}

func SortedCheck(nums []int) {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			log.Fatalf("Error ")
		}
	}
}
