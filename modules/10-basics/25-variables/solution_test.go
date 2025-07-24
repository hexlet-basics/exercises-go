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
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	os.Stdout = w

	// Вызываем тестируемую функцию
	PrintBody()

	// Закрываем запись и читаем вывод
	err = w.Close()
	assert.NoError(t, err)
	os.Stdout = old

	_, err = buf.ReadFrom(r)
	assert.NoError(t, err)

	expected := "Go Go 1.21\n"
	assert.Equal(t, expected, buf.String(), "Вывод функции PrintBody не совпадает с ожидаемым")
}
