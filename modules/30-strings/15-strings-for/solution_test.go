package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShiftASCII(t *testing.T) {
	a := assert.New(t)
	a.Equal("abc", shiftASCII("abc", 0))
	a.Equal("bcd2", shiftASCII("abc1", 1))
	a.Equal("abc1", shiftASCII("bcd2", -1))
	a.Equal("rs", shiftASCII("hi", 10))
	a.Equal("abc", shiftASCII("abc", 256))
	a.Equal("bcd", shiftASCII("abc", -511))
}
