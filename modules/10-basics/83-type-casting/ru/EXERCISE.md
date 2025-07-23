Создайте функцию `BuildProfile(name string, age int, rating float64) string`,  
которая возвращает строку профиля пользователя в формате:

    "Name: <name>, Age: <age>, Rating: <rating>"

**Примеры**

```go
fmt.Println(BuildProfile("Alice", 30, 4.73))
// "Name: Alice, Age: 30, Rating: 4.7"

fmt.Println(BuildProfile("Bob", 25, 5))
// "Name: Bob, Age: 25, Rating: 5.0"
```
