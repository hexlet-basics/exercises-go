Создайте структуру `Product` с полями:

- `Name` — название товара (строка),
- `Price` — цена (целое число).

Реализуйте функцию `NewDiscountedProduct(name string, price int, discount int) *Product`, которая возвращает **указатель** на новый товар с учётом скидки `discount` (в процентах). Скидка меньше нуля считается нулевой, больше ста — стопроцентной.

**Пример**

```go
p := NewDiscountedProduct("Laptop", 1000, 10)
fmt.Println(p.Price) // 900

p = NewDiscountedProduct("Laptop", 1000, -10)
fmt.Println(p.Price) // 1000

p = NewDiscountedProduct("Laptop", 1000, 150)
fmt.Println(p.Price) // 0
```
