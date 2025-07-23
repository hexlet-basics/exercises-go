package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateEmail(t *testing.T) {
	users := map[int]*User{
		1: {Name: "Alice", Email: "alice@example.com"},
		2: {Name: "Bob", Email: "bob@example.com"},
	}

	// Тест успешного обновления
	err := UpdateEmail(users, 1, "alice@newmail.com")
	assert.NoError(t, err)
	assert.Equal(t, "alice@newmail.com", users[1].Email)

	// Тест обновления несуществующего пользователя
	err = UpdateEmail(users, 3, "charlie@mail.com")
	assert.EqualError(t, err, "user not found")

	// Проверка, что существующие данные не изменились
	assert.Equal(t, "bob@example.com", users[2].Email)
}
