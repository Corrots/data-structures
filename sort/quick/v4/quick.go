package v4

import (
	"math/rand"
	"time"

	"github.com/corrots/data-structures/sort/helper"
)

func QuickSort(nums []int) {
	rand.Seed(time.Now().UnixNano())
	quickSort2Ways(nums, 0, len(nums)-1)
}

// 双路快排
func quickSort2Ways(nums []int, l, r int) {
	if r-l <= 8 {
		helper.InsertionSort(nums, l, r)
		return
	}
	p := partition(nums, l, r)
	quickSort2Ways(nums, l, p-1)
	quickSort2Ways(nums, p+1, r)
}

func randomPivot(nums []int, l, r int) {
	randIndex := rand.Int()%(r-l+1) + l
	nums[randIndex], nums[l] = nums[l], nums[randIndex]
}

func partition(nums []int, l, r int) int {
	randomPivot(nums, l, r)
	v := nums[l]
	// 循环不变量: nums[l+1:i-1] <= v; nums[j+1:r] >= v
	i, j := l+1, r
	for {
		for i <= r && nums[i] < v {
			i++
		}
		for j >= l+1 && nums[j] > v {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[j], nums[l] = nums[l], nums[j]
	return j
}
