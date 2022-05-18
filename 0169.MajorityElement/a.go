package a169

func majorityElement(nums []int) int {
	index := 0
	count := 0

	for index < len(nums) {
		if index == len(nums)-1 {
			break
		}

		number := nums[index]
		count = 1

		for i := index + 1; i < len(nums); i++ {
			if count == 0 {
				index = i
				break
			}

			if nums[i] == number {
				count++
			} else {
				count--
			}

			if i == len(nums)-1 && count != 0 {
				return nums[index]
			}
		}
	}

	return nums[index]
}
