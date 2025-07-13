package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueUserIDs(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int64{}, UniqueUserIDs([]int64{}))
	a.Equal([]int64{10}, UniqueUserIDs([]int64{10}))
	a.Equal([]int64{55}, UniqueUserIDs([]int64{55, 55}))
	a.Equal([]int64{55, 33, 22}, UniqueUserIDs([]int64{55, 55, 33, 22}))
	a.Equal([]int64{55, 2, 88, 33, 103}, UniqueUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}))
}
