package solution_test

import (
	"testing"

	solution "exercises-go/modules/10-basics/80-strings"

	"github.com/stretchr/testify/assert"
)

func TestMinInt(t *testing.T) {
	a := assert.New(t)
	a.Equal("Привет, Иван!", solution.Greetings(" иван"))
	a.Equal("Привет, Петр!", solution.Greetings("ПЕТР"))
	a.Equal("Привет, Василий!", solution.Greetings("     вАсиЛИЙ       "))
}
