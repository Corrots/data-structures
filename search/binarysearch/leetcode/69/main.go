package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(0))
}

// https://leetcode-cn.com/problems/sqrtx/
// 求i*i <= x时，i的最大值
func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := (r-l+1)/2 + l
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}
