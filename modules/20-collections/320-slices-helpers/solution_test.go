package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueSortedUserIDs(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int64{}, UniqueSortedUserIDs([]int64{}))
	a.Equal([]int64{10}, UniqueSortedUserIDs([]int64{10}))
	a.Equal([]int64{55}, UniqueSortedUserIDs([]int64{55, 55}))
	a.Equal([]int64{22, 33, 55}, UniqueSortedUserIDs([]int64{55, 55, 33, 22}))
	a.Equal([]int64{2, 33, 55, 88, 103}, UniqueSortedUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}))
}
