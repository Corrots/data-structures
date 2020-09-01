package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/corrots/data-structures/heap"
)

func main() {
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

	for i := 1; i < opCount; i++ {
		if nums[i-1] < nums[i] {
			log.Fatalf("Error ")
		}
	}
	fmt.Printf("spent %d ms\n", time.Since(start).Milliseconds())
}
