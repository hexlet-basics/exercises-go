package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateProgress(t *testing.T) {
	a := assert.New(t)

	a.InDelta(0.0, CalculateProgress(0, 10), 0.0001, "CalculateProgress(0, 10) should return 0.0")
	a.InDelta(0.25, CalculateProgress(1, 4), 0.0001, "CalculateProgress(1, 4) should return 0.25")
	a.InDelta(0.4, CalculateProgress(2, 5), 0.0001, "CalculateProgress(2, 5) should return 0.4")
	a.InDelta(0.75, CalculateProgress(3, 4), 0.0001, "CalculateProgress(3, 4) should return 0.75")
	a.InDelta(1.0, CalculateProgress(5, 5), 0.0001, "CalculateProgress(5, 5) should return 1.0")
}
