Реализуйте функцию `MakeGreeting()`, которая принимает на вход строку — приветствие (например, "Hello") и возвращает новую функцию. Эта возвращённая функция должна принимать имя и возвращать строку-приветствие, например: "Hello, Hexlet!".

**Примеры:**

```go
greeter := MakeGreeting("Hello")
greeter("Hexlet") // "Hello, Hexlet!"

hiGreeter := MakeGreeting("Hi")
hiGreeter("Go")   // "Hi, Go!"
```
