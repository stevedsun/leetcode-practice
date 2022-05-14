package L5

func longestPalindrome(s string) string {
	result := ""
	left := -1
	right := -1
	for i := 0; i < 2*len(s)-1; i++ {
		r := ""
		if i%2 != 0 {
			left = (i - 1) / 2
			right = (i + 1) / 2
			r = getLongestPalindrome(s, left, right)
		} else {
			left = i / 2
			right = -1
			r = getLongestPalindrome(s, left, right)
		}

		if len(r) > len(result) {
			result = r
		}
	}

	return result
}

func getLongestPalindrome(s string, c1 int, c2 int) string {
	if len(s) == 1 {
		return s
	}

	if c1 < 0 || c2 > len(s)-1 {
		return ""
	}

	left := -1
	right := -1

	if c2 == -1 {
		left = c1
		right = c1
	} else {
		left = c1
		right = c2
	}

	if s[left] != s[right] {
		return ""
	}

	if left-1 < 0 || right+1 > len(s)-1 {
		return s[left : right+1]
	}

	if res := getLongestPalindrome(s, left-1, right+1); res != "" {
		return res
	} else {
		return s[left : right+1]
	}
}
