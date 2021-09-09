package solution_test

import (
	"testing"

	solution "exercises-go/modules/10-basics/60-structs"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	a := assert.New(t)
	a.Equal("invalid request", solution.Validate(solution.UserCreateRequest{
		FirstName: "",
		Age:       0,
	}))
	a.Equal("invalid request", solution.Validate(solution.UserCreateRequest{
		FirstName: "John",
		Age:       -5,
	}))
	a.Equal("invalid request", solution.Validate(solution.UserCreateRequest{
		FirstName: "Andy",
		Age:       0,
	}))
	a.Equal("invalid request", solution.Validate(solution.UserCreateRequest{
		FirstName: "Karl",
		Age:       151,
	}))
	a.Equal("invalid request", solution.Validate(solution.UserCreateRequest{
		FirstName: "",
		Age:       5,
	}))
	a.Equal("invalid request", solution.Validate(solution.UserCreateRequest{
		FirstName: " Henry",
		Age:       15,
	}))
	a.Equal("", solution.Validate(solution.UserCreateRequest{
		FirstName: "John",
		Age:       150,
	}))
	a.Equal("", solution.Validate(solution.UserCreateRequest{
		FirstName: "Susan",
		Age:       30,
	}))
}
