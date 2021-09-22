package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinInt(t *testing.T) {
	a := assert.New(t)
	a.Equal(5, MinInt(5, 20))
	a.Equal(-30, MinInt(-30, 30))
	a.Equal(0, MinInt(0, 0))
	a.Equal(0, MinInt(2, 0))
}
