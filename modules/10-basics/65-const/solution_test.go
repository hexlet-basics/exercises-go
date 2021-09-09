package solution_test

import (
	"testing"

	solution "exercises-go/modules/10-basics/65-const"

	"github.com/stretchr/testify/assert"
)

func TestErrorMessageToCode(t *testing.T) {
	a := assert.New(t)
	a.Equal(0, solution.ErrorMessageToCode("OK"))
	a.Equal(1, solution.ErrorMessageToCode("CANCELLED"))
	a.Equal(2, solution.ErrorMessageToCode("UNKNOWN"))
	a.Equal(2, solution.ErrorMessageToCode("err"))
}
