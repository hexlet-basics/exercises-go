package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToString(t *testing.T) {
	a := assert.New(t)
	a.Equal("0", IntToString(0))
	a.Equal("-42", IntToString(-42))
	a.Equal("100500", IntToString(100500))
}
