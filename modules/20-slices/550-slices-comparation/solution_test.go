// Тесты
package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreOrderHistoriesEqual(t *testing.T) {
	tests := []struct {
		name     string
		h1       [][]string
		h2       [][]string
		expected bool
	}{
		{
			name: "Equal histories",
			h1: [][]string{
				{"milk", "bread"},
				{"apple", "banana"},
			},
			h2: [][]string{
				{"milk", "bread"},
				{"apple", "banana"},
			},
			expected: true,
		},
		{
			name: "Different orders",
			h1: [][]string{
				{"milk", "bread"},
				{"apple", "banana"},
			},
			h2: [][]string{
				{"milk", "bread", "butter"},
				{"apple", "banana"},
			},
			expected: false,
		},
		{
			name: "Different sequence",
			h1: [][]string{
				{"milk", "bread"},
				{"apple", "banana"},
			},
			h2: [][]string{
				{"apple", "banana"},
				{"milk", "bread"},
			},
			expected: false,
		},
		{
			name:     "Nil vs empty",
			h1:       nil,
			h2:       [][]string{},
			expected: false,
		},
		{
			name:     "Both empty",
			h1:       [][]string{},
			h2:       [][]string{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AreOrderHistoriesEqual(tt.h1, tt.h2)
			assert.Equal(t, tt.expected, result)
		})
	}
}
