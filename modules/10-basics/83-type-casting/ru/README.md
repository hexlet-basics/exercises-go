В языке Go преобразование типов всегда выполняется **явно**. Это означает, что нельзя просто сложить `int` и `float64`, или передать `int` в функцию, ожидающую `string`. Явное преобразование требуется везде, где участвуют разные типы.

## Преобразование между числовыми типами

```go
var a int = 10
var b float64 = float64(a)
```

Преобразование `int` в `float64`. Без него компилятор выдаст ошибку:

```go
var a int = 10
var b float64 = a // ошибка: нельзя неявно преобразовать int в float64
```

## Преобразование `float64` в `int`

```go
var f float64 = 3.9
var i int = int(f)
fmt.Println(i) // 3
```

Преобразование отбрасывает дробную часть.

## Преобразование `int` в `string`

Чтобы получить строковое представление числа, используется `strconv.Itoa`:

```go
import "strconv"

var n int = 65
fmt.Println(strconv.Itoa(n)) // "65"
```

## Преобразование `string` в `int`

```go
import "strconv"

var s string = "42"
n, err := strconv.Atoi(s)
fmt.Println(n) // 42
```

Строку нельзя просто привести к числу. Используется `strconv.Atoi`, потому что строка может содержать любое значение, не обязательно число.

## Преобразование `bool`

`bool` не может быть приведён ни к числу, ни к строке напрямую.

```go
var flag bool = true
// var x int = int(flag) // ошибка: нельзя преобразовать bool в int
// var s string = string(flag) // ошибка: нельзя преобразовать bool в string
```

Для строк используется `fmt.Sprintf`:

```go
import "fmt"

var flag bool = true
s := fmt.Sprintf("%v", flag)
fmt.Println(s) // "true"
```

## Ошибки при несовместимых типах

Go не делает неявных преобразований между несовместимыми типами:

```go
var x int = 10
var y uint = uint(x) // допустимо: явное преобразование

var z string = string(x)
fmt.Println(z) // символ с кодом 10, а не "10"
```
