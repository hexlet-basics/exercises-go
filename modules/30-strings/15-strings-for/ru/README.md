
Так как строка — это **массив** байт, ее можно обойти с помощью цикла `for`:

```go
package main

import (
	"fmt"
)

func main() {
	s := "hello"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}

}
```

Вывод:

```
h
e
l
l
o
```

Таким способом можно обойти только строки, состоящие из ASCII символов. Если строка содержит мультибайтовые символы, вывод будет некорректен:

```go
package main

import (
	"fmt"
)

func main() {
	s := "привет"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}

}
```

Вывод проверьте сами в [Go Playground](https://play.golang.org/p/-G3ygH0rTIv)
