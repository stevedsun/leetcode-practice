package a121

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	res := maxProfit([]int{7, 1, 5, 3, 6, 4})

	assert.Equal(t, 5, res)
}

func Test_maxProfit_2(t *testing.T) {
	res := maxProfit([]int{7, 6, 4, 3, 1})

	assert.Equal(t, 0, res)
}

func Test_maxProfit_3(t *testing.T) {
	res := maxProfit([]int{2, 1, 4})

	assert.Equal(t, 3, res)
}
