В пакете unicode есть [функция `unicode.Is(unicode.Latin, rune)`](https://pkg.go.dev/unicode#Is), которая проверяет, что руна является латинским символом или нет.

Реализуйте функцию `latinLetters(s string) string`, которая возвращает только латинские символы из строки `s`. Например:

```go
latinLetters("привет world!") // "world"
```
