package solution_test

import (
	solution "exercises-go/modules/20-array-slice-map/25-slice-copy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{}, solution.IntsCopy([]int{}, 0))
	a.Equal([]int{}, solution.IntsCopy([]int{1, 2}, 0))
	a.Equal([]int{}, solution.IntsCopy([]int{1, 2}, -1))
	a.Equal([]int{}, solution.IntsCopy([]int{1, 2}, -5))
	a.Equal([]int{1, 2}, solution.IntsCopy([]int{1, 2, 3}, 2))
	a.Equal([]int{1, 2, 3, 4}, solution.IntsCopy([]int{1, 2, 3, 4}, 4))
	a.Equal([]int{1, 2, 3, 4, 5}, solution.IntsCopy([]int{1, 2, 3, 4, 5}, 10))
}
