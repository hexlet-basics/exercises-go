package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{}, IntsCopy([]int{}, 0))
	a.Equal([]int{}, IntsCopy([]int{1, 2}, 0))
	a.Equal([]int{}, IntsCopy([]int{1, 2}, -1))
	a.Equal([]int{}, IntsCopy([]int{1, 2}, -5))
	a.Equal([]int{1, 2}, IntsCopy([]int{1, 2, 3}, 2))
	a.Equal([]int{1, 2, 3, 4}, IntsCopy([]int{1, 2, 3, 4}, 4))
	a.Equal([]int{1, 2, 3, 4, 5}, IntsCopy([]int{1, 2, 3, 4, 5}, 10))
}
