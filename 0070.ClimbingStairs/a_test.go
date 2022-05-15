package a70

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbingStairs(t *testing.T) {
	res := climbStairs(2)
	assert.Equal(t, 2, res)
}

func Test_climbingStairs_2(t *testing.T) {
	res := climbStairs(44)
	assert.Equal(t, 1134903170, res)
}
