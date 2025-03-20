Реализуйте функцию `ModifySpaces(s, mode string) string`, которая изменяет строку `s` в зависимости от переданного режима `mode`. Подробности в примере:

```go
// когда передается mode "dash", нужно заменить все пробелы на дефисы "-"
ModifySpaces("hello world", "dash") // hello-world

// когда передается mode "underscore", нужно заменить все пробелы на нижние подчеркивания "_"
ModifySpaces("hello world", "underscore") // hello_world

// когда передается неизвестный или пустой mode, заменяем все пробелы на звездочки "*"
ModifySpaces("hello world", "unknown") // hello*world
ModifySpaces("hello world", "")        // hello*world
```

Для замены символов в строке существует функция `ReplaceAll(s, old, new string) string` из пакета `strings`:

```go
strings.ReplaceAll("hello world!", "world!", "buddy!") // hello buddy!

strings.ReplaceAll("one two three", " ", "_") // one_two_three
```
