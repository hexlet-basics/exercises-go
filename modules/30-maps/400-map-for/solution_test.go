package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:     "Simple case",
			input:    "Go Go go!",
			expected: map[string]int{"go": 3},
		},
		{
			name:     "Mixed words",
			input:    "Hello world hello",
			expected: map[string]int{"hello": 2, "world": 1},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:     "With punctuation",
			input:    "Hi, hi. Hi?",
			expected: map[string]int{"hi": 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWords(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
