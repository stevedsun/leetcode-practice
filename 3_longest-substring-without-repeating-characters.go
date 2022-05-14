package L3

import "strings"

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	var res int = 1
	var i int = 0
	var j int = 0
	for {
		if i >= len(s) || j >= len(s) {
			break
		}
		if i >= j {
			j = j + 1
			continue
		}
		if strings.Contains(s[i:j], string(s[j])) {
			if len(s[i:j]) > res {
				res = len(s[i:j])
			}
			i = i + 1
		} else {
			if len(s[i:j+1]) > res {
				res = len(s[i : j+1])
			}
			j = j + 1
		}
	}

	return res
}

func main() {
	lengthOfLongestSubstring("aabbcc")
}
