Напиши функцию FilterString(), которая получает строку и символ. Она должна вернуть новую строку без всех букв, совпадающих с этим символом — неважно, большие они или маленькие (например, 'A' и 'a' считаются одинаковыми).

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
