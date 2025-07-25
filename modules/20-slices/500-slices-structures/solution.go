package solution

import "strings"

type User struct {
	Name  string
	Email string
}

// BEGIN

func FindUserNames(users []User) []string {
	result := []string{} // Явно создаём пустой срез, а не nil
	for _, u := range users {
		if strings.HasSuffix(strings.TrimSpace(u.Email), "@example.com") {
			result = append(result, u.Name)
		}
	}
	return result
}

// END
