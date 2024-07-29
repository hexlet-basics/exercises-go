
Реализуйте интерфейс `Voicer` для структур `Cat`, `Cow` и `Dog` так, чтобы при вызове метода `Voice` экземпляр структуры `Cat` возвращал строку "Мяу", экземпляр `Cow` строку "Мууу", а экземпляр `Dog` сообщение `Гав`:

```go
cat := Cat{} 
dog := Dog{}
cow := Cow{}

fmt.Println(cat.Voice()) // Мяу
fmt.Println(dog.Voice()) // Гав
fmt.Println(cow.Voice()) // Мууу
```
