package solution

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintFullName(t *testing.T) {
	// Перехватываем вывод функции в буфер
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	stdout := os.Stdout
	os.Stdout = w

	PrintFullName()

	// Закрываем writer и проверяем ошибку
	closeErr := w.Close()
	assert.NoError(t, closeErr)

	// Восстанавливаем Stdout
	os.Stdout = stdout

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	assert.NoError(t, err)

	// Проверяем вывод
	output := buf.String()
	assert.Equal(t, "John Smith\n", output)
}
