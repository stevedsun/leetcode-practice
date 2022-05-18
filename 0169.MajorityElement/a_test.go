package a169

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	res := majorityElement([]int{3, 2, 3})

	assert.Equal(t, 3, res)
}

func Test_majorityElement1(t *testing.T) {
	res := majorityElement([]int{3, 3, 4})

	assert.Equal(t, 3, res)
}

func Test_majorityElement2(t *testing.T) {
	res := majorityElement([]int{6, 5, 5})

	assert.Equal(t, 5, res)
}
