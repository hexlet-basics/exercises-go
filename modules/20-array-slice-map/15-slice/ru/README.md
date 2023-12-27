
На практике не часто сталкиваешься с массивами из-за ограниченной длины при строгой типизации. Вместо этого повсеместно используются слайсы. Слайс — это массив *неопределенной* длины (или динамический массив):

```go
var nums = []int{1,2,3}

nums := []int{1,2,3}
```

Чтение и запись осуществляется точно так же как в массивах:

```go
nums := []int{1,2,3}

nums[2] // 3

nums[0] = 10 // [10, 2, 3]

nums[:2] // [10, 2]
```

В слайсы можно добавлять элементы с помощью встроенной функции `func append(slice []Type, elems ...Type) []Type`, которая возвращает новый слайс с добавленным элементом:

```go
words := []string{"hello"}

words = append(words, "world") // ["hello", "world"]
```

Так как слайс имеет нефиксированную длину, "под капотом" лежит более сложная структура, чем у массива. Помимо самих значений слайс хранит 2 дополнительных свойства: длину массива *len* (длина) и *cap* (вместимость). Благодаря этому возможно инициализировать слайс нужной длины с помощью встроенной функции `func make(t Type, len, cap IntegerType) Type`. Понимание, где лучше использовать какой способ инициализации, приходит с опытом, но для старта рекомендуется использовать `make` везде, где можно:

```go
// len = 5. Массив сразу будет заполнен 5-ю нулевыми значениями
nums := make([]int, 5, 5) // [0, 0, 0, 0, 0]

// len = 0, но cap = 5. Массив будет пустым, однако заполнение слайса через append будет эффективным, потому что в памяти уже выделен массив нужной длины
nums := make([]int, 0, 5) // []
```

Передача слайса как аргумента функции происходит хитро. Длина и вместимость передаются по значению, но массив значений передается по ссылке. Вследствие этого получается неявное поведение: добавленные элементы не сохранятся в исходный слайс, но изменение существующих останется:

```go
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	modifySlice(nums)

	fmt.Println(nums) // [1 2 10 4 5]
}

func modifySlice(nums []int) {
	nums[2] = 10 // элемент будет и в исходном слайсе
	nums = append(nums, 6) // элемент не добавится в исходный слайс
}
```