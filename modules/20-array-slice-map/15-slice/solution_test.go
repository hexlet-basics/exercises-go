package solution

import (
	"sort"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	a := assert.New(t)
	equalSlices(a, []int{1, 2}, Remove([]int{1, 2}, -1))
	equalSlices(a, []int{1, 2}, Remove([]int{1, 2}, -5))
	equalSlices(a, []int{2, 3}, Remove([]int{1, 2, 3}, 0))
	equalSlices(a, []int{1, 2}, Remove([]int{1, 2, 3}, 2))
	equalSlices(a, []int{1, 2, 3}, Remove([]int{1, 2, 3}, 3))
	equalSlices(a, []int{1, 2, 3}, Remove([]int{1, 2, 3}, 5))
}

func equalSlices(a *assert.Assertions, expected, actual []int) {
	sort.Ints(expected)
	sort.Ints(actual)
	a.Equal(expected, actual)
}
