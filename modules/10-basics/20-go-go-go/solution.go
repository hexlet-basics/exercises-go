package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	// BEGIN
	wg := sync.WaitGroup{}

	for i := range 3 {
		wg.Add(1)
		go func() {
			fmt.Println("Go! " + strconv.Itoa(i))
			wg.Done()
		}()
	}

	wg.Wait()
	// END
}
