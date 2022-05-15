package a53

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray_1(t *testing.T) {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})

	assert.Equal(t, 6, res)
}

func Test_maxSubArray_2(t *testing.T) {
	res := maxSubArray([]int{1})

	assert.Equal(t, 1, res)
}

func Test_maxSubArray_3(t *testing.T) {
	res := maxSubArray([]int{5, 4, -1, 7, 8})

	assert.Equal(t, 23, res)
}
