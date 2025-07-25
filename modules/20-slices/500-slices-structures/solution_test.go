package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindUserNames(t *testing.T) {
	users := []User{
		{Name: "Семен", Email: "semen@example.com"},
		{Name: "Юля", Email: "yulia@gmail.com"},
		{Name: "Алексей", Email: "alex@example.com"},
	}

	expected := []string{"Семен", "Алексей"}
	result := FindUserNames(users)

	assert.Equal(t, expected, result, "Должны быть возвращены только пользователи с email @example.com")
}

func TestFindUserNamesEmpty(t *testing.T) {
	users := []User{
		{Name: "Юля", Email: "yulia@gmail.com"},
		{Name: "Виктор", Email: "victor@mail.ru"},
	}

	expected := []string{}
	result := FindUserNames(users)

	assert.Equal(t, expected, result, "Если нет email с @example.com, должен вернуться пустой срез")
}
