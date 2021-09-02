package solution_test

import (
	"testing"

	solution "exercises-go/modules/10-basics/70-bool"

	"github.com/stretchr/testify/assert"
)

func TestMinInt(t *testing.T) {
	a := assert.New(t)
	a.Equal(false, solution.Valid(0, ""))
	a.Equal(false, solution.Valid(5, ""))
	a.Equal(false, solution.Valid(0, "hey!"))
	a.Equal(true, solution.Valid(10, "hey!"))
}
