Создайте структуру `Product` с полями:

- `Name` — название товара (строка),
- `Price` — цена (целое число).

Реализуйте функцию `NewDiscountedProduct(name string, price int, discount int) *Product`,  
которая возвращает **указатель** на новый товар с учётом скидки `discount` (в процентах).

**Пример**

```go
p := NewDiscountedProduct("Laptop", 1000, 10)
fmt.Println(p.Price) // 900
```
