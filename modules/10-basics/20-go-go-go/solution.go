package solution

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// work изображает медленную работу: ждёт 100 миллисекунд и печатает номер
func work(n int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Go! " + strconv.Itoa(n))
}

func RunAll() {
	wg := sync.WaitGroup{}

	// BEGIN
	wg.Add(3)

	go func() {
		work(0)
		wg.Done()
	}()

	go func() {
		work(1)
		wg.Done()
	}()

	go func() {
		work(2)
		wg.Done()
	}()
	// END

	wg.Wait()
}
