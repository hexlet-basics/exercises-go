Создайте структуру `Package`, которая описывает посылку и содержит два поля:

- `ID` — строка: идентификатор посылки (например, "PKG-1042")
- `Delivered` — логическое значение (bool), указывающее доставлена ли посылка

Реализуйте метод `Status()`, который возвращает статус в виде строки:

- Если `Delivered == true`, метод возвращает: "Package PKG-1042 has been delivered"
- Если `Delivered == false`, метод возвращает: "Package PKG-1042 is still in transit"

```go
p1 := Package{ID: "PKG-1042", Delivered: true}
p1.Status() // "Package PKG-1042 has been delivered"

p2 := Package{ID: "PKG-2048", Delivered: false}
p2.Status() // "Package PKG-2048 is still in transit"
```
