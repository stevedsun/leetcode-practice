package a136

func singleNumber(nums []int) int {
	ans := 0
	for _, i := range nums {
		ans = i ^ ans
	}

	return ans
}
