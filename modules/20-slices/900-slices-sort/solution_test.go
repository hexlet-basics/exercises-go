// Тесты
package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortOrdersByCustomerID(t *testing.T) {
	tests := []struct {
		name     string
		input    []Order
		expected []Order
	}{
		{
			name: "Sort by CustomerID",
			input: []Order{
				{CustomerID: 3, Price: 200},
				{CustomerID: 1, Price: 300},
				{CustomerID: 2, Price: 100},
			},
			expected: []Order{
				{CustomerID: 1, Price: 300},
				{CustomerID: 2, Price: 100},
				{CustomerID: 3, Price: 200},
			},
		},
		{
			name: "Sort by CustomerID and price",
			input: []Order{
				{CustomerID: 2, Price: 300},
				{CustomerID: 2, Price: 100},
				{CustomerID: 1, Price: 500},
			},
			expected: []Order{
				{CustomerID: 1, Price: 500},
				{CustomerID: 2, Price: 100},
				{CustomerID: 2, Price: 300},
			},
		},
		{
			name:     "Empty slice",
			input:    []Order{},
			expected: []Order{},
		},
		{
			name:     "Single element",
			input:    []Order{{CustomerID: 1, Price: 100}},
			expected: []Order{{CustomerID: 1, Price: 100}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SortOrdersByCustomerID(tt.input)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.input, tt.input, "original slice should remain unchanged")
		})
	}
}
