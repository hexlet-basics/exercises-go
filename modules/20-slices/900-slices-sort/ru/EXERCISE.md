Реализуйте функцию `SortOrdersByCustomerID(orders []Order) []Order`, которая возвращает новый срез заказов, отсортированных по **ID клиента** (от меньшего к большему). Если у нескольких заказов одинаковый `CustomerID` — сортируйте их по цене (от меньшей к большей).

Структура `Order` уже определена:

```go
type Order struct {
	CustomerID int
	Price      int
}
```

Оригинальный срез изменяться не должен.

**Примеры:**

```go
orders := []Order{
	{CustomerID: 3, Price: 200},
	{CustomerID: 1, Price: 300},
	{CustomerID: 3, Price: 100},
}

result := SortOrdersByCustomerID(orders)
fmt.Println(result)
// [{1 300} {3 100} {3 200}]
fmt.Println(orders)
// [{3 200} {1 300} {3 100}] — оригинал не изменился
```
