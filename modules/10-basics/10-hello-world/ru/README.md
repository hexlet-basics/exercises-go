По традиции начнем с написания программы 'Hello, World!'. Эта программа будет выводить на экран текст:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

## Пояснение

Пакеты в Go выполняют роль неймспейсов — логических единиц, объединяющих функции и типы. Каждый файл в Go должен принадлежать какому-то пакету. Если вы хотите, чтобы программа была исполняемой, файл должен принадлежать пакету main — это сигнал компилятору, что здесь содержится точка входа в программу.

Затем мы импортировали пакет `fmt`:

```go
import "fmt"
```

Go устроен минималистично: чтобы использовать функции из стандартной библиотеки или сторонних модулей, их необходимо явно импортировать. Пакет fmt содержит функции форматированного ввода и вывода, включая Print, Println, Printf и другие. Импорт всегда указывается в начале файла, сразу после названия пакета.

Далее идёт определение главной функции:

```go
func main() {
    ...
}
```

`main()` — это точка входа в программу. Она вызывается автоматически при запуске и должна обязательно присутствовать в исполняемом Go-приложении. Функция main не принимает аргументов и ничего не возвращает. Внутри неё описывается логика запуска программы.

Для вывода текста на экран используется функция:

```go
fmt.Print("Hello, World!")
```

Обратите внимание, что fmt — это не объект, а имя пакета. В Go нет классов, и всё устроено проще: функции организованы по пакетам, и к ним обращаются по следующему синтаксису: пакет.функция. Вызов fmt.Print(...) означает: "вызвать функцию Print из пакета fmt".

## Дополнительно

- В конце строки не ставится точка с запятой — компилятор Go автоматически расставляет ; за вас в большинстве случаев.
- Отступы важны: по соглашению в Go используется 1 таб, а не пробелы.
- Если бы вы забыли импортировать fmt, программа не скомпилируется — Go строго требует явного указания зависимостей.
