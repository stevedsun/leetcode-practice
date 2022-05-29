package a0448

func findDisappearedNumbers(nums []int) []int {
	cache := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		cache[nums[i]-1] = 1
	}

	res := []int{}
	for i := 0; i < len(cache); i++ {
		if cache[i] != 1 {
			res = append(res, i+1)
		}
	}

	return res
}
