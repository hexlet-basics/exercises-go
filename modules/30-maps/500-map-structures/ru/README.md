Карты (`map`) часто используются для хранения данных, связанных с уникальными ключами. Частый пример — хранение структур (например, пользователей, заказов, товаров), где ключ — это уникальный идентификатор (ID).

## Базовый пример

Предположим, у нас есть структура `User`:

```go
type User struct {
	Name  string
	Email string
}
```

Создадим карту, где ключ — целое число (ID), а значение — структура `User`:

```go
users := map[int]User{
	1: {Name: "Alice", Email: "alice@example.com"},
	2: {Name: "Bob", Email: "bob@example.com"},
}
fmt.Println(users[1]) // => {Alice alice@example.com}
```

## Проблема копий

При получении элемента из карты возвращается **копия значения**, а не ссылка на оригинал:

```go
user := users[1]
user.Name = "Alicia"
fmt.Println(users[1].Name) // => Alice — оригинал не изменился
```

Чтобы изменить данные в карте, нужно явно присвоить изменённую структуру обратно:

```go
user := users[1]
user.Name = "Alicia"
users[1] = user
fmt.Println(users[1].Name) // => Alicia
```

## Использование указателей

Если нужно часто изменять данные, удобнее хранить в карте **указатели на структуры**:

```go
users := map[int]*User{
	1: &User{Name: "Alice", Email: "alice@example.com"},
	2: &User{Name: "Bob", Email: "bob@example.com"},
}

users[1].Name = "Alicia"
fmt.Println(users[1].Name) // Alicia — оригинал изменился
```

В этом случае карта хранит ссылки на объекты, и изменения применяются напрямую.

## Перебор карты со структурами

Перебор элементов ничем не отличается от обычных карт:

```go
type User struct {
	Name  string
	Email string
}

users := map[int]*User{
	1: &User{Name: "Alice", Email: "alice@example.com"},
	2: &User{Name: "Bob", Email: "bob@example.com"},
}

for id, user := range users {
	fmt.Printf("ID=%d, Name=%s\n", id, user.Name)
}

// => ID=1, Name=Alice
// => ID=2, Name=Bob
```

Если карта хранит указатели, цикл будет работать аналогично, только `user` уже будет указателем.
