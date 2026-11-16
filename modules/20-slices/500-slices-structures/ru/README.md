В Go часто возникает необходимость хранить набор однотипных структур. Для этого удобно использовать **срезы структур**. Такой подход позволяет работать с коллекцией объектов, каждый из которых имеет одинаковый набор полей.

## Объявление среза структур

Срез структур можно объявить так же, как и срез базовых типов:

```go
type User struct {
	Name  string
	Email string
}

func main() {
	users := []User{
		{Name: "Семен", Email: "semen@example.com"},
		{Name: "Юля", Email: "yulia@example.com"},
	}

	for _, u := range users {
		fmt.Println(u.Name, u.Email)
	}
}
```

## Добавление новых элементов

Элементы добавляются с помощью `append`:

```go
users = append(users, User{Name: "Алексей", Email: "alex@example.com"})
```

## Изменение значений внутри среза

Срез хранит сами структуры. Обращение по индексу дает доступ к элементу среза, и поле меняется прямо в нем.

```go
users[0].Name = "Семен Иванов"
fmt.Println(users[0].Name) // => Семен Иванов
```

Переменная цикла `range` получает копию элемента. Изменение копии срез не затрагивает.

```go
for _, u := range users {
	u.Name = "Аноним"
}
fmt.Println(users[0].Name) // => Семен Иванов
```

Чтобы изменить элементы в цикле, обращайтесь к ним по индексу.

```go
for i := range users {
	users[i].Name = "Аноним"
}
fmt.Println(users[0].Name) // => Аноним
```

## Срез указателей на структуры

Если нам нужно хранить и изменять объекты по ссылке, можно использовать срез указателей:

```go
type User struct {
	Name  string
	Email string
}

func main() {
	users := []*User{
		{Name: "Семен", Email: "semen@example.com"},
		{Name: "Юля", Email: "yulia@example.com"},
	}

	users[0].Name = "Семен Иванов"  // Изменение отразится на оригинальном объекте
}
```

Использование указателей полезно, если структуры большие и их копирование неэффективно.

## Поиск и фильтрация

Для поиска элементов в срезе структур можно использовать обычные циклы:

```go
func FindUserByEmail(users []User, email string) *User {
	for i := range users {
		if users[i].Email == email {
			return &users[i]
		}
	}
	return nil
}
```
