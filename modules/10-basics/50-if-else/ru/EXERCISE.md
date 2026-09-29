На веб-сайтах часто используются разные поддомены для языков. Например, сайт _site.com_ на английском располагается по адресу _en.site.com_, а на русском — _ru.site.com_.

Реализуйте функцию `DomainForLocale(domain, locale string) string`, которая добавляет язык `locale` как поддомен к домену `domain` и возвращает получившийся адрес строкой. Язык может прийти пустым, тогда нужно добавить поддомен _en._. Например:

```go
DomainForLocale("site.com", "")   // "en.site.com"
DomainForLocale("site.com", "fr") // "fr.site.com"
DomainForLocale("site.com", "es") // "es.site.com"
```
