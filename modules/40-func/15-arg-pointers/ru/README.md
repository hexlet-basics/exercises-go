
Указатели — очень обширная и непростая тема, выходящая за рамки данного курса. В этом уроке будут рассмотренны только основы передачи указателей на аргументы в функции:

```go
package main

import (
	"fmt"
)

type User struct {
	email    string
	password string
}

// при объявлении указываем,
// что переменная должна быть указателем.
// Для этого ставим звездочку * перед типом данных
func fillUserData(u *User, email string, pass string) {
	u.email = email
	u.password = pass
}

func main() {
	u := User{}

	// передаем указатель с помощью амперсанда
	// & перед переменной
	fillUserData(&u, "test@test.com", "qwerty")

	fmt.Printf("points on func call %+v\n", u)
	// points on func call {email:test@test.com password:qwerty}

	// сразу инициализируем переменную с указателем
	up := &User{}

	fillUserData(up, "test@test.com", "qwerty")

	fmt.Printf("points on init %+v\n", up)
	// points on init {email:test@test.com password:qwerty}
}
```

Мапы по умолчанию передаются с указателем:

```go
package main

import (
	"fmt"
)

func main() {
	m := map[string]int{}

	fillMap(m)

	fmt.Println(m) // map[random:1]
}

func fillMap(m map[string]int) {
	m["random"] = 1
}
```

Разработчики, пришедшие из других языков, часто используют фразы "передача по ссылке" или "ссылка на переменную". Строго говоря, в Go нет ссылок, только указатели:

```go
package main

import "fmt"

func main() {
	a := 1
	b := &a
	c := &b

	fmt.Printf("%p %p %p\n", &a, &b, &c)
	// 0xc000018030 0xc00000e028 0xc00000e030
}
```

В этом примере *b* и *c* содержат разные значения — *b* содержит адрес переменной *a*, *c* содержит адрес переменной *b*. И *b* и *c* хранятся в разных адресах. Из-за этого обновление переменной *b* не изменит *c*. Поэтому если кто-то говорит про ссылки в Go, он имеет в виду указатели.
