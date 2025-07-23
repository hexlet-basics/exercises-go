package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGrade(t *testing.T) {
	scores := map[string]int{
		"Tom":   90,
		"Alice": 75,
	}

	tests := []struct {
		name      string
		student   string
		want      string
		wantError bool
	}{
		{"Existing student", "Tom", "Tom has 90 points", false},
		{"Another student", "Alice", "Alice has 75 points", false},
		{"Non-existing student", "Charlie", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetGrade(scores, tt.student)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}
		})
	}
}
