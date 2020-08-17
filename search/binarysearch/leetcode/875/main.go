package main

import (
	"fmt"
	"sort"
)

func main() {
	piles1 := []int{3, 6, 7, 11}
	piles2 := []int{30, 11, 23, 4, 20}
	piles3 := []int{30, 11, 23, 4, 20}
	fmt.Println(minEatingSpeed(piles1, 8))
	fmt.Println(minEatingSpeed(piles2, 5))
	fmt.Println(minEatingSpeed(piles3, 6))
	//fmt.Println(maxPile(piles2))
	//fmt.Println(eatingTime(piles1, 4))
}

// https://leetcode-cn.com/problems/koko-eating-bananas/
func minEatingSpeed(piles []int, H int) int {
	sort.Ints(piles)
	l, r := 1, piles[len(piles)-1]
	for l < r {
		//fmt.Printf("l: %d, r: %d\n", l, r)
		mid := (r-l)/2 + l
		k := eatingTime(piles, mid)
		if k > H {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// 每小时吃k根香蕉，吃完所有香蕉所需的时间(H)
func eatingTime(nums []int, k int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		if nums[i]%k != 0 {
			sum += nums[i]/k + 1
		} else {
			sum += nums[i] / k
		}
	}
	return sum
}

func maxPile(nums []int) int {
	for i := 1; i < len(nums); i++ {
		e := nums[i]
		j := 0
		for j = i; j > 0 && nums[j-1] > e; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = e
	}
	return nums[len(nums)-1]
}
