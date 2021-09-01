package solution_test

import (
	"testing"

	solution "exercises-go/modules/10-basics/60-math"

	"github.com/stretchr/testify/assert"
)

func TestMinInt(t *testing.T) {
	a := assert.New(t)
	a.Equal(5, solution.MinInt(5, 20))
	a.Equal(-30, solution.MinInt(-30, 30))
	a.Equal(0, solution.MinInt(0, 0))
	a.Equal(0, solution.MinInt(2, 0))
}
