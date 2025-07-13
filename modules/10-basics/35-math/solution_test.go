package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateProgress(t *testing.T) {
	a := assert.New(t)

	a.InDelta(0.0, CalculateProgress(0, 10), 0.0001)
	a.InDelta(0.25, CalculateProgress(1, 4), 0.0001)
	a.InDelta(0.4, CalculateProgress(2, 5), 0.0001)
	a.InDelta(0.75, CalculateProgress(3, 4), 0.0001)
	a.InDelta(1.0, CalculateProgress(5, 5), 0.0001)
}
