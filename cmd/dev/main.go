package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	errors := make(chan error)

	// Start server
	go func() {
		cmd := exec.Command("go", "run", "main.go", "--dir=.", "--port=4000")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			errors <- err
		}
	}()

	// Start dev-proxy
	go func() {
		cmd := exec.Command("go", "run", "cmd/dev-proxy/main.go")
		out, err := cmd.CombinedOutput()
		if err != nil {
			errors <- fmt.Errorf("%s:\n%s\n", err, out)
		}
	}()

	err := <-errors

	if err != nil {
		panic(err)
	}
}
