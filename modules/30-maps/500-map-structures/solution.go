package solution

import "errors"

type User struct {
	Name  string
	Email string
}

// BEGIN

func UpdateEmail(users map[int]*User, id int, newEmail string) error {
	user, ok := users[id]
	if !ok {
		return errors.New("user not found")
	}
	user.Email = newEmail
	return nil
}

// END
