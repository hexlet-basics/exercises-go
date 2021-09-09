package solution_test

import (
	"testing"

	solution "exercises-go/modules/10-basics/55-switch"

	"github.com/stretchr/testify/assert"
)

func TestModifySpaces(t *testing.T) {
	a := assert.New(t)
	a.Equal("-hello-world-", solution.ModifySpaces(" hello world ", "dash"))
	a.Equal("_how_r_u_doing_", solution.ModifySpaces(" how r u doing ", "underscore"))
	a.Equal("*great*", solution.ModifySpaces(" great ", ""))
	a.Equal("*nice*", solution.ModifySpaces(" nice ", "unknwn"))
}
