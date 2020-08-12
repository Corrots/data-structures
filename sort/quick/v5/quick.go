package v5

import (
	"math/rand"
	"time"

	"github.com/corrots/data-structures/sort/helper"
)

func QuickSort(nums []int) {
	rand.Seed(time.Now().UnixNano())
	quickSort3Ways(nums, 0, len(nums)-1)
}

// 三路快排
func quickSort3Ways(nums []int, l, r int) {
	if r-l <= 8 {
		helper.InsertionSort(nums, l, r)
		return
	}
	//if l >= r {
	//	return
	//}
	lt, gt := partition(nums, l, r)
	quickSort3Ways(nums, l, lt-1)
	quickSort3Ways(nums, gt, r)
}

func randomPivot(nums []int, l, r int) {
	randIndex := rand.Int()%(r-l+1) + l
	nums[randIndex], nums[l] = nums[l], nums[randIndex]
}

func partition(nums []int, l, r int) (int, int) {
	randomPivot(nums, l, r)
	v := nums[l]
	// 循环不变量: nums[l+1:lt] < v; nums[lt+1:i-1] == v; nums[gt:r] > v
	lt, i, gt := l, l+1, r+1
	for i < gt {
		if nums[i] == v {
			i++
		} else if nums[i] < v {
			lt++
			nums[lt], nums[i] = nums[i], nums[lt]
			i++
		} else {
			gt--
			nums[gt], nums[i] = nums[i], nums[gt]
		}
	}
	nums[l], nums[lt] = nums[lt], nums[l]
	// 此时 nums[l, lt-1] < v; nums[lt, gt-1] == v; nums[gt, r] > v
	return lt, gt
}
