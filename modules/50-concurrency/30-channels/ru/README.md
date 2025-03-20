
В Go существует постулат: "Do not communicate by sharing memory; instead, share memory by communicating" (Не общайтесь разделением памяти. Разделяйте память через общение). Для безопасной коммуникации между горутинами используется специальный тип данных: `chan` (канал).

Как слайсы и мапы, каналы инициализируются с помощью функции `make`:

```go
numCh := make(chan int)
```

Чтение и запись в канал происходит через конструкцию `<-`. Стрелка ставится перед, если канал читается и после, если записывается:

```go
numCh := make(chan int)

numCh <- 10 // записали значение в канал

num := <-numCh // прочитали значение из канала и записали в переменную "num"
```

Чтение из канала блокирует текущую горутину, пока не вернется значение:

```go
package main

import (
	"fmt"
)

func main() {
	numCh := make(chan int)

	<-numCh // программа зависнет здесь и будет ошибка: fatal error: all goroutines are asleep - deadlock!

	fmt.Println("program has ended") // эта строка никогда не выведется
}
```

Запись в канал так же блокирует текущую горутину, пока кто-то не прочтет значение.

Каналы также можно использовать для задачи из прошлого урока:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSum([]int{1, 2, 3}, []int{10, 20, 50})) // [10 20 50]
}

// суммирует значения каждого слайса nums и возвращает тот, который имеет наибольшую сумму
func maxSum(nums1, nums2 []int) []int {
	// канал для результата первой суммы
	s1Ch := make(chan int)
	go sumParallel(nums1, s1Ch)

	// канал для результата второй суммы
	s2Ch := make(chan int)
	go sumParallel(nums2, s2Ch)

	// присваиваем результаты в переменные. Здесь программа будет заблокирована, пока не придут результаты из обоих каналов.
	s1, s2 := <-s1Ch, <-s2Ch

	if s1 > s2 {
		return nums1
	}

	return nums2
}

func sumParallel(nums []int, resCh chan int) {
	s := 0
	for _, n := range nums {
		s += n
	}

	// результат суммы передаем в канал
	resCh <- s
}
```

Иногда требуется запустить обработчика в отдельной горутине, который будет выполнять работу на протяжении всего цикла жизни программы. С помощью конструкции `for range` можно читать из канала до того момента, пока он не будет закрыт:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// создаем канал, в который будем отправлять сообщения
	msgCh := make(chan string)

	// вызываем функцию асинхронно в горутине
	go printer(msgCh)

	msgCh <- "hello"
	msgCh <- "concurrent"
	msgCh <- "world"

	// закрываем канал
	close(msgCh)

	// и ждем, пока printer закончит работу
	time.Sleep(100 * time.Millisecond)
}

func printer(msgCh chan string) {
	// читаем из канала, пока он открыт
	for msg := range msgCh {
		fmt.Println(msg)
	}

	fmt.Println("printer has finished")
}
```
