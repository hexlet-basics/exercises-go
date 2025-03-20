На веб-сайтах часто используются разные поддомены для языков. Например, сайт *site.com* на английском располагается по адресу *en.site.com*, а на русском — *ru.site.com*.

Реализуйте функцию `DomainForLocale(domain, locale string) string`, которая добавляет язык `locale` как поддомен к домену `domain`. Язык может прийти пустым, тогда нужно добавить поддомен *en.*. Например:

```go
DomainForLocale("site.com", "")   // en.site.com
DomainForLocale("site.com", "ru") // ru.site.com
```
