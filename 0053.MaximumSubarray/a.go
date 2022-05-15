package a53

func maxSubArray(nums []int) int {
	result := 0
	preSubMax := 0
	for i := range nums {
		if i == 0 {
			result = nums[i]
			preSubMax = nums[i]
			continue
		}

		if preSubMax+nums[i] > nums[i] {
			preSubMax = preSubMax + nums[i]
		} else {
			preSubMax = nums[i]
		}

		if preSubMax > result {
			result = preSubMax
		}
	}
	return result
}
