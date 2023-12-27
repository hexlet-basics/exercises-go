
Реализуйте функцию `shiftASCII(s string, step int) string`, которая принимает на вход состоящую из ASCII символов строку `s` и возвращает новую строку, где каждый символ из входящей строки сдвинут вперед на число `step`. Например:

```go
shiftASCII("abc", 0) // "abc"
shiftASCII("abc1", 1) // "bcd2"
shiftASCII("bcd2", -1) // "abc1"
shiftASCII("hi", 10) // "rs"
```

Если после сдвига код символа выходит за рамки ASCII, то число должно быть взято по модулю 256:

```go
shiftASCII("abc", 256) // "abc"
shiftASCII("abc", -511) // "bcd"
```
