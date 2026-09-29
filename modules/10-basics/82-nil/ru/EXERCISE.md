Создайте структуру `Config` с полями:

- `Host` — указатель на строку (`*string`),
- `Port` — число (`int`).

Реализуйте функцию `PrintConfig(cfg *Config)`, которая печатает конфигурацию на экран и ничего не возвращает. Функция должна:

1. Проверять, инициализировано ли поле `Host`.
    - Если `Host == nil`, выводить `"Host is not set"`.
    - Если `Host != nil`, выводить значение строки `Host` в виде `Host: <значение>`.

2. Всегда выводить значение поля `Port` в виде `Port: <значение>`.

**Пример использования:**

```go
host := "localhost"
cfg := &Config{Host: &host, Port: 8080}
PrintConfig(cfg)
// Вывод:
// Host: localhost
// Port: 8080

cfg2 := &Config{Port: 3000}
PrintConfig(cfg2)
// Вывод:
// Host is not set
// Port: 3000
```
