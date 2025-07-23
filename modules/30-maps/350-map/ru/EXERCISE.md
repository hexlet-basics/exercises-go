Реализуйте функцию `GetGrade(scores map[string]int, name string) (string, error)`, которая принимает карту `scores` с именами студентов и их оценками (число от 0 до 100) и имя студента `name`.

Функция должна:

- Вернуть строку вида `"<имя> has <оценка> points"`, если студент найден.
- Вернуть ошибку, если указанного студента нет в карте.

**Примеры**

```go
scores := map[string]int{
	"Tom":  90,
	"Alice": 75,
}

res, err := GetGrade(scores, "Tom")
// res = "Tom has 90 points", err = nil

res, err := GetGrade(scores, "Charlie")
// res = "", err != nil
```
