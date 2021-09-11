package solution

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	a := assert.New(t)
	a.False(Valid(0, ""))
	a.False(Valid(5, ""))
	a.False(Valid(0, "hey!"))
	a.True(Valid(10, "hey!"))
}
