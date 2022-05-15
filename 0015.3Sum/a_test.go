package a15

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {

	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	resString := fmt.Sprintf("%v", res)
	assert.Equal(t, "[[-1 -1 2] [0 -1 1]]", resString)
}

func Test_threeSum_1(t *testing.T) {

	res := threeSum([]int{0, 0, 0, 0})
	resString := fmt.Sprintf("%v", res)
	assert.Equal(t, "[[0 0 0]]", resString)
}
