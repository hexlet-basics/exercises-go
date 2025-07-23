Реализуйте функцию `SetUserSetting(settings map[string]map[string]string, user, key, value string)`, которая добавляет или обновляет настройку `key` для пользователя `user` в карте `settings`.

- Если пользователь отсутствует в `settings`, необходимо создать для него вложенную карту.
- Функция не возвращает значение, но изменяет карту по ссылке.

**Пример использования:**

```go
settings := map[string]map[string]string{
	"alice": {
		"theme": "dark",
		"lang":  "en",
	},
}

SetUserSetting(settings, "alice", "lang", "ru")
SetUserSetting(settings, "bob", "theme", "light")

fmt.Println(settings["alice"]) // map[lang:ru theme:dark]
fmt.Println(settings["bob"])   // map[theme:light]
```
