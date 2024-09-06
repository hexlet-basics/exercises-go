
В Go присутствует единственная альтернатива `if` — конструкция `switch`. Для этой конструкции используется стандартный синтаксис, но логика работы отличается от С-подобных языков. Когда срабатывает условие какого-либо `case`, программа выполняет блок и выходит из конструкции `switch` без необходимости писать `break`:

```go
x := 10

switch x {
	default: // default всегда выполняется последним независимо от расположения в конструкции
		fmt.Println("default case")
	case 10:
		fmt.Println("case 10")
}
```

Output:

```go
case 10
```

Однако при необходимости можно реализовать логику С-подобных языков и «провалиться» в следующий `case`:

```go
x := 10

switch { // выражение отсутствует. Для компилятора выглядит как: switch true
	default:
		fmt.Println("default case")
	case x == 10:
		fmt.Println("equal 10 case")
		fallthrough
	case x < 10:
		fmt.Println("less 10 case")
}
```

Output:

```go
equal 10 case
less or equal 10 case
```
