package a46

import "fmt"

func permute(nums []int) [][]int {
	var result = [][]int{}
	if len(nums) == 1 {
		result = append(result, []int{nums[0]})
		return result
	}
	for i := 0; i < len(nums); i++ {
		temp := make([]int, len(nums))
		copy(temp, nums)
		temp[i] = temp[len(temp)-1]
		temp = temp[:len(temp)-1]
		fmt.Println("nums=", nums, "temp=", temp)
		for _, res := range permute(temp) {
			result = append(result, append([]int{nums[i]}, res...))
		}
	}

	return result
}
