package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatWord(t *testing.T) {
	a := assert.New(t)

	a.Equal("go", RepeatWord("go", 1))
	a.Equal("gogogo", RepeatWord("go", 3))
	a.Equal("aaaaa", RepeatWord("a", 5))
	a.Equal("", RepeatWord("any", 0)) // дополнительный тест: повторить 0 раз
}
