package main

import (
	"fmt"
	"sort"
)

func main() {
	weights1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	weights2 := []int{3, 2, 2, 4, 1, 4}
	weights3 := []int{1, 2, 3, 1, 1}
	fmt.Println(shipWithinDays(weights1, 5))
	fmt.Println(shipWithinDays(weights2, 3))
	fmt.Println(shipWithinDays(weights3, 4))

	//fmt.Println(days(weights1, 14))
	//fmt.Println(days(weights1, 15))
}

// https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/
func shipWithinDays(weights []int, D int) int {
	w := make([]int, len(weights))
	copy(w, weights)
	sort.Ints(w)
	max := w[len(w)-1]
	l, r := 1, sum(weights)
	for l < r {
		//fmt.Printf("l: %d, r: %d\n", l, r)
		mid := (r-l)/2 + l
		if days(weights, mid) <= D && mid >= max {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// 当载重量为k时，运送完所有货物所需的天数(D)
func days(weights []int, k int) int {
	sum, d := 0, 1
	for i := 0; i < len(weights); i++ {
		if sum+weights[i] <= k {
			sum += weights[i]
		} else {
			d++
			sum = weights[i]
		}
	}
	return d
}

func sum(weights []int) int {
	var sum int
	for i := range weights {
		sum += weights[i]
	}
	return sum
}
