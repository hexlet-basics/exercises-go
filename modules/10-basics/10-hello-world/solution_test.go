package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	cmd := exec.Command("go", "run", "solution.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("the program failed to run: %s\n%s", err, out)
	}
	fmt.Println(string(out))

	expected := "Hello, World!"
	if actual := strings.TrimSpace(string(out)); actual != expected {
		t.Errorf("the program should print %q, got %q", expected, actual)
	}
}
