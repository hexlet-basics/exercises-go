package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorMessageToCode(t *testing.T) {
	a := assert.New(t)
	a.Equal(0, ErrorMessageToCode("OK"))
	a.Equal(1, ErrorMessageToCode("CANCELLED"))
	a.Equal(2, ErrorMessageToCode("UNKNOWN"))
	a.Equal(2, ErrorMessageToCode("err"))
}
