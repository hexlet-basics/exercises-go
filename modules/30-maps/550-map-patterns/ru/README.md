# Полезные приёмы работы с картами

Рассмотрим наиболее полезные приёмы (паттерны) при работе с картами. В некоторых случаях мы будем использовать библиотеку [samber/lo](https://github.com/samber/lo), которая предоставляет готовые функции для работы с коллекциями и картами.

## Инкремент значений (счётчики)

Карты отлично подходят для подсчёта чего-либо. Если ключ отсутствует, доступ `map[key]` возвращает нулевое значение:

```go
words := map[string]int{}
words["go"]++
words["go"]++

fmt.Println(words["go"]) // => 2
```

## Установка значения по умолчанию

С помощью `lo.GetOrElse()` можно задать значение по умолчанию, если ключ отсутствует:

```go
import "github.com/samber/lo"

defaults := map[string]string{
	"theme": "light",
}

lang := lo.GetOrElse(defaults, "lang", "en")
fmt.Println(lang) // => en
```

## Сбор ключей или значений

Чтобы получить список всех ключей карты, используйте `maps.Keys()`:

```go
import "maps"

keys := maps.Keys(users)
fmt.Println(keys) // => [alice bob]
```

Чтобы собрать значения:

```go
import (
  "fmt"
  "maps"
)

values := maps.Values(users)
fmt.Println(values) // => [25 30]
```

Если нужно отсортировать ключи:

```go
import (
  "fmt"
  "slices"
  "maps"
)

keys := maps.Keys(users)
slices.Sort(keys)
fmt.Println(keys) // => [alice bob]
```

## Вложенные карты (инициализация на лету)

Когда значение карты само является картой, полезен следующий паттерн:

```go
settings := map[string]map[string]string{}

user := "alice"
if settings[user] == nil {
	settings[user] = make(map[string]string)
}
settings[user]["theme"] = "dark"
```

При первом обращении к `settings["alice"]`, если такого ключа ещё нет, Go вернёт `nil`, но не ошибку.
Однако если мы попытаемся сразу обратиться к `settings["alice"]["theme"]`, не проверяя, инициализирована ли внутренняя карта, будет паника времени выполнения.

Этот шаблон — инициализация на лету: если для ключа ещё нет вложенной карты, она создаётся сразу перед использованием.
