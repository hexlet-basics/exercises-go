package solution

import (
	"encoding/json"
	"errors"
)

// CreateUserRequest is a request to create a new user.
type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// validation errors
var (
	errEmailRequired                = errors.New("email is required")
	errPasswordRequired             = errors.New("password is required")
	errPasswordConfirmationRequired = errors.New("password confirmation is required")
	errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

// BEGIN

// DecodeAndValidateRequest decodes the JSON body and validates the user creation request.
func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {
	req := CreateUserRequest{}

	err := json.Unmarshal(requestBody, &req)
	if err != nil {
		return CreateUserRequest{}, err
	}

	err = validateCreateUserRequest(req)
	if err != nil {
		return CreateUserRequest{}, err
	}

	return req, nil
}

func validateCreateUserRequest(req CreateUserRequest) error {
	if req.Email == "" {
		return errEmailRequired
	}

	if req.Password == "" {
		return errPasswordRequired
	}

	if req.PasswordConfirmation == "" {
		return errPasswordConfirmationRequired
	}

	if req.Password != req.PasswordConfirmation {
		return errPasswordDoesNotMatch
	}

	return nil
}

// END
