Создайте функцию `BuildProfile(name string, age int, rating float64) string`, которая возвращает строку профиля пользователя в формате:

```text
Name: <name>, Age: <age>, Rating: <rating>
```

Рейтинг выводится с одним знаком после точки, и `4.73` превращается в `4.7`, а `5` в `5.0`.

## Примеры

```go
fmt.Println(BuildProfile("Alice", 30, 4.73))
// => "Name: Alice, Age: 30, Rating: 4.7"

fmt.Println(BuildProfile("Bob", 25, 5))
// => "Name: Bob, Age: 25, Rating: 5.0"
```

## Подсказки

- Число с плавающей точкой с заданным количеством знаков после точки переводит в строку функция `strconv.FormatFloat()` из урока про строки. Третий аргумент задаёт число знаков.
- В файле подключены пакеты `fmt` и `strconv`. Неиспользованный импорт Go считает ошибкой (`imported and not used`), поэтому пакет, который не понадобился, убирают из `import`.
