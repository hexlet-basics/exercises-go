В любой программе часто требуется выполнять разные действия в зависимости от условий. Для этого в Go, как и в большинстве языков программирования, используется конструкция `if`.

**Базовый синтаксис**

```go
if условие {
	// код, который выполнится, если условие истинно (true)
}
```

**Пример**

```go
age := 18

if age >= 18 {
	fmt.Println("Access granted")
}

// => Access granted
```

## Условие else

Если нужно выполнить другой код в случае, когда условие ложно (`false`), добавляется блок `else`:

```go
age := 14

if age >= 18 {
	fmt.Println("Access granted")
} else {
	fmt.Println("Access denied")
}

// => Access denied
```

Когда условий больше двух, можно использовать цепочку `else if`:

```go
score := 85

if score >= 90 {
	fmt.Println("Grade: A")
} else if score >= 80 {
	fmt.Println("Grade: B")
} else {
	fmt.Println("Grade: C or lower")
}

// Grade: B
```

## Особенности if в Go

- В Go не ставятся круглые скобки вокруг условия.
- Тело `if` обязательно должно быть в фигурных скобках `{}`, даже если там одна строка.
- Условие должно быть логического типа `bool`. Например, такой код вызовет ошибку компиляции:
    ```go
    if 1 {
        fmt.Println("Ошибка") //  non-boolean condition in if statement
    }
    ```

## Логические выражения в условиях

В условии `if` можно использовать логические операторы (`&&`, `||`, `!`) для составления сложных проверок:

```go
age := 20
isRegistered := true

if age >= 18 && isRegistered {
	fmt.Println("Access granted")
}
// => Access granted
```

## Краткая форма объявления переменной

Go позволяет объявить переменную прямо в условии `if`. Эта переменная будет доступна только внутри блока `if` и его `else`:

```go
if n := len("hello"); n > 3 {
	// n существует
	fmt.Println("String is long enough:", n)
} else {
	// и здесь n существует тоже
	fmt.Println("String too short:", n)
}
// Здесь n уже не существует
// => String is long enough: 5
```

## Отсутствие тернарного оператора

В Go, как и в Python нет тернарного оператора (`?:`), который встречается в других языках. Для выбора значения, используется обычный `if`:

```go
age := 16
status := ""

if age >= 18 {
	status = "adult"
} else {
	status = "minor"
}

fmt.Println(status)

// => minor
```
