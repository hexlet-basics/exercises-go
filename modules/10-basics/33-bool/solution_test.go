package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEven(t *testing.T) {
	a := assert.New(t)

	a.False(IsEven(5), "IsEven(5) should return false")
	a.True(IsEven(6), "IsEven(6) should return true")
	a.True(IsEven(0), "IsEven(0) should return true")
	a.True(IsEven(-2), "IsEven(-2) should return true")
	a.False(IsEven(-3), "IsEven(-3) should return false")
}
