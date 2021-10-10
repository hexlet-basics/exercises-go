package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSum(t *testing.T) {
	a := assert.New(t)
	a.Nil(MaxSum(nil, nil))
	a.Equal([]int{3, 4, 5, 6}, MaxSum([]int{1, 2, 3}, []int{3, 4, 5, 6}))
	a.Equal([]int{1, 2, 3}, MaxSum([]int{1, 2, 3}, []int{3, 2, 1}))
	a.Equal([]int{10, 20}, MaxSum([]int{10, 20}, []int{4, 6}))
}
