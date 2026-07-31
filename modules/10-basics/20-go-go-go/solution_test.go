package solution

import (
	"bytes"
	"io"
	"os"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunAll(t *testing.T) {
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	stdout := os.Stdout
	os.Stdout = w

	done := make(chan struct{})
	start := time.Now()

	go func() {
		RunAll()
		close(done)
	}()

	var elapsed time.Duration

	select {
	case <-done:
		elapsed = time.Since(start)
	case <-time.After(2 * time.Second):
		os.Stdout = stdout
		t.Fatal("RunAll не завершилась: счётчик WaitGroup должен уменьшаться для каждой горутины")
	}

	assert.NoError(t, w.Close())
	os.Stdout = stdout

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	assert.NoError(t, err)

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	slices.Sort(lines)

	assert.Equal(t, []string{"Go! 0", "Go! 1", "Go! 2"}, lines)
	assert.Less(t, elapsed, 250*time.Millisecond,
		"работа должна выполняться параллельно, а не одна за другой")
}
