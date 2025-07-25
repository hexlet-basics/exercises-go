Вложенные карты (`map` внутри `map`) используются, когда нужно хранить данные с **двумя уровнями ключей**. Например, если у нас есть пользователи, и у каждого из них есть набор настроек, можно использовать карту вида: `map[string]map[string]string`.

## Пример: настройки пользователей

Допустим, у нас есть несколько пользователей, и мы хотим хранить для каждого их настройки:

```go
settings := map[string]map[string]string{
	"alice": {
		"theme":"dark",
		"lang": "en",
	},
	"bob": {
		"theme":"light",
		"lang": "fr",
	},
}

fmt.Println(settings["alice"]["theme"]) // => dark
```

В этом примере:

- Ключи верхнего уровня — это имена пользователей (`"alice"`, `"bob"`).
- Значения — карты с настройками (пары `"ключ-настройки" : "значение"`).

## Добавление элементов

Вот так будет выглядеть добавление вложенной карты:

```go
settings["charlie"] = map[string]string{
	"theme": "dark",
	"lang":  "es",
}
```

Чтобы добавить или изменить конкретную настройку у существующего пользователя:

```go
settings["alice"]["lang"] = "ru"
fmt.Println(settings["alice"]) // => map[lang:ru theme:dark]
```

## Инициализация вложенной карты

Если вложенная карта для пользователя ещё не создана, при обращении к ней будет возвращён `nil`.
Перед изменением такой карты её нужно инициализировать:

```go
user := "david"

fmt.Println(settings[user] == nil) // => true

if settings[user] == nil {
	settings[user] = make(map[string]string)
}
settings[user]["theme"] = "light"

fmt.Println(settings[user] == nil) // => false
```

## Удаление элементов

Удаление вложенного элемента карты можно выполнить с `delete()`:

* Удалить одну настройку пользователя:
    ```go
    delete(settings["bob"], "lang")
    ```

* Удалить все настройки пользователя:
    ```go
    delete(settings, "alice")
    ```

## Перебор вложенных карт

Можно пройтись по всем пользователям и их настройкам:

```go
settings := map[string]map[string]string{
	"alice": {
		"theme": "dark",
		"lang":  "en",
	},
	"bob": {
		"theme": "light",
		"lang":  "fr",
	},
}

for user, userSettings := range settings {
	fmt.Printf("User: %s\n", user)
	for key, value := range userSettings {
		fmt.Printf("  %s = %s\n", key, value)
	}
}
```

Пример вывода:

```text
User: alice
  lang = ru
  theme = dark
User: bob
  theme = light
```
