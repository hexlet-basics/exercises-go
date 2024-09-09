package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFirstElement(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{2, 3, 4}, RemoveFirstElement([]int{1, 2, 3, 4}))
	a.Equal([]int{}, RemoveFirstElement([]int{}))
	a.Equal([]int{-3, -2}, RemoveFirstElement([]int{-4, -3, -2}))
	a.Equal([]int{}, RemoveFirstElement([]int{3}))
}
