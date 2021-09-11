package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainForLocale(t *testing.T) {
	a := assert.New(t)
	a.Equal("en.code-basics.com", DomainForLocale("code-basics.com", ""))
	a.Equal("ru.code-basics.com", DomainForLocale("code-basics.com", "ru"))
	a.Equal("vn.code-basics.com", DomainForLocale("code-basics.com", "vn"))
}
