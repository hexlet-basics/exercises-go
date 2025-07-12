package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDouble(t *testing.T) {
	a := assert.New(t)
	a.Equal(6, Double(3))
	a.Equal(16, Double(8))
	a.Equal(0, Double(0))
	a.Equal(-4, Double(-2))
}
