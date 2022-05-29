package a0448

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDisappearedNumbers(t *testing.T) {
	res := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})

	assert.Equal(t, []int{5, 6}, res)
}

func Test_findDisappearedNumbers_1(t *testing.T) {
	res := findDisappearedNumbers([]int{1, 1})

	assert.Equal(t, []int{2}, res)
}
