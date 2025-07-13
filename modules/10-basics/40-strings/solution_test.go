package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildGreeting(t *testing.T) {
	a := assert.New(t)

	a.Equal("Hello, Ivan! You have 5 new messages.", BuildGreeting("Ivan", 5))
	a.Equal("Hello, Hexlet! You have 0 new messages.", BuildGreeting("Hexlet", 0))
	a.Equal("Hello, Anna! You have 3 new messages.", BuildGreeting("Anna", 3))
	a.Equal("Hello, GoUser! You have 100 new messages.", BuildGreeting("GoUser", 100))
}
