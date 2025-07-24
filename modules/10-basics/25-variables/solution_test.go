package solution

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintBody(t *testing.T) {
	// Создаём буфер и перенаправляем os.Stdout
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Вызываем тестируемую функцию
	PrintBody()

	// Закрываем запись и читаем вывод
	w.Close()
	os.Stdout = old
	_, _ = buf.ReadFrom(r)

	expected := "Go Go 1.21\n"
	assert.Equal(t, expected, buf.String(), "Вывод функции PrintBody не совпадает с ожидаемым")
}
