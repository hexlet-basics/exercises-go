package main

import (
	"fmt"
	"log"
	"os/exec"
)

func ExampleGoGoGo() {
	cmd := exec.Command("go", "run", "solution.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("exec command: %s\n", err)
	}
	fmt.Println(string(out))
	// Output:
	// Go!
	// Go!
	// Go!
}
