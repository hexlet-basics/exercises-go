В рамках обучения вы разрабатываете командную утилиту, которая выводит приветствие пользователю и включает в себя информацию о количестве новых сообщений. Реализуйте функцию `BuildGreeting()`, которая принимает два аргумента:

- `name` (тип `string`) — имя пользователя
- `count` (тип `int`) — количество новых сообщений

Функция должна возвращать строку в следующем формате:

```
Hello, <имя>! You have <число> new messages.
```

**Примеры**

```go
BuildGreeting("Ivan", 5) // "Hello, Ivan! You have 5 new messages."
BuildGreeting("Hexlet", 0) // "Hello, Hexlet! You have 0 new messages."
```


