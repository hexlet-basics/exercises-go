package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountLanguages(t *testing.T) {
	users := map[string]string{
		"Alice": "Go",
		"Bob":   "Python",
		"Tom":   "Go",
		"Kate":  "Java",
	}

	expected := map[string]int{
		"Go":     2,
		"Python": 1,
		"Java":   1,
	}

	assert.Equal(t, expected, CountLanguages(users))
}

func TestEmptyUsers(t *testing.T) {
	users := map[string]string{}
	expected := map[string]int{}
	assert.Equal(t, expected, CountLanguages(users))
}
