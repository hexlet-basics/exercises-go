package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHiddenCard(t *testing.T) {
	a := assert.New(t)

	a.Equal("**5678", GetHiddenCard("1234567812345678", 2))
	a.Equal("***5678", GetHiddenCard("1234567812345678", 3))
	a.Equal("*1100", GetHiddenCard("2034399002121100", 1))
	a.Equal("*****9999", GetHiddenCard("1111222233339999", 5))
}
