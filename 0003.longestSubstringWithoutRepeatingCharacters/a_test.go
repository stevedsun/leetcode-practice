package a3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	res := lengthOfLongestSubstring("aabbcc")
	assert.Equal(t, 2, res)
}
