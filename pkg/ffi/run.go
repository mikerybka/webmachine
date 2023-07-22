package ffi

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func Run(input io.Reader, name string, args ...string) (status int, body []byte, err error) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = input
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err = cmd.Run()
	stdoutbytes, _ := io.ReadAll(stdout)
	stderrbytes, _ := io.ReadAll(stderr)
	if err != nil {
		status = http.StatusInternalServerError
		if len(stdoutbytes) > 0 {
			err = fmt.Errorf("%s\nstdout:\n%s", err, stdoutbytes)
		}
		if len(stderrbytes) > 0 {
			err = fmt.Errorf("%s\nstderr:\n%s", err, stderrbytes)
		}
		return status, nil, err
	}

	exitCode := cmd.ProcessState.ExitCode()
	fmt.Println(exitCode)
	if exitCode == 0 {
		status = 200
	} else if exitCode >= 100 && exitCode < 600 {
		status = exitCode
	} else {
		status = http.StatusInternalServerError
	}

	body = stdoutbytes

	return status, body, nil
}
