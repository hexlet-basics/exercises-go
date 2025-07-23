// Тесты
package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareProductLists(t *testing.T) {
	tests := []struct {
		name          string
		oldList       []string
		newList       []string
		expectedAdded []string
		expectedRem   []string
	}{
		{
			name:          "One added and one removed",
			oldList:       []string{"apple", "banana", "orange"},
			newList:       []string{"banana", "kiwi", "apple"},
			expectedAdded: []string{"kiwi"},
			expectedRem:   []string{"orange"},
		},
		{
			name:          "All new",
			oldList:       []string{},
			newList:       []string{"milk"},
			expectedAdded: []string{"milk"},
			expectedRem:   []string{},
		},
		{
			name:          "All removed",
			oldList:       []string{"bread"},
			newList:       []string{},
			expectedAdded: []string{},
			expectedRem:   []string{"bread"},
		},
		{
			name:          "No changes",
			oldList:       []string{"water"},
			newList:       []string{"water"},
			expectedAdded: []string{},
			expectedRem:   []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			added, removed := CompareProductLists(tt.oldList, tt.newList)
			assert.ElementsMatch(t, tt.expectedAdded, added)
			assert.ElementsMatch(t, tt.expectedRem, removed)
		})
	}
}
