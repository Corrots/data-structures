package main

import (
	"fmt"

	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	//nums := []int{3, 2, 5, 1, 4, 0}
	//MergeSort(nums)
	//fmt.Println(nums)
	opCount := 1000000
	t1 := helper.TestSort(opCount, MergeSort)
	fmt.Printf("MergeSort spend %d ms", t1)
}

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) {
	//if l >= r {
	//	return
	//}
	if r-l <= 15 {
		helper.InsertionSort(nums, l, r)
		return
	}
	mid := (l + r) / 2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	if nums[mid] > nums[mid+1] {
		merge(nums, l, mid, r)
	}
}

func merge(nums []int, l, mid, r int) {
	aux := make([]int, (r-l)+1)
	for i := l; i <= r; i++ {
		aux[i-l] = nums[i]
	}
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			nums[k] = aux[j-l]
			j++
		} else if j > r {
			nums[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			nums[k] = aux[i-l]
			i++
		} else {
			nums[k] = aux[j-l]
			j++
		}
	}
}
