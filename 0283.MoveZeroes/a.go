package a283

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	t := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && t == -1 {
			continue
		}

		if nums[i] == 0 && t == -1 {
			t = i
		}

		if nums[i] != 0 && t != -1 {
			nums[t], nums[i] = nums[i], nums[t]
			t = t + 1
		}
	}
}
