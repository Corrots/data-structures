package v2

func ShellSort(nums []int) {
	h := len(nums) / 2
	for h >= 1 {
		// 对nums[h,n)进行插入排序
		for i := h; i < len(nums); i++ {
			tmp := nums[i]
			j := 0
			for j = i; j-h >= 0 && nums[j-h] > tmp; j -= h {
				nums[j] = nums[j-h]
			}
			nums[j] = tmp
		}
		h /= 2
	}
}
