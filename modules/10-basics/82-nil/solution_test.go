package solution

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintConfig(t *testing.T) {
	tests := []struct {
		name     string
		cfg      *Config
		expected string
	}{
		{
			name:     "With host",
			cfg:      &Config{Host: strPtr("localhost"), Port: 8080},
			expected: "Host: localhost\nPort: 8080\n",
		},
		{
			name:     "Without host",
			cfg:      &Config{Port: 3000},
			expected: "Host is not set\nPort: 3000\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаём pipe для перехвата вывода
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			stdout := os.Stdout
			os.Stdout = w

			// Выполняем тестируемую функцию
			PrintConfig(tt.cfg)

			// Закрываем writer и проверяем ошибку
			closeErr := w.Close()
			assert.NoError(t, closeErr)

			// Восстанавливаем Stdout
			os.Stdout = stdout

			// Читаем перехваченный вывод
			var buf bytes.Buffer
			_, err = io.Copy(&buf, r)
			assert.NoError(t, err)

			// Проверка результата
			assert.Equal(t, tt.expected, buf.String())
		})
	}
}

func strPtr(s string) *string {
	return &s
}
