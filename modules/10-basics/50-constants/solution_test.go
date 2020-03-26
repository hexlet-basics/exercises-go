package main

import (
	"fmt"
	"log"
	"os/exec"
)

func ExampleConstants() {
	cmd := exec.Command("go", "run", "solution.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("error on running the command: %v\n", err)
	}
	fmt.Println(string(out))
	// Output:
	// 20000
}
