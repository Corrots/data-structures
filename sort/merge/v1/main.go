package main

import (
	"fmt"

	"github.com/corrots/data-structures/sort/helper"
)

func main() {
	opCount := 100000
	t1 := helper.TestSort(opCount, MergeSort)
	fmt.Printf("MergeSort spend %d ms", t1)
}

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

// 对nums[l:r]范围内的元素进行归并排序
func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (r-l)/2 + l
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	merge(nums, l, mid, r)
}

// 将nums[l:mid]和nums[mid+1:r]的数组进行归并排序
func merge(nums []int, l, mid, r int) {
	aux := make([]int, r-l+1)
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
