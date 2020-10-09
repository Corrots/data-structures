package numarray

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{}
	}
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 1; i < len(sum); i++ {
		// sum[i]为前i-1个元素的和，= 前i-2个元素的和+第i-1个元素
		sum[i] = sum[i-1] + nums[i-1]
	}
	return NumArray{sum: sum}
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
 * param_1 := obj.SumRange(i,j);
 */
