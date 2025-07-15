package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeGreeting(t *testing.T) {
	a := assert.New(t)

	greeter := MakeGreeting("Hello")
	a.Equal("Hello, Hexlet!", greeter("Hexlet"))
	a.Equal("Hello, Go!", greeter("Go"))

	hi := MakeGreeting("Hi")
	a.Equal("Hi, Hexlet!", hi("Hexlet"))
	a.Equal("Hi, Dev!", hi("Dev"))
}

