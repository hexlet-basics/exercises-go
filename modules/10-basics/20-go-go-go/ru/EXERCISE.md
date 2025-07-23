Написать интересный код самостоятельно на текущем уровне будет затруднительно, поэтому скопируйте код ниже.

```go
wg := sync.WaitGroup{}

for i := 0; i < 3; i++ {
	wg.Add(1)
	go func() {
	fmt.Println("Go! " + string(i))
		wg.Done()
	}()
}

wg.Wait()
```
