Реализуйте функцию `AreOrderHistoriesEqual(history1, history2 [][]string) bool`, которая проверяет, совпадают ли два списка историй заказов.  

Каждая история — это список товаров (срез строк). Истории считаются равными, если количество заказов одинаково и каждый заказ (список товаров) содержит те же товары в том же порядке.

**Примеры**

```go
history1 := [][]string{
	{"milk", "bread"},
	{"apple", "banana"},
}
history2 := [][]string{
	{"milk", "bread"},
	{"apple", "banana"},
}
AreOrderHistoriesEqual(history1, history2) // true

AreOrderHistoriesEqual(history1, [][]string{
	{"milk", "bread", "butter"},
	{"apple", "banana"},
}) // false
AreOrderHistoriesEqual(nil, [][]string{})  // false
```
