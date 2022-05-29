package a338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countBits(t *testing.T) {
	res := countBits(0)

	assert.Equal(t, []int{0}, res)
}

func Test_countBits_1(t *testing.T) {
	res := countBits(6)

	assert.Equal(t, []int{0, 1, 1, 2, 1, 2, 2}, res)
}
