package a121

func maxProfit(prices []int) int {
	result := 0
	minVal := prices[0]

	for _, i := range prices {
		if i < minVal {
			minVal = i
		} else {
			profit := i - minVal
			if profit > result {
				result = profit
			}
		}
	}

	return result
}
