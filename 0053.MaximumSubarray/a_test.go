package L53

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {
	res := maxSubArray([]int{1, 2, 3})
	assert.Equal(t, 0, res)
}
