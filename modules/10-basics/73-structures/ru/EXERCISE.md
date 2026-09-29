Создайте структуру `Package`, которая описывает посылку и содержит два поля:

- `ID` — строка: идентификатор посылки (например, "PKG-1042")
- `Delivered` — логическое значение (bool), указывающее доставлена ли посылка

Реализуйте метод `Status() string` структуры `Package`, который возвращает статус в виде строки, например:

- Если `Delivered == true`, метод возвращает: "Package PKG-1042 has been delivered"
- Если `Delivered == false`, метод возвращает: "Package PKG-1042 is still in transit"

Вместо `PKG-1042` в строку подставляется `ID` посылки. Метод возвращает строку и ничего не печатает.

```go
p1 := Package{ID: "PKG-1042", Delivered: true}
p1.Status() // "Package PKG-1042 has been delivered"

p2 := Package{ID: "PKG-2048", Delivered: false}
p2.Status() // "Package PKG-2048 is still in transit"
```

## Подсказки

- Строку можно собрать функцией `fmt.Sprintf()`, пакет `fmt` в файле уже подключён. Неиспользованный импорт Go считает ошибкой (`imported and not used`), поэтому пакет, который не понадобился, убирают из `import`.
