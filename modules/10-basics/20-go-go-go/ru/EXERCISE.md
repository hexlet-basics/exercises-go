Написать интересный код самостоятельно на текущем уровне будет затруднительно, поэтому скопируйте код ниже.

```go
wg := sync.WaitGroup{}

for i := range 3 {
	wg.Add(1)
	go func() {
	fmt.Println("Go! " + strconv.Itoa(i))
	wg.Done()
	}()
}

wg.Wait()
```
