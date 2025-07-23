Реализуйте функцию `FilterExpensiveOrders(orders []int, limit int) []int`, которая возвращает новый срез, содержащий только те заказы, сумма которых строго больше значения limit.
Если в функцию передан пустой срез или нет заказов выше лимита, функция должна вернуть пустой срез.

**Примеры**:

```go
orders := []int{120, 35, 70, 400, 15, 220, 90}
FilterExpensiveOrders(orders, 100) // [120 400 220]

FilterExpensiveOrders([]int{}, 100) // []
FilterExpensiveOrders([]int{10, 20, 30}, 100) // []
```
