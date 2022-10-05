package solution

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TestMap(t *testing.T) {
	a := assert.New(t)
	testMap(a, []string{"John", "Peter", "Fedor"}, []string{"john", "peter", "fedor"}, func(s string) string {
		return cases.Title(language.Russian).String(s)
	})
	testMap(a, []string{"hello", "world"}, []string{"HELLO", "WORLD"}, func(s string) string {
		return strings.ToLower(s)
	})
}

func testMap(a *assert.Assertions, expected, input []string, mapFunc func(s string) string) {
	inputCopy := make([]string, len(input))
	copy(inputCopy, input)

	a.Equal(expected, Map(input, mapFunc))
	a.Equal(inputCopy, input) // check that the initial slice hasn't been modified.
}
