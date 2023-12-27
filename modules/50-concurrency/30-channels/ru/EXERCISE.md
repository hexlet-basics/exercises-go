
Реализуйте функцию-воркера `SumWorker(numsCh chan []int, sumCh chan int)`, которая суммирует переданные числа из канала *numsCh* и передает результат в канал *sumCh*:

```go
numsCh := make(chan []int)
sumCh := make(chan int)

go SumWorker(numsCh, sumCh)
numsCh <- []int{10, 10, 10}

res := <- sumCh // 30
```
