
Константы — это постоянные значения, которые инициализируются один раз и не изменяются в течение всего выполнения программы. В Go константы объявляются через ключевое слово `const`:

```go
const [название] [тип данных] = [значение]
```

На практике тип данных не указывается, и несколько констант принято объявлять в рамках одного блока `const`:

```go
const (
	StatusOk       = 200
	StatusNotFound = 404
)
```

Только некоторые типы данных можно присвоить константе: строки, символы, числа, логический тип:

```go
package main

type Person struct {
}

func main() {
	// такие константы допустимы
	const (
		num     = 20
		str     = "hey"
		isValid = true
	)

	// нельзя объявить структуру как константу
	const p = Person{} // ошибка компиляции: const initializer Person{} is not a constant
}
```

Регистр первой буквы указывает на публичность/приватность константы:

```go
const (
	// публичная константа, которую можно использовать во внешних пакетах
	StatusOk = 200

	// приватная константа, доступная только в рамках текущего пакета
	statusInternalError = 500
)
```

Константы можно объявлять на уровне функции, либо пакета:

```go
package main

import "fmt"

const defaultStatus = 200

func main() {
	const status = 404

	fmt.Println("default status:", defaultStatus) // default status: 200
	fmt.Println("current status:", status)        // current status: 404
}
```

Для последовательных числовых констант следует использовать идентификатор iota, который присвоит для списка чисел значения от его текущей позиции:

```go
package main

import "fmt"

const (
	zero = iota
	one
	two
	three
)

const (
	a = iota
	b = 42
	c = iota
	d
)

func main() {
	fmt.Println(zero, one, two, three) // 0 1 2 3
	fmt.Println(a, b, c, d)            // 0 42 2 3
}
```
