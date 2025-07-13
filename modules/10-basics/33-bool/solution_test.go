package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEven(t *testing.T) {
	a := assert.New(t)

	a.False(IsEven(5))
	a.True(IsEven(6))
	a.True(IsEven(0))
	a.True(IsEven(-2))
	a.False(IsEven(-3))
}
