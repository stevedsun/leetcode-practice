package a136

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	res := singleNumber([]int{4, 1, 2, 1, 2})

	assert.Equal(t, 4, res)
}
