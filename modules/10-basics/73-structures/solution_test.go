package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageStatus(t *testing.T) {
	a := assert.New(t)

	p1 := Package{ID: "PKG-1042", Delivered: true}
	a.Equal("Package PKG-1042 has been delivered", p1.Status())

	p2 := Package{ID: "PKG-2048", Delivered: false}
	a.Equal("Package PKG-2048 is still in transit", p2.Status())
}
