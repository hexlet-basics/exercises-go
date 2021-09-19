package solution

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	a := assert.New(t)
	a.Equal("invalid request", Validate(UserCreateRequest{
		FirstName: "",
		Age:       0,
	}))
	a.Equal("invalid request", Validate(UserCreateRequest{
		FirstName: "John",
		Age:       -5,
	}))
	a.Equal("invalid request", Validate(UserCreateRequest{
		FirstName: "Andy",
		Age:       0,
	}))
	a.Equal("invalid request", Validate(UserCreateRequest{
		FirstName: "Karl",
		Age:       151,
	}))
	a.Equal("invalid request", Validate(UserCreateRequest{
		FirstName: "",
		Age:       5,
	}))
	a.Equal("invalid request", Validate(UserCreateRequest{
		FirstName: " Henry",
		Age:       15,
	}))
	a.Equal("", Validate(UserCreateRequest{
		FirstName: "John",
		Age:       150,
	}))
	a.Equal("", Validate(UserCreateRequest{
		FirstName: "Susan",
		Age:       30,
	}))
}
