package v2

// 调整步长序列h
func ShellSort2(nums []int) {
	// 确定h的最大值
	h := 1
	for h < len(nums)/3 {
		h = h*3 + 1
	}
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
		h /= 3
	}
}

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
