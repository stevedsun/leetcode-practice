package a11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})

	assert.Equal(t, 49, res)
}
