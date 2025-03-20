Условия в Go представлены привычной конструкцией `if else`. В условии должно быть строго выражение логического типа. Следующий пример вернет ошибку компиляции:

```go
if "hi" { // non-bool "hi" (type string) used as if condition
    // какой-то код
}
```

Корректный пример:

```go
package main

import (
    "fmt"
    "strings"
)

func statusByName(name string) string {
    // функция проверяет, что строка name начинается с подстроки "Mr."
    if strings.HasPrefix(name, "Mr.") {
        return "married man"
    } else if strings.HasPrefix(name, "Mrs.") {
        return "married woman"
    } else {
        return "single person"
    }
}

func main() {
    n := "Mr. Doe"
    fmt.Printf("%s is a %s\n", n, statusByName(n))
    // => Mr. Doe is a married man

    n = "Mrs. Berry"
    fmt.Printf("%s is a %s\n", n, statusByName(n))
    // => Mrs. Berry is a married woman

    n = "Karl"
    fmt.Printf("%s is a %s\n", n, statusByName(n))
    // => Karl is a single person
}
```

Логическое выражение пишется после `if` без скобок. `else if` можно написать только раздельно.
