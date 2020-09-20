package v1

func ShellSort(nums []int) {
	h := len(nums) / 2
	for h >= 1 {
		for start := 0; start < h; start++ {
			for i := start + h; i < len(nums); i += h {
				tmp := nums[i]
				j := 0
				for j = i; j-h >= 0 && nums[j-h] > tmp; j -= h {
					nums[j] = nums[j-h]
				}
				nums[j] = tmp
			}
		}
		h /= 2
	}
}

//func ShellSort(nums []int) {
//	h := len(nums) / 2
//	for h >= 1 {
//		for i := 0; i < len(nums); i += h {
//			tmp := nums[i]
//			j := 0
//			for j = i; j-h >= 0 && nums[j-h] > tmp; j -= h {
//				nums[j] = nums[j-h]
//			}
//			nums[j] = tmp
//		}
//		h /= 2
//	}
//}
