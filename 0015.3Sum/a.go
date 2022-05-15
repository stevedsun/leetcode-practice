package a15

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	if len(nums) == 3 && nums[0]+nums[1]+nums[2] == 0 {
		return [][]int{nums}
	}

	sort.Ints(nums)
	result := [][]int{}

	if nums[0] > 0 {
		return result
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[i]+nums[right] == 0 {
				result = append(result, []int{nums[left], nums[i], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left]+nums[i]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return result
}
