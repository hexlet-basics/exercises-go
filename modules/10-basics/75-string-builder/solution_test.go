package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterString(t *testing.T) {
	a := assert.New(t)

	a.Equal("f  look forward  wn", FilterString("If I look forward I win", 'i'))
	a.Equal("If I lk frward I win", FilterString("If I look forward I win", 'O'))
	a.Equal("Hxlt", FilterString("Hexlet", 'e'))
	a.Equal("Hello", FilterString("Hello", 'x')) // Символ отсутствует
	a.Equal("", FilterString("", 'a'))           // Пустая строка
}
