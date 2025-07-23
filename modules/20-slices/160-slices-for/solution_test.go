package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterExpensiveOrders(t *testing.T) {
	tests := []struct {
		name     string
		orders   []int
		limit    int
		expected []int
	}{
		{
			name:     "Orders above limit",
			orders:   []int{120, 35, 70, 400, 15, 220, 90},
			limit:    100,
			expected: []int{120, 400, 220},
		},
		{
			name:     "No orders above limit",
			orders:   []int{10, 20, 30},
			limit:    100,
			expected: []int{},
		},
		{
			name:     "Empty orders list",
			orders:   []int{},
			limit:    100,
			expected: []int{},
		},
		{
			name:     "All orders above limit",
			orders:   []int{200, 300, 400},
			limit:    100,
			expected: []int{200, 300, 400},
		},
		{
			name:     "Orders equal to limit are excluded",
			orders:   []int{50, 100, 150},
			limit:    100,
			expected: []int{150},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterExpensiveOrders(tt.orders, tt.limit)
			assert.Equal(t, tt.expected, result)
		})
	}
}
