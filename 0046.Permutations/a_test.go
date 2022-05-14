package a46

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permute(t *testing.T) {
	res := permute([]int{1, 2, 3})
	s := fmt.Sprintf("%v", res)
	assert.Equal(t, "[[1 3 2] [1 2 3] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]", s)
}
