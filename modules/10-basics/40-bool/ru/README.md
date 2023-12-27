
Логический тип в Go представлен привычными значениями `true` и `false` c операторами:
— `&&` (и)
— `==` (равно)
— `||` (или)
— `!` (не):

```go
true && false // false
false || true // true
```

Объявление переменных происходит через ключевое слово `bool`:

```go
var b bool = true

// сокращенная запись
bs := false
```

Из-за строгой типизации в Go можно сравнивать только одинаковые типы данных:

```go
true == false // false

false == false // true
```

То есть нельзя сравнить строку с логическим типом, как это происходит в динамических языках:

```go
true == "hello" // invalid operation: false == "hello" (mismatched types untyped bool and untyped string)
```

Когда необходимо проверить на истинность разные значения, нелогические типы нужно привести к логическому:

```go
flag := true
text := "hello"

// вариант не сработает, потому что нельзя конвертировать строку в bool
flag && bool(text) // cannot convert text (type string) to type bool

// правильный вариант: если строка не пустая, то в ней есть текст
flag && text != "" // true
```
