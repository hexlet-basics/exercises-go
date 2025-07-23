Реализуйте функцию `CountLanguages(users map[string]string) map[string]int`, которая принимает карту `users`, где ключ — имя пользователя, а значение — язык программирования, и возвращает карту с подсчётом количества пользователей для каждого языка.

**Пример:**

```go
users := map[string]string{
	"Alice": "Go",
	"Bob":   "Python",
	"Tom":   "Go",
	"Kate":  "Java",
}

result := CountLanguages(users)
fmt.Println(result)
// map[Go:2 Java:1 Python:1]
```
