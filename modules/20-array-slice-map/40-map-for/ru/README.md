
Как и слайс, мапу можно обойти с помощью конструкции `for range`:

```go
idToName := map[int64]string{1: "Alex", 2: "Dan", 3: "George"}

// первый аргумент — ключ, второй — значение
for id, name := range idToName {
	fmt.Println("id: ", id, "name: ", name)
}
```

Вывод:

```go
id:  1 name:  Alex
id:  2 name:  Dan
id:  3 name:  George
```

Стоит учитывать, что порядок ключей в мапе **рандомизирован**:

```go
numExistence := make(map[int]bool, 0)

// записали по порядку числа от 0 до 9
for i := 0; i < 10; i++ {
	numExistence[i] = true
}

// обходим мапу и выводим ключи
for num := range numExistence {
	fmt.Println(num)
}
```

Вывод:

```go
8
1
2
3
6
7
9
0
4
5
```
