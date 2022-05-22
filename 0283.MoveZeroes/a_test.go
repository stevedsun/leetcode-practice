package a283

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)

	assert.Equal(t, []int{1, 3, 12, 0, 0}, a)
}

func Test_moveZeroes_1(t *testing.T) {
	a := []int{0}
	moveZeroes(a)

	assert.Equal(t, []int{0}, a)
}
