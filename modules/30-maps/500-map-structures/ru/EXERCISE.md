Реализуйте функцию `UpdateEmail(users map[int]*User, id int, newEmail string) error`, которая обновляет email пользователя с указанным `id` в карте `users`.

- Если пользователь с таким `id` существует, функция должна изменить его `Email` на `newEmail` и вернуть `nil`.
- Если пользователь не найден, функция должна вернуть ошибку с сообщением `"user not found"`.

**Пример использования:**

```go
users := map[int]*User{
	1: {Name: "Alice", Email: "alice@example.com"},
	2: {Name: "Bob", Email: "bob@example.com"},
}

err := UpdateEmail(users, 1, "alice@newmail.com")
fmt.Println(users[1].Email) // "alice@newmail.com"
fmt.Println(err)            // <nil>

err = UpdateEmail(users, 3, "charlie@mail.com")
fmt.Println(err)            // "user not found"
```
