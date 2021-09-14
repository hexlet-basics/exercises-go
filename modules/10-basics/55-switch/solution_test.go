package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModifySpaces(t *testing.T) {
	a := assert.New(t)
	a.Equal("-hello-world-", ModifySpaces(" hello world ", "dash"))
	a.Equal("_how_r_u_doing_", ModifySpaces(" how r u doing ", "underscore"))
	a.Equal("*great*", ModifySpaces(" great ", ""))
	a.Equal("*nice*", ModifySpaces(" nice ", "unknwn"))
}
