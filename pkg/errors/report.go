package errors

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"
)

func Report(err error) {
	stack := debug.Stack()
	errString := err.Error()
	timestamp := time.Now().UnixNano()
	filename := filepath.Join("/errors", fmt.Sprintf("%d.txt", timestamp))
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "%s\n\n%s", errString, stack)
	if err != nil {
		panic(err)
	}
}
