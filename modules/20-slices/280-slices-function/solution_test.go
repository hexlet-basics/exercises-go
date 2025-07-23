package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDiscount(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		discount int
		expected []int
	}{
		{
			name:     "No discount",
			prices:   []int{100, 200},
			discount: 0,
			expected: []int{100, 200},
		},
		{
			name:     "10 percent discount",
			prices:   []int{100, 200, 300},
			discount: 10,
			expected: []int{90, 180, 270},
		},
		{
			name:     "100 percent discount",
			prices:   []int{100, 50},
			discount: 100,
			expected: []int{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := append([]int(nil), tt.prices...) // копия для проверки
			result := AddDiscount(tt.prices, tt.discount)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, orig, tt.prices, "original slice should remain unchanged")
		})
	}
}
