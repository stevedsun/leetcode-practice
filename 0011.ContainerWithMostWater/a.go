package a11

func maxArea(height []int) int {
	result := 0
	i := 0
	j := len(height) - 1
	for i < j {
		h := 0
		w := j - i
		if height[i] > height[j] {
			h = height[j]
			j--
		} else {
			h = height[i]
			i++
		}
		area := w * h
		if result < area {
			result = area
		}
	}

	return result
}
