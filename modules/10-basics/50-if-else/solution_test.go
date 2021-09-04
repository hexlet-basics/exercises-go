package solution_test

import (
	"testing"

	solution "exercises-go/modules/10-basics/50-if-else"

	"github.com/stretchr/testify/assert"
)

func TestMinInt(t *testing.T) {
	a := assert.New(t)
	a.Equal("en.code-basics.com", solution.DomainForLocale("code-basics.com", ""))
	a.Equal("ru.code-basics.com", solution.DomainForLocale("code-basics.com", "ru"))
	a.Equal("vn.code-basics.com", solution.DomainForLocale("code-basics.com", "vn"))
}
