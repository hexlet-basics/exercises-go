package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWorkHours(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Full week",
			input:    []int{8, 8, 8, 8, 6, 0, 0},
			expected: []int{8, 8, 8, 8, 6},
		},
		{
			name:     "Only five days",
			input:    []int{9, 9, 9, 9, 9},
			expected: []int{9, 9, 9, 9, 9},
		},
		{
			name:     "Less than five days",
			input:    []int{8, 8, 8, 8},
			expected: []int{},
		},
		{
			name:     "Empty schedule",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetWorkHours(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
