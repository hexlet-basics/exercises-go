
Ошибки в Go — это особенность языка, которая позволяет работать с неожиданным поведением кода в явном виде:

```go
import "errors"

func validateName(name string) error {
	if name == "" {
		// errors.New создает новый объект ошибки
		return errors.New("empty name")
	}

	if len([]rune(name)) > 50 {
		return errors.New("a name cannot be more than 50 characters")
	}

	return nil
}
```

Тип `error` является интерфейсом. Интерфейс — это отдельный тип данных в Go, представляющий набор методов. Любая структура реализует интерфейс неявно через структурную типизацию. Структурная типизация (в динамических языках это называют *утиной типизацией*) — это связывание типа с реализацией во время компиляции без явного указания связи в коде:

```go
package main

import (
	"fmt"
)

// объявление интерфейса
type Printer interface {
	Print()
}

// нигде не указано, что User реализует интерфейс Printer
type User struct {
	email string
}

// структура User имеет метод Print, как в интерфейсе Printer. Следовательно, во время компиляции запишется связь между User и Printer
func (u *User) Print() {
	fmt.Println("My email is", u.email)
}

// функция принимает как аргумент интерфейс Printer
func TestPrint(p Printer) {
	p.Print()
}

func main() {
	// в функцию TestPrint передается структура User, и так как она реализует интерфейс Printer, все работает без ошибок
	TestPrint(&User{email: "test@test.com"})
}
```

Интерфейс `error` содержит только один метод *Error*, который возвращает строковое представление ошибки:

```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```

Следовательно, легко можно создавать свои реализации ошибок:

```go
type TimeoutErr struct {
	msg string
}

// структура TimeoutErr реализует интерфейс error и может быть использована как обычная ошибка
func (e *TimeoutErr) Error() string {
	return e.msg
}
```

Следует запомнить, что если функция возвращает ошибку, то она всегда возвращается последним аргументом:

```go
// функция возвращает несколько аргументов, и ошибка возвращается последней
func DoHTTPCall(r Request) (Response, error) {
	...
}
```

Нулевое значение для интерфейса — это пустое значение `nil`. Следовательно, когда код работает верно, возвращается `nil` вместо ошибки.
