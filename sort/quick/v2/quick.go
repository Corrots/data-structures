package v2

import (
	"math/rand"
	"time"

	"github.com/corrots/data-structures/sort/helper"
)

func QuickSort(nums []int) {
	rand.Seed(time.Now().UnixNano())
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l, r int) {
	if r-l <= 8 {
		helper.InsertionSort(nums, l, r)
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func randomPivot(nums []int, l, r int) {
	randIndex := rand.Int()%(r-l+1) + l
	nums[randIndex], nums[l] = nums[l], nums[randIndex]
}

func partition(nums []int, l, r int) int {
	randomPivot(nums, l, r)
	v := nums[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if nums[i] < v {
			nums[i], nums[j+1] = nums[j+1], nums[i]
			j++
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}
