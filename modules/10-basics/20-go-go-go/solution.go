package main

import (
	"fmt"
	"sync"
)

func main() {
	// BEGIN
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Go! " + string(i))
			wg.Done()
		}()
	}

	wg.Wait()
	// END
}
