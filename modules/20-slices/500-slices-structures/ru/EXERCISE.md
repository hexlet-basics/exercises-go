Создайте функцию `FindUserNames()`, которая принимает срез структур `User` и возвращает **срез имен** всех пользователей, у которых email оканчивается на `@example.com`.

1. Структура `User` уже объявлена:

   ```go
   type User struct {
			Name  string
			Email string
   }
   ```

2. Если пользователей с таким email нет, функция должна вернуть пустой срез (а не `nil`).
3. Порядок имен в возвращаемом срезе должен соответствовать исходному порядку.

**Пример**

```go
users := []User{
	{Name: "Семен", Email: "semen@example.com"},
	{Name: "Юля", Email: "yulia@gmail.com"},
	{Name: "Алексей", Email: "alex@example.com"},
}

names := FindUserNames(users)
// Ожидаемый результат: []string{"Семен", "Алексей"}
```
