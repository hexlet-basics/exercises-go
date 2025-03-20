Go — это компилируемый строго типизированный язык программирования, разработанный в Google. Язык спроектирован для быстрой разработки высоконагруженных бэкендов. Если вы знакомы с императивными языками (например, C++, PHP, Java), то синтаксис Go будет понятен практически сразу:

```go
import (
	"encoding/json"
	"errors"
	"fmt"
)

type Message struct {
	Sender string `json:"sender"` // ставим тег с описанием JSON поля
	Text   string `json:"text"`
}

// инициализация ошибки через конструктор стандартного пакета errors
var errEmptyMessage = errors.New("empty message")

// возвращаем ошибку в случае неожиданного поведения
func DecodeJSON(rawMsg string) (Message, error) {
	// если нам передали пустую строку, возвращаем ошибку об этом
	if len(rawMsg) == 0 {
		return Message{}, errEmptyMessage
	}

	msg := Message{}

	// декодируем строку в структуру
	err := json.Unmarshal([]byte(rawMsg), &msg)
	if err != nil {
		return Message{}, fmt.Errorf("unmarshal: %w", err)
	}

	return msg, nil
}
```

В Go нет исключений. Вместо них используется встроенный интерфейс `error`. Ошибки возвращаются явно последним аргументом из функции. Поэтому Go-код выглядит как череда вызовов функций и проверок на ошибки:

```go
func main() {
	msg, err := DecodeJSON("")
	if errors.Is(err, errEmptyMessage) {
		// { } empty message
		fmt.Println(msg, err)
	}

	msg, err = DecodeJSON("hello")
	if err != nil {
		// { } unmarshal: invalid character 'h' looking for beginning of value
		fmt.Println(msg, err)
	}

	msg, err = DecodeJSON(`{"sender":"hexlet","text":"Go,Go,Go"}`)
	// {hexlet Go,Go,Go} <nil>
	fmt.Println(msg, err)
}
```

Такой подход может показаться «неизящным» из-за постоянного повторения условного блока `if err != nil`, однако он позволяет увидеть и контролировать все потенциальные ошибки в коде.

Самая сильная сторона Go — простое написание конкурентных программ. Для этого в языке используются легковесные потоки — горутины. Мы разберем эту тему подробно в соответствующем модуле, а пока можем оценить синтаксис программы, которая суммирует 10 значений из разных внешних источников:

```go
import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	sum := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)

		// ставим перед любой функцией слово «go», и она выполняется конкурентно в горутине
		go func() {
			// делаем долгий вызов к стороннему API. Так как каждый вызов происходит в своей горутине, мы делаем 10 вызовов одновременно
			n := externalHTTPNum()

			mu.Lock()
			sum += n
			mu.Unlock()

			wg.Done()
		}()
	}

	// ждем, пока все 10 горутин вернут ответ
	wg.Wait()

	fmt.Println(sum) // 55
}
```

Не стоит расстраиваться, если сейчас что-то непонятно, или кажется сложным. После разбора концепций конкурентного программирования в Go и небольшой практики, вы будете с легкостью писать высокопроизводительный код.
