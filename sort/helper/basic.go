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
		for j = i; j > l && nums[j-1] > e; j-- {
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

func TestSort(nums []int, fn func([]int)) int64 {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())
	fn(nums)
	//fmt.Println(nums)
	if !IsSorted(nums) {
		log.Fatal("sorted failed")
	}
	return time.Since(start).Milliseconds()
}

func GenerateRandArray(count, rangeLeft, rangeRight int) []int {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		nums = append(nums, rand.Int()%(rangeRight-rangeLeft+1)+rangeLeft)
	}
	return nums
}

func GenerateNearlyOrderedArray(count, rangeLeft, rangeRight int) []int {
	nums := make([]int, count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		nums[i] = i
	}
	for i := 0; i < 10; i++ {
		randomIndex := rand.Intn(count)
		nums[i], nums[randomIndex] = nums[randomIndex], nums[i]
	}
	return nums
}

func GenerateOrderedArray(count int) []int {
	var nums []int
	for i := 0; i < count; i++ {
		nums = append(nums, i)
	}
	return nums
}
