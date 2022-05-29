package a338

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i == 0 {
			res[i] = 0
			continue
		}

		if i == 1 {
			res[i] = 1
			continue
		}

		if i%2 == 1 {
			res[i] = res[i-1] + 1
		}

		if i%2 == 0 {
			res[i] = res[i/2]
		}
	}

	return res
}
