
Какая-то функция возвращает критичные и некритичные ошибки:

```go
// некритичная ошибка валидации
type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
	return "validation error"
}

// критичные ошибки
var (
	errBadConnection = errors.New("bad connection")
	errBadRequest    = errors.New("bad request")
)
```

Реализуйте функцию `GetErrorMsg(err error) string`, которая возвращает текст ошибки, если она критичная. В случае неизвестной ошибки возвращается строка *unknown error*:

```go
GetErrorMsg(errors.New("bad connection"))
// "bad connection"

GetErrorMsg(errors.New("bad request"))
// "bad request"

GetErrorMsg(nonCriticalError{})
// ""

GetErrorMsg(errors.New("random error"))
// "unknown error"
```
