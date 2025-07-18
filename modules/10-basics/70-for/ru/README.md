В Go есть только одна конструкция для организации циклов — for. Она заменяет собой и `while`, и `do while` из других языков. С её помощью можно реализовать любые виды повторений.

## Классический цикл с счётчиком

Наиболее распространённый способ перебора:

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

**Здесь:**

`i := 0` — инициализация переменной;
`i < 5` — условие выполнения;
`i++` — шаг итерации.

## Цикл с только условием

Такой for работает как while в других языках:

```go
i := 0
for i < 5 {
	fmt.Println(i)
	i++
}
```

## Бесконечный цикл

```go
for {
	fmt.Println("Hello, Hexlet")
}
```

Обычно используется с `break` или `return` внутри тела, чтобы остановить выполнение.

## Повторение действий с помощью range

В Go существует альтернативный вариант цикла для задачи повторения какого-то действия `n` раз. Вместо привычного счётчика `i := 0; i < n; i++`, можно использовать конструкцию `for range`.

```go
for range 3 {
	fmt.Println("Hello!")
}
```

Если понадобится использовать номер итерации, вы можете явно объявить переменную:

```go
for i := range 5 {
	fmt.Println("Итерация номер", i)
}
```

```go
Итерация номер 0
Итерация номер 1
Итерация номер 2
Итерация номер 3
Итерация номер 4
```

Такой цикл начинается с 0 и заканчивается на n - 1 — как в обычном счётчике.

## Перебор строки посимвольно

Цикл for ... range позволяет пройтись по строке символ за символом. Внутри цикла доступны:

- индекс символа
- сам символ в виде rune (Unicode)

```go
text := "go"

for i, r := range text {
	fmt.Printf("Индекс: %d, Символ: %c\n", i, r)
}
```

**Вывод:**

```
Индекс: 0, Символ: g
Индекс: 1, Символ: o
```

**Что делает range?**

- Возвращает индекс текущего элемента (в байтах)
- Возвращает символ в виде rune — это Unicode-код символа

Строка в Go — это последовательность байт, но range работает корректно с UTF-8 и возвращает именно символы (а не байты):

```go
word := "го"

for i, r := range word {
	fmt.Println(i, string(r))
}
```

**Вывод:**

```
0 г
2 о
```

Индексы 0 и 2 — потому что символы г и о занимают по два байта каждый.

Если не нужен индекс, его можно опустить с помощью _:

```go
for _, r := range "hexlet" {
	fmt.Printf("%c\n", r)
}
```

`range` будет особенно полезен при работе с коллекциями (срезами, словарями и т.п.), к которым мы перейдём в следующих уроках.
