package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyParent(t *testing.T) {
	a := assert.New(t)

	// nil case
	cp := CopyParent(nil)
	a.Equal(Parent{}, cp)

	// filled struct case
	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
			{
				Name: "Vasya",
				Age:  22,
			},
		},
	}

	cp = CopyParent(p)
	cp.Children[0] = Child{}

	a.Equal("Harry", cp.Name)
	a.Equal("Harry", p.Name)
	a.Len(p.Children, 2)
	a.Len(cp.Children, 2)
	a.Equal([]Child{
		{
			Name: "Andy",
			Age:  18,
		},
		{
			Name: "Vasya",
			Age:  22,
		},
	}, p.Children)
	a.Equal([]Child{
		{},
		{
			Name: "Vasya",
			Age:  22,
		},
	}, cp.Children)
}
