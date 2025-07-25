Чтобы пройтись по всем элементам карты, в Go используется цикл `for range`.

## Полный перебор ключей и значений

Обычно при работе с картой нужны и ключ, и значение. Для этого используются две переменные:

```go
ages := map[string]int{
	"Alice": 25,
	"Bob":   30,
	"Tom":   19,
}

for name, age := range ages {
	fmt.Printf("%s is %d years old\n", name, age)
}
```

Пример вывода (порядок может быть любым):

```text
Alice is 25 years old
Tom is 19 years old
Bob is 30 years old
```

В отличие от срезов, порядок обхода элементов карты **не гарантируется** — он может быть разным при каждом запуске программы. Это связано с внутренней реализацией `map` и важно учитывать, если требуется получить данные в определённом порядке. Для этого ключи карты можно извлечь в срез, отсортировать и обойти уже его.

<!-- FIXME: как будто бы текст про порядок в начало. или даже в урок где мы рассказываем про мапы. Типа "отличие мап от слайсов" -->

## Перебор только ключей

Если нужно обойти только ключи, можно использовать одну переменную:

```go
ages := map[string]int{
	"Alice": 25,
	"Bob":   30,
	"Tom":   19,
}

for name := range ages {
	fmt.Println(name)
}

// => Alice
// => Bob
// => Tom

```

## Перебор только значений

Если нужен только список значений, ключ можно пропустить, заменив его на `_`:

```go
ages := map[string]int{
	"Alice": 25,
	"Bob":   30,
	"Tom":   19,
}

for _, age := range ages {
	fmt.Println(age)
}

// => 25
// => 30
// => 19
```

## Проверка содержимого при переборе

Иногда нужно не просто вывести все элементы, а проверить условие для каждого из них.
Например, найти первого человека старше 20 лет:

```go
ages := map[string]int{
	"Alice": 25,
	"Bob":   30,
	"Tom":   19,
}

for name, age := range ages {
	if age > 20 {
		fmt.Printf("%s is older than 20\n", name)
		break
	}
}

// => Alice is older than 20
```

## Пример с сортировкой ключей

Чтобы получить элементы карты в определённом порядке, нужно отдельно собрать ключи, отсортировать их и пройтись по ним:

```go
import "slices"

ages := map[string]int{
	"Alice": 25,
	"Bob":   30,
	"Tom":   19,
}

keys := []string{}
for name := range ages {
	keys = append(keys, name)
}
slices.Sort(keys)

for _, name := range keys {
	fmt.Printf("%s is %d years old\n", name, ages[name])
}

// => Alice is 25 years old
// => Bob is 30 years old
// => Tom is 19 years old
```

## Преобразование карты в срез пар

Иногда бывает полезно сконвертировать карту в срез структур для дальнейшей обработки (например, сортировки по значениям):

```go
import "slices"

type Person struct {
	Name string
	Age  int
}

var people []Person
for name, age := range ages {
	people = append(people, Person{Name: name, Age: age})
}

// Теперь people можно отсортировать по возрасту:
slices.SortFunc(people, func(a, b Person) int {
	return a.Age - b.Age
})
```

<!-- FIXME: структуры + слайсы не рассматривались.   -->
