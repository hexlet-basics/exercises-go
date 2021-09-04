package solution

import (
	"strings"
)

// UserCreateRequest is a request to create a new user.
type UserCreateRequest struct {
	FirstName string
	Age       int
}

// BEGIN

var (
	invalidRequest = "invalid request"
)

// Validate validates the UserCreateRequest and returns an error if the model is invalid.
func Validate(req UserCreateRequest) string {
	if req.FirstName == "" || strings.Contains(req.FirstName, " ") {
		return invalidRequest
	}

	if req.Age <= 0 || req.Age > 150 {
		return invalidRequest
	}

	return ""
}

// END
