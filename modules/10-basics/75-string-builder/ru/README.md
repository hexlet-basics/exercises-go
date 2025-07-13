В Go строки — неизменяемы (immutable). Это значит, что при каждой операции конкатенации создаётся новая строка, а не изменяется существующая:

```go
s := "Hello"
s += " "
s += "world"
```

Код выше создает три строки в памяти, потому что строки в Go не модифицируются напрямую. Если таких операций много (например, при сборке HTML, SQL или при генерации отчёта), это приводит к лишним аллокациям и снижению производительности.

Для таких случаев стандартная библиотека Go предлагает эффективное решение — `strings.Builder`.

`strings.Builder` — это структура из пакета strings, предназначенная для эффективного построения строк. Он накапливает содержимое во внутреннем буфере, а финальную строку можно получить с помощью метода `.String()`.

```go
var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("world")

	result := builder.String()
	fmt.Println(result) // "Hello world"
```

Для работы с рунами используется свой метод:

```go
builder.WriteString("Hi")
builder.WriteByte(' ')
builder.WriteRune('世') // китайский символ
builder.WriteRune('界') // "界"
fmt.Println(builder.String()) // "Hi 世界"
```

## Почему Builder быстрее

Каждый вызов + над строками в Go создает новую строку, копируя содержимое в новую память. Это дорого по производительности при большом числе операций. Builder же использует один буфер и растёт по мере необходимости, что сокращает число выделений памяти и операций копирования.

## Пример: Объединение с разделителем

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "hexlet"

	var bld strings.Builder

	for i, r := range text {
		bld.WriteRune(r)

		// Добавляем запятую, если это не последний символ
		if i < len(text)-1 {
			bld.WriteString(",")
		}
	}

	result := bld.String()
	fmt.Println(result) // h,e,x,l,e,t
}
```
