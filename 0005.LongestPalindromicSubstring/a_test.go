package L5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome_abbcdcb(t *testing.T) {
	res := longestPalindrome("abbcdcb")
	assert.Equal(t, "bcdcb", res)
}

func Test_longestPalindrome_ab(t *testing.T) {
	res := longestPalindrome("ab")
	assert.Equal(t, "a", res)
}
