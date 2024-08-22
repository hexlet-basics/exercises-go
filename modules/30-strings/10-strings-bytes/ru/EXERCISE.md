
Реализуйте функции `nextASCII(b byte) byte` и `prevASCII(b byte) byte`, которые возвращают следующий или предыдущий символ ASCII таблицы соответственно. Например:

```go
string(nextASCII(byte('a'))) // 'b'
string(prevASCII(byte('b'))) // 'a'
```

Допускаем, что в функцию `prevASCII` *не* может прийти нулевой символ, а в функцию `nextASCII` — последний символ ASCII таблицы.
