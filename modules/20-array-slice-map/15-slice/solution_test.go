package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{4, 3, 2, 1}, Reverse([]int{1, 2, 3, 4}))
	a.Equal([]int{}, Reverse([]int{}))
	a.Equal([]int{-1, -2, -3, -4}, Reverse([]int{-4, -3, -2, -1}))
	a.Equal([]int{3}, Reverse([]int{3}))
}
