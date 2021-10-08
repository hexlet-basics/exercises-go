package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeAndValidateRequest(t *testing.T) {
	a := assert.New(t)

	req, err := DecodeAndValidateRequest(nil)
	a.Error(err)
	a.Equal(CreateUserRequest{}, req)

	req, err = DecodeAndValidateRequest([]byte(""))
	a.Error(err)
	a.Equal(CreateUserRequest{}, req)

	req, err = DecodeAndValidateRequest([]byte("{}"))
	a.Error(err)
	a.Equal(CreateUserRequest{}, req)

	req, err = DecodeAndValidateRequest([]byte("{\"email\":\"\",\"password\":\"test\",\"password_confirmation\":\"test\"}"))
	a.Equal(errEmailRequired, err)
	a.Equal(CreateUserRequest{}, req)

	req, err = DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"\",\"password_confirmation\":\"test\"}"))
	a.Equal(errPasswordRequired, err)
	a.Equal(CreateUserRequest{}, req)

	req, err = DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"\"}"))
	a.Equal(errPasswordConfirmationRequired, err)
	a.Equal(CreateUserRequest{}, req)

	req, err = DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"test2\"}"))
	a.Equal(errPasswordDoesNotMatch, err)
	a.Equal(CreateUserRequest{}, req)

	req, err = DecodeAndValidateRequest([]byte("{\"email\":\"email@test.com\",\"password\":\"passwordtest\",\"password_confirmation\":\"passwordtest\"}"))
	a.NoError(err)
	a.Equal(CreateUserRequest{
		Email:                "email@test.com",
		Password:             "passwordtest",
		PasswordConfirmation: "passwordtest",
	}, req)
}
