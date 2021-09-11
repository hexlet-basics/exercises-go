package solution

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGreetings(t *testing.T) {
	a := assert.New(t)
	a.Equal("Привет, Иван!", Greetings(" иван"))
	a.Equal("Привет, Петр!", Greetings("ПЕТР"))
	a.Equal("Привет, Василий!", Greetings("     вАсиЛИЙ       "))
}
