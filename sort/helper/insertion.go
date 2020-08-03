package helper

import (
	"log"
	"math/rand"
	"time"
)

func InsertionSort(nums []int, l, r int) {
	for i := l + 1; i <= r; i++ {
		e := nums[i]
		j := 0
		for j := i; j > l && nums[j-1] > e; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = e
	}
}

func IsSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}

func TestSort(opCount int, fn func([]int)) int64 {
	var nums []int
	start := time.Now()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < opCount; i++ {
		nums = append(nums, rand.Int()%(opCount*10))
	}
	fn(nums)
	//fmt.Println(nums)
	if !IsSorted(nums) {
		log.Fatal("sorted failed")
	}
	return time.Since(start).Milliseconds()
}
