package v1

import "log"

type NumArray struct {
	sum  []int
	data []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{}
	}

	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	return NumArray{
		sum:  sum,
		data: nums,
	}
}

func (this *NumArray) Update(i int, val int) {
	if i < 0 || i > len(this.data) {
		log.Fatalf("invalid index %v\n", i)
	}
	this.data[i] = val
	for j := i + 1; j < len(this.sum); j++ {
		this.sum[j] = this.sum[j-1] + this.data[j-1]
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if this == nil {
		return 0
	}
	return this.sum[j+1] - this.sum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
