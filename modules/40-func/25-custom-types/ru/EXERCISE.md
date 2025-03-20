
Представим, что есть структура *Person*, содержащая возраст человека:

```go
type Person struct {
	Age uint8
}
```

Реализуйте тип `PersonList` (слайс структур *Person*), с методом `(pl PersonList) GetAgePopularity() map[uint8]int`, который возвращает мапу, где ключ — возраст, а значение — кол-во таких возрастов:

```go
pl := PersonList{
	{Age: 18},
	{Age: 44},
	{Age: 18},
}

pl.GetAgePopularity() // map[18:2 44:1]
```
