package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostPopularWord(t *testing.T) {
	a := assert.New(t)
	a.Equal("hello", MostPopularWord([]string{"hello", "world", "hello"}))
	a.Equal("world", MostPopularWord([]string{"hello", "world", "hello", "world", "world"}))
	a.Equal("one", MostPopularWord([]string{"one", "two", "three", "four", "five"}))
	a.Equal("c", MostPopularWord([]string{"a", "b", "c", "c", "d", "e", "e", "d"}))
	a.Equal("a", MostPopularWord([]string{"a", "b", "b", "a"}))
}
