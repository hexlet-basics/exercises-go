Go предоставляет набор функций для работы со строками в пакете `strings`. Эти функции не требуют создания новых типов и работают напрямую со значениями типа `string`.

Чтобы использовать их, нужно импортировать пакет:

```go
import "strings"
```

## Приведение к нижнему и верхнему регистру

```go
strings.ToLower("HeLLo") // "hello"
strings.ToUpper("HeLLo") // "HELLO"
```

## Проверка начала и конца строки

```go
strings.HasPrefix("golang", "go")       // true
strings.HasSuffix("version.go", ".go")  // true
```

## Поиск подстроки

```go
strings.Contains("hello world", "world") // true
strings.Index("hello world", "lo")       // 3
strings.Index("hello world", "go")       // -1
```

## Замена подстроки

```go
strings.Replace("foo bar foo", "foo", "baz", 1)  // "baz bar foo"
strings.Replace("foo bar foo", "foo", "baz", -1) // "baz bar baz"
```

Четвёртый аргумент — число замен. `-1` означает заменить все вхождения.

## Разделение строки

```go
strings.Split("a,b,c", ",") // []string{"a", "b", "c"}
```

Если разделителя нет — вернётся срез из одной строки. Срезы изучаются далее по курсу.

## Соединение строк

```go
// Тут тоже срез
parts := []string{"a", "b", "c"}
strings.Join(parts, "-") // "a-b-c"
```

## Повторение строки

```go
strings.Repeat("na", 3) // "nananana"
```

Функция `Repeat` возвращает новую строку, повторяя исходную указанное число раз. Если указать 0 — вернётся пустая строка.

## Удаление пробелов в начале и конце

```go
strings.TrimSpace("  hello\n") // "hello"
```

Удаляются пробелы, табуляция, переносы строк.
