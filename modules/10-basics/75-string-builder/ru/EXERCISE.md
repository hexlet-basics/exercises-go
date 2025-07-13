Реализуйте функцию `FilterString()`, которая принимает строку и символ, а возвращает новую строку, в которой удалены все вхождения этого символа. Сравнение символов должно быть без учёта регистра — то есть 'A' и 'a' считаются одинаковыми.

```go
FilterString("If I look forward I win", 'i') // "f  look forward  wn"
FilterString("If I look forward I win", 'O') // "If I lk frward I win"
```

Для приведения символов к нижнему регистру используйте функцию `unicode.ToLower()`:

```go
import "unicode"

r := 'A'
unicode.ToLower(r) // 'a'
```
