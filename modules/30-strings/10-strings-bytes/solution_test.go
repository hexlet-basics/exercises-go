package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextASCII(t *testing.T) {
	a := assert.New(t)
	a.Equal(byte('b'), nextASCII(byte('a')))
	a.Equal(byte('y'), nextASCII(byte('x')))
	a.Equal(byte('6'), nextASCII(byte('5')))
}

func TestPrevASCII(t *testing.T) {
	a := assert.New(t)
	a.Equal(byte('7'), prevASCII(byte('8')))
	a.Equal(byte('c'), prevASCII(byte('d')))
	a.Equal(byte('m'), prevASCII(byte('n')))
}
