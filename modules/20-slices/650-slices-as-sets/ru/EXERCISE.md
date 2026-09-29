Реализуйте функцию `CompareProductLists(oldList, newList []string) (added, removed []string)`, которая сравнивает два списка товаров и возвращает два среза:

- `added` — товары, появившиеся в `newList`, которых не было в `oldList`.
- `removed` — товары, которые были в `oldList`, но отсутствуют в `newList`.

Порядок элементов в результатах не имеет значения. Библиотека `github.com/samber/lo` из урока подключена к упражнению и импортирована в файле решения, её функции можно использовать в решении. Неиспользованный импорт Go считает ошибкой (`imported and not used`), поэтому без `lo` этот импорт убирают.

**Примеры:**

```go
oldList := []string{"apple", "banana", "orange"}
newList := []string{"banana", "kiwi", "apple"}

added, removed := CompareProductLists(oldList, newList)
// added = ["kiwi"], removed = ["orange"]

added, removed = CompareProductLists([]string{}, []string{"milk"})
// added = ["milk"], removed = []

added, removed = CompareProductLists([]string{"bread"}, []string{})
// added = [], removed = ["bread"]
```
