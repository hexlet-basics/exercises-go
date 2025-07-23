package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDiscount(t *testing.T) {
	a := assert.New(t)
	a.Equal([]float64{100, 200, 270}, AddDiscount([]float64{100, 200, 300}, 10))
	a.Equal([]float64{}, AddDiscount([]float64{}, 15))
	a.Equal([]float64{400}, AddDiscount([]float64{500}, 20))
	a.Equal([]float64{50, 90}, AddDiscount([]float64{50, 100}, 10))
}
