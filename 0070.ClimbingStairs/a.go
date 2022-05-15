package a70

var cache = make(map[int]int)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	var r1 int
	var r2 int
	if cache[n-1] != 0 {
		r1 = cache[n-1]
	} else {
		r1 = climbStairs(n - 1)
		cache[n-1] = r1
	}

	if cache[n-2] != 0 {
		r2 = cache[n-2]
	} else {
		r2 = climbStairs(n - 2)
		cache[n-2] = r2
	}

	return r1 + r2
}
