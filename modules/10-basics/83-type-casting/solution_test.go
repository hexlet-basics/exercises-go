package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildProfile(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		rating   float64
		expected string
	}{
		{"Alice", 30, 4.73, "Name: Alice, Age: 30, Rating: 4.7"},
		{"Bob", 25, 5, "Name: Bob, Age: 25, Rating: 5.0"},
		{"Charlie", 40, 3.123, "Name: Charlie, Age: 40, Rating: 3.1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildProfile(tt.name, tt.age, tt.rating)
			assert.Equal(t, tt.expected, result)
		})
	}
}
