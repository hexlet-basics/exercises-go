package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestGoGoGo(t *testing.T) {
	cmd := exec.Command("go", "run", "solution.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("the program failed to run: %s\n%s", err, out)
	}
	fmt.Println(string(out))

	for _, line := range []string{"Go! 0", "Go! 1", "Go! 2"} {
		if !strings.Contains(string(out), line) {
			t.Errorf("the program should print the line %q", line)
		}
	}
}
